package btpevents

import (
	"context"
	"github.com/nnicora/sap-sdk-go/internal/times"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"time"
)

const events = "Events"

// GET /cloud-management/v1/events
// Get events
type GetEventsInput struct {
	//The ID of the event.
	Id []int `dest:"querystring" dest-name:"id"`

	//The ID of the entity associated with the event.
	EntityId string `dest:"querystring" dest-name:"entityId"`

	//The type of entity associated with the event.
	//For example: Subaccount, Directory, Tenant.
	EntityType []string `dest:"querystring" dest-name:"entityType"`

	//The type of the event that was triggered.
	//There are two groups of event types, Central Events and Local Events group.
	//The group you get depends on the scopes granted to you after you authorized to use the API.
	//You can query any event listed in the group that is relevant for your scope.
	//The examples of some event types in each of the groups:
	//
	//Central Events group: GlobalAccount_Update, AccountDirectory_Creation, AccountDirectory_Update,
	//	AccountDirectory_Update_Type, AccountDirectory_Deletion, Subaccount_Creation, Subaccount_Deletion,
	//	Subaccount_Update, Subaccount_Move, AccountDirectoryTenant_Creation, AccountDirectoryTenant_Deletion,
	//	GlobalAccountEntitlements_Update, EntityEntitlements_Update, EntityEntitlements_Move
	//
	//Local Events group: SubaccountAppSubscription_Creation, SubaccountAppSubscription_Deletion, SubaccountAppSubscription_Update,
	//	AppRegistration_Creation, AppRegistration_Deletion, AppRegistration_Update, SubaccountTenant_Creation,
	//	SubaccountTenant_Update, SubaccountTenant_Deletion, EnvironmentInstance_Creation, EnvironmentInstance_Deletion,
	//	EnvironmentInstances_Deletion
	EventType []string `dest:"querystring" dest-name:"eventType"`

	//Start date and time to query the events by the action that triggered them.
	//Use the Unix epoch time in milliseconds (you can find an online converter from a regular date-time format to the Unix epoch time format).
	//For example:
	//	Monday, June 1, 2020 9:40:22 AM is 1590993622000 in Unix epoch milliseconds time.
	FromActionTime time.Time `dest:"querystring" dest-name:"fromActionTime" timestampFormat:"unixTimestamp"`

	//Start date and time to query the events by when they were created.
	//Use the Unix epoch time in milliseconds (you can find an online converter from a regular date-time format to the Unix epoch time format).
	//For example:
	//	Monday, June 10, 2020 04:32:22 AM is 1591752742000 in Unix epoch milliseconds time.
	FromCreationTime time.Time `dest:"querystring" dest-name:"fromCreationTime" timestampFormat:"unixTimestamp"`

	//The page number to retrieve.
	PageNum uint32 `dest:"querystring" dest-name:"pageNum"`

	//The number of events to retrieve per page (max = 150).
	PageSize uint32 `dest:"querystring" dest-name:"pageSize"`

	//Field by which to sort the events.
	SortField string `dest:"querystring" dest-name:"sortField"`

	//Sort order for the events.
	//Can be ascending or descending.
	//
	//Available values : ASC, DESC
	SortOrder string `dest:"querystring" dest-name:"sortOrder"`

	//End date and time to query the events by the action that triggered them.
	//Use the Unix epoch time in milliseconds (you can find an online converter from a regular date-time format to the Unix epoch time format).
	//For example:
	//	Monday, June 4, 2020 11:40:22 AM is 1591260022000 in Unix epoch milliseconds time.
	ToActionTime time.Time `dest:"querystring" dest-name:"toActionTime" timestampFormat:"unixTimestamp"`

	//End date and time to query the events by when they were created.
	//Use the Unix epoch time in milliseconds (you can find an online converter from a regular date-time format to the Unix epoch time format).
	//For example:
	//	Monday, June 6, 2020 12:32:22 AM is 1591392742000 in Unix epoch milliseconds time.
	ToCreationTime time.Time `dest:"querystring" dest-name:"toCreationTime" timestampFormat:"unixTimestamp"`
}
type GetEventsOutput struct {
	//Lists of the events associated with the API call and used scopes.
	Events []Event `json:"events"`

	//Whether there are more pages.
	MorePages bool `json:"morePages"`

	//The current page number.
	PageNum int32 `json:"pageNum"`

	//Total numbers of results.
	Total int64 `json:"total"`

	//Total numbers of pages.
	TotalPages int64 `json:"totalPages"`
}
type Event struct {
	//The ID of the event.
	Id int64 `json:"id"`

	//The time the action triggered the event.
	//The format is Unix epoch time in milliseconds.
	ActionTime times.JavaTime `json:"actionTime"`

	//The time when the event record was created.
	//The format is Unix epoch time in milliseconds.
	CreationTime times.JavaTime `json:"creationTime"`

	//JSON object that contains description and details about the requested events.
	Details map[string]interface{} `json:"details"`

	//The ID of the entity associated with the event.
	EntityId string `json:"entityId"`

	//The type of entity associated with the event.
	EntityType string `json:"entityType"`

	//The service that reported the event.
	EventOrigin string `json:"eventOrigin"`

	//The type of the event that was triggered.
	//There are two groups of event types: Local Events and Central Events group.
	//Only event types that belong to one of the groups are returned as the result of a single API call.
	//The event types group you get depends on the scope you used to access the API.
	//The examples of some of the events for each of the groups:
	//
	//Central Events group: GlobalAccount_Update, AccountDirectory_Creation, AccountDirectory_Update, AccountDirectory_Update_Type,
	//	AccountDirectory_Deletion, Subaccount_Creation, Subaccount_Deletion, Subaccount_Update, Subaccount_Move, AccountDirectoryTenant_Creation,
	//	AccountDirectoryTenant_Deletion, GlobalAccountEntitlements_Update, EntityEntitlements_Update, EntityEntitlements_Move
	//
	//Local Events group: SubaccountAppSubscription_Creation, SubaccountAppSubscription_Deletion, SubaccountAppSubscription_Update,
	//	AppRegistration_Creation, AppRegistration_Deletion, AppRegistration_Update, SubaccountTenant_Creation, SubaccountTenant_Update,
	//	SubaccountTenant_Deletion, EnvironmentInstance_Creation, EnvironmentInstance_Deletion, EnvironmentInstances_Deletion
	EventType string `json:"eventType"`

	//The unique ID of the global account associated with the event.
	GlobalAccountGuid string `json:"globalAccountGUID"`
}

