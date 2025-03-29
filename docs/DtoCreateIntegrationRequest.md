# DtoCreateIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | **map[string]string** |  | 
**Name** | **string** |  | 
**Provider** | [**TypesSecretProvider**](TypesSecretProvider.md) |  | 

## Methods

### NewDtoCreateIntegrationRequest

`func NewDtoCreateIntegrationRequest(credentials map[string]string, name string, provider TypesSecretProvider, ) *DtoCreateIntegrationRequest`

NewDtoCreateIntegrationRequest instantiates a new DtoCreateIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateIntegrationRequestWithDefaults

`func NewDtoCreateIntegrationRequestWithDefaults() *DtoCreateIntegrationRequest`

NewDtoCreateIntegrationRequestWithDefaults instantiates a new DtoCreateIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *DtoCreateIntegrationRequest) GetCredentials() map[string]string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *DtoCreateIntegrationRequest) GetCredentialsOk() (*map[string]string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *DtoCreateIntegrationRequest) SetCredentials(v map[string]string)`

SetCredentials sets Credentials field to given value.


### GetName

`func (o *DtoCreateIntegrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreateIntegrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreateIntegrationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *DtoCreateIntegrationRequest) GetProvider() TypesSecretProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DtoCreateIntegrationRequest) GetProviderOk() (*TypesSecretProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DtoCreateIntegrationRequest) SetProvider(v TypesSecretProvider)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


