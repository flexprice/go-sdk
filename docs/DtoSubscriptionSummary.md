# DtoSubscriptionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivedAt** | Pointer to **string** | archived_at timestamp (for old subscriptions) | [optional] 
**BillingAnchor** | Pointer to **string** | billing_anchor of the subscription | [optional] 
**CreatedAt** | Pointer to **string** | created_at timestamp | [optional] 
**CurrentPeriodEnd** | Pointer to **string** | current_period_end of the subscription | [optional] 
**CurrentPeriodStart** | Pointer to **string** | current_period_start of the subscription | [optional] 
**Id** | Pointer to **string** | id of the subscription | [optional] 
**PlanId** | Pointer to **string** | plan_id of the subscription | [optional] 
**Status** | Pointer to [**TypesSubscriptionStatus**](TypesSubscriptionStatus.md) |  | [optional] 

## Methods

### NewDtoSubscriptionSummary

`func NewDtoSubscriptionSummary() *DtoSubscriptionSummary`

NewDtoSubscriptionSummary instantiates a new DtoSubscriptionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionSummaryWithDefaults

`func NewDtoSubscriptionSummaryWithDefaults() *DtoSubscriptionSummary`

NewDtoSubscriptionSummaryWithDefaults instantiates a new DtoSubscriptionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivedAt

`func (o *DtoSubscriptionSummary) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *DtoSubscriptionSummary) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *DtoSubscriptionSummary) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *DtoSubscriptionSummary) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetBillingAnchor

`func (o *DtoSubscriptionSummary) GetBillingAnchor() string`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *DtoSubscriptionSummary) GetBillingAnchorOk() (*string, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *DtoSubscriptionSummary) SetBillingAnchor(v string)`

SetBillingAnchor sets BillingAnchor field to given value.

### HasBillingAnchor

`func (o *DtoSubscriptionSummary) HasBillingAnchor() bool`

HasBillingAnchor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoSubscriptionSummary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoSubscriptionSummary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoSubscriptionSummary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoSubscriptionSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *DtoSubscriptionSummary) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *DtoSubscriptionSummary) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *DtoSubscriptionSummary) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *DtoSubscriptionSummary) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *DtoSubscriptionSummary) GetCurrentPeriodStart() string`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *DtoSubscriptionSummary) GetCurrentPeriodStartOk() (*string, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *DtoSubscriptionSummary) SetCurrentPeriodStart(v string)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *DtoSubscriptionSummary) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetId

`func (o *DtoSubscriptionSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoSubscriptionSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoSubscriptionSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoSubscriptionSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoSubscriptionSummary) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoSubscriptionSummary) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoSubscriptionSummary) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoSubscriptionSummary) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetStatus

`func (o *DtoSubscriptionSummary) GetStatus() TypesSubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoSubscriptionSummary) GetStatusOk() (*TypesSubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoSubscriptionSummary) SetStatus(v TypesSubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoSubscriptionSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


