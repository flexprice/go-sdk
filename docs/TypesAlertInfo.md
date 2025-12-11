# TypesAlertInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSettings** | Pointer to [**TypesAlertSettings**](TypesAlertSettings.md) |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**ValueAtTime** | Pointer to **float32** |  | [optional] 

## Methods

### NewTypesAlertInfo

`func NewTypesAlertInfo() *TypesAlertInfo`

NewTypesAlertInfo instantiates a new TypesAlertInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesAlertInfoWithDefaults

`func NewTypesAlertInfoWithDefaults() *TypesAlertInfo`

NewTypesAlertInfoWithDefaults instantiates a new TypesAlertInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSettings

`func (o *TypesAlertInfo) GetAlertSettings() TypesAlertSettings`

GetAlertSettings returns the AlertSettings field if non-nil, zero value otherwise.

### GetAlertSettingsOk

`func (o *TypesAlertInfo) GetAlertSettingsOk() (*TypesAlertSettings, bool)`

GetAlertSettingsOk returns a tuple with the AlertSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSettings

`func (o *TypesAlertInfo) SetAlertSettings(v TypesAlertSettings)`

SetAlertSettings sets AlertSettings field to given value.

### HasAlertSettings

`func (o *TypesAlertInfo) HasAlertSettings() bool`

HasAlertSettings returns a boolean if a field has been set.

### GetTimestamp

`func (o *TypesAlertInfo) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TypesAlertInfo) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TypesAlertInfo) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TypesAlertInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValueAtTime

`func (o *TypesAlertInfo) GetValueAtTime() float32`

GetValueAtTime returns the ValueAtTime field if non-nil, zero value otherwise.

### GetValueAtTimeOk

`func (o *TypesAlertInfo) GetValueAtTimeOk() (*float32, bool)`

GetValueAtTimeOk returns a tuple with the ValueAtTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueAtTime

`func (o *TypesAlertInfo) SetValueAtTime(v float32)`

SetValueAtTime sets ValueAtTime field to given value.

### HasValueAtTime

`func (o *TypesAlertInfo) HasValueAtTime() bool`

HasValueAtTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


