# DtoCreateEntityIntegrationMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **string** |  | 
**EntityType** | [**TypesIntegrationEntityType**](TypesIntegrationEntityType.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ProviderEntityId** | **string** |  | 
**ProviderType** | **string** |  | 

## Methods

### NewDtoCreateEntityIntegrationMappingRequest

`func NewDtoCreateEntityIntegrationMappingRequest(entityId string, entityType TypesIntegrationEntityType, providerEntityId string, providerType string, ) *DtoCreateEntityIntegrationMappingRequest`

NewDtoCreateEntityIntegrationMappingRequest instantiates a new DtoCreateEntityIntegrationMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateEntityIntegrationMappingRequestWithDefaults

`func NewDtoCreateEntityIntegrationMappingRequestWithDefaults() *DtoCreateEntityIntegrationMappingRequest`

NewDtoCreateEntityIntegrationMappingRequestWithDefaults instantiates a new DtoCreateEntityIntegrationMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *DtoCreateEntityIntegrationMappingRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DtoCreateEntityIntegrationMappingRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DtoCreateEntityIntegrationMappingRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetEntityType

`func (o *DtoCreateEntityIntegrationMappingRequest) GetEntityType() TypesIntegrationEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoCreateEntityIntegrationMappingRequest) GetEntityTypeOk() (*TypesIntegrationEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoCreateEntityIntegrationMappingRequest) SetEntityType(v TypesIntegrationEntityType)`

SetEntityType sets EntityType field to given value.


### GetMetadata

`func (o *DtoCreateEntityIntegrationMappingRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateEntityIntegrationMappingRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateEntityIntegrationMappingRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateEntityIntegrationMappingRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProviderEntityId

`func (o *DtoCreateEntityIntegrationMappingRequest) GetProviderEntityId() string`

GetProviderEntityId returns the ProviderEntityId field if non-nil, zero value otherwise.

### GetProviderEntityIdOk

`func (o *DtoCreateEntityIntegrationMappingRequest) GetProviderEntityIdOk() (*string, bool)`

GetProviderEntityIdOk returns a tuple with the ProviderEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderEntityId

`func (o *DtoCreateEntityIntegrationMappingRequest) SetProviderEntityId(v string)`

SetProviderEntityId sets ProviderEntityId field to given value.


### GetProviderType

`func (o *DtoCreateEntityIntegrationMappingRequest) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *DtoCreateEntityIntegrationMappingRequest) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *DtoCreateEntityIntegrationMappingRequest) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


