package btpsaasprovisioning

import (
	"context"
	"fmt"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

const applicationOperations = "Application Operations for App Providers"

// GET /saas-manager/v1/application
// Get application registration details
type GetApplicationRegistrationInput struct {
}
type GetApplicationRegistrationOutput struct {
	//The ID of the multitenant application that is registered to the SAP SaaS Provisioning service registry.
	ServiceInstanceId string `json:"serviceInstanceId,omitempty"`
	//The unique ID of the Cloud Foundry org where the app provider has deployed and registered the
	//multitenant application.
	OrganizationGuid string `json:"organizationGuid,omitempty"`
	//The unique ID of the Cloud Foundry space where the app provider has deployed and registered the
	//multitenant application.
	SpaceGuid string `json:"spaceGuid,omitempty"`
	//The xsappname configured in the security descriptor file used to create the xsuaa service instance for the
	//multitenant application.
	XSAppName string `json:"xsappname,omitempty"`
	//The ID returned by an xsuaa service instance after the app provider has connected the multitenant
	//application to an xsuaa service instance.
	AppId string `json:"appId,omitempty"`
	//The unique registration name of the deployed multitenant application as defined by the app developer.
	AppName string `json:"appName,omitempty"`
	//The unique commercial registration name of the deployed multitenant application as defined by the app developer.
	CommercialAppName string `json:"commercialAppName,omitempty"`
	//Any callback URLs that the multitenant application exposes.
	AppUrls string `json:"appUrls,omitempty"`
	//The unique ID of the tenant that provides the multitenant application.
	ProviderTenantId string `json:"providerTenantId,omitempty"`
	//The plan used to register the multitenant application or reusable service.
	//- saasApplication: Registered entity is a multitenant application.
	//- saasService: Registered entity is a reuse service.
	AppType string `json:"appType,omitempty"`
	//The display name of the application for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`
	//The description of the multitenant application for customer-facing UIs.
	Description string `json:"description,omitempty"`
	//The category to which the application is grouped in the Subscriptions page in the cockpit.
	//If left empty, it gets assigned to the default category.
	Category string `json:"category,omitempty"`
	//ID of the global account associated with the multitenant application.
	GlobalAccountId string `json:"globalAccountId,omitempty"`
	//Name of the formations solution associated with the multitenant application.
	FormationSolutionName string `json:"formationSolutionName,omitempty"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *SaaSProvisioningV1) GetApplicationRegistration(ctx context.Context,
	input *GetApplicationRegistrationInput) (*GetApplicationRegistrationOutput, error) {
	req, out := c.getApplicationRegistrationRequest(ctx, input)
	return out, req.Send()
}
func (c *SaaSProvisioningV1) getApplicationRegistrationRequest(ctx context.Context,
	input *GetApplicationRegistrationInput) (*request.Request, *GetApplicationRegistrationOutput) {
	op := &request.Operation{
		Name: "Get Application Registration",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/application",
		},
	}

	if input == nil {
		input = &GetApplicationRegistrationInput{}
	}

	output := &GetApplicationRegistrationOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /saas-manager/v1/application/subscriptions
// Get application subscriptions
type GetApplicationSubscriptionsInput struct {
	//Get subscriptions by associated global account ID.
	GlobalAccountId string `dest:"querystring" dest-name:"globalAccountId"`
	//Get subscriptions by state.
	//Available values : IN_PROCESS, SUBSCRIBED, SUBSCRIBE_FAILED, UNSUBSCRIBE_FAILED, UPDATE_FAILED, NOT_SUBSCRIBED
	State string `dest:"querystring" dest-name:"state"`
	//Get subscriptions by the associated subaccount ID.
	SubAccountId string `dest:"querystring" dest-name:"subaccountId"`
	//Get subscriptions by tenant ID.
	TenantId string `dest:"querystring" dest-name:"tenantId"`
}
type GetApplicationSubscriptionsOutput struct {
	//Specifies the ability to use the service plan of the subscribed application. The actual amount has no bearing on
	//the maximum consumption limit of the application.
	Amount int64 `json:"amount,omitempty"`
	//The unique registration name of the deployed multitenant application, as defined by the app developer.
	AppName string `json:"appName,omitempty"`
	//The date and time the subscription was last modified. Dates and times are in UTC format.
	ChangedOn string `json:"changedOn,omitempty"`
	//A subscription code for the application.
	Code string `json:"code,omitempty"`
	//Tenant ID of the global account or subaccount of the consumer that has subscribed to the multitenant application.
	ConsumerTenantId string `json:"consumerTenantId,omitempty"`
	//The date and time the subscription was created. Dates and times are in UTC format.
	CreatedOn string `json:"createdOn,omitempty"`
	//Any reuse services used or required by a subscribed application and its services.
	Dependencies []Dependency `json:"dependencies,omitempty"`
	//Error description for the following statuses: SUBSCRIBE_FAILED, UNSUBSCRIBE_FAILED, UPDATE_FAILED.
	Error string `json:"error,omitempty"`
	//ID of the associated global account.
	GlobalAccountId string `json:"globalAccountId,omitempty"`
	//Whether the consumer tenant is active. This field is returned only if one of the following query parameters was
	//used during the API call: tenantId, subaccountId
	IsConsumerTenantActive bool `json:"isConsumerTenantActive,omitempty"`
	//The license type of the associated global account.
	LicenseType string `json:"licenseType,omitempty"`
	//The ID of the multitenant application that is registered to the SAP SaaS Provisioning registry.
	ServiceInstanceId string `json:"serviceInstanceId,omitempty"`
	//State of the subscriptions. Possible states: IN_PROCESS, SUBSCRIBED, SUBSCRIBE_FAILED, UPDATE_FAILED.
	State string `json:"state,omitempty"`
	//ID of the associated subaccount.
	SubAccountId string `json:"subaccountId,omitempty"`
	//Consumer Subdomain
	Subdomain string `json:"subdomain,omitempty"`
	//Application URL
	Url string `json:"url,omitempty"`

	//Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}
type Dependency struct {
	//The unique registration name of the linked dependency application.
	AppName string `json:"appName,omitempty"`
	//The list of relevant dependencies and their descriptions.
	Dependencies []interface{} `json:"dependencies,omitempty"`
	//In case there are errors during dependencies' assignments, the descriptions are shown here.
	Error string `json:"error,omitempty"`
	//The xsappname configured in the security descriptor file used to create the XSUAA instance.
	XSAppName string `json:"xsappname,omitempty"`
}

func (c *SaaSProvisioningV1) GetApplicationSubscriptions(ctx context.Context,
	input *GetApplicationRegistrationInput) (*GetApplicationRegistrationOutput, error) {
	req, out := c.getApplicationSubscriptionsRequest(ctx, input)
	return out, req.Send()
}
func (c *SaaSProvisioningV1) getApplicationSubscriptionsRequest(ctx context.Context,
	input *GetApplicationRegistrationInput) (*request.Request, *GetApplicationRegistrationOutput) {
	op := &request.Operation{
		Name: "Get Application Subscriptions",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/application/subscriptions",
		},
	}

	if input == nil {
		input = &GetApplicationRegistrationInput{}
	}

	output := &GetApplicationRegistrationOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /saas-manager/v1/application/tenants/{tenantId}/subscriptions
// Subscribe tenant to an application
type SubscribeTenantToApplicationInput struct {
	//The ID of the tenant to subscribe.
	TenantId string `dest:"uri" dest-name:"tenantId"`
}
type SubscribeTenantToApplicationOutput struct {
	Location string `src:"header" src-name:"Location"`

	JobStatusId string `json:"jobStatusId"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *SaaSProvisioningV1) SubscribeTenantToApplication(ctx context.Context,
	input *SubscribeTenantToApplicationInput) (*SubscribeTenantToApplicationOutput, error) {
	req, out := c.subscribeTenantToApplicationRequest(ctx, input)
	return out, req.Send()
}
func (c *SaaSProvisioningV1) subscribeTenantToApplicationRequest(ctx context.Context,
	input *SubscribeTenantToApplicationInput) (*request.Request, *SubscribeTenantToApplicationOutput) {
	op := &request.Operation{
		Name: "Subscribe Tenant To Application",
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/application/tenants/{tenantId}/subscriptions",
		},
	}

	if input == nil {
		input = &SubscribeTenantToApplicationInput{}
	}

	output := &SubscribeTenantToApplicationOutput{}
	request := c.newRequest(ctx, op, input, output)

	// TODO: This is a hack should not be use on good designed API
	wrapperBody := "{ \"jobStatusId\": \"%v\" }"
	request.ResponseBodyHandler = func(statusCode int, body []byte) ([]byte, error) {
		if statusCode == 202 {
			return []byte(fmt.Sprintf(wrapperBody, string(body))), nil
		}
		return body, nil
	}

	return request, output
}

