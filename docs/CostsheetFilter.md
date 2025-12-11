# CostsheetFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostsheetIDs** | Pointer to **[]string** | CostsheetIDs allows filtering by specific costsheet IDs | [optional] 
**EnvironmentID** | Pointer to **string** | EnvironmentID filters by specific environment ID | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) | Filters contains custom filtering conditions | [optional] 
**LookupKey** | Pointer to **string** | LookupKey filters by lookup key | [optional] 
**Name** | Pointer to **string** | Name filters by costsheet name | [optional] 
**QueryFilter** | Pointer to [**TypesQueryFilter**](TypesQueryFilter.md) |  | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) | Sort specifies result ordering preferences | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TenantID** | Pointer to **string** | TenantID filters by specific tenant ID | [optional] 
**TimeRangeFilter** | Pointer to [**TypesTimeRangeFilter**](TypesTimeRangeFilter.md) |  | [optional] 

## Methods

### NewCostsheetFilter

`func NewCostsheetFilter() *CostsheetFilter`

NewCostsheetFilter instantiates a new CostsheetFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostsheetFilterWithDefaults

`func NewCostsheetFilterWithDefaults() *CostsheetFilter`

NewCostsheetFilterWithDefaults instantiates a new CostsheetFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostsheetIDs

`func (o *CostsheetFilter) GetCostsheetIDs() []string`

GetCostsheetIDs returns the CostsheetIDs field if non-nil, zero value otherwise.

### GetCostsheetIDsOk

`func (o *CostsheetFilter) GetCostsheetIDsOk() (*[]string, bool)`

GetCostsheetIDsOk returns a tuple with the CostsheetIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostsheetIDs

`func (o *CostsheetFilter) SetCostsheetIDs(v []string)`

SetCostsheetIDs sets CostsheetIDs field to given value.

### HasCostsheetIDs

`func (o *CostsheetFilter) HasCostsheetIDs() bool`

HasCostsheetIDs returns a boolean if a field has been set.

### GetEnvironmentID

`func (o *CostsheetFilter) GetEnvironmentID() string`

GetEnvironmentID returns the EnvironmentID field if non-nil, zero value otherwise.

### GetEnvironmentIDOk

`func (o *CostsheetFilter) GetEnvironmentIDOk() (*string, bool)`

GetEnvironmentIDOk returns a tuple with the EnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentID

`func (o *CostsheetFilter) SetEnvironmentID(v string)`

SetEnvironmentID sets EnvironmentID field to given value.

### HasEnvironmentID

`func (o *CostsheetFilter) HasEnvironmentID() bool`

HasEnvironmentID returns a boolean if a field has been set.

### GetFilters

`func (o *CostsheetFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CostsheetFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CostsheetFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CostsheetFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLookupKey

`func (o *CostsheetFilter) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *CostsheetFilter) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *CostsheetFilter) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *CostsheetFilter) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetName

`func (o *CostsheetFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CostsheetFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CostsheetFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CostsheetFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueryFilter

`func (o *CostsheetFilter) GetQueryFilter() TypesQueryFilter`

GetQueryFilter returns the QueryFilter field if non-nil, zero value otherwise.

### GetQueryFilterOk

`func (o *CostsheetFilter) GetQueryFilterOk() (*TypesQueryFilter, bool)`

GetQueryFilterOk returns a tuple with the QueryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFilter

`func (o *CostsheetFilter) SetQueryFilter(v TypesQueryFilter)`

SetQueryFilter sets QueryFilter field to given value.

### HasQueryFilter

`func (o *CostsheetFilter) HasQueryFilter() bool`

HasQueryFilter returns a boolean if a field has been set.

### GetSort

`func (o *CostsheetFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CostsheetFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CostsheetFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *CostsheetFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStatus

`func (o *CostsheetFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CostsheetFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CostsheetFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CostsheetFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantID

`func (o *CostsheetFilter) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *CostsheetFilter) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *CostsheetFilter) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *CostsheetFilter) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.

### GetTimeRangeFilter

`func (o *CostsheetFilter) GetTimeRangeFilter() TypesTimeRangeFilter`

GetTimeRangeFilter returns the TimeRangeFilter field if non-nil, zero value otherwise.

### GetTimeRangeFilterOk

`func (o *CostsheetFilter) GetTimeRangeFilterOk() (*TypesTimeRangeFilter, bool)`

GetTimeRangeFilterOk returns a tuple with the TimeRangeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeFilter

`func (o *CostsheetFilter) SetTimeRangeFilter(v TypesTimeRangeFilter)`

SetTimeRangeFilter sets TimeRangeFilter field to given value.

### HasTimeRangeFilter

`func (o *CostsheetFilter) HasTimeRangeFilter() bool`

HasTimeRangeFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


