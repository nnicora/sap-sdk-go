package btpresources

import (
	"context"
	"github.com/nnicora/sap-sdk-go/sap/http/request"
	"github.com/nnicora/sap-sdk-go/service/types"
)

const resourceConsumption = "Resource Consumption"

// GET /reports/v1/cloudCreditsDetails
// Get cloud credit data for a global account

type GetCloudCreditsDetailsInput struct {
	//Show the cloud credit history:
	//CURRENT: For the current phase; the default setting.
	//ALL: For all phases.
	//Available values : ALL, CURRENT
	ViewPhases string `dest:"querystring" dest-name:"viewPhases"`
}
type GetCloudCreditsDetailsOutput struct {
	Contracts []Contract `json:"contracts,omitempty"`
	//The unique ID of the global account.
	GlobalAccountId string `json:"globalAccountId,omitempty"`
	//The display name of the global account.
	GlobalAccountName string `json:"globalAccountName,omitempty"`

	types.StatusAndBodyFromResponse
}
type Contract struct {
	//The date that the contract finishes. Date is in the format YYYY-MM-DD
	ContractEndDate string `json:"contractEndDate,omitempty"`
	//The date that the contract begins. Date is in the format YYYY-MM-DD.
	ContractStartDate string `json:"contractStartDate,omitempty"`
	//The currency used to pay for the contract.
	Currency string `json:"currency,omitempty"`
	//The period for which a contract is purchased is broken down into smaller parts and each part is called a phase.
	Phases []Phase `json:"phases,omitempty"`
}
type Phase struct {
	//End date is in the format YYYY-MM-DD.
	EndDate string `json:"phaseEndDate,omitempty"`
	//Start date is in the format YYYY-MM-DD.
	StartDate string `json:"phaseStartDate,omitempty"`
	//History relating to phase updates.
	Updates []PhaseUpdate `json:"phaseUpdates,omitempty"`
}
type PhaseUpdate struct {
	//The residual amount of cloud credits available.
	Balance float64 `json:"balance,omitempty"`
	//The complete amount of cloud credits available in this phase.
	CloudCreditsForPhase float64 `json:"cloudCreditsForPhase,omitempty"`
	//The date that the phase was updated. Date is in the format YYYY-MM-DD.
	UpdatedOn string `json:"phaseUpdatedOn,omitempty"`
}

func (c *ResourceV1) GetCloudCreditsDetails(ctx context.Context, input *GetCloudCreditsDetailsInput) (*GetCloudCreditsDetailsOutput, error) {
	req, out := c.getCloudCreditsDetailsRequest(ctx, input)
	return out, req.Send()
}
func (c *ResourceV1) getCloudCreditsDetailsRequest(ctx context.Context, input *GetCloudCreditsDetailsInput) (*request.Request, *GetCloudCreditsDetailsOutput) {
	op := &request.Operation{
		Name: resourceConsumption,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/cloudCreditsDetails",
		},
	}

	if input == nil {
		input = &GetCloudCreditsDetailsInput{}
	}

	output := &GetCloudCreditsDetailsOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /reports/v1/monthlySubaccountsCost
