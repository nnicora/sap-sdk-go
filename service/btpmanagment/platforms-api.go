package btpmanagment

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

const platforms = "Service Management - Platforms"

// GET /v1/platforms
// Get all platforms
type GetPlatformsInput struct {
	//Filters the response based on the field query.
	//If used, must be a nonempty string.
	//For example:
	//	type eq 'kubernetes'
	FieldQuery string `dest:"querystring" dest-name:"fieldQuery"`

	//Filters the response based on the label query.
	//If used, must be a nonempty string.
	//For example:
	//	environment eq 'dev'
	LabelQuery string `dest:"querystring" dest-name:"labelQuery"`

	//You get this parameter in the response list of the API if the total number of items to return (num_items) is larger
	//than the number of items returned in a single API call (max_items).
	//You get a different token in each response to be used in each consecutive call as long as there are more items to list.
	//Use the returned tokens to get the full list of items associated with your subaccount.
	//If this is the first time you are calling the API, leave this field empty.
	Token string `dest:"querystring" dest-name:"token"`

	//The maximum number of platforms to return in the response.
	MaxItems int64 `dest:"querystring" dest-name:"max_items"`
}
type GetPlatformsOutput struct {
	Error

	//Use this token when you call the API again to get more platforms associated with your subaccount.
	//The token field indicates that the total number of platforms to view in the list (num_items) is larger than the
	//defined maximum number of platforms to be returned after a single API call (max_items).
	//If the field is not present, either all the platforms were included in the first response, or you have
	//reached the end of the list.
	Token string `json:"token,omitempty"`

	//The number of platforms associated with the subaccount.
	NumItems int64 `json:"num_items,omitempty"`

	//The list of response objects that contains details about the platforms.
	Items []PlatformItem `json:"items,omitempty"`

	types.StatusAndBodyFromResponse
}
type PlatformItem struct {
	//The ID of the platform.
	//You can use this ID to get details about the platform, to update or to delete it.
	//See the GET, PATCH, or DELETE APIs for the Platforms group.
	Id string `json:"id,omitempty"`

	//Whether the platform is ready for consumption.
	Ready bool `json:"ready,omitempty"`

	//The type of the platform.
	//Possible values:
	//Enum:
	//	[ kubernetes, cloud foundry ]
	Type string `json:"type,omitempty"`

	//The name of the platform.
	Name string `json:"name,omitempty"`

	//The description of the platform.
	Description string `json:"description,omitempty"`

	//The time the platform was created.
	//In ISO 8601 format:
	//	YYYY-MM-DDThh:mm:ssTZD
	CreatedAt string `json:"created_at,omitempty"`

	//The last time the platform was updated.
	//In ISO 8601 format.
	UpdatedAt string `json:"updated_at,omitempty"`

	//Additional data associated with the resource entity.
	Labels map[string][]string `json:"labels,omitempty"`
}

