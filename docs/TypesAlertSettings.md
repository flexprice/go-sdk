# TypesAlertSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertEnabled** | Pointer to **bool** |  | [optional] 
**Critical** | Pointer to [**TypesAlertThreshold**](TypesAlertThreshold.md) |  | [optional] 
**Info** | Pointer to [**TypesAlertThreshold**](TypesAlertThreshold.md) |  | [optional] 
**Warning** | Pointer to [**TypesAlertThreshold**](TypesAlertThreshold.md) |  | [optional] 

## Methods

### NewTypesAlertSettings

`func NewTypesAlertSettings() *TypesAlertSettings`

NewTypesAlertSettings instantiates a new TypesAlertSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesAlertSettingsWithDefaults

`func NewTypesAlertSettingsWithDefaults() *TypesAlertSettings`

NewTypesAlertSettingsWithDefaults instantiates a new TypesAlertSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertEnabled

`func (o *TypesAlertSettings) GetAlertEnabled() bool`

GetAlertEnabled returns the AlertEnabled field if non-nil, zero value otherwise.

### GetAlertEnabledOk

`func (o *TypesAlertSettings) GetAlertEnabledOk() (*bool, bool)`

GetAlertEnabledOk returns a tuple with the AlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEnabled

`func (o *TypesAlertSettings) SetAlertEnabled(v bool)`

SetAlertEnabled sets AlertEnabled field to given value.

### HasAlertEnabled

`func (o *TypesAlertSettings) HasAlertEnabled() bool`

HasAlertEnabled returns a boolean if a field has been set.

### GetCritical

`func (o *TypesAlertSettings) GetCritical() TypesAlertThreshold`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *TypesAlertSettings) GetCriticalOk() (*TypesAlertThreshold, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *TypesAlertSettings) SetCritical(v TypesAlertThreshold)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *TypesAlertSettings) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetInfo

`func (o *TypesAlertSettings) GetInfo() TypesAlertThreshold`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TypesAlertSettings) GetInfoOk() (*TypesAlertThreshold, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TypesAlertSettings) SetInfo(v TypesAlertThreshold)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TypesAlertSettings) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetWarning

`func (o *TypesAlertSettings) GetWarning() TypesAlertThreshold`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *TypesAlertSettings) GetWarningOk() (*TypesAlertThreshold, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *TypesAlertSettings) SetWarning(v TypesAlertThreshold)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *TypesAlertSettings) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


