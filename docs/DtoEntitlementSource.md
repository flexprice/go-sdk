# DtoEntitlementSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementId** | Pointer to **string** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityName** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to [**DtoEntitlementSourceEntityType**](DtoEntitlementSourceEntityType.md) |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**StaticValue** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**UsageLimit** | Pointer to **int32** |  | [optional] 
**UsageResetPeriod** | Pointer to [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | [optional] 

## Methods

### NewDtoEntitlementSource

`func NewDtoEntitlementSource() *DtoEntitlementSource`

NewDtoEntitlementSource instantiates a new DtoEntitlementSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoEntitlementSourceWithDefaults

`func NewDtoEntitlementSourceWithDefaults() *DtoEntitlementSource`

NewDtoEntitlementSourceWithDefaults instantiates a new DtoEntitlementSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementId

`func (o *DtoEntitlementSource) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *DtoEntitlementSource) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *DtoEntitlementSource) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.

### HasEntitlementId

`func (o *DtoEntitlementSource) HasEntitlementId() bool`

HasEntitlementId returns a boolean if a field has been set.

### GetEntityId

`func (o *DtoEntitlementSource) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DtoEntitlementSource) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DtoEntitlementSource) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *DtoEntitlementSource) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityName

`func (o *DtoEntitlementSource) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *DtoEntitlementSource) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *DtoEntitlementSource) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.

### HasEntityName

`func (o *DtoEntitlementSource) HasEntityName() bool`

HasEntityName returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoEntitlementSource) GetEntityType() DtoEntitlementSourceEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoEntitlementSource) GetEntityTypeOk() (*DtoEntitlementSourceEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoEntitlementSource) SetEntityType(v DtoEntitlementSourceEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DtoEntitlementSource) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetIsEnabled

`func (o *DtoEntitlementSource) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DtoEntitlementSource) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DtoEntitlementSource) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DtoEntitlementSource) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetQuantity

`func (o *DtoEntitlementSource) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *DtoEntitlementSource) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *DtoEntitlementSource) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *DtoEntitlementSource) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStaticValue

`func (o *DtoEntitlementSource) GetStaticValue() string`

GetStaticValue returns the StaticValue field if non-nil, zero value otherwise.

### GetStaticValueOk

`func (o *DtoEntitlementSource) GetStaticValueOk() (*string, bool)`

GetStaticValueOk returns a tuple with the StaticValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticValue

`func (o *DtoEntitlementSource) SetStaticValue(v string)`

SetStaticValue sets StaticValue field to given value.

### HasStaticValue

`func (o *DtoEntitlementSource) HasStaticValue() bool`

HasStaticValue returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoEntitlementSource) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoEntitlementSource) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoEntitlementSource) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoEntitlementSource) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUsageLimit

`func (o *DtoEntitlementSource) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *DtoEntitlementSource) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *DtoEntitlementSource) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *DtoEntitlementSource) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetUsageResetPeriod

`func (o *DtoEntitlementSource) GetUsageResetPeriod() TypesBillingPeriod`

GetUsageResetPeriod returns the UsageResetPeriod field if non-nil, zero value otherwise.

### GetUsageResetPeriodOk

`func (o *DtoEntitlementSource) GetUsageResetPeriodOk() (*TypesBillingPeriod, bool)`

GetUsageResetPeriodOk returns a tuple with the UsageResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageResetPeriod

`func (o *DtoEntitlementSource) SetUsageResetPeriod(v TypesBillingPeriod)`

SetUsageResetPeriod sets UsageResetPeriod field to given value.

### HasUsageResetPeriod

`func (o *DtoEntitlementSource) HasUsageResetPeriod() bool`

HasUsageResetPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


