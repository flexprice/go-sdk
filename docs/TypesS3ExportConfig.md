# TypesS3ExportConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | S3 bucket name | [optional] 
**Compression** | Pointer to [**TypesS3CompressionType**](TypesS3CompressionType.md) |  | [optional] 
**Encryption** | Pointer to [**TypesS3EncryptionType**](TypesS3EncryptionType.md) |  | [optional] 
**IsFlexpriceManaged** | Pointer to **bool** | If true, use Flexprice-managed S3 credentials instead of user-provided | [optional] 
**KeyPrefix** | Pointer to **string** | Optional prefix for S3 keys (e.g., \&quot;flexprice-exports/\&quot;) | [optional] 
**Region** | Pointer to **string** | AWS region (e.g., \&quot;us-west-2\&quot;) | [optional] 

## Methods

### NewTypesS3ExportConfig

`func NewTypesS3ExportConfig() *TypesS3ExportConfig`

NewTypesS3ExportConfig instantiates a new TypesS3ExportConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesS3ExportConfigWithDefaults

`func NewTypesS3ExportConfigWithDefaults() *TypesS3ExportConfig`

NewTypesS3ExportConfigWithDefaults instantiates a new TypesS3ExportConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *TypesS3ExportConfig) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *TypesS3ExportConfig) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *TypesS3ExportConfig) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *TypesS3ExportConfig) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCompression

`func (o *TypesS3ExportConfig) GetCompression() TypesS3CompressionType`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *TypesS3ExportConfig) GetCompressionOk() (*TypesS3CompressionType, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *TypesS3ExportConfig) SetCompression(v TypesS3CompressionType)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *TypesS3ExportConfig) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetEncryption

`func (o *TypesS3ExportConfig) GetEncryption() TypesS3EncryptionType`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *TypesS3ExportConfig) GetEncryptionOk() (*TypesS3EncryptionType, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *TypesS3ExportConfig) SetEncryption(v TypesS3EncryptionType)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *TypesS3ExportConfig) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetIsFlexpriceManaged

`func (o *TypesS3ExportConfig) GetIsFlexpriceManaged() bool`

GetIsFlexpriceManaged returns the IsFlexpriceManaged field if non-nil, zero value otherwise.

### GetIsFlexpriceManagedOk

`func (o *TypesS3ExportConfig) GetIsFlexpriceManagedOk() (*bool, bool)`

GetIsFlexpriceManagedOk returns a tuple with the IsFlexpriceManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlexpriceManaged

`func (o *TypesS3ExportConfig) SetIsFlexpriceManaged(v bool)`

SetIsFlexpriceManaged sets IsFlexpriceManaged field to given value.

### HasIsFlexpriceManaged

`func (o *TypesS3ExportConfig) HasIsFlexpriceManaged() bool`

HasIsFlexpriceManaged returns a boolean if a field has been set.

### GetKeyPrefix

`func (o *TypesS3ExportConfig) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *TypesS3ExportConfig) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *TypesS3ExportConfig) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *TypesS3ExportConfig) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.

### GetRegion

`func (o *TypesS3ExportConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TypesS3ExportConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TypesS3ExportConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TypesS3ExportConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


