package provider

import (
	"fmt"
	err2 "github.com/castillofranciscodaniel/golang-beers/infrastructure/err"
	"github.com/castillofranciscodaniel/golang-beers/utils"
	"github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

type Currencies struct {
	Success bool               `json:"success"`
	Quotes  map[string]float64 `json:"Quotes"`
}

const (
	getCurrenciesFunc = "GetCurrencies"
)

//go:generate mockgen -destination=./mockCurrencyClient.go -package=provider github.com/castillofranciscodaniel/golang-beers/infrastructure/provider CurrencyClient
type CurrencyClient interface {
	GetCurrencies() (map[string]float64, error)
}

type CurrencyClientDefault struct {
	log              zerolog.Logger
	accessKey        string
	currencyEndpoint string
}

func NewCurrencyClientDefault() CurrencyClientDefault {
	return CurrencyClientDefault{
		log:              log.With().Str(utils.Struct, "CurrencyClientDefault").Logger(),
		accessKey:        os.Getenv("KEY_CURRENCY_LAYER"),
		currencyEndpoint: os.Getenv("CURRENCY_API"),
	}
}

func (c CurrencyClientDefault) GetCurrencies() (map[string]float64, error) {
	c.log.Info().Str(utils.Method, getCurrenciesFunc).Msg(utils.InitStr)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/%v", c.currencyEndpoint, "live"), nil)
	req.Header.Set("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("access_key", c.accessKey)

	req.URL.RawQuery = query.Encode()

	client := &http.Client{
		Timeout: time.Second * 60,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var result Currencies
	err = jsoniter.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, err2.ErrorTakingCurrencies
	}
	return result.Quotes, nil
}
