//go:build InitializeServer
// +build InitializeServer

package main

import (
	app2 "github.com/castillofranciscodaniel/golang-beers/pkg/app"
	config2 "github.com/castillofranciscodaniel/golang-beers/pkg/config"
	domain2 "github.com/castillofranciscodaniel/golang-beers/pkg/domain"
	"github.com/google/wire"
)

func InitializeServer() app2.ContainerServiceImp {
	wire.Build(

		wire.Bind(new(config2.DbManager), new(config2.DbManagerPostgres)),
		wire.Bind(new(domain2.BeerRepository), new(domain2.BeerRepositoryDb)),

		config2.NewDbManagerImpl,
		app2.NewHealthHandler,
		domain2.NewBeersRepositoryDb,
		domain2.NewBeersServiceImpl,
		app2.NewBeersHandler,
		app2.NewContainerServiceImp,
	)

	return app2.ContainerServiceImp{}
}
