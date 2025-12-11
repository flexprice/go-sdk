# DtoUpdateConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedSecretData** | Pointer to [**TypesConnectionMetadata**](TypesConnectionMetadata.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SyncConfig** | Pointer to [**TypesSyncConfig**](TypesSyncConfig.md) |  | [optional] 

## Methods

### NewDtoUpdateConnectionRequest

`func NewDtoUpdateConnectionRequest() *DtoUpdateConnectionRequest`

NewDtoUpdateConnectionRequest instantiates a new DtoUpdateConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdateConnectionRequestWithDefaults

`func NewDtoUpdateConnectionRequestWithDefaults() *DtoUpdateConnectionRequest`

NewDtoUpdateConnectionRequestWithDefaults instantiates a new DtoUpdateConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedSecretData

`func (o *DtoUpdateConnectionRequest) GetEncryptedSecretData() TypesConnectionMetadata`

GetEncryptedSecretData returns the EncryptedSecretData field if non-nil, zero value otherwise.

### GetEncryptedSecretDataOk

`func (o *DtoUpdateConnectionRequest) GetEncryptedSecretDataOk() (*TypesConnectionMetadata, bool)`

GetEncryptedSecretDataOk returns a tuple with the EncryptedSecretData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSecretData

`func (o *DtoUpdateConnectionRequest) SetEncryptedSecretData(v TypesConnectionMetadata)`

SetEncryptedSecretData sets EncryptedSecretData field to given value.

### HasEncryptedSecretData

`func (o *DtoUpdateConnectionRequest) HasEncryptedSecretData() bool`

HasEncryptedSecretData returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoUpdateConnectionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdateConnectionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdateConnectionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdateConnectionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DtoUpdateConnectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoUpdateConnectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoUpdateConnectionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoUpdateConnectionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSyncConfig

`func (o *DtoUpdateConnectionRequest) GetSyncConfig() TypesSyncConfig`

GetSyncConfig returns the SyncConfig field if non-nil, zero value otherwise.

### GetSyncConfigOk

`func (o *DtoUpdateConnectionRequest) GetSyncConfigOk() (*TypesSyncConfig, bool)`

GetSyncConfigOk returns a tuple with the SyncConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncConfig

`func (o *DtoUpdateConnectionRequest) SetSyncConfig(v TypesSyncConfig)`

SetSyncConfig sets SyncConfig field to given value.

### HasSyncConfig

`func (o *DtoUpdateConnectionRequest) HasSyncConfig() bool`

HasSyncConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


