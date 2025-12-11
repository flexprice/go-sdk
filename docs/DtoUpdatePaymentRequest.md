# DtoUpdatePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** |  | [optional] 
**FailedAt** | Pointer to **string** |  | [optional] 
**GatewayPaymentId** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PaymentGateway** | Pointer to **string** |  | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**PaymentStatus** | Pointer to **string** |  | [optional] 
**SucceededAt** | Pointer to **string** |  | [optional] 

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

### GetErrorMessage

`func (o *DtoUpdatePaymentRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *DtoUpdatePaymentRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *DtoUpdatePaymentRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *DtoUpdatePaymentRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetFailedAt

`func (o *DtoUpdatePaymentRequest) GetFailedAt() string`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *DtoUpdatePaymentRequest) GetFailedAtOk() (*string, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *DtoUpdatePaymentRequest) SetFailedAt(v string)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *DtoUpdatePaymentRequest) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetGatewayPaymentId

`func (o *DtoUpdatePaymentRequest) GetGatewayPaymentId() string`

GetGatewayPaymentId returns the GatewayPaymentId field if non-nil, zero value otherwise.

### GetGatewayPaymentIdOk

`func (o *DtoUpdatePaymentRequest) GetGatewayPaymentIdOk() (*string, bool)`

GetGatewayPaymentIdOk returns a tuple with the GatewayPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentId

`func (o *DtoUpdatePaymentRequest) SetGatewayPaymentId(v string)`

SetGatewayPaymentId sets GatewayPaymentId field to given value.

### HasGatewayPaymentId

`func (o *DtoUpdatePaymentRequest) HasGatewayPaymentId() bool`

HasGatewayPaymentId returns a boolean if a field has been set.

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

### GetPaymentGateway

`func (o *DtoUpdatePaymentRequest) GetPaymentGateway() string`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *DtoUpdatePaymentRequest) GetPaymentGatewayOk() (*string, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *DtoUpdatePaymentRequest) SetPaymentGateway(v string)`

SetPaymentGateway sets PaymentGateway field to given value.

### HasPaymentGateway

`func (o *DtoUpdatePaymentRequest) HasPaymentGateway() bool`

HasPaymentGateway returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *DtoUpdatePaymentRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *DtoUpdatePaymentRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *DtoUpdatePaymentRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *DtoUpdatePaymentRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

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

### GetSucceededAt

`func (o *DtoUpdatePaymentRequest) GetSucceededAt() string`

GetSucceededAt returns the SucceededAt field if non-nil, zero value otherwise.

### GetSucceededAtOk

`func (o *DtoUpdatePaymentRequest) GetSucceededAtOk() (*string, bool)`

GetSucceededAtOk returns a tuple with the SucceededAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededAt

`func (o *DtoUpdatePaymentRequest) SetSucceededAt(v string)`

SetSucceededAt sets SucceededAt field to given value.

### HasSucceededAt

`func (o *DtoUpdatePaymentRequest) HasSucceededAt() bool`

HasSucceededAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


