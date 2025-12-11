# TypesChargebeeConnectionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | Chargebee API key (encrypted) | [optional] 
**Site** | Pointer to **string** | Chargebee site name (not encrypted) | [optional] 
**WebhookPassword** | Pointer to **string** | Basic Auth password for webhooks (encrypted) | [optional] 
**WebhookSecret** | Pointer to **string** | Chargebee Webhook Secret (encrypted, optional, NOT USED in v2) | [optional] 
**WebhookUsername** | Pointer to **string** | Basic Auth username for webhooks (encrypted) | [optional] 

## Methods

### NewTypesChargebeeConnectionMetadata

`func NewTypesChargebeeConnectionMetadata() *TypesChargebeeConnectionMetadata`

NewTypesChargebeeConnectionMetadata instantiates a new TypesChargebeeConnectionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesChargebeeConnectionMetadataWithDefaults

`func NewTypesChargebeeConnectionMetadataWithDefaults() *TypesChargebeeConnectionMetadata`

NewTypesChargebeeConnectionMetadataWithDefaults instantiates a new TypesChargebeeConnectionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TypesChargebeeConnectionMetadata) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TypesChargebeeConnectionMetadata) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TypesChargebeeConnectionMetadata) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *TypesChargebeeConnectionMetadata) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSite

`func (o *TypesChargebeeConnectionMetadata) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *TypesChargebeeConnectionMetadata) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *TypesChargebeeConnectionMetadata) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *TypesChargebeeConnectionMetadata) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetWebhookPassword

`func (o *TypesChargebeeConnectionMetadata) GetWebhookPassword() string`

GetWebhookPassword returns the WebhookPassword field if non-nil, zero value otherwise.

### GetWebhookPasswordOk

`func (o *TypesChargebeeConnectionMetadata) GetWebhookPasswordOk() (*string, bool)`

GetWebhookPasswordOk returns a tuple with the WebhookPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookPassword

`func (o *TypesChargebeeConnectionMetadata) SetWebhookPassword(v string)`

SetWebhookPassword sets WebhookPassword field to given value.

### HasWebhookPassword

`func (o *TypesChargebeeConnectionMetadata) HasWebhookPassword() bool`

HasWebhookPassword returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *TypesChargebeeConnectionMetadata) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *TypesChargebeeConnectionMetadata) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *TypesChargebeeConnectionMetadata) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *TypesChargebeeConnectionMetadata) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### GetWebhookUsername

`func (o *TypesChargebeeConnectionMetadata) GetWebhookUsername() string`

GetWebhookUsername returns the WebhookUsername field if non-nil, zero value otherwise.

### GetWebhookUsernameOk

`func (o *TypesChargebeeConnectionMetadata) GetWebhookUsernameOk() (*string, bool)`

GetWebhookUsernameOk returns a tuple with the WebhookUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUsername

`func (o *TypesChargebeeConnectionMetadata) SetWebhookUsername(v string)`

SetWebhookUsername sets WebhookUsername field to given value.

### HasWebhookUsername

`func (o *TypesChargebeeConnectionMetadata) HasWebhookUsername() bool`

HasWebhookUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