// DELETE /saas-manager/v1/application/tenants/{tenantId}/subscriptions
// Unsubscribe tenant from an application
type UnSubscribeTenantFromApplicationInput struct {
	//The ID of the tenant to unsubscribe
	TenantId string `dest:"uri" dest-name:"tenantId"`
}
type UnSubscribeTenantFromApplicationOutput struct {
	Location string `src:"header" src-name:"Location"`

	JobStatusId string `json:"jobStatusId"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *SaaSProvisioningV1) UnSubscribeTenantFromApplication(ctx context.Context,
	input *UnSubscribeTenantFromApplicationInput) (*UnSubscribeTenantFromApplicationOutput, error) {
	req, out := c.unsubscribeTenantFromApplicationRequest(ctx, input)
	return out, req.Send()
}
func (c *SaaSProvisioningV1) unsubscribeTenantFromApplicationRequest(ctx context.Context,
	input *UnSubscribeTenantFromApplicationInput) (*request.Request, *UnSubscribeTenantFromApplicationOutput) {
	op := &request.Operation{
		Name: "UnSubscribe Tenant From Application",
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/application/tenants/{tenantId}/subscriptions",
		},
	}

	if input == nil {
		input = &UnSubscribeTenantFromApplicationInput{}
	}

	output := &UnSubscribeTenantFromApplicationOutput{}
	request := c.newRequest(ctx, op, input, output)

	// TODO: This is a hack should not be use on good designed API
	wrapperBody := "{ \"jobStatusId\": \"%v\" }"
	request.ResponseBodyHandler = func(statusCode int, body []byte) ([]byte, error) {
		if statusCode == 202 {
			return []byte(fmt.Sprintf(wrapperBody, string(body))), nil
		}
		return body, nil
	}

	return request, output
}

// PATCH /saas-manager/v1/application/tenants/{tenantId}/subscriptions
// Update subscription dependencies
type UpdateSubscriptionDependenciesInput struct {
	//The ID of the tenant for which to update dependencies
	TenantId string `dest:"uri" dest-name:"tenantId"`

	//Whether to skip updating the dependencies that haven’t changed.
	SkipUnchangedDependencies bool `dest:"querystring" dest-name:"skipUnchangedDependencies"`
	//Whether to skip updating dependencies. If set to true, updateApplicationURL must also be set to true.
	//This way, you can update the application URL without updating its dependencies.
	SkipUpdatingDependencies bool `dest:"querystring" dest-name:"skipUpdatingDependencies"`
	//Whether to update the application URL returned from the app callback. If set to true together with
	//skipUpdatingDependencies, the API call becomes synchronous.
	UpdateApplicationURL bool `dest:"querystring" dest-name:"updateApplicationURL"`

	//Send custom property values in the form of key-value pairs to dependent services (provider applications) during
	//the update to notify them about a change related to an existing subscription.
	UpdateApplicationDependencies map[string]interface{} `json:"appProviderCustomProperties"`
}
type UpdateSubscriptionDependenciesOutput struct {
	Location string `src:"header" src-name:"Location"`

	JobStatusId string `json:"jobStatusId"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *SaaSProvisioningV1) UpdateSubscriptionDependencies(ctx context.Context,
	input *UpdateSubscriptionDependenciesInput) (*UpdateSubscriptionDependenciesOutput, error) {
	req, out := c.updateSubscriptionDependenciesRequest(ctx, input)
	return out, req.Send()
}
func (c *SaaSProvisioningV1) updateSubscriptionDependenciesRequest(ctx context.Context,
	input *UpdateSubscriptionDependenciesInput) (*request.Request, *UpdateSubscriptionDependenciesOutput) {
	op := &request.Operation{
		Name: "Update Subscription Dependencies",
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/application/tenants/{tenantId}/subscriptions",
		},
	}

	if input == nil {
		input = &UpdateSubscriptionDependenciesInput{}
	}

	output := &UpdateSubscriptionDependenciesOutput{}
	request := c.newRequest(ctx, op, input, output)

	// TODO: This is a hack should not be use on good designed API
	wrapperBody := "{ \"jobStatusId\": \"%v\" }"
	request.ResponseBodyHandler = func(statusCode int, body []byte) ([]byte, error) {
		if statusCode == 202 {
			return []byte(fmt.Sprintf(wrapperBody, string(body))), nil
		}
		return body, nil
	}

	return request, output
}

