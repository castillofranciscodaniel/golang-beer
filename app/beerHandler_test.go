package app

import (
	"context"
	"github.com/castillofranciscodaniel/golang-beers/mocks/domain"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_should_return_beers_with_code_status_ok(t *testing.T) {
	ctrol := gomock.NewController(t)
	defer ctrol.Finish()
	mockService := domain.NewMockBeerService(ctrol)

	dummyBeers := []BeerRequest{}

	mockService.EXPECT().Get(context.Background()).Return(dummyBeers, nil)
	bh := BeerHandler{
		beersService: mockService,
	}
	router := chi.NewRouter()
	router.Get("/beers", bh.Get)
}
