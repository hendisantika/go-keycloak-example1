package main

import (
	"go-keycloak-example1/controller"
	"go-keycloak-example1/internal/server"
	"go-keycloak-example1/internal/server/router"
	keycloack "go-keycloak-example1/pkp/keycloak"
	"go-keycloak-example1/server/middleware"
)

func main() {
	// dependency injection
	kc := keycloack.New()
	ctrl := controller.New(kc)
	m := middleware.New(kc)
	route := router.New(ctrl, m)

	// starting server
	server.Start(route)
}
