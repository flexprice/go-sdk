# DtoCreateTaxAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApply** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EntityId** | **string** |  | 
**EntityType** | [**TypesTaxRateEntityType**](TypesTaxRateEntityType.md) |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**TaxRateCode** | **string** |  | 

## Methods

### NewDtoCreateTaxAssociationRequest

`func NewDtoCreateTaxAssociationRequest(entityId string, entityType TypesTaxRateEntityType, taxRateCode string, ) *DtoCreateTaxAssociationRequest`

NewDtoCreateTaxAssociationRequest instantiates a new DtoCreateTaxAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateTaxAssociationRequestWithDefaults

`func NewDtoCreateTaxAssociationRequestWithDefaults() *DtoCreateTaxAssociationRequest`

NewDtoCreateTaxAssociationRequestWithDefaults instantiates a new DtoCreateTaxAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApply

`func (o *DtoCreateTaxAssociationRequest) GetAutoApply() bool`

GetAutoApply returns the AutoApply field if non-nil, zero value otherwise.

### GetAutoApplyOk

`func (o *DtoCreateTaxAssociationRequest) GetAutoApplyOk() (*bool, bool)`

GetAutoApplyOk returns a tuple with the AutoApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApply

`func (o *DtoCreateTaxAssociationRequest) SetAutoApply(v bool)`

SetAutoApply sets AutoApply field to given value.

### HasAutoApply

`func (o *DtoCreateTaxAssociationRequest) HasAutoApply() bool`

HasAutoApply returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoCreateTaxAssociationRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCreateTaxAssociationRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCreateTaxAssociationRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoCreateTaxAssociationRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEntityId

`func (o *DtoCreateTaxAssociationRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DtoCreateTaxAssociationRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DtoCreateTaxAssociationRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *DtoCreateTaxAssociationRequest) GetEntityType() TypesTaxRateEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoCreateTaxAssociationRequest) GetEntityTypeOk() (*TypesTaxRateEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoCreateTaxAssociationRequest) SetEntityType(v TypesTaxRateEntityType)`

SetEntityType sets EntityType field to given value.


### GetMetadata

`func (o *DtoCreateTaxAssociationRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateTaxAssociationRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateTaxAssociationRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateTaxAssociationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPriority

`func (o *DtoCreateTaxAssociationRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtoCreateTaxAssociationRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtoCreateTaxAssociationRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtoCreateTaxAssociationRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTaxRateCode

`func (o *DtoCreateTaxAssociationRequest) GetTaxRateCode() string`

GetTaxRateCode returns the TaxRateCode field if non-nil, zero value otherwise.

### GetTaxRateCodeOk

`func (o *DtoCreateTaxAssociationRequest) GetTaxRateCodeOk() (*string, bool)`

GetTaxRateCodeOk returns a tuple with the TaxRateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateCode

`func (o *DtoCreateTaxAssociationRequest) SetTaxRateCode(v string)`

SetTaxRateCode sets TaxRateCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


