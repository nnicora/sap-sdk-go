package btpentitlements

import (
	"context"
	"fmt"
	"github.com/nnicora/sap-sdk-go/internal/times"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

// GET /entitlements/v1/globalAccountAllowedDataCenters
// Get available data centers
type DataCentersInput struct {
	AcceptLanguage string `dest:"header" dest-name:"Accept-Language"`

	//Region for which to get data centers.
	Region string `dest:"querystring" dest-name:"region"`
}
type DataCentersOutput struct {
	//Contains information about the available data centers for a specified global account.
	DataCenters []DataCenter `json:"datacenters,omitempty"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}
type DataCenter struct {
	//Technical name of the data center. Must be unique within the cloud deployment.
	Name string `json:"name,omitempty"`

	//Descriptive name of the data center for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`

	//The region in which the data center is located.
	Region string `json:"region,omitempty"`

	//The environment that the data center supports. For example: Kubernetes, Cloud Foundry.
	Environment string `json:"environment,omitempty"`

	//The infrastructure provider for the data center. Valid values:
	//
	//AWS: Amazon Web Services.
	//GCP: Google Cloud Platform.
	//AZURE: Microsoft Azure.
	//SAP: SAP BTP (Neo).
	//ALI: Alibaba Cloud.
	//IBM: IBM Cloud.
	//Enum:
	//	[ AWS, GCP, AZURE, SAP, ALI, IBM ]
	IaasProvider string `json:"iaasProvider,omitempty"`

	//Whether the specified datacenter supports trial accounts.
	SupportsTrial bool `json:"supportsTrial,omitempty"`

	//Provisioning service URL.
	ProvisioningServiceUrl string `json:"provisioningServiceUrl,omitempty"`

	//Saas-Registry service URL.
	SaasRegistryServiceUrl string `json:"saasRegistryServiceUrl,omitempty"`

	//The domain of the data center
	Domain string `json:"domain,omitempty"`
}

func (c *EntitlementsV1) GetDataCenters(ctx context.Context) (*DataCentersOutput, error) {
	req, out := c.getDataCentersRequest(ctx, nil)
	return out, req.Send()
}
func (c *EntitlementsV1) getDataCentersRequest(ctx context.Context, input *DataCentersInput) (*request.Request, *DataCentersOutput) {
	op := &request.Operation{
		Name: "Regions for Global Account",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/globalAccountAllowedDataCenters",
		},
	}

	if input == nil {
		input = &DataCentersInput{}
	}

	output := &DataCentersOutput{}
	return c.newRequest(ctx, op, input, output), output
}

func (e *EntitlementsV1) GetProvidersRegions(ctx context.Context) (map[string][]string, error) {
	dcs, err := e.GetDataCenters(ctx)
	if err != nil {
		return nil, err
	}

	providers := make(map[string][]string, 0)
	for _, dc := range dcs.DataCenters {
		provider := providers[dc.IaasProvider]
		if provider == nil {
			providers[dc.IaasProvider] = make([]string, 0)
		}
		providers[dc.IaasProvider] = append(providers[dc.IaasProvider], dc.Region)
	}
	return providers, nil
}
func (e *EntitlementsV1) GetProviderRegions(ctx context.Context, provider string) ([]string, error) {
	providers, err := e.GetProvidersRegions(ctx)
	if err != nil {
		return nil, err
	}
	return providers[provider], nil
}