func (c *EventsV1) GetEvents(ctx context.Context, input *GetEventsInput) (*GetEventsOutput, error) {
	req, out := c.getEventRequest(ctx, input)
	return out, req.Send()
}
func (c *EventsV1) getEventRequest(ctx context.Context, input *GetEventsInput) (*request.Request, *GetEventsOutput) {
	op := &request.Operation{
		Name: events,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/events",
		},
	}

	if input == nil {
		input = &GetEventsInput{}
	}

	output := &GetEventsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /cloud-management/v1/events/types
// Get events
type GetEventsTypesInput struct {
}
type GetEventsTypesOutput struct {
	//Category to which the event type belongs.
	//
	//LOCAL: The event is associated with the local region within a multi-region universe.
	//CENTRAL: The event is associated with the central region within a multi-region universe.
	//Enum:
	//	[ LOCAL, CENTRAL ]
	Category string `json:"category"`

	//The description of the event type.
	Description string `json:"description"`

	//List of all the search parameters for the event type.
	SearchParams []string `json:"searchParams"`

	//The type of the event that was triggered.
	//There are two groups of event types: Local Events and Central Events group.
	//Only event types that belong to one of the groups are returned as the result of a single API call.
	//The event types group you get depends on the scope you used to access the API.
	//The examples of some of the events for each of the groups:
	//
	//Central Events group: GlobalAccount_Update, AccountDirectory_Creation, AccountDirectory_Update, AccountDirectory_Update_Type,
	//	AccountDirectory_Deletion, Subaccount_Creation, Subaccount_Deletion, Subaccount_Update, Subaccount_Move,
	//	AccountDirectoryTenant_Creation, AccountDirectoryTenant_Deletion, GlobalAccountEntitlements_Update,
	//	EntityEntitlements_Update, EntityEntitlements_Move
	//
	//Local Events group: SubaccountAppSubscription_Creation, SubaccountAppSubscription_Deletion, SubaccountAppSubscription_Update,
	//	AppRegistration_Creation, AppRegistration_Deletion, AppRegistration_Update, SubaccountTenant_Creation,
	//	SubaccountTenant_Update, SubaccountTenant_Deletion, EnvironmentInstance_Creation, EnvironmentInstance_Deletion,
	//	EnvironmentInstances_Deletion
	Type string `json:"type"`
}

func (c *EventsV1) GetEventsTypes(ctx context.Context) (*GetEventsTypesOutput, error) {
	req, out := c.getEventsTypesRequest(ctx, nil)
	return out, req.Send()
}
func (c *EventsV1) getEventsTypesRequest(ctx context.Context, input *GetEventsTypesInput) (*request.Request, *GetEventsTypesOutput) {
	op := &request.Operation{
		Name: events,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/events/types",
		},
	}

	if input == nil {
		input = &GetEventsTypesInput{}
	}

	output := &GetEventsTypesOutput{}
	return c.newRequest(ctx, op, input, output), output
}
