# DtoUpdateTaxRateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | code is the updated unique alphanumeric identifier for the tax rate | [optional] 
**Description** | Pointer to **string** | description is the updated text description for the tax rate | [optional] 
**Metadata** | Pointer to **map[string]string** | metadata contains updated key-value pairs that will replace existing metadata | [optional] 
**Name** | Pointer to **string** | name is the updated human-readable name for the tax rate | [optional] 
**TaxRateStatus** | Pointer to [**TypesTaxRateStatus**](TypesTaxRateStatus.md) |  | [optional] 

## Methods

### NewDtoUpdateTaxRateRequest

`func NewDtoUpdateTaxRateRequest() *DtoUpdateTaxRateRequest`

NewDtoUpdateTaxRateRequest instantiates a new DtoUpdateTaxRateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdateTaxRateRequestWithDefaults

`func NewDtoUpdateTaxRateRequestWithDefaults() *DtoUpdateTaxRateRequest`

NewDtoUpdateTaxRateRequestWithDefaults instantiates a new DtoUpdateTaxRateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DtoUpdateTaxRateRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DtoUpdateTaxRateRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DtoUpdateTaxRateRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DtoUpdateTaxRateRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *DtoUpdateTaxRateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoUpdateTaxRateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoUpdateTaxRateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoUpdateTaxRateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdateTaxRateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdateTaxRateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdateTaxRateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdateTaxRateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoUpdateTaxRateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoUpdateTaxRateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoUpdateTaxRateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoUpdateTaxRateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxRateStatus

`func (o *DtoUpdateTaxRateRequest) GetTaxRateStatus() TypesTaxRateStatus`

GetTaxRateStatus returns the TaxRateStatus field if non-nil, zero value otherwise.

### GetTaxRateStatusOk

`func (o *DtoUpdateTaxRateRequest) GetTaxRateStatusOk() (*TypesTaxRateStatus, bool)`

GetTaxRateStatusOk returns a tuple with the TaxRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateStatus

`func (o *DtoUpdateTaxRateRequest) SetTaxRateStatus(v TypesTaxRateStatus)`

SetTaxRateStatus sets TaxRateStatus field to given value.

### HasTaxRateStatus

`func (o *DtoUpdateTaxRateRequest) HasTaxRateStatus() bool`

HasTaxRateStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


