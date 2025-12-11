# DtoWalletResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertConfig** | Pointer to [**TypesAlertConfig**](TypesAlertConfig.md) |  | [optional] 
**AlertEnabled** | Pointer to **bool** |  | [optional] 
**AlertState** | Pointer to **string** |  | [optional] 
**AutoTopupAmount** | Pointer to **string** |  | [optional] 
**AutoTopupMinBalance** | Pointer to **string** |  | [optional] 
**AutoTopupTrigger** | Pointer to [**TypesAutoTopupTrigger**](TypesAutoTopupTrigger.md) |  | [optional] 
**Balance** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**TypesWalletConfig**](TypesWalletConfig.md) |  | [optional] 
**ConversionRate** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreditBalance** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**WalletStatus** | Pointer to [**TypesWalletStatus**](TypesWalletStatus.md) |  | [optional] 
**WalletType** | Pointer to [**TypesWalletType**](TypesWalletType.md) |  | [optional] 

## Methods

### NewDtoWalletResponse

`func NewDtoWalletResponse() *DtoWalletResponse`

NewDtoWalletResponse instantiates a new DtoWalletResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoWalletResponseWithDefaults

`func NewDtoWalletResponseWithDefaults() *DtoWalletResponse`

NewDtoWalletResponseWithDefaults instantiates a new DtoWalletResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertConfig

`func (o *DtoWalletResponse) GetAlertConfig() TypesAlertConfig`

GetAlertConfig returns the AlertConfig field if non-nil, zero value otherwise.

### GetAlertConfigOk

`func (o *DtoWalletResponse) GetAlertConfigOk() (*TypesAlertConfig, bool)`

GetAlertConfigOk returns a tuple with the AlertConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertConfig

`func (o *DtoWalletResponse) SetAlertConfig(v TypesAlertConfig)`

SetAlertConfig sets AlertConfig field to given value.

### HasAlertConfig

`func (o *DtoWalletResponse) HasAlertConfig() bool`

HasAlertConfig returns a boolean if a field has been set.

### GetAlertEnabled

`func (o *DtoWalletResponse) GetAlertEnabled() bool`

GetAlertEnabled returns the AlertEnabled field if non-nil, zero value otherwise.

### GetAlertEnabledOk

`func (o *DtoWalletResponse) GetAlertEnabledOk() (*bool, bool)`

GetAlertEnabledOk returns a tuple with the AlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEnabled

`func (o *DtoWalletResponse) SetAlertEnabled(v bool)`

SetAlertEnabled sets AlertEnabled field to given value.

### HasAlertEnabled

`func (o *DtoWalletResponse) HasAlertEnabled() bool`

HasAlertEnabled returns a boolean if a field has been set.

### GetAlertState

`func (o *DtoWalletResponse) GetAlertState() string`

GetAlertState returns the AlertState field if non-nil, zero value otherwise.

### GetAlertStateOk

`func (o *DtoWalletResponse) GetAlertStateOk() (*string, bool)`

GetAlertStateOk returns a tuple with the AlertState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertState

`func (o *DtoWalletResponse) SetAlertState(v string)`

SetAlertState sets AlertState field to given value.

### HasAlertState

`func (o *DtoWalletResponse) HasAlertState() bool`

HasAlertState returns a boolean if a field has been set.

### GetAutoTopupAmount

`func (o *DtoWalletResponse) GetAutoTopupAmount() string`

GetAutoTopupAmount returns the AutoTopupAmount field if non-nil, zero value otherwise.

### GetAutoTopupAmountOk

`func (o *DtoWalletResponse) GetAutoTopupAmountOk() (*string, bool)`

GetAutoTopupAmountOk returns a tuple with the AutoTopupAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupAmount

`func (o *DtoWalletResponse) SetAutoTopupAmount(v string)`

SetAutoTopupAmount sets AutoTopupAmount field to given value.

### HasAutoTopupAmount

`func (o *DtoWalletResponse) HasAutoTopupAmount() bool`

HasAutoTopupAmount returns a boolean if a field has been set.

### GetAutoTopupMinBalance

`func (o *DtoWalletResponse) GetAutoTopupMinBalance() string`

GetAutoTopupMinBalance returns the AutoTopupMinBalance field if non-nil, zero value otherwise.

### GetAutoTopupMinBalanceOk

