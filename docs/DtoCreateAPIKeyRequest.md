# DtoCreateAPIKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Type** | [**TypesSecretType**](TypesSecretType.md) |  | 

## Methods

### NewDtoCreateAPIKeyRequest

`func NewDtoCreateAPIKeyRequest(name string, type_ TypesSecretType, ) *DtoCreateAPIKeyRequest`

NewDtoCreateAPIKeyRequest instantiates a new DtoCreateAPIKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateAPIKeyRequestWithDefaults

`func NewDtoCreateAPIKeyRequestWithDefaults() *DtoCreateAPIKeyRequest`

NewDtoCreateAPIKeyRequestWithDefaults instantiates a new DtoCreateAPIKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *DtoCreateAPIKeyRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DtoCreateAPIKeyRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DtoCreateAPIKeyRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DtoCreateAPIKeyRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetName

`func (o *DtoCreateAPIKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreateAPIKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreateAPIKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *DtoCreateAPIKeyRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DtoCreateAPIKeyRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DtoCreateAPIKeyRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DtoCreateAPIKeyRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetType

`func (o *DtoCreateAPIKeyRequest) GetType() TypesSecretType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoCreateAPIKeyRequest) GetTypeOk() (*TypesSecretType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoCreateAPIKeyRequest) SetType(v TypesSecretType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