// GET /entitlements/v1/globalAccountAssignments
// Get available data centers
type GlobalAccountAssignmentsInput struct {
	AcceptLanguage string `dest:"header" dest-name:"Accept-Language"`

	//Specify if to include also services that are automatically assigned to a subaccount when the subaccount is created.
	//	Default is false.
	IncludeAutoManagedPlans bool `dest:"querystring" dest-name:"includeAutoManagedPlans"`

	//Use the parameter to specify for which subaccount to view assigned entitlements.
	//If left empty, the API returns the entitlements for the global account and all its subaccounts.
	SubAccountGuid string `dest:"querystring" dest-name:"subaccountGUID"`
}
type GlobalAccountAssignmentsOutput struct {
	//Services entitled to global account, its directories and subaccounts.
	EntitledServices []EntitledService `json:"entitledServices,omitempty"`

	//The list of services that are assigned to subaccounts located under a global account.
	AssignedServices []AssignedService `json:"assignedServices,omitempty"`

	//Whether the External Provider Registry (XPR) is available.
	FetchErrorFromExternalProviderRegistry bool `json:"fetchErrorFromExternalProviderRegistry,omitempty"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *EntitlementsV1) GetGlobalAccountAssignments(ctx context.Context, input *GlobalAccountAssignmentsInput) (*GlobalAccountAssignmentsOutput, error) {
	req, out := c.getGlobalAccountAssignmentsRequest(ctx, input)
	return out, req.Send()
}
func (c *EntitlementsV1) getGlobalAccountAssignmentsRequest(ctx context.Context, input *GlobalAccountAssignmentsInput) (*request.Request, *GlobalAccountAssignmentsOutput) {
	op := &request.Operation{
		Name: "Get Global Account Assignments",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/globalAccountAssignments",
		},
	}

	if input == nil {
		input = &GlobalAccountAssignmentsInput{}
	}

	output := &GlobalAccountAssignmentsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /entitlements/v1/assignments
// Get all the entitlements and quota assignments
type GetAssignmentsInput struct {
	//The ID of the directory for which to show the entitlements and quota assignments.
	DirectoryGuid string `dest:"querystring" dest-name:"directoryGUID"`

	//Specify if to include also services that are automatically entitled to a global account when the global account is created.
	//Default is false.
	IncludeAutoManagedPlans bool `dest:"querystring" dest-name:"includeAutoManagedPlans"`

	//The ID of the subaccount for which to show the entitlements and quota assignments.
	SubAccountGuid string `dest:"querystring" dest-name:"subaccountGUID"`
}
type GetAssignmentsOutput struct {
	//Services entitled to global account, its directories and subaccounts.
	EntitledServices []EntitledService `json:"entitledServices,omitempty"`

	//The list of services that are assigned to subaccounts located under a global account.
	AssignedServices []AssignedService `json:"assignedServices,omitempty"`

	//Whether the External Provider Registry (XPR) is available.
	FetchErrorFromExternalProviderRegistry bool `json:"fetchErrorFromExternalProviderRegistry,omitempty"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}
