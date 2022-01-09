//go:build InitializeServer
// +build InitializeServer

package main

import (
	"github.com/castillofranciscodaniel/golang-beers/app"
	"github.com/castillofranciscodaniel/golang-beers/domain"
	"github.com/castillofranciscodaniel/golang-beers/infrastructure/persistence"
	provider2 "github.com/castillofranciscodaniel/golang-beers/infrastructure/provider"
	"github.com/google/wire"
)

func InitializeServer() app.ContainerServiceImp {
	wire.Build(

		wire.Bind(new(persistence.DbManager), new(persistence.DbManagerPostgres)),
		wire.Bind(new(domain.BeerRepository), new(domain.BeerRepositoryDb)),
		wire.Bind(new(provider2.CurrencyClient), new(provider2.CurrencyClientDefault)),

		persistence.NewDbManagerImpl,
		app.NewHealthHandler,
		domain.NewBeersRepositoryDb,
		domain.NewBeersServiceDefault,
		app.NewBeersHandler,

		provider2.NewCurrencyClientDefault,

		app.NewContainerServiceImp,
	)

	return app.ContainerServiceImp{}
}
