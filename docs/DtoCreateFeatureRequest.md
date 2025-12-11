# DtoCreateFeatureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSettings** | Pointer to [**TypesAlertSettings**](TypesAlertSettings.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LookupKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Meter** | Pointer to [**DtoCreateMeterRequest**](DtoCreateMeterRequest.md) |  | [optional] 
**MeterId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Type** | [**TypesFeatureType**](TypesFeatureType.md) |  | 
**UnitPlural** | Pointer to **string** |  | [optional] 
**UnitSingular** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCreateFeatureRequest

`func NewDtoCreateFeatureRequest(name string, type_ TypesFeatureType, ) *DtoCreateFeatureRequest`

NewDtoCreateFeatureRequest instantiates a new DtoCreateFeatureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateFeatureRequestWithDefaults

`func NewDtoCreateFeatureRequestWithDefaults() *DtoCreateFeatureRequest`

NewDtoCreateFeatureRequestWithDefaults instantiates a new DtoCreateFeatureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSettings

`func (o *DtoCreateFeatureRequest) GetAlertSettings() TypesAlertSettings`

GetAlertSettings returns the AlertSettings field if non-nil, zero value otherwise.

### GetAlertSettingsOk

`func (o *DtoCreateFeatureRequest) GetAlertSettingsOk() (*TypesAlertSettings, bool)`

GetAlertSettingsOk returns a tuple with the AlertSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSettings

`func (o *DtoCreateFeatureRequest) SetAlertSettings(v TypesAlertSettings)`

SetAlertSettings sets AlertSettings field to given value.

### HasAlertSettings

`func (o *DtoCreateFeatureRequest) HasAlertSettings() bool`

HasAlertSettings returns a boolean if a field has been set.

### GetDescription

`func (o *DtoCreateFeatureRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoCreateFeatureRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoCreateFeatureRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoCreateFeatureRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLookupKey

`func (o *DtoCreateFeatureRequest) GetLookupKey() string`

GetLookupKey returns the LookupKey field if non-nil, zero value otherwise.

### GetLookupKeyOk

`func (o *DtoCreateFeatureRequest) GetLookupKeyOk() (*string, bool)`

GetLookupKeyOk returns a tuple with the LookupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupKey

`func (o *DtoCreateFeatureRequest) SetLookupKey(v string)`

SetLookupKey sets LookupKey field to given value.

### HasLookupKey

`func (o *DtoCreateFeatureRequest) HasLookupKey() bool`

HasLookupKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateFeatureRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateFeatureRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateFeatureRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateFeatureRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMeter

`func (o *DtoCreateFeatureRequest) GetMeter() DtoCreateMeterRequest`

GetMeter returns the Meter field if non-nil, zero value otherwise.

### GetMeterOk

`func (o *DtoCreateFeatureRequest) GetMeterOk() (*DtoCreateMeterRequest, bool)`

GetMeterOk returns a tuple with the Meter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeter

`func (o *DtoCreateFeatureRequest) SetMeter(v DtoCreateMeterRequest)`

SetMeter sets Meter field to given value.

### HasMeter

`func (o *DtoCreateFeatureRequest) HasMeter() bool`

HasMeter returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoCreateFeatureRequest) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoCreateFeatureRequest) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoCreateFeatureRequest) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *DtoCreateFeatureRequest) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetName

`func (o *DtoCreateFeatureRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreateFeatureRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreateFeatureRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *DtoCreateFeatureRequest) GetType() TypesFeatureType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoCreateFeatureRequest) GetTypeOk() (*TypesFeatureType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoCreateFeatureRequest) SetType(v TypesFeatureType)`

SetType sets Type field to given value.


### GetUnitPlural

`func (o *DtoCreateFeatureRequest) GetUnitPlural() string`

GetUnitPlural returns the UnitPlural field if non-nil, zero value otherwise.

### GetUnitPluralOk

`func (o *DtoCreateFeatureRequest) GetUnitPluralOk() (*string, bool)`

GetUnitPluralOk returns a tuple with the UnitPlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPlural

`func (o *DtoCreateFeatureRequest) SetUnitPlural(v string)`

SetUnitPlural sets UnitPlural field to given value.

### HasUnitPlural

`func (o *DtoCreateFeatureRequest) HasUnitPlural() bool`

HasUnitPlural returns a boolean if a field has been set.

### GetUnitSingular

`func (o *DtoCreateFeatureRequest) GetUnitSingular() string`

GetUnitSingular returns the UnitSingular field if non-nil, zero value otherwise.

### GetUnitSingularOk

`func (o *DtoCreateFeatureRequest) GetUnitSingularOk() (*string, bool)`

GetUnitSingularOk returns a tuple with the UnitSingular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSingular

`func (o *DtoCreateFeatureRequest) SetUnitSingular(v string)`

SetUnitSingular sets UnitSingular field to given value.

### HasUnitSingular

`func (o *DtoCreateFeatureRequest) HasUnitSingular() bool`

HasUnitSingular returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


