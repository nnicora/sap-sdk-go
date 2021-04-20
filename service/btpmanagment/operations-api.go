package btpmanagment

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

const operations = "Service Management - Operations"

// GET /v1/{resourceType}/{resourceID}/operations/{operationID}
// Get operation status
type GetOperationStatusInput struct {
	//The type of the SAP Cloud Service Management service resource.
	//Available values : platforms, service_brokers, service_bindings, service_instances
	ResourceType string `dest:"uri" dest-name:"resourceType"`
	//The ID of the previously created entity of the specified resource type.
	ResourceID string `dest:"uri" dest-name:"resourceID"`
	//The ID of the operation for which to get status.
	OperationID string `dest:"uri" dest-name:"operationID"`
}
type GetOperationStatusOutput struct {
	Operation

	Error
	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) GetOperationStatus(ctx context.Context,
	input *GetOperationStatusInput) (*GetOperationStatusOutput, error) {
	req, out := c.getOperationStatusRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getOperationStatusRequest(ctx context.Context,
	input *GetOperationStatusInput) (*request.Request, *GetOperationStatusOutput) {
	op := &request.Operation{
		Name: operations,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/{resourceType}/{resourceID}/operations/{operationID}",
		},
	}

	if input == nil {
		input = &GetOperationStatusInput{}
	}

	output := &GetOperationStatusOutput{}
	return c.newRequest(ctx, op, input, output), output
}
