package btpmanagment

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

const serviceBindings = "Service Management - Bindings"

// GET /v1/service_bindings
// Get all service bindings
type GetServiceBindingsInput struct {
	//Filters the response based on the field query.
	//If used, must be a nonempty string.
	//For example:
	//	ready eq 'true'
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
	//The maximum number of service bindings to return in the response.
	MaxItems int64 `dest:"querystring" dest-name:"max_items"`
}
type GetServiceBindingsOutput struct {
	Error

	//Use this token when you call the API again to get more service bindings associated with your subaccount.
	//The token field indicates that the total number of service bindings to view in the list (num_items) is larger
	//than the defined maximum number of service bindings to be returned after a single API call (max_items).
	//If the field is not present, either all the service bindings were included in the first response, or you have
	//reached the end of the list.
	Token string `json:"token"`
	//The number of the service bindings associated with the subaccount.
	NumItems int64 `json:"num_items"`
	//The list of the response objects that contains details about the service bindings.
	Items []BindingItem `json:"items"`

	types.StatusAndBodyFromResponse
}
type BindingItem struct {
	//The ID of the service binding.
	Id string `json:"id"`
	//Whether the service binding is ready.
	Ready bool `json:"ready"`
	//The name of the service binding.
	Name string `json:"name"`
	//The ID of the service instance associated with the binding.
	ServiceInstanceId string `json:"service_instance_id"`
	//Contextual data for the resource.
	Context map[string]string `json:"context"`
	//Contains the resources associated with the binding.
	BindResource map[string]string `json:"bind_resource"`
	//Credentials to access the binding.
	Credentials Credentials `json:"credentials"`
	//The time the binding was created.
	//In ISO 8601 format:
	//	YYYY-MM-DDThh:mm:ssTZD
	CreatedAt string `json:"created_at"`
	//The last time the binding was updated.
	//In ISO 8601 format.
	UpdatedAt string `json:"updated_at"`
	//Additional data associated with the resource entity.
	Labels map[string][]string `json:"labels"`
}

