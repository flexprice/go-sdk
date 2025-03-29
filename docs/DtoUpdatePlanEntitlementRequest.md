# DtoUpdatePlanEntitlementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureId** | **string** |  | 
**FeatureType** | [**TypesFeatureType**](TypesFeatureType.md) |  | 
**Id** | Pointer to **string** | The ID of the entitlement to update (present if the entitlement is being updated) | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsSoftLimit** | Pointer to **bool** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**StaticValue** | Pointer to **string** |  | [optional] 
**UsageLimit** | Pointer to **int32** |  | [optional] 
**UsageResetPeriod** | Pointer to [**TypesBillingPeriod**](TypesBillingPeriod.md) |  | [optional] 

## Methods

### NewDtoUpdatePlanEntitlementRequest

`func NewDtoUpdatePlanEntitlementRequest(featureId string, featureType TypesFeatureType, ) *DtoUpdatePlanEntitlementRequest`

NewDtoUpdatePlanEntitlementRequest instantiates a new DtoUpdatePlanEntitlementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdatePlanEntitlementRequestWithDefaults

`func NewDtoUpdatePlanEntitlementRequestWithDefaults() *DtoUpdatePlanEntitlementRequest`

NewDtoUpdatePlanEntitlementRequestWithDefaults instantiates a new DtoUpdatePlanEntitlementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureId

`func (o *DtoUpdatePlanEntitlementRequest) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *DtoUpdatePlanEntitlementRequest) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *DtoUpdatePlanEntitlementRequest) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.


### GetFeatureType

`func (o *DtoUpdatePlanEntitlementRequest) GetFeatureType() TypesFeatureType`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *DtoUpdatePlanEntitlementRequest) GetFeatureTypeOk() (*TypesFeatureType, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *DtoUpdatePlanEntitlementRequest) SetFeatureType(v TypesFeatureType)`

SetFeatureType sets FeatureType field to given value.


### GetId

`func (o *DtoUpdatePlanEntitlementRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoUpdatePlanEntitlementRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoUpdatePlanEntitlementRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoUpdatePlanEntitlementRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEnabled

`func (o *DtoUpdatePlanEntitlementRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DtoUpdatePlanEntitlementRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DtoUpdatePlanEntitlementRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DtoUpdatePlanEntitlementRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsSoftLimit

`func (o *DtoUpdatePlanEntitlementRequest) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *DtoUpdatePlanEntitlementRequest) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *DtoUpdatePlanEntitlementRequest) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *DtoUpdatePlanEntitlementRequest) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetPlanId

`func (o *DtoUpdatePlanEntitlementRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *DtoUpdatePlanEntitlementRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *DtoUpdatePlanEntitlementRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *DtoUpdatePlanEntitlementRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetStaticValue

`func (o *DtoUpdatePlanEntitlementRequest) GetStaticValue() string`

GetStaticValue returns the StaticValue field if non-nil, zero value otherwise.

### GetStaticValueOk

`func (o *DtoUpdatePlanEntitlementRequest) GetStaticValueOk() (*string, bool)`

GetStaticValueOk returns a tuple with the StaticValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticValue

`func (o *DtoUpdatePlanEntitlementRequest) SetStaticValue(v string)`

SetStaticValue sets StaticValue field to given value.

### HasStaticValue

`func (o *DtoUpdatePlanEntitlementRequest) HasStaticValue() bool`

HasStaticValue returns a boolean if a field has been set.

### GetUsageLimit

`func (o *DtoUpdatePlanEntitlementRequest) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *DtoUpdatePlanEntitlementRequest) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *DtoUpdatePlanEntitlementRequest) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *DtoUpdatePlanEntitlementRequest) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetUsageResetPeriod

`func (o *DtoUpdatePlanEntitlementRequest) GetUsageResetPeriod() TypesBillingPeriod`

GetUsageResetPeriod returns the UsageResetPeriod field if non-nil, zero value otherwise.

### GetUsageResetPeriodOk

`func (o *DtoUpdatePlanEntitlementRequest) GetUsageResetPeriodOk() (*TypesBillingPeriod, bool)`

GetUsageResetPeriodOk returns a tuple with the UsageResetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageResetPeriod

`func (o *DtoUpdatePlanEntitlementRequest) SetUsageResetPeriod(v TypesBillingPeriod)`

SetUsageResetPeriod sets UsageResetPeriod field to given value.

### HasUsageResetPeriod

`func (o *DtoUpdatePlanEntitlementRequest) HasUsageResetPeriod() bool`

HasUsageResetPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


