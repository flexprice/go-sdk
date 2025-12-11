# DtoTaxAssociationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoApply** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 

## Methods

### NewDtoTaxAssociationUpdateRequest

`func NewDtoTaxAssociationUpdateRequest() *DtoTaxAssociationUpdateRequest`

NewDtoTaxAssociationUpdateRequest instantiates a new DtoTaxAssociationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTaxAssociationUpdateRequestWithDefaults

`func NewDtoTaxAssociationUpdateRequestWithDefaults() *DtoTaxAssociationUpdateRequest`

NewDtoTaxAssociationUpdateRequestWithDefaults instantiates a new DtoTaxAssociationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoApply

`func (o *DtoTaxAssociationUpdateRequest) GetAutoApply() bool`

GetAutoApply returns the AutoApply field if non-nil, zero value otherwise.

### GetAutoApplyOk

`func (o *DtoTaxAssociationUpdateRequest) GetAutoApplyOk() (*bool, bool)`

GetAutoApplyOk returns a tuple with the AutoApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApply

`func (o *DtoTaxAssociationUpdateRequest) SetAutoApply(v bool)`

SetAutoApply sets AutoApply field to given value.

### HasAutoApply

`func (o *DtoTaxAssociationUpdateRequest) HasAutoApply() bool`

HasAutoApply returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoTaxAssociationUpdateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoTaxAssociationUpdateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoTaxAssociationUpdateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoTaxAssociationUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPriority

`func (o *DtoTaxAssociationUpdateRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtoTaxAssociationUpdateRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtoTaxAssociationUpdateRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtoTaxAssociationUpdateRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


