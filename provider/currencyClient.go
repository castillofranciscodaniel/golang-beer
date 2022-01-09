package provider

import (
	err2 "github.com/castillofranciscodaniel/golang-beers/err"
	"github.com/castillofranciscodaniel/golang-beers/utils"
	"github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

const (
	currencyEndpoint  = "https://api.currencylayer.com/live"
	getCurrenciesFunc = "GetCurrencies"
)

type CurrencyClient interface {
	GetCurrencies() (interface{}, error)
}

type CurrencyClientDefault struct {
	log       zerolog.Logger
	accessKey string
}

func NewCurrencyClientDefault() CurrencyClientDefault {
	return CurrencyClientDefault{
		log: log.With().Str(utils.Struct, "CurrencyClientDefault").Logger(),
		//accessKey:   os.Getenv("KEY_CURRENCY_LAYER"),
		accessKey: "3fbc64e43d0f0af2089938650bd3635b",
	}
}

func (c CurrencyClientDefault) GetCurrencies() (interface{}, error) {
	c.log.Info().Str(utils.Method, getCurrenciesFunc).Msg(utils.InitStr)

	req, err := http.NewRequest(http.MethodGet, currencyEndpoint, nil)
	req.Header.Set("Content-Type", "application/json")

	query := req.URL.Query()
	query.Add("access_key", c.accessKey)

	req.URL.RawQuery = query.Encode()

	client := &http.Client{
		Timeout: time.Second * 60,
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	var result map[string]interface{}
	err = jsoniter.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	currencies, isOk := result["quotes"]
	if isOk {
		return currencies, nil
	}
	return nil, err2.ErrorTakingCurrencies
}
