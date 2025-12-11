# DtoCreditNoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreditNoteNumber** | Pointer to **string** | credit_note_number is the unique identifier for credit notes | [optional] 
**CreditNoteStatus** | Pointer to [**TypesCreditNoteStatus**](TypesCreditNoteStatus.md) |  | [optional] 
**CreditNoteType** | Pointer to [**TypesCreditNoteType**](TypesCreditNoteType.md) |  | [optional] 
**Currency** | Pointer to **string** | currency is the three-letter ISO currency code (e.g., USD, EUR) for the credit note | [optional] 
**Customer** | Pointer to [**GithubComFlexpriceFlexpriceInternalDomainCustomerCustomer**](GithubComFlexpriceFlexpriceInternalDomainCustomerCustomer.md) |  | [optional] 
**CustomerId** | Pointer to **string** | customer_id is the unique identifier of the customer who owns this credit note | [optional] 
**EnvironmentId** | Pointer to **string** | environment_id is the unique identifier of the environment this credit note belongs to | [optional] 
**FinalizedAt** | Pointer to **string** | finalized_at is the timestamp when the credit note was finalized | [optional] 
**Id** | Pointer to **string** | id is the unique identifier for the credit note | [optional] 
**IdempotencyKey** | Pointer to **string** | idempotency_key is an optional key used to prevent duplicate credit note creation | [optional] 
**Invoice** | Pointer to [**DtoInvoiceResponse**](DtoInvoiceResponse.md) |  | [optional] 
**InvoiceId** | Pointer to **string** | invoice_id is the id of the invoice resource that this credit note is applied to | [optional] 
**LineItems** | Pointer to [**[]CreditnoteCreditNoteLineItem**](CreditnoteCreditNoteLineItem.md) | line_items contains all of the line items associated with this credit note | [optional] 
**Memo** | Pointer to **string** | memo is an optional memo supplied on the credit note | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Reason** | Pointer to [**TypesCreditNoteReason**](TypesCreditNoteReason.md) |  | [optional] 
**RefundStatus** | Pointer to [**TypesPaymentStatus**](TypesPaymentStatus.md) |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**Subscription** | Pointer to [**DtoSubscriptionResponse**](DtoSubscriptionResponse.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id is the optional unique identifier of the subscription related to this credit note | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TotalAmount** | Pointer to **float32** | total_amount is the total including creditable invoice-level discounts or minimums, and tax | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**VoidedAt** | Pointer to **string** | voided_at is the timestamp when the credit note was voided | [optional] 

## Methods

### NewDtoCreditNoteResponse

`func NewDtoCreditNoteResponse() *DtoCreditNoteResponse`

NewDtoCreditNoteResponse instantiates a new DtoCreditNoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreditNoteResponseWithDefaults

`func NewDtoCreditNoteResponseWithDefaults() *DtoCreditNoteResponse`

NewDtoCreditNoteResponseWithDefaults instantiates a new DtoCreditNoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DtoCreditNoteResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoCreditNoteResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoCreditNoteResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoCreditNoteResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoCreditNoteResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoCreditNoteResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoCreditNoteResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoCreditNoteResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreditNoteNumber

`func (o *DtoCreditNoteResponse) GetCreditNoteNumber() string`

GetCreditNoteNumber returns the CreditNoteNumber field if non-nil, zero value otherwise.

### GetCreditNoteNumberOk

`func (o *DtoCreditNoteResponse) GetCreditNoteNumberOk() (*string, bool)`

GetCreditNoteNumberOk returns a tuple with the CreditNoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteNumber

`func (o *DtoCreditNoteResponse) SetCreditNoteNumber(v string)`

SetCreditNoteNumber sets CreditNoteNumber field to given value.

### HasCreditNoteNumber

`func (o *DtoCreditNoteResponse) HasCreditNoteNumber() bool`

HasCreditNoteNumber returns a boolean if a field has been set.

### GetCreditNoteStatus

`func (o *DtoCreditNoteResponse) GetCreditNoteStatus() TypesCreditNoteStatus`

GetCreditNoteStatus returns the CreditNoteStatus field if non-nil, zero value otherwise.

### GetCreditNoteStatusOk

`func (o *DtoCreditNoteResponse) GetCreditNoteStatusOk() (*TypesCreditNoteStatus, bool)`

GetCreditNoteStatusOk returns a tuple with the CreditNoteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteStatus

`func (o *DtoCreditNoteResponse) SetCreditNoteStatus(v TypesCreditNoteStatus)`

SetCreditNoteStatus sets CreditNoteStatus field to given value.

### HasCreditNoteStatus

`func (o *DtoCreditNoteResponse) HasCreditNoteStatus() bool`

HasCreditNoteStatus returns a boolean if a field has been set.

### GetCreditNoteType

`func (o *DtoCreditNoteResponse) GetCreditNoteType() TypesCreditNoteType`

GetCreditNoteType returns the CreditNoteType field if non-nil, zero value otherwise.

### GetCreditNoteTypeOk

`func (o *DtoCreditNoteResponse) GetCreditNoteTypeOk() (*TypesCreditNoteType, bool)`

GetCreditNoteTypeOk returns a tuple with the CreditNoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditNoteType

`func (o *DtoCreditNoteResponse) SetCreditNoteType(v TypesCreditNoteType)`

SetCreditNoteType sets CreditNoteType field to given value.

### HasCreditNoteType

`func (o *DtoCreditNoteResponse) HasCreditNoteType() bool`

HasCreditNoteType returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoCreditNoteResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoCreditNoteResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoCreditNoteResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoCreditNoteResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomer

`func (o *DtoCreditNoteResponse) GetCustomer() GithubComFlexpriceFlexpriceInternalDomainCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *DtoCreditNoteResponse) GetCustomerOk() (*GithubComFlexpriceFlexpriceInternalDomainCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *DtoCreditNoteResponse) SetCustomer(v GithubComFlexpriceFlexpriceInternalDomainCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *DtoCreditNoteResponse) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoCreditNoteResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoCreditNoteResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoCreditNoteResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoCreditNoteResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoCreditNoteResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoCreditNoteResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoCreditNoteResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoCreditNoteResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetFinalizedAt

`func (o *DtoCreditNoteResponse) GetFinalizedAt() string`

GetFinalizedAt returns the FinalizedAt field if non-nil, zero value otherwise.

### GetFinalizedAtOk

`func (o *DtoCreditNoteResponse) GetFinalizedAtOk() (*string, bool)`

GetFinalizedAtOk returns a tuple with the FinalizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedAt

`func (o *DtoCreditNoteResponse) SetFinalizedAt(v string)`

SetFinalizedAt sets FinalizedAt field to given value.

### HasFinalizedAt

`func (o *DtoCreditNoteResponse) HasFinalizedAt() bool`

HasFinalizedAt returns a boolean if a field has been set.

### GetId

`func (o *DtoCreditNoteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoCreditNoteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoCreditNoteResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoCreditNoteResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *DtoCreditNoteResponse) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoCreditNoteResponse) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoCreditNoteResponse) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *DtoCreditNoteResponse) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetInvoice

