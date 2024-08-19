package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	keycloack "go-keycloak-example1/pkg/keycloak"
	"net/http"
	"strings"
)

type Authentication struct {
	keycloak *keycloack.Keycloak
}

func New(keycloak *keycloack.Keycloak) *Authentication {

	return &Authentication{keycloak: keycloak}
}

func (m *Authentication) VerifyToken(next http.Handler) http.Handler {

	f := func(w http.ResponseWriter, r *http.Request) {

		// try to extract Authorization parameter from the HTTP header
		token := r.Header.Get("Authorization")

		if token == "" {

			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// extract Bearer token
		token = extractBearerToken(token)

		if token == "" {
			http.Error(w, "Bearer Token missing", http.StatusUnauthorized)
			return
		}

		//// call Keycloak API to verify the access token
		result, err := m.keycloak.Client.RetrospectToken(context.Background(), token, m.keycloak.ClientId, m.keycloak.ClientSecret, m.keycloak.Realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		jwt, _, err := m.keycloak.Client.DecodeAccessToken(context.Background(), token, m.keycloak.Realm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid or malformed token: %s", err.Error()), http.StatusUnauthorized)
			return
		}

		jwtj, _ := json.Marshal(jwt)
		fmt.Printf("token: %v\n", string(jwtj))

		// check if the token isn't expired and valid
		if !*result.Active {
			http.Error(w, "Invalid or expired Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}

func extractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}