type AssignedService struct {
	//The unique registration name of the deployed service as defined by the service provider.
	Name string `json:"name,omitempty"`

	//Display name of the service for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`

	BusinessCategory BusinessCategory `json:"businessCategory,omitempty"`

	//List of service plans associated with the assigned service.
	ServicePlans []AssignedServicePlan `json:"servicePlans,omitempty"`

	//The icon of the service in Base64 format.
	IconBase64 string `json:"iconBase64,omitempty"`

	//The application coordinates provided in metadata.
	ApplicationCoordinates interface{} `json:"applicationCoordinates,omitempty"`
}
type EntitledService struct {
	//The unique registration name of the deployed service as defined by the service provider.
	Name string `json:"name,omitempty"`

	//Display name of the service for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`

	//Description of the service for customer-facing UIs.
	Description      string           `json:"description,omitempty"`
	BusinessCategory BusinessCategory `json:"businessCategory,omitempty"`

	//The owner type of the service. Possible values:
	//
	//VENDOR: The owner is a service owner, who is affiliated with the cloud operator, that added the service to the product catalog for general consumption.
	//CUSTOMER: The owner is an SAP customer that added a custom service to the product catalog, and it is available only for consumption within the customer's global account.
	//PARTNER: The owner is an SAP partner that added the service to the product catalog, and it is available only to their customers for consumption.
	//Enum:
	//	[ VENDOR, CUSTOMER, PARTNER ]
	OwnerType string `json:"ownerType,omitempty"`

	//List of service plans associated with the entitled service.
	ServicePlans []ServicePlan `json:"servicePlans,omitempty"`

	//The icon of the service in Base64 format.
	IconBase64 string `json:"iconBase64,omitempty"`

	//The application coordinates provided in metadata.
	ApplicationCoordinates interface{} `json:"applicationCoordinates,omitempty"`
}
type BusinessCategory struct {
	//Unique ID of the business category.
	Id string `json:"id,omitempty"`

	//Display name of the business category for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`
}
type AssignedServicePlan struct {
	//The unique registration name of the service plan.
	Name string `json:"name,omitempty"`

	//The name of the service plan for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`

	//A unique identifier for service plans that can distinguish between the same service plans with different pricing plans.
	UniqueIdentifier string `json:"uniqueIdentifier,omitempty"`

	//The type of service offering. Possible values:
	//
	//PLATFORM: A service required for using a specific platform; for example, Application Runtime is required for the
	//	Cloud Foundry platform.
	//SERVICE: A commercial or technical service. that has a numeric quota (amount) when entitled or assigned to a resource.
	//	When assigning entitlements of this type, use the 'amount' option instead of 'enable'.
	//	See: PUT/entitlements/v1/directories/{directoryGUID}/assignments.
	//ELASTIC_SERVICE: A commercial or technical service that has no numeric quota (amount) when entitled or assigned
	//	to a resource. Generally this type of service can be as many times as needed when enabled, but may in some
	//	cases be restricted by the service owner. When assigning entitlements of this type, use the 'enable' option
	//	instead of 'amount'. See: PUT/entitlements/v1/directories/{directoryGUID}/assignments.
	//ELASTIC_LIMITED: An elastic service that can be enabled for only one subaccount per global account.
	//APPLICATION: A multitenant application to which consumers can subscribe. As opposed to applications defined as a
	//	'QUOTA_BASED_APPLICATION', these applications do not have a numeric quota and are simply enabled or disabled
	//	as entitlements per subaccount.
	//QUOTA_BASED_APPLICATION: A multitenant application to which consumers can subscribe. As opposed to applications
	//	defined as 'APPLICATION', these applications have an numeric quota that limits consumer usage of the subscribed
	//		application per subaccount. When maxAllowedSubaccountQuota is > 0, this is the limit that can be set when
	//		assigning the max quota entitlement of the app to any subaccount. If maxAllowedSubaccountQuota is = 0 or null,
	//		the max quota that can be entitled to any subaccount is the amount purchased by the customer (the global account quota).
	//ENVIRONMENT: An environment service; for example, Cloud Foundry.
	//Enum:
	//	[ PLATFORM, SERVICE, ELASTIC_SERVICE, ELASTIC_LIMITED, APPLICATION, QUOTA_BASED_APPLICATION, ENVIRONMENT ]
	Category string `json:"category,omitempty"`

	//Whether the service plan is a beta feature.
	Beta bool `json:"beta,omitempty"`

	//The maximum allowed usage quota per subaccount for multitenant applications and environments that are defined as
	//"quota-based". This quota limits the usage of the application and/or environment per subaccount per a given
	//usage metric that is defined within the application or environment by the service provider. If null, the usage
	//limit per subaccount is the maximum free quota in the global account.
	//For example, a value of 1 could:
	//	(1) limit the number of subscriptions to a quota-based multitenant application within a global account according
	//		to the purchased quota, or
	//	(2) restrict the enablement of a single instance of an environment per subaccount.
	MaxAllowedSubAccountQuota int32 `json:"maxAllowedSubaccountQuota,omitempty"`

	//Is the quota of this service plan entitled to the global account with unlimited usage.
	Unlimited bool `json:"unlimited,omitempty"`

	//Assignment detailed information
	AssignmentInfo []AssignedServicePlanSubAccount `json:"assignmentInfo,omitempty"`
}
type AssignedServicePlanSubAccount struct {
	//Specifies if the plan was automatically assigned regardless of any action by an admin. This applies to entitlements
	//that are always available to subaccounts and cannot be removed.
	AutoAssigned bool `json:"autoAssigned,omitempty"`

	//example: GUID of GLOBAL_ACCOUNT or SUBACCOUNT
	//The unique ID of the global account or directory to which the entitlement is assigned.
	EntityId              string  `json:"entityId,omitempty"`
	ParentAmount          float32 `json:"parentAmount,omitempty"`
	ParentId              string  `json:"parentId,omitempty"`
	ParentRemainingAmount float32 `json:"parentRemainingAmount,omitempty"`

	//Enum:
	//	[ SUBACCOUNT, GLOBAL_ACCOUNT, DIRECTORY ]
	ParentType string `json:"parentType,omitempty"`

	//example: GLOBAL_ACCOUNT or SUBACCOUNT
	//The type of entity to which the entitlement is assigned.
	//
	//SUBACCOUNT: The entitlement is assigned to a subaccount.
	//GLOBAL_ACCOUNT: The entitlement is assigned to a root global account.
	//DIRECTORY: The entitlement is assigned to a directory.
	//Enum:
	//	[ SUBACCOUNT, GLOBAL_ACCOUNT, DIRECTORY ]
	EntityType string `json:"entityType,omitempty"`

	//The quantity of the entitlement that is assigned to the root global account or directory.
	Amount float32 `json:"amount,omitempty"`

	//The requested amount when it is different from the actual amount because the request state is still in process or failed.
	RequestedAmount float32 `json:"requestedAmount,omitempty"`

	//The current state of the service plan assignment.
	//
	//STARTED: CRUD operation on an entity has started.
	//PROCESSING: A series of operations related to the entity is in progress.
	//PROCESSING_FAILED: The processing operations failed.
	//OK: The CRUD operation or series of operations completed successfully.
	//Enum:
	//	[ STARTED, PROCESSING, PROCESSING_FAILED, OK ]
	EntityState string `json:"entityState,omitempty"`

	//Information about the current state.
	StateMessage string `json:"stateMessage,omitempty"`

	//Whether the plan is automatically distributed to the subaccounts that are located in the directory.
	AutoAssign bool `json:"autoAssign,omitempty"`

	//The amount of the entitlement to automatically assign to subaccounts that are added in the future to the
	//entitlement's assigned directory.
	//Requires that autoAssign is set to TRUE, and there is remaining quota for the entitlement. To automatically
	//distribute to subaccounts that are added in the future to the directory, distribute must be set to TRUE.
	AutoDistributeAmount int32 `json:"autoDistributeAmount,omitempty"`

	//Date the subaccount has been created. Dates and times are in UTC format.
	CreatedDate times.JavaTime `json:"createdDate,omitempty"`

	//Date the subaccount has been modified. Dates and times are in UTC format.
	ModifiedDate times.JavaTime `json:"modifiedDate,omitempty"`

	//Global account resource details
	Resources []Resource `json:"resources,omitempty"`

	//True, if an unlimited quota of this service plan assigned to the directory or subaccount in the global account.
	//False, if the service plan is assigned to the directory or subaccount with a limited numeric quota, even if the
	//	service plan has an unlimited usage entitled on the level of the global account.
	UnlimitedAmountAssigned bool `json:"unlimitedAmountAssigned,omitempty"`
}
type ServicePlan struct {
	//The unique registration name of the service plan.
	Name      string `json:"name,omitempty"`
	Unlimited bool   `json:"unlimited,omitempty"`

	//Display name of the service plan for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`

	//Description of the service plan for customer-facing UIs.
	Description string `json:"description,omitempty"`

	//A unique identifier for service plans that can distinguish between the same service plans with different pricing plans.
	UniqueIdentifier string `json:"uniqueIdentifier,omitempty"`

	//The method used to provision the service plan.
	//
	//SERVICE_BROKER: Provisioning of NEO or CF quotas done by the service broker.
	//NONE_REQUIRED: Provisioning of CF quotas done by setting amount at provisioning-service.
	//COMMERCIAL_SOLUTION_SCRIPT: Provisioning is done by a script provided by the service owner and run by the Core Commercial Foundation service.
	//GLOBAL_COMMERCIAL_SOLUTION_SCRIPT: Provisioning is done by a script provided by the service owner and run by the Core Commercial Foundation service used for Global Account level.
	//GLOBAL_QUOTA_DOMAIN_DB: Provisioning is done by setting amount at Domain DB, this is relevant for non-ui quotas only.
	//Enum:
	//	[ SERVICE_BROKER, NONE_REQUIRED, COMMERCIAL_SOLUTION_SCRIPT, GLOBAL_COMMERCIAL_SOLUTION_SCRIPT, GLOBAL_QUOTA_DOMAIN_DB ]
	ProvisioningMethod string `json:"provisioningMethod,omitempty"`

	//The assigned quota for maximum allowed consumption of the plan. Relevant for services that have a numeric quota assignment.
	Amount float32 `json:"amount,omitempty"`

	//The remaining amount of the plan that can still be assigned. For plans that don't have a numeric quota,
	//the remaining amount is always the maximum allowed quota.
	RemainingAmount float32 `json:"remainingAmount,omitempty"`

	//[DEPRECATED] The source that added the service. Possible values:
	//
	//VENDOR: The product has been added by SAP or the cloud operator to the product catalog for general use.
	//GLOBAL_ACCOUNT_OWNER: Custom services that are added by a customer and are available only for that customer’s global account.
	//PARTNER: Service that are added by partners. And only available to its customers.
	//Note: This property is deprecated. Please use the ownerType attribute on the entitledService level instead.
	//
	//Enum:
	//	[ VENDOR, GLOBAL_ACCOUNT_OWNER, PARTNER ]
	ProvidedBy string `json:"providedBy,omitempty"`

	//Whether the service plan is a beta feature.
	Beta bool `json:"beta,omitempty"`

	//Whether the service plan is available internally to SAP users.
	AvailableForInternal bool `json:"availableForInternal,omitempty"`

	//The quota limit that is allowed for this service plan for SAP internal users.
	//If null, the default quota limit is set to 200.
	//Applies only when the availableForInternal property is set to TRUE.
	InternalQuotaLimit int32 `json:"internalQuotaLimit,omitempty"`

	//Whether to automatically assign a quota of the entitlement to a subaccount when the subaccount is
	//created in the entitlement's assigned directory.
	AutoAssign bool `json:"autoAssign,omitempty"`

	//The amount of the entitlement to automatically assign to a subaccount when the subaccount is created in the
	//entitlement's assigned directory.
	//Requires that autoAssign is set to TRUE, and there is remaining quota for the entitlement.
	AutoDistributeAmount int32 `json:"autoDistributeAmount,omitempty"`

	//The maximum allowed usage quota per subaccount for multitenant applications and environments that are defined as
	//"quota-based". This quota limits the usage of the application and/or environment per subaccount per a given usage
	//metric that is defined within the application or environment by the service provider. If null, the usage limit per
	//subaccount is the maximum free quota in the global account.
	//For example, a value of 1 could:
	//	(1) limit the number of subscriptions to a quota-based multitenant application within a global account according
	//		to the purchased quota, or
	//	(2) restrict the enablement of a single instance of an environment per subaccount.
	MaxAllowedSubAccountQuota int32 `json:"maxAllowedSubaccountQuota,omitempty"`

	//The type of service offering. Possible values:
	//
	//PLATFORM: A service required for using a specific platform; for example, Application Runtime is required for the
	//	Cloud Foundry platform.
	//SERVICE: A commercial or technical service. that has a numeric quota (amount) when entitled or assigned to a resource.
	//	When assigning entitlements of this type, use the 'amount' option instead of 'enable'.
	//	See: PUT/entitlements/v1/directories/{directoryGUID}/assignments.
	//ELASTIC_SERVICE: A commercial or technical service that has no numeric quota (amount) when entitled or assigned
	//	to a resource. Generally this type of service can be as many times as needed when enabled, but may in some cases
	//	be restricted by the service owner. When assigning entitlements of this type, use the 'enable' option instead
	//	of 'amount'.
	//	See: PUT/entitlements/v1/directories/{directoryGUID}/assignments.
	//ELASTIC_LIMITED: An elastic service that can be enabled for only one subaccount per global account.
	//APPLICATION: A multitenant application to which consumers can subscribe. As opposed to applications defined as a
	//	'QUOTA_BASED_APPLICATION', these applications do not have a numeric quota and are simply enabled or disabled as
	//	entitlements per subaccount.
	//QUOTA_BASED_APPLICATION: A multitenant application to which consumers can subscribe. As opposed to applications
	//	defined as 'APPLICATION', these applications have an numeric quota that limits consumer usage of the subscribed
	//	application per subaccount. When maxAllowedSubaccountQuota is > 0, this is the limit that can be set when
	//	assigning the max quota entitlement of the app to any subaccount. If maxAllowedSubaccountQuota is = 0 or null,
	//	the max quota that can be entitled to any subaccount is the amount purchased by the customer (the global account quota).
	//ENVIRONMENT: An environment service; for example, Cloud Foundry.
	//Enum:
	//	[ PLATFORM, SERVICE, ELASTIC_SERVICE, ELASTIC_LIMITED, APPLICATION, QUOTA_BASED_APPLICATION, ENVIRONMENT ]
	Category string `json:"category,omitempty"`

	//Relevant entitlements for the source that added the product.
	SourceEntitlements []SourceEntitlement `json:"sourceEntitlements,omitempty"`

	//Contains information about the data centers and regions in the cloud landscape
	DataCenters []DataCenter `json:"dataCenters,omitempty"`

	//Used to service plan external resources
	Resources []Resource `json:"resources,omitempty"`
}
type SourceEntitlement struct {
	//The technical name of the product.
	EntitlementName string `json:"entitlementName,omitempty"`

	//The quantity of the entitlement that is assigned to the root global account or directory.
	Amount float32 `json:"amount,omitempty"`

	//The product ID of the assigned entitlement.
	ProductId       string          `json:"productId,omitempty"`
	CommercialModel CommercialModel `json:"commercialModel,omitempty"`

	//Specifies if a plan associated with this entitlement will be automatically assigned by the system to any new
	//subaccount. For example, free plans that are available to all subaccounts.
	AutoAssign bool `json:"autoAssign,omitempty"`
}
type CommercialModel struct {
	//Whether a customer pays only for services that they actually use (consumption-based) or pays for subscribed
	//services at a fixed cost irrespective of consumption (subscription-based).
	//True: Consumption-based commercial model.False: Subscription-based commercial model.
	ConsumptionBased bool `json:"consumptionBased,omitempty"`

	//Directly contained commercial models.
	ContainedCommercialModels []CommercialModel `json:"containedCommercialModels,omitempty"`

	//A description of the commercial model
	Description string `json:"description,omitempty"`

	//A descriptive name of the commercial model for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`

	//Technical name of the commercial model.
	Name string `json:"name,omitempty"`
}
type Resource struct {
	//The name of the resource.
	Name string `json:"resourceName,omitempty"`
	//The name of the provider.
	Provider string `json:"resourceProvider,omitempty"`
	//The unique name of the resource.
	TechnicalName string `json:"resourceTechnicalName,omitempty"`
	//The type of the provider. For example infrastructure-as-a-service (IaaS).
	Type string `json:"resourceType,omitempty"`
	//Any additional data to include.
	Data interface{} `json:"resourceData,omitempty"`
}

