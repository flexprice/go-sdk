# DtoSignUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | Pointer to **string** |  | [optional] 
**TenantName** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoSignUpRequest

`func NewDtoSignUpRequest(email string, ) *DtoSignUpRequest`

NewDtoSignUpRequest instantiates a new DtoSignUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSignUpRequestWithDefaults

`func NewDtoSignUpRequestWithDefaults() *DtoSignUpRequest`

NewDtoSignUpRequestWithDefaults instantiates a new DtoSignUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *DtoSignUpRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DtoSignUpRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DtoSignUpRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *DtoSignUpRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DtoSignUpRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DtoSignUpRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DtoSignUpRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetTenantName

`func (o *DtoSignUpRequest) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *DtoSignUpRequest) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *DtoSignUpRequest) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *DtoSignUpRequest) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### GetToken

`func (o *DtoSignUpRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DtoSignUpRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DtoSignUpRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DtoSignUpRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


