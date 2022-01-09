package domain

import (
	"database/sql"
	"github.com/castillofranciscodaniel/golang-beers/err"
	"github.com/castillofranciscodaniel/golang-beers/provider"
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

func Test_Post_should_return_errors_if_the_new_beer_can_not_be_created(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	beer, _ := NewBeer(2, "Patagonia", "Norte", "Chile", 740, "CL")

	mockBeerRepository.EXPECT().Post(beer).Return(nil, err.NotFoundError)

	beerService.Post(beer)

	t.Error("Failed while testing the status code")

}

func Test_BoxPrice_same_visa(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	quantity := 6

	beer, _ := NewBeer(2, "Patagonia", "Norte", "Chile", 740, "USD")
	beerSql := BeerSql{
		Id:       sql.NullInt64{Int64: 2, Valid: true},
		Name:     sql.NullString{String: "Patagonia", Valid: true},
		Brewery:  sql.NullString{String: "Norte", Valid: true},
		Country:  sql.NullString{String: "Chile", Valid: true},
		Price:    sql.NullFloat64{Float64: 740, Valid: true},
		Currency: sql.NullString{String: "USD", Valid: true},
	}

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