// Get monthly cost reporting data for all subaccounts
type GetMonthlySubAccountsCostInput struct {
	// Start date for querying the global account’s monthly usage data
	//
	//Example:
	// fromDate=201901, toDate=201912
	// This query will return the usage data for the period between January 2019 and December 2019.
	FromDate uint32 `dest:"querystring" dest-name:"fromDate"`

	// End date for querying the global account’s monthly usage data
	//
	//Example:
	// fromDate=201901, toDate=201912
	// This query will return the usage data for the period between January 2019 and December 2019.
	ToDate uint32 `dest:"querystring" dest-name:"fromDate"`
}
type GetMonthlySubAccountsCostOutput struct {
	Content []MonthlySubAccountsCost `json:"content,omitempty"`

	types.StatusAndBodyFromResponse
}
type MonthlySubAccountsCost struct {
	//The subaccount usage cost for a specified month.
	Cost float64 `json:"cost,omitempty"`
	//The SKU of the service consumed.
	CrmSku string `json:"crmSku,omitempty"`
	//The currency in which costs are shown. Defined on the global account level, upon signing the contract.
	Currency string `json:"currency,omitempty"`
	//The technical name of the landscape, (as identified by core services for SAP BTP), on which the usage was
	//originally initialized. Example values: cf-us10-staging, cf-eu10-canary, cf-eu20.
	DataCenter string `json:"dataCenter,omitempty"`
	//The descriptive name of the data center.
	DataCenterName string `json:"dataCenterName,omitempty"`
	//The unique ID of the directory.
	DirectoryId string `json:"directoryId,omitempty"`
	//The descriptive name of the directory for customer-facing UIs.
	DirectoryName string `json:"directoryName,omitempty"`
	//The billing status of the billable item. If TRUE the item was not billed.
	Estimated bool `json:"estimated,omitempty"`
	//The unique ID of the global account to which the subaccounts belong, and which is the context for billing the customer.
	GlobalAccountId string `json:"globalAccountId,omitempty"`
	//The descriptive name of the global account for customer-facing UIs.
	GlobalAccountName string `json:"globalAccountName,omitempty"`
	//The original measure of the usage as reported by the technical usage API payload.
	MeasureId string `json:"measureId,omitempty"`
	//The name of the metric used by cloud services for customer-facing UIs.
	MetricName string `json:"metricName,omitempty"`
	//The ID of the service plan to which the measured usage data is related.
	Plan string `json:"plan,omitempty"`
	//The name of the plan for customer-facing UIs.
	PlanName string `json:"planName,omitempty"`
	//The year and month for which the cost is reported.
	ReportYearMonth int32 `json:"reportYearMonth,omitempty"`
	//The ID of the service to which the measured usage data is related.
	ServiceId string `json:"serviceId,omitempty"`
	//The name of the service for customer-facing UIs.
	ServiceName string `json:"serviceName,omitempty"`
	//The unique ID of the subaccount for which to get the usage data.
	SubAccountId string `json:"subaccountId,omitempty"`
	//The descriptive name of the subaccount for customer-facing UIs.
	SubAccountName string `json:"subaccountName,omitempty"`
	//Predefined name for more than one unit of usage for the given metric. Generally a short name for use in customer-facing UIs.
	UnitPlural string `json:"unitPlural,omitempty"`
	//Pre-defined name for one unit of usage.
	UnitSingular string `json:"unitSingular,omitempty"`
	//The reported usage in numbers for the given metric.
	Usage string `json:"usage,omitempty"`
}

func (c *ResourceV1) GetMonthlySubAccountsCost(ctx context.Context, input *GetMonthlySubAccountsCostInput) (*GetMonthlySubAccountsCostOutput, error) {
	req, out := c.getMonthlySubAccountsCostRequest(ctx, input)
	return out, req.Send()
}
func (c *ResourceV1) getMonthlySubAccountsCostRequest(ctx context.Context, input *GetMonthlySubAccountsCostInput) (*request.Request, *GetMonthlySubAccountsCostOutput) {
	op := &request.Operation{
		Name: resourceConsumption,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/monthlySubaccountsCost",
		},
	}

	if input == nil {
		input = &GetMonthlySubAccountsCostInput{}
	}

	output := &GetMonthlySubAccountsCostOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /reports/v1/monthlyUsage