func (c *ServiceManagementV1) GetPlatforms(ctx context.Context, input *GetPlatformsInput) (*GetPlatformsOutput, error) {
	req, out := c.getPlatformsRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getPlatformsRequest(ctx context.Context,
	input *GetPlatformsInput) (*request.Request, *GetPlatformsOutput) {
	op := &request.Operation{
		Name: platforms,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/platforms",
		},
	}

	if input == nil {
		input = &GetPlatformsInput{}
	}

	output := &GetPlatformsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /v1/platforms
// Register a platform
type RegisterPlatformInput struct {
	//The CLI-friendly name of the platform.
	//A CLI-friendly name is a short string that only contains alphanumeric characters, periods, and hyphens.
	//It can't contain white spaces.
	//The name must not exceed 255 characters, but it is recommended to keep it much shorter, for the convenience
	//of using short names in CLI commands.
	Name string `json:"name,omitempty"`

	//The type of the platform.
	//Possible values:
	//Enum:
	//	[ kubernetes ]
	Type string `json:"type,omitempty"`

	//The description of the platform for customer-facing UIs.
	Description string `json:"description,omitempty"`

	//Additional data associated with the resource entity.
	Labels map[string][]string `json:"labels,omitempty"`
}
type RegisterPlatformOutput struct {
	Error

	// The ID of the platform.
	//You can use this ID to get details about the platform, to update, or to delete the platform.
	//See the GET, PATCH, and DELETE APIs for the Platforms group.
	Id string `json:"id,omitempty"`

	//Whether the platform is ready for consumption.
	Ready bool `json:"ready,omitempty"`

	//The type of the platform.
	//Possible values:
	//Enum:
	//	[ kubernetes ]
	Type string `json:"type,omitempty"`

	//The technical name of the platform.
	Name string `json:"name,omitempty"`

	//The description of the platform.
	Description string `json:"description,omitempty"`

	//Credentials to authenticate with the SAP Cloud Service Management service.
	Credentials Credentials `json:"credentials,omitempty"`

	//The time the platform was created.
	//In ISO 8601 format:
	//	YYYY-MM-DDThh:mm:ssTZD
	CreatedAt string `json:"created_at,omitempty"`

	//The last time the platform was updated.
	//In ISO 8601 format.
	UpdatedAt string `json:"updated_at,omitempty"`

	//Additional data associated with the resource entity.
	Labels map[string][]string `json:"labels,omitempty"`

	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) RegisterPlatform(ctx context.Context,
	input *RegisterPlatformInput) (*RegisterPlatformOutput, error) {
	req, out := c.registerPlatformRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) registerPlatformRequest(ctx context.Context,
	input *RegisterPlatformInput) (*request.Request, *RegisterPlatformOutput) {
	op := &request.Operation{
		Name: platforms,
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/platforms",
		},
	}

	if input == nil {
		input = &RegisterPlatformInput{}
	}

	output := &RegisterPlatformOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /v1/platforms/{platformID}
// Get all platforms
type GetPlatformDetailsInput struct {
	//The ID of the registered platform for which to get details.
	PlatformID string `dest:"uri" dest-name:"platformID"`
}
type GetPlatformDetailsOutput struct {
	Error

	//The ID of the platform.
	//You can use this ID to update or to delete the platform.
	//See the PATCH and DELETE calls for the Platforms group.
	Id string `json:"id,omitempty"`

	//Whether the platform is ready for consumption.
	Ready bool `json:"ready,omitempty"`

	LastOperation Operation `json:"last_operation,omitempty"`

	//The type of the platform.
	//Possible values:
	//Enum:
	//	[ kubernetes, cloud foundry ]
	Type string `json:"type,omitempty"`

	//The name of the platform.
	Name string `json:"name,omitempty"`

	//The description of the platform.
	Description string `json:"description,omitempty"`

	//The time the platform was created.
	//In ISO 8601 format:
	//	YYYY-MM-DDThh:mm:ssTZD
	CreatedAt string `json:"created_at,omitempty"`

	//The last time the platform was updated.
	//In ISO 8601 format.
	UpdatedAt string `json:"updated_at,omitempty"`

	//Additional data associated with the resource entity.
	Labels map[string][]string `json:"labels,omitempty"`

	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) GetPlatformDetails(ctx context.Context,
	input *GetPlatformDetailsInput) (*GetPlatformDetailsOutput, error) {
	req, out := c.getPlatformDetailsRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) getPlatformDetailsRequest(ctx context.Context,
	input *GetPlatformDetailsInput) (*request.Request, *GetPlatformDetailsOutput) {
	op := &request.Operation{
		Name: platforms,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/platforms/{platformID}",
		},
	}

	if input == nil {
		input = &GetPlatformDetailsInput{}
	}

	output := &GetPlatformDetailsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /v1/platforms/{platformID}
// Unregister a platform
type UnregisterPlatformInput struct {
	//The ID of the platform to unregister.
	PlatformID string `dest:"uri" dest-name:"platformID"`

	//Whether to cascade-delete all the services and bindings that are related to the platform.
	Cascade string `dest:"querystring" dest-name:"cascade"`
}
type UnregisterPlatformOutput struct {
	Error

	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) UnregisterPlatform(ctx context.Context,
	input *UnregisterPlatformInput) (*UnregisterPlatformOutput, error) {
	req, out := c.unregisterPlatformRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) unregisterPlatformRequest(ctx context.Context,
	input *UnregisterPlatformInput) (*request.Request, *UnregisterPlatformOutput) {
	op := &request.Operation{
		Name: platforms,
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/platforms/{platformID}",
		},
	}

	if input == nil {
		input = &UnregisterPlatformInput{}
	}

	output := &UnregisterPlatformOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PATCH /v1/platforms/{platformID}
// Update a platform
type UpdatePlatformInput struct {
	//The ID of the registered platform to update.
	PlatformID string `dest:"uri" dest-name:"platformID"`

	//The ID of the platform to update.
	//Platform ID is a globally unique identifier (GUID).
	//GUID must be longer than 50 characters and only contains uppercase and lowercase letters, decimal digits, hyphens, periods, and underscores.
	Id string `json:"id,omitempty"`

	//The CLI-friendly name of the platform to update.
	//A CLI-friendly name is a short string that only contains alphanumeric characters, periods, and hyphens.
	//It can't contain white spaces.
	//The name must not exceed 255 characters, but it is recommended to keep it much shorter, for the convenience of using short names in CLI commands.
	Name string `json:"name,omitempty"`

	//The type of the platform.
	//Possible values:
	//Enum:
	//	[ kubernetes ]
	Type string `json:"type,omitempty"`

	//The description of the platform for customer-facing UIs.
	Description string `json:"description,omitempty"`

	//The list of labels to update for the resource.
	Labels []Label `json:"labels,omitempty"`
}
type Label struct {
	//The operation to perform on a label.
	//Possible values:
	//Enum:
	//	[ add, remove ]
	Op string `json:"op,omitempty"`

	//The name of the label.
	Key string `json:"key,omitempty"`

	//The list of values for the label
	Values []string `json:"values,omitempty"`
}
type UpdatePlatformOutput struct {
	Error

	//The ID of the platform to update.
	//Platform ID is a globally unique identifier (GUID).
	//GUID must be longer than 50 characters and only contains uppercase and lowercase letters, decimal digits, hyphens, periods, and underscores.
	Id string `json:"id,omitempty"`

	//Whether the resource is ready for consumption.
	Ready string `json:"ready,omitempty"`

	//The type of the platform.
	//Possible values:
	//Enum:
	//	[ kubernetes ]
	Type string `json:"type,omitempty"`

	//The technical name of the platform.
	Name string `json:"name,omitempty"`

	//The description of the platform for customer-facing UIs.
	Description string `json:"description,omitempty"`

	//The time the platform was created.
	//In ISO 8601 format:
	//	YYYY-MM-DDThh:mm:ssTZD
	CreatedAt string `json:"created_at,omitempty"`

	//The last time the platform was updated.
	//In ISO 8601 format.
	UpdatedAt string `json:"updated_at,omitempty"`

	//The list of labels to update for the resource.
	Labels map[string][]string `json:"labels,omitempty"`

	types.StatusAndBodyFromResponse
}

func (c *ServiceManagementV1) UpdatePlatform(ctx context.Context,
	input *UpdatePlatformInput) (*UpdatePlatformOutput, error) {
	req, out := c.updatePlatformRequest(ctx, input)
	return out, req.Send()
}
func (c *ServiceManagementV1) updatePlatformRequest(ctx context.Context,
	input *UpdatePlatformInput) (*request.Request, *UpdatePlatformOutput) {
	op := &request.Operation{
		Name: platforms,
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/platforms/{platformID}",
		},
	}

	if input == nil {
		input = &UpdatePlatformInput{}
	}

	output := &UpdatePlatformOutput{}
	return c.newRequest(ctx, op, input, output), output
}
