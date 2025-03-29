# DtoUpdateEntitlementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsSoftLimit** | Pointer to **bool** |  | [optional] 
**StaticValue** | Pointer to **string** |  | [optional] 
**UsageLimit** | Pointer to **int32** |  | [optional] 
**UsageResetPeriod** | Pointer to [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | [optional] 

## Methods

### NewDtoUpdateEntitlementRequest

`func NewDtoUpdateEntitlementRequest() *DtoUpdateEntitlementRequest`

NewDtoUpdateEntitlementRequest instantiates a new DtoUpdateEntitlementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdateEntitlementRequestWithDefaults

`func NewDtoUpdateEntitlementRequestWithDefaults() *DtoUpdateEntitlementRequest`

NewDtoUpdateEntitlementRequestWithDefaults instantiates a new DtoUpdateEntitlementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *DtoUpdateEntitlementRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DtoUpdateEntitlementRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DtoUpdateEntitlementRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DtoUpdateEntitlementRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsSoftLimit

`func (o *DtoUpdateEntitlementRequest) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *DtoUpdateEntitlementRequest) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *DtoUpdateEntitlementRequest) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *DtoUpdateEntitlementRequest) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetStaticValue

`func (o *DtoUpdateEntitlementRequest) GetStaticValue() string`

GetStaticValue returns the StaticValue field if non-nil, zero value otherwise.

### GetStaticValueOk

`func (o *DtoUpdateEntitlementRequest) GetStaticValueOk() (*string, bool)`

GetStaticValueOk returns a tuple with the StaticValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticValue

`func (o *DtoUpdateEntitlementRequest) SetStaticValue(v string)`

SetStaticValue sets StaticValue field to given value.

### HasStaticValue

`func (o *DtoUpdateEntitlementRequest) HasStaticValue() bool`

HasStaticValue returns a boolean if a field has been set.

### GetUsageLimit

`func (o *DtoUpdateEntitlementRequest) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *DtoUpdateEntitlementRequest) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *DtoUpdateEntitlementRequest) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *DtoUpdateEntitlementRequest) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetUsageResetPeriod

`func (o *DtoUpdateEntitlementRequest) GetUsageResetPeriod() TypesBillingPeriod`

GetUsageResetPeriod returns the UsageResetPeriod field if non-nil, zero value otherwise.

### GetUsageResetPeriodOk

`func (o *DtoUpdateEntitlementRequest) GetUsageResetPeriodOk() (*TypesBillingPeriod, bool)`

GetUsageResetPeriodOk returns a tuple with the UsageResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageResetPeriod

`func (o *DtoUpdateEntitlementRequest) SetUsageResetPeriod(v TypesBillingPeriod)`

SetUsageResetPeriod sets UsageResetPeriod field to given value.

### HasUsageResetPeriod

`func (o *DtoUpdateEntitlementRequest) HasUsageResetPeriod() bool`

HasUsageResetPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


