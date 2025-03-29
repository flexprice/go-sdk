# DtoUpdateWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTopupAmount** | Pointer to **float32** |  | [optional] 
**AutoTopupMinBalance** | Pointer to **float32** |  | [optional] 
**AutoTopupTrigger** | Pointer to [**TypesAutoTopupTrigger**](TypesAutoTopupTrigger.md) |  | [optional] 
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

### GetAutoTopupAmount

`func (o *DtoUpdateWalletRequest) GetAutoTopupAmount() float32`

GetAutoTopupAmount returns the AutoTopupAmount field if non-nil, zero value otherwise.

### GetAutoTopupAmountOk

`func (o *DtoUpdateWalletRequest) GetAutoTopupAmountOk() (*float32, bool)`

GetAutoTopupAmountOk returns a tuple with the AutoTopupAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupAmount

`func (o *DtoUpdateWalletRequest) SetAutoTopupAmount(v float32)`

SetAutoTopupAmount sets AutoTopupAmount field to given value.

### HasAutoTopupAmount

`func (o *DtoUpdateWalletRequest) HasAutoTopupAmount() bool`

HasAutoTopupAmount returns a boolean if a field has been set.

### GetAutoTopupMinBalance

`func (o *DtoUpdateWalletRequest) GetAutoTopupMinBalance() float32`

GetAutoTopupMinBalance returns the AutoTopupMinBalance field if non-nil, zero value otherwise.

### GetAutoTopupMinBalanceOk

`func (o *DtoUpdateWalletRequest) GetAutoTopupMinBalanceOk() (*float32, bool)`

GetAutoTopupMinBalanceOk returns a tuple with the AutoTopupMinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupMinBalance

`func (o *DtoUpdateWalletRequest) SetAutoTopupMinBalance(v float32)`

SetAutoTopupMinBalance sets AutoTopupMinBalance field to given value.

### HasAutoTopupMinBalance

`func (o *DtoUpdateWalletRequest) HasAutoTopupMinBalance() bool`

HasAutoTopupMinBalance returns a boolean if a field has been set.

### GetAutoTopupTrigger

`func (o *DtoUpdateWalletRequest) GetAutoTopupTrigger() TypesAutoTopupTrigger`

GetAutoTopupTrigger returns the AutoTopupTrigger field if non-nil, zero value otherwise.

### GetAutoTopupTriggerOk

`func (o *DtoUpdateWalletRequest) GetAutoTopupTriggerOk() (*TypesAutoTopupTrigger, bool)`

GetAutoTopupTriggerOk returns a tuple with the AutoTopupTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupTrigger

`func (o *DtoUpdateWalletRequest) SetAutoTopupTrigger(v TypesAutoTopupTrigger)`

SetAutoTopupTrigger sets AutoTopupTrigger field to given value.

### HasAutoTopupTrigger

`func (o *DtoUpdateWalletRequest) HasAutoTopupTrigger() bool`

HasAutoTopupTrigger returns a boolean if a field has been set.

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


