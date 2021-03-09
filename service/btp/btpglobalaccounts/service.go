package btpglobalaccounts

import (
	"fmt"
	"net/http"
	"strings"
)

type GlobalAccounts struct {
	globalAccountsHost string

	httpClient *http.Client
}

func New(host string, httpClient *http.Client) *GlobalAccounts {
	host = strings.TrimSuffix(host, "/")
	return &GlobalAccounts{
		globalAccountsHost: fmt.Sprintf("%s%s", host, "/accounts/v1/globalAccount"),

		httpClient: httpClient,
	}
}
