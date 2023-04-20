package data

import "github.com/uptrace/bun"

type DALProvider struct {
	db *bun.DB
}

func NewDALProvider(db *bun.DB) *DALProvider {
	return &DALProvider{db: db}
}
