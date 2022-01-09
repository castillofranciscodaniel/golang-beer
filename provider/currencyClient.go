package provider

import (
	"fmt"
	"github.com/castillofranciscodaniel/golang-beers/utils"
	"github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

const (
	url               = "https://api.currencylayer.com/convert?from=%V&to=%V&amount=%V"
	convertAmountFunc = "ConvertAmount"
)

type CurrencyClient interface {
	ConvertAmount(fromCurrency, toCurrency string, amount float64) (float64, error)
}

type CurrencyClientDefault struct {
	log zerolog.Logger
}

func NewCurrencyClientDefault() CurrencyClientDefault {
	return CurrencyClientDefault{
		log: log.With().Str(utils.Struct, "CurrencyClientDefault").Logger(),
	}
}

func (c CurrencyClientDefault) ConvertAmount(fromCurrency, toCurrency string, amount float64) (float64, error) {
	c.log.Info().Str(utils.Method, convertAmountFunc).Msg(utils.InitStr)

	url := fmt.Sprint(url, fromCurrency, toCurrency, amount)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 60,
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	var mapita map[string]float64
	err = jsoniter.NewDecoder(res.Body).Decode(&mapita)
	if err != nil {
		return 0, err
	}

	return mapita[""], nil
}
