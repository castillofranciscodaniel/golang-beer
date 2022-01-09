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

func makeBeer() (Beer, error) {
	return NewBeer(2, "Patagonia", "Norte", "Chile", 740, "USD")
}

func makeValidBeerSql() BeerSql {
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

//func Test_Post_should_return_errors_if_the_new_beer_can_not_be_created(t *testing.T) {
//	tearDown := setUp(t)
//	defer tearDown()
//
//	beer, _ := NewBeer(2, "Patagonia", "Norte", "Chile", 740, "CL")
//
//	mockBeerRepository.EXPECT().Post(beer).Return(nil, err.NotFoundError)
//
//	beerService.Post(beer)
//
//	t.Error("Failed while testing the status code")
//
//}

func Test_BoxPrice_same_visa(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	quantity := 6

	beer, _ := makeBeer()
	beerSql := makeValidBeerSql()

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

	beer, _ := makeBeer()
	beerSql := makeValidBeerSql()
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

	beer, _ := makeBeer()
	beerSql := makeValidBeerSql()
	currency := makeCurrencies()
	toCurrency := ""

	mockCurrencyClient.EXPECT().GetCurrencies().Return(currency.Quotes, nil)
	mockBeerRepository.EXPECT().GetById(beerSql.Id.Int64).Return(&beerSql, nil)

	_, err := beerService.BoxPrice(beer.GetId(), toCurrency, quantity)
	if err != err2.CurrencyNotAllowedError {
		t.Error("unexpected error: ", err.Error())
	}

}
