# TypesEntitlementFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** |  | [optional] 
**EntityIds** | Pointer to **[]string** |  | [optional] 
**EntityType** | Pointer to [**TypesEntitlementEntityType**](TypesEntitlementEntityType.md) |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**FeatureIds** | Pointer to **[]string** |  | [optional] 
**FeatureType** | Pointer to [**TypesFeatureType**](TypesFeatureType.md) |  | [optional] 
**Filters** | Pointer to [**[]TypesFilterCondition**](TypesFilterCondition.md) | Specific filters for entitlements | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**PlanIds** | Pointer to **[]string** |  | [optional] 
**Sort** | Pointer to [**[]TypesSortCondition**](TypesSortCondition.md) |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 

## Methods

### NewTypesEntitlementFilter

`func NewTypesEntitlementFilter() *TypesEntitlementFilter`

NewTypesEntitlementFilter instantiates a new TypesEntitlementFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesEntitlementFilterWithDefaults

`func NewTypesEntitlementFilterWithDefaults() *TypesEntitlementFilter`

NewTypesEntitlementFilterWithDefaults instantiates a new TypesEntitlementFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *TypesEntitlementFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesEntitlementFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesEntitlementFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesEntitlementFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntityIds

`func (o *TypesEntitlementFilter) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *TypesEntitlementFilter) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *TypesEntitlementFilter) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *TypesEntitlementFilter) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### GetEntityType

`func (o *TypesEntitlementFilter) GetEntityType() TypesEntitlementEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TypesEntitlementFilter) GetEntityTypeOk() (*TypesEntitlementEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TypesEntitlementFilter) SetEntityType(v TypesEntitlementEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *TypesEntitlementFilter) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetExpand

`func (o *TypesEntitlementFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesEntitlementFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesEntitlementFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesEntitlementFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetFeatureIds

`func (o *TypesEntitlementFilter) GetFeatureIds() []string`

GetFeatureIds returns the FeatureIds field if non-nil, zero value otherwise.

### GetFeatureIdsOk

`func (o *TypesEntitlementFilter) GetFeatureIdsOk() (*[]string, bool)`

GetFeatureIdsOk returns a tuple with the FeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIds

`func (o *TypesEntitlementFilter) SetFeatureIds(v []string)`

SetFeatureIds sets FeatureIds field to given value.

### HasFeatureIds

`func (o *TypesEntitlementFilter) HasFeatureIds() bool`

HasFeatureIds returns a boolean if a field has been set.

### GetFeatureType

`func (o *TypesEntitlementFilter) GetFeatureType() TypesFeatureType`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *TypesEntitlementFilter) GetFeatureTypeOk() (*TypesFeatureType, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *TypesEntitlementFilter) SetFeatureType(v TypesFeatureType)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *TypesEntitlementFilter) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.

### GetFilters

`func (o *TypesEntitlementFilter) GetFilters() []TypesFilterCondition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TypesEntitlementFilter) GetFiltersOk() (*[]TypesFilterCondition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TypesEntitlementFilter) SetFilters(v []TypesFilterCondition)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TypesEntitlementFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetIsEnabled

`func (o *TypesEntitlementFilter) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *TypesEntitlementFilter) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *TypesEntitlementFilter) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *TypesEntitlementFilter) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetLimit

`func (o *TypesEntitlementFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesEntitlementFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesEntitlementFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesEntitlementFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *TypesEntitlementFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesEntitlementFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesEntitlementFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesEntitlementFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesEntitlementFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesEntitlementFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesEntitlementFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesEntitlementFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPlanIds

`func (o *TypesEntitlementFilter) GetPlanIds() []string`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *TypesEntitlementFilter) GetPlanIdsOk() (*[]string, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *TypesEntitlementFilter) SetPlanIds(v []string)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *TypesEntitlementFilter) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetSort

`func (o *TypesEntitlementFilter) GetSort() []TypesSortCondition`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesEntitlementFilter) GetSortOk() (*[]TypesSortCondition, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesEntitlementFilter) SetSort(v []TypesSortCondition)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesEntitlementFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesEntitlementFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesEntitlementFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesEntitlementFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesEntitlementFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesEntitlementFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesEntitlementFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesEntitlementFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesEntitlementFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