`func (o *DtoCreditNoteResponse) GetInvoice() DtoInvoiceResponse`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *DtoCreditNoteResponse) GetInvoiceOk() (*DtoInvoiceResponse, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *DtoCreditNoteResponse) SetInvoice(v DtoInvoiceResponse)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *DtoCreditNoteResponse) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetInvoiceId

`func (o *DtoCreditNoteResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DtoCreditNoteResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DtoCreditNoteResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *DtoCreditNoteResponse) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetLineItems

`func (o *DtoCreditNoteResponse) GetLineItems() []CreditnoteCreditNoteLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *DtoCreditNoteResponse) GetLineItemsOk() (*[]CreditnoteCreditNoteLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *DtoCreditNoteResponse) SetLineItems(v []CreditnoteCreditNoteLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *DtoCreditNoteResponse) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMemo

`func (o *DtoCreditNoteResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *DtoCreditNoteResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *DtoCreditNoteResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *DtoCreditNoteResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoCreditNoteResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreditNoteResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreditNoteResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreditNoteResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReason

`func (o *DtoCreditNoteResponse) GetReason() TypesCreditNoteReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DtoCreditNoteResponse) GetReasonOk() (*TypesCreditNoteReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DtoCreditNoteResponse) SetReason(v TypesCreditNoteReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *DtoCreditNoteResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRefundStatus

`func (o *DtoCreditNoteResponse) GetRefundStatus() TypesPaymentStatus`

GetRefundStatus returns the RefundStatus field if non-nil, zero value otherwise.

### GetRefundStatusOk

`func (o *DtoCreditNoteResponse) GetRefundStatusOk() (*TypesPaymentStatus, bool)`

GetRefundStatusOk returns a tuple with the RefundStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundStatus

`func (o *DtoCreditNoteResponse) SetRefundStatus(v TypesPaymentStatus)`

SetRefundStatus sets RefundStatus field to given value.

### HasRefundStatus

`func (o *DtoCreditNoteResponse) HasRefundStatus() bool`

HasRefundStatus returns a boolean if a field has been set.

### GetStatus

`func (o *DtoCreditNoteResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoCreditNoteResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoCreditNoteResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoCreditNoteResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscription

`func (o *DtoCreditNoteResponse) GetSubscription() DtoSubscriptionResponse`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *DtoCreditNoteResponse) GetSubscriptionOk() (*DtoSubscriptionResponse, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *DtoCreditNoteResponse) SetSubscription(v DtoSubscriptionResponse)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *DtoCreditNoteResponse) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoCreditNoteResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoCreditNoteResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoCreditNoteResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoCreditNoteResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoCreditNoteResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoCreditNoteResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoCreditNoteResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoCreditNoteResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTotalAmount

`func (o *DtoCreditNoteResponse) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *DtoCreditNoteResponse) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *DtoCreditNoteResponse) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *DtoCreditNoteResponse) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoCreditNoteResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoCreditNoteResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoCreditNoteResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoCreditNoteResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoCreditNoteResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoCreditNoteResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoCreditNoteResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoCreditNoteResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVoidedAt

`func (o *DtoCreditNoteResponse) GetVoidedAt() string`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *DtoCreditNoteResponse) GetVoidedAtOk() (*string, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *DtoCreditNoteResponse) SetVoidedAt(v string)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *DtoCreditNoteResponse) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


