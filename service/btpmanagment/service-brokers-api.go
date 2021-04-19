package btpmanagment

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

const serviceBrokers = "Service Management - Brokers"

// GET /v1/service_brokers
// Get all service brokers
type GetServiceBrokersInput struct {
	//Filters the response based on the field query.
	//If used, must be a nonempty string.
	//For example:
	//name eq 'my service broker'.
	FieldQuery string `dest:"querystring" dest-name:"fieldQuery"`

	//Filters the response based on the label query.
	//If used, must be a nonempty string.
	//For example:
	//environment eq 'dev'.
	LabelQuery string `dest:"querystring" dest-name:"labelQuery"`

	//You get this parameter in the response list of the API if the total number of items to return (num_items) is
	//larger than the number of items returned in a single API call (max_items).
	//You get a different token in each response to be used in each consecutive call as long as there are more items to list.
	//Use the returned tokens to get the full list of resources associated with your subaccount.
	//Leave the field empty if this is the first time you are calling the API.
	Token string `dest:"querystring" dest-name:"token"`

	//The maximum number of service brokers to return in the response.
	MaxItems int64 `dest:"querystring" dest-name:"max_items"`
}
type GetServiceBrokersOutput struct {
	Error

	//Use this token when you call the API again to get more service brokers associated with your subaccount.
	//The token field indicates that the total number of service brokers to view in the list (num_items) is larger than
	//the defined maximum number of service brokers to be returned after a single API call (max_items).
	//If the field is not present, either all the service brokers were included in the first response, or you have
	//reached the end of the list.
	Token string `json:"token,omitempty"`

	//The number of the service brokers associated with the subaccount.
	NumItems int64 `json:"num_items,omitempty"`

	//The list of response objects that contains details about the service brokers.
	Items []BrokerItem `json:"items,omitempty"`

	types.StatusAndBodyFromResponse
}
type BrokerItem struct {
	//The ID of the service broker.
	Id string `json:"id,omitempty"`
	//Whether the service broker is ready.
	Ready bool `json:"ready,omitempty"`
	//The name of the service broker.
	Name string `json:"name,omitempty"`
	//The description of the service broker.
	Description string `json:"description,omitempty"`
	//The URL of the service broker.
	BrokerUrl string `json:"broker_url,omitempty"`

	//The time the service broker was created.
	//In ISO 8601 format:
	//	YYYY-MM-DDThh:mm:ssTZD
	CreatedAt string `json:"created_at,omitempty"`

	//The last time the service broker was updated.
	//In ISO 8601 format.
	UpdatedAt string `json:"updated_at,omitempty"`

	//Additional data associated with the resource entity.
	Labels map[string][]string `json:"labels,omitempty"`
}

func (c *ServiceManagementV1) GetServiceBrokers(ctx context.Context,
	input *GetServiceBrokersInput) (*GetServiceBrokersOutput, error) {
	req, out := c.getServiceBrokersRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServiceBrokersRequest(ctx context.Context,
	input *GetServiceBrokersInput) (*request.Request, *GetServiceBrokersOutput) {
	op := &request.Operation{
		Name: serviceBrokers,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_brokers",
		},
	}

	if input == nil {
		input = &GetServiceBrokersInput{}
	}

	output := &GetServiceBrokersOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /v1/service_brokers/{serviceBrokerID}
// Get service broker details
type GetServiceBrokerDetailsInput struct {
	//The ID of the service broker for which to get details.
	ServiceBrokerID string `dest:"uri" dest-name:"serviceBrokerID"`
}
type GetServiceBrokerDetailsOutput struct {
	Error

	BrokerItem
	LastOperation Operation `json:"last_operation,omitempty"`

	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) GetServiceBrokerDetails(ctx context.Context,
	input *GetServiceBrokerDetailsInput) (*GetServiceBrokerDetailsOutput, error) {
	req, out := c.getServiceBrokerDetailsRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getServiceBrokerDetailsRequest(ctx context.Context,
	input *GetServiceBrokerDetailsInput) (*request.Request, *GetServiceBrokerDetailsOutput) {
	op := &request.Operation{
		Name: serviceBrokers,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/service_brokers/{serviceBrokerID}",
		},
	}

	if input == nil {
		input = &GetServiceBrokerDetailsInput{}
	}

	output := &GetServiceBrokerDetailsOutput{}
	return c.newRequest(ctx, op, input, output), output
}
