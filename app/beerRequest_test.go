package app

import (
	"testing"
)

func Test_MapToDomain(t *testing.T) {
	beer := BeerRequest{
		Id:       2,
		Name:     "Patagonia",
		Brewery:  "Norte",
		Country:  "Chile",
		Price:    270,
		Currency: "CL",
	}

	_, err := beer.MapToDomain()
	if err != nil {
		t.Error("should be ok")
	}
}
