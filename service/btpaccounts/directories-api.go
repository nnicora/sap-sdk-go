package btpaccounts

import (
	"context"
	"github.com/nnicora/sap-sdk-go/internal/times"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

type Directory struct {
	// The response object containing information about the directories.
	Children []map[string]interface{} `json:"children,omitempty"`

	// The status of the customer contract and its associated root global account.
	//
	//ACTIVE: The customer contract and its associated global account is currently active.
	//PENDING_TERMINATION: A termination process has been triggered for a customer contract
	//	(the customer contract has expired, or a customer has given notification that they wish to
	//	terminate their contract), and the global account is currently in the validation period.
	//	The customer can still access their global account until the end of the validation period.
	//SUSPENDED: For enterprise accounts, specifies that the customer's global account is currently
	//	in the grace period of the termination process. Access to the global account by the customer is blocked.
	//	No data is deleted until the deletion date is reached at the end of the grace period.
	//	For trial accounts, specifies that the account is suspended, and the account owner
	//	has not yet extended the trial period.
	//Enum:
	//	[ ACTIVE, PENDING_TERMINATION, SUSPENDED ]
	ContractStatus string `json:"contractStatus,omitempty"`

	// Details of the user that created the directory.
	CreatedBy string `json:"createdBy,omitempty"`

	// The date the directory was created. Dates and times are in UTC format.
	CreatedDate times.JavaTime `json:"createdDate,omitempty"`

	// Custom properties assigned to the directory as key-value pairs.
	CustomProperties []CustomProperties `json:"customProperties,omitempty"`

	// A description of the directory.
	Description string `json:"description,omitempty"`

	// The features that are enabled for the directory. Valid values:
	//
	//DEFAULT: All directories have the following basic features enabled:
	//	(1) Group and filter subaccounts for reports and filters,
	//	(2) monitor usage and costs on a directory level (costs only available for contracts that use the
	//		consumption-based commercial model), and
	//	(3) set custom properties and tags to the directory for identification and reporting purposes.
	//ENTITLEMENTS: (Optional) Enables the assignment of a quota for services and applications
	//	to the directory from the global account quota for distribution to the directory's subaccounts.
	//AUTHORIZATIONS: (Optional) Enables a custom identity provider and/or authorization management for the directory.
	//	For example, it allows certain users to manage directory entitlements. You can only use this feature in
	//	combination with the ENTITLEMENTS feature.
	//
	//Examples:
	//	[DEFAULT]
	//	[DEFAULT,ENTITLEMENTS]
	//	[DEFAULT,ENTITLEMENTS,AUTHORIZATIONS]
	//Enum:
	//	[ DEFAULT, ENTITLEMENTS, AUTHORIZATIONS ]
	DirectoryFeatures []string `json:"directoryFeatures,omitempty"`

	// The display name of the directory.
	DisplayName string `json:"displayName,omitempty"`

	// The current state of the directory.
	//
	//STARTED: CRUD operation on an entity has started.
	//CREATING: Creating entity operation is in progress.
	//UPDATING: Updating entity operation is in progress.
	//MOVING: Moving entity operation is in progress.
	//PROCESSING: A series of operations related to the entity is in progress.
	//DELETING: Deleting entity operation is in progress.
	//OK: The CRUD operation or series of operations completed successfully.
	//PENDING REVIEW: The processing operation has been stopped for reviewing and can be restarted by the operator.
	//CANCELLED: The operation or processing was canceled by the operator.
	//CREATION_FAILED: The creation operation failed, and the entity was not created or was created but cannot be used.
	//UPDATE_FAILED: The update operation failed, and the entity was not updated.
	//PROCESSING_FAILED: The processing operations failed.
	//DELETION_FAILED: The delete operation failed, and the entity was not deleted.
	//MOVE_FAILED: Entity could not be moved to a different location.
	//MIGRATING: Migrating entity from NEO to CF.
	//Enum:
	//	[ STARTED, CREATING, UPDATING, MOVING, PROCESSING, DELETING, OK, PENDING_REVIEW, CANCELED, CREATION_FAILED,
	//	UPDATE_FAILED, UPDATE_ACCOUNT_TYPE_FAILED, UPDATE_DIRECTORY_TYPE_FAILED, PROCESSING_FAILED, DELETION_FAILED,
	//	MOVE_FAILED, MIGRATING, MIGRATION_FAILED, ROLLBACK_MIGRATION_PROCESSING, MIGRATED ]
	EntityState string `json:"entityState,omitempty"`

	// The unique ID of the directory.
	Guid       string     `json:"guid,omitempty"`
	LegalLinks LegalLinks `json:"legalLinks,omitempty"`

	// The date the directory was last modified. Dates and times are in UTC format.
	ModifiedDate times.JavaTime `json:"modifiedDate,omitempty"`

	// The GUID of the directory's parent entity. Typically this is the global account.
	ParentGuid string `json:"parentGuid,omitempty"`

	// Information about the state.
	StateMessage string `json:"stateMessage,omitempty"`

	// The subaccounts contained in the directory.
	SubAccounts []SubAccount `json:"subaccounts,omitempty"`

	// Relevant only for directories that are enabled to manage their authorizations. The subdomain that becomes part
	// of the path used to access the authorization tenant of the directory. Unique within the defined region.
	Subdomain string `json:"subdomain,omitempty"`
}

// POST /accounts/v1/directories
// Create a directory
type CreateDirectoryInput struct {
	// Additional properties of the directory.
	CustomProperties []CustomProperties `json:"customProperties,omitempty"`

	// A description of the directory.
	Description string `json:"description,omitempty"`

	// Additional admins of the directory. Do not add yourself as you are assigned as a directory admin by default.
	// Use only with directories that are configured to manage their authorizations.
	//Example: ["admin1@example.com", "admin2@example.com"]
	DirectoryAdmins []string `json:"directoryAdmins,omitempty"`

	// The features to enable for the directory.
	//
	//DEFAULT: (Mandatory) All directories provide the following basic features:
	//	(1) Group and filter subaccounts for reports and filters,
	//	(2) monitor usage and costs on a directory level (costs only available for contracts that use
	//		the consumption-based commercial model), and
	//	(3) set custom properties and tags to the directory for identification and reporting purposes.
	//ENTITLEMENTS: (Optional) Enables the assignment of a quota for services and applications to the
	//	directory from the global account quota for distribution to the directory's subaccounts.
	//AUTHORIZATIONS: (Optional) Enables a custom identity provider and/or authorization management for the directory.
	//	For example, it allows certain users to manage directory entitlements. NOTE: You can only use this feature in combination with the ENTITLEMENTS feature.
	//IMPORTANT: Once a feature has been enabled for a directory, you cannot disable it. If you are not sure which
	//	features to enable, we recommend that you set only the default features, and then add features later on as they are needed.
	//
	//Examples:
	//	[DEFAULT]
	//	[DEFAULT,ENTITLEMENTS]
	//	[DEFAULT,ENTITLEMENTS,AUTHORIZATIONS]
	//Enum:
	//	[ DEFAULT, ENTITLEMENTS, AUTHORIZATIONS ]
	DirectoryFeatures []string `json:"directoryFeatures,omitempty"`

	// The display name of the directory.
	DisplayName string `json:"displayName,omitempty"`

	//Relevant only for directories that are enabled to manage their authorizations. The subdomain that becomes part
	//of the path used to access the authorization tenant of the directory. Must be unique in the defined region.
	//Use only letters (a-z), digits (0-9), and hyphens (not at start or end). Maximum length is 63 characters.
	//Cannot be changed after the directory has been created.
	Subdomain string `json:"subdomain,omitempty"`
}
type CreateDirectoryOutput struct {
	Directory

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) CreateDirectory(ctx context.Context,
	input *CreateDirectoryInput) (*CreateDirectoryOutput, error) {
	req, out := c.createDirectoryRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) createDirectoryRequest(ctx context.Context,
	input *CreateDirectoryInput) (*request.Request, *CreateDirectoryOutput) {
	op := &request.Operation{
		Name: "Create Account Directory",
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/directories",
		},
	}

	if input == nil {
		input = &CreateDirectoryInput{}
	}

	output := &CreateDirectoryOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /accounts/v1/directories/{directoryGUID}
// Get a directory
type GetDirectoryInput struct {
	//The GUID of the directory for which to get details.
	DirectoryGuid string `dest:"uri" dest-name:"directoryGUID"`

	// The range of authorizations for which to return information.
	//
	//any: Returns a directory for which the user has authorizations on any of the subaccounts
	//(for example, user is a subaccount admin) or Cloud Foundry roles (for example, user is a Cloud Foundry space manager).
	//(empty value): Returns a directory for which the user has explicit authorization.
	DerivedAuthorizations string `dest:"querystring" dest-name:"derivedAuthorizations"`

	//Whether to get the contents of the directory, for example the subaccounts it contains.
	Expand bool `dest:"querystring" dest-name:"expand"`
}
type GetDirectoryOutput struct {
	Directory

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) GetDirectory(ctx context.Context,
	input *GetDirectoryInput) (*GetDirectoryOutput, error) {
	req, out := c.getDirectoryRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getDirectoryRequest(ctx context.Context,
	input *GetDirectoryInput) (*request.Request, *GetDirectoryOutput) {
	op := &request.Operation{
		Name: "Get Account Directory",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/directories/{directoryGUID}",
		},
	}

	if input == nil {
		input = &GetDirectoryInput{}
	}

	output := &GetDirectoryOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /accounts/v1/directories/{directoryGUID}
// Delete a directory
type DeleteDirectoryInput struct {
	//The GUID of the directory to update.
	DirectoryGuid string `dest:"uri" dest-name:"directoryGUID"`

	//Whether to delete the directory even if it contains data. If not set to true,
	//the request fails when the directory contains data.
	ForceDelete bool `dest:"querystring" dest-name:"forceDelete"`
}
type DeleteDirectoryOutput struct {
	Directory

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) DeleteDirectory(ctx context.Context,
	input *DeleteDirectoryInput) (*DeleteDirectoryOutput, error) {
	req, out := c.deleteDirectoryRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) deleteDirectoryRequest(ctx context.Context,
	input *DeleteDirectoryInput) (*request.Request, *DeleteDirectoryOutput) {
	op := &request.Operation{
		Name: "Delete Account Directory",
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/directories/{directoryGUID}",
		},
	}

	if input == nil {
		input = &DeleteDirectoryInput{}
	}

	output := &DeleteDirectoryOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PATCH /accounts/v1/directories/{directoryGUID}
// Update a directory
type UpdateDirectoryInput struct {
	//The GUID of the directory to update.
	DirectoryGuid string `dest:"uri" dest-name:"directoryGUID"`

	//Additional Properties of the directory.
	CustomProperties []CustomProperties `json:"customProperties,omitempty"`

	//The description of the directory for the customer-facing UIs.
	Description string `json:"description,omitempty"`

	//The new descriptive name of the directory.
	DisplayName string `json:"displayName,omitempty"`
}
type UpdateDirectoryOutput struct {
	Directory

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) UpdateDirectory(ctx context.Context,
	input *UpdateDirectoryInput) (*UpdateDirectoryOutput, error) {
	req, out := c.updateDirectoryRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) updateDirectoryRequest(ctx context.Context,
	input *UpdateDirectoryInput) (*request.Request, *UpdateDirectoryOutput) {
	op := &request.Operation{
		Name: "Update Account Directory",
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/directories/{directoryGUID}",
		},
	}

	if input == nil {
		input = &UpdateDirectoryInput{}
	}

	output := &UpdateDirectoryOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PATCH /accounts/v1/directories/{directoryGUID}/changeDirectoryFeatures
// Add features to a directory
type UpdateDirectoryFeaturesInput struct {
	//The GUID of the directory to update.
	DirectoryGuid string `dest:"uri" dest-name:"directoryGUID"`

	//Additional admins of the directory. Do not add yourself as you are assigned as a directory admin by default.
	//Use only with directories that are configured to manage their authorizations.
	//Example: ["admin1@example.com", "admin2@example.com"]
	DirectoryAdmins []string `json:"directoryAdmins,omitempty"`

	//The features to enable for the directory.
	//
	//ENTITLEMENTS: (Optional) Enables the assignment of a quota for services and applications to the directory
	//	from the global account quota for distribution to the directory's subaccounts.
	//AUTHORIZATIONS: (Optional) Enables a custom identity provider and/or authorization management for the directory.
	//	For example, it allows certain users to manage directory entitlements. NOTE: You can only use this feature in
	//	combination with the ENTITLEMENTS feature.
	//IMPORTANT: Once a feature has been enabled for a directory, you cannot disable it. If you are not sure which
	//	features to enable, we recommend that you set only the default features, and then add features later on as they are needed.
	//
	//Examples:
	//	[DEFAULT]
	//	[DEFAULT,ENTITLEMENTS]
	//	[DEFAULT,ENTITLEMENTS,AUTHORIZATIONS]
	//Enum:
	//	[ ENTITLEMENTS, AUTHORIZATIONS ]
	DirectoryFeatures []string `json:"directoryFeatures,omitempty"`

	//Relevant only for directories that are enabled to manage their authorizations. The subdomain that becomes
	//part of the path used to access the authorization tenant of the directory. Must be unique within the defined region.
	//Use only letters (a-z), digits (0-9), and hyphens (not at start or end). Maximum length is 63 characters.
	//Cannot be changed after the directory has been created.
	Subdomain string `json:"subdomain,omitempty"`
}
type UpdateDirectoryFeaturesOutput struct {
	Directory

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) UpdateDirectoryFeatures(ctx context.Context,
	input *UpdateDirectoryFeaturesInput) (*UpdateDirectoryFeaturesOutput, error) {
	req, out := c.updateDirectoryFeaturesRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) updateDirectoryFeaturesRequest(ctx context.Context,
	input *UpdateDirectoryFeaturesInput) (*request.Request, *UpdateDirectoryFeaturesOutput) {
	op := &request.Operation{
		Name: "Add Feature To Account Directory",
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/directories/{directoryGUID}/changeDirectoryFeatures",
		},
	}

	if input == nil {
		input = &UpdateDirectoryFeaturesInput{}
	}

	output := &UpdateDirectoryFeaturesOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /accounts/v1/directories/{directoryGUID}/customProperties
// Get directory custom properties
type GetDirectoryCustomPropertiesInput struct {
	//The GUID of the directory to update.
	DirectoryGuid string `dest:"uri" dest-name:"directoryGUID"`
}
type GetDirectoryCustomPropertiesOutput struct {
	Value []CustomProperties `json:"value,omitempty"`

	Error *types.Error `json:"error,omitempty"`
	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) GetDirectorCustomProperties(ctx context.Context,
	input *GetDirectoryCustomPropertiesInput) (*GetDirectoryCustomPropertiesOutput, error) {
	req, out := c.getDirectorCustomPropertiesRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getDirectorCustomPropertiesRequest(ctx context.Context,
	input *GetDirectoryCustomPropertiesInput) (*request.Request, *GetDirectoryCustomPropertiesOutput) {
	op := &request.Operation{
		Name: "Get Account Directory Custom Directory",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/directories/{directoryGUID}/changeDirectoryFeatures",
		},
	}

	if input == nil {
		input = &GetDirectoryCustomPropertiesInput{}
	}

	output := &GetDirectoryCustomPropertiesOutput{}
	return c.newRequest(ctx, op, input, output), output
}
