# DtoUpdatePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PaymentStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoUpdatePaymentRequest

`func NewDtoUpdatePaymentRequest() *DtoUpdatePaymentRequest`

NewDtoUpdatePaymentRequest instantiates a new DtoUpdatePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdatePaymentRequestWithDefaults

`func NewDtoUpdatePaymentRequestWithDefaults() *DtoUpdatePaymentRequest`

NewDtoUpdatePaymentRequestWithDefaults instantiates a new DtoUpdatePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *DtoUpdatePaymentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoUpdatePaymentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoUpdatePaymentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoUpdatePaymentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *DtoUpdatePaymentRequest) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *DtoUpdatePaymentRequest) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *DtoUpdatePaymentRequest) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *DtoUpdatePaymentRequest) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


