package btpaccounts

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

// GET /jobs-management/v1/jobs/{jobInstanceIdOrUniqueId}/status
// Get available jobs
type GetJobStatusInput struct {
	//ID of the job for which to get status
	JobId string `dest:"uri" dest-name:"jobInstanceIdOrUniqueId"`
}
type GetJobStatusOutput struct {
	//A description of the exit status of a job when it ends.
	Description string `json:"description,omitempty"`

	//The current state of the job.
	//
	//IN_PROGRESS: The job is being executed.
	//COMPLETED: The job has completed.
	//FAILED: The job failed and did not complete. The job can be restarted.
	//Enum:
	//	[ IN_PROGRESS, COMPLETED, FAILED ]
	Status string `json:"status,omitempty"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) GetJobStatus(ctx context.Context, input *GetJobStatusInput) (*GetJobStatusOutput, error) {
	req, out := c.getJobStatusRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getJobStatusRequest(ctx context.Context, input *GetJobStatusInput) (*request.Request, *GetJobStatusOutput) {
	op := &request.Operation{
		Name: "Get Job Status",
		Http: request.HTTP{
			Method:      request.GET,
			Path:        "/jobs-management/v1/jobs/{jobInstanceIdOrUniqueId}/status",
			UsePathAsIs: true,
		},
	}

	if input == nil {
		input = &GetJobStatusInput{}
	}

	output := &GetJobStatusOutput{}
	return c.newRequest(ctx, op, input, output), output
}
