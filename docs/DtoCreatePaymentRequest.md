# DtoCreatePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** |  | 
**CancelUrl** | Pointer to **string** |  | [optional] 
**Currency** | **string** |  | 
**DestinationId** | **string** |  | 
**DestinationType** | [**TypesPaymentDestinationType**](TypesPaymentDestinationType.md) |  | 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PaymentGateway** | Pointer to [**TypesPaymentGatewayType**](TypesPaymentGatewayType.md) |  | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**PaymentMethodType** | [**TypesPaymentMethodType**](TypesPaymentMethodType.md) |  | 
**ProcessPayment** | Pointer to **bool** |  | [optional] [default to true]
**SaveCardAndMakeDefault** | Pointer to **bool** |  | [optional] [default to false]
**SuccessUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCreatePaymentRequest

`func NewDtoCreatePaymentRequest(amount string, currency string, destinationId string, destinationType TypesPaymentDestinationType, paymentMethodType TypesPaymentMethodType, ) *DtoCreatePaymentRequest`

NewDtoCreatePaymentRequest instantiates a new DtoCreatePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreatePaymentRequestWithDefaults

`func NewDtoCreatePaymentRequestWithDefaults() *DtoCreatePaymentRequest`

NewDtoCreatePaymentRequestWithDefaults instantiates a new DtoCreatePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoCreatePaymentRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoCreatePaymentRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoCreatePaymentRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCancelUrl

`func (o *DtoCreatePaymentRequest) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *DtoCreatePaymentRequest) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *DtoCreatePaymentRequest) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *DtoCreatePaymentRequest) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoCreatePaymentRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCreatePaymentRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCreatePaymentRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDestinationId

`func (o *DtoCreatePaymentRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DtoCreatePaymentRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DtoCreatePaymentRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetDestinationType

`func (o *DtoCreatePaymentRequest) GetDestinationType() TypesPaymentDestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *DtoCreatePaymentRequest) GetDestinationTypeOk() (*TypesPaymentDestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *DtoCreatePaymentRequest) SetDestinationType(v TypesPaymentDestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetIdempotencyKey

`func (o *DtoCreatePaymentRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoCreatePaymentRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoCreatePaymentRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *DtoCreatePaymentRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreatePaymentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreatePaymentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreatePaymentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreatePaymentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentGateway

`func (o *DtoCreatePaymentRequest) GetPaymentGateway() TypesPaymentGatewayType`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *DtoCreatePaymentRequest) GetPaymentGatewayOk() (*TypesPaymentGatewayType, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *DtoCreatePaymentRequest) SetPaymentGateway(v TypesPaymentGatewayType)`

SetPaymentGateway sets PaymentGateway field to given value.

### HasPaymentGateway

`func (o *DtoCreatePaymentRequest) HasPaymentGateway() bool`

HasPaymentGateway returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *DtoCreatePaymentRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *DtoCreatePaymentRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *DtoCreatePaymentRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *DtoCreatePaymentRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *DtoCreatePaymentRequest) GetPaymentMethodType() TypesPaymentMethodType`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *DtoCreatePaymentRequest) GetPaymentMethodTypeOk() (*TypesPaymentMethodType, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *DtoCreatePaymentRequest) SetPaymentMethodType(v TypesPaymentMethodType)`

SetPaymentMethodType sets PaymentMethodType field to given value.


### GetProcessPayment

`func (o *DtoCreatePaymentRequest) GetProcessPayment() bool`

GetProcessPayment returns the ProcessPayment field if non-nil, zero value otherwise.

### GetProcessPaymentOk

`func (o *DtoCreatePaymentRequest) GetProcessPaymentOk() (*bool, bool)`

GetProcessPaymentOk returns a tuple with the ProcessPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessPayment

`func (o *DtoCreatePaymentRequest) SetProcessPayment(v bool)`

SetProcessPayment sets ProcessPayment field to given value.

### HasProcessPayment

`func (o *DtoCreatePaymentRequest) HasProcessPayment() bool`

HasProcessPayment returns a boolean if a field has been set.

### GetSaveCardAndMakeDefault

`func (o *DtoCreatePaymentRequest) GetSaveCardAndMakeDefault() bool`

GetSaveCardAndMakeDefault returns the SaveCardAndMakeDefault field if non-nil, zero value otherwise.

### GetSaveCardAndMakeDefaultOk

`func (o *DtoCreatePaymentRequest) GetSaveCardAndMakeDefaultOk() (*bool, bool)`

GetSaveCardAndMakeDefaultOk returns a tuple with the SaveCardAndMakeDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveCardAndMakeDefault

`func (o *DtoCreatePaymentRequest) SetSaveCardAndMakeDefault(v bool)`

SetSaveCardAndMakeDefault sets SaveCardAndMakeDefault field to given value.

### HasSaveCardAndMakeDefault

`func (o *DtoCreatePaymentRequest) HasSaveCardAndMakeDefault() bool`

HasSaveCardAndMakeDefault returns a boolean if a field has been set.

### GetSuccessUrl

`func (o *DtoCreatePaymentRequest) GetSuccessUrl() string`

GetSuccessUrl returns the SuccessUrl field if non-nil, zero value otherwise.

### GetSuccessUrlOk

`func (o *DtoCreatePaymentRequest) GetSuccessUrlOk() (*string, bool)`

GetSuccessUrlOk returns a tuple with the SuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessUrl

`func (o *DtoCreatePaymentRequest) SetSuccessUrl(v string)`

SetSuccessUrl sets SuccessUrl field to given value.

### HasSuccessUrl

`func (o *DtoCreatePaymentRequest) HasSuccessUrl() bool`

HasSuccessUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


