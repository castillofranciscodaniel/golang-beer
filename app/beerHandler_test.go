package app

import (
	"bytes"
	"errors"
	"github.com/castillofranciscodaniel/golang-beers/domain"
	"github.com/castillofranciscodaniel/golang-beers/err"
	mockDomain "github.com/castillofranciscodaniel/golang-beers/mocks/domain"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	jsoniter "github.com/json-iterator/go"
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

func Test_Get_should_return_beers_with_code_status_ok(t *testing.T) {

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

func Test_Get_should_return_beers_with_code_status_conflict(t *testing.T) {
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

func Test_Post_should_return_beers_with_code_status_created(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()
	beer, _ := domain.NewBeer(5, "Patagonia", "Norte", "Chile", 270, "CL")
	mockService.EXPECT().Post(beer).Return(nil)

	router.Post("/beers", beerHandler.Post)
	beerRequest := BeerRequest{
		Id:       5,
		Name:     "Patagonia",
		Brewery:  "Norte",
		Country:  "Chile",
		Price:    270,
		Currency: "CL",
	}
	json, _ := jsoniter.Marshal(beerRequest)
	request, _ := http.NewRequest(http.MethodPost, "/beers", bytes.NewBuffer(json))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusCreated {
		t.Error("Failed while testing the status code")
	}
}

func Test_Post_should_return_beers_with_code_status_conflict(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()
	beer, _ := domain.NewBeer(5, "Patagonia", "Norte", "Chile", 270, "CL")
	mockService.EXPECT().Post(beer).Return(err.DuplicatedIdError)

	router.Post("/beers", beerHandler.Post)
	beerRequest := BeerRequest{
		Id:       5,
		Name:     "Patagonia",
		Brewery:  "Norte",
		Country:  "Chile",
		Price:    270,
		Currency: "CL",
	}
	json, _ := jsoniter.Marshal(beerRequest)
	request, _ := http.NewRequest(http.MethodPost, "/beers", bytes.NewBuffer(json))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusConflict {
		t.Error("Failed while testing the status code")
	}
}

func Test_Post_should_return_beers_with_code_status_bad_request(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()
	beer, _ := domain.NewBeer(5, "Patagonia", "Norte", "Chile", 270, "CL")
	mockService.EXPECT().Post(beer).Return(errors.New("any error"))

	router.Post("/beers", beerHandler.Post)
	beerRequest := BeerRequest{
		Id:       5,
		Name:     "Patagonia",
		Brewery:  "Norte",
		Country:  "Chile",
		Price:    270,
		Currency: "CL",
	}
	json, _ := jsoniter.Marshal(beerRequest)
	request, _ := http.NewRequest(http.MethodPost, "/beers", bytes.NewBuffer(json))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusBadRequest {
		t.Error("Failed while testing the status code")
	}
}

func Test_GetById_should_return_beers_with_code_status_ok(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()
	beer, _ := domain.NewBeer(5, "Patagonia", "Norte", "Chile", 270, "CL")
	mockService.EXPECT().GetById(int64(5)).Return(&beer, nil)

	router.Post("/beers/{beerId:[0-9]+}", beerHandler.GetById)

	request, _ := http.NewRequest(http.MethodPost, "/beers/5", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_GetById_should_return_beers_with_code_status_not_found(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()
	mockService.EXPECT().GetById(int64(5)).Return(nil, err.NotFoundError)

	router.Post("/beers/{beerId:[0-9]+}", beerHandler.GetById)

	request, _ := http.NewRequest(http.MethodPost, "/beers/5", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusNotFound {
		t.Error("Failed while testing the status code")
	}
}
