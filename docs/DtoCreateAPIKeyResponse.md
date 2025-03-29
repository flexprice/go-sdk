# DtoCreateAPIKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to [**DtoSecretResponse**](DtoSecretResponse.md) |  | [optional] 

## Methods

### NewDtoCreateAPIKeyResponse

`func NewDtoCreateAPIKeyResponse() *DtoCreateAPIKeyResponse`

NewDtoCreateAPIKeyResponse instantiates a new DtoCreateAPIKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateAPIKeyResponseWithDefaults

`func NewDtoCreateAPIKeyResponseWithDefaults() *DtoCreateAPIKeyResponse`

NewDtoCreateAPIKeyResponseWithDefaults instantiates a new DtoCreateAPIKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *DtoCreateAPIKeyResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *DtoCreateAPIKeyResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *DtoCreateAPIKeyResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *DtoCreateAPIKeyResponse) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecret

`func (o *DtoCreateAPIKeyResponse) GetSecret() DtoSecretResponse`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DtoCreateAPIKeyResponse) GetSecretOk() (*DtoSecretResponse, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DtoCreateAPIKeyResponse) SetSecret(v DtoSecretResponse)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DtoCreateAPIKeyResponse) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


