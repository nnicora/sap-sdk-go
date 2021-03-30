package btpmanagment

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
)

const serviceInstances = "Service Management - Instances"

// GET /v1/service_instances
// Get all service instances
type GetServiceInstancesInput struct {
	//Filters the response based on the field query.
	//If used, must be a nonempty string.
	//For example:
	//	usable eq 'true'
	FieldQuery string `dest:"querystring" dest-name:"fieldQuery"`
	//Filters the response based on the label query.
	//If used, must be a nonempty string.
	//For example:
	//	environment eq 'dev'
	LabelQuery string `dest:"querystring" dest-name:"labelQuery"`
	//You get this parameter in the response list of the API if the total number of items to return (num_items) is larger
	//than the number of items returned in a single API call (max_items).
	//You get a different token in each response to be used in each consecutive call as long as there are more items to list.
	//Use the returned tokens to get the full list of resources associated with your subaccount.
	//If this is the first time you are calling the API, leave this field empty.
	Token string `dest:"querystring" dest-name:"token"`
	//The maximum number of service instances to return in the response.
	MaxItems int64 `dest:"querystring" dest-name:"max_items"`
}
type GetServiceInstancesOutput struct {
	Error

	//Use this token when you call the API again to get more service instances associated with your subaccount.
	//The token field indicates that the total number of service instances to view in the list (num_items) is larger
	//than the defined maximum number of service instances to be returned after a single API call (max_items).
	//If the field is not present, either all the instances were included in the first response, or you have reached the
	//end of the list.
	Token string `json:"token"`
	//The number of service instances associated with the subaccount.
	NumItems int64 `json:"num_items"`
	//The list of response objects that contains details about the service instances.
	Items []InstanceItem `json:"items"`
}
type InstanceItem struct {
	//The ID of the service instance.
	Id string `json:"id"`
	//Whether the service instance is ready.
	Ready bool `json:"ready"`
	//The name of the service instance.
	Name string `json:"name"`
	//The ID of the service plan associated with the service instance.
	ServicePlanId string `json:"service_plan_id"`
	//The ID of the platform to which the service instance belongs.
	PlatformId string `json:"platform_id"`
	//The URL of the web-based management UI for the service instance.
	DashboardUrl string `json:"dashboard_url"`
	//Contextual data for the resource.
	Context map[string]string `json:"context"`
	//The maintenance information associated with the service instance.
	MaintenanceInfo map[string]string `json:"maintenance_info"`
	//Whether the service instance can be used.
	Usable bool `json:"usable"`
	//The time the service instance was created.
	//In ISO 8601 format:
	//	YYYY-MM-DDThh:mm:ssTZD
	CreatedAt string `json:"created_at"`
	//The last time the service instance was updated.
	//In ISO 8601 format.
	UpdatedAt string `json:"updated_at"`
	//Additional data associated with the resource entity.
	Labels map[string][]string `json:"labels"`
}

