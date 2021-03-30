package btpaccounts

import "github.com/nnicora/sap-sdk-go/internal/times"

type SubAccount struct {
	//Whether the subaccount can use beta services and applications.
	BetaEnabled bool `json:"betaEnabled"`

	//Details of the user that created the subaccount.
	CreatedBy string `json:"createdBy"`

	//The date the subaccount was created. Dates and times are in UTC format.
	CreatedDate times.JavaTime `json:"createdDate"`

	//The custom properties assigned to the subaccount.
	CustomProperties []CustomProperties `json:"customProperties"`

	//A description of the subaccount for customer-facing UIs.
	Description string `json:"description"`

	//A descriptive name of the subaccount for customer-facing UIs.
	DisplayName string `json:"displayName"`

	//The unique ID of the subaccount's global account.
	GlobalAccountGuid string `json:"globalAccountGUID"`

	//Unique ID of the subaccount.
	Guid string `json:"guid"`

	//The date the subaccount was last modified. Dates and times are in UTC format.
	ModifiedDate times.JavaTime `json:"modifiedDate"`

	//The features of parent entity of the subaccount.
	//
	//Enum:
	//	[ DEFAULT, ENTITLEMENTS, AUTHORIZATIONS, CRM ]
	ParentFeatures []string `json:"parentFeatures"`

	//The GUID of the subaccount’s parent entity. If the subaccount is located directly in the global account
	//	(not in a directory), then this is the GUID of the global account.
	ParentGuid string `json:"parentGUID"`

	//The region in which the subaccount was created.
	Region string `json:"region"`

	//The current state of the subaccount.
	//
	//Enum:
	//	[ STARTED, CREATING, UPDATING, MOVING, PROCESSING, DELETING, OK, PENDING_REVIEW, CANCELED, CREATION_FAILED,
	//	UPDATE_FAILED, UPDATE_ACCOUNT_TYPE_FAILED, UPDATE_DIRECTORY_TYPE_FAILED, PROCESSING_FAILED, DELETION_FAILED,
	//	MOVE_FAILED, MIGRATING, MIGRATION_FAILED, ROLLBACK_MIGRATION_PROCESSING, MIGRATED ]
	State string `json:"state"`

	//Information about the state of the subaccount.
	StateMessage string `json:"stateMessage"`

	//The subdomain that becomes part of the path used to access the authorization tenant of the subaccount.
	//Must be unique within the defined region. Use only letters (a-z), digits (0-9), and hyphens (not at the start or end).
	//Maximum length is 63 characters. Cannot be changed after the subaccount has been created.
	Subdomain string `json:"subdomain"`

	//Whether the subaccount is used for production purposes. This flag can help your cloud operator to take appropriate
	// action when handling incidents that are related to mission-critical accounts in production systems.
	// Do not apply for subaccounts that are used for non-production purposes, such as development, testing, and demos.
	// Applying this setting this does not modify the subaccount.
	//
	//UNSET: Global account or subaccount admin has not set the production-relevancy flag. Default value.
	//NOT_USED_FOR_PRODUCTION: Subaccount is not used for production purposes.
	//USED_FOR_PRODUCTION: Subaccount is used for production purposes.
	//Enum:
	//	[ UNSET, USED_FOR_PRODUCTION, NOT_USED_FOR_PRODUCTION ]
	UsedForProduction string `json:"usedForProduction"`

	//The zoneId of the subaccount.
	ZoneId string `json:"zoneId"`
}

type CustomProperties struct {
	KeyValue

	//The unique id for the corresponding entity.
	AccountGuid string `json:"accountGUID"`
}

type KeyValue struct {
	//A name for the custom property.
	Key string `json:"key"`
	//A value for the corresponding key.
	Value string `json:"value"`
}
