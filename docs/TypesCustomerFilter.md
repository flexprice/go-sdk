# TypesCustomerFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerIds** | Pointer to **[]string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalIds** | Pointer to **[]string** |  | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**ParentCustomerIds** | Pointer to **[]string** |  | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 

## Methods

### NewTypesCustomerFilter

`func NewTypesCustomerFilter() *TypesCustomerFilter`

NewTypesCustomerFilter instantiates a new TypesCustomerFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesCustomerFilterWithDefaults

`func NewTypesCustomerFilterWithDefaults() *TypesCustomerFilter`

NewTypesCustomerFilterWithDefaults instantiates a new TypesCustomerFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerIds

`func (o *TypesCustomerFilter) GetCustomerIds() []string`

GetCustomerIds returns the CustomerIds field if non-nil, zero value otherwise.

### GetCustomerIdsOk

`func (o *TypesCustomerFilter) GetCustomerIdsOk() (*[]string, bool)`

GetCustomerIdsOk returns a tuple with the CustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIds

`func (o *TypesCustomerFilter) SetCustomerIds(v []string)`

SetCustomerIds sets CustomerIds field to given value.

### HasCustomerIds

`func (o *TypesCustomerFilter) HasCustomerIds() bool`

HasCustomerIds returns a boolean if a field has been set.

### GetEmail

`func (o *TypesCustomerFilter) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TypesCustomerFilter) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TypesCustomerFilter) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TypesCustomerFilter) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEndTime

`func (o *TypesCustomerFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesCustomerFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesCustomerFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesCustomerFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExpand

`func (o *TypesCustomerFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesCustomerFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesCustomerFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesCustomerFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetExternalId

`func (o *TypesCustomerFilter) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TypesCustomerFilter) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TypesCustomerFilter) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *TypesCustomerFilter) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalIds

`func (o *TypesCustomerFilter) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *TypesCustomerFilter) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *TypesCustomerFilter) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *TypesCustomerFilter) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetFilters

`func (o *TypesCustomerFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TypesCustomerFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TypesCustomerFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TypesCustomerFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *TypesCustomerFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesCustomerFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesCustomerFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesCustomerFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *TypesCustomerFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesCustomerFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesCustomerFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesCustomerFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesCustomerFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesCustomerFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesCustomerFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesCustomerFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetParentCustomerIds

`func (o *TypesCustomerFilter) GetParentCustomerIds() []string`

GetParentCustomerIds returns the ParentCustomerIds field if non-nil, zero value otherwise.

### GetParentCustomerIdsOk

`func (o *TypesCustomerFilter) GetParentCustomerIdsOk() (*[]string, bool)`

GetParentCustomerIdsOk returns a tuple with the ParentCustomerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCustomerIds

`func (o *TypesCustomerFilter) SetParentCustomerIds(v []string)`

SetParentCustomerIds sets ParentCustomerIds field to given value.

### HasParentCustomerIds

`func (o *TypesCustomerFilter) HasParentCustomerIds() bool`

HasParentCustomerIds returns a boolean if a field has been set.

### GetSort

`func (o *TypesCustomerFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesCustomerFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesCustomerFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesCustomerFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesCustomerFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesCustomerFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesCustomerFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesCustomerFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesCustomerFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesCustomerFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesCustomerFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesCustomerFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