func (c *ServiceManagementV1) GetServiceInstances(ctx context.Context,
	input *GetServiceInstancesInput) (*GetServiceInstancesOutput, error) {
	req, out := c.getServiceInstancesRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServiceInstancesRequest(ctx context.Context,
	input *GetServiceInstancesInput) (*request.Request, *GetServiceInstancesOutput) {
	op := &request.Operation{
		Name: serviceInstances,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_instances",
		},
	}

	if input == nil {
		input = &GetServiceInstancesInput{}
	}

	output := &GetServiceInstancesOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /v1/service_instances
// Create a service instance
type CreateServiceInstanceInput struct {
	//Whether to perform this operation asynchronously.
	Async bool `dest:"querystring" dest-name:"async"`

	//The name of the new service instance.
	//Can't be an empty object.
	Name string `json:"name"`
	//The ID of the service plan to use for the service instance.
	ServicePlanId string `json:"service_plan_id"`
	//Some services support providing of additional configuration parameters during instance creation.
	//Pass these parameters as key-value pairs.
	//For the list of supported configuration parameters, see the documentation of a particular service offering.
	//You can also use the GET /v1/service_instances/{serviceInstanceID}/parameters API later to view the parameters
	//defined during this step.
	Parameters map[string]string `json:"parameters"`
	//Additional data associated with the resource entity.
	Labels map[string][]string `json:"labels"`
}
type CreateServiceInstanceOutput struct {
	Error

	InstanceItem
}

func (c *ServiceManagementV1) CreateServiceInstance(ctx context.Context,
	input *CreateServiceInstanceInput) (*CreateServiceInstanceOutput, error) {
	req, out := c.createServiceInstanceRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) createServiceInstanceRequest(ctx context.Context,
	input *CreateServiceInstanceInput) (*request.Request, *CreateServiceInstanceOutput) {
	op := &request.Operation{
		Name: serviceInstances,
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/service_instances",
		},
	}

	if input == nil {
		input = &CreateServiceInstanceInput{}
	}

	output := &CreateServiceInstanceOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /v1/service_instances/{serviceInstanceID}
// Get service instance details
type GetServiceInstanceDetailsInput struct {
	//The ID of the provisioned service instance for which to get details.
	ServiceInstanceID string `dest:"uri" dest-name:"serviceInstanceID"`
}
type GetServiceInstanceDetailsOutput struct {
	Error

	InstanceItem
	LastOperation Operation `json:"last_operation"`
}

func (c *ServiceManagementV1) GetServiceInstanceDetails(ctx context.Context,
	input *GetServiceInstanceDetailsInput) (*GetServiceInstanceDetailsOutput, error) {
	req, out := c.getServiceInstanceDetailsRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServiceInstanceDetailsRequest(ctx context.Context,
	input *GetServiceInstanceDetailsInput) (*request.Request, *GetServiceInstanceDetailsOutput) {
	op := &request.Operation{
		Name: serviceInstances,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_instances/{serviceInstanceID}",
		},
	}

	if input == nil {
		input = &GetServiceInstanceDetailsInput{}
	}

	output := &GetServiceInstanceDetailsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /v1/service_instances/{serviceInstanceID}
// Delete a service instance
type DeleteServiceInstanceInput struct {
	//The ID of the provisioned service instance to delete.
	ServiceInstanceID string `dest:"uri" dest-name:"serviceInstanceID"`

	//Whether to perform this operation asynchronously.
	Async string `dest:"querystring" dest-name:"async"`
}
type DeleteServiceInstanceOutput struct {
	Error
}

func (c *ServiceManagementV1) DeleteServiceInstance(ctx context.Context,
	input *DeleteServiceInstanceInput) (*DeleteServiceInstanceOutput, error) {
	req, out := c.deleteServiceInstanceRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) deleteServiceInstanceRequest(ctx context.Context,
	input *DeleteServiceInstanceInput) (*request.Request, *DeleteServiceInstanceOutput) {
	op := &request.Operation{
		Name: serviceInstances,
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/service_instances/{serviceInstanceID}",
		},
	}

	if input == nil {
		input = &DeleteServiceInstanceInput{}
	}

	output := &DeleteServiceInstanceOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PATCH /v1/service_instances/{serviceInstanceID}
// Update a service instance
type UpdateServiceInstanceInput struct {
	//The ID of the provisioned service instance to update.
	ServiceInstanceID string `dest:"uri" dest-name:"serviceInstanceID"`
	//Whether to perform this operation asynchronously.
	Async string `dest:"querystring" dest-name:"async"`

	//The name of the service instance to update.
	Name string `json:"name"`
	//The ID of the service plan for the service instance to update.
	ServicePlanId string `json:"service_plan_id"`
	//Some services support providing of additional configuration parameters during instance creation.
	//You can update these parameters.
	//For the list of supported configuration parameters, see the documentation of a particular service offering.
	//You can also use the GET /v1/service_instances/{serviceInstanceID}/parameters API later to view the parameters
	//defined during this step
	Parameters map[string]string `json:"parameters"`
	//The list of labels to update for the resource.
	Labels map[string][]string `json:"labels"`
}
type UpdateServiceInstanceOutput struct {
	Error
}

func (c *ServiceManagementV1) UpdateServiceInstance(ctx context.Context,
	input *UpdateServiceInstanceInput) (*UpdateServiceInstanceOutput, error) {
	req, out := c.updateServiceInstanceRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) updateServiceInstanceRequest(ctx context.Context,
	input *UpdateServiceInstanceInput) (*request.Request, *UpdateServiceInstanceOutput) {
	op := &request.Operation{
		Name: serviceInstances,
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/service_instances/{serviceInstanceID}",
		},
	}

	if input == nil {
		input = &UpdateServiceInstanceInput{}
	}

	output := &UpdateServiceInstanceOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /v1/service_instances/{serviceInstanceID}/parameters
// Get service instance parameters
type GetServiceInstanceParametersInput struct {
	//The ID of the provisioned service instance for which to get parameters.
	ServiceInstanceID string `dest:"uri" dest-name:"serviceInstanceID"`
}
type GetServiceInstanceParametersOutput struct {
	Error

	Parameters map[string]string `json:"-"`
}

func (c *ServiceManagementV1) GetServiceInstanceParameters(ctx context.Context,
	input *GetServiceInstanceParametersInput) (*GetServiceInstanceParametersOutput, error) {
	req, out := c.getServiceInstanceParametersRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServiceInstanceParametersRequest(ctx context.Context,
	input *GetServiceInstanceParametersInput) (*request.Request, *GetServiceInstanceParametersOutput) {
	op := &request.Operation{
		Name: serviceInstances,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_instances/{serviceInstanceID}/parameters",
		},
	}

	if input == nil {
		input = &GetServiceInstanceParametersInput{}
	}

	output := &GetServiceInstanceParametersOutput{}
	return c.newRequest(ctx, op, input, output), output
}