func (c *ServiceManagementV1) GetServiceBindings(ctx context.Context,
	input *GetServiceBindingsInput) (*GetServiceBindingsOutput, error) {
	req, out := c.getServiceBindingsRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServiceBindingsRequest(ctx context.Context,
	input *GetServiceBindingsInput) (*request.Request, *GetServiceBindingsOutput) {
	op := &request.Operation{
		Name: serviceBindings,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_bindings",
		},
	}

	if input == nil {
		input = &GetServiceBindingsInput{}
	}

	output := &GetServiceBindingsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /v1/service_bindings
// Create a service binding
type CreateServiceBindingInput struct {
	//Whether to perform this operation asynchronously.
	Async bool `dest:"querystring" dest-name:"async"`

	//The name of the service binding.
	Name string `json:"name"`
	//The id of the service instance associated with the binding.
	ServiceInstanceId string `json:"service_instance_id"`
	//Some services support providing of additional configuration parameters during binding creation.
	//Pass these parameters as key-value pairs.
	//For the list of supported configuration parameters, see the documentation of a particular service offering.
	//You can also use the GET /v1/service_bindings/{serviceBindingID}/parameters API later to view the parameters
	//defined during this step.
	Parameters map[string]string `json:"parameters"`
	//The bind_resource object contains platform-specific information related to the context in which the service is used.
	//The examples of some common fields to use:
	//app_guid - A string GUID of an application associated with the binding. For credentials bindings.
	//Must be unique within the scope of the platform.
	//app_guid - Represents the scope to which the binding applies within the platform.
	//For example, in Kubernetes it can map to a namespace.
	//The scope of what the platform maps the app_guid to is platform-specific and can vary across binding requests.
	//route - URL of the intermediate application. For route services bindings.
	BindResource map[string]string `json:"bind_resource"`
	//Additional data associated with the resource entity.
	Labels map[string][]string `json:"labels"`
}
type CreateServiceBindingOutput struct {
	Error

	BindingItem

	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) CreateServiceBinding(ctx context.Context,
	input *CreateServiceBindingInput) (*CreateServiceBindingOutput, error) {
	req, out := c.createServiceBindingRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) createServiceBindingRequest(ctx context.Context,
	input *CreateServiceBindingInput) (*request.Request, *CreateServiceBindingOutput) {
	op := &request.Operation{
		Name: serviceBindings,
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/service_bindings",
		},
	}

	if input == nil {
		input = &CreateServiceBindingInput{}
	}

	output := &CreateServiceBindingOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /v1/service_bindings/{serviceBindingID}
// Get service Binding details
type GetServiceBindingDetailsInput struct {
	//The ID of the service binding for which to get details.
	ServiceBindingID string `dest:"uri" dest-name:"serviceBindingID"`
}
type GetServiceBindingDetailsOutput struct {
	Error

	BindingItem
	LastOperation Operation `json:"last_operation"`

	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) GetServiceBindingDetails(ctx context.Context,
	input *GetServiceBindingDetailsInput) (*GetServiceBindingDetailsOutput, error) {
	req, out := c.getServiceBindingDetailsRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServiceBindingDetailsRequest(ctx context.Context,
	input *GetServiceBindingDetailsInput) (*request.Request, *GetServiceBindingDetailsOutput) {
	op := &request.Operation{
		Name: serviceBindings,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_bindings/{serviceBindingID}",
		},
	}

	if input == nil {
		input = &GetServiceBindingDetailsInput{}
	}

	output := &GetServiceBindingDetailsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /v1/service_bindings/{serviceBindingID}
// Delete a service Binding
type DeleteServiceBindingInput struct {
	//The ID of the service binding to delete.
	ServiceBindingID string `dest:"uri" dest-name:"serviceBindingID"`
	//Whether to perform this operation asynchronously.
	Async string `dest:"querystring" dest-name:"async"`
}
type DeleteServiceBindingOutput struct {
	Error

	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) DeleteServiceBinding(ctx context.Context,
	input *DeleteServiceBindingInput) (*DeleteServiceBindingOutput, error) {
	req, out := c.deleteServiceBindingRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) deleteServiceBindingRequest(ctx context.Context,
	input *DeleteServiceBindingInput) (*request.Request, *DeleteServiceBindingOutput) {
	op := &request.Operation{
		Name: serviceBindings,
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/service_bindings/{serviceBindingID}",
		},
	}

	if input == nil {
		input = &DeleteServiceBindingInput{}
	}

	output := &DeleteServiceBindingOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PATCH /v1/service_bindings/{serviceBindingID}
// Update a service Binding
/*type UpdateServiceBindingInput struct {
	ServiceBindingID string `dest:"uri" dest-name:"serviceBindingID"`

	Label          string              `json:"name"`
	ServicePlanId string              `json:"service_plan_id"`
	Parameters    map[string]string   `json:"parameters"`
	Labels        map[string][]string `json:"labels"`
}
type UpdateServiceBindingOutput struct {
	Error
}

func (c *ServiceManagementV1) UpdateServiceBinding(ctx context.Context,
	input *UpdateServiceBindingInput) (*UpdateServiceBindingOutput, error) {
	req, out := c.updateServiceBindingRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) updateServiceBindingRequest(ctx context.Context,
	input *UpdateServiceBindingInput) (*request.Request, *UpdateServiceBindingOutput) {
	op := &request.Operation{
		Label: serviceBindings,
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/service_bindings/{serviceBindingID}",
		},
	}

	if input == nil {
		input = &UpdateServiceBindingInput{}
	}

	output := &UpdateServiceBindingOutput{}
	return c.newRequest(ctx, op, input, output), output
}*/

// GET /v1/service_bindings/{serviceBindingID}/parameters
// Get service Binding parameters
type GetServiceBindingParametersInput struct {
	//The ID of the service binding for which to get parameters.
	ServiceBindingID string `dest:"uri" dest-name:"serviceBindingID"`
}
type GetServiceBindingParametersOutput struct {
	Error

	Parameters map[string]string `json:"-"`

	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) GetServiceBindingParameters(ctx context.Context,
	input *GetServiceBindingParametersInput) (*GetServiceBindingParametersOutput, error) {
	req, out := c.getServiceBindingParametersRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServiceBindingParametersRequest(ctx context.Context,
	input *GetServiceBindingParametersInput) (*request.Request, *GetServiceBindingParametersOutput) {
	op := &request.Operation{
		Name: serviceBindings,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_bindings/{serviceBindingID}/parameters",
		},
	}

	if input == nil {
		input = &GetServiceBindingParametersInput{}
	}

	output := &GetServiceBindingParametersOutput{}
	return c.newRequest(ctx, op, input, output), output
}
