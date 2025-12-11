# DtoCreateTaxRateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | code is the unique alphanumeric case sensitive identifier for the tax rate (required) | 
**Description** | Pointer to **string** | description is an optional text description providing details about the tax rate | [optional] 
**FixedValue** | Pointer to **string** | fixed_value is the fixed monetary amount when tax_rate_type is \&quot;fixed\&quot; | [optional] 
**Metadata** | Pointer to **map[string]string** | metadata contains additional key-value pairs for storing extra information | [optional] 
**Name** | **string** | name is the human-readable name for the tax rate (required) | 
**PercentageValue** | Pointer to **string** | percentage_value is the percentage value (0-100) when tax_rate_type is \&quot;percentage\&quot; | [optional] 
**Scope** | Pointer to [**TypesTaxRateScope**](TypesTaxRateScope.md) |  | [optional] 
**TaxRateType** | Pointer to [**TypesTaxRateType**](TypesTaxRateType.md) |  | [optional] 

## Methods

### NewDtoCreateTaxRateRequest

`func NewDtoCreateTaxRateRequest(code string, name string, ) *DtoCreateTaxRateRequest`

NewDtoCreateTaxRateRequest instantiates a new DtoCreateTaxRateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateTaxRateRequestWithDefaults

`func NewDtoCreateTaxRateRequestWithDefaults() *DtoCreateTaxRateRequest`

NewDtoCreateTaxRateRequestWithDefaults instantiates a new DtoCreateTaxRateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DtoCreateTaxRateRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DtoCreateTaxRateRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DtoCreateTaxRateRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *DtoCreateTaxRateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoCreateTaxRateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoCreateTaxRateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoCreateTaxRateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFixedValue

`func (o *DtoCreateTaxRateRequest) GetFixedValue() string`

GetFixedValue returns the FixedValue field if non-nil, zero value otherwise.

### GetFixedValueOk

`func (o *DtoCreateTaxRateRequest) GetFixedValueOk() (*string, bool)`

GetFixedValueOk returns a tuple with the FixedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedValue

`func (o *DtoCreateTaxRateRequest) SetFixedValue(v string)`

SetFixedValue sets FixedValue field to given value.

### HasFixedValue

`func (o *DtoCreateTaxRateRequest) HasFixedValue() bool`

HasFixedValue returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateTaxRateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateTaxRateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateTaxRateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateTaxRateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoCreateTaxRateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreateTaxRateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreateTaxRateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPercentageValue

`func (o *DtoCreateTaxRateRequest) GetPercentageValue() string`

GetPercentageValue returns the PercentageValue field if non-nil, zero value otherwise.

### GetPercentageValueOk

`func (o *DtoCreateTaxRateRequest) GetPercentageValueOk() (*string, bool)`

GetPercentageValueOk returns a tuple with the PercentageValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageValue

`func (o *DtoCreateTaxRateRequest) SetPercentageValue(v string)`

SetPercentageValue sets PercentageValue field to given value.

### HasPercentageValue

`func (o *DtoCreateTaxRateRequest) HasPercentageValue() bool`

HasPercentageValue returns a boolean if a field has been set.

### GetScope

`func (o *DtoCreateTaxRateRequest) GetScope() TypesTaxRateScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DtoCreateTaxRateRequest) GetScopeOk() (*TypesTaxRateScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DtoCreateTaxRateRequest) SetScope(v TypesTaxRateScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *DtoCreateTaxRateRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTaxRateType

`func (o *DtoCreateTaxRateRequest) GetTaxRateType() TypesTaxRateType`

GetTaxRateType returns the TaxRateType field if non-nil, zero value otherwise.

### GetTaxRateTypeOk

`func (o *DtoCreateTaxRateRequest) GetTaxRateTypeOk() (*TypesTaxRateType, bool)`

GetTaxRateTypeOk returns a tuple with the TaxRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateType

`func (o *DtoCreateTaxRateRequest) SetTaxRateType(v TypesTaxRateType)`

SetTaxRateType sets TaxRateType field to given value.

### HasTaxRateType

`func (o *DtoCreateTaxRateRequest) HasTaxRateType() bool`

HasTaxRateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


