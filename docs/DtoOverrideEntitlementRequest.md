# DtoOverrideEntitlementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementId** | **string** | EntitlementID references the plan/addon entitlement to override | 
**IsEnabled** | Pointer to **bool** | IsEnabled determines if the entitlement is enabled or disabled | [optional] 
**StaticValue** | Pointer to **string** | StaticValue is the static value for static features | [optional] 
**UsageLimit** | Pointer to **int32** | UsageLimit is the new usage limit (only these 3 fields can be overridden) For metered features, nil means unlimited usage | [optional] 

## Methods

### NewDtoOverrideEntitlementRequest

`func NewDtoOverrideEntitlementRequest(entitlementId string, ) *DtoOverrideEntitlementRequest`

NewDtoOverrideEntitlementRequest instantiates a new DtoOverrideEntitlementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoOverrideEntitlementRequestWithDefaults

`func NewDtoOverrideEntitlementRequestWithDefaults() *DtoOverrideEntitlementRequest`

NewDtoOverrideEntitlementRequestWithDefaults instantiates a new DtoOverrideEntitlementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementId

`func (o *DtoOverrideEntitlementRequest) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *DtoOverrideEntitlementRequest) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *DtoOverrideEntitlementRequest) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetIsEnabled

`func (o *DtoOverrideEntitlementRequest) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DtoOverrideEntitlementRequest) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DtoOverrideEntitlementRequest) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DtoOverrideEntitlementRequest) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetStaticValue

`func (o *DtoOverrideEntitlementRequest) GetStaticValue() string`

GetStaticValue returns the StaticValue field if non-nil, zero value otherwise.

### GetStaticValueOk

`func (o *DtoOverrideEntitlementRequest) GetStaticValueOk() (*string, bool)`

GetStaticValueOk returns a tuple with the StaticValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticValue

`func (o *DtoOverrideEntitlementRequest) SetStaticValue(v string)`

SetStaticValue sets StaticValue field to given value.

### HasStaticValue

`func (o *DtoOverrideEntitlementRequest) HasStaticValue() bool`

HasStaticValue returns a boolean if a field has been set.

### GetUsageLimit

`func (o *DtoOverrideEntitlementRequest) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *DtoOverrideEntitlementRequest) GetUsageLimitOk() (*int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *DtoOverrideEntitlementRequest) SetUsageLimit(v int32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *DtoOverrideEntitlementRequest) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


