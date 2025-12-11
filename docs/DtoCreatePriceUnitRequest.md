# DtoCreatePriceUnitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseCurrency** | **string** |  | 
**Code** | **string** |  | 
**ConversionRate** | **string** |  | 
**Name** | **string** |  | 
**Precision** | Pointer to **int32** |  | [optional] 
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


### GetPrecision

`func (o *DtoCreatePriceUnitRequest) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *DtoCreatePriceUnitRequest) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *DtoCreatePriceUnitRequest) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *DtoCreatePriceUnitRequest) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

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


