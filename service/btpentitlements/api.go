package btpentitlements

import (
	"context"
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
	DataCenters []DataCenter `json:"datacenters"`

	types.StatusAndBodyFromResponse
}
type DataCenter struct {
	//Technical name of the data center. Must be unique within the cloud deployment.
	Name string `json:"name"`

	//Descriptive name of the data center for customer-facing UIs.
	DisplayName string `json:"displayName"`

	//The region in which the data center is located.
	Region string `json:"region"`

	//The environment that the data center supports. For example: Kubernetes, Cloud Foundry.
	Environment string `json:"environment"`

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
	IaasProvider string `json:"iaasProvider"`

	//Whether the specified datacenter supports trial accounts.
	SupportsTrial bool `json:"supportsTrial"`

	//Provisioning service URL.
	ProvisioningServiceUrl string `json:"provisioningServiceUrl"`

	//Saas-Registry service URL.
	SaasRegistryServiceUrl string `json:"saasRegistryServiceUrl"`

	//The domain of the data center
	Domain string `json:"domain"`
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
	EntitledServices []EntitledService `json:"entitledServices"`

	//The list of services that are assigned to subaccounts located under a global account.
	AssignedServices []AssignedService `json:"assignedServices"`

	//Whether the External Provider Registry (XPR) is available.
	FetchErrorFromExternalProviderRegistry bool `json:"fetchErrorFromExternalProviderRegistry"`

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
	EntitledServices []EntitledService `json:"entitledServices"`

	//The list of services that are assigned to subaccounts located under a global account.
	AssignedServices []AssignedService `json:"assignedServices"`

	//Whether the External Provider Registry (XPR) is available.
	FetchErrorFromExternalProviderRegistry bool `json:"fetchErrorFromExternalProviderRegistry"`

	types.StatusAndBodyFromResponse
}
type AssignedService struct {
	//The unique registration name of the deployed service as defined by the service provider.
	Name string `json:"name"`

	//Display name of the service for customer-facing UIs.
	DisplayName string `json:"displayName"`

	BusinessCategory BusinessCategory `json:"businessCategory"`

	//List of service plans associated with the assigned service.
	ServicePlans []AssignedServicePlan `json:"servicePlans"`

	//The icon of the service in Base64 format.
	IconBase64 string `json:"iconBase64"`

	//The application coordinates provided in metadata.
	ApplicationCoordinates interface{} `json:"applicationCoordinates"`
}
type EntitledService struct {
	//The unique registration name of the deployed service as defined by the service provider.
	Name string `json:"name"`

	//Display name of the service for customer-facing UIs.
	DisplayName string `json:"displayName"`

	//Description of the service for customer-facing UIs.
	Description      string           `json:"description"`
	BusinessCategory BusinessCategory `json:"businessCategory"`

	//The owner type of the service. Possible values:
	//
	//VENDOR: The owner is a service owner, who is affiliated with the cloud operator, that added the service to the product catalog for general consumption.
	//CUSTOMER: The owner is an SAP customer that added a custom service to the product catalog, and it is available only for consumption within the customer's global account.
	//PARTNER: The owner is an SAP partner that added the service to the product catalog, and it is available only to their customers for consumption.
	//Enum:
	//	[ VENDOR, CUSTOMER, PARTNER ]
	OwnerType string `json:"ownerType"`

	//List of service plans associated with the entitled service.
	ServicePlans []ServicePlan `json:"servicePlans"`

	//The icon of the service in Base64 format.
	IconBase64 string `json:"iconBase64"`

	//The application coordinates provided in metadata.
	ApplicationCoordinates interface{} `json:"applicationCoordinates"`
}
type BusinessCategory struct {
	//Unique ID of the business category.
	Id string `json:"id"`

	//Display name of the business category for customer-facing UIs.
	DisplayName string `json:"displayName"`
}
type AssignedServicePlan struct {
	//The unique registration name of the service plan.
	Name string `json:"name"`

	//The name of the service plan for customer-facing UIs.
	DisplayName string `json:"displayName"`

	//A unique identifier for service plans that can distinguish between the same service plans with different pricing plans.
	UniqueIdentifier string `json:"uniqueIdentifier"`

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
	Category string `json:"category"`

	//Whether the service plan is a beta feature.
	Beta bool `json:"beta"`

	//The maximum allowed usage quota per subaccount for multitenant applications and environments that are defined as
	//"quota-based". This quota limits the usage of the application and/or environment per subaccount per a given
	//usage metric that is defined within the application or environment by the service provider. If null, the usage
	//limit per subaccount is the maximum free quota in the global account.
	//For example, a value of 1 could:
	//	(1) limit the number of subscriptions to a quota-based multitenant application within a global account according
	//		to the purchased quota, or
	//	(2) restrict the enablement of a single instance of an environment per subaccount.
	MaxAllowedSubAccountQuota int32 `json:"maxAllowedSubaccountQuota"`

	//Is the quota of this service plan entitled to the global account with unlimited usage.
	Unlimited bool `json:"unlimited"`

	//Assignment detailed information
	AssignmentInfo []AssignedServicePlanSubAccount `json:"assignmentInfo"`
}
type AssignedServicePlanSubAccount struct {
	//Specifies if the plan was automatically assigned regardless of any action by an admin. This applies to entitlements
	//that are always available to subaccounts and cannot be removed.
	AutoAssigned bool `json:"autoAssigned"`

	//example: GUID of GLOBAL_ACCOUNT or SUBACCOUNT
	//The unique ID of the global account or directory to which the entitlement is assigned.
	EntityId              string      `json:"entityId"`
	ParentAmount          int64       `json:"parentAmount"`
	ParentId              string      `json:"parentId"`
	ParentRemainingAmount interface{} `json:"parentRemainingAmount"`

	//Enum:
	//	[ SUBACCOUNT, GLOBAL_ACCOUNT, DIRECTORY ]
	ParentType string `json:"parentType"`

	//example: GLOBAL_ACCOUNT or SUBACCOUNT
	//The type of entity to which the entitlement is assigned.
	//
	//SUBACCOUNT: The entitlement is assigned to a subaccount.
	//GLOBAL_ACCOUNT: The entitlement is assigned to a root global account.
	//DIRECTORY: The entitlement is assigned to a directory.
	//Enum:
	//	[ SUBACCOUNT, GLOBAL_ACCOUNT, DIRECTORY ]
	EntityType string `json:"entityType"`

	//The quantity of the entitlement that is assigned to the root global account or directory.
	Amount int64 `json:"amount"`

	//The requested amount when it is different from the actual amount because the request state is still in process or failed.
	RequestedAmount int64 `json:"requestedAmount"`

	//The current state of the service plan assignment.
	//
	//STARTED: CRUD operation on an entity has started.
	//PROCESSING: A series of operations related to the entity is in progress.
	//PROCESSING_FAILED: The processing operations failed.
	//OK: The CRUD operation or series of operations completed successfully.
	//Enum:
	//	[ STARTED, PROCESSING, PROCESSING_FAILED, OK ]
	EntityState string `json:"entityState"`

	//Information about the current state.
	StateMessage string `json:"stateMessage"`

	//Whether the plan is automatically distributed to the subaccounts that are located in the directory.
	AutoAssign bool `json:"autoAssign"`

	//The amount of the entitlement to automatically assign to subaccounts that are added in the future to the
	//entitlement's assigned directory.
	//Requires that autoAssign is set to TRUE, and there is remaining quota for the entitlement. To automatically
	//distribute to subaccounts that are added in the future to the directory, distribute must be set to TRUE.
	AutoDistributeAmount int32 `json:"autoDistributeAmount"`

	//Date the subaccount has been created. Dates and times are in UTC format.
	CreatedDate times.JavaTime `json:"createdDate"`

	//Date the subaccount has been modified. Dates and times are in UTC format.
	ModifiedDate times.JavaTime `json:"modifiedDate"`

	//Global account resource details
	Resources []Resource `json:"resources"`

	//True, if an unlimited quota of this service plan assigned to the directory or subaccount in the global account.
	//False, if the service plan is assigned to the directory or subaccount with a limited numeric quota, even if the
	//	service plan has an unlimited usage entitled on the level of the global account.
	UnlimitedAmountAssigned bool `json:"unlimitedAmountAssigned"`
}
type ServicePlan struct {
	//The unique registration name of the service plan.
	Name      string `json:"name"`
	Unlimited bool   `json:"unlimited"`

	//Display name of the service plan for customer-facing UIs.
	DisplayName string `json:"displayName"`

	//Description of the service plan for customer-facing UIs.
	Description string `json:"description"`

	//A unique identifier for service plans that can distinguish between the same service plans with different pricing plans.
	UniqueIdentifier string `json:"uniqueIdentifier"`

	//The method used to provision the service plan.
	//
	//SERVICE_BROKER: Provisioning of NEO or CF quotas done by the service broker.
	//NONE_REQUIRED: Provisioning of CF quotas done by setting amount at provisioning-service.
	//COMMERCIAL_SOLUTION_SCRIPT: Provisioning is done by a script provided by the service owner and run by the Core Commercial Foundation service.
	//GLOBAL_COMMERCIAL_SOLUTION_SCRIPT: Provisioning is done by a script provided by the service owner and run by the Core Commercial Foundation service used for Global Account level.
	//GLOBAL_QUOTA_DOMAIN_DB: Provisioning is done by setting amount at Domain DB, this is relevant for non-ui quotas only.
	//Enum:
	//	[ SERVICE_BROKER, NONE_REQUIRED, COMMERCIAL_SOLUTION_SCRIPT, GLOBAL_COMMERCIAL_SOLUTION_SCRIPT, GLOBAL_QUOTA_DOMAIN_DB ]
	ProvisioningMethod string `json:"provisioningMethod"`

	//The assigned quota for maximum allowed consumption of the plan. Relevant for services that have a numeric quota assignment.
	Amount interface{} `json:"amount"`

	//The remaining amount of the plan that can still be assigned. For plans that don't have a numeric quota,
	//the remaining amount is always the maximum allowed quota.
	RemainingAmount interface{} `json:"remainingAmount"`

	//[DEPRECATED] The source that added the service. Possible values:
	//
	//VENDOR: The product has been added by SAP or the cloud operator to the product catalog for general use.
	//GLOBAL_ACCOUNT_OWNER: Custom services that are added by a customer and are available only for that customer’s global account.
	//PARTNER: Service that are added by partners. And only available to its customers.
	//Note: This property is deprecated. Please use the ownerType attribute on the entitledService level instead.
	//
	//Enum:
	//	[ VENDOR, GLOBAL_ACCOUNT_OWNER, PARTNER ]
	ProvidedBy string `json:"providedBy"`

	//Whether the service plan is a beta feature.
	Beta bool `json:"beta"`

	//Whether the service plan is available internally to SAP users.
	AvailableForInternal bool `json:"availableForInternal"`

	//The quota limit that is allowed for this service plan for SAP internal users.
	//If null, the default quota limit is set to 200.
	//Applies only when the availableForInternal property is set to TRUE.
	InternalQuotaLimit int32 `json:"internalQuotaLimit"`

	//Whether to automatically assign a quota of the entitlement to a subaccount when the subaccount is
	//created in the entitlement's assigned directory.
	AutoAssign bool `json:"autoAssign"`

	//The amount of the entitlement to automatically assign to a subaccount when the subaccount is created in the
	//entitlement's assigned directory.
	//Requires that autoAssign is set to TRUE, and there is remaining quota for the entitlement.
	AutoDistributeAmount int32 `json:"autoDistributeAmount"`

	//The maximum allowed usage quota per subaccount for multitenant applications and environments that are defined as
	//"quota-based". This quota limits the usage of the application and/or environment per subaccount per a given usage
	//metric that is defined within the application or environment by the service provider. If null, the usage limit per
	//subaccount is the maximum free quota in the global account.
	//For example, a value of 1 could:
	//	(1) limit the number of subscriptions to a quota-based multitenant application within a global account according
	//		to the purchased quota, or
	//	(2) restrict the enablement of a single instance of an environment per subaccount.
	MaxAllowedSubAccountQuota int32 `json:"maxAllowedSubaccountQuota"`

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
	Category string `json:"category"`

	//Relevant entitlements for the source that added the product.
	SourceEntitlements []SourceEntitlement `json:"sourceEntitlements"`

	//Contains information about the data centers and regions in the cloud landscape
	DataCenters []DataCenter `json:"dataCenters"`

	//Used to service plan external resources
	Resources []Resource `json:"resources"`
}
type SourceEntitlement struct {
	//The technical name of the product.
	EntitlementName string `json:"entitlementName"`

	//The quantity of the entitlement that is assigned to the root global account or directory.
	Amount int64 `json:"amount"`

	//The product ID of the assigned entitlement.
	ProductId       string          `json:"productId"`
	CommercialModel CommercialModel `json:"commercialModel"`

	//Specifies if a plan associated with this entitlement will be automatically assigned by the system to any new
	//subaccount. For example, free plans that are available to all subaccounts.
	AutoAssign bool `json:"autoAssign"`
}
type CommercialModel struct {
	//Whether a customer pays only for services that they actually use (consumption-based) or pays for subscribed
	//services at a fixed cost irrespective of consumption (subscription-based).
	//True: Consumption-based commercial model.False: Subscription-based commercial model.
	ConsumptionBased bool `json:"consumptionBased"`

	//Directly contained commercial models.
	ContainedCommercialModels []CommercialModel `json:"containedCommercialModels"`

	//A description of the commercial model
	Description string `json:"description"`

	//A descriptive name of the commercial model for customer-facing UIs.
	DisplayName string `json:"displayName"`

	//Technical name of the commercial model.
	Name string `json:"name"`
}
type Resource struct {
	//The name of the resource.
	Name string `json:"resourceName"`
	//The name of the provider.
	Provider string `json:"resourceProvider"`
	//The unique name of the resource.
	TechnicalName string `json:"resourceTechnicalName"`
	//The type of the provider. For example infrastructure-as-a-service (IaaS).
	Type string `json:"resourceType"`
	//Any additional data to include.
	Data interface{} `json:"resourceData"`
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
	SubAccountServicePlans []SubAccountServicePlan `json:"subaccountServicePlans"`
}
type SubAccountServicePlan struct {
	//The technical name of the entitlement to assign to a subaccount.
	ServiceName string `json:"serviceName"`

	//The technical name of the entitlement's plan.
	ServicePlanName string `json:"servicePlanName"`

	//List of assigned entitlements and their specifications.
	AssignmentInfo []AssignmentInfo `json:"assignmentInfo"`
}
type AssignmentInfo struct {
	//The quantity of the plan that is assigned to the specified subaccount. Relevant and mandatory only for plans that
	//have a numeric quota. Do not set if enable=TRUE is specified.
	Amount int64 `json:"amount"`

	//Whether to enable the service plan assignment to the specified subaccount without quantity restrictions.
	//Relevant and mandatory only for plans that do not have a numeric quota. Do not set if amount is specified.
	Enable bool `json:"enable"`

	//The unique ID of the subaccount to which to assign a service plan.
	SubAccountGuid string `json:"subaccountGUID"`

	//External resources to assign to subaccount
	Resources []Resource `json:"resources"`
}
type UpdateSubAccountServicePlanOutput struct {
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
	return c.newRequest(ctx, op, input, output), output
}

