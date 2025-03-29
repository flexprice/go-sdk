# DtoCreateWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTopupAmount** | Pointer to **float32** |  | [optional] 
**AutoTopupMinBalance** | Pointer to **float32** |  | [optional] 
**AutoTopupTrigger** | Pointer to [**TypesAutoTopupTrigger**](TypesAutoTopupTrigger.md) |  | [optional] 
**Config** | Pointer to [**TypesWalletConfig**](TypesWalletConfig.md) |  | [optional] 
**ConversionRate** | Pointer to **float32** |  | [optional] 
**Currency** | **string** |  | 
**CustomerId** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WalletType** | Pointer to [**TypesWalletType**](TypesWalletType.md) |  | [optional] 

## Methods

### NewDtoCreateWalletRequest

`func NewDtoCreateWalletRequest(currency string, customerId string, ) *DtoCreateWalletRequest`

NewDtoCreateWalletRequest instantiates a new DtoCreateWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateWalletRequestWithDefaults

`func NewDtoCreateWalletRequestWithDefaults() *DtoCreateWalletRequest`

NewDtoCreateWalletRequestWithDefaults instantiates a new DtoCreateWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoTopupAmount

`func (o *DtoCreateWalletRequest) GetAutoTopupAmount() float32`

GetAutoTopupAmount returns the AutoTopupAmount field if non-nil, zero value otherwise.

### GetAutoTopupAmountOk

`func (o *DtoCreateWalletRequest) GetAutoTopupAmountOk() (*float32, bool)`

GetAutoTopupAmountOk returns a tuple with the AutoTopupAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupAmount

`func (o *DtoCreateWalletRequest) SetAutoTopupAmount(v float32)`

SetAutoTopupAmount sets AutoTopupAmount field to given value.

### HasAutoTopupAmount

`func (o *DtoCreateWalletRequest) HasAutoTopupAmount() bool`

HasAutoTopupAmount returns a boolean if a field has been set.

### GetAutoTopupMinBalance

`func (o *DtoCreateWalletRequest) GetAutoTopupMinBalance() float32`

GetAutoTopupMinBalance returns the AutoTopupMinBalance field if non-nil, zero value otherwise.

### GetAutoTopupMinBalanceOk

`func (o *DtoCreateWalletRequest) GetAutoTopupMinBalanceOk() (*float32, bool)`

GetAutoTopupMinBalanceOk returns a tuple with the AutoTopupMinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupMinBalance

`func (o *DtoCreateWalletRequest) SetAutoTopupMinBalance(v float32)`

SetAutoTopupMinBalance sets AutoTopupMinBalance field to given value.

### HasAutoTopupMinBalance

`func (o *DtoCreateWalletRequest) HasAutoTopupMinBalance() bool`

HasAutoTopupMinBalance returns a boolean if a field has been set.

### GetAutoTopupTrigger

`func (o *DtoCreateWalletRequest) GetAutoTopupTrigger() TypesAutoTopupTrigger`

GetAutoTopupTrigger returns the AutoTopupTrigger field if non-nil, zero value otherwise.

### GetAutoTopupTriggerOk

`func (o *DtoCreateWalletRequest) GetAutoTopupTriggerOk() (*TypesAutoTopupTrigger, bool)`

GetAutoTopupTriggerOk returns a tuple with the AutoTopupTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupTrigger

`func (o *DtoCreateWalletRequest) SetAutoTopupTrigger(v TypesAutoTopupTrigger)`

SetAutoTopupTrigger sets AutoTopupTrigger field to given value.

### HasAutoTopupTrigger

`func (o *DtoCreateWalletRequest) HasAutoTopupTrigger() bool`

HasAutoTopupTrigger returns a boolean if a field has been set.

### GetConfig

`func (o *DtoCreateWalletRequest) GetConfig() TypesWalletConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DtoCreateWalletRequest) GetConfigOk() (*TypesWalletConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DtoCreateWalletRequest) SetConfig(v TypesWalletConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DtoCreateWalletRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConversionRate

`func (o *DtoCreateWalletRequest) GetConversionRate() float32`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *DtoCreateWalletRequest) GetConversionRateOk() (*float32, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *DtoCreateWalletRequest) SetConversionRate(v float32)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *DtoCreateWalletRequest) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoCreateWalletRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCreateWalletRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCreateWalletRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *DtoCreateWalletRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoCreateWalletRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoCreateWalletRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDescription

`func (o *DtoCreateWalletRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoCreateWalletRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoCreateWalletRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoCreateWalletRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateWalletRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateWalletRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateWalletRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateWalletRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoCreateWalletRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreateWalletRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreateWalletRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoCreateWalletRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWalletType

`func (o *DtoCreateWalletRequest) GetWalletType() TypesWalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *DtoCreateWalletRequest) GetWalletTypeOk() (*TypesWalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *DtoCreateWalletRequest) SetWalletType(v TypesWalletType)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *DtoCreateWalletRequest) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


