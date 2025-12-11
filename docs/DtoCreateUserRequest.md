# DtoCreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | **[]string** | Roles are required | 
**Type** | [**TypesUserType**](TypesUserType.md) |  | 

## Methods

### NewDtoCreateUserRequest

`func NewDtoCreateUserRequest(roles []string, type_ TypesUserType, ) *DtoCreateUserRequest`

NewDtoCreateUserRequest instantiates a new DtoCreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateUserRequestWithDefaults

`func NewDtoCreateUserRequestWithDefaults() *DtoCreateUserRequest`

NewDtoCreateUserRequestWithDefaults instantiates a new DtoCreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *DtoCreateUserRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DtoCreateUserRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DtoCreateUserRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetType

`func (o *DtoCreateUserRequest) GetType() TypesUserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoCreateUserRequest) GetTypeOk() (*TypesUserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoCreateUserRequest) SetType(v TypesUserType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


