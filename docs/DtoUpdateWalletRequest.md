# DtoUpdateWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertConfig** | Pointer to [**DtoAlertConfig**](DtoAlertConfig.md) |  | [optional] 
**AlertEnabled** | Pointer to **bool** |  | [optional] 
**AutoTopup** | Pointer to [**TypesAutoTopup**](TypesAutoTopup.md) |  | [optional] 
**Config** | Pointer to [**TypesWalletConfig**](TypesWalletConfig.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoUpdateWalletRequest

`func NewDtoUpdateWalletRequest() *DtoUpdateWalletRequest`

NewDtoUpdateWalletRequest instantiates a new DtoUpdateWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdateWalletRequestWithDefaults

`func NewDtoUpdateWalletRequestWithDefaults() *DtoUpdateWalletRequest`

NewDtoUpdateWalletRequestWithDefaults instantiates a new DtoUpdateWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertConfig

`func (o *DtoUpdateWalletRequest) GetAlertConfig() DtoAlertConfig`

GetAlertConfig returns the AlertConfig field if non-nil, zero value otherwise.

### GetAlertConfigOk

`func (o *DtoUpdateWalletRequest) GetAlertConfigOk() (*DtoAlertConfig, bool)`

GetAlertConfigOk returns a tuple with the AlertConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfig

`func (o *DtoUpdateWalletRequest) SetAlertConfig(v DtoAlertConfig)`

SetAlertConfig sets AlertConfig field to given value.

### HasAlertConfig

`func (o *DtoUpdateWalletRequest) HasAlertConfig() bool`

HasAlertConfig returns a boolean if a field has been set.

### GetAlertEnabled

`func (o *DtoUpdateWalletRequest) GetAlertEnabled() bool`

GetAlertEnabled returns the AlertEnabled field if non-nil, zero value otherwise.

### GetAlertEnabledOk

`func (o *DtoUpdateWalletRequest) GetAlertEnabledOk() (*bool, bool)`

GetAlertEnabledOk returns a tuple with the AlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEnabled

`func (o *DtoUpdateWalletRequest) SetAlertEnabled(v bool)`

SetAlertEnabled sets AlertEnabled field to given value.

### HasAlertEnabled

`func (o *DtoUpdateWalletRequest) HasAlertEnabled() bool`

HasAlertEnabled returns a boolean if a field has been set.

### GetAutoTopup

`func (o *DtoUpdateWalletRequest) GetAutoTopup() TypesAutoTopup`

GetAutoTopup returns the AutoTopup field if non-nil, zero value otherwise.

### GetAutoTopupOk

`func (o *DtoUpdateWalletRequest) GetAutoTopupOk() (*TypesAutoTopup, bool)`

GetAutoTopupOk returns a tuple with the AutoTopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopup

`func (o *DtoUpdateWalletRequest) SetAutoTopup(v TypesAutoTopup)`

SetAutoTopup sets AutoTopup field to given value.

### HasAutoTopup

`func (o *DtoUpdateWalletRequest) HasAutoTopup() bool`

HasAutoTopup returns a boolean if a field has been set.

### GetConfig

`func (o *DtoUpdateWalletRequest) GetConfig() TypesWalletConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DtoUpdateWalletRequest) GetConfigOk() (*TypesWalletConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DtoUpdateWalletRequest) SetConfig(v TypesWalletConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DtoUpdateWalletRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *DtoUpdateWalletRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoUpdateWalletRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoUpdateWalletRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoUpdateWalletRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdateWalletRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdateWalletRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdateWalletRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdateWalletRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoUpdateWalletRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoUpdateWalletRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoUpdateWalletRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoUpdateWalletRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