func (c *EntitlementsV1) GetAssignments(ctx context.Context, input *GetAssignmentsInput) (*GetAssignmentsOutput, error) {
	req, out := c.getAssignmentsRequest(ctx, input)
	return out, req.Send()
}
func (c *EntitlementsV1) getAssignmentsRequest(ctx context.Context, input *GetAssignmentsInput) (*request.Request, *GetAssignmentsOutput) {
	op := &request.Operation{
		Name: "Get Entitlements",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/assignments",
		},
	}

	if input == nil {
		input = &GetAssignmentsInput{}
	}

	output := &GetAssignmentsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PUT /entitlements/v1/subaccountServicePlans
// Assign entitlements to subaccounts
type UpdateSubAccountServicePlanInput struct {
	//The details of entitlement's name, plan, amount and subaccount GUIDs to assign to a subaccount. The entitlement
	//can be a service, multitenant application, or environment. Note that some environments, such as Cloud Foundry,
	//are available by default to all subaccounts, and therefore are not displayed as entitlements.
	SubAccountServicePlans []SubAccountServicePlan `json:"subaccountServicePlans,omitempty"`
}
type SubAccountServicePlan struct {
	//The technical name of the entitlement to assign to a subaccount.
	ServiceName string `json:"serviceName,omitempty"`

	//The technical name of the entitlement's plan.
	ServicePlanName string `json:"servicePlanName,omitempty"`

	//List of assigned entitlements and their specifications.
	AssignmentInfo []AssignmentInfo `json:"assignmentInfo,omitempty"`
}
type AssignmentInfo struct {
	//The quantity of the plan that is assigned to the specified subaccount. Relevant and mandatory only for plans that
	//have a numeric quota. Do not set if enable=TRUE is specified.
	Amount *float32 `json:"amount,omitempty"`

	//Whether to enable the service plan assignment to the specified subaccount without quantity restrictions.
	//Relevant and mandatory only for plans that do not have a numeric quota. Do not set if amount is specified.
	Enable *bool `json:"enable,omitempty"`

	//The unique ID of the subaccount to which to assign a service plan.
	SubAccountGuid string `json:"subaccountGUID,omitempty"`

	//External resources to assign to subaccount
	Resources []Resource `json:"resources,omitempty"`
}
type UpdateSubAccountServicePlanOutput struct {
	JobStatusId *string `json:"jobStatusId,omitempty"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *EntitlementsV1) UpdateSubAccountServicePlan(ctx context.Context, input *UpdateSubAccountServicePlanInput) (*UpdateSubAccountServicePlanOutput, error) {
	req, out := c.updateUpdateSubAccountServicePlanRequest(ctx, input)
	return out, req.Send()
}
func (c *EntitlementsV1) updateUpdateSubAccountServicePlanRequest(ctx context.Context, input *UpdateSubAccountServicePlanInput) (*request.Request, *UpdateSubAccountServicePlanOutput) {
	op := &request.Operation{
		Name: "Update Sub Account ServicePlan",
		Http: request.HTTP{
			Method: request.PUT,
			Path:   "/subaccountServicePlans",
		},
	}

	if input == nil {
		input = &UpdateSubAccountServicePlanInput{}
	}

	output := &UpdateSubAccountServicePlanOutput{}

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

// PUT /entitlements/v1/directories/{directoryGUID}/assignments
// Assign or update an entitlement in a directory
type UpdateDirectoryEntitlementsInput struct {
	//The unique ID of the directory to which the entitlement is assigned.
	DirectoryGuid string `dest:"uri" dest-name:"directoryGUID"`

	//JSON object that contains the specifications of assignment, such as the name of the assigned plan, the quantity
	//to distribute, and whether to distribute the quota and how much to subaccounts that currently exist in the
	//directory and to subaccounts that will added to the directory in the future.
	DirectoryEntitlements []DirectoryEntitlement `json:"entitlements,omitempty"`
}
type DirectoryEntitlement struct {
	//The quantity of the plan to assign to the specified directory. Relevant and mandatory only for plans that have a
	//numeric quota. Do not set if enable=TRUE is specified.
	Amount *int64 `json:"amount,omitempty"`

	//The technical name of the entitlement to assign to the directory.
	Plan string `json:"plan,omitempty"`

	//Whether to allocate the plan to the to the specified directory without quantity restrictions.
	//Relevant and mandatory only for plans that don't have a numeric quota. Do not use if amount is specified.
	Enable *bool `json:"enable,omitempty"`

	//The technical name of the entitlement (service, application, environment) to assign.
	Service string `json:"service,omitempty"`

	//Whether to assign the plan with the quota specified in autoDistributeAmount to subaccounts currently located in
	//the specified directory. For entitlements without a numeric quota, such as multitenant apps, the plan is assigned
	//to the subaccounts currently located in the directory (autoDistributeAmount is not relevant in this case).
	//In both cases, autoAssign must be set to TRUE.
	Distribute bool `json:"distribute,omitempty"`

	//Whether to automatically allocate the plans of entitlements that have a numeric quota with the amount specified
	//in auto-distribute-amount to any new subaccount that is added to the directory in the future. For entitlements
	//without a numeric quota, the plan is assigned to any new subaccount that is added to the directory in the future
	//with the condition that enable=TRUE is set (autoDistributeAmount is not relevant in this case). If distribute=TRUE,
	//the same assignment is also made to all subaccounts currently in the directory. Entitlements are subject to
	//available quota in the directory.
	AutoAssign bool `json:"autoAssign,omitempty"`

	//The quota of the specified plan to automatically allocate to any new subaccount that is created in the future in the directory.
	//When applying this option, you must set autoAssign=TRUE and/or distribute=TRUE. Applies only to entitlements
	//that have a numeric quota. Entitlements are subject to available quota in the directory.
	AutoDistributeAmount int32 `json:"autoDistributeAmount,omitempty"`
}
type UpdateDirectoryEntitlementsOutput struct {
	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *EntitlementsV1) UpdateDirectoryEntitlements(ctx context.Context, input *UpdateDirectoryEntitlementsInput) (*UpdateDirectoryEntitlementsOutput, error) {
	req, out := c.updateDirectoryEntitlementsRequest(ctx, input)
	return out, req.Send()
}
func (c *EntitlementsV1) updateDirectoryEntitlementsRequest(ctx context.Context, input *UpdateDirectoryEntitlementsInput) (*request.Request, *UpdateDirectoryEntitlementsOutput) {
	op := &request.Operation{
		Name: "Update Directory Entitlements",
		Http: request.HTTP{
			Method: request.PUT,
			Path:   "/directories/{directoryGUID}/assignments",
		},
	}

	if input == nil {
		input = &UpdateDirectoryEntitlementsInput{}
	}

	output := &UpdateDirectoryEntitlementsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PATCH /entitlements/v1/directories/{directoryGUID}/assignments
// Update an existing entitlement in a directory
/*type PatchDirectoryEntitlementInput struct {
	//The unique ID of the directory to which the entitlement is assigned.
	DirectoryGuid string `dest:"uri" dest-name:"directoryGUID"`

	//JSON object that contains the specifications of an assignment, such as the name of the assigned plan and whether
	//to distribute the quota and how much to subaccounts that currently exist in the directory and to subaccounts that
	//will be added to the directory in the future.
	UpdateDirectoryAssignments []UpdateDirectoryAssignment `json:"entitlementUpdates,omitempty"`
}
type UpdateDirectoryAssignment struct {
	//The technical name of the entitlement to assign to the directory.
	Plan string `json:"plan,omitempty"`

	//The technical name of the entitlement (service, application, environment) to assign.
	Service string `json:"service,omitempty"`

	//Whether to assign the plan with the quota specified in autoDistributeAmount to subaccounts currently located in
	//the specified directory. For entitlements without a numeric quota, such as multitenant apps, the plan is assigned
	//to the subaccounts currently located in the directory (autoDistributeAmount is not relevant in this case).
	//In both cases, autoAssign must be set to TRUE.
	Distribute bool `json:"distribute,omitempty"`

	//Whether to automatically allocate the plans of entitlements that have a numeric quota with the amount specified in
	//auto-distribute-amount to any new subaccount that is added to the directory in the future.
	//For entitlements without a numeric quota, the plan is assigned to any new subaccount that is added to the directory
	//in the future with the condition that enable=TRUE is set (autoDistributeAmount is not relevant in this case).
	//If distribute=TRUE, the same assignment is also made to all subaccounts currently in the directory.
	//Entitlements are subject to available quota in the directory.
	AutoAssign bool `json:"autoAssign,omitempty"`

	//The quota of the specified plan to automatically allocate to any new subaccount that is created in the future in the directory.
	//When applying this option, you must set autoAssign=TRUE and/or distribute=TRUE.
	//Applies only to entitlements that have a numeric quota. Entitlements are subject to available quota in the directory.
	AutoDistributeAmount int32 `json:"autoDistributeAmount,omitempty"`
}
type UpdateDirectoryAssignmentOutput struct {
	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *EntitlementsV1) UpdateDirectoryAssignment(ctx context.Context, input *PatchDirectoryEntitlementInput) (*UpdateDirectoryAssignmentOutput, error) {
	req, out := c.updateUpdateDirectoryAssignmentRequest(ctx, input)
	return out, req.Send()
}
func (c *EntitlementsV1) updateUpdateDirectoryAssignmentRequest(ctx context.Context, input *PatchDirectoryEntitlementInput) (*request.Request, *UpdateDirectoryAssignmentOutput) {
	op := &request.Operation{
		Name: "Update Directory Assignment",
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/directories/{directoryGUID}/assignments",
		},
	}
	if input == nil {
		input = &PatchDirectoryEntitlementInput{}
	}
	output := &UpdateDirectoryAssignmentOutput{}
	return c.newRequest(ctx, op, input, output), output
}*/
