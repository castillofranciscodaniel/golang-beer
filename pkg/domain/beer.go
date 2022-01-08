package domain

import (
	"context"
	"github.com/castillofranciscodaniel/golang-beers/pkg/err"
	"gopkg.in/guregu/null.v4"
)

type Beer struct {
	id       int64
	name     string
	brewery  string
	country  string
	price    float64
	currency string
}

func NewBeer(id int64, name string, brewery string, country string, price float64, currency string) (Beer, error) {
	if id < 0 {
		return Beer{}, err.IdCanNotBeMinorThanOne
	}

	if id < 0 {
		return Beer{}, err.PriceCanNotBeMinorThanZero
	}

	return Beer{
		id:       id,
		name:     name,
		brewery:  brewery,
		country:  country,
		price:    price,
		currency: currency,
	}, nil
}

func (b *Beer) getId() int64 {
	return b.id
}

func (b *Beer) getName() string {
	return b.name
}

func (b *Beer) getCountry() string {
	return b.country
}

func (b *Beer) getPrice() float64 {
	return b.price
}

func (b *Beer) getCurrency() string {
	return b.currency
}

func (b *Beer) getBrewery() string {
	return b.brewery
}

type BeerSql struct {
	Id       null.Int
	Name     null.String
	Brewery  null.String
	Country  null.String
	Price    null.Float
	Currency null.String
}

type BeerRepository interface {
	Get(ctx context.Context) ([]BeerSql, error)
	Post(ctx context.Context, beer Beer) error
	ById(ctx context.Context, id int64) (*BeerSql, error)
}
