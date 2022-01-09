package domain

import (
	"fmt"
	err2 "github.com/castillofranciscodaniel/golang-beers/infrastructure/err"
	"github.com/castillofranciscodaniel/golang-beers/infrastructure/provider"
	"github.com/castillofranciscodaniel/golang-beers/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	productNameLog  = "productName"
	productIdLog    = "productId"
	productPriceLog = "productPrice"
	usd             = "USD"
)

//go:generate mockgen -destination=./mockBeerService.go -package=domain github.com/castillofranciscodaniel/golang-beers/domain BeerService
type BeerService interface {
	Get() ([]Beer, error)
	Post(beerDto Beer) error
	GetById(id int64) (*Beer, error)
	BoxPrice(id int64, toCurrency string, quantity int) (float64, error)
}

type DefaultBeerService struct {
	log             zerolog.Logger
	beersRepository BeerRepository

	currencyClient provider.CurrencyClient
}

func NewBeersServiceDefault(beersRepository BeerRepository, currencyClient provider.CurrencyClient) BeerService {
	return &DefaultBeerService{
		log:             log.With().Str(utils.Struct, "DefaultBeerService").Logger(),
		beersRepository: beersRepository,
		currencyClient:  currencyClient,
	}
}

// Get -
func (d DefaultBeerService) Get() ([]Beer, error) {
	d.log.Info().Str(utils.Method, utils.GetFunc).Msg(utils.InitStr)

	beersSql, err := d.beersRepository.Get()
	if err != nil {
		d.log.Error().Err(err).Str(utils.Method, utils.GetFunc).Send()
		return nil, err
	}
	beers := make([]Beer, 0, len(beersSql))
	for _, beerSql := range beersSql {
		beer, err := beerSql.MapToDomain()
		if err != nil {
			d.log.Error().Err(err).Str(utils.Method, utils.GetFunc).Send()
			return nil, err
		}
		beers = append(beers, beer)
	}

	d.log.Info().Str(utils.Method, utils.GetFunc).Msg(utils.EndStr)
	return beers, err

}

// Post -
func (d DefaultBeerService) Post(beer Beer) error {
	d.log.Info().Str(utils.Method, utils.PostFunc).
		Int64(productIdLog, beer.GetId()).
		Str(productNameLog, beer.GetName()).
		Float64(productPriceLog, beer.GetPrice()).
		Msg(utils.InitStr)

	err := d.beersRepository.Post(beer)
	if err != nil {
		d.log.Error().Err(err).Str(utils.Method, utils.PostFunc).Send()
	} else {
		d.log.Info().Str(utils.Method, utils.PostFunc).Msg(utils.EndStr)
	}
	return err
}

func (d DefaultBeerService) GetById(id int64) (*Beer, error) {
	d.log.Info().Str(utils.Method, utils.ByIdFunc).
		Int64(productIdLog, id).
		Msg(utils.InitStr)

	beerSql, err := d.beersRepository.GetById(id)
	if err != nil {
		d.log.Error().Err(err).Str(utils.Method, utils.ByIdFunc).Send()
		return nil, err
	}

	beer, err := beerSql.MapToDomain()
	if err != nil {
		d.log.Error().Err(err).Str(utils.Method, utils.ByIdFunc).Send()
		return nil, err
	}

	d.log.Info().Str(utils.Method, utils.ByIdFunc).Msg(utils.EndStr)
	return &beer, err
}

func (d DefaultBeerService) BoxPrice(id int64, toCurrency string, quantity int) (float64, error) {
	d.log.Info().Str(utils.Method, utils.ByIdFunc).
		Int64(productIdLog, id).
		Msg(utils.InitStr)

	var changeTotal, totalPrice float64

	beer, err := d.GetById(id)
	if err != nil {
		d.log.Error().Err(err).Str(utils.Method, utils.ByIdFunc).Send()
		return changeTotal, err
	}

	currencies, err := d.currencyClient.GetCurrencies()
	if err != nil {
		d.log.Error().Err(err).Str(utils.Method, utils.ByIdFunc).Send()
		return totalPrice, err
	}

	fromCurrency := beer.GetCurrency()

	if toCurrency == beer.currency { // same visa
		changeTotal = beer.price
	} else if fromCurrency == "USD" { // USD --> any visa
		priceVisa, isOK := currencies[fmt.Sprintf("%v%v", beer.GetCurrency(), toCurrency)]
		if isOK {
			changeTotal = beer.price * priceVisa
		}
		return 0, err2.CurrencyNotAllowedError

	} else if toCurrency == "USD" { // any visa --> USD
		priceVisa, isOK := currencies[fmt.Sprintf("%v%v", toCurrency, beer.GetCurrency())]
		if !isOK {
			return 0, err2.CurrencyNotAllowedError
		}
		changeTotal = beer.price / priceVisa
	} else { // any visa not usd --> any visa not usd
		priceUsdToFrom, isOK := currencies[fmt.Sprintf("%v%v", "USD", beer.GetCurrency())]
		priceUsdToFinish, isOK2 := currencies[fmt.Sprintf("%v%v", "USD", toCurrency)]

		if !isOK || !isOK2 {
			return 0, err2.CurrencyNotAllowedError
		}
		changeTotal = beer.price / priceUsdToFrom * priceUsdToFinish
	}

	totalPrice = changeTotal * float64(quantity)
	d.log.Info().Str(utils.Method, utils.ByIdFunc).Msg(utils.EndStr)
	return totalPrice, nil
}
