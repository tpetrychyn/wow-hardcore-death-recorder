package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tpetrychyn/wow-hardcore-recorder/cmd/player_deaths/data"
	"github.com/tpetrychyn/wow-hardcore-recorder/cmd/player_deaths/data/models"
	"github.com/yuin/gluamapper"
	"github.com/yuin/gopher-lua"
	"io"
	"net/http"
	"time"
)

type PlayerDeathHandler struct {
	Dp *data.DALProvider
}

func (h *PlayerDeathHandler) GetByGuid(c *gin.Context) {
	guid := c.Param("guid")
	if guid == "" {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("guid cannot be blank"))
		return
	}

	deaths, err := h.Dp.GetPlayerDeathsByGuid(c, guid)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, deaths)
}

func (h *PlayerDeathHandler) Insert(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// Initialize a new Lua state
	L := lua.NewState()
	defer L.Close()

	if err := L.DoString(string(body)); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	table := L.GetGlobal("Recorded_Deaths")
	arr, ok := table.(*lua.LTable)
	if !ok {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("cannot find Recorded_Deaths prefix"))
		return
	}

	var deaths models.PlayerDeaths
	arr.ForEach(func(key, value lua.LValue) {
		if tbl, ok := value.(*lua.LTable); ok {
			var death *models.PlayerDeath
			if err := gluamapper.Map(tbl, &death); err != nil {
				_ = c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			death.RecordedAt = time.Unix(int64(death.Time), 0)
			deaths = append(deaths, death)
		}
	})

	err = h.Dp.InsertPlayerDeaths(c, deaths)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
