package btpaccounts

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
)

const subAccounts = "SubAccounts"

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
	DirectoryGUID string `dest:"querystring" dest-name:"directoryGUID"`
}
type GetSubAccountsOutput struct {
	Value []SubAccount `json:"value"`
}

func (c *AccountsV1) GetSubAccounts(ctx context.Context, input *GetSubAccountsInput) (*GetSubAccountsOutput, error) {
	req, out := c.getSubAccountsRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getSubAccountsRequest(ctx context.Context, input *GetSubAccountsInput) (*request.Request, *GetSubAccountsOutput) {
	op := &request.Operation{
		Name: subAccounts,
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
	BetaEnabled bool `json:"betaEnabled"`

	// Additional properties of the subaccount.
	CustomProperties []KeyValue `json:"customProperties"`

	// A description of the subaccount for customer-facing UIs.
	Description string `json:"description"`

	// The display name of the subaccount for customer-facing UIs.
	DisplayName string `json:"displayName"`

	//The origin of the subaccount creation.
	//
	//REGION_SETUP: Created automatically as part of the region setup.
	//COCKPIT: Created in the cockpit.
	//Enum:
	//	[ REGION_SETUP, COCKPIT, MIGRATED_TO_CP_FOUNDATION_V2, DOMAINDB_SYNC ]
	Origin string `json:"origin"`

	//The unique ID subaccount’s parent entity.
	ParentGuid string `json:"parentGUID"`

	//The region in which the subaccount was created.
	Region string `json:"region"`

	//Additional admins of the subaccount. Do not add yourself as you are assigned as a subaccount admin by default.
	//Enter as a valid JSON array containing the list of admin e-mails (as required by your identity provider).
	//To add admins to Neo subaccounts, use instead the SAP BTP cockpit or the APIs in the SDK for SAP BTP, Neo environment.
	//Example: ["admin1@example.com", "admin2@example.com"]
	SubaccountAdmins []string `json:"subaccountAdmins"`

	//The subdomain that becomes part of the path used to access the authorization tenant of the subaccount.
	//Must be unique within the defined region. Use only letters (a-z), digits (0-9), and hyphens (not at start or end).
	//Maximum length is 63 characters. Cannot be changed after the subaccount has been created. Does not apply to Neo subaccounts.
	Subdomain string `json:"subdomain"`

	//Whether the subaccount is used for production purposes. This flag can help your cloud operator to take appropriate
	//action when handling incidents that are related to mission-critical accounts in production systems.
	//Do not apply for subaccounts that are used for non-production purposes, such as development, testing, and demos.
	//Applying this setting this does not modify the subaccount.
	//
	//NOT_USED_FOR_PRODUCTION: Subaccount is not used for production purposes.
	//USED_FOR_PRODUCTION: Subaccount is used for production purposes.
	//Enum:
	//	[ USED_FOR_PRODUCTION, NOT_USED_FOR_PRODUCTION ]
	UsedForProduction string `json:"usedForProduction"`
}

func (c *AccountsV1) CreateSubAccount(ctx context.Context, input *CreateSubAccountInput) (*SubAccount, error) {
	req, out := c.createSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) createSubAccountRequest(ctx context.Context, input *CreateSubAccountInput) (*request.Request, *SubAccount) {
	op := &request.Operation{
		Name: subAccounts,
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/subaccounts",
		},
	}

	if input == nil {
		input = &CreateSubAccountInput{}
	}

	output := &SubAccount{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /accounts/v1/subaccounts/clone/{sourceSubaccountGUID}
// Clone a Neo subaccount
type CloneSubAccountInput struct {
	SourceSubAccountGuid string `dest:"uri" dest-name:"sourceSubaccountGUID"`

	//Enables the subaccount to use beta services and applications. Not to be used in a production environment.
	//Cannot be reverted once set. Any use of beta functionality is at the customer's own risk, and SAP shall not be
	//liable for errors or damages caused by the use of beta features.
	BetaEnabled bool `json:"betaEnabled"`

	//Clone configuration of the subaccount.
	CloneConfigurations []string `json:"cloneConfigurations"`

	// Additional properties of the subaccount.
	CustomProperties []KeyValue `json:"customProperties"`

	// A description of the subaccount for customer-facing UIs.
	Description string `json:"description"`

	// The display name of the subaccount for customer-facing UIs.
	DisplayName string `json:"displayName"`

	//The origin of the subaccount creation.
	//
	//REGION_SETUP: Created automatically as part of the region setup.
	//COCKPIT: Created in the cockpit.
	//Enum:
	//	[ REGION_SETUP, COCKPIT, MIGRATED_TO_CP_FOUNDATION_V2, DOMAINDB_SYNC ]
	Origin string `json:"origin"`

	//The unique ID subaccount’s parent entity.
	ParentGuid string `json:"parentGUID"`

	//The region in which the subaccount was created.
	Region string `json:"region"`

	//Additional admins of the subaccount. Do not add yourself as you are assigned as a subaccount admin by default.
	//Enter as a valid JSON array containing the list of admin e-mails (as required by your identity provider).
	//To add admins to Neo subaccounts, use instead the SAP BTP cockpit or the APIs in the SDK for SAP BTP, Neo environment.
	//Example: ["admin1@example.com", "admin2@example.com"]
	SubaccountAdmins []string `json:"subaccountAdmins"`

	//The subdomain that becomes part of the path used to access the authorization tenant of the subaccount.
	//Must be unique within the defined region. Use only letters (a-z), digits (0-9), and hyphens (not at start or end).
	//Maximum length is 63 characters. Cannot be changed after the subaccount has been created. Does not apply to Neo subaccounts.
	Subdomain string `json:"subdomain"`

	//Whether the subaccount is used for production purposes. This flag can help your cloud operator to take appropriate
	//action when handling incidents that are related to mission-critical accounts in production systems.
	//Do not apply for subaccounts that are used for non-production purposes, such as development, testing, and demos.
	//Applying this setting this does not modify the subaccount.
	//
	//NOT_USED_FOR_PRODUCTION: Subaccount is not used for production purposes.
	//USED_FOR_PRODUCTION: Subaccount is used for production purposes.
	//Enum:
	//	[ USED_FOR_PRODUCTION, NOT_USED_FOR_PRODUCTION ]
	UsedForProduction string `json:"usedForProduction"`
}

func (c *AccountsV1) CloneSubAccount(ctx context.Context, input *CloneSubAccountInput) (*SubAccount, error) {
	req, out := c.cloneSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) cloneSubAccountRequest(ctx context.Context, input *CloneSubAccountInput) (*request.Request, *SubAccount) {
	op := &request.Operation{
		Name: subAccounts,
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/subaccounts/clone/{sourceSubaccountGUID}",
		},
	}

	if input == nil {
		input = &CloneSubAccountInput{}
	}

	output := &SubAccount{}
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

func (c *AccountsV1) GetSubAccount(ctx context.Context, input *GetSubAccountInput) (*SubAccount, error) {
	req, out := c.getSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getSubAccountRequest(ctx context.Context, input *GetSubAccountInput) (*request.Request, *SubAccount) {
	op := &request.Operation{
		Name: subAccounts,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/subaccounts/{subaccountGUID}",
		},
	}

	if input == nil {
		input = &GetSubAccountInput{}
	}

	output := &SubAccount{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /accounts/v1/subaccounts/{subaccountGUID}
// Delete a subaccount
type DeleteSubAccountInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`
}

func (c *AccountsV1) DeleteSubAccount(ctx context.Context, input *DeleteSubAccountInput) (*SubAccount, error) {
	req, out := c.deleteSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) deleteSubAccountRequest(ctx context.Context, input *DeleteSubAccountInput) (*request.Request, *SubAccount) {
	op := &request.Operation{
		Name: subAccounts,
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/subaccounts/{subaccountGUID}",
		},
	}

	if input == nil {
		input = &DeleteSubAccountInput{}
	}

	output := &SubAccount{}
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
	BetaEnabled bool `json:"betaEnabled"`

	//Additional Properties of the subaccount.
	CustomProperties  []UpdateSubAccountProperties `json:"customProperties"`
	Description       string                       `json:"description"`
	DisplayName       string                       `json:"displayName"`
	UsedForProduction string                       `json:"usedForProduction"`
}

//Custom properties as key-value pairs to assign, update, and remove from the subaccount.
type UpdateSubAccountProperties struct {
	KeyValue

	//Whether to delete a property according to the provided key.
	Delete bool `json:"delete"`
}

func (c *AccountsV1) UpdateSubAccount(ctx context.Context, input *UpdateSubAccountInput) (*SubAccount, error) {
	req, out := c.updateSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) updateSubAccountRequest(ctx context.Context, input *UpdateSubAccountInput) (*request.Request, *SubAccount) {
	op := &request.Operation{
		Name: subAccounts,
		Http: request.HTTP{
			Method: request.DELETE,
			Path:   "/subaccounts/{subaccountGUID}",
		},
	}

	if input == nil {
		input = &UpdateSubAccountInput{}
	}

	output := &SubAccount{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /accounts/v1/subaccounts/{subaccountGUID}/customProperties
// Get custom properties for a subaccount
type GetCustomPropertiesInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`
}
type GetCustomPropertiesOutput struct {
	Value []CustomProperties `json:"value"`
}

func (c *AccountsV1) GetSubAccountCustomProperties(ctx context.Context, input *GetCustomPropertiesInput) (*GetCustomPropertiesOutput, error) {
	req, out := c.getSubAccountCustomPropertiesRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getSubAccountCustomPropertiesRequest(ctx context.Context, input *GetCustomPropertiesInput) (*request.Request, *GetCustomPropertiesOutput) {
	op := &request.Operation{
		Name: subAccounts,
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
	SubAccountsToMove []MoveSubAccountsRequestPayload `json:"subaccountsToMoveCollection"`
}
type MoveSubAccountsRequestPayload struct {
	//The GUID of the current location of the subaccounts. If empty, then GUID of root global account is used.
	SourceGuid string `json:"sourceGuid"`

	//GUIDs of the subaccounts to move.
	SubaccountGuids []string `json:"subaccountGuids"`

	//The GUID of the new location of the subaccounts. To move to a directory, enter the GUID of the directory.
	//To move out of a directory to the root global account, enter the GUID of the global account.
	TargetGuid string `json:"targetGuid"`
}

func (c *AccountsV1) MoveManySubAccounts(ctx context.Context, input *MoveManySubAccountsInput) (*SubAccount, error) {
	req, out := c.moveManySubAccountsRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) moveManySubAccountsRequest(ctx context.Context, input *MoveManySubAccountsInput) (*request.Request, *SubAccount) {
	op := &request.Operation{
		Name: subAccounts,
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/subaccounts/move",
		},
	}

	if input == nil {
		input = &MoveManySubAccountsInput{}
	}

	output := &SubAccount{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /accounts/v1/subaccounts/{subaccountGUID}/move
// Move a subaccount
type MoveSubAccountInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`

	//The GUID of the new location of the subaccount. To move to a directory, enter the GUID of the directory.
	//To move out of a directory to the root global account, enter the GUID of the global account.
	TargetAccountGuid string `json:"targetAccountGUID"`
}

func (c *AccountsV1) MoveSubAccount(ctx context.Context, input *MoveSubAccountInput) (*SubAccount, error) {
	req, out := c.moveSubAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) moveSubAccountRequest(ctx context.Context, input *MoveSubAccountInput) (*request.Request, *SubAccount) {
	op := &request.Operation{
		Name: subAccounts,
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/subaccounts/{subaccountGUID}/move",
		},
	}

	if input == nil {
		input = &MoveSubAccountInput{}
	}

	output := &SubAccount{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /accounts/v1/subaccounts/{subaccountGUID}/serviceManagementBinding
// Get a Service Management binding
type GetServiceManagementBindingInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`
}

//OAuth 2.0 Client Credentials Grant Type to obtain an access token to use the Service Management APIs in a subaccount context.
type ServiceManagementBindingOutput struct {
	//A public identifier of the app.
	ClientId string `json:"clientid"`

	//Secret known only to the app and the authorization server.
	ClientSecret string `json:"clientsecret"`

	//The URL of Service Management APIs to access with the obtained token.
	SMUrl string `json:"sm_url"`

	//The URL to authentication server to get a token to authenticate with Service Management using the obtained client ID and secret.
	Url string `json:"url"`

	//The name of the xsapp used to get the access token.
	XsAppName string `json:"xsappname"`
}

func (c *AccountsV1) GetSubAccountServiceManagementBinding(ctx context.Context, input *GetServiceManagementBindingInput) (*ServiceManagementBindingOutput, error) {
	req, out := c.getSubAccountServiceManagementBindingRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getSubAccountServiceManagementBindingRequest(ctx context.Context, input *GetServiceManagementBindingInput) (*request.Request, *ServiceManagementBindingOutput) {
	op := &request.Operation{
		Name: subAccounts,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/subaccounts/{subaccountGUID}/serviceManagementBinding",
		},
	}

	if input == nil {
		input = &GetServiceManagementBindingInput{}
	}

	output := &ServiceManagementBindingOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// POST /accounts/v1/subaccounts/{subaccountGUID}/serviceManagementBinding
// Create a Service Management binding
type CreateServiceManagementBindingInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`
}

func (c *AccountsV1) CreateSubAccountServiceManagementBinding(ctx context.Context, input *CreateServiceManagementBindingInput) (*ServiceManagementBindingOutput, error) {
	req, out := c.createSubAccountServiceManagementBindingRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) createSubAccountServiceManagementBindingRequest(ctx context.Context, input *CreateServiceManagementBindingInput) (*request.Request, *ServiceManagementBindingOutput) {
	op := &request.Operation{
		Name: subAccounts,
		Http: request.HTTP{
			Method: request.POST,
			Path:   "/subaccounts/{subaccountGUID}/serviceManagementBinding",
		},
	}

	if input == nil {
		input = &CreateServiceManagementBindingInput{}
	}

	output := &ServiceManagementBindingOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// DELETE /accounts/v1/subaccounts/{subaccountGUID}/serviceManagementBinding
// Delete a Service Management binding
type DeleteServiceManagementBindingInput struct {
	//The GUID of the subaccount for which to get details.
	SubAccountGuid string `dest:"uri" dest-name:"subaccountGUID"`
}

type DeleteServiceManagementBindingOutput struct {
}

func (c *AccountsV1) DeleteSubAccountServiceManagementBinding(ctx context.Context,
	input *DeleteServiceManagementBindingInput) (*DeleteServiceManagementBindingOutput, error) {
	req, out := c.deleteSubAccountServiceManagementBindingRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) deleteSubAccountServiceManagementBindingRequest(ctx context.Context,
	input *DeleteServiceManagementBindingInput) (*request.Request, *DeleteServiceManagementBindingOutput) {
	op := &request.Operation{
		Name: subAccounts,
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
