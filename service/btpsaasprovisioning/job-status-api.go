package btpsaasprovisioning

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

const jobManagement = "Job Management"

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

	types.StatusAndBodyFromResponse
}

func (c *SaaSProvisioningV1) GetJobStatus(ctx context.Context, input *GetJobStatusInput) (*GetJobStatusOutput, error) {
	req, out := c.getJobStatusRequest(ctx, input)
	return out, req.Send()
}
func (c *SaaSProvisioningV1) getJobStatusRequest(ctx context.Context, input *GetJobStatusInput) (*request.Request, *GetJobStatusOutput) {
	op := &request.Operation{
		Name: "Job Status Management",
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

// GET /api/v2.0/jobs/{jobUuid}
// Get job errorStatusCode
type GetErrorJobStatusInput struct {
	//The unique ID of a job for which to get information.
	JobUuid string `dest:"uri" dest-name:"jobUuid"`
}
type GetErrorJobStatusOutput struct {
	//The service instance ID of the SAP SaaS Provisioning service (saas-registry) that the application is using.
	CreatedBy string `json:"description,omitempty"`
	//ID of the corresponding job.
	Id string `json:"id,omitempty"`

	//The current state of the corresponding job. Possible values:
	//CREATED: Job processing has created.
	//STARTED: Job processing has started.
	//SUCCEEDED: The job has completed.
	//FAILED: The job failed and did not complete.
	//RETRY: Subscription has timed out and job processing is pending a retry.
	//Enum:
	//	[ CREATED, STARTED, SUCCEEDED, FAILED, RETRY ]
	State string   `json:"state,omitempty"`
	Error ErrorJob `json:"error,omitempty"`

	types.StatusAndBodyFromResponse
}
type ErrorJob struct {
	//Description of the error.
	Description string `json:"error,omitempty"`
	//The runtime exception for the error.
	Exception string `json:"exception,omitempty"`
	//The message associated with the current error.
	Message string `json:"message,omitempty"`
	//Path of the exception received from the server.
	Paths string `json:"paths,omitempty"`
	//Error status code.
	Status    int32     `json:"status,omitempty"`
	Timestamp Timestamp `json:"timestamp,omitempty"`
}
type Timestamp struct {
	Date           int32 `json:"date,omitempty"`
	Day            int32 `json:"day,omitempty"`
	Hours          int32 `json:"hours,omitempty"`
	Minutes        int32 `json:"minutes,omitempty"`
	Month          int32 `json:"month,omitempty"`
	Nanos          int32 `json:"nanos,omitempty"`
	Seconds        int32 `json:"seconds,omitempty"`
	Time           int64 `json:"time,omitempty"`
	TimezoneOffset int32 `json:"timezoneOffset,omitempty"`
	Year           int32 `json:"year,omitempty"`
}

func (c *SaaSProvisioningV1) GetErrorJobStatus(ctx context.Context, input *GetErrorJobStatusInput) (*GetErrorJobStatusOutput, error) {
	req, out := c.getErrorJobStatusRequest(ctx, input)
	return out, req.Send()
}
func (c *SaaSProvisioningV1) getErrorJobStatusRequest(ctx context.Context, input *GetErrorJobStatusInput) (*request.Request, *GetErrorJobStatusOutput) {
	op := &request.Operation{
		Name: jobManagement,
		Http: request.HTTP{
			Method:      request.GET,
			Path:        "/api/v2.0/jobs/{jobUuid}",
			UsePathAsIs: true,
		},
	}

	if input == nil {
		input = &GetErrorJobStatusInput{}
	}

	output := &GetErrorJobStatusOutput{}
	return c.newRequest(ctx, op, input, output), output
}
