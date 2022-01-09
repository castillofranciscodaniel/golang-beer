package app

import (
	"github.com/castillofranciscodaniel/golang-beers/domain"
	"github.com/castillofranciscodaniel/golang-beers/infrastructure/persistence"
	provider2 "github.com/castillofranciscodaniel/golang-beers/infrastructure/provider"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Start() {
	routes := chi.NewRouter()

	container := InitializeServer()

	routes.Use(middleware.AllowContentType("application/json", "multipart/form-data"))
	routes.Use(middleware.RealIP)
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)

	routes.Mount("/debug", middleware.Profiler())

	routes.Get("/health", container.HealthHandler.Health)

	routes.Route("/beers", func(r chi.Router) {
		r.Get("/", container.BeerHandler.Get)
		r.Post("/", container.BeerHandler.Post)
		r.Get("/{beerId:[0-9]+}", container.BeerHandler.GetById)
		r.Get("/{beerId:[0-9]+}/boxprice", container.BeerHandler.BoxPrice)
	})

	log.Error().Err(http.ListenAndServe(":8080", routes)).Send()
}

func InitializeServer() ContainerServiceImp {
	healthHandler := NewHealthHandler()
	dbManagerPostgres := persistence.NewDbManagerImpl()
	beerRepositoryDb := domain.NewBeersRepositoryDb(dbManagerPostgres)
	currencyClientDefault := provider2.NewCurrencyClientDefault()
	beerService := domain.NewBeersServiceDefault(beerRepositoryDb, currencyClientDefault)
	beerHandler := NewBeersHandler(beerService)
	containerServiceImp := NewContainerServiceImp(healthHandler, beerHandler)
	return containerServiceImp
}
