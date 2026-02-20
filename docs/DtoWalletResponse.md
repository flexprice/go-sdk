# DtoWalletResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSettings** | Pointer to [**TypesAlertSettings**](TypesAlertSettings.md) |  | [optional] 
**AlertState** | Pointer to [**TypesAlertState**](TypesAlertState.md) |  | [optional] 
**AutoTopup** | Pointer to [**TypesAutoTopup**](TypesAutoTopup.md) |  | [optional] 
**Balance** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**TypesWalletConfig**](TypesWalletConfig.md) |  | [optional] 
**ConversionRate** | Pointer to **string** | amount in the currency &#x3D;  number of credits * conversion_rate ex if conversion_rate is 1, then 1 USD &#x3D; 1 credit ex if conversion_rate is 2, then 1 USD &#x3D; 0.5 credits ex if conversion_rate is 0.5, then 1 USD &#x3D; 2 credits | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreditBalance** | Pointer to **string** |  | [optional] 
**CreditsAvailableBreakdown** | Pointer to [**TypesCreditBreakdown**](TypesCreditBreakdown.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TopupConversionRate** | Pointer to **string** | topup_conversion_rate is the conversion rate for the topup to the currency ex if topup_conversion_rate is 1, then 1 USD &#x3D; 1 credit ex if topup_conversion_rate is 2, then 1 USD &#x3D; 0.5 credits ex if topup_conversion_rate is 0.5, then 1 USD &#x3D; 2 credits | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
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

### GetAlertSettings

`func (o *DtoWalletResponse) GetAlertSettings() TypesAlertSettings`

GetAlertSettings returns the AlertSettings field if non-nil, zero value otherwise.

### GetAlertSettingsOk

`func (o *DtoWalletResponse) GetAlertSettingsOk() (*TypesAlertSettings, bool)`

GetAlertSettingsOk returns a tuple with the AlertSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSettings

`func (o *DtoWalletResponse) SetAlertSettings(v TypesAlertSettings)`

SetAlertSettings sets AlertSettings field to given value.

### HasAlertSettings

`func (o *DtoWalletResponse) HasAlertSettings() bool`

HasAlertSettings returns a boolean if a field has been set.

### GetAlertState

`func (o *DtoWalletResponse) GetAlertState() TypesAlertState`

GetAlertState returns the AlertState field if non-nil, zero value otherwise.

### GetAlertStateOk

`func (o *DtoWalletResponse) GetAlertStateOk() (*TypesAlertState, bool)`

GetAlertStateOk returns a tuple with the AlertState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertState

`func (o *DtoWalletResponse) SetAlertState(v TypesAlertState)`

SetAlertState sets AlertState field to given value.

### HasAlertState

`func (o *DtoWalletResponse) HasAlertState() bool`

HasAlertState returns a boolean if a field has been set.

### GetAutoTopup

`func (o *DtoWalletResponse) GetAutoTopup() TypesAutoTopup`

GetAutoTopup returns the AutoTopup field if non-nil, zero value otherwise.

### GetAutoTopupOk

`func (o *DtoWalletResponse) GetAutoTopupOk() (*TypesAutoTopup, bool)`

GetAutoTopupOk returns a tuple with the AutoTopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopup

`func (o *DtoWalletResponse) SetAutoTopup(v TypesAutoTopup)`

SetAutoTopup sets AutoTopup field to given value.

### HasAutoTopup

`func (o *DtoWalletResponse) HasAutoTopup() bool`

HasAutoTopup returns a boolean if a field has been set.

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

### GetCreatedBy

`func (o *DtoWalletResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoWalletResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoWalletResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoWalletResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

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

### GetCreditsAvailableBreakdown

`func (o *DtoWalletResponse) GetCreditsAvailableBreakdown() TypesCreditBreakdown`

GetCreditsAvailableBreakdown returns the CreditsAvailableBreakdown field if non-nil, zero value otherwise.

### GetCreditsAvailableBreakdownOk

`func (o *DtoWalletResponse) GetCreditsAvailableBreakdownOk() (*TypesCreditBreakdown, bool)`

GetCreditsAvailableBreakdownOk returns a tuple with the CreditsAvailableBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsAvailableBreakdown

`func (o *DtoWalletResponse) SetCreditsAvailableBreakdown(v TypesCreditBreakdown)`

SetCreditsAvailableBreakdown sets CreditsAvailableBreakdown field to given value.

### HasCreditsAvailableBreakdown

`func (o *DtoWalletResponse) HasCreditsAvailableBreakdown() bool`

HasCreditsAvailableBreakdown returns a boolean if a field has been set.

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

### GetEnvironmentId

`func (o *DtoWalletResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoWalletResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoWalletResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoWalletResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

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

### GetStatus

`func (o *DtoWalletResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoWalletResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoWalletResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoWalletResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoWalletResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoWalletResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoWalletResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoWalletResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTopupConversionRate

`func (o *DtoWalletResponse) GetTopupConversionRate() string`

GetTopupConversionRate returns the TopupConversionRate field if non-nil, zero value otherwise.

### GetTopupConversionRateOk

`func (o *DtoWalletResponse) GetTopupConversionRateOk() (*string, bool)`

GetTopupConversionRateOk returns a tuple with the TopupConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupConversionRate

`func (o *DtoWalletResponse) SetTopupConversionRate(v string)`

SetTopupConversionRate sets TopupConversionRate field to given value.

### HasTopupConversionRate

`func (o *DtoWalletResponse) HasTopupConversionRate() bool`

HasTopupConversionRate returns a boolean if a field has been set.

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

### GetUpdatedBy

`func (o *DtoWalletResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoWalletResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoWalletResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoWalletResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

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


