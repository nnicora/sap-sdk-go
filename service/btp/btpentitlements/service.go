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

	if ok, err := utils.IsUrl(host); err != nil {
		return nil, fmt.Errorf("invalid Entitlements host; %v", err)
	} else if !ok {
		return nil, fmt.Errorf("invalid Entitlements host '%s'", host)
	}

	return &Entitlements{
		dataCentersHost: fmt.Sprintf("%s%s", host, "/entitlements/v1/globalAccountAllowedDataCenters"),

		httpClient: httpClient,
	}, nil
}
