package btpaccounts

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

// GET /accounts/v1/subaccounts
// Get all subaccounts
type GetSubAccountsInput struct {
	//The range of authorizations for which to return information.
	//
	//any: Returns all global accounts for which the user has authorizations on any of the accounts' entities,
	//such as its subaccounts (for example, user is a subaccount admin) or spaces (for example, user is a Cloud Foundry space manager).
	//(empty value): Returns all subaccounts for which the user has explicit authorization on the global account or directory.
	DerivedAuthorizations string `dest:"querystring" dest-name:"derivedAuthorizations"`

	//Returns only the subaccounts in a given directory. Provide the unique ID of the directory.
	DirectoryGuid string `dest:"querystring" dest-name:"directoryGUID"`
}
type GetSubAccountsOutput struct {
	Value []SubAccount `json:"value,omitempty"`

	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) GetSubAccounts(ctx context.Context, input *GetSubAccountsInput) (*GetSubAccountsOutput, error) {
	req, out := c.getSubAccountsRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getSubAccountsRequest(ctx context.Context, input *GetSubAccountsInput) (*request.Request, *GetSubAccountsOutput) {
	op := &request.Operation{
		Name: "Get Sub Accounts",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/subaccounts",
		},
	}

	if input == nil {
		input = &GetSubAccountsInput{}
	}

	output := &GetSubAccountsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST accounts/v1/subaccounts
// Create a subaccount
type CreateSubAccountInput struct {
	//Enables the subaccount to use beta services and applications. Not to be used in a production environment.
	//Cannot be reverted once set. Any use of beta functionality is at the customer's own risk, and SAP shall not be
	//liable for errors or damages caused by the use of beta features.
	BetaEnabled bool `json:"betaEnabled,omitempty"`

	// Additional properties of the subaccount.
	CustomProperties []KeyValue `json:"customProperties,omitempty"`

	// A description of the subaccount for customer-facing UIs.
	Description string `json:"description,omitempty"`

	// The display name of the subaccount for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`

	//The origin of the subaccount creation.
	//
	//REGION_SETUP: Created automatically as part of the region setup.
	//COCKPIT: Created in the cockpit.
	//Enum:
	//	[ REGION_SETUP, COCKPIT, MIGRATED_TO_CP_FOUNDATION_V2, DOMAINDB_SYNC ]
	Origin string `json:"origin,omitempty"`

	//The unique ID subaccount’s parent entity.
	ParentGuid string `json:"parentGUID,omitempty"`

	//The region in which the subaccount was created.
	Region string `json:"region,omitempty"`

	//Additional admins of the subaccount. Do not add yourself as you are assigned as a subaccount admin by default.
	//Enter as a valid JSON array containing the list of admin e-mails (as required by your identity provider).
	//To add admins to Neo subaccounts, use instead the SAP BTP cockpit or the APIs in the SDK for SAP BTP, Neo environment.
	//Example: ["admin1@example.com", "admin2@example.com"]
	SubaccountAdmins []string `json:"subaccountAdmins,omitempty"`

	//The subdomain that becomes part of the path used to access the authorization tenant of the subaccount.
	//Must be unique within the defined region. Use only letters (a-z), digits (0-9), and hyphens (not at start or end).
	//Maximum length is 63 characters. Cannot be changed after the subaccount has been created. Does not apply to Neo subaccounts.
	Subdomain string `json:"subdomain,omitempty"`

	//Whether the subaccount is used for production purposes. This flag can help your cloud operator to take appropriate
	//action when handling incidents that are related to mission-critical accounts in production systems.
	//Do not apply for subaccounts that are used for non-production purposes, such as development, testing, and demos.
	//Applying this setting this does not modify the subaccount.
	//
	//NOT_USED_FOR_PRODUCTION: Subaccount is not used for production purposes.
	//USED_FOR_PRODUCTION: Subaccount is used for production purposes.
	//Enum:
	//	[ USED_FOR_PRODUCTION, NOT_USED_FOR_PRODUCTION ]
	UsedForProduction string `json:"usedForProduction,omitempty"`
}
type CreateSubAccountOutput struct {
	SubAccount

	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) CreateSubAccount(ctx context.Context, input *CreateSubAccountInput) (*CreateSubAccountOutput, error) {
	req, out := c.createSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) createSubAccountRequest(ctx context.Context, input *CreateSubAccountInput) (*request.Request, *CreateSubAccountOutput) {
	op := &request.Operation{
		Name: "Create Sub Account",
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/subaccounts",
		},
	}

	if input == nil {
		input = &CreateSubAccountInput{}
	}

	output := &CreateSubAccountOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /accounts/v1/subaccounts/clone/{sourceSubaccountGUID}
// Clone a Neo subaccount
type CloneSubAccountInput struct {
	SourceSubAccountGuid string `dest:"uri" dest-name:"sourceSubaccountGUID"`

	//Enables the subaccount to use beta services and applications. Not to be used in a production environment.
	//Cannot be reverted once set. Any use of beta functionality is at the customer's own risk, and SAP shall not be
	//liable for errors or damages caused by the use of beta features.
	BetaEnabled bool `json:"betaEnabled,omitempty"`

	//Clone configuration of the subaccount.
	CloneConfigurations []string `json:"cloneConfigurations,omitempty"`

	// Additional properties of the subaccount.
	CustomProperties []KeyValue `json:"customProperties,omitempty"`

	// A description of the subaccount for customer-facing UIs.
	Description string `json:"description,omitempty"`

	// The display name of the subaccount for customer-facing UIs.
	DisplayName string `json:"displayName,omitempty"`

	//The origin of the subaccount creation.
	//
	//REGION_SETUP: Created automatically as part of the region setup.
	//COCKPIT: Created in the cockpit.
	//Enum:
	//	[ REGION_SETUP, COCKPIT, MIGRATED_TO_CP_FOUNDATION_V2, DOMAINDB_SYNC ]
	Origin string `json:"origin,omitempty"`

	//The unique ID subaccount’s parent entity.
	ParentGuid string `json:"parentGUID,omitempty"`

	//The region in which the subaccount was created.
	Region string `json:"region,omitempty"`

	//Additional admins of the subaccount. Do not add yourself as you are assigned as a subaccount admin by default.
	//Enter as a valid JSON array containing the list of admin e-mails (as required by your identity provider).
	//To add admins to Neo subaccounts, use instead the SAP BTP cockpit or the APIs in the SDK for SAP BTP, Neo environment.
	//Example: ["admin1@example.com", "admin2@example.com"]
	SubaccountAdmins []string `json:"subaccountAdmins,omitempty"`

	//The subdomain that becomes part of the path used to access the authorization tenant of the subaccount.
	//Must be unique within the defined region. Use only letters (a-z), digits (0-9), and hyphens (not at start or end).
	//Maximum length is 63 characters. Cannot be changed after the subaccount has been created. Does not apply to Neo subaccounts.
	Subdomain string `json:"subdomain,omitempty"`

	//Whether the subaccount is used for production purposes. This flag can help your cloud operator to take appropriate
	//action when handling incidents that are related to mission-critical accounts in production systems.
	//Do not apply for subaccounts that are used for non-production purposes, such as development, testing, and demos.
	//Applying this setting this does not modify the subaccount.
	//
	//NOT_USED_FOR_PRODUCTION: Subaccount is not used for production purposes.
	//USED_FOR_PRODUCTION: Subaccount is used for production purposes.
	//Enum:
	//	[ USED_FOR_PRODUCTION, NOT_USED_FOR_PRODUCTION ]
	UsedForProduction string `json:"usedForProduction,omitempty"`
}
type CloneSubAccountOutput struct {
	SubAccount

	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) CloneSubAccount(ctx context.Context, input *CloneSubAccountInput) (*CloneSubAccountOutput, error) {
	req, out := c.cloneSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) cloneSubAccountRequest(ctx context.Context, input *CloneSubAccountInput) (*request.Request, *CloneSubAccountOutput) {
	op := &request.Operation{
		Name: "Clone Sub Account",
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/subaccounts/clone/{sourceSubaccountGUID}",
		},
	}

	if input == nil {
		input = &CloneSubAccountInput{}
	}

	output := &CloneSubAccountOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /accounts/v1/subaccounts/{subaccountGUID}
// Get a subaccount
type GetSubAccountInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`

	//The range of authorizations for which to return information.
	//
	//any: Returns all global accounts for which the user has authorizations on any of the accounts' entities, such as
	//	its subaccounts (for example, user is a subaccount admin) or spaces (for example, user is a Cloud Foundry space manager).
	//(empty value): Returns all subaccounts for which the user has explicit authorization on the global account or directory.
	DerivedAuthorizations string `dest:"querystring" dest-name:"derivedAuthorizations"`
}
type GetSubAccountOutput struct {
	SubAccount

	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) GetSubAccount(ctx context.Context, input *GetSubAccountInput) (*GetSubAccountOutput, error) {
	req, out := c.getSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getSubAccountRequest(ctx context.Context, input *GetSubAccountInput) (*request.Request, *GetSubAccountOutput) {
	op := &request.Operation{
		Name: "Get Sub Account",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/subaccounts/{subaccountGUID}",
		},
	}

	if input == nil {
		input = &GetSubAccountInput{}
	}

	output := &GetSubAccountOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /accounts/v1/subaccounts/{subaccountGUID}
// Delete a subaccount
type DeleteSubAccountInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`
}
type DeleteSubAccountOutput struct {
	SubAccount

	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) DeleteSubAccount(ctx context.Context, input *DeleteSubAccountInput) (*DeleteSubAccountOutput, error) {
	req, out := c.deleteSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) deleteSubAccountRequest(ctx context.Context, input *DeleteSubAccountInput) (*request.Request, *DeleteSubAccountOutput) {
	op := &request.Operation{
		Name: "Delete Sub Account",
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/subaccounts/{subaccountGUID}",
		},
	}

	if input == nil {
		input = &DeleteSubAccountInput{}
	}

	output := &DeleteSubAccountOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PATCH /accounts/v1/subaccounts/{subaccountGUID}
