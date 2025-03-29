# DtoUpdatePaymentStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**PaymentStatus** | [**TypesPaymentStatus**](TypesPaymentStatus.md) |  | 

## Methods

### NewDtoUpdatePaymentStatusRequest

`func NewDtoUpdatePaymentStatusRequest(paymentStatus TypesPaymentStatus, ) *DtoUpdatePaymentStatusRequest`

NewDtoUpdatePaymentStatusRequest instantiates a new DtoUpdatePaymentStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoUpdatePaymentStatusRequestWithDefaults

`func NewDtoUpdatePaymentStatusRequestWithDefaults() *DtoUpdatePaymentStatusRequest`

NewDtoUpdatePaymentStatusRequestWithDefaults instantiates a new DtoUpdatePaymentStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoUpdatePaymentStatusRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoUpdatePaymentStatusRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoUpdatePaymentStatusRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoUpdatePaymentStatusRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *DtoUpdatePaymentStatusRequest) GetPaymentStatus() TypesPaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *DtoUpdatePaymentStatusRequest) GetPaymentStatusOk() (*TypesPaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *DtoUpdatePaymentStatusRequest) SetPaymentStatus(v TypesPaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


