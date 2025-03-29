# DtoCreateEntitlementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureId** | **string** |  | 
**FeatureType** | [**TypesFeatureType**](TypesFeatureType.md) |  | 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsSoftLimit** | Pointer to **bool** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**StaticValue** | Pointer to **string** |  | [optional] 
**UsageLimit** | Pointer to **int32** |  | [optional] 
**UsageResetPeriod** | Pointer to [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | [optional] 

## Methods

### NewDtoCreateEntitlementRequest

`func NewDtoCreateEntitlementRequest(featureId string, featureType TypesFeatureType, ) *DtoCreateEntitlementRequest`

NewDtoCreateEntitlementRequest instantiates a new DtoCreateEntitlementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateEntitlementRequestWithDefaults

`func NewDtoCreateEntitlementRequestWithDefaults() *DtoCreateEntitlementRequest`

NewDtoCreateEntitlementRequestWithDefaults instantiates a new DtoCreateEntitlementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureId

`func (o *DtoCreateEntitlementRequest) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *DtoCreateEntitlementRequest) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *DtoCreateEntitlementRequest) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.


### GetFeatureType

`func (o *DtoCreateEntitlementRequest) GetFeatureType() TypesFeatureType`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *DtoCreateEntitlementRequest) GetFeatureTypeOk() (*TypesFeatureType, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *DtoCreateEntitlementRequest) SetFeatureType(v TypesFeatureType)`

SetFeatureType sets FeatureType field to given value.


### GetIsEnabled

`func (o *DtoCreateEntitlementRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DtoCreateEntitlementRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DtoCreateEntitlementRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DtoCreateEntitlementRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsSoftLimit

`func (o *DtoCreateEntitlementRequest) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *DtoCreateEntitlementRequest) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *DtoCreateEntitlementRequest) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *DtoCreateEntitlementRequest) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoCreateEntitlementRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoCreateEntitlementRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoCreateEntitlementRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoCreateEntitlementRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetStaticValue

`func (o *DtoCreateEntitlementRequest) GetStaticValue() string`

GetStaticValue returns the StaticValue field if non-nil, zero value otherwise.

### GetStaticValueOk

`func (o *DtoCreateEntitlementRequest) GetStaticValueOk() (*string, bool)`

GetStaticValueOk returns a tuple with the StaticValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticValue

`func (o *DtoCreateEntitlementRequest) SetStaticValue(v string)`

SetStaticValue sets StaticValue field to given value.

### HasStaticValue

`func (o *DtoCreateEntitlementRequest) HasStaticValue() bool`

HasStaticValue returns a boolean if a field has been set.

### GetUsageLimit

`func (o *DtoCreateEntitlementRequest) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *DtoCreateEntitlementRequest) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *DtoCreateEntitlementRequest) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *DtoCreateEntitlementRequest) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetUsageResetPeriod

`func (o *DtoCreateEntitlementRequest) GetUsageResetPeriod() TypesBillingPeriod`

GetUsageResetPeriod returns the UsageResetPeriod field if non-nil, zero value otherwise.

### GetUsageResetPeriodOk

`func (o *DtoCreateEntitlementRequest) GetUsageResetPeriodOk() (*TypesBillingPeriod, bool)`

GetUsageResetPeriodOk returns a tuple with the UsageResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageResetPeriod

`func (o *DtoCreateEntitlementRequest) SetUsageResetPeriod(v TypesBillingPeriod)`

SetUsageResetPeriod sets UsageResetPeriod field to given value.

### HasUsageResetPeriod

`func (o *DtoCreateEntitlementRequest) HasUsageResetPeriod() bool`

HasUsageResetPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


