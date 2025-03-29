# DtoAggregatedEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsSoftLimit** | Pointer to **bool** |  | [optional] 
**StaticValues** | Pointer to **[]string** | For static/SLA features | [optional] 
**UsageLimit** | Pointer to **int32** |  | [optional] 
**UsageResetPeriod** | Pointer to [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | [optional] 

## Methods

### NewDtoAggregatedEntitlement

`func NewDtoAggregatedEntitlement() *DtoAggregatedEntitlement`

NewDtoAggregatedEntitlement instantiates a new DtoAggregatedEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoAggregatedEntitlementWithDefaults

`func NewDtoAggregatedEntitlementWithDefaults() *DtoAggregatedEntitlement`

NewDtoAggregatedEntitlementWithDefaults instantiates a new DtoAggregatedEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *DtoAggregatedEntitlement) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DtoAggregatedEntitlement) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DtoAggregatedEntitlement) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DtoAggregatedEntitlement) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsSoftLimit

`func (o *DtoAggregatedEntitlement) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *DtoAggregatedEntitlement) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *DtoAggregatedEntitlement) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *DtoAggregatedEntitlement) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetStaticValues

`func (o *DtoAggregatedEntitlement) GetStaticValues() []string`

GetStaticValues returns the StaticValues field if non-nil, zero value otherwise.

### GetStaticValuesOk

`func (o *DtoAggregatedEntitlement) GetStaticValuesOk() (*[]string, bool)`

GetStaticValuesOk returns a tuple with the StaticValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticValues

`func (o *DtoAggregatedEntitlement) SetStaticValues(v []string)`

SetStaticValues sets StaticValues field to given value.

### HasStaticValues

`func (o *DtoAggregatedEntitlement) HasStaticValues() bool`

HasStaticValues returns a boolean if a field has been set.

### GetUsageLimit

`func (o *DtoAggregatedEntitlement) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *DtoAggregatedEntitlement) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *DtoAggregatedEntitlement) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *DtoAggregatedEntitlement) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetUsageResetPeriod

`func (o *DtoAggregatedEntitlement) GetUsageResetPeriod() TypesBillingPeriod`

GetUsageResetPeriod returns the UsageResetPeriod field if non-nil, zero value otherwise.

### GetUsageResetPeriodOk

`func (o *DtoAggregatedEntitlement) GetUsageResetPeriodOk() (*TypesBillingPeriod, bool)`

GetUsageResetPeriodOk returns a tuple with the UsageResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageResetPeriod

`func (o *DtoAggregatedEntitlement) SetUsageResetPeriod(v TypesBillingPeriod)`

SetUsageResetPeriod sets UsageResetPeriod field to given value.

### HasUsageResetPeriod

`func (o *DtoAggregatedEntitlement) HasUsageResetPeriod() bool`

HasUsageResetPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


