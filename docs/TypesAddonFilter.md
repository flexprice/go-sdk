# TypesAddonFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonIds** | Pointer to **[]string** |  | [optional] 
**AddonType** | Pointer to [**TypesAddonType**](TypesAddonType.md) |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) | filters allows complex filtering based on multiple fields | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**LookupKeys** | Pointer to **[]string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 

## Methods

### NewTypesAddonFilter

`func NewTypesAddonFilter() *TypesAddonFilter`

NewTypesAddonFilter instantiates a new TypesAddonFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesAddonFilterWithDefaults

`func NewTypesAddonFilterWithDefaults() *TypesAddonFilter`

NewTypesAddonFilterWithDefaults instantiates a new TypesAddonFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonIds

`func (o *TypesAddonFilter) GetAddonIds() []string`

GetAddonIds returns the AddonIds field if non-nil, zero value otherwise.

### GetAddonIdsOk

`func (o *TypesAddonFilter) GetAddonIdsOk() (*[]string, bool)`

GetAddonIdsOk returns a tuple with the AddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonIds

`func (o *TypesAddonFilter) SetAddonIds(v []string)`

SetAddonIds sets AddonIds field to given value.

### HasAddonIds

`func (o *TypesAddonFilter) HasAddonIds() bool`

HasAddonIds returns a boolean if a field has been set.

### GetAddonType

`func (o *TypesAddonFilter) GetAddonType() TypesAddonType`

GetAddonType returns the AddonType field if non-nil, zero value otherwise.

### GetAddonTypeOk

`func (o *TypesAddonFilter) GetAddonTypeOk() (*TypesAddonType, bool)`

GetAddonTypeOk returns a tuple with the AddonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonType

`func (o *TypesAddonFilter) SetAddonType(v TypesAddonType)`

SetAddonType sets AddonType field to given value.

### HasAddonType

`func (o *TypesAddonFilter) HasAddonType() bool`

HasAddonType returns a boolean if a field has been set.

### GetEndTime

`func (o *TypesAddonFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesAddonFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesAddonFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesAddonFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpand

`func (o *TypesAddonFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesAddonFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesAddonFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesAddonFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFilters

`func (o *TypesAddonFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TypesAddonFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TypesAddonFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TypesAddonFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *TypesAddonFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesAddonFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesAddonFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesAddonFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLookupKeys

`func (o *TypesAddonFilter) GetLookupKeys() []string`

GetLookupKeys returns the LookupKeys field if non-nil, zero value otherwise.

### GetLookupKeysOk

`func (o *TypesAddonFilter) GetLookupKeysOk() (*[]string, bool)`

GetLookupKeysOk returns a tuple with the LookupKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKeys

`func (o *TypesAddonFilter) SetLookupKeys(v []string)`

SetLookupKeys sets LookupKeys field to given value.

### HasLookupKeys

`func (o *TypesAddonFilter) HasLookupKeys() bool`

HasLookupKeys returns a boolean if a field has been set.

### GetOffset

`func (o *TypesAddonFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesAddonFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesAddonFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesAddonFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesAddonFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesAddonFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesAddonFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesAddonFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSort

`func (o *TypesAddonFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesAddonFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesAddonFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesAddonFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesAddonFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesAddonFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesAddonFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesAddonFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesAddonFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesAddonFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesAddonFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesAddonFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


