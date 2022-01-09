//go:build InitializeServer
// +build InitializeServer

package main

import (
	"github.com/castillofranciscodaniel/golang-beers/app"
	"github.com/castillofranciscodaniel/golang-beers/domain"
	"github.com/castillofranciscodaniel/golang-beers/provider"
	"github.com/google/wire"
)

func InitializeServer() app.ContainerServiceImp {
	wire.Build(

		wire.Bind(new(provider.DbManager), new(provider.DbManagerPostgres)),
		wire.Bind(new(domain.BeerRepository), new(domain.BeerRepositoryDb)),
		wire.Bind(new(provider.CurrencyClient), new(provider.CurrencyClientDefault)),

		provider.NewDbManagerImpl,
		app.NewHealthHandler,
		domain.NewBeersRepositoryDb,
		domain.NewBeersServiceDefault,
		app.NewBeersHandler,

		provider.NewCurrencyClientDefault,

		app.NewContainerServiceImp,
	)

	return app.ContainerServiceImp{}
}
