# DtoCreatePriceUnitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCurrency** | **string** | base_currency  is the currency that the price unit is based on | 
**Code** | **string** |  | 
**ConversionRate** | **string** | ConversionRate defines the exchange rate from this price unit to the base currency. This rate is used to convert amounts in the custom price unit to the base currency for storage and billing.  Conversion formula:   price_unit_amount * conversion_rate &#x3D; base_currency_amount  Example:   If conversion_rate &#x3D; \&quot;0.01\&quot; and base_currency &#x3D; \&quot;usd\&quot;:   100 price_unit tokens * 0.01 &#x3D; 1.00 USD  Note: Rounding precision is determined by the base currency (e.g., USD uses 2 decimal places, JPY uses 0). | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 
**Symbol** | **string** |  | 

## Methods

### NewDtoCreatePriceUnitRequest

`func NewDtoCreatePriceUnitRequest(baseCurrency string, code string, conversionRate string, name string, symbol string, ) *DtoCreatePriceUnitRequest`

NewDtoCreatePriceUnitRequest instantiates a new DtoCreatePriceUnitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreatePriceUnitRequestWithDefaults

`func NewDtoCreatePriceUnitRequestWithDefaults() *DtoCreatePriceUnitRequest`

NewDtoCreatePriceUnitRequestWithDefaults instantiates a new DtoCreatePriceUnitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseCurrency

`func (o *DtoCreatePriceUnitRequest) GetBaseCurrency() string`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *DtoCreatePriceUnitRequest) GetBaseCurrencyOk() (*string, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *DtoCreatePriceUnitRequest) SetBaseCurrency(v string)`

SetBaseCurrency sets BaseCurrency field to given value.


### GetCode

`func (o *DtoCreatePriceUnitRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DtoCreatePriceUnitRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DtoCreatePriceUnitRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetConversionRate

`func (o *DtoCreatePriceUnitRequest) GetConversionRate() string`

GetConversionRate returns the ConversionRate field if non-nil, zero value otherwise.

### GetConversionRateOk

`func (o *DtoCreatePriceUnitRequest) GetConversionRateOk() (*string, bool)`

GetConversionRateOk returns a tuple with the ConversionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRate

`func (o *DtoCreatePriceUnitRequest) SetConversionRate(v string)`

SetConversionRate sets ConversionRate field to given value.


### GetMetadata

`func (o *DtoCreatePriceUnitRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreatePriceUnitRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreatePriceUnitRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreatePriceUnitRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoCreatePriceUnitRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreatePriceUnitRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreatePriceUnitRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *DtoCreatePriceUnitRequest) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *DtoCreatePriceUnitRequest) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *DtoCreatePriceUnitRequest) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


