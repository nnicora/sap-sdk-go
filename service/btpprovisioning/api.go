package btpprovisioning

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
)

const quotaAssignments = "Quota Assignments"
const environments = "Environments"

// GET /provisioning/v1/servicePlanAssignments
// Get subaccount quota assignments
type GetServicePlanAssignmentsInput struct {
}
type GetServicePlanAssignmentsOutput struct {
	Quotas []ServicePlanQuotaAssignment `json:"quotas"`
}
type ServicePlanQuotaAssignment struct {
	//The quantity of consumed quota. The service owner can provide a value for the consumed quota if it is different
	//from the provisioned quota calculation.
	ConsumedQuota []string `json:"consumedQuota"`
	//Unique ID of the global account associated with the subaccount.
	GlobalAccountGuid string `json:"globalAccountGUID"`
	//The ID of the associated global account.
	GlobalAccountId string `json:"globalAccountId"`
	//The name of the plan of the provisioned quota.
	Plan string `json:"plan"`
	//Enum:
	//	[ SERVICE_BROKER, NONE_REQUIRED, COMMERCIAL_SOLUTION_SCRIPT, GLOBAL_COMMERCIAL_SOLUTION_SCRIPT,
	//	GLOBAL_QUOTA_DOMAIN_DB ]
	ProvisioningMethod string `json:"provisioningMethod"`
	//The quantity of provisioned quota.
	Quota     int32      `json:"quota"`
	Resources []Resource `json:"resources"`
	//The name of the service of the provisioned quota.
	Service string `json:"service"`
	//Enum:
	//	[ PLATFORM, SERVICE, ELASTIC_SERVICE, ELASTIC_LIMITED, APPLICATION, QUOTA_BASED_APPLICATION, ENVIRONMENT ]
	ServiceCategory string `json:"serviceCategory"`
	//Unique ID of the subaccount for which to get quota.
	SubAccountGuid string `json:"subaccountGUID"`
	//The ID of the tenant for the subaccount.
	TenantId string `json:"tenantId"`
	//Whether an unlimited quantity of quota can be provisioned.
	Unlimited bool `json:"unlimited"`
}
type Resource struct {
	//any relevant information about the resource that is not provided by other parameter values.
	AdditionalInfo interface{} `json:"additionalInfo"`
	//Description of the resource.
	Description string `json:"description"`
	//Descriptive name of the resource for customer-facing UIs.
	DisplayName string `json:"displayName"`
	//Data associated with the resource.
	Data interface{} `json:"resourceData"`
	//Provider of the requested resource. For example, IaaS provider: AWS.
	Provider string `json:"resourceProvider"`
	//Unique technical name of the resource.
	TechnicalName string `json:"resourceTechnicalName"`
	//Type of the resource.
	Type string `json:"resourceType"`
}

func (c *ProvisioningV1) GetServicePlanQuotaAssignments(ctx context.Context) (*GetServicePlanAssignmentsOutput, error) {
	req, out := c.getServicePlanQuotaAssignmentsRequest(ctx, nil)
	return out, req.Send()
}
func (c *ProvisioningV1) getServicePlanQuotaAssignmentsRequest(ctx context.Context, input *GetServicePlanAssignmentsInput) (*request.Request, *GetServicePlanAssignmentsOutput) {
	op := &request.Operation{
		Name: quotaAssignments,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/servicePlanAssignments",
		},
	}

	if input == nil {
		input = &GetServicePlanAssignmentsInput{}
	}

	output := &GetServicePlanAssignmentsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /provisioning/v1/availableEnvironments
