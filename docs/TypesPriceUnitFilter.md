# TypesPriceUnitFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) | filters allows complex filtering based on multiple fields | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**PriceUnitIds** | Pointer to **[]string** |  | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 

## Methods

### NewTypesPriceUnitFilter

`func NewTypesPriceUnitFilter() *TypesPriceUnitFilter`

NewTypesPriceUnitFilter instantiates a new TypesPriceUnitFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesPriceUnitFilterWithDefaults

`func NewTypesPriceUnitFilterWithDefaults() *TypesPriceUnitFilter`

NewTypesPriceUnitFilterWithDefaults instantiates a new TypesPriceUnitFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *TypesPriceUnitFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesPriceUnitFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesPriceUnitFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesPriceUnitFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpand

`func (o *TypesPriceUnitFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesPriceUnitFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesPriceUnitFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesPriceUnitFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFilters

`func (o *TypesPriceUnitFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TypesPriceUnitFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TypesPriceUnitFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TypesPriceUnitFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *TypesPriceUnitFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesPriceUnitFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesPriceUnitFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesPriceUnitFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *TypesPriceUnitFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesPriceUnitFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesPriceUnitFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesPriceUnitFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesPriceUnitFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesPriceUnitFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesPriceUnitFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesPriceUnitFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPriceUnitIds

`func (o *TypesPriceUnitFilter) GetPriceUnitIds() []string`

GetPriceUnitIds returns the PriceUnitIds field if non-nil, zero value otherwise.

### GetPriceUnitIdsOk

`func (o *TypesPriceUnitFilter) GetPriceUnitIdsOk() (*[]string, bool)`

GetPriceUnitIdsOk returns a tuple with the PriceUnitIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnitIds

`func (o *TypesPriceUnitFilter) SetPriceUnitIds(v []string)`

SetPriceUnitIds sets PriceUnitIds field to given value.

### HasPriceUnitIds

`func (o *TypesPriceUnitFilter) HasPriceUnitIds() bool`

HasPriceUnitIds returns a boolean if a field has been set.

### GetSort

`func (o *TypesPriceUnitFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesPriceUnitFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesPriceUnitFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesPriceUnitFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesPriceUnitFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesPriceUnitFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesPriceUnitFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesPriceUnitFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesPriceUnitFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesPriceUnitFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesPriceUnitFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesPriceUnitFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