// Get monthly usage reporting data for a global account
type GetMonthlyUsageInput struct {
	// Start date for querying the global account’s monthly usage data
	//
	//Example:
	// fromDate=201901, toDate=201912
	// This query will return the usage data for the period between January 2019 and December 2019.
	FromDate uint32 `dest:"querystring" dest-name:"fromDate"`

	// End date for querying the global account’s monthly usage data
	//
	//Example:
	// fromDate=201901, toDate=201912
	// This query will return the usage data for the period between January 2019 and December 2019.
	ToDate uint32 `dest:"querystring" dest-name:"fromDate"`
}
type GetMonthlyUsageOutput struct {
	Content []MonthlyUsage `json:"content,omitempty"`

	types.StatusAndBodyFromResponse
}
type MonthlyUsage struct {
	// The technical name of the landscape, (as identified by core services for SAP BTP),
	// on which the usage was originally initialized. Example values: cf-us10-staging, cf-eu10-canary, cf-eu20.
	DataCenter string `json:"dataCenter,omitempty"`
	// The descriptive name of the data center.
	DataCenterName string `json:"dataCenterName,omitempty"`
	//The unique ID of the directory.
	DirectoryId string `json:"directoryId,omitempty"`
	//The descriptive name of the directory for customer-facing UIs.
	DirectoryName string `json:"directoryName,omitempty"`
	//The unique ID of the consumer environment instance.
	EnvironmentInstanceId string `json:"environmentInstanceId,omitempty"`
	//The name of the consumer environment instance for customer-facing UIs.
	EnvironmentInstanceName string `json:"environmentInstanceName,omitempty"`
	//The unique ID of the global account to which the subaccounts belong, and which is the context for billing the customer.
	GlobalAccountId string `json:"globalAccountId,omitempty"`
	//The descriptive name of the global account for customer-facing UIs.
	GlobalAccountName string `json:"globalAccountName,omitempty"`
	//Consumer identity zone.
	IdentityZone string `json:"identityZone,omitempty"`
	//Consumer instance ID.
	InstanceId string `json:"instanceId,omitempty"`
	//The original measure of the usage as reported by the technical usage API payload.
	MeasureId string `json:"measureId,omitempty"`
	//The name of the metric used by cloud services for customer-facing UIs.
	MetricName string `json:"metricName,omitempty"`
	//The ID of the service plan to which the measured usage data is related.
	Plan string `json:"plan,omitempty"`
	//The name of the plan for customer-facing UIs.
	PlanName string `json:"planName,omitempty"`
	//The year and month for which the cost is reported.
	ReportYearMonth int32 `json:"reportYearMonth,omitempty"`
	//The ID of the service to which the measured usage data is related.
	ServiceId string `json:"serviceId,omitempty"`
	//The name of the service for customer-facing UIs.
	ServiceName string `json:"serviceName,omitempty"`
	//The ID of the consumer space.
	SpaceId string `json:"spaceId,omitempty"`
	//The descriptive name of the consumer space for customer-facing UIs.
	SpaceName string `json:"spaceName,omitempty"`
	//The unique ID of the subaccount for which to get the usage data.
	SubAccountId string `json:"subaccountId,omitempty"`
	//The descriptive name of the subaccount for customer-facing UIs.
	SubAccountName string `json:"subaccountName,omitempty"`
	//Predefined name for more than one unit of usage for the given metric. Generally a short name for use in
	//customer-facing UIs.
	UnitPlural string `json:"unitPlural,omitempty"`
	//Pre-defined name for one unit of usage.
	UnitSingular string `json:"unitSingular,omitempty"`
	//The reported usage in numbers for the given metric.
	Usage float64 `json:"usage,omitempty"`
}

func (c *ResourceV1) GetMonthlyUsage(ctx context.Context, input *GetMonthlyUsageInput) (*GetMonthlyUsageOutput, error) {
	req, out := c.getMonthlyUsageRequest(ctx, input)
	return out, req.Send()
}
func (c *ResourceV1) getMonthlyUsageRequest(ctx context.Context, input *GetMonthlyUsageInput) (*request.Request, *GetMonthlyUsageOutput) {
	op := &request.Operation{
		Name: resourceConsumption,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/monthlyUsage",
		},
	}

	if input == nil {
		input = &GetMonthlyUsageInput{}
	}

	output := &GetMonthlyUsageOutput{}
	return c.newRequest(ctx, op, input, output), output
}