// Get available environments
type GetAvailableEnvironmentsInput struct {
	//Security token that contains authentication declarations of an end user by the authorization server
	//(SAP Identity and Authentication Service).
	XIDToken string `dest:"header" dest-name:"X-ID-Token"`
}
type GetAvailableEnvironmentsOutput struct {
	Environments []AvailableEnvironment `json:"availableEnvironments"`
}
type AvailableEnvironment struct {
	//The availability level of the environment broker.
	AvailabilityLevel string `json:"availabilityLevel"`
	//The create schema of the environment broker.
	CreateSchema string `json:"createSchema"`
	//Description of the service plan for the available environment.
	Description string `json:"description"`
	//The type of environment that is available (for example: cloudfoundry).
	EnvironmentType string `json:"environmentType"`
	//The landscape label of the environment broker.
	LandscapeLabel string `json:"landscapeLabel"`
	//Name of the service plan for the available environment.
	PlanName string `json:"planName"`
	//The description of the service.
	ServiceDescription string `json:"serviceDescription"`
	//The display name of the service.
	ServiceDisplayName string `json:"serviceDisplayName"`
	//The URL of the documentation link for the service.
	ServiceDocumentationUrl string `json:"serviceDocumentationUrl"`
	//The URL of the image for the service.
	ServiceImageUrl string `json:"serviceImageUrl"`
	//Name of the service offered in the catalog of the corresponding environment broker (for example, cloudfoundry).
	ServiceName string `json:"serviceName"`
	//The URL of the support link for the service.
	ServiceSupportUrl string `json:"serviceSupportUrl"`
	//Technical key of the corresponding environment broker.
	TechnicalKey string `json:"technicalKey"`
}

func (c *ProvisioningV1) GetAvailableEnvironments(ctx context.Context) (*GetAvailableEnvironmentsOutput, error) {
	req, out := c.getAvailableEnvironmentsRequest(ctx, nil)
	return out, req.Send()
}
func (c *ProvisioningV1) getAvailableEnvironmentsRequest(ctx context.Context, input *GetAvailableEnvironmentsInput) (*request.Request, *GetAvailableEnvironmentsOutput) {
	op := &request.Operation{
		Name: environments,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/availableEnvironments",
		},
	}

	if input == nil {
		input = &GetAvailableEnvironmentsInput{}
	}

	output := &GetAvailableEnvironmentsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /provisioning/v1/environments
// Get environment instances
type GetEnvironmentInstancesInput struct {
	//Security token that contains authentication declarations of an end user by the authorization server
	//(SAP Identity and Authentication Service).
	XIDToken string `dest:"header" dest-name:"X-ID-Token"`
}
type GetEnvironmentInstancesOutput struct {
	//The list of all the environment instances to delete
	Environments []EnvironmentInstance `json:"environmentInstances"`
}
type EnvironmentInstance struct {
	//The ID of the associated environment broker.
	BrokerId string `json:"brokerId"`
	//The commercial type of the environment broker.
	CommercialType string `json:"commercialType"`
	//The date the environment instance was created. Dates and times are in UTC format.
	CreatedDate string `json:"createdDate"`
	//The URL of the service dashboard, which is a web-based management user interface for the service instances.
	DashboardUrl string `json:"dashboardUrl"`
	//The description of the environment instance.
	Description string `json:"description"`
	//Type of the environment instance that is used.
	//Enum:
	//	[ cloudfoundry, kubernetes, neo ]
	EnvironmentType string `json:"environmentType"`
	//The GUID of the global account that is associated with the environment instance.
	GlobalAccountGuid string `json:"globalAccountGUID"`
	//Automatically generated unique identifier for the environment instance.
	Id string `json:"id"`
	//Broker-specified key-value pairs that specify attributes of a service instance.
	Labels string `json:"labels"`
	//The name of the landscape within the logged-in region on which the environment instance is created.
	LandscapeLabel string `json:"landscapeLabel"`
	//The last date the environment instance was last modified. Dates and times are in UTC format.
	ModifiedDate string `json:"modifiedDate"`
	//Name of the environment instance.
	Name string `json:"name"`
	//An identifier that represents the last operation. This ID is returned by the environment brokers.
	Operation string `json:"operation"`
	//Configuration parameters for the environment instance.
	Parameters string `json:"parameters"`
	//ID of the service plan for the environment instance in the corresponding service broker's catalog.
	PlanId string `json:"planId"`
	//Name of the service plan for the environment instance in the corresponding service broker's catalog.
	PlanName string `json:"planName"`
	//ID of the platform for the environment instance in the corresponding service broker's catalog.
	PlatformId string `json:"platformId"`
	//ID of the service for the environment instance in the corresponding service broker's catalog.
	ServiceId string `json:"serviceId"`
	//Name of the service for the environment instance in the corresponding service broker's catalog.
	ServiceName string `json:"serviceName"`
	//Current state of the environment instance.
	//Enum:
	//	[ CREATING, UPDATING, DELETING, OK, CREATION_FAILED, DELETION_FAILED, UPDATE_FAILED ]
	State string `json:"state"`
	//Information about the current state of the environment instance.
	StateMessage string `json:"stateMessage"`
	//The GUID of the subaccount associated with the environment instance.
	SubAccountGuid string `json:"subaccountGUID"`
	//The ID of the tenant that owns the environment instance.
	TenantId string `json:"tenantId"`
	//The last provisioning operation on the environment instance.
	//Provision: Environment instance created.
	//Update: Environment instance changed.
	//Deprovision: Environment instance deleted.
	//Enum:
	//	[ Provision, Update, Deprovision ]
	Type string `json:"type"`
}

