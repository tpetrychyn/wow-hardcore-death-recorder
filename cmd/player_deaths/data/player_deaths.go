package data

import (
	"context"
	"github.com/pkg/errors"
	"github.com/tpetrychyn/wow-hardcore-recorder/cmd/player_deaths/data/models"
)

func (dp *DALProvider) GetPlayerDeathsByGuid(ctx context.Context, guid string) (*models.PlayerDeaths, error) {
	deaths := new(models.PlayerDeaths)
	err := dp.db.NewSelect().
		Model(deaths).
		Where("guid = ?", guid).
		Scan(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch player deaths")
	}
	return deaths, nil
}

func (dp *DALProvider) InsertPlayerDeaths(ctx context.Context, deaths models.PlayerDeaths) error {
	_, err := dp.db.NewInsert().
		Model(&deaths).
		Ignore().
		Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to insert player deaths")
	}
	return nil
}
