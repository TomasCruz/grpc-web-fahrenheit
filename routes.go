package main

import (
	"net/http"

	"github.com/TomasCruz/grpc-web-fahrenheit/presenter"

	"github.com/TomasCruz/grpc-web-fahrenheit/configuration"
)

func newRoutes() configuration.Routes {
	return configuration.Routes{
		HealthRoute: "/health/",
		C2FRoute:    "/c2f/",
		F2CRoute:    "/f2c/",
	}
}

func bindRoutes(routes configuration.Routes) {
	http.HandleFunc(routes.HealthRoute, presenter.HealthHandler)

	http.HandleFunc(routes.C2FRoute, presenter.Float64InputValidator(
		routes.C2FRoute, presenter.C2FHandler))

	http.HandleFunc(routes.F2CRoute, presenter.Float64InputValidator(
		routes.F2CRoute, presenter.F2CHandler))
}
