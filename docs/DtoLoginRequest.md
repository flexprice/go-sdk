# DtoLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Password** | **string** |  | 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoLoginRequest

`func NewDtoLoginRequest(email string, password string, ) *DtoLoginRequest`

NewDtoLoginRequest instantiates a new DtoLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoLoginRequestWithDefaults

`func NewDtoLoginRequestWithDefaults() *DtoLoginRequest`

NewDtoLoginRequestWithDefaults instantiates a new DtoLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *DtoLoginRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DtoLoginRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DtoLoginRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *DtoLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DtoLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DtoLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetToken

`func (o *DtoLoginRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DtoLoginRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DtoLoginRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DtoLoginRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