// GET /reports/v1/subaccountUsage
// Get usage reporting data for a subaccount
type GetSubAccountUsageInput struct {
	//Security token containing claims about the authentication of an end user by the
	// authorization server (Identity Authentication).
	XIDToken string `dest:"header" dest-name:"X-ID-Token"`

	//Start date for querying the subaccount usage data using the format YYYYMMDD.
	//
	//Example:
	// fromDate=20190101, toDate=20191201
	// This query returns the subaccount usage data for the period January 1st, 2019 to December 1st, 2019.
	FromDate uint32 `dest:"querystring" dest-name:"fromDate"`

	//The time division of the subaccount usage report, namely, DAY, WEEK and MONTH according to the specified
	// time period. If no period perspective is defined, then the subaccount usage data is returned for the entire period as a single element.
	// If you select DAY, the maximum search period is four months. If you select WEEK, the search period must not exceed one year.
	//Example:
	// periodPerspective=WEEK
	// This query returns the subaccount usage data aggregated by week.
	//
	//Available values : DAY, WEEK, MONTH
	PeriodPerspective string `dest:"querystring" dest-name:"periodPerspective"`

	// Unique ID of the subaccount.
	SubAccountId string `dest:"querystring" dest-name:"subaccountId"`

	//End date for querying the subaccount usage data using the format YYYYMMDD.
	//
	//Example:
	// fromDate=20190101, toDate=20191201
	// This query returns the subaccount usage data for the period January 1st, 2019 to December 1st, 2019.
	ToDate uint32 `dest:"querystring" dest-name:"fromDate"`
}
type GetSubAccountUsageOutput struct {
	Content []SubAccountUsage `json:"content,omitempty"`

	types.StatusAndBodyFromResponse
}
type SubAccountUsage struct {
	//The unique ID of the product category.
	CategoryId float64 `json:"categoryId,omitempty"`
	//The name of the product category.
	CategoryName string `json:"categoryName,omitempty"`
	//The technical name of the landscape, (as identified by core services for SAP BTP), on which the usage was
	//originally initialized. Example values: cf-us10-staging, cf-eu10-canary, cf-eu20.
	DataCenter string `json:"dataCenter,omitempty"`
	//The descriptive name of the data center.
	DataCenterName string `json:"dataCenterName,omitempty"`
	//The unique ID of the directory.
	DirectoryId string `json:"directoryId,omitempty"`
	//The descriptive name of the directory for customer-facing UIs.
	DirectoryName string `json:"directoryName,omitempty"`
	//The unique ID of the consumer environment instance.
	EnvironmentInstanceId string `json:"environmentInstanceId,omitempty"`
	//The name of the consumer environment instance for customer-facing UIs.
	EnvironmentInstanceName string `json:"environmentInstanceName,omitempty"`
	//The unique ID of the global account to which the subaccounts belong, and which is the context for billing the customer.
	GlobalAccountId string `json:"globalAccountId,omitempty"`
	//The descriptive name of the global account for customer-facing UIs.
	GlobalAccountName string `json:"globalAccountName,omitempty"`
	//Consumer identity zone.
	IdentityZone string `json:"identityZone,omitempty"`
	//Consumer instance ID.
	InstanceId string `json:"instanceId,omitempty"`
	//The original measure of the usage as reported by the technical usage API payload.
	MeasureId string `json:"measureId,omitempty"`
	//The name of the metric used by cloud services for customer-facing UIs.
	MetricName string `json:"metricName,omitempty"`
	//The last day of the time division requested for the subaccount usage report.
	PeriodEndDate uint32 `json:"periodEndDate,omitempty"`
	//The first day of the time division requested for the subaccount usage report.
	PeriodStartDate uint32 `json:"periodStartDate,omitempty"`
	//The ID of the service plan to which the measured usage data is related.
	Plan string `json:"plan,omitempty"`
	//The name of the plan for customer-facing UIs.
	PlanName string `json:"planName,omitempty"`
	//The ID of the service to which the measured usage data is related.
	ServiceId string `json:"serviceId,omitempty"`
	//The name of the service for customer-facing UIs.
	ServiceName string `json:"serviceName,omitempty"`
	//The ID of the consumer space.
	SpaceId string `json:"spaceId,omitempty"`
	//The descriptive name of the consumer space for customer-facing UIs.
	SpaceName string `json:"spaceName,omitempty"`
	//The unique ID of the subaccount for which to get the usage data.
	SubAccountId string `json:"subaccountId,omitempty"`
	//The descriptive name of the subaccount for customer-facing UIs.
	SubAccountName string `json:"subaccountName,omitempty"`
	//Predefined name for more than one unit of usage for the given metric. Generally a short name for use in customer-facing UIs.
	UnitPlural string `json:"unitPlural,omitempty"`
	//Pre-defined name for one unit of usage.
	UnitSingular string `json:"unitSingular,omitempty"`
	//The reported usage in numbers for the given metric.
	Usage float64 `json:"usage,omitempty"`
}

func (c *ResourceV1) GetSubAccountUsage(ctx context.Context, input *GetSubAccountUsageInput) (*GetSubAccountUsageOutput, error) {
	req, out := c.getSubAccountUsageRequest(ctx, input)
	return out, req.Send()
}
func (c *ResourceV1) getSubAccountUsageRequest(ctx context.Context, input *GetSubAccountUsageInput) (*request.Request, *GetSubAccountUsageOutput) {
	op := &request.Operation{
		Name: resourceConsumption,
		Http: request.HTTP{
			Method: request.GET,
			Path:   "/subaccountUsage",
		},
	}

	if input == nil {
		input = &GetSubAccountUsageInput{}
	}

	output := &GetSubAccountUsageOutput{}
	return c.newRequest(ctx, op, input, output), output
}
