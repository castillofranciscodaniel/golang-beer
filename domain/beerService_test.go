package domain

import (
	"database/sql"
	"fmt"
	err2 "github.com/castillofranciscodaniel/golang-beers/infrastructure/err"
	"github.com/castillofranciscodaniel/golang-beers/infrastructure/provider"
	"github.com/golang/mock/gomock"
	"testing"
)

var mockBeerRepository *MockBeerRepository
var beerService DefaultBeerService
var mockCurrencyClient *provider.MockCurrencyClient

func setUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockBeerRepository = NewMockBeerRepository(ctrl)
	mockCurrencyClient = provider.NewMockCurrencyClient(ctrl)

	beerService = DefaultBeerService{beersRepository: mockBeerRepository, currencyClient: mockCurrencyClient}

	return func() {
		defer ctrl.Finish()
	}
}

func makeBeerUsd() (Beer, error) {
	return NewBeer(2, "Patagonia", "Norte", "Chile", 740, "USD")
}

func makeBeerClp() (Beer, error) {
	return NewBeer(2, "Patagonia", "Norte", "Chile", 740, "CLP")
}

func makeValidBeerSqlClp() BeerSql {
	return BeerSql{
		Id:       sql.NullInt64{Int64: 2, Valid: true},
		Name:     sql.NullString{String: "Patagonia", Valid: true},
		Brewery:  sql.NullString{String: "Norte", Valid: true},
		Country:  sql.NullString{String: "Chile", Valid: true},
		Price:    sql.NullFloat64{Float64: 740, Valid: true},
		Currency: sql.NullString{String: "CLP", Valid: true},
	}
}

func makeValidBeerSqlUsd() BeerSql {
	return BeerSql{
		Id:       sql.NullInt64{Int64: 2, Valid: true},
		Name:     sql.NullString{String: "Patagonia", Valid: true},
		Brewery:  sql.NullString{String: "Norte", Valid: true},
		Country:  sql.NullString{String: "Chile", Valid: true},
		Price:    sql.NullFloat64{Float64: 740, Valid: true},
		Currency: sql.NullString{String: "USD", Valid: true},
	}
}

func makeCurrencies() provider.Currencies {
	return provider.Currencies{
		Success: true,
		Quotes: map[string]float64{
			"USDCLP": 10,
			"USDARS": 50,
		},
	}

}

func Test_Post_should_return_errors_if_the_id_exist(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	beer, _ := makeBeerUsd()

	mockBeerRepository.EXPECT().Post(beer).Return(err2.DuplicatedIdError)

	err := beerService.Post(beer)

	if err != nil && err != err2.DuplicatedIdError {
		t.Error("the error should be: ", err.Error())
	}

}

func Test_Post_should_return_any_error_when_beer_is_created(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	beer, _ := makeBeerUsd()

	mockBeerRepository.EXPECT().Post(beer).Return(nil)

	err := beerService.Post(beer)

	if err != nil {
		t.Error("the error should be: ", err.Error())
	}

}

func Test_BoxPrice_same_visa(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	quantity := 6

	beer, _ := makeBeerUsd()
	beerSql := makeValidBeerSqlUsd()

	mockCurrencyClient.EXPECT().GetCurrencies().Return(nil, nil)
	mockBeerRepository.EXPECT().GetById(beerSql.Id.Int64).Return(&beerSql, nil)

	price, err := beerService.BoxPrice(beer.GetId(), "USD", quantity)
	if err != nil {
		t.Error(err.Error())
	}

	if price != float64(quantity)*beer.price {
		t.Error("Failed distinct price total")
	}

}

func Test_BoxPrice_usd_to_another_visa(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	quantity := 6

	beer, _ := makeBeerUsd()
	beerSql := makeValidBeerSqlUsd()
	currency := makeCurrencies()
	toCurrency := "CLP"

	mockCurrencyClient.EXPECT().GetCurrencies().Return(currency.Quotes, nil)
	mockBeerRepository.EXPECT().GetById(beerSql.Id.Int64).Return(&beerSql, nil)

	price, err := beerService.BoxPrice(beer.GetId(), toCurrency, quantity)
	if err != nil {
		t.Error(err.Error())
	}

	expectedPrice := float64(quantity) * beer.price * currency.Quotes[fmt.Sprintf("%v%v", beer.GetCurrency(), toCurrency)]
	if price != expectedPrice {
		t.Error("Failed distinct price total")
	}

}

func Test_BoxPrice_currency_not_allowed(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	quantity := 6

	beer, _ := makeBeerUsd()
	beerSql := makeValidBeerSqlUsd()
	currency := makeCurrencies()
	toCurrency := ""

	mockCurrencyClient.EXPECT().GetCurrencies().Return(currency.Quotes, nil)
	mockBeerRepository.EXPECT().GetById(beerSql.Id.Int64).Return(&beerSql, nil)

	_, err := beerService.BoxPrice(beer.GetId(), toCurrency, quantity)
	if err != err2.CurrencyNotAllowedError {
		t.Error("unexpected error: ", err.Error())
	}

}

func Test_BoxPrice_any_visa_to_usd(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	quantity := 6

	beer, _ := makeBeerClp()
	beerSql := makeValidBeerSqlClp()
	currency := makeCurrencies()
	toCurrency := "USD"

	mockCurrencyClient.EXPECT().GetCurrencies().Return(currency.Quotes, nil)
	mockBeerRepository.EXPECT().GetById(beerSql.Id.Int64).Return(&beerSql, nil)

	price, err := beerService.BoxPrice(beer.GetId(), toCurrency, quantity)
	if err != nil {
		t.Error(err.Error())
	}

	expectedPrice := float64(quantity) * beer.price / currency.Quotes[fmt.Sprintf("%v%v", toCurrency, beer.GetCurrency())]
	if price != expectedPrice {
		t.Error("Failed distinct price total")
	}

}

func Test_BoxPrice_any_visa_not_usd_to_any_visa_not_usd(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	quantity := 6

	beer, _ := makeBeerClp()
	beerSql := makeValidBeerSqlClp()
	currency := makeCurrencies()
	toCurrency := "ARS"

	mockCurrencyClient.EXPECT().GetCurrencies().Return(currency.Quotes, nil)
	mockBeerRepository.EXPECT().GetById(beerSql.Id.Int64).Return(&beerSql, nil)

	price, err := beerService.BoxPrice(beer.GetId(), toCurrency, quantity)
	if err != nil {
		t.Error(err.Error())
	}

	priceUsdToFrom := currency.Quotes[fmt.Sprintf("%v%v", "USD", beer.GetCurrency())]
	priceUsdToFinish := currency.Quotes[fmt.Sprintf("%v%v", "USD", toCurrency)]

	expectedPrice := float64(quantity) * beer.price / priceUsdToFrom * priceUsdToFinish
	if price != expectedPrice {
		t.Error("Failed distinct price total")
	}

}

func Test_Get_should_return_a_list(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	dummyBeersSql := []BeerSql{
		makeValidBeerSqlClp(),
		makeValidBeerSqlUsd(),
	}

	mockBeerRepository.EXPECT().Get().Return(dummyBeersSql, nil)

	beers, err := beerService.Get()

	if err != nil {
		t.Error("the error should be: ", err.Error())
	}

	if beers[0].id != dummyBeersSql[0].Id.Int64 {
		t.Error("the id should be the same")
	}

	if beers[1].id != dummyBeersSql[1].Id.Int64 {
		t.Error("the id should be the same")
	}
}
