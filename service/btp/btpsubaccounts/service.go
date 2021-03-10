package btpsubaccounts

import (
	"fmt"
	"github.com/nnicora/sap-sdk-go/internal/utils"
	"net/http"
	"strings"
)

type SubAccounts struct {
	subAccountsHost                         string
	cloneSubAccountsHost                    string
	moveSubAccountHost                      string
	moveSubAccountsHost                     string
	customPropertiesSubAccountsHost         string
	serviceManagementBindingSubAccountsHost string

	httpClient *http.Client
}

func New(host string, httpClient *http.Client) (*SubAccounts, error) {
	host = strings.TrimSuffix(host, "/")

	if !utils.IsUrl(host) {
		return nil, fmt.Errorf("invalid Sub Accounts url '%s'", host)
	}

	if !utils.CheckUrlSchema(host) {
		return nil, fmt.Errorf("invalid Sub Accounts url schema '%s'", host)
	}

	return &SubAccounts{
		subAccountsHost:                         fmt.Sprintf("%s%s", host, "/accounts/v1/subaccounts"),
		cloneSubAccountsHost:                    fmt.Sprintf("%s%s", host, "/accounts/v1/subaccounts/clone/%s"),
		moveSubAccountHost:                      fmt.Sprintf("%s%s", host, "/accounts/v1/subaccounts/move"),
		moveSubAccountsHost:                     fmt.Sprintf("%s%s", host, "/accounts/v1/subaccounts/%s/move"),
		customPropertiesSubAccountsHost:         fmt.Sprintf("%s%s", host, "/accounts/v1/subaccounts/%s/customProperties"),
		serviceManagementBindingSubAccountsHost: fmt.Sprintf("%s%s", host, "/accounts/v1/subaccounts/%s/serviceManagementBinding"),

		httpClient: httpClient,
	}, nil
}
