package main

import "github.com/Nerzal/gocloak/v13"

type keycloak struct {
	client       *gocloak.GoCloak // keycloak client
	realm        string           // realm specified in Keycloak
	clientId     string           // clientId specified in Keycloak
	clientSecret string           // client secret specified in Keycloak
}

func newKeycloak() *keycloak {
	return &keycloak{
		client:       gocloak.NewClient("http://localhost:8080"),
		realm:        "PowerRanger",
		clientId:     "my-go-service",
		clientSecret: "PE73ia15f68FLHIM9UqhZf729lEqijhR",
	}
}
