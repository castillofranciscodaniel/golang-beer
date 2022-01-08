package app

import (
	"errors"
	"github.com/castillofranciscodaniel/golang-beers/domain"
	mockDomain "github.com/castillofranciscodaniel/golang-beers/mocks/domain"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockService *mockDomain.MockBeerService
var beerHandler BeerHandler
var router *chi.Mux

func setUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = mockDomain.NewMockBeerService(ctrl)
	beerHandler = BeerHandler{beersService: mockService}

	router = chi.NewRouter()

	return func() {
		defer ctrl.Finish()
	}
}

func Test_should_return_beers_with_code_status_ok(t *testing.T) {

	tearDown := setUp(t)
	defer tearDown()

	beer1, _ := domain.NewBeer(2, "Patagonia", "Norte", "Chile", 740, "CL")
	beer2, _ := domain.NewBeer(2, "Quilmes", "Sur", "Chile", 710, "CL")
	dummyBeers := []domain.Beer{
		beer1,
		beer2,
	}

	mockService.EXPECT().Get().Return(dummyBeers, nil)
	router.Get("/beers", beerHandler.Get)

	request, _ := http.NewRequest(http.MethodGet, "/beers", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_should_return_beers_with_code_status_conflict(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	mockService.EXPECT().Get().Return(nil, errors.New("any error"))

	router.Get("/beers", beerHandler.Get)

	request, _ := http.NewRequest(http.MethodGet, "/beers", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusConflict {
		t.Error("Failed while testing the status code")
	}
}
