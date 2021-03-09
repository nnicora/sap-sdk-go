package btpentitlements

import (
	"fmt"
	"net/http"
	"strings"
)

type Entitlements struct {
	dataCentersHost string

	httpClient *http.Client
}

func New(host string, httpClient *http.Client) *Entitlements {
	host = strings.TrimSuffix(host, "/")
	return &Entitlements{
		dataCentersHost: fmt.Sprintf("%s%s", host, "/entitlements/v1/globalAccountAllowedDataCenters"),

		httpClient: httpClient,
	}
}