// Update a subaccount
type UpdateSubAccountInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`

	//Enables the subaccount to use beta services and applications. Not to be used in a production environment.
	//Cannot be reverted once set. Any use of beta functionality is at the customer's own risk, and SAP shall not be
	//liable for errors or damages caused by the use of beta features.
	BetaEnabled bool `json:"betaEnabled,omitempty"`

	//Additional Properties of the subaccount.
	CustomProperties  []UpdateSubAccountProperties `json:"customProperties,omitempty"`
	Description       string                       `json:"description,omitempty"`
	DisplayName       string                       `json:"displayName,omitempty"`
	UsedForProduction string                       `json:"usedForProduction,omitempty"`
}

//Custom properties as key-value pairs to assign, update, and remove from the subaccount.
type UpdateSubAccountProperties struct {
	KeyValue

	//Whether to delete a property according to the provided key.
	Delete bool `json:"delete,omitempty"`
}
type UpdateSubAccountOutput struct {
	SubAccount

	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) UpdateSubAccount(ctx context.Context, input *UpdateSubAccountInput) (*UpdateSubAccountOutput, error) {
	req, out := c.updateSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) updateSubAccountRequest(ctx context.Context, input *UpdateSubAccountInput) (*request.Request, *UpdateSubAccountOutput) {
	op := &request.Operation{
		Name: "Update Sub Account",
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/subaccounts/{subaccountGUID}",
		},
	}

	if input == nil {
		input = &UpdateSubAccountInput{}
	}

	output := &UpdateSubAccountOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /accounts/v1/subaccounts/{subaccountGUID}/customProperties
// Get custom properties for a subaccount
type GetCustomPropertiesInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`
}
type GetCustomPropertiesOutput struct {
	Value []CustomProperties `json:"value,omitempty"`

	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) GetSubAccountCustomProperties(ctx context.Context, input *GetCustomPropertiesInput) (*GetCustomPropertiesOutput, error) {
	req, out := c.getSubAccountCustomPropertiesRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getSubAccountCustomPropertiesRequest(ctx context.Context, input *GetCustomPropertiesInput) (*request.Request, *GetCustomPropertiesOutput) {
	op := &request.Operation{
		Name: "Get Sub Account Custom Properties",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/subaccounts/{subaccountGUID}/customProperties",
		},
	}

	if input == nil {
		input = &GetCustomPropertiesInput{}
	}

	output := &GetCustomPropertiesOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /accounts/v1/subaccounts/move
// Batch move subaccounts

type MoveManySubAccountsInput struct {
	//Details of which subaccounts to move and where to move them to. All subaccounts must be moved to the same location.
	SubAccountsToMove []MoveSubAccountsRequestPayload `json:"subaccountsToMoveCollection,omitempty"`
}
type MoveSubAccountsRequestPayload struct {
	//The GUID of the current location of the subaccounts. If empty, then GUID of root global account is used.
	SourceGuid string `json:"sourceGuid,omitempty"`

	//GUIDs of the subaccounts to move.
	SubaccountGuids []string `json:"subaccountGuids,omitempty"`

	//The GUID of the new location of the subaccounts. To move to a directory, enter the GUID of the directory.
	//To move out of a directory to the root global account, enter the GUID of the global account.
	TargetGuid string `json:"targetGuid,omitempty"`
}
type MoveManySubAccountsOutput struct {
	SubAccount

	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) MoveManySubAccounts(ctx context.Context, input *MoveManySubAccountsInput) (*MoveManySubAccountsOutput, error) {
	req, out := c.moveManySubAccountsRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) moveManySubAccountsRequest(ctx context.Context, input *MoveManySubAccountsInput) (*request.Request, *MoveManySubAccountsOutput) {
	op := &request.Operation{
		Name: "Move Many Sub Accounts",
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/subaccounts/move",
		},
	}

	if input == nil {
		input = &MoveManySubAccountsInput{}
	}

	output := &MoveManySubAccountsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /accounts/v1/subaccounts/{subaccountGUID}/move
// Move a subaccount
type MoveSubAccountInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`

	//The GUID of the new location of the subaccount. To move to a directory, enter the GUID of the directory.
	//To move out of a directory to the root global account, enter the GUID of the global account.
	TargetAccountGuid string `json:"targetAccountGUID,omitempty"`
}

type MoveSubAccountOutput struct {
	SubAccount

	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) MoveSubAccount(ctx context.Context, input *MoveSubAccountInput) (*MoveSubAccountOutput, error) {
	req, out := c.moveSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) moveSubAccountRequest(ctx context.Context, input *MoveSubAccountInput) (*request.Request, *MoveSubAccountOutput) {
	op := &request.Operation{
		Name: "Move Many Sub Accounts",
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/subaccounts/{subaccountGUID}/move",
		},
	}

	if input == nil {
		input = &MoveSubAccountInput{}
	}

	output := &MoveSubAccountOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /accounts/v1/subaccounts/{subaccountGUID}/serviceManagementBinding
// Get a Service Management binding
type GetServiceManagementBindingInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`
}

