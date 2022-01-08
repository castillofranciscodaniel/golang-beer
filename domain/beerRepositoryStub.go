package domain

import (
	"github.com/castillofranciscodaniel/golang-beers/config"
	"github.com/castillofranciscodaniel/golang-beers/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BeerRepositoryStub struct {
	log       zerolog.Logger
	dbManager config.DbManager
}

func NewBeerRepositoryStub(dbManager config.DbManager) BeerRepositoryStub {
	return BeerRepositoryStub{
		dbManager: dbManager,
		log:       log.With().Str(utils.Struct, "BeerRepositoryDb").Logger(),
	}
}

// Get -
func (b *BeerRepositoryStub) Get() ([]BeerSql, error) {
	b.log.Info().Str(utils.Method, utils.GetFunc).Msgf(utils.InitStr)
	var beersSql []BeerSql
	b.log.Info().Str(utils.Method, utils.GetFunc).Send()
	return beersSql, nil
}

// Post -
func (b *BeerRepositoryStub) Post(beer Beer) error {
	b.log.Info().Str(utils.Method, utils.PostFunc).Msgf(utils.InitStr)
	b.log.Info().Str(utils.Method, utils.PostFunc).Send()
	return nil
}

func (b *BeerRepositoryStub) ById(id int64) (*BeerSql, error) {
	b.log.Info().Str(utils.Method, utils.PostFunc).Msgf(utils.InitStr)
	b.log.Info().Str(utils.Method, utils.PostFunc).Send()
	return nil, nil
}
