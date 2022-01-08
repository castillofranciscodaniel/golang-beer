//go:build InitializeServer
// +build InitializeServer

package main

import (
	"github.com/castillofranciscodaniel/golang-beers/app"
	"github.com/castillofranciscodaniel/golang-beers/config"
	"github.com/castillofranciscodaniel/golang-beers/domain"
	"github.com/google/wire"
)

func InitializeServer() app.ContainerServiceImp {
	wire.Build(

		wire.Bind(new(config.DbManager), new(config.DbManagerPostgres)),
		wire.Bind(new(domain.BeerRepository), new(domain.BeerRepositoryDb)),

		config.NewDbManagerImpl,
		app.NewHealthHandler,
		domain.NewBeersRepositoryDb,
		domain.NewBeersServiceImpl,
		app.NewBeersHandler,
		app.NewContainerServiceImp,
	)

	return app.ContainerServiceImp{}
}
