package app

import (
	"github.com/castillofranciscodaniel/golang-beers/pkg/domain"
)

type BeerRequest struct {
	Id       int64   `json:"Id"`
	Name     string  `json:"Name"`
	Brewery  string  `json:"Brewery"`
	Country  string  `json:"Country"`
	Price    float64 `json:"Price"`
	Currency string  `json:"Currency"`
}

func (b *BeerRequest) MapToDomain() (domain.Beer, error) {
	return domain.NewBeer(b.Id, b.Name, b.Brewery, b.Country, b.Price, b.Currency)
}
