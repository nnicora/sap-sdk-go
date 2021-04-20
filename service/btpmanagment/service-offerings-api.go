package btpmanagment

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

const serviceOfferings = "Service Management - Offerings"

// GET /v1/service_offerings
// Get all service offerings
type GetServiceOfferingsInput struct {
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
	//The maximum number of service offerings to return in the response.
	MaxItems int64 `dest:"querystring" dest-name:"max_items"`
}
type GetServiceOfferingsOutput struct {
	//Use this token when you call the API again to get more service offerings associated with your subaccount.
	//The token field indicates that the total number of service offerings to view in the list (num_items) is larger
	//than the defined maximum number of service offerings to be returned after a single API call (max_items).
	//If the field is not present, either all the service offerings were included in the first response, or you have
	//reached the end of the list.
	Token string `json:"token,omitempty"`
	//The number of service offerings associated with the subaccount.
	NumItems int64 `json:"num_items,omitempty"`
	//The list of the service offerings.
	Items []OfferingItem `json:"items,omitempty"`

	Error
	types.StatusAndBodyFromResponse
}
type OfferingItem struct {
	//The ID of the service offering.
	Id string `json:"id,omitempty"`
	//Whether the service offering is ready to be advertised.
	Ready bool `json:"ready,omitempty"`
	//The name of the service offering.
	Name string `json:"name,omitempty"`
	//The description of the service offering.
	Description string `json:"description,omitempty"`
	//Whether the service offering is bindable.
	Bindable bool `json:"bindable,omitempty"`
	//Whether the service instances associated with the service offering can be retrieved.
	InstancesRetrievable bool `json:"instances_retrievable,omitempty"`
	//Whether the bindings associated with the service offering can be retrieved.
	BindingsRetrievable bool `json:"bindings_retrievable,omitempty"`
	//Whether the offered plan can be updated.
	PlanUpdateable bool `json:"plan_updateable,omitempty"`
	//Whether the context for the service offering can be updated.
	AllowContextUpdates bool `json:"allow_context_updates,omitempty"`
	//The list of tags for the service offering.
	Tags     []string         `json:"tags,omitempty"`
	Metadata OfferingMetadata `json:"metadata,omitempty"`
	//The ID of the broker that provides the service plan.
	BrokerId string `json:"broker_id,omitempty"`
	//The ID of the service offering as provided by the catalog.
	CatalogId string `json:"catalog_id,omitempty"`
	//The catalog name of the service offering.
	CatalogName string `json:"catalog_name,omitempty"`
	//The time the service offering was created.
	//In ISO 8601 format:
	//	YYYY-MM-DDThh:mm:ssTZD
	CreatedAt string `json:"created_at,omitempty"`
	//The last time the service offering was updated.
	//In ISO 8601 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}
type OfferingMetadata struct {
	//The description of the service offering.
	LongDescription string `json:"longDescription,omitempty"`
	//The URL to the associated documentation.
	DocumentationUrl string `json:"documentationUrl,omitempty"`
	//The name of the service offering for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`
	//The URL to the associated image.
	ImageUrl string `json:"imageUrl,omitempty"`
	//The support URL for the service offering.
	SupportUrl string `json:"supportUrl,omitempty"`
}

func (c *ServiceManagementV1) GetServiceOfferings(ctx context.Context,
	input *GetServiceOfferingsInput) (*GetServiceOfferingsOutput, error) {
	req, out := c.getServiceOfferingsRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServiceOfferingsRequest(ctx context.Context,
	input *GetServiceOfferingsInput) (*request.Request, *GetServiceOfferingsOutput) {
	op := &request.Operation{
		Name: serviceOfferings,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_offerings",
		},
	}

	if input == nil {
		input = &GetServiceOfferingsInput{}
	}

	output := &GetServiceOfferingsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /v1/service_offerings/{serviceOfferingID}
// Get service offering details
type GetServiceOfferingDetailsInput struct {
	//The ID of the service offering for which to get details.
	ServiceOfferingID string `dest:"uri" dest-name:"serviceOfferingID"`
}
type GetServiceOfferingDetailsOutput struct {
	OfferingItem

	Error
	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) GetServiceOfferingDetails(ctx context.Context,
	input *GetServiceOfferingDetailsInput) (*GetServiceOfferingDetailsOutput, error) {
	req, out := c.getServiceOfferingDetailsRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServiceOfferingDetailsRequest(ctx context.Context,
	input *GetServiceOfferingDetailsInput) (*request.Request, *GetServiceOfferingDetailsOutput) {
	op := &request.Operation{
		Name: serviceOfferings,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_offerings/{serviceOfferingID}",
		},
	}

	if input == nil {
		input = &GetServiceOfferingDetailsInput{}
	}

	output := &GetServiceOfferingDetailsOutput{}
	return c.newRequest(ctx, op, input, output), output
}
