# TypesNomodConnectionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | Nomod API Key (encrypted) | [optional] 
**WebhookSecret** | Pointer to **string** | Basic Auth secret for webhooks (encrypted, optional) | [optional] 

## Methods

### NewTypesNomodConnectionMetadata

`func NewTypesNomodConnectionMetadata() *TypesNomodConnectionMetadata`

NewTypesNomodConnectionMetadata instantiates a new TypesNomodConnectionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesNomodConnectionMetadataWithDefaults

`func NewTypesNomodConnectionMetadataWithDefaults() *TypesNomodConnectionMetadata`

NewTypesNomodConnectionMetadataWithDefaults instantiates a new TypesNomodConnectionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TypesNomodConnectionMetadata) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TypesNomodConnectionMetadata) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TypesNomodConnectionMetadata) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *TypesNomodConnectionMetadata) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *TypesNomodConnectionMetadata) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *TypesNomodConnectionMetadata) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *TypesNomodConnectionMetadata) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *TypesNomodConnectionMetadata) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


