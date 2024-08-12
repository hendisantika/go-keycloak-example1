package main

import "strings"

type keyCloakMiddleware struct {
	keycloak *keycloak
}

func newMiddleware(keycloak *keycloak) *keyCloakMiddleware {
	return &keyCloakMiddleware{keycloak: keycloak}
}

func (auth *keyCloakMiddleware) extractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}