func (c *ProvisioningV1) GetEnvironmentInstances(ctx context.Context) (*GetEnvironmentInstancesOutput, error) {
	req, out := c.getEnvironmentInstancesRequest(ctx, nil)
	return out, req.Send()
}
func (c *ProvisioningV1) getEnvironmentInstancesRequest(ctx context.Context, input *GetEnvironmentInstancesInput) (*request.Request, *GetEnvironmentInstancesOutput) {
	op := &request.Operation{
		Name: environments,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/environments",
		},
	}

	if input == nil {
		input = &GetEnvironmentInstancesInput{}
	}

	output := &GetEnvironmentInstancesOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /provisioning/v1/environments
// Get environment instances
type CreateEnvironmentInstancesInput struct {
	//The description of the environment instance.
	Description string `json:"description"`
	//Type of the environment instance that is used. Must match the type of the environment instance broker
	//(for example: cloudfoundry). Use GET /provisioning/v1/availableEnvironments to view the valid values.
	EnvironmentType string `json:"environmentType"`
	//The name of the landscape within the logged-in region on which to create the environment instance. Only required
	//only if the region has multiple landscapes. To see which landscapes are available for this environment, use the
	//GET /provisioning/v1/availableEnvironments API.
	LandscapeLabel string `json:"landscapeLabel"`
	//The name of the created environment instance.
	Name string `json:"name"`
	//The origin of the user in case of a custom IdP configuration. This parameter is only required if the
	//OAuth 2.0 client credentials grant flow is used, a user parameter is provided and the target environment supports
	//custom IdP, otherwise it is ignored.
	Origin string `json:"origin"`

	//If needed, you can pass environment-specific configuration parameters using a valid embedded JSON object.
	//For a list of supported configuration parameters, see the documentation of the particular environment offering.
	//In this example, additional configuration parameters 'id' and 'email' are specified:
	//{
	//"instance_name": "myOrg"
	//}
	Parameters map[string]interface{} `json:"parameters"`
	//Name of the service plan for the environment instance. Must match the name in the corresponding service broker's
	//catalog. (for example: standard)
	PlanName string `json:"planName"`
	//The name of service offered in the catalog of the corresponding environment broker. (for example: cloudfoundry)
	ServiceName string `json:"serviceName"`
	//Technical key of the corresponding environment broker.
	TechnicalKey string `json:"technicalKey"`
	//The e-mail of the user that owns the environment instance. In some environments, this user might be assigned as
	//the initial admin of the provisioned environment. For example, for a Cloud Foundry environment, this user is
	//assigned to the Org Manager role. This parameter is required only when OAuth 2.0 client credentials grant flow
	//is used, otherwise it is ignored.
	User string `json:"user"`
}
type CreateEnvironmentInstancesOutput struct {
	//ID of the created environment instance
	Id string `json:"id"`
}

func (c *ProvisioningV1) CreateEnvironmentInstances(ctx context.Context) (*CreateEnvironmentInstancesOutput, error) {
	req, out := c.createEnvironmentInstancesRequest(ctx, nil)
	return out, req.Send()
}
func (c *ProvisioningV1) createEnvironmentInstancesRequest(ctx context.Context, input *CreateEnvironmentInstancesInput) (*request.Request, *CreateEnvironmentInstancesOutput) {
	op := &request.Operation{
		Name: environments,
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/environments",
		},
	}

	if input == nil {
		input = &CreateEnvironmentInstancesInput{}
	}

	output := &CreateEnvironmentInstancesOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /provisioning/v1/environments
// Delete all environment instances
type DeleteEnvironmentInstancesInput struct {
}
type DeleteEnvironmentInstancesOutput struct {
	//The list of all the environment instances to delete
	Environments []EnvironmentInstance `json:"environmentInstances"`
}