// GET /saas-manager/v1/applications
// Get all entitled multitenant applications
type GetEntitledApplicationsInput struct {
	AcceptLanguage string `dest:"header" dest-name:"Accept-Language"`
}
type GetEntitledApplicationsOutput struct {
	//The response list of all the multitenant applications to which a specified subaccount is entitled to subscribe.
	Applications []Application `json:"applications,omitempty"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}
type Application struct {
	//The ID returned by XSUAA after the app provider has performed a bind of the multitenant application
	//to a XSUAA service instance.
	AppId string `json:"appId,omitempty"`
	//The unique registration name of the deployed multitenant application as defined by the app developer.
	AppName string `json:"appName,omitempty"`
	//The application coordinates provided in metadata for customer-facing UIs.
	ApplicationCoordinates map[string]interface{} `json:"applicationCoordinates,omitempty"`
	//The authentication provider of the multitenant application.
	//XSUAA is the SAP Authorization and Trust Management service that defines scopes and permissions for users as tenants at the global account level.
	//IAS is Identity Authentication Service that defines scopes and permissions for users in zones (common data isolation systems across systems, SaaS tenants, and services).
	//Enum:
	//	[ XSUAA, IAS ]
	AuthenticationProvider string `json:"authenticationProvider,omitempty"`
	//The technical name of the category defined by the app developer to which the multitenant
	//application is grouped in customer-facing UIs.
	Category string `json:"category,omitempty"`
	//The display name of the category for customer-facing UIs.
	CategoryDisplayName string `json:"categoryDisplayName,omitempty"`
	//The commercial name of the deployed multitenant application as defined by the app developer.
	CommercialAppName string `json:"commercialAppName,omitempty"`
	//The date the subscription was created. Dates and times are in UTC format.
	CreatedDate string `json:"createdDate,omitempty"`
	//Whether the application was developed by a customer. If not, then the application
	//is developed by the cloud operator, such as SAP.
	CustomerDeveloped bool `json:"customerDeveloped,omitempty"`
	//The description of the multitenant application for customer-facing UIs.
	Description string `json:"description,omitempty"`
	//The display name of the application for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`
	//Name of the formations solution associated with the multitenant application.
	FormationSolutionName string `json:"formationSolutionName,omitempty"`
	//ID of the associated global account.
	GlobalAccountId string `json:"globalAccountId,omitempty"`
	//The icon of the multitenant application for customer-facing UIs.
	IconBase64 string `json:"iconBase64,omitempty"`
	//The application's incident-tracking component provided in metadata for customer-facing UIs.
	IncidentTrackingComponent string `json:"incidentTrackingComponent,omitempty"`
	//The date the subscription was last modified. Dates and times are in UTC format.
	ModifiedDate string `json:"modifiedDate,omitempty"`
	//The plan name of the application to which the consumer has subscribed.
	PlanName string `json:"planName,omitempty"`
	//ID of the landscape-specific environment.
	PlatformEntityId string `json:"platformEntityId,omitempty"`
	//Total amount the subscribed subaccount is entitled to consume.
	Quota int32 `json:"quota,omitempty"`
	//The short description of the multitenant application for customer-facing UIs.
	ShortDescription string `json:"shortDescription,omitempty"`
	//The subscription state of the subaccount regarding the multitenant application.
	//Enum:
	//	[ IN_PROCESS, SUBSCRIBED, SUBSCRIBE_FAILED, UNSUBSCRIBE_FAILED, UPDATE_FAILED, NOT_SUBSCRIBED ]
	State string `json:"state,omitempty"`
	//The ID of the subaccount which is subscribed to the multitenant application.
	SubscribedSubAccountId string `json:"subscribedSubaccountId,omitempty"`
	//The ID of the tenant which is subscribed to a multitenant application.
	SubscribedTenantId string `json:"subscribedTenantId,omitempty"`

	SubscriptionError SubscriptionError `json:"subscriptionError,omitempty"`
	//Technical ID generated by XSUAA for a multitenant application when a consumer subscribes to the application.
	SubscriptionId string `json:"subscriptionId,omitempty"`
	//URL for app users to launch the subscribed application.
	SubscriptionUrl string `json:"subscriptionUrl,omitempty"`
	//Tenant ID of the application provider.
	TenantId string `json:"tenantId,omitempty"`
}
type SubscriptionError struct {
	//A response object that contains details about the error an app provider returns to the subscriber.
	//It contains the error code, a user-friendly, customer-oriented error message,
	//technical details about the error, and more.
	AppError string `json:"appError,omitempty"`
	//The message that describes the error that occurred during the subscription.
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (c *SaaSProvisioningV1) GetEntitledApplications(ctx context.Context,
	input *GetEntitledApplicationsInput) (*GetEntitledApplicationsOutput, error) {
	req, out := c.getEntitledApplicationsRequest(ctx, input)
	return out, req.Send()
}
func (c *SaaSProvisioningV1) getEntitledApplicationsRequest(ctx context.Context,
	input *GetEntitledApplicationsInput) (*request.Request, *GetEntitledApplicationsOutput) {
	op := &request.Operation{
		Name: "Get Entitled Applications",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/applications",
		},
	}

	if input == nil {
		input = &GetEntitledApplicationsInput{}
	}

	output := &GetEntitledApplicationsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /saas-manager/v1/applications/{appName}
// Get details of a multitenant application
type GetDetailsApplicationsInput struct {
	AcceptLanguage string `dest:"header" dest-name:"Accept-Language"`

	//The name of the multitenant application to which a subaccount is entitled to subscribe.
	AppName string `dest:"uri" dest-name:"appName"`
	//The name of the subscription plan to the multitenant application.
	PlanName string `dest:"querystring" dest-name:"planName"`
}
type GetDetailsApplicationsOutput struct {
	Application

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *SaaSProvisioningV1) GetDetailsApplications(ctx context.Context,
	input *GetDetailsApplicationsInput) (*GetDetailsApplicationsOutput, error) {
	req, out := c.getDetailsApplicationsRequest(ctx, input)
	return out, req.Send()
}
func (c *SaaSProvisioningV1) getDetailsApplicationsRequest(ctx context.Context,
	input *GetDetailsApplicationsInput) (*request.Request, *GetDetailsApplicationsOutput) {
	op := &request.Operation{
		Name: "Get Details Applications",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/applications/{appName}",
		},
	}

	if input == nil {
		input = &GetDetailsApplicationsInput{}
	}

	output := &GetDetailsApplicationsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /saas-manager/v1/applications/{appName}/subscription
// Subscribe to an application from a subaccount
type SubscribeToApplicationInput struct {
	//The name of the multitenant application to subscribe to.
	AppName string `dest:"uri" dest-name:"appName"`

	//The name of the subscription plan to a multitenant application
	PlanName string `json:"planName,omitempty"`
}
type SubscribeToApplicationOutput struct {
	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *SaaSProvisioningV1) SubscribeToApplication(ctx context.Context,
	input *SubscribeToApplicationInput) (*SubscribeToApplicationOutput, error) {
	req, out := c.getSubscribeToApplicationRequest(ctx, input)
	return out, req.Send()
}
func (c *SaaSProvisioningV1) getSubscribeToApplicationRequest(ctx context.Context,
	input *SubscribeToApplicationInput) (*request.Request, *SubscribeToApplicationOutput) {
	op := &request.Operation{
		Name: "Subscribe To Application",
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/applications/{appName}/subscription",
		},
	}

	if input == nil {
		input = &SubscribeToApplicationInput{}
	}

	output := &SubscribeToApplicationOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /saas-manager/v1/applications/{appName}/subscription
// Unsubscribe an application from a subaccount
type UnSubscribeFromApplicationInput struct {
	//The name of the multitenant application from which to unsubscribe the subaccount.
	AppName string `dest:"uri" dest-name:"appName"`
}
type UnSubscribeFromApplicationOutput struct {
	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *SaaSProvisioningV1) UnSubscribeFromApplication(ctx context.Context,
	input *UnSubscribeFromApplicationInput) error {
	req, _ := c.unSubscribeFromApplicationRequest(ctx, input)
	return req.Send()
}
func (c *SaaSProvisioningV1) unSubscribeFromApplicationRequest(ctx context.Context,
	input *UnSubscribeFromApplicationInput) (*request.Request, *UnSubscribeFromApplicationOutput) {
	op := &request.Operation{
		Name: "UnSubscribe From Application",
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/applications/{appName}/subscription",
		},
	}

	if input == nil {
		input = &UnSubscribeFromApplicationInput{}
	}

	output := &UnSubscribeFromApplicationOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PUT /saas-manager/v1/subscription-callback/{identifier}/result
// Subscribe a subaccount tenant to an application
type SubscribeSubAccountTenantToApplicationInput struct {
	//Unique identifier of the current subscription job.
	Identifier string `dest:"uri" dest-name:"identifier"`

	//Additional details accompanying the subscription process. Relates mostly to the
	//cases when the subscription process status is FAILED.
	Message string `json:"message,omitempty"`
	//Status of the subscription job.
	//Enum:
	//	[ SUCCEEDED, FAILED ]
	Status string `json:"status,omitempty"`
	//The URL the multitenant application is exposing for a subscription.
	SubscriptionUrl string `json:"subscriptionUrl,omitempty"`
}
type SubscribeSubAccountTenantToApplicationOutput struct {
	JobStatusId string `json:"jobStatusId"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *SaaSProvisioningV1) SubscribeSubAccountTenantToApplication(ctx context.Context,
	input *SubscribeSubAccountTenantToApplicationInput) error {
	req, _ := c.subscribeSubAccountTenantToApplicationRequest(ctx, input)
	return req.Send()
}
func (c *SaaSProvisioningV1) subscribeSubAccountTenantToApplicationRequest(ctx context.Context,
	input *SubscribeSubAccountTenantToApplicationInput) (*request.Request, *UnSubscribeFromApplicationOutput) {
	op := &request.Operation{
		Name: "Subscribe SubAccount Tenant To Application",
		Http: request.HTTP{
			Method: request.PUT,
			Path:   "/applications/{appName}/subscription",
		},
	}

	if input == nil {
		input = &SubscribeSubAccountTenantToApplicationInput{}
	}

	output := &UnSubscribeFromApplicationOutput{}
	request := c.newRequest(ctx, op, input, output)

	// TODO: This is a hack should not be use on good designed API
	wrapperBody := "{ \"jobStatusId\": \"%v\" }"
	request.ResponseBodyHandler = func(statusCode int, body []byte) ([]byte, error) {
		if statusCode == 202 {
			return []byte(fmt.Sprintf(wrapperBody, string(body))), nil
		}
		return body, nil
	}

	return request, output
}
