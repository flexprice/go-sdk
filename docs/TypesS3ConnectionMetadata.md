# TypesS3ConnectionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKeyId** | Pointer to **string** | AWS access key (encrypted) | [optional] 
**AwsSecretAccessKey** | Pointer to **string** | AWS secret access key (encrypted) | [optional] 
**AwsSessionToken** | Pointer to **string** | AWS session token for temporary credentials (encrypted) | [optional] 

## Methods

### NewTypesS3ConnectionMetadata

`func NewTypesS3ConnectionMetadata() *TypesS3ConnectionMetadata`

NewTypesS3ConnectionMetadata instantiates a new TypesS3ConnectionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesS3ConnectionMetadataWithDefaults

`func NewTypesS3ConnectionMetadataWithDefaults() *TypesS3ConnectionMetadata`

NewTypesS3ConnectionMetadataWithDefaults instantiates a new TypesS3ConnectionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKeyId

`func (o *TypesS3ConnectionMetadata) GetAwsAccessKeyId() string`

GetAwsAccessKeyId returns the AwsAccessKeyId field if non-nil, zero value otherwise.

### GetAwsAccessKeyIdOk

`func (o *TypesS3ConnectionMetadata) GetAwsAccessKeyIdOk() (*string, bool)`

GetAwsAccessKeyIdOk returns a tuple with the AwsAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKeyId

`func (o *TypesS3ConnectionMetadata) SetAwsAccessKeyId(v string)`

SetAwsAccessKeyId sets AwsAccessKeyId field to given value.

### HasAwsAccessKeyId

`func (o *TypesS3ConnectionMetadata) HasAwsAccessKeyId() bool`

HasAwsAccessKeyId returns a boolean if a field has been set.

### GetAwsSecretAccessKey

`func (o *TypesS3ConnectionMetadata) GetAwsSecretAccessKey() string`

GetAwsSecretAccessKey returns the AwsSecretAccessKey field if non-nil, zero value otherwise.

### GetAwsSecretAccessKeyOk

`func (o *TypesS3ConnectionMetadata) GetAwsSecretAccessKeyOk() (*string, bool)`

GetAwsSecretAccessKeyOk returns a tuple with the AwsSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretAccessKey

`func (o *TypesS3ConnectionMetadata) SetAwsSecretAccessKey(v string)`

SetAwsSecretAccessKey sets AwsSecretAccessKey field to given value.

### HasAwsSecretAccessKey

`func (o *TypesS3ConnectionMetadata) HasAwsSecretAccessKey() bool`

HasAwsSecretAccessKey returns a boolean if a field has been set.

### GetAwsSessionToken

`func (o *TypesS3ConnectionMetadata) GetAwsSessionToken() string`

GetAwsSessionToken returns the AwsSessionToken field if non-nil, zero value otherwise.

### GetAwsSessionTokenOk

`func (o *TypesS3ConnectionMetadata) GetAwsSessionTokenOk() (*string, bool)`

GetAwsSessionTokenOk returns a tuple with the AwsSessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSessionToken

`func (o *TypesS3ConnectionMetadata) SetAwsSessionToken(v string)`

SetAwsSessionToken sets AwsSessionToken field to given value.

### HasAwsSessionToken

`func (o *TypesS3ConnectionMetadata) HasAwsSessionToken() bool`

HasAwsSessionToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


