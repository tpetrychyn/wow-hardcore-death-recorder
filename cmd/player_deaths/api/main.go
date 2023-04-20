package main

import (
	"database/sql"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tpetrychyn/wow-hardcore-recorder/cmd/player_deaths/api/handler"
	"github.com/tpetrychyn/wow-hardcore-recorder/cmd/player_deaths/data"
	"github.com/tpetrychyn/wow-hardcore-recorder/internal/middleware"
	"github.com/tpetrychyn/wow-hardcore-recorder/internal/rate_limit"
	"github.com/tpetrychyn/wow-hardcore-recorder/internal/service"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

const EnvKey = "DOT_ENV_PATH"
const ConfPrefix = "PLAYER_DEATH_API"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Error("failed to initialize logger", zap.Error(err))
		os.Exit(1)
	}

	conf := &service.Config{}
	err = service.LoadConfig(logger, EnvKey, ConfPrefix, conf)
	if err != nil {
		logger.Error("failed to load env", zap.Error(err))
		os.Exit(1)
	}

	sqldb, err := sql.Open("mysql", conf.DSN)
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		os.Exit(1)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	dalProvider := data.NewDALProvider(db)

	router := gin.New()
	router.Use(ginzap.Ginzap(logger, service.TimeFormat, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(middleware.ErrorHandler)
	router.GET("/health", service.HealthCheckHandler)

	rateLimitMiddleware, err := rate_limit.NewInMemoryRateLimitMiddleware(conf.RateLimit)
	router.Use(rateLimitMiddleware)

	playerDeathHandler := &handler.PlayerDeathHandler{
		Dp: dalProvider,
	}
	router.POST("/deaths", playerDeathHandler.Insert)
	router.GET("/deaths/:guid", playerDeathHandler.GetByGuid)

	if os.Getenv("PORT") != "" {
		conf.Port = os.Getenv("PORT")
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), router))
}
