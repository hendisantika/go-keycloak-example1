package main

import "github.com/Nerzal/gocloak/v13"

type keycloak struct {
	//gocloak      gocloak.GoCloak // keycloak client
	client       gocloak.GoCloak // keycloak client
	realm        string          // realm specified in Keycloak
	clientId     string          // clientId specified in Keycloak
	clientSecret string          // client secret specified in Keycloak
}

func newKeycloak() *keycloak {
	return &keycloak{
		client: gocloak.NewClient("http://localhost:8080"),
		//client := gocloak.NewClient(URL, gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("realms")),
		realm:        "PowerRanger",
		clientId:     "my-go-service",
		clientSecret: "abfa2984-9125-486b-b360-03386ad13e08",
	}
}