func (c *ProvisioningV1) DeleteEnvironmentInstances(ctx context.Context) (*DeleteEnvironmentInstancesOutput, error) {
	req, out := c.deleteEnvironmentInstancesRequest(ctx, nil)
	return out, req.Send()
}
func (c *ProvisioningV1) deleteEnvironmentInstancesRequest(ctx context.Context, input *DeleteEnvironmentInstancesInput) (*request.Request, *DeleteEnvironmentInstancesOutput) {
	op := &request.Operation{
		Name: environments,
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/environments",
		},
	}

	if input == nil {
		input = &DeleteEnvironmentInstancesInput{}
	}

	output := &DeleteEnvironmentInstancesOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /provisioning/v1/environments/{environmentInstanceId}
// Get an environment instance
type GetEnvironmentInstanceInput struct {
	//The ID of the environment instance to view.
	EnvironmentInstanceId string `dest:"uri" dest-name:"environmentInstanceId"`
}
type GetEnvironmentInstanceOutput struct {
	EnvironmentInstance
}

func (c *ProvisioningV1) GetEnvironmentInstance(ctx context.Context, input *GetEnvironmentInstanceInput) (*GetEnvironmentInstanceOutput, error) {
	req, out := c.getEnvironmentInstanceRequest(ctx, input)
	return out, req.Send()
}
func (c *ProvisioningV1) getEnvironmentInstanceRequest(ctx context.Context, input *GetEnvironmentInstanceInput) (*request.Request, *GetEnvironmentInstanceOutput) {
	op := &request.Operation{
		Name: environments,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/environments/{environmentInstanceId}",
		},
	}

	if input == nil {
		input = &GetEnvironmentInstanceInput{}
	}

	output := &GetEnvironmentInstanceOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /provisioning/v1/environments/{environmentInstanceId}
// Delete an environment instance
type DeleteEnvironmentInstanceInput struct {
	//ID of the environment instance to delete
	EnvironmentInstanceId string `dest:"uri" dest-name:"environmentInstanceId"`
}
type DeleteEnvironmentInstanceOutput struct {
	EnvironmentInstance
}

func (c *ProvisioningV1) DeleteEnvironmentInstance(ctx context.Context, input *DeleteEnvironmentInstanceInput) (*DeleteEnvironmentInstanceOutput, error) {
	req, out := c.deleteEnvironmentInstanceRequest(ctx, input)
	return out, req.Send()
}
func (c *ProvisioningV1) deleteEnvironmentInstanceRequest(ctx context.Context, input *DeleteEnvironmentInstanceInput) (*request.Request, *DeleteEnvironmentInstanceOutput) {
	op := &request.Operation{
		Name: environments,
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/environments/{environmentInstanceId}",
		},
	}

	if input == nil {
		input = &DeleteEnvironmentInstanceInput{}
	}

	output := &DeleteEnvironmentInstanceOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PATCH /provisioning/v1/environments/{environmentInstanceId}
// Update an environment instance
type UpdateEnvironmentInstanceInput struct {
	//ID of the environment instance to delete
	EnvironmentInstanceId string `dest:"uri" dest-name:"environmentInstanceId"`

	//Name of the service plan for the environment instance. Must match the name in the corresponding service broker's
	//catalog. (for example: Subscription)
	PlanName string `json:"planName"`
	//If needed, you can pass environment-specific configuration parameters using a valid embedded JSON object.
	//For a list of supported configuration parameters, see the documentation of the particular environment offering.
	//In this example, additional configuration parameter 'instance_name' is specified:
	//{
	//"instance_name": "myOrg"
	//}
	Parameters map[string]interface{} `json:"parameters"`
}
type UpdateEnvironmentInstanceOutput struct {
	EnvironmentInstance
}

func (c *ProvisioningV1) UpdateEnvironmentInstance(ctx context.Context, input *UpdateEnvironmentInstanceInput) (*UpdateEnvironmentInstanceOutput, error) {
	req, out := c.updateEnvironmentInstanceRequest(ctx, input)
	return out, req.Send()
}
func (c *ProvisioningV1) updateEnvironmentInstanceRequest(ctx context.Context, input *UpdateEnvironmentInstanceInput) (*request.Request, *UpdateEnvironmentInstanceOutput) {
	op := &request.Operation{
		Name: environments,
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/environments/{environmentInstanceId}",
		},
	}

	if input == nil {
		input = &UpdateEnvironmentInstanceInput{}
	}

	output := &UpdateEnvironmentInstanceOutput{}
	return c.newRequest(ctx, op, input, output), output
}
