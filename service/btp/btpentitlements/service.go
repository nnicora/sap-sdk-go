package btpentitlements

import (
	"fmt"
	"github.com/nnicora/sap-sdk-go/internal/utils"
	"net/http"
	"strings"
)

type Entitlements struct {
	dataCentersHost string

	httpClient *http.Client
}

func New(host string, httpClient *http.Client) (*Entitlements, error) {
	host = strings.TrimSuffix(host, "/")

	if !utils.IsUrl(host) {
		return nil, fmt.Errorf("invalid Entitlements url '%s'", host)
	}

	if !utils.CheckUrlSchema(host) {
		return nil, fmt.Errorf("invalid Entitlements url schema '%s'", host)
	}

	return &Entitlements{
		dataCentersHost: fmt.Sprintf("%s%s", host, "/entitlements/v1/globalAccountAllowedDataCenters"),

		httpClient: httpClient,
	}, nil
}
