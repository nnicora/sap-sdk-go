package btp

import (
	"github.com/nnicora/sap-sdk-go/sap"
	"github.com/nnicora/sap-sdk-go/sap/oauth2"
	"github.com/nnicora/sap-sdk-go/sap/session"
	"os"
)

var (
	globalAccountGuid = os.Getenv("SAP_BTP_GLOBAL_ACCOUNT")
)

var sess *session.RuntimeSession

func init() {
	const (
		SAP_OAUTH2_USERNAME      = "SAP_OAUTH2_USERNAME"
		SAP_OAUTH2_PASSWORD      = "SAP_OAUTH2_PASSWORD"
		SAP_OAUTH2_GRANT_TYPE    = "SAP_OAUTH2_GRANT_TYPE"
		SAP_OAUTH2_CLIENT_ID     = "SAP_OAUTH2_CLIENT_ID"
		SAP_OAUTH2_CLIENT_SECRET = "SAP_OAUTH2_CLIENT_SECRET"
		SAP_OAUTH2_TOKEN_URL     = "SAP_OAUTH2_TOKEN_URL"

		SAP_ACCOUNTS_HOST_SERVICE     = "SAP_ACCOUNTS_HOST_SERVICE"
		SAP_ENTITLEMENTS_HOST_SERVICE = "SAP_ENTITLEMENTS_HOST_SERVICE"
		SAP_EVENTS_HOST_SERVICE       = "SAP_EVENTS_HOST_SERVICE"
	)

	var cfg = &sap.Config{
		Endpoints: map[string]*sap.EndpointConfig{
			"accounts": {
				Host: os.Getenv(SAP_ACCOUNTS_HOST_SERVICE),
			},
			"entitlements": {
				Host: os.Getenv(SAP_ENTITLEMENTS_HOST_SERVICE),
			},
			"events": {
				Host: os.Getenv(SAP_EVENTS_HOST_SERVICE),
			},
		},

		DefaultOAuth2: &oauth2.Config{
			GrantType:    os.Getenv(SAP_OAUTH2_GRANT_TYPE),
			ClientID:     os.Getenv(SAP_OAUTH2_CLIENT_ID),
			ClientSecret: os.Getenv(SAP_OAUTH2_CLIENT_SECRET),
			TokenURL:     os.Getenv(SAP_OAUTH2_TOKEN_URL),
			Username:     os.Getenv(SAP_OAUTH2_USERNAME),
			Password:     os.Getenv(SAP_OAUTH2_PASSWORD),
		},
	}

	sessTmp, err := session.BuildFromConfig(cfg)
	if err != nil {
		panic(err)
	}
	sess = sessTmp
}
