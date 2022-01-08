package app

import (
	"github.com/castillofranciscodaniel/golang-beers/config"
	domain2 "github.com/castillofranciscodaniel/golang-beers/pkg/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Start() {
	routes := chi.NewRouter()

	container := InitializeServer()

	routes.Use(middleware.AllowContentType("application/json", "multipart/form-data"))
	routes.Use(middleware.RequestID)
	routes.Use(middleware.RealIP)
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)
	routes.Mount("/debug", middleware.Profiler())

	routes.Get("/health", container.HealthHandler.Health)

	routes.Route("/beers", func(r chi.Router) {
		r.Get("/", container.BeerHandler.Get)
		r.Post("/", container.BeerHandler.Post)
		r.Get("/{beerId:[0-9]+}", container.BeerHandler.ById)
	})

	log.Error().Err(http.ListenAndServe(":3000", routes)).Send()
}

func InitializeServer() ContainerServiceImp {
	healthHandler := NewHealthHandler()
	dbManagerPostgres := config.NewDbManagerImpl()
	beerRepositoryDb := domain2.NewBeersRepositoryDb(dbManagerPostgres)
	beerService := domain2.NewBeersServiceImpl(beerRepositoryDb)
	beerHandler := NewBeersHandler(beerService)
	containerServiceImp := NewContainerServiceImp(healthHandler, beerHandler)
	return containerServiceImp
}
