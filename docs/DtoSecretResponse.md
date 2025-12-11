# DtoSecretResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**DisplayId** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUsedAt** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**TypesSecretProvider**](TypesSecretProvider.md) |  | [optional] 
**Roles** | Pointer to **[]string** | RBAC roles | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**Type** | Pointer to [**TypesSecretType**](TypesSecretType.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to [**TypesUserType**](TypesUserType.md) |  | [optional] 

## Methods

### NewDtoSecretResponse

`func NewDtoSecretResponse() *DtoSecretResponse`

NewDtoSecretResponse instantiates a new DtoSecretResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSecretResponseWithDefaults

`func NewDtoSecretResponseWithDefaults() *DtoSecretResponse`

NewDtoSecretResponseWithDefaults instantiates a new DtoSecretResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DtoSecretResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoSecretResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoSecretResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoSecretResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayId

`func (o *DtoSecretResponse) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *DtoSecretResponse) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *DtoSecretResponse) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *DtoSecretResponse) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *DtoSecretResponse) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DtoSecretResponse) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DtoSecretResponse) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DtoSecretResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *DtoSecretResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoSecretResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoSecretResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoSecretResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *DtoSecretResponse) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *DtoSecretResponse) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *DtoSecretResponse) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *DtoSecretResponse) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetName

`func (o *DtoSecretResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoSecretResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoSecretResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtoSecretResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *DtoSecretResponse) GetProvider() TypesSecretProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DtoSecretResponse) GetProviderOk() (*TypesSecretProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DtoSecretResponse) SetProvider(v TypesSecretProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DtoSecretResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRoles

`func (o *DtoSecretResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DtoSecretResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DtoSecretResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *DtoSecretResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStatus

`func (o *DtoSecretResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoSecretResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoSecretResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoSecretResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *DtoSecretResponse) GetType() TypesSecretType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoSecretResponse) GetTypeOk() (*TypesSecretType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoSecretResponse) SetType(v TypesSecretType)`

SetType sets Type field to given value.

### HasType

`func (o *DtoSecretResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoSecretResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoSecretResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoSecretResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoSecretResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserType

`func (o *DtoSecretResponse) GetUserType() TypesUserType`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *DtoSecretResponse) GetUserTypeOk() (*TypesUserType, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *DtoSecretResponse) SetUserType(v TypesUserType)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *DtoSecretResponse) HasUserType() bool`

HasUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


