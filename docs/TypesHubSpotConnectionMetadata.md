# TypesHubSpotConnectionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | Private App Access Token (encrypted) | [optional] 
**AppId** | Pointer to **string** | HubSpot App ID (optional, not encrypted) | [optional] 
**ClientSecret** | Pointer to **string** | Private App Client Secret for webhook verification (encrypted) | [optional] 

## Methods

### NewTypesHubSpotConnectionMetadata

`func NewTypesHubSpotConnectionMetadata() *TypesHubSpotConnectionMetadata`

NewTypesHubSpotConnectionMetadata instantiates a new TypesHubSpotConnectionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesHubSpotConnectionMetadataWithDefaults

`func NewTypesHubSpotConnectionMetadataWithDefaults() *TypesHubSpotConnectionMetadata`

NewTypesHubSpotConnectionMetadataWithDefaults instantiates a new TypesHubSpotConnectionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *TypesHubSpotConnectionMetadata) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TypesHubSpotConnectionMetadata) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TypesHubSpotConnectionMetadata) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TypesHubSpotConnectionMetadata) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAppId

`func (o *TypesHubSpotConnectionMetadata) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *TypesHubSpotConnectionMetadata) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *TypesHubSpotConnectionMetadata) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *TypesHubSpotConnectionMetadata) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetClientSecret

`func (o *TypesHubSpotConnectionMetadata) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TypesHubSpotConnectionMetadata) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TypesHubSpotConnectionMetadata) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *TypesHubSpotConnectionMetadata) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


