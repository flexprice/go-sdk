# DtoFeatureUsageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUsage** | Pointer to **string** |  | [optional] 
**Feature** | Pointer to [**DtoFeatureResponse**](DtoFeatureResponse.md) |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**IsSoftLimit** | Pointer to **bool** |  | [optional] 
**IsUnlimited** | Pointer to **bool** |  | [optional] 
**NextUsageResetAt** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to [**[]DtoEntitlementSource**](DtoEntitlementSource.md) |  | [optional] 
**TotalLimit** | Pointer to **int32** |  | [optional] 
**UsagePercent** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoFeatureUsageSummary

`func NewDtoFeatureUsageSummary() *DtoFeatureUsageSummary`

NewDtoFeatureUsageSummary instantiates a new DtoFeatureUsageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoFeatureUsageSummaryWithDefaults

`func NewDtoFeatureUsageSummaryWithDefaults() *DtoFeatureUsageSummary`

NewDtoFeatureUsageSummaryWithDefaults instantiates a new DtoFeatureUsageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUsage

`func (o *DtoFeatureUsageSummary) GetCurrentUsage() string`

GetCurrentUsage returns the CurrentUsage field if non-nil, zero value otherwise.

### GetCurrentUsageOk

`func (o *DtoFeatureUsageSummary) GetCurrentUsageOk() (*string, bool)`

GetCurrentUsageOk returns a tuple with the CurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsage

`func (o *DtoFeatureUsageSummary) SetCurrentUsage(v string)`

SetCurrentUsage sets CurrentUsage field to given value.

### HasCurrentUsage

`func (o *DtoFeatureUsageSummary) HasCurrentUsage() bool`

HasCurrentUsage returns a boolean if a field has been set.

### GetFeature

`func (o *DtoFeatureUsageSummary) GetFeature() DtoFeatureResponse`

GetFeature returns the Feature field if non-nil, zero value otherwise.

### GetFeatureOk

`func (o *DtoFeatureUsageSummary) GetFeatureOk() (*DtoFeatureResponse, bool)`

GetFeatureOk returns a tuple with the Feature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeature

`func (o *DtoFeatureUsageSummary) SetFeature(v DtoFeatureResponse)`

SetFeature sets Feature field to given value.

### HasFeature

`func (o *DtoFeatureUsageSummary) HasFeature() bool`

HasFeature returns a boolean if a field has been set.

### GetIsEnabled

`func (o *DtoFeatureUsageSummary) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *DtoFeatureUsageSummary) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *DtoFeatureUsageSummary) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *DtoFeatureUsageSummary) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsSoftLimit

`func (o *DtoFeatureUsageSummary) GetIsSoftLimit() bool`

GetIsSoftLimit returns the IsSoftLimit field if non-nil, zero value otherwise.

### GetIsSoftLimitOk

`func (o *DtoFeatureUsageSummary) GetIsSoftLimitOk() (*bool, bool)`

GetIsSoftLimitOk returns a tuple with the IsSoftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftLimit

`func (o *DtoFeatureUsageSummary) SetIsSoftLimit(v bool)`

SetIsSoftLimit sets IsSoftLimit field to given value.

### HasIsSoftLimit

`func (o *DtoFeatureUsageSummary) HasIsSoftLimit() bool`

HasIsSoftLimit returns a boolean if a field has been set.

### GetIsUnlimited

`func (o *DtoFeatureUsageSummary) GetIsUnlimited() bool`

GetIsUnlimited returns the IsUnlimited field if non-nil, zero value otherwise.

### GetIsUnlimitedOk

`func (o *DtoFeatureUsageSummary) GetIsUnlimitedOk() (*bool, bool)`

GetIsUnlimitedOk returns a tuple with the IsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnlimited

`func (o *DtoFeatureUsageSummary) SetIsUnlimited(v bool)`

SetIsUnlimited sets IsUnlimited field to given value.

### HasIsUnlimited

`func (o *DtoFeatureUsageSummary) HasIsUnlimited() bool`

HasIsUnlimited returns a boolean if a field has been set.

### GetNextUsageResetAt

`func (o *DtoFeatureUsageSummary) GetNextUsageResetAt() string`

GetNextUsageResetAt returns the NextUsageResetAt field if non-nil, zero value otherwise.

### GetNextUsageResetAtOk

`func (o *DtoFeatureUsageSummary) GetNextUsageResetAtOk() (*string, bool)`

GetNextUsageResetAtOk returns a tuple with the NextUsageResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUsageResetAt

`func (o *DtoFeatureUsageSummary) SetNextUsageResetAt(v string)`

SetNextUsageResetAt sets NextUsageResetAt field to given value.

### HasNextUsageResetAt

`func (o *DtoFeatureUsageSummary) HasNextUsageResetAt() bool`

HasNextUsageResetAt returns a boolean if a field has been set.

### GetSources

`func (o *DtoFeatureUsageSummary) GetSources() []DtoEntitlementSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *DtoFeatureUsageSummary) GetSourcesOk() (*[]DtoEntitlementSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *DtoFeatureUsageSummary) SetSources(v []DtoEntitlementSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *DtoFeatureUsageSummary) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTotalLimit

`func (o *DtoFeatureUsageSummary) GetTotalLimit() int32`

GetTotalLimit returns the TotalLimit field if non-nil, zero value otherwise.

### GetTotalLimitOk

`func (o *DtoFeatureUsageSummary) GetTotalLimitOk() (*int32, bool)`

GetTotalLimitOk returns a tuple with the TotalLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLimit

`func (o *DtoFeatureUsageSummary) SetTotalLimit(v int32)`

SetTotalLimit sets TotalLimit field to given value.

### HasTotalLimit

`func (o *DtoFeatureUsageSummary) HasTotalLimit() bool`

HasTotalLimit returns a boolean if a field has been set.

### GetUsagePercent

`func (o *DtoFeatureUsageSummary) GetUsagePercent() string`

GetUsagePercent returns the UsagePercent field if non-nil, zero value otherwise.

### GetUsagePercentOk

`func (o *DtoFeatureUsageSummary) GetUsagePercentOk() (*string, bool)`

GetUsagePercentOk returns a tuple with the UsagePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePercent

`func (o *DtoFeatureUsageSummary) SetUsagePercent(v string)`

SetUsagePercent sets UsagePercent field to given value.

### HasUsagePercent

`func (o *DtoFeatureUsageSummary) HasUsagePercent() bool`

HasUsagePercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


