package keycloack

import (
	"github.com/Nerzal/gocloak/v13"
	"go-keycloak-example1/internal/config"
)

// Keycloack data model
type Keycloak struct {
	Client       *gocloak.GoCloak // keycloak client
	Realm        string           // Realm specified in Keycloak
	ClientId     string           // ClientId specified in Keycloak
	ClientSecret string           // client secret specified in Keycloak
}

// New : create new keycloack instance
func New() *Keycloak {
	return &Keycloak{
		Client:       gocloak.NewClient("http://localhost:8080"),
		Realm:        config.Realm,
		ClientId:     config.ClientID,
		ClientSecret: config.ClientSecret,
	}
}
