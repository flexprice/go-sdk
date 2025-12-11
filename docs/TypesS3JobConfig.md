# TypesS3JobConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | S3 bucket name | [optional] 
**Compression** | Pointer to [**TypesS3CompressionType**](TypesS3CompressionType.md) |  | [optional] 
**Encryption** | Pointer to [**TypesS3EncryptionType**](TypesS3EncryptionType.md) |  | [optional] 
**EndpointUrl** | Pointer to **string** | Custom S3 endpoint URL (e.g., \&quot;http://minio:9000\&quot; for MinIO) | [optional] 
**KeyPrefix** | Pointer to **string** | Optional prefix for S3 keys (e.g., \&quot;flexprice-exports/\&quot;) | [optional] 
**Region** | Pointer to **string** | AWS region (e.g., \&quot;us-west-2\&quot;) | [optional] 
**UsePathStyle** | Pointer to **bool** | Use path-style addressing instead of virtual-hosted-style (required for MinIO) | [optional] 

## Methods

### NewTypesS3JobConfig

`func NewTypesS3JobConfig() *TypesS3JobConfig`

NewTypesS3JobConfig instantiates a new TypesS3JobConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesS3JobConfigWithDefaults

`func NewTypesS3JobConfigWithDefaults() *TypesS3JobConfig`

NewTypesS3JobConfigWithDefaults instantiates a new TypesS3JobConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *TypesS3JobConfig) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *TypesS3JobConfig) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *TypesS3JobConfig) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *TypesS3JobConfig) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCompression

`func (o *TypesS3JobConfig) GetCompression() TypesS3CompressionType`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *TypesS3JobConfig) GetCompressionOk() (*TypesS3CompressionType, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *TypesS3JobConfig) SetCompression(v TypesS3CompressionType)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *TypesS3JobConfig) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetEncryption

`func (o *TypesS3JobConfig) GetEncryption() TypesS3EncryptionType`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *TypesS3JobConfig) GetEncryptionOk() (*TypesS3EncryptionType, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *TypesS3JobConfig) SetEncryption(v TypesS3EncryptionType)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *TypesS3JobConfig) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetEndpointUrl

`func (o *TypesS3JobConfig) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *TypesS3JobConfig) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *TypesS3JobConfig) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.

### HasEndpointUrl

`func (o *TypesS3JobConfig) HasEndpointUrl() bool`

HasEndpointUrl returns a boolean if a field has been set.

### GetKeyPrefix

`func (o *TypesS3JobConfig) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *TypesS3JobConfig) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *TypesS3JobConfig) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *TypesS3JobConfig) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.

### GetRegion

`func (o *TypesS3JobConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TypesS3JobConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TypesS3JobConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TypesS3JobConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUsePathStyle

`func (o *TypesS3JobConfig) GetUsePathStyle() bool`

GetUsePathStyle returns the UsePathStyle field if non-nil, zero value otherwise.

### GetUsePathStyleOk

`func (o *TypesS3JobConfig) GetUsePathStyleOk() (*bool, bool)`

GetUsePathStyleOk returns a tuple with the UsePathStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePathStyle

`func (o *TypesS3JobConfig) SetUsePathStyle(v bool)`

SetUsePathStyle sets UsePathStyle field to given value.

### HasUsePathStyle

`func (o *TypesS3JobConfig) HasUsePathStyle() bool`

HasUsePathStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


