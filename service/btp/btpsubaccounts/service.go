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

	if ok, err := utils.IsUrl(host); err != nil {
		return nil, fmt.Errorf("invalid Sub Accounts host; %v", err)
	} else if !ok {
		return nil, fmt.Errorf("invalid Sub Accounts host '%s'", host)
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