`func (o *DtoWalletResponse) GetAutoTopupMinBalanceOk() (*string, bool)`

GetAutoTopupMinBalanceOk returns a tuple with the AutoTopupMinBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupMinBalance

`func (o *DtoWalletResponse) SetAutoTopupMinBalance(v string)`

SetAutoTopupMinBalance sets AutoTopupMinBalance field to given value.

### HasAutoTopupMinBalance

`func (o *DtoWalletResponse) HasAutoTopupMinBalance() bool`

HasAutoTopupMinBalance returns a boolean if a field has been set.

### GetAutoTopupTrigger

`func (o *DtoWalletResponse) GetAutoTopupTrigger() TypesAutoTopupTrigger`

GetAutoTopupTrigger returns the AutoTopupTrigger field if non-nil, zero value otherwise.

### GetAutoTopupTriggerOk

`func (o *DtoWalletResponse) GetAutoTopupTriggerOk() (*TypesAutoTopupTrigger, bool)`

GetAutoTopupTriggerOk returns a tuple with the AutoTopupTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopupTrigger

`func (o *DtoWalletResponse) SetAutoTopupTrigger(v TypesAutoTopupTrigger)`

SetAutoTopupTrigger sets AutoTopupTrigger field to given value.

### HasAutoTopupTrigger

`func (o *DtoWalletResponse) HasAutoTopupTrigger() bool`

HasAutoTopupTrigger returns a boolean if a field has been set.

### GetBalance

`func (o *DtoWalletResponse) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *DtoWalletResponse) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *DtoWalletResponse) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *DtoWalletResponse) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetConfig

`func (o *DtoWalletResponse) GetConfig() TypesWalletConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DtoWalletResponse) GetConfigOk() (*TypesWalletConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DtoWalletResponse) SetConfig(v TypesWalletConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DtoWalletResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConversionRate

`func (o *DtoWalletResponse) GetConversionRate() string`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *DtoWalletResponse) GetConversionRateOk() (*string, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *DtoWalletResponse) SetConversionRate(v string)`

SetConversionRate sets ConversionRate field to given value.

### HasConversionRate

`func (o *DtoWalletResponse) HasConversionRate() bool`

HasConversionRate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoWalletResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoWalletResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoWalletResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoWalletResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditBalance

`func (o *DtoWalletResponse) GetCreditBalance() string`

GetCreditBalance returns the CreditBalance field if non-nil, zero value otherwise.

### GetCreditBalanceOk

`func (o *DtoWalletResponse) GetCreditBalanceOk() (*string, bool)`

GetCreditBalanceOk returns a tuple with the CreditBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalance

`func (o *DtoWalletResponse) SetCreditBalance(v string)`

SetCreditBalance sets CreditBalance field to given value.

### HasCreditBalance

`func (o *DtoWalletResponse) HasCreditBalance() bool`

HasCreditBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoWalletResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoWalletResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoWalletResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoWalletResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoWalletResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoWalletResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoWalletResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoWalletResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *DtoWalletResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoWalletResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoWalletResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoWalletResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DtoWalletResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoWalletResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoWalletResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoWalletResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoWalletResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoWalletResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoWalletResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoWalletResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoWalletResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoWalletResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoWalletResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoWalletResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoWalletResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoWalletResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoWalletResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoWalletResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWalletStatus

`func (o *DtoWalletResponse) GetWalletStatus() TypesWalletStatus`

GetWalletStatus returns the WalletStatus field if non-nil, zero value otherwise.

### GetWalletStatusOk

`func (o *DtoWalletResponse) GetWalletStatusOk() (*TypesWalletStatus, bool)`

GetWalletStatusOk returns a tuple with the WalletStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletStatus

`func (o *DtoWalletResponse) SetWalletStatus(v TypesWalletStatus)`

SetWalletStatus sets WalletStatus field to given value.

### HasWalletStatus

`func (o *DtoWalletResponse) HasWalletStatus() bool`

HasWalletStatus returns a boolean if a field has been set.

### GetWalletType

`func (o *DtoWalletResponse) GetWalletType() TypesWalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *DtoWalletResponse) GetWalletTypeOk() (*TypesWalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *DtoWalletResponse) SetWalletType(v TypesWalletType)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *DtoWalletResponse) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


