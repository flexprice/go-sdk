# TypesSubscriptionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAt** | Pointer to **string** | ActiveAt filters subscriptions that are active at the given time | [optional] 
**BillingCadence** | Pointer to [**[]TypesBillingCadence**](TypesBillingCadence.md) | BillingCadence filters by billing cadence | [optional] 
**BillingPeriod** | Pointer to [**[]TypesBillingPeriod**](TypesBillingPeriod.md) | BillingPeriod filters by billing period | [optional] 
**CustomerId** | Pointer to **string** | CustomerID filters by customer ID | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**ExternalCustomerId** | Pointer to **string** | ExternalCustomerID filters by external customer ID | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) |  | [optional] 
**InvoicingCustomerIds** | Pointer to **[]string** | InvoicingCustomerIDs filters by invoicing customer ID | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** | PlanID filters by plan ID | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionIds** | Pointer to **[]string** |  | [optional] 
**SubscriptionStatus** | Pointer to [**[]TypesSubscriptionStatus**](TypesSubscriptionStatus.md) | SubscriptionStatus filters by subscription status | [optional] 
**WithLineItems** | Pointer to **bool** | WithLineItems includes line items in the response | [optional] 

## Methods

### NewTypesSubscriptionFilter

`func NewTypesSubscriptionFilter() *TypesSubscriptionFilter`

NewTypesSubscriptionFilter instantiates a new TypesSubscriptionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesSubscriptionFilterWithDefaults

`func NewTypesSubscriptionFilterWithDefaults() *TypesSubscriptionFilter`

NewTypesSubscriptionFilterWithDefaults instantiates a new TypesSubscriptionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAt

`func (o *TypesSubscriptionFilter) GetActiveAt() string`

GetActiveAt returns the ActiveAt field if non-nil, zero value otherwise.

### GetActiveAtOk

`func (o *TypesSubscriptionFilter) GetActiveAtOk() (*string, bool)`

GetActiveAtOk returns a tuple with the ActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAt

`func (o *TypesSubscriptionFilter) SetActiveAt(v string)`

SetActiveAt sets ActiveAt field to given value.

### HasActiveAt

`func (o *TypesSubscriptionFilter) HasActiveAt() bool`

HasActiveAt returns a boolean if a field has been set.

### GetBillingCadence

`func (o *TypesSubscriptionFilter) GetBillingCadence() []TypesBillingCadence`

GetBillingCadence returns the BillingCadence field if non-nil, zero value otherwise.

### GetBillingCadenceOk

`func (o *TypesSubscriptionFilter) GetBillingCadenceOk() (*[]TypesBillingCadence, bool)`

GetBillingCadenceOk returns a tuple with the BillingCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCadence

`func (o *TypesSubscriptionFilter) SetBillingCadence(v []TypesBillingCadence)`

SetBillingCadence sets BillingCadence field to given value.

### HasBillingCadence

`func (o *TypesSubscriptionFilter) HasBillingCadence() bool`

HasBillingCadence returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *TypesSubscriptionFilter) GetBillingPeriod() []TypesBillingPeriod`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *TypesSubscriptionFilter) GetBillingPeriodOk() (*[]TypesBillingPeriod, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *TypesSubscriptionFilter) SetBillingPeriod(v []TypesBillingPeriod)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *TypesSubscriptionFilter) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetCustomerId

`func (o *TypesSubscriptionFilter) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TypesSubscriptionFilter) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TypesSubscriptionFilter) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *TypesSubscriptionFilter) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEndTime

`func (o *TypesSubscriptionFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesSubscriptionFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesSubscriptionFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesSubscriptionFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpand

`func (o *TypesSubscriptionFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesSubscriptionFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesSubscriptionFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesSubscriptionFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetExternalCustomerId

`func (o *TypesSubscriptionFilter) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *TypesSubscriptionFilter) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *TypesSubscriptionFilter) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *TypesSubscriptionFilter) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetFilters

`func (o *TypesSubscriptionFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TypesSubscriptionFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TypesSubscriptionFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TypesSubscriptionFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetInvoicingCustomerIds

`func (o *TypesSubscriptionFilter) GetInvoicingCustomerIds() []string`

GetInvoicingCustomerIds returns the InvoicingCustomerIds field if non-nil, zero value otherwise.

### GetInvoicingCustomerIdsOk

`func (o *TypesSubscriptionFilter) GetInvoicingCustomerIdsOk() (*[]string, bool)`

GetInvoicingCustomerIdsOk returns a tuple with the InvoicingCustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingCustomerIds

`func (o *TypesSubscriptionFilter) SetInvoicingCustomerIds(v []string)`

SetInvoicingCustomerIds sets InvoicingCustomerIds field to given value.

### HasInvoicingCustomerIds

`func (o *TypesSubscriptionFilter) HasInvoicingCustomerIds() bool`

HasInvoicingCustomerIds returns a boolean if a field has been set.

### GetLimit

`func (o *TypesSubscriptionFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesSubscriptionFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesSubscriptionFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesSubscriptionFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *TypesSubscriptionFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesSubscriptionFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesSubscriptionFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesSubscriptionFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesSubscriptionFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesSubscriptionFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesSubscriptionFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesSubscriptionFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPlanId

`func (o *TypesSubscriptionFilter) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *TypesSubscriptionFilter) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *TypesSubscriptionFilter) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *TypesSubscriptionFilter) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetSort

`func (o *TypesSubscriptionFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesSubscriptionFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesSubscriptionFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesSubscriptionFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesSubscriptionFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesSubscriptionFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesSubscriptionFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesSubscriptionFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesSubscriptionFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesSubscriptionFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesSubscriptionFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesSubscriptionFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *TypesSubscriptionFilter) GetSubscriptionIds() []string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *TypesSubscriptionFilter) GetSubscriptionIdsOk() (*[]string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *TypesSubscriptionFilter) SetSubscriptionIds(v []string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *TypesSubscriptionFilter) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *TypesSubscriptionFilter) GetSubscriptionStatus() []TypesSubscriptionStatus`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *TypesSubscriptionFilter) GetSubscriptionStatusOk() (*[]TypesSubscriptionStatus, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *TypesSubscriptionFilter) SetSubscriptionStatus(v []TypesSubscriptionStatus)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *TypesSubscriptionFilter) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetWithLineItems

`func (o *TypesSubscriptionFilter) GetWithLineItems() bool`

GetWithLineItems returns the WithLineItems field if non-nil, zero value otherwise.

### GetWithLineItemsOk

`func (o *TypesSubscriptionFilter) GetWithLineItemsOk() (*bool, bool)`

GetWithLineItemsOk returns a tuple with the WithLineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithLineItems

`func (o *TypesSubscriptionFilter) SetWithLineItems(v bool)`

SetWithLineItems sets WithLineItems field to given value.

### HasWithLineItems

`func (o *TypesSubscriptionFilter) HasWithLineItems() bool`

HasWithLineItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


