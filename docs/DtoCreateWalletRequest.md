# DtoCreateWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSettings** | Pointer to [**TypesAlertSettings**](TypesAlertSettings.md) |  | [optional] 
**AutoTopup** | Pointer to [**TypesAutoTopup**](TypesAutoTopup.md) |  | [optional] 
**ConversionRate** | Pointer to **string** | amount in the currency &#x3D;  number of credits * conversion_rate ex if conversion_rate is 1, then 1 USD &#x3D; 1 credit ex if conversion_rate is 2, then 1 USD &#x3D; 0.5 credits ex if conversion_rate is 0.5, then 1 USD &#x3D; 2 credits | [optional] [default to "1"]
**Currency** | **string** |  | 
**CustomerId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalCustomerId** | Pointer to **string** | external_customer_id is the customer id in the external system | [optional] 
**InitialCreditsExpiryDateUtc** | Pointer to **string** | initial_credits_expiry_date_utc is the expiry date in UTC timezone (optional to set nil means no expiry) ex 2025-01-01 00:00:00 UTC | [optional] 
**InitialCreditsToLoad** | Pointer to **string** | initial_credits_to_load is the number of credits to load to the wallet if not provided, the wallet will be created with 0 balance NOTE: this is not the amount in the currency, but the number of credits | [optional] [default to "0"]
**InitialCreditsToLoadExpiryDate** | Pointer to **int32** | initial_credits_to_load_expiry_date YYYYMMDD format in UTC timezone (optional to set nil means no expiry) for ex 20250101 means the credits will expire on 2025-01-01 00:00:00 UTC hence they will be available for use until 2024-12-31 23:59:59 UTC | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PriceUnit** | Pointer to **string** | price_unit is the code of the price unit to use for wallet creation If provided, the price unit will be used to set the currency and conversion rate of the wallet: - currency: set to price unit&#39;s base_currency - conversion_rate: set to price unit&#39;s conversion_rate | [optional] 
**TopupConversionRate** | Pointer to **string** | topup_conversion_rate is the conversion rate for the topup to the currency ex if topup_conversion_rate is 1, then 1 USD &#x3D; 1 credit ex if topup_conversion_rate is 2, then 1 USD &#x3D; 0.5 credits ex if topup_conversion_rate is 0.5, then 1 USD &#x3D; 2 credits | [optional] 
**WalletType** | Pointer to [**TypesWalletType**](TypesWalletType.md) |  | [optional] 

## Methods

### NewDtoCreateWalletRequest

`func NewDtoCreateWalletRequest(currency string, ) *DtoCreateWalletRequest`

NewDtoCreateWalletRequest instantiates a new DtoCreateWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateWalletRequestWithDefaults

`func NewDtoCreateWalletRequestWithDefaults() *DtoCreateWalletRequest`

NewDtoCreateWalletRequestWithDefaults instantiates a new DtoCreateWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSettings

`func (o *DtoCreateWalletRequest) GetAlertSettings() TypesAlertSettings`

GetAlertSettings returns the AlertSettings field if non-nil, zero value otherwise.

### GetAlertSettingsOk

`func (o *DtoCreateWalletRequest) GetAlertSettingsOk() (*TypesAlertSettings, bool)`

GetAlertSettingsOk returns a tuple with the AlertSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSettings

`func (o *DtoCreateWalletRequest) SetAlertSettings(v TypesAlertSettings)`

SetAlertSettings sets AlertSettings field to given value.

### HasAlertSettings

`func (o *DtoCreateWalletRequest) HasAlertSettings() bool`

HasAlertSettings returns a boolean if a field has been set.

### GetAutoTopup

`func (o *DtoCreateWalletRequest) GetAutoTopup() TypesAutoTopup`

GetAutoTopup returns the AutoTopup field if non-nil, zero value otherwise.

### GetAutoTopupOk

`func (o *DtoCreateWalletRequest) GetAutoTopupOk() (*TypesAutoTopup, bool)`

GetAutoTopupOk returns a tuple with the AutoTopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopup

`func (o *DtoCreateWalletRequest) SetAutoTopup(v TypesAutoTopup)`

SetAutoTopup sets AutoTopup field to given value.

### HasAutoTopup

`func (o *DtoCreateWalletRequest) HasAutoTopup() bool`

HasAutoTopup returns a boolean if a field has been set.

### GetConversionRate

