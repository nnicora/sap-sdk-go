package btpglobalaccounts

import (
	"fmt"
	"github.com/nnicora/sap-sdk-go/internal/utils"
	"net/http"
	"strings"
)

type GlobalAccounts struct {
	globalAccountsHost string

	httpClient *http.Client
}

func New(host string, httpClient *http.Client) (*GlobalAccounts, error) {
	host = strings.TrimSuffix(host, "/")

	if ok, err := utils.IsUrl(host); err != nil {
		return nil, fmt.Errorf("invalid Global Accounts host; %v", err)
	} else if !ok {
		return nil, fmt.Errorf("invalid Global Accounts host '%s'", host)
	}

	return &GlobalAccounts{
		globalAccountsHost: fmt.Sprintf("%s%s", host, "/accounts/v1/globalAccount"),

		httpClient: httpClient,
	}, nil
}
