package auth

import (
	"os"

	"github.com/Nerzal/gocloak/v13"
)

type keycloak struct {
	Gocloak      *gocloak.GoCloak
	ClientId     string
	ClientSecret string
	Realm        string
}

func newKeycloak() *keycloak {
	return &keycloak{
		Gocloak:      gocloak.NewClient(os.Getenv("KEYCLOAK_HOST")),
		ClientId:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		Realm:        os.Getenv("KEYCLOAK_REALM"),
	}
}

// Returns the keycloak client
var KC = newKeycloak()
