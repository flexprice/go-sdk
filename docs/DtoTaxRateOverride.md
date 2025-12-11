# DtoTaxRateOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApply** | Pointer to **bool** |  | [optional] [default to true]
**Currency** | **string** |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**TaxRateCode** | **string** |  | 

## Methods

### NewDtoTaxRateOverride

`func NewDtoTaxRateOverride(currency string, taxRateCode string, ) *DtoTaxRateOverride`

NewDtoTaxRateOverride instantiates a new DtoTaxRateOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTaxRateOverrideWithDefaults

`func NewDtoTaxRateOverrideWithDefaults() *DtoTaxRateOverride`

NewDtoTaxRateOverrideWithDefaults instantiates a new DtoTaxRateOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApply

`func (o *DtoTaxRateOverride) GetAutoApply() bool`

GetAutoApply returns the AutoApply field if non-nil, zero value otherwise.

### GetAutoApplyOk

`func (o *DtoTaxRateOverride) GetAutoApplyOk() (*bool, bool)`

GetAutoApplyOk returns a tuple with the AutoApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApply

`func (o *DtoTaxRateOverride) SetAutoApply(v bool)`

SetAutoApply sets AutoApply field to given value.

### HasAutoApply

`func (o *DtoTaxRateOverride) HasAutoApply() bool`

HasAutoApply returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoTaxRateOverride) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoTaxRateOverride) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoTaxRateOverride) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetMetadata

`func (o *DtoTaxRateOverride) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoTaxRateOverride) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoTaxRateOverride) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoTaxRateOverride) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPriority

`func (o *DtoTaxRateOverride) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtoTaxRateOverride) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtoTaxRateOverride) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtoTaxRateOverride) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTaxRateCode

`func (o *DtoTaxRateOverride) GetTaxRateCode() string`

GetTaxRateCode returns the TaxRateCode field if non-nil, zero value otherwise.

### GetTaxRateCodeOk

`func (o *DtoTaxRateOverride) GetTaxRateCodeOk() (*string, bool)`

GetTaxRateCodeOk returns a tuple with the TaxRateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateCode

`func (o *DtoTaxRateOverride) SetTaxRateCode(v string)`

SetTaxRateCode sets TaxRateCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


