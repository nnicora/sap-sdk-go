package btpaccounts

import (
	"context"
	"github.com/nnicora/sap-sdk-go/internal/times"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

const globalAccounts = "GlobalAccounts"

// GET /accounts/v1/globalAccount
// Update a global account
type GetGlobalAccountInput struct {
	//The range of authorizations for which to return information.
	//
	//any: Returns a global account for which the user has authorizations on any of the accounts'
	//entities, such as its directories (for example, directory admin), subaccounts (for example, user is a subaccount admin)
	//or Cloud Foundry roles (for example, user is a Cloud Foundry space manager).
	//(empty value): Returns a global account for which the user has explicit authorization
	DerivedAuthorizations string `dest:"querystring" dest-name:"derivedAuthorizations"`

	//If true, returns the structure of the global account including all its children, such as subaccounts and directories,
	//in the account model. The structure content may vary from user to user and depends on users’ authorizations.
	Expand bool `dest:"querystring" dest-name:"expand"`
}
type GlobalAccountOutput struct {
	//The list of directories associated with the specified global account.
	Children []Directory `json:"children"`

	//The type of the commercial contract that was signed.
	CommercialModel string `json:"commercialModel"`

	//Whether the customer of the global account pays only for services that they actually use (consumption-based)
	//or pay for subscribed services at a fixed cost irrespective of consumption (subscription-based).
	//
	//TRUE: Consumption-based commercial model.
	//FALSE: Subscription-based commercial model.
	ConsumptionBased bool `json:"consumptionBased"`

	//The status of the customer contract and its associated root global account.
	//
	//ACTIVE: The customer contract and its associated global account is currently active.
	//PENDING_TERMINATION: A termination process has been triggered for a customer contract (the customer contract
	//	has expired, or a customer has given notification that they wish to terminate their contract), and the
	//	global account is currently in the validation period. The customer can still access their global account
	//	until the end of the validation period.
	//SUSPENDED: For enterprise accounts, specifies that the customer's global account is currently in the grace period
	//	of the termination process. Access to the global account by the customer is blocked. No data is deleted until
	//	the deletion date is reached at the end of the grace period. For trial accounts, specifies that the account
	//	is suspended, and the account owner has not yet extended the trial period.
	//Enum:
	//	[ ACTIVE, PENDING_TERMINATION, SUSPENDED ]
	ContractStatus string `json:"contractStatus"`

	//For internal accounts, the cost center that is associated with the global account owner.
	//A cost center represents a set of users belonging to the same business unit and is charged for
	//the creation and usage of the global account.
	CostCenter string `json:"costCenter"`

	//The date the global account was created. Dates and times are in UTC format.
	CreatedDate times.JavaTime `json:"createdDate"`

	//The ID of the customer as registered in the CRM system.
	CrmCustomerId string `json:"crmCustomerId"`

	//The ID of the customer tenant as registered in the CRM system.
	CrmTenantId string `json:"crmTenantId"`

	//Contains information about the additional properties related to a specified global account.
	CustomProperties []CustomProperties `json:"customProperties"`

	//A description of the global account.
	Description string `json:"description"`

	//The display name of the global account.
	DisplayName string `json:"displayName"`

	//The current state of the global account.
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
	EntityState string `json:"entityState"`

	//The planned date that the global account expires. This is the same date as the Contract End Date,
	//unless a manual adjustment has been made to the actual expiration date of the global account.
	//Typically, this property is automatically populated only when a formal termination order is received from the CRM system.
	//From a customer perspective, this date marks the start of the grace period, which is typically 30 days before the actual deletion of the account.
	ExpiryDate times.JavaTime `json:"expiryDate"`

	//The geographic locations from where the global account can be accessed.
	//
	//STANDARD: The global account can be accessed from any geographic location.
	//EU_ACCESS: The global account can be accessed only within locations in the EU.
	//Enum:
	//	[ STANDARD, EU_ACCESS ]
	GeoAccess string `json:"geoAccess"`

	//The unique ID of the global account.
	Guid string `json:"guid"`

	// Legal Description
	LegalLinks LegalLinks `json:"legalLinks"`

	//The type of license for the global account. The license type affects the scope of functions of the account.
	//
	//DEVELOPER: For internal developer global accounts on Staging or Canary landscapes.
	//CUSTOMER: For customer global accounts.
	//PARTNER: For partner global accounts.
	//INTERNAL_DEV: For internal global accounts on the Dev landscape.
	//INTERNAL_PROD: For internal global accounts on the Live landscape.
	//TRIAL: For customer trial accounts.
	//Enum:
	//	[ DEVELOPER, CUSTOMER, PARTNER, INTERNAL_DEV, INTERNAL_PROD, SYSTEM, TRIAL, SAPDEV, SAPPROD ]
	LicenseType string `json:"licenseType"`

	//The date the global account was last modified. Dates and times are in UTC format.
	ModifiedDate times.JavaTime `json:"modifiedDate"`

	//The origin of the account.
	//
	//ORDER: Created by the Order Processing API or Submit Order wizard.
	//OPERATOR: Created by the Global Account wizard.
	//REGION_SETUP: Created automatically as part of the region setup.
	//Enum:
	//	[ ORDER, OPERATOR, REGION_SETUP, MIGRATED_TO_CP_FOUNDATION_V2 ]
	Origin string `json:"origin"`

	//The GUID of the global account's parent entity. Typically this is the global account.
	ParentGuid string `json:"parentGuid"`

	//The Type of the global account's parent entity.
	//
	//Enum:
	//	[ ROOT, GLOBAL_ACCOUNT, PROJECT, GROUP, FOLDER ]
	ParentType string `json:"parentType"`

	// The date that an expired contract was renewed. Dates and times are in UTC format.
	RenewalDate times.JavaTime `json:"renewalDate"`

	// For internal accounts, the service for which the global account was created.
	ServiceId string `json:"serviceId"`

	//Information about the state.
	StateMessage string `json:"stateMessage"`

	//The subaccounts contained in the global account.
	Subaccounts []SubAccount `json:"subaccounts"`

	//Relevant only for entities that require authorization (e.g. global account). The subdomain that becomes part of
	//the path used to access the authorization tenant of the global account. Unique within the defined region.
	Subdomain string `json:"subdomain"`

	//Specifies the current stage of the termination notifications sequence.
	//
	//PENDING_FIRST_NOTIFICATION: A notification has not yet been sent to the global account owner informing them of
	//	the expired contract or termination request.
	//FIRST_NOTIFICATION_PROCESSED: A first notification has been sent to the global account owner informing them of
	//	the expired contract, and the termination date when the global account will be closed.
	//SECOND_NOTIFICATION_PROCESSED: A follow-up notification has been sent to the global account owner.
	//Your mail server must be configured so that termination notifications can be sent by the Core Commercialization Foundation service.
	//
	//Enum:
	//	[ PENDING_FIRST_NOTIFICATION, FIRST_NOTIFICATION_PROCESSED, SECOND_NOTIFICATION_PROCESSED ]
	TerminationNotificationStatus string `json:"terminationNotificationStatus"`

	//For internal accounts, the intended purpose of the global account. Possible purposes:
	//
	//Development: For development of a service.
	//Testing: For testing development.
	//Demo: For creating demos.
	//Production: For delivering a service in a production landscape.
	UseFor string `json:"useFor"`

	//A unique ID to track this event.
	XCorrelationId string `src:"header" src-name:"x-correlationid"`

	types.StatusAndBodyFromResponse
}
type LegalLinks struct {
	Privacy string `json:"privacy"`
}

func (c *AccountsV1) GetGlobalAccount(ctx context.Context, input *GetGlobalAccountInput) (*GlobalAccountOutput, error) {
	req, out := c.getGlobalAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) getGlobalAccountRequest(ctx context.Context, input *GetGlobalAccountInput) (*request.Request, *GlobalAccountOutput) {
	op := &request.Operation{
		Name: globalAccounts,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/globalAccount",
		},
	}

	if input == nil {
		input = &GetGlobalAccountInput{}
	}

	output := &GlobalAccountOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// PATCH /accounts/v1/globalAccount
// Update a global account
type UpdateGlobalAccountInput struct {
	//A new display of the global account.
	Description string `json:"description"`

	//The new descriptive name of the global account.
	DisplayName string `json:"displayName"`
}

func (c *AccountsV1) UpdateGlobalAccount(ctx context.Context, input *UpdateGlobalAccountInput) (*GlobalAccountOutput, error) {
	req, out := c.updateGlobalAccountRequest(ctx, input)
	return out, req.Send()
}
func (c *AccountsV1) updateGlobalAccountRequest(ctx context.Context, input *UpdateGlobalAccountInput) (*request.Request, *GlobalAccountOutput) {
	op := &request.Operation{
		Name: globalAccounts,
		Http: request.HTTP{
			Method: request.PATCH,
			Path:   "/globalAccount",
		},
	}

	if input == nil {
		input = &UpdateGlobalAccountInput{}
	}

	output := &GlobalAccountOutput{}
	return c.newRequest(ctx, op, input, output), output
}
