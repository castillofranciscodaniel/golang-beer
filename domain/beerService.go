package domain

import (
	"context"
	"github.com/castillofranciscodaniel/golang-beers/utils"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	productNameLog  = "productName"
	productIdLog    = "productId"
	productPriceLog = "productPrice"
)

//go:generate mockgen -destination=../mocks/service/mockBeerService.go -package=domain github.com/castillofranciscodaniel/golang-beer/domain BeerService
type BeerService interface {
	Get(ctx context.Context) ([]BeerSql, error)
	Post(ctx context.Context, beerDto Beer) error
	ById(ctx context.Context, id int64) (*BeerSql, error)
}

type DefaultBeerService struct {
	log             zerolog.Logger
	beersRepository BeerRepository
}

func NewBeersServiceImpl(beersRepository BeerRepository) BeerService {
	return &DefaultBeerService{
		beersRepository: beersRepository,
		log:             log.With().Str(utils.Struct, "DefaultBeerService").Logger(),
	}
}

// Get -
func (d DefaultBeerService) Get(ctx context.Context) ([]BeerSql, error) {
	d.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.GetFunc).Msg(utils.InitStr)

	beers, err := d.beersRepository.Get(ctx)
	if err != nil {
		d.log.Error().Err(err).Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.GetFunc).Send()

	} else {
		d.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.GetFunc).Msg(utils.EndStr)
	}
	return beers, err

}

// Post -
func (d DefaultBeerService) Post(ctx context.Context, beer Beer) error {
	d.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.PostFunc).
		Int64(productIdLog, beer.getId()).
		Str(productNameLog, beer.getName()).
		Float64(productPriceLog, beer.getPrice()).
		Msg(utils.InitStr)

	err := d.beersRepository.Post(ctx, beer)
	if err != nil {
		d.log.Error().Err(err).Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.PostFunc).Send()
	} else {
		d.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.PostFunc).Msg(utils.EndStr)
	}
	return err
}

func (d DefaultBeerService) ById(ctx context.Context, id int64) (*BeerSql, error) {
	d.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.ByIdFunc).
		Int64(productIdLog, id).
		Msg(utils.InitStr)

	beer, err := d.beersRepository.ById(ctx, id)
	if err != nil {
		d.log.Error().Err(err).Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.ByIdFunc).Send()
	} else {
		d.log.Info().Str(utils.Thread, middleware.GetReqID(ctx)).Str(utils.Method, utils.ByIdFunc).Msg(utils.EndStr)
	}
	return beer, err

}
