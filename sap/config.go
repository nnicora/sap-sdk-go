package sap

import (
	"github.com/nnicora/sap-sdk-go/sap/endpoints"
	"github.com/nnicora/sap-sdk-go/sap/oauth2"
)

// Runtime configuration used during running the services; Having all prepared configuration,
// will be used for running, in order to save time to prepare again and again for each call.
type RuntimeConfig struct {
	Endpoints map[string]*endpoints.Endpoint

	MaxRetries uint8
}

// Raw Config coming from outside
type Config struct {
	Endpoints  map[string]*EndpointConfig
	MaxRetries uint8

	DefaultOAuth2 *oauth2.Config
}

type EndpointConfig struct {
	Host   string
	OAuth2 *oauth2.Config
}
