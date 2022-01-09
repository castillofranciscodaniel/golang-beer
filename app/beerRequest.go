package app

import (
	"github.com/castillofranciscodaniel/golang-beers/domain"
)

type BeerRequest struct {
	Id       int64   `json:"Id"`
	Name     string  `json:"Name"`
	Brewery  string  `json:"Brewery"`
	Country  string  `json:"Country"`
	Price    float64 `json:"Price"`
	Currency string  `json:"Currency"`
}

type BeerBoxRequest struct {
	PriceTotal float64 `json:"PriceTotal"`
}

func (b *BeerRequest) MapToDomain() (domain.Beer, error) {
	return domain.NewBeer(b.Id, b.Name, b.Brewery, b.Country, b.Price, b.Currency)
}

func (b *BeerRequest) DomainToRequest(beer domain.Beer) BeerRequest {

	b.Id = beer.GetId()
	b.Name = beer.GetName()
	b.Brewery = beer.GetBrewery()
	b.Country = beer.GetCountry()
	b.Price = beer.GetPrice()
	b.Currency = beer.GetCurrency()
	return *b
}
