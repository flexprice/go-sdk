# DtoUsageAnalyticItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationType** | Pointer to [**TypesAggregationType**](TypesAggregationType.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EventName** | Pointer to **string** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Points** | Pointer to [**[]DtoUsageAnalyticPoint**](DtoUsageAnalyticPoint.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**TotalCost** | Pointer to **float32** |  | [optional] 
**TotalUsage** | Pointer to **float32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**UnitPlural** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoUsageAnalyticItem

`func NewDtoUsageAnalyticItem() *DtoUsageAnalyticItem`

NewDtoUsageAnalyticItem instantiates a new DtoUsageAnalyticItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUsageAnalyticItemWithDefaults

`func NewDtoUsageAnalyticItemWithDefaults() *DtoUsageAnalyticItem`

NewDtoUsageAnalyticItemWithDefaults instantiates a new DtoUsageAnalyticItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationType

`func (o *DtoUsageAnalyticItem) GetAggregationType() TypesAggregationType`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *DtoUsageAnalyticItem) GetAggregationTypeOk() (*TypesAggregationType, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *DtoUsageAnalyticItem) SetAggregationType(v TypesAggregationType)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *DtoUsageAnalyticItem) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoUsageAnalyticItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoUsageAnalyticItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoUsageAnalyticItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoUsageAnalyticItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEventName

`func (o *DtoUsageAnalyticItem) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DtoUsageAnalyticItem) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DtoUsageAnalyticItem) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *DtoUsageAnalyticItem) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetFeatureId

`func (o *DtoUsageAnalyticItem) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *DtoUsageAnalyticItem) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *DtoUsageAnalyticItem) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *DtoUsageAnalyticItem) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetName

`func (o *DtoUsageAnalyticItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoUsageAnalyticItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoUsageAnalyticItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoUsageAnalyticItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoints

`func (o *DtoUsageAnalyticItem) GetPoints() []DtoUsageAnalyticPoint`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *DtoUsageAnalyticItem) GetPointsOk() (*[]DtoUsageAnalyticPoint, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *DtoUsageAnalyticItem) SetPoints(v []DtoUsageAnalyticPoint)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *DtoUsageAnalyticItem) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetSource

`func (o *DtoUsageAnalyticItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DtoUsageAnalyticItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DtoUsageAnalyticItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DtoUsageAnalyticItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTotalCost

`func (o *DtoUsageAnalyticItem) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *DtoUsageAnalyticItem) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *DtoUsageAnalyticItem) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *DtoUsageAnalyticItem) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalUsage

`func (o *DtoUsageAnalyticItem) GetTotalUsage() float32`

GetTotalUsage returns the TotalUsage field if non-nil, zero value otherwise.

### GetTotalUsageOk

`func (o *DtoUsageAnalyticItem) GetTotalUsageOk() (*float32, bool)`

GetTotalUsageOk returns a tuple with the TotalUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsage

`func (o *DtoUsageAnalyticItem) SetTotalUsage(v float32)`

SetTotalUsage sets TotalUsage field to given value.

### HasTotalUsage

`func (o *DtoUsageAnalyticItem) HasTotalUsage() bool`

HasTotalUsage returns a boolean if a field has been set.

### GetUnit

`func (o *DtoUsageAnalyticItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DtoUsageAnalyticItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DtoUsageAnalyticItem) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *DtoUsageAnalyticItem) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPlural

`func (o *DtoUsageAnalyticItem) GetUnitPlural() string`

GetUnitPlural returns the UnitPlural field if non-nil, zero value otherwise.

### GetUnitPluralOk

`func (o *DtoUsageAnalyticItem) GetUnitPluralOk() (*string, bool)`

GetUnitPluralOk returns a tuple with the UnitPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPlural

`func (o *DtoUsageAnalyticItem) SetUnitPlural(v string)`

SetUnitPlural sets UnitPlural field to given value.

### HasUnitPlural

`func (o *DtoUsageAnalyticItem) HasUnitPlural() bool`

HasUnitPlural returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


