package domain

import (
	"context"
	"github.com/castillofranciscodaniel/golang-beers/config"
	"github.com/castillofranciscodaniel/golang-beers/pkg/utils"
	"github.com/go-chi/chi/v5/middleware"
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
func (b *BeerRepositoryStub) Get(ctx context.Context) ([]BeerSql, error) {
	b.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.GetFunc).Msgf(utils.InitStr)
	var beersSql []BeerSql
	b.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.GetFunc).Send()
	return beersSql, nil
}

// Post -
func (b *BeerRepositoryStub) Post(ctx context.Context, beer Beer) error {
	b.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.PostFunc).Msgf(utils.InitStr)
	b.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.PostFunc).Send()
	return nil
}

func (b *BeerRepositoryStub) ById(ctx context.Context, id int64) (*BeerSql, error) {
	b.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.PostFunc).Msgf(utils.InitStr)
	b.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.PostFunc).Send()
	return nil, nil
}
