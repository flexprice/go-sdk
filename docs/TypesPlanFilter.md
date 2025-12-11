# TypesPlanFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) | filters allows complex filtering based on multiple fields | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**LookupKey** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**PlanIds** | Pointer to **[]string** |  | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 

## Methods

### NewTypesPlanFilter

`func NewTypesPlanFilter() *TypesPlanFilter`

NewTypesPlanFilter instantiates a new TypesPlanFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesPlanFilterWithDefaults

`func NewTypesPlanFilterWithDefaults() *TypesPlanFilter`

NewTypesPlanFilterWithDefaults instantiates a new TypesPlanFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *TypesPlanFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesPlanFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesPlanFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesPlanFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpand

`func (o *TypesPlanFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesPlanFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesPlanFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesPlanFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFilters

`func (o *TypesPlanFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TypesPlanFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TypesPlanFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TypesPlanFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *TypesPlanFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesPlanFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesPlanFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesPlanFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLookupKey

`func (o *TypesPlanFilter) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *TypesPlanFilter) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *TypesPlanFilter) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *TypesPlanFilter) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetOffset

`func (o *TypesPlanFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesPlanFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesPlanFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesPlanFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesPlanFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesPlanFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesPlanFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesPlanFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPlanIds

`func (o *TypesPlanFilter) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *TypesPlanFilter) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *TypesPlanFilter) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *TypesPlanFilter) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetSort

`func (o *TypesPlanFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesPlanFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesPlanFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesPlanFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesPlanFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesPlanFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesPlanFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesPlanFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesPlanFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesPlanFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesPlanFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesPlanFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


