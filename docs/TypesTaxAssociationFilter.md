# TypesTaxAssociationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApply** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to [**TypesTaxRateEntityType**](TypesTaxRateEntityType.md) |  | [optional] 
**Expand** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**Sort** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TaxAssociationIds** | Pointer to **[]string** |  | [optional] 
**TaxRateIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTypesTaxAssociationFilter

`func NewTypesTaxAssociationFilter() *TypesTaxAssociationFilter`

NewTypesTaxAssociationFilter instantiates a new TypesTaxAssociationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesTaxAssociationFilterWithDefaults

`func NewTypesTaxAssociationFilterWithDefaults() *TypesTaxAssociationFilter`

NewTypesTaxAssociationFilterWithDefaults instantiates a new TypesTaxAssociationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApply

`func (o *TypesTaxAssociationFilter) GetAutoApply() bool`

GetAutoApply returns the AutoApply field if non-nil, zero value otherwise.

### GetAutoApplyOk

`func (o *TypesTaxAssociationFilter) GetAutoApplyOk() (*bool, bool)`

GetAutoApplyOk returns a tuple with the AutoApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApply

`func (o *TypesTaxAssociationFilter) SetAutoApply(v bool)`

SetAutoApply sets AutoApply field to given value.

### HasAutoApply

`func (o *TypesTaxAssociationFilter) HasAutoApply() bool`

HasAutoApply returns a boolean if a field has been set.

### GetCurrency

`func (o *TypesTaxAssociationFilter) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TypesTaxAssociationFilter) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TypesTaxAssociationFilter) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TypesTaxAssociationFilter) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEndTime

`func (o *TypesTaxAssociationFilter) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TypesTaxAssociationFilter) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TypesTaxAssociationFilter) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TypesTaxAssociationFilter) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntityId

`func (o *TypesTaxAssociationFilter) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *TypesTaxAssociationFilter) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *TypesTaxAssociationFilter) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *TypesTaxAssociationFilter) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *TypesTaxAssociationFilter) GetEntityType() TypesTaxRateEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TypesTaxAssociationFilter) GetEntityTypeOk() (*TypesTaxRateEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TypesTaxAssociationFilter) SetEntityType(v TypesTaxRateEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *TypesTaxAssociationFilter) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetExpand

`func (o *TypesTaxAssociationFilter) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *TypesTaxAssociationFilter) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *TypesTaxAssociationFilter) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *TypesTaxAssociationFilter) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetLimit

`func (o *TypesTaxAssociationFilter) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TypesTaxAssociationFilter) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TypesTaxAssociationFilter) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TypesTaxAssociationFilter) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *TypesTaxAssociationFilter) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TypesTaxAssociationFilter) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TypesTaxAssociationFilter) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TypesTaxAssociationFilter) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *TypesTaxAssociationFilter) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *TypesTaxAssociationFilter) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *TypesTaxAssociationFilter) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *TypesTaxAssociationFilter) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSort

`func (o *TypesTaxAssociationFilter) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TypesTaxAssociationFilter) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TypesTaxAssociationFilter) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TypesTaxAssociationFilter) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetStartTime

`func (o *TypesTaxAssociationFilter) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TypesTaxAssociationFilter) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TypesTaxAssociationFilter) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TypesTaxAssociationFilter) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *TypesTaxAssociationFilter) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TypesTaxAssociationFilter) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TypesTaxAssociationFilter) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TypesTaxAssociationFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxAssociationIds

`func (o *TypesTaxAssociationFilter) GetTaxAssociationIds() []string`

GetTaxAssociationIds returns the TaxAssociationIds field if non-nil, zero value otherwise.

### GetTaxAssociationIdsOk

`func (o *TypesTaxAssociationFilter) GetTaxAssociationIdsOk() (*[]string, bool)`

GetTaxAssociationIdsOk returns a tuple with the TaxAssociationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAssociationIds

`func (o *TypesTaxAssociationFilter) SetTaxAssociationIds(v []string)`

SetTaxAssociationIds sets TaxAssociationIds field to given value.

### HasTaxAssociationIds

`func (o *TypesTaxAssociationFilter) HasTaxAssociationIds() bool`

HasTaxAssociationIds returns a boolean if a field has been set.

### GetTaxRateIds

`func (o *TypesTaxAssociationFilter) GetTaxRateIds() []string`

GetTaxRateIds returns the TaxRateIds field if non-nil, zero value otherwise.

### GetTaxRateIdsOk

`func (o *TypesTaxAssociationFilter) GetTaxRateIdsOk() (*[]string, bool)`

GetTaxRateIdsOk returns a tuple with the TaxRateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateIds

`func (o *TypesTaxAssociationFilter) SetTaxRateIds(v []string)`

SetTaxRateIds sets TaxRateIds field to given value.

### HasTaxRateIds

`func (o *TypesTaxAssociationFilter) HasTaxRateIds() bool`

HasTaxRateIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


