# PriceunitPriceUnitFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | EnvironmentID filters by specific environment ID | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) | Filters allows complex filtering based on multiple fields | [optional] 
**QueryFilter** | Pointer to [**TypesQueryFilter**](TypesQueryFilter.md) |  | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) | Sort allows sorting by multiple fields | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** | TenantID filters by specific tenant ID | [optional] 
**TimeRangeFilter** | Pointer to [**TypesTimeRangeFilter**](TypesTimeRangeFilter.md) |  | [optional] 

## Methods

### NewPriceunitPriceUnitFilter

`func NewPriceunitPriceUnitFilter() *PriceunitPriceUnitFilter`

NewPriceunitPriceUnitFilter instantiates a new PriceunitPriceUnitFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceunitPriceUnitFilterWithDefaults

`func NewPriceunitPriceUnitFilterWithDefaults() *PriceunitPriceUnitFilter`

NewPriceunitPriceUnitFilterWithDefaults instantiates a new PriceunitPriceUnitFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *PriceunitPriceUnitFilter) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *PriceunitPriceUnitFilter) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *PriceunitPriceUnitFilter) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *PriceunitPriceUnitFilter) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetFilters

`func (o *PriceunitPriceUnitFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PriceunitPriceUnitFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PriceunitPriceUnitFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PriceunitPriceUnitFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetQueryFilter

`func (o *PriceunitPriceUnitFilter) GetQueryFilter() TypesQueryFilter`

GetQueryFilter returns the QueryFilter field if non-nil, zero value otherwise.

### GetQueryFilterOk

`func (o *PriceunitPriceUnitFilter) GetQueryFilterOk() (*TypesQueryFilter, bool)`

GetQueryFilterOk returns a tuple with the QueryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFilter

`func (o *PriceunitPriceUnitFilter) SetQueryFilter(v TypesQueryFilter)`

SetQueryFilter sets QueryFilter field to given value.

### HasQueryFilter

`func (o *PriceunitPriceUnitFilter) HasQueryFilter() bool`

HasQueryFilter returns a boolean if a field has been set.

### GetSort

`func (o *PriceunitPriceUnitFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PriceunitPriceUnitFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PriceunitPriceUnitFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PriceunitPriceUnitFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStatus

`func (o *PriceunitPriceUnitFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PriceunitPriceUnitFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PriceunitPriceUnitFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PriceunitPriceUnitFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *PriceunitPriceUnitFilter) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PriceunitPriceUnitFilter) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PriceunitPriceUnitFilter) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PriceunitPriceUnitFilter) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTimeRangeFilter

`func (o *PriceunitPriceUnitFilter) GetTimeRangeFilter() TypesTimeRangeFilter`

GetTimeRangeFilter returns the TimeRangeFilter field if non-nil, zero value otherwise.

### GetTimeRangeFilterOk

`func (o *PriceunitPriceUnitFilter) GetTimeRangeFilterOk() (*TypesTimeRangeFilter, bool)`

GetTimeRangeFilterOk returns a tuple with the TimeRangeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeFilter

`func (o *PriceunitPriceUnitFilter) SetTimeRangeFilter(v TypesTimeRangeFilter)`

SetTimeRangeFilter sets TimeRangeFilter field to given value.

### HasTimeRangeFilter

`func (o *PriceunitPriceUnitFilter) HasTimeRangeFilter() bool`

HasTimeRangeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