// PUT /entitlements/v1/directories/{directoryGUID}/assignments
// Assign or update an entitlement in a directory
type AssignDirectoryAssignmentInput struct {
	//The unique ID of the directory to which the entitlement is assigned.
	DirectoryGuid string `dest:"uri" dest-name:"directoryGUID"`

	//JSON object that contains the specifications of assignment, such as the name of the assigned plan, the quantity
	//to distribute, and whether to distribute the quota and how much to subaccounts that currently exist in the
	//directory and to subaccounts that will added to the directory in the future.
	DirectoryAssignments []AssignDirectoryAssignment `json:"entitlements"`
}
type AssignDirectoryAssignment struct {
	//The quantity of the plan to assign to the specified directory. Relevant and mandatory only for plans that have a
	//numeric quota. Do not set if enable=TRUE is specified.
	Amount int64 `json:"amount"`

	//The technical name of the entitlement to assign to the directory.
	Plan string `json:"plan"`

	//Whether to allocate the plan to the to the specified directory without quantity restrictions.
	//Relevant and mandatory only for plans that don't have a numeric quota. Do not use if amount is specified.
	Enable bool `json:"enable"`

	//The technical name of the entitlement (service, application, environment) to assign.
	Service string `json:"service"`

	//Whether to assign the plan with the quota specified in autoDistributeAmount to subaccounts currently located in
	//the specified directory. For entitlements without a numeric quota, such as multitenant apps, the plan is assigned
	//to the subaccounts currently located in the directory (autoDistributeAmount is not relevant in this case).
	//In both cases, autoAssign must be set to TRUE.
	Distribute bool `json:"distribute"`

	//Whether to automatically allocate the plans of entitlements that have a numeric quota with the amount specified
	//in auto-distribute-amount to any new subaccount that is added to the directory in the future. For entitlements
	//without a numeric quota, the plan is assigned to any new subaccount that is added to the directory in the future
	//with the condition that enable=TRUE is set (autoDistributeAmount is not relevant in this case). If distribute=TRUE,
	//the same assignment is also made to all subaccounts currently in the directory. Entitlements are subject to
	//available quota in the directory.
	AutoAssign bool `json:"autoAssign"`

	//The quota of the specified plan to automatically allocate to any new subaccount that is created in the future in the directory.
	//When applying this option, you must set autoAssign=TRUE and/or distribute=TRUE. Applies only to entitlements
	//that have a numeric quota. Entitlements are subject to available quota in the directory.
	AutoDistributeAmount int32 `json:"autoDistributeAmount"`
}
type AssignDirectoryAssignmentOutput struct {
	types.StatusAndBodyFromResponse
}

