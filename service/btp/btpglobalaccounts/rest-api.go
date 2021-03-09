package btpglobalaccounts

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nnicora/sap-sdk-go/sap/times"
	"github.com/nnicora/sap-sdk-go/service"
	"github.com/nnicora/sap-sdk-go/service/btp/btpsubaccounts"
	"io/ioutil"
	"net/http"
)

type LegalLinks struct {
	privacy string `json:"privacy"`
}
type Children struct {
	Children          []Children                        `json:"children"`
	ContractStatus    string                            `json:"contractStatus"`
	CreatedBy         string                            `json:"createdBy"`
	CreatedDate       times.JavaTime                    `json:"createdDate"`
	CustomProperties  []btpsubaccounts.CustomProperties `json:"customProperties"`
	Description       string                            `json:"description"`
	DirectoryFeatures []string                          `json:"directoryFeatures"`
	DisplayName       string                            `json:"displayName"`
	EntityState       string                            `json:"entityState"`
	Guid              string                            `json:"guid"`
	LegalLinks        LegalLinks                        `json:"legalLinks"`
	ModifiedDate      times.JavaTime                    `json:"modifiedDate"`
	ParentGuid        string                            `json:"parentGuid"`
	StateMessage      string                            `json:"stateMessage"`
	Subaccounts       []btpsubaccounts.SubAccount       `json:"subaccounts"`
	Subdomain         string                            `json:"subdomain"`
}

type GlobalAccount struct {
	Children                      []Children                        `json:"children"`
	CommercialModel               string                            `json:"commercialModel"`
	ConsumptionBased              bool                              `json:"consumptionBased"`
	ContractStatus                string                            `json:"contractStatus"`
	CostCenter                    string                            `json:"costCenter"`
	CreatedDate                   times.JavaTime                    `json:"createdDate"`
	CrmCustomerId                 string                            `json:"crmCustomerId"`
	CrmTenantId                   string                            `json:"crmTenantId"`
	CustomProperties              []btpsubaccounts.CustomProperties `json:"customProperties"`
	Description                   string                            `json:"description"`
	DisplayName                   string                            `json:"displayName"`
	EntityState                   string                            `json:"entityState"`
	ExpiryDate                    times.JavaTime                    `json:"expiryDate"`
	GeoAccess                     string                            `json:"geoAccess"`
	Guid                          string                            `json:"guid"`
	LegalLinks                    LegalLinks                        `json:"legalLinks"`
	LicenseType                   string                            `json:"licenseType"`
	ModifiedDate                  times.JavaTime                    `json:"modifiedDate"`
	Origin                        string                            `json:"origin"`
	ParentGuid                    string                            `json:"parentGuid"`
	ParentType                    string                            `json:"parentType"`
	RenewalDate                   times.JavaTime                    `json:"renewalDate"`
	ServiceId                     string                            `json:"serviceId"`
	StateMessage                  string                            `json:"stateMessage"`
	Subaccounts                   []btpsubaccounts.SubAccount       `json:"subaccounts"`
	Subdomain                     string                            `json:"subdomain"`
	TerminationNotificationStatus string                            `json:"terminationNotificationStatus"`
	UseFor                        string                            `json:"useFor"`
}

type GlobalAccountParams struct {
	DerivedAuthorizations string
	Expand                bool
}

// GET /accounts/v1/globalAccount
// Get a global account
func (ga *GlobalAccounts) GetRestApiWithParams(ctx context.Context, params *GlobalAccountParams) (*GlobalAccount, error) {
	reqUrl := ga.globalAccountsHost
	if params != nil {
		if len(params.DerivedAuthorizations) > 0 && params.Expand {
			reqUrl = fmt.Sprintf("%s?derivedAuthorizations=%s&expand=%v",
				ga.globalAccountsHost, params.DerivedAuthorizations, params.Expand)
		} else if params.Expand {
			reqUrl = fmt.Sprintf("%s?expand=%v", reqUrl, params.Expand)
		}
	}

	request, err := http.NewRequestWithContext(ctx, "GET", reqUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := ga.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	obj := &GlobalAccount{}
	if err := json.Unmarshal(body, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

type GlobalAccountRequest struct {
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
}

// PATCH /accounts/v1/globalAccount
// Update a global account
func (ga *GlobalAccounts) UpdateRestApi(ctx context.Context, req *GlobalAccountRequest) (*GlobalAccount, error) {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(req); err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "PATCH", ga.globalAccountsHost, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := ga.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	obj := &GlobalAccount{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}
	return obj, nil
}
