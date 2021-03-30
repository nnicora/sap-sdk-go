package btp

import (
	"github.com/nnicora/sap-sdk-go/sap"
	"github.com/nnicora/sap-sdk-go/sap/oauth2"
	"github.com/nnicora/sap-sdk-go/sap/session"
)

// TODO: - IMPORTANT -- Don't Commit Credentials
const (
	oauth2Username    = ""
	oauth2Password    = ""
	globalAccountGuid = ""
)

var oauth2Config = &oauth2.Config{
	GrantType: "password",
	Username:  oauth2Username,
	Password:  oauth2Password,
}

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
