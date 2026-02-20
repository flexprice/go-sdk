# TypesMoyasarConnectionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublishableKey** | Pointer to **string** | Moyasar Publishable Key (encrypted, for frontend use) | [optional] 
**SecretKey** | Pointer to **string** | Moyasar Secret Key (encrypted) | [optional] 
**WebhookSecret** | Pointer to **string** | Moyasar Webhook Secret (encrypted, optional) | [optional] 

## Methods

### NewTypesMoyasarConnectionMetadata

`func NewTypesMoyasarConnectionMetadata() *TypesMoyasarConnectionMetadata`

NewTypesMoyasarConnectionMetadata instantiates a new TypesMoyasarConnectionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesMoyasarConnectionMetadataWithDefaults

`func NewTypesMoyasarConnectionMetadataWithDefaults() *TypesMoyasarConnectionMetadata`

NewTypesMoyasarConnectionMetadataWithDefaults instantiates a new TypesMoyasarConnectionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublishableKey

`func (o *TypesMoyasarConnectionMetadata) GetPublishableKey() string`

GetPublishableKey returns the PublishableKey field if non-nil, zero value otherwise.

### GetPublishableKeyOk

`func (o *TypesMoyasarConnectionMetadata) GetPublishableKeyOk() (*string, bool)`

GetPublishableKeyOk returns a tuple with the PublishableKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishableKey

`func (o *TypesMoyasarConnectionMetadata) SetPublishableKey(v string)`

SetPublishableKey sets PublishableKey field to given value.

### HasPublishableKey

`func (o *TypesMoyasarConnectionMetadata) HasPublishableKey() bool`

HasPublishableKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *TypesMoyasarConnectionMetadata) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *TypesMoyasarConnectionMetadata) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *TypesMoyasarConnectionMetadata) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *TypesMoyasarConnectionMetadata) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *TypesMoyasarConnectionMetadata) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *TypesMoyasarConnectionMetadata) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *TypesMoyasarConnectionMetadata) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *TypesMoyasarConnectionMetadata) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


