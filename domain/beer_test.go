package domain

import (
	err2 "github.com/castillofranciscodaniel/golang-beers/err"
	"testing"
)

func Test_NewBeerOK(t *testing.T) {
	_, err := NewBeer(1, "Patagonia", "Norte", "Chile", 270, "CL")
	if err != nil {
		t.Error("Its is should be ok")
	}
}

func Test_NewBeerBadId(t *testing.T) {
	_, err := NewBeer(-1, "Patagonia", "Norte", "Chile", 270, "CL")
	if err != err2.IdCanNotBeMinorThanOne {
		t.Error("id can not be minor than one")
	}

	_, err = NewBeer(0, "Patagonia", "Norte", "Chile", 270, "CL")
	if err != err2.IdCanNotBeMinorThanOne {
		t.Error("id can not be minor than one")
	}
}

func Test_NewBeerBadPrice(t *testing.T) {
	_, err := NewBeer(1, "Patagonia", "Norte", "Chile", -1, "CL")
	if err != err2.PriceCanNotBeMinorThanZero {
		t.Error("price can not be minor than zero")
	}
}

func Test_NewBeerPriceZeroOk(t *testing.T) {
	_, err := NewBeer(1, "Patagonia", "Norte", "Chile", 0, "CL")
	if err != nil {
		t.Error("Its is should be ok")
	}
}
