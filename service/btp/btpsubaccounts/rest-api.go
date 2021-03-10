package btpsubaccounts

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nnicora/sap-sdk-go/internal/utils"
	"github.com/nnicora/sap-sdk-go/sap/times"
	"github.com/nnicora/sap-sdk-go/service"
	"io/ioutil"
	"net/http"
)

type SubAccount struct {
	BetaEnabled       bool               `json:"betaEnabled"`
	CreatedBy         string             `json:"createdBy"`
	CreatedDate       times.JavaTime     `json:"createdDate"`
	CustomProperties  []CustomProperties `json:"customProperties"`
	Description       string             `json:"description"`
	DisplayName       string             `json:"displayName"`
	GlobalAccountGUID string             `json:"globalAccountGUID"`
	Guid              string             `json:"guid"`
	ModifiedDate      times.JavaTime     `json:"modifiedDate"`
	ParentFeatures    []string           `json:"parentFeatures"`
	ParentGUID        string             `json:"parentGUID"`
	Region            string             `json:"region"`
	State             string             `json:"state"`
	StateMessage      string             `json:"stateMessage"`
	Subdomain         string             `json:"subdomain"`
	UsedForProduction string             `json:"usedForProduction"`
	ZoneId            string             `json:"zoneId"`
}

type CustomProperties struct {
	KeyValue
	AccountGUID string `json:"accountGUID"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type subAccountResponse struct {
	Value []SubAccount `json:"value"`
}

// GET /accounts/v1/subaccounts
// Get all subaccounts
func (a *SubAccounts) GetAllRestApi(ctx context.Context) ([]SubAccount, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", a.subAccountsHost, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
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

	obj := &subAccountResponse{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}
	return obj.Value, nil
}

type SubAccountRequest struct {
	BetaEnabled       bool       `json:"betaEnabled"`
	CustomProperties  []KeyValue `json:"customProperties"`
	Description       string     `json:"description"`
	DisplayName       string     `json:"displayName"`
	Origin            string     `json:"origin"`
	ParentGUID        string     `json:"parentGUID"`
	Region            string     `json:"region"`
	SubaccountAdmins  []string   `json:"subaccountAdmins"`
	Subdomain         string     `json:"subdomain"`
	UsedForProduction string     `json:"usedForProduction"`
}

// POST accounts/v1/subaccounts
// Create a subaccount
func (a *SubAccounts) CreateRestApi(ctx context.Context, req *SubAccountRequest) (*SubAccount, error) {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(req); err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", a.subAccountsHost, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
	if err != nil {
		return nil, utils.GetDnsError(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, errors.New(string(responseData))
	}

	obj := &SubAccount{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type CloneSubAccountRequest struct {
	BetaEnabled         bool       `json:"betaEnabled"`
	CloneConfigurations []string   `json:"cloneConfigurations"`
	CustomProperties    []KeyValue `json:"customProperties"`
	Description         string     `json:"description"`
	DisplayName         string     `json:"displayName"`
	Origin              string     `json:"origin"`
	Region              string     `json:"region"`
	SubaccountAdmins    []string   `json:"subaccountAdmins"`
	Subdomain           string     `json:"subdomain"`
	UsedForProduction   string     `json:"usedForProduction"`
}

// POST /accounts/v1/subaccounts/clone/{sourceSubaccountGUID}
// Clone a Neo subaccount
func (a *SubAccounts) CloneRestApi(ctx context.Context, sourceSubAccountGUID string, req *CloneSubAccountRequest) (*SubAccount, error) {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(req); err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf(a.cloneSubAccountsHost, sourceSubAccountGUID), body)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
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

	obj := &SubAccount{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// GET /accounts/v1/subaccounts/{subaccountGUID}
// Get a subaccount
func (a *SubAccounts) GetRestApi(ctx context.Context, subAccountGUID string) (*SubAccount, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/%s", a.subAccountsHost, subAccountGUID), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
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

	obj := &SubAccount{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// DELETE /accounts/v1/subaccounts/{subaccountGUID}
// Delete a subaccount
func (a *SubAccounts) DeleteRestApi(ctx context.Context, subAccountGUID string) (*SubAccount, error) {
	sa, err := a.GetRestApi(ctx, subAccountGUID)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf("%s/%s", a.subAccountsHost, subAccountGUID), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
	if err != nil {
		return nil, utils.GetDnsError(err)
	}

	if resp.StatusCode != 200 {
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(responseData))
	}

	return sa, nil
}

type UpdateSubAccountRequest struct {
	BetaEnabled       bool                         `json:"betaEnabled"`
	CustomProperties  []UpdateSubAccountProperties `json:"customProperties"`
	Description       string                       `json:"description"`
	DisplayName       string                       `json:"displayName"`
	UsedForProduction string                       `json:"usedForProduction"`
}
type UpdateSubAccountProperties struct {
	KeyValue
	Delete bool `json:"delete"`
}

// PATCH /accounts/v1/subaccounts/{subaccountGUID}
// Update a subaccount
func (a *SubAccounts) PatchRestApi(ctx context.Context, subAccountGUID string, req *UpdateSubAccountRequest) (*SubAccount, error) {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(req); err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "PATCH", fmt.Sprintf("%s/%s", a.subAccountsHost, subAccountGUID), body)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
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

	obj := &SubAccount{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// GET /accounts/v1/subaccounts/{subaccountGUID}/customProperties
// Get custom properties for a subaccount
func (a *SubAccounts) GetCustomPropertiesRestApi(ctx context.Context, subAccountGUID string) ([]CustomProperties, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(a.customPropertiesSubAccountsHost, subAccountGUID), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
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

	type customPropertiesResponse struct {
		value []CustomProperties `json:"value"`
	}

	obj := &customPropertiesResponse{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}
	return obj.value, nil
}

type MoveSubAccountRequest struct {
	TargetAccountGUID string `json:"targetAccountGUID"`
}

type MoveSubAccountsRequest struct {
	subaccountsToMove []MoveSubAccountsRequestPayload `json:"subaccountsToMove"`
}

type MoveSubAccountsRequestPayload struct {
	SourceGuid      string   `json:"sourceGuid"`
	SubaccountGuids []string `json:"subaccountGuids"`
	TargetGuid      string   `json:"targetGuid"`
}

// POST /accounts/v1/subaccounts/move
// Batch move subaccounts
func (a *SubAccounts) MoveManyRestApi(ctx context.Context, req *MoveSubAccountsRequest) (*SubAccount, error) {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(req); err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", a.moveSubAccountHost, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
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

	obj := &SubAccount{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

// POST /accounts/v1/subaccounts/{subaccountGUID}/move
// Move a subaccount
func (a *SubAccounts) MoveRestApi(ctx context.Context, subAccountGUID string, req *MoveSubAccountRequest) (*SubAccount, error) {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(req); err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf(a.moveSubAccountsHost, subAccountGUID), body)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
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

	obj := &SubAccount{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}

	return obj, nil
}

type ServiceManagementBinding struct {
	Description  string `json:"description"`
	ClientId     string `json:"clientid"`
	ClientSecret string `json:"clientsecret"`
	SMUrl        string `json:"sm_url"`
	Url          string `json:"url"`
	XsAppName    string `json:"xsappname"`
}

// GET /accounts/v1/subaccounts/{subaccountGUID}/serviceManagementBinding
// Get a Service Management binding
func (a *SubAccounts) GetServiceManagementBindingRestApi(ctx context.Context, subAccountGUID string) (*ServiceManagementBinding, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf(a.serviceManagementBindingSubAccountsHost, subAccountGUID), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
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

	obj := &ServiceManagementBinding{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// POST /accounts/v1/subaccounts/{subaccountGUID}/serviceManagementBinding
// Create a Service Management binding
func (a *SubAccounts) CreateServiceManagementBindingRestApi(ctx context.Context, subAccountGUID string) (*ServiceManagementBinding, error) {
	request, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf(a.serviceManagementBindingSubAccountsHost, subAccountGUID), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
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

	obj := &ServiceManagementBinding{}
	if err := json.Unmarshal(responseData, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// DELETE /accounts/v1/subaccounts/{subaccountGUID}/serviceManagementBinding
// Delete a Service Management binding
func (a *SubAccounts) DeleteServiceManagementBindingRestApi(ctx context.Context, subAccountGUID string) (bool, error) {
	request, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf(a.serviceManagementBindingSubAccountsHost, subAccountGUID), nil)
	if err != nil {
		return false, err
	}
	request.Header.Set(service.ContentType, service.JSON)

	resp, err := a.httpClient.Do(request)
	if err != nil {
		return false, utils.GetDnsError(err)
	}

	if resp.StatusCode != 200 {
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}
		return false, errors.New(string(responseData))
	}

	return resp.StatusCode == 200, nil
}
