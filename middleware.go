package main

type keyCloakMiddleware struct {
	keycloak *keycloak
}

func newMiddleware(keycloak *keycloak) *keyCloakMiddleware {
	return &keyCloakMiddleware{keycloak: keycloak}
}
