package btpmanagment

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

const servicePlans = "Service Management - Plans"

// GET /v1/service_plans
// Get all service plans
type GetServicePlansInput struct {
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
	//You get this parameter in the response list of the API if the total number of items to return (num_items) is
	//larger than the number of items returned in a single API call (max_items).
	//You get a different token in each response to be used in each consecutive call as long as there are more items to list.
	//Use the returned tokens to get the full list of resources associated with your subaccount.
	//If this is the first time you are calling this API, leave this field empty.
	Token string `dest:"querystring" dest-name:"token"`
	//The maximum number of service plans to return in the response.
	MaxItems int64 `dest:"querystring" dest-name:"max_items"`
}
type GetServicePlansOutput struct {
	//Use this token when you call the API again to get more service plans associated with your subaccount.
	//The token field indicates that the total number of service plans to view in the list (num_items) is larger than
	//the defined maximum number of ervice plans to be returned after a single API call (max_items).
	//If the field is not present, either all the service plans were included in the first response, or you have reached
	//the end of the list.
	Token string `json:"token,omitempty"`
	//The number of service plans associated with the subaccount.
	NumItems int64 `json:"num_items,omitempty"`
	//The list of the response objects that contain details about the service plans.
	Items []PlanItem `json:"items,omitempty"`

	Error
	types.StatusAndBodyFromResponse
}
type PlanItem struct {
	//The ID of the service plan.
	Id string `json:"id,omitempty"`
	//Whether the service plan is ready.
	Ready bool `json:"ready,omitempty"`
	//The name of the service plan.
	Name string `json:"name,omitempty"`
	//The description of the service plan.
	Description string `json:"description,omitempty"`
	//The ID of the service plan in the service broker catalog.
	CatalogId string `json:"catalog_id,omitempty"`
	//The name of the associated service broker catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	//Whether the service plan is free.
	Free bool `json:"free,omitempty"`
	//Whether the service plan is bindable.
	Bindable bool         `json:"bindable,omitempty"`
	Metadata PlanMetadata `json:"metadata,omitempty"`
	//The ID of the service offering.
	ServiceOfferingId string `json:"service_offering_id,omitempty"`
	//The time the service plan was created.
	//In ISO 8601 format:
	//	YYYY-MM-DDThh:mm:ssTZD
	CreatedAt string `json:"created_at,omitempty"`
	//The last time the service plan was updated.
	//In ISO 8601 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}
type PlanMetadata struct {
	//Platforms supported by the service plan.
	//Enum:
	//	[ kubernetes, cloudfoundry ]
	SupportedPlatforms []string `json:"supportedPlatforms,omitempty"`
	//The earliest supported OSB version.
	SupportedMinOSBVersion string `json:"supportedMinOSBVersion,omitempty"`
	//The latest supported OSB version.
	SupportedMaxOSBVersion string `json:"supportedMaxOSBVersion,omitempty"`
}

func (c *ServiceManagementV1) GetServicePlans(ctx context.Context,
	input *GetServicePlansInput) (*GetServicePlansOutput, error) {
	req, out := c.getServicePlansRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServicePlansRequest(ctx context.Context,
	input *GetServicePlansInput) (*request.Request, *GetServicePlansOutput) {
	op := &request.Operation{
		Name: servicePlans,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_plans",
		},
	}

	if input == nil {
		input = &GetServicePlansInput{}
	}

	output := &GetServicePlansOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /v1/service_plans/{servicePlanID}
// Get service plan details
type GetServicePlanInput struct {
	//The ID of the service plan for which to get details.
	ServicePlanID string `dest:"uri" dest-name:"servicePlanID"`
}
type GetServicePlanOutput struct {
	PlanItem

	Error
	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) GetServicePlanDetails(ctx context.Context,
	input *GetServicePlanInput) (*GetServicePlanOutput, error) {
	req, out := c.getServicePlanRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServicePlanRequest(ctx context.Context,
	input *GetServicePlanInput) (*request.Request, *GetServicePlanOutput) {
	op := &request.Operation{
		Name: servicePlans,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_plans/{servicePlanID}",
		},
	}

	if input == nil {
		input = &GetServicePlanInput{}
	}

	output := &GetServicePlanOutput{}
	return c.newRequest(ctx, op, input, output), output
}
