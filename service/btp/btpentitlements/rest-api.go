package btpentitlements

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nnicora/sap-sdk-go/internal/utils"
	"github.com/nnicora/sap-sdk-go/sap/http/headerkey"
	"github.com/nnicora/sap-sdk-go/sap/http/mimetype"
	"io/ioutil"
	"net/http"
)

type DataCenter struct {
	Name                   string `json:"name"`
	DisplayName            string `json:"displayName"`
	Region                 string `json:"region"`
	Environment            string `json:"environment"`
	IaasProvider           string `json:"iaasProvider"`
	SupportsTrial          bool   `json:"supportsTrial"`
	ProvisioningServiceUrl string `json:"provisioningServiceUrl"`
	SaasRegistryServiceUrl string `json:"saasRegistryServiceUrl"`
	Domain                 string `json:"domain"`
}

// GET /entitlements/v1/globalAccountAllowedDataCenters
// Get available data centers
type dataCentersResponse struct {
	DataCenters []DataCenter `json:"datacenters"`
}

func (e *Entitlements) GetDataCentersRestApi(ctx context.Context) ([]DataCenter, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", e.dataCentersHost, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(headerkey.ContentType, mimetype.ApplicationJson)

	resp, err := e.httpClient.Do(request)
	if err != nil {
		return nil, utils.GetDnsError(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(responseData))
	}

	obj := &dataCentersResponse{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}
	return obj.DataCenters, nil
}

func (e *Entitlements) GetProvidersRegionsRestApi(ctx context.Context) (map[string][]string, error) {
	dcs, err := e.GetDataCentersRestApi(ctx)
	if err != nil {
		return nil, err
	}

	providers := make(map[string][]string, 0)
	for _, dc := range dcs {
		provider := providers[dc.IaasProvider]
		if provider == nil {
			providers[dc.IaasProvider] = make([]string, 0)
		}
		providers[dc.IaasProvider] = append(providers[dc.IaasProvider], dc.Region)
	}
	return providers, nil
}

func (e *Entitlements) GetProviderRegionsRestApi(ctx context.Context, provider string) ([]string, error) {
	providers, err := e.GetProvidersRegionsRestApi(ctx)
	if err != nil {
		return nil, err
	}
	return providers[provider], nil
}
