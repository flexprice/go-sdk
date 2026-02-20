# TypesPriceFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowExpiredPrices** | Pointer to **bool** |  | [optional] [default to false]
**EndTime** | Pointer to **string** |  | [optional] 
**EntityIds** | Pointer to **[]string** |  | [optional] 
**EntityType** | Pointer to [**TypesPriceEntityType**](TypesPriceEntityType.md) |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) | DSL filters | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**MeterIds** | Pointer to **[]string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**ParentPriceId** | Pointer to **string** |  | [optional] 
**PlanIds** | Pointer to **[]string** | Price override filtering fields | [optional] 
**PriceIds** | Pointer to **[]string** |  | [optional] 
**Sort** | Pointer to **string** |  | [optional] 
**StartDateLt** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewTypesPriceFilter

`func NewTypesPriceFilter() *TypesPriceFilter`

NewTypesPriceFilter instantiates a new TypesPriceFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesPriceFilterWithDefaults

`func NewTypesPriceFilterWithDefaults() *TypesPriceFilter`

NewTypesPriceFilterWithDefaults instantiates a new TypesPriceFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowExpiredPrices

`func (o *TypesPriceFilter) GetAllowExpiredPrices() bool`

GetAllowExpiredPrices returns the AllowExpiredPrices field if non-nil, zero value otherwise.

### GetAllowExpiredPricesOk

`func (o *TypesPriceFilter) GetAllowExpiredPricesOk() (*bool, bool)`

GetAllowExpiredPricesOk returns a tuple with the AllowExpiredPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExpiredPrices

`func (o *TypesPriceFilter) SetAllowExpiredPrices(v bool)`

SetAllowExpiredPrices sets AllowExpiredPrices field to given value.

### HasAllowExpiredPrices

`func (o *TypesPriceFilter) HasAllowExpiredPrices() bool`

HasAllowExpiredPrices returns a boolean if a field has been set.

### GetEndTime

`func (o *TypesPriceFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesPriceFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesPriceFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesPriceFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntityIds

`func (o *TypesPriceFilter) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *TypesPriceFilter) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *TypesPriceFilter) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *TypesPriceFilter) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### GetEntityType

`func (o *TypesPriceFilter) GetEntityType() TypesPriceEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TypesPriceFilter) GetEntityTypeOk() (*TypesPriceEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TypesPriceFilter) SetEntityType(v TypesPriceEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *TypesPriceFilter) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetExpand

`func (o *TypesPriceFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesPriceFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesPriceFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesPriceFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFilters

`func (o *TypesPriceFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TypesPriceFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TypesPriceFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TypesPriceFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *TypesPriceFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesPriceFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesPriceFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesPriceFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMeterIds

`func (o *TypesPriceFilter) GetMeterIds() []string`

GetMeterIds returns the MeterIds field if non-nil, zero value otherwise.

### GetMeterIdsOk

`func (o *TypesPriceFilter) GetMeterIdsOk() (*[]string, bool)`

GetMeterIdsOk returns a tuple with the MeterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterIds

`func (o *TypesPriceFilter) SetMeterIds(v []string)`

SetMeterIds sets MeterIds field to given value.

### HasMeterIds

`func (o *TypesPriceFilter) HasMeterIds() bool`

HasMeterIds returns a boolean if a field has been set.

### GetOffset

`func (o *TypesPriceFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesPriceFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesPriceFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesPriceFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesPriceFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesPriceFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesPriceFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesPriceFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetParentPriceId

`func (o *TypesPriceFilter) GetParentPriceId() string`

GetParentPriceId returns the ParentPriceId field if non-nil, zero value otherwise.

### GetParentPriceIdOk

`func (o *TypesPriceFilter) GetParentPriceIdOk() (*string, bool)`

GetParentPriceIdOk returns a tuple with the ParentPriceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPriceId

`func (o *TypesPriceFilter) SetParentPriceId(v string)`

SetParentPriceId sets ParentPriceId field to given value.

### HasParentPriceId

`func (o *TypesPriceFilter) HasParentPriceId() bool`

HasParentPriceId returns a boolean if a field has been set.

### GetPlanIds

`func (o *TypesPriceFilter) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *TypesPriceFilter) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *TypesPriceFilter) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *TypesPriceFilter) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetPriceIds

`func (o *TypesPriceFilter) GetPriceIds() []string`

GetPriceIds returns the PriceIds field if non-nil, zero value otherwise.

### GetPriceIdsOk

`func (o *TypesPriceFilter) GetPriceIdsOk() (*[]string, bool)`

GetPriceIdsOk returns a tuple with the PriceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceIds

`func (o *TypesPriceFilter) SetPriceIds(v []string)`

SetPriceIds sets PriceIds field to given value.

### HasPriceIds

`func (o *TypesPriceFilter) HasPriceIds() bool`

HasPriceIds returns a boolean if a field has been set.

### GetSort

`func (o *TypesPriceFilter) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesPriceFilter) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesPriceFilter) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesPriceFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartDateLt

`func (o *TypesPriceFilter) GetStartDateLt() string`

GetStartDateLt returns the StartDateLt field if non-nil, zero value otherwise.

### GetStartDateLtOk

`func (o *TypesPriceFilter) GetStartDateLtOk() (*string, bool)`

GetStartDateLtOk returns a tuple with the StartDateLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateLt

`func (o *TypesPriceFilter) SetStartDateLt(v string)`

SetStartDateLt sets StartDateLt field to given value.

### HasStartDateLt

`func (o *TypesPriceFilter) HasStartDateLt() bool`

HasStartDateLt returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesPriceFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesPriceFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesPriceFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesPriceFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesPriceFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesPriceFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesPriceFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesPriceFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *TypesPriceFilter) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *TypesPriceFilter) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *TypesPriceFilter) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *TypesPriceFilter) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


