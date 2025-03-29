# DtoUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Tenant** | Pointer to [**DtoTenantResponse**](DtoTenantResponse.md) |  | [optional] 

## Methods

### NewDtoUserResponse

`func NewDtoUserResponse() *DtoUserResponse`

NewDtoUserResponse instantiates a new DtoUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUserResponseWithDefaults

`func NewDtoUserResponseWithDefaults() *DtoUserResponse`

NewDtoUserResponseWithDefaults instantiates a new DtoUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *DtoUserResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DtoUserResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DtoUserResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DtoUserResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *DtoUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoUserResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoUserResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenant

`func (o *DtoUserResponse) GetTenant() DtoTenantResponse`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *DtoUserResponse) GetTenantOk() (*DtoTenantResponse, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *DtoUserResponse) SetTenant(v DtoTenantResponse)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *DtoUserResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