func (c *EntitlementsV1) AssignDirectoryAssignment(ctx context.Context, input *AssignDirectoryAssignmentInput) (*AssignDirectoryAssignmentOutput, error) {
	req, out := c.assignUpdateDirectoryAssignmentRequest(ctx, input)
	return out, req.Send()
}
func (c *EntitlementsV1) assignUpdateDirectoryAssignmentRequest(ctx context.Context, input *AssignDirectoryAssignmentInput) (*request.Request, *AssignDirectoryAssignmentOutput) {
	op := &request.Operation{
		Name: "Assign Directory Assignment",
		Http: request.HTTP{
			Method: request.PUT,
			Path:   "/directories/{directoryGUID}/assignments",
		},
	}

	if input == nil {
		input = &AssignDirectoryAssignmentInput{}
	}

	output := &AssignDirectoryAssignmentOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PATCH /entitlements/v1/directories/{directoryGUID}/assignments
// Update an existing entitlement in a directory
type UpdateDirectoryAssignmentInput struct {
	//The unique ID of the directory to which the entitlement is assigned.
	DirectoryGuid string `dest:"uri" dest-name:"directoryGUID"`

	//JSON object that contains the specifications of an assignment, such as the name of the assigned plan and whether
	//to distribute the quota and how much to subaccounts that currently exist in the directory and to subaccounts that
	//will be added to the directory in the future.
	UpdateDirectoryAssignments []UpdateDirectoryAssignment `json:"entitlementUpdates"`
}
type UpdateDirectoryAssignment struct {
	//The technical name of the entitlement to assign to the directory.
	Plan string `json:"plan"`

	//The technical name of the entitlement (service, application, environment) to assign.
	Service string `json:"service"`

	//Whether to assign the plan with the quota specified in autoDistributeAmount to subaccounts currently located in
	//the specified directory. For entitlements without a numeric quota, such as multitenant apps, the plan is assigned
	//to the subaccounts currently located in the directory (autoDistributeAmount is not relevant in this case).
	//In both cases, autoAssign must be set to TRUE.
	Distribute bool `json:"distribute"`

	//Whether to automatically allocate the plans of entitlements that have a numeric quota with the amount specified in
	//auto-distribute-amount to any new subaccount that is added to the directory in the future.
	//For entitlements without a numeric quota, the plan is assigned to any new subaccount that is added to the directory
	//in the future with the condition that enable=TRUE is set (autoDistributeAmount is not relevant in this case).
	//If distribute=TRUE, the same assignment is also made to all subaccounts currently in the directory.
	//Entitlements are subject to available quota in the directory.
	AutoAssign bool `json:"autoAssign"`

	//The quota of the specified plan to automatically allocate to any new subaccount that is created in the future in the directory.
	//When applying this option, you must set autoAssign=TRUE and/or distribute=TRUE.
	//Applies only to entitlements that have a numeric quota. Entitlements are subject to available quota in the directory.
	AutoDistributeAmount int32 `json:"autoDistributeAmount"`
}
type UpdateDirectoryAssignmentOutput struct {
	types.StatusAndBodyFromResponse
}

func (c *EntitlementsV1) UpdateDirectoryAssignment(ctx context.Context, input *UpdateDirectoryAssignmentInput) (*UpdateDirectoryAssignmentOutput, error) {
	req, out := c.updateUpdateDirectoryAssignmentRequest(ctx, input)
	return out, req.Send()
}
func (c *EntitlementsV1) updateUpdateDirectoryAssignmentRequest(ctx context.Context, input *UpdateDirectoryAssignmentInput) (*request.Request, *UpdateDirectoryAssignmentOutput) {
	op := &request.Operation{
		Name: "Update Directory Assignment",
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/directories/{directoryGUID}/assignments",
		},
	}
	if input == nil {
		input = &UpdateDirectoryAssignmentInput{}
	}
	output := &UpdateDirectoryAssignmentOutput{}
	return c.newRequest(ctx, op, input, output), output
}
