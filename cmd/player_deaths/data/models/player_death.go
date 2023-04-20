package models

import (
	"github.com/uptrace/bun"
	"time"
)

type PlayerDeath struct {
	bun.BaseModel `bun:"table:player_deaths"`
	Sender        string
	SourceID      string
	Time          int `bun:"-" json:"-"`
	RecordedAt    time.Time
	GUID          string
}

type PlayerDeaths []*PlayerDeath