type GetServiceManagementBindingOutput struct {
	ServiceManagementBinding

	types.StatusAndBodyFromResponse
}

//OAuth 2.0 Client Credentials Grant Type to obtain an access token to use the Service Management APIs in a subaccount context.
type ServiceManagementBinding struct {
	//A public identifier of the app.
	ClientId string `json:"clientid,omitempty"`

	//Secret known only to the app and the authorization server.
	ClientSecret string `json:"clientsecret,omitempty"`

	//The URL of Service Management APIs to access with the obtained token.
	SMUrl string `json:"sm_url,omitempty"`

	//The URL to authentication server to get a token to authenticate with Service Management using the obtained client ID and secret.
	Url string `json:"url,omitempty"`

	//The name of the xsapp used to get the access token.
	XsAppName string `json:"xsappname,omitempty"`
}

func (c *AccountsV1) GetSubAccountServiceManagementBinding(ctx context.Context,
	input *GetServiceManagementBindingInput) (*GetServiceManagementBindingOutput, error) {
	req, out := c.getSubAccountServiceManagementBindingRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getSubAccountServiceManagementBindingRequest(ctx context.Context,
	input *GetServiceManagementBindingInput) (*request.Request, *GetServiceManagementBindingOutput) {
	op := &request.Operation{
		Name: "Get Sub Account Service Management Bindings",
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/subaccounts/{subaccountGUID}/serviceManagementBinding",
		},
	}

	if input == nil {
		input = &GetServiceManagementBindingInput{}
	}

	output := &GetServiceManagementBindingOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /accounts/v1/subaccounts/{subaccountGUID}/serviceManagementBinding
// Create a Service Management binding
type CreateServiceManagementBindingInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`
}
type CreateServiceManagementBindingOutput struct {
	ServiceManagementBinding

	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) CreateSubAccountServiceManagementBinding(ctx context.Context,
	input *CreateServiceManagementBindingInput) (*CreateServiceManagementBindingOutput, error) {
	req, out := c.createSubAccountServiceManagementBindingRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) createSubAccountServiceManagementBindingRequest(ctx context.Context,
	input *CreateServiceManagementBindingInput) (*request.Request, *CreateServiceManagementBindingOutput) {
	op := &request.Operation{
		Name: "Create Sub Account Service Management Bindings",
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/subaccounts/{subaccountGUID}/serviceManagementBinding",
		},
	}

	if input == nil {
		input = &CreateServiceManagementBindingInput{}
	}

	output := &CreateServiceManagementBindingOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /accounts/v1/subaccounts/{subaccountGUID}/serviceManagementBinding
// Delete a Service Management binding
type DeleteServiceManagementBindingInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`
}

type DeleteServiceManagementBindingOutput struct {
	types.StatusAndBodyFromResponse
}

func (c *AccountsV1) DeleteSubAccountServiceManagementBinding(ctx context.Context,
	input *DeleteServiceManagementBindingInput) (*DeleteServiceManagementBindingOutput, error) {
	req, out := c.deleteSubAccountServiceManagementBindingRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) deleteSubAccountServiceManagementBindingRequest(ctx context.Context,
	input *DeleteServiceManagementBindingInput) (*request.Request, *DeleteServiceManagementBindingOutput) {
	op := &request.Operation{
		Name: "Delete Sub Account Service Management Bindings",
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/subaccounts/{subaccountGUID}/serviceManagementBinding",
		},
	}

	if input == nil {
		input = &DeleteServiceManagementBindingInput{}
	}
	output := &DeleteServiceManagementBindingOutput{}
	return c.newRequest(ctx, op, input, output), output
}
