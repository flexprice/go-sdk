# DtoCreateInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDue** | **float32** |  | 
**AmountPaid** | Pointer to **float32** |  | [optional] 
**BillingPeriod** | Pointer to **string** |  | [optional] 
**BillingReason** | Pointer to [**TypesInvoiceBillingReason**](TypesInvoiceBillingReason.md) |  | [optional] 
**Currency** | **string** |  | 
**CustomerId** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**InvoiceNumber** | Pointer to **string** |  | [optional] 
**InvoiceStatus** | Pointer to [**TypesInvoiceStatus**](TypesInvoiceStatus.md) |  | [optional] 
**InvoiceType** | Pointer to [**TypesInvoiceType**](TypesInvoiceType.md) |  | [optional] 
**LineItems** | Pointer to [**[]DtoCreateInvoiceLineItemRequest**](DtoCreateInvoiceLineItemRequest.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PaymentStatus** | Pointer to [**TypesPaymentStatus**](TypesPaymentStatus.md) |  | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoCreateInvoiceRequest

`func NewDtoCreateInvoiceRequest(amountDue float32, currency string, customerId string, ) *DtoCreateInvoiceRequest`

NewDtoCreateInvoiceRequest instantiates a new DtoCreateInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateInvoiceRequestWithDefaults

`func NewDtoCreateInvoiceRequestWithDefaults() *DtoCreateInvoiceRequest`

NewDtoCreateInvoiceRequestWithDefaults instantiates a new DtoCreateInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDue

`func (o *DtoCreateInvoiceRequest) GetAmountDue() float32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *DtoCreateInvoiceRequest) GetAmountDueOk() (*float32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *DtoCreateInvoiceRequest) SetAmountDue(v float32)`

SetAmountDue sets AmountDue field to given value.


### GetAmountPaid

`func (o *DtoCreateInvoiceRequest) GetAmountPaid() float32`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *DtoCreateInvoiceRequest) GetAmountPaidOk() (*float32, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *DtoCreateInvoiceRequest) SetAmountPaid(v float32)`

SetAmountPaid sets AmountPaid field to given value.

### HasAmountPaid

`func (o *DtoCreateInvoiceRequest) HasAmountPaid() bool`

HasAmountPaid returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *DtoCreateInvoiceRequest) GetBillingPeriod() string`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoCreateInvoiceRequest) GetBillingPeriodOk() (*string, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoCreateInvoiceRequest) SetBillingPeriod(v string)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *DtoCreateInvoiceRequest) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetBillingReason

`func (o *DtoCreateInvoiceRequest) GetBillingReason() TypesInvoiceBillingReason`

GetBillingReason returns the BillingReason field if non-nil, zero value otherwise.

### GetBillingReasonOk

`func (o *DtoCreateInvoiceRequest) GetBillingReasonOk() (*TypesInvoiceBillingReason, bool)`

GetBillingReasonOk returns a tuple with the BillingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingReason

`func (o *DtoCreateInvoiceRequest) SetBillingReason(v TypesInvoiceBillingReason)`

SetBillingReason sets BillingReason field to given value.

### HasBillingReason

`func (o *DtoCreateInvoiceRequest) HasBillingReason() bool`

HasBillingReason returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoCreateInvoiceRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCreateInvoiceRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCreateInvoiceRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerId

`func (o *DtoCreateInvoiceRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoCreateInvoiceRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoCreateInvoiceRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetDescription

`func (o *DtoCreateInvoiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoCreateInvoiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoCreateInvoiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoCreateInvoiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *DtoCreateInvoiceRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *DtoCreateInvoiceRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *DtoCreateInvoiceRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *DtoCreateInvoiceRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoCreateInvoiceRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoCreateInvoiceRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoCreateInvoiceRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoCreateInvoiceRequest) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *DtoCreateInvoiceRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoCreateInvoiceRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoCreateInvoiceRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *DtoCreateInvoiceRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *DtoCreateInvoiceRequest) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *DtoCreateInvoiceRequest) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *DtoCreateInvoiceRequest) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *DtoCreateInvoiceRequest) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *DtoCreateInvoiceRequest) GetInvoiceStatus() TypesInvoiceStatus`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *DtoCreateInvoiceRequest) GetInvoiceStatusOk() (*TypesInvoiceStatus, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *DtoCreateInvoiceRequest) SetInvoiceStatus(v TypesInvoiceStatus)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *DtoCreateInvoiceRequest) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetInvoiceType

`func (o *DtoCreateInvoiceRequest) GetInvoiceType() TypesInvoiceType`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *DtoCreateInvoiceRequest) GetInvoiceTypeOk() (*TypesInvoiceType, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *DtoCreateInvoiceRequest) SetInvoiceType(v TypesInvoiceType)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *DtoCreateInvoiceRequest) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetLineItems

`func (o *DtoCreateInvoiceRequest) GetLineItems() []DtoCreateInvoiceLineItemRequest`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *DtoCreateInvoiceRequest) GetLineItemsOk() (*[]DtoCreateInvoiceLineItemRequest, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *DtoCreateInvoiceRequest) SetLineItems(v []DtoCreateInvoiceLineItemRequest)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *DtoCreateInvoiceRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreateInvoiceRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateInvoiceRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateInvoiceRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateInvoiceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *DtoCreateInvoiceRequest) GetPaymentStatus() TypesPaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *DtoCreateInvoiceRequest) GetPaymentStatusOk() (*TypesPaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *DtoCreateInvoiceRequest) SetPaymentStatus(v TypesPaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *DtoCreateInvoiceRequest) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *DtoCreateInvoiceRequest) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *DtoCreateInvoiceRequest) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *DtoCreateInvoiceRequest) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *DtoCreateInvoiceRequest) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *DtoCreateInvoiceRequest) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *DtoCreateInvoiceRequest) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *DtoCreateInvoiceRequest) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *DtoCreateInvoiceRequest) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoCreateInvoiceRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoCreateInvoiceRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoCreateInvoiceRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoCreateInvoiceRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


