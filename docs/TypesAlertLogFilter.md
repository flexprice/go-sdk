# TypesAlertLogFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertStatus** | Pointer to [**TypesAlertState**](TypesAlertState.md) |  | [optional] 
**AlertType** | Pointer to [**TypesAlertType**](TypesAlertType.md) |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to [**TypesAlertEntityType**](TypesAlertEntityType.md) |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) | filters allows complex filtering based on multiple fields | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 

## Methods

### NewTypesAlertLogFilter

`func NewTypesAlertLogFilter() *TypesAlertLogFilter`

NewTypesAlertLogFilter instantiates a new TypesAlertLogFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesAlertLogFilterWithDefaults

`func NewTypesAlertLogFilterWithDefaults() *TypesAlertLogFilter`

NewTypesAlertLogFilterWithDefaults instantiates a new TypesAlertLogFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertStatus

`func (o *TypesAlertLogFilter) GetAlertStatus() TypesAlertState`

GetAlertStatus returns the AlertStatus field if non-nil, zero value otherwise.

### GetAlertStatusOk

`func (o *TypesAlertLogFilter) GetAlertStatusOk() (*TypesAlertState, bool)`

GetAlertStatusOk returns a tuple with the AlertStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertStatus

`func (o *TypesAlertLogFilter) SetAlertStatus(v TypesAlertState)`

SetAlertStatus sets AlertStatus field to given value.

### HasAlertStatus

`func (o *TypesAlertLogFilter) HasAlertStatus() bool`

HasAlertStatus returns a boolean if a field has been set.

### GetAlertType

`func (o *TypesAlertLogFilter) GetAlertType() TypesAlertType`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *TypesAlertLogFilter) GetAlertTypeOk() (*TypesAlertType, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *TypesAlertLogFilter) SetAlertType(v TypesAlertType)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *TypesAlertLogFilter) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetCustomerId

`func (o *TypesAlertLogFilter) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TypesAlertLogFilter) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TypesAlertLogFilter) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *TypesAlertLogFilter) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEndTime

`func (o *TypesAlertLogFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesAlertLogFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesAlertLogFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesAlertLogFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntityId

`func (o *TypesAlertLogFilter) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *TypesAlertLogFilter) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *TypesAlertLogFilter) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *TypesAlertLogFilter) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *TypesAlertLogFilter) GetEntityType() TypesAlertEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TypesAlertLogFilter) GetEntityTypeOk() (*TypesAlertEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TypesAlertLogFilter) SetEntityType(v TypesAlertEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *TypesAlertLogFilter) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetExpand

`func (o *TypesAlertLogFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesAlertLogFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesAlertLogFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesAlertLogFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFilters

`func (o *TypesAlertLogFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TypesAlertLogFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TypesAlertLogFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TypesAlertLogFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *TypesAlertLogFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesAlertLogFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesAlertLogFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesAlertLogFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *TypesAlertLogFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesAlertLogFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesAlertLogFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesAlertLogFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesAlertLogFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesAlertLogFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesAlertLogFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesAlertLogFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSort

`func (o *TypesAlertLogFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesAlertLogFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesAlertLogFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesAlertLogFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesAlertLogFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesAlertLogFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesAlertLogFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesAlertLogFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesAlertLogFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesAlertLogFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesAlertLogFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesAlertLogFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


