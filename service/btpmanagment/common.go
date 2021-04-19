package btpmanagment

type Operation struct {
	//The ID of the operation.
	Id string `json:"id,omitempty"`

	//Whether the resource is ready.
	Ready bool `json:"ready,omitempty"`

	//The type of the operation.
	//Possible values:
	//Enum:
	//	[ CREATE, UPDATE, DELETE ]
	Type string `json:"type,omitempty"`

	//Valid values are: in progress, succeeded, and failed.
	//While the state is "in progress", the platform should continue polling.
	//The responses: "state": "succeeded" or "state": "failed" must cause the platform to stop polling.
	//Enum:
	//	[ in progress, succeeded, failed ]
	State string `json:"state,omitempty"`

	//Details about the operation for customer-facing UI.
	Description string `json:"description,omitempty"`

	//The ID of the resource.
	//Exists if: "state": "succeeded", and also for PATCH and DELETE requests
	ResourceId string `json:"resource_id,omitempty"`

	TransitiveResources []TransitiveResource `json:"transitive_resources,omitempty"`

	//The type of the resource (e.g. service_brokers, service_instances).
	ResourceType string `json:"resource_type,omitempty"`

	//The ID of the platform associated with the operation.
	PlatformId string `json:"platform_id,omitempty"`

	//The correlation ID received from the request related to this operation.
	CorrelationId string `json:"correlation_id,omitempty"`

	//Whether the operation has reached a checkpoint and can be executed again.
	Reschedule bool `json:"reschedule,omitempty"`

	//The time the resource is scheduled for deletion.
	//In ISO 8601 format:
	//	YYYY-MM-DDThh:mm:ssTZD
	DeletionScheduled string `json:"deletion_scheduled,omitempty"`

	//The time the resource was created.
	//In ISO 8601 format.
	CreatedAt string `json:"created_at,omitempty"`

	//The last time the resource was updated.
	//In ISO 8601 format.
	//Recommended field if "state": "succeeded" or "state": "failed".
	UpdatedAt string `json:"updated_at,omitempty"`

	//The list of the errors if the operation has failed.
	Errors []Error `json:"errors,omitempty"`

	//Additional data associated with the resource entity.
	Labels map[string][]string `json:"labels,omitempty"`
}
type TransitiveResource struct {
	//The ID of the resource.
	Id string `json:"id,omitempty"`

	//The type of the resource.
	Type string `json:"type,omitempty"`

	//The type of the operation associated with the resource.
	OperationType string `json:"operation_type,omitempty"`

	//The minimum criteria required to use the resource in the context of the platform.
	Criteria string `json:"criteria,omitempty"`
}

//A response object that contains details about the error.
type Error struct {
	//The name of the error.
	ErrorMessage string `json:"error,omitempty"`

	//The description of the error.
	ErrorDescription string `json:"description,omitempty"`
}

type Credentials struct {
	Basic Basic `json:"basic,omitempty"`
}
type Basic struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
