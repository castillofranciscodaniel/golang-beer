package app

type ContainerServiceImp struct {
	//restyClient   *resty.Client
	HealthHandler HealthHandler
	BeerHandler   BeerHandler
}

func NewContainerServiceImp(healthHandler HealthHandler, beerHandler BeerHandler) ContainerServiceImp {
	return ContainerServiceImp{
		HealthHandler: healthHandler,
		BeerHandler:   beerHandler,
	}
}
