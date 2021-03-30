package btp

import (
	"github.com/nnicora/sap-sdk-go/sap"
	"github.com/nnicora/sap-sdk-go/sap/oauth2"
	"github.com/nnicora/sap-sdk-go/sap/session"
)

var cfg = &sap.Config{
	Endpoints: map[string]*sap.EndpointConfig{
		"accounts": {
			Host: "https://accounts-service.cfapps.eu10.hana.ondemand.com",
		},
		"entitlements": {
			Host: "https://accounts-service.cfapps.eu10.hana.ondemand.com",
		},
		"events": {
			Host: "https://events-service.cfapps.eu10.hana.ondemand.com",
		},
	},

	DefaultOAuth2: oauth2Config,
}

var sess *session.RuntimeSession

func init() {
	sessTmp, err := session.BuildFromConfig(cfg)
	if err != nil {
		panic(err)
	}
	sess = sessTmp
}
