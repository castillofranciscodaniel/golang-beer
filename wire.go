//go:build InitializeServer
// +build InitializeServer

package main

import (
	"github.com/castillofranciscodaniel/golang-beers/config"
	port2 "github.com/castillofranciscodaniel/golang-beers/pkg/app"
	domain2 "github.com/castillofranciscodaniel/golang-beers/pkg/domain"
	"github.com/google/wire"
)

func InitializeServer() port2.ContainerServiceImp {
	wire.Build(

		wire.Bind(new(config.DbManager), new(config.DbManagerPostgres)),
		wire.Bind(new(domain2.BeerRepository), new(domain2.BeerRepositoryDb)),

		config.NewDbManagerImpl,
		port2.NewHealthHandler,
		domain2.NewBeersRepositoryDb,
		domain2.NewBeersServiceImpl,
		port2.NewBeersHandler,
		port2.NewContainerServiceImp,
	)

	return port2.ContainerServiceImp{}
}