`func (o *DtoCreateWalletRequest) GetConversionRate() string`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *DtoCreateWalletRequest) GetConversionRateOk() (*string, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *DtoCreateWalletRequest) SetConversionRate(v string)`

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

### HasCustomerId

`func (o *DtoCreateWalletRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

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

### GetExternalCustomerId

`func (o *DtoCreateWalletRequest) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *DtoCreateWalletRequest) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *DtoCreateWalletRequest) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *DtoCreateWalletRequest) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetInitialCreditsExpiryDateUtc

`func (o *DtoCreateWalletRequest) GetInitialCreditsExpiryDateUtc() string`

GetInitialCreditsExpiryDateUtc returns the InitialCreditsExpiryDateUtc field if non-nil, zero value otherwise.

### GetInitialCreditsExpiryDateUtcOk

`func (o *DtoCreateWalletRequest) GetInitialCreditsExpiryDateUtcOk() (*string, bool)`

GetInitialCreditsExpiryDateUtcOk returns a tuple with the InitialCreditsExpiryDateUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCreditsExpiryDateUtc

`func (o *DtoCreateWalletRequest) SetInitialCreditsExpiryDateUtc(v string)`

SetInitialCreditsExpiryDateUtc sets InitialCreditsExpiryDateUtc field to given value.

### HasInitialCreditsExpiryDateUtc

`func (o *DtoCreateWalletRequest) HasInitialCreditsExpiryDateUtc() bool`

HasInitialCreditsExpiryDateUtc returns a boolean if a field has been set.

### GetInitialCreditsToLoad

`func (o *DtoCreateWalletRequest) GetInitialCreditsToLoad() string`

GetInitialCreditsToLoad returns the InitialCreditsToLoad field if non-nil, zero value otherwise.

### GetInitialCreditsToLoadOk

`func (o *DtoCreateWalletRequest) GetInitialCreditsToLoadOk() (*string, bool)`

GetInitialCreditsToLoadOk returns a tuple with the InitialCreditsToLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCreditsToLoad

`func (o *DtoCreateWalletRequest) SetInitialCreditsToLoad(v string)`

SetInitialCreditsToLoad sets InitialCreditsToLoad field to given value.

### HasInitialCreditsToLoad

`func (o *DtoCreateWalletRequest) HasInitialCreditsToLoad() bool`

HasInitialCreditsToLoad returns a boolean if a field has been set.

### GetInitialCreditsToLoadExpiryDate

`func (o *DtoCreateWalletRequest) GetInitialCreditsToLoadExpiryDate() int32`

GetInitialCreditsToLoadExpiryDate returns the InitialCreditsToLoadExpiryDate field if non-nil, zero value otherwise.

### GetInitialCreditsToLoadExpiryDateOk

`func (o *DtoCreateWalletRequest) GetInitialCreditsToLoadExpiryDateOk() (*int32, bool)`

GetInitialCreditsToLoadExpiryDateOk returns a tuple with the InitialCreditsToLoadExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCreditsToLoadExpiryDate

`func (o *DtoCreateWalletRequest) SetInitialCreditsToLoadExpiryDate(v int32)`

SetInitialCreditsToLoadExpiryDate sets InitialCreditsToLoadExpiryDate field to given value.

### HasInitialCreditsToLoadExpiryDate

`func (o *DtoCreateWalletRequest) HasInitialCreditsToLoadExpiryDate() bool`

HasInitialCreditsToLoadExpiryDate returns a boolean if a field has been set.

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

### GetPriceUnit

`func (o *DtoCreateWalletRequest) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *DtoCreateWalletRequest) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *DtoCreateWalletRequest) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *DtoCreateWalletRequest) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetTopupConversionRate

`func (o *DtoCreateWalletRequest) GetTopupConversionRate() string`

GetTopupConversionRate returns the TopupConversionRate field if non-nil, zero value otherwise.

### GetTopupConversionRateOk

`func (o *DtoCreateWalletRequest) GetTopupConversionRateOk() (*string, bool)`

GetTopupConversionRateOk returns a tuple with the TopupConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopupConversionRate

`func (o *DtoCreateWalletRequest) SetTopupConversionRate(v string)`

SetTopupConversionRate sets TopupConversionRate field to given value.

### HasTopupConversionRate

`func (o *DtoCreateWalletRequest) HasTopupConversionRate() bool`

HasTopupConversionRate returns a boolean if a field has been set.

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


