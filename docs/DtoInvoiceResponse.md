# DtoInvoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDue** | Pointer to **string** | amount_due is the total amount that needs to be paid for this invoice | [optional] 
**AmountPaid** | Pointer to **string** | amount_paid is the amount that has been paid towards this invoice | [optional] 
**AmountRemaining** | Pointer to **string** | amount_remaining is the amount still outstanding on this invoice | [optional] 
**BillingPeriod** | Pointer to **string** | billing_period is the period this invoice covers (e.g., \&quot;monthly\&quot;, \&quot;yearly\&quot;) | [optional] 
**BillingReason** | Pointer to **string** | billing_reason indicates why this invoice was created (subscription_cycle, manual, etc.) | [optional] 
**BillingSequence** | Pointer to **int32** | billing_sequence is the optional sequence number for billing cycles | [optional] 
**CouponApplications** | Pointer to [**[]DtoCouponApplicationResponse**](DtoCouponApplicationResponse.md) | coupon_applications contains the coupon applications associated with this invoice | [optional] 
**CreatedAt** | Pointer to **string** | created_at is the timestamp when this invoice was created | [optional] 
**CreatedBy** | Pointer to **string** | created_by is the identifier of the user who created this invoice | [optional] 
**Currency** | Pointer to **string** | currency is the three-letter ISO currency code (e.g., USD, EUR) for the invoice | [optional] 
**Customer** | Pointer to [**DtoCustomerResponse**](DtoCustomerResponse.md) |  | [optional] 
**CustomerId** | Pointer to **string** | customer_id is the unique identifier of the customer this invoice belongs to | [optional] 
**Description** | Pointer to **string** | description is the optional text description of the invoice | [optional] 
**DueDate** | Pointer to **string** | due_date is the date by which payment is expected | [optional] 
**FinalizedAt** | Pointer to **string** | finalized_at is the timestamp when this invoice was finalized | [optional] 
**Id** | Pointer to **string** | id is the unique identifier for this invoice | [optional] 
**IdempotencyKey** | Pointer to **string** | idempotency_key is the optional key used to prevent duplicate invoice creation | [optional] 
**InvoiceNumber** | Pointer to **string** | invoice_number is the optional human-readable identifier for the invoice | [optional] 
**InvoicePdfUrl** | Pointer to **string** | invoice_pdf_url is the optional URL to the PDF version of this invoice | [optional] 
**InvoiceStatus** | Pointer to [**TypesInvoiceStatus**](TypesInvoiceStatus.md) |  | [optional] 
**InvoiceType** | Pointer to [**TypesInvoiceType**](TypesInvoiceType.md) |  | [optional] 
**LineItems** | Pointer to [**[]DtoInvoiceLineItemResponse**](DtoInvoiceLineItemResponse.md) | line_items contains the individual items that make up this invoice | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**OverpaidAmount** | Pointer to **string** | overpaid_amount is the amount overpaid if payment_status is OVERPAID (amount_paid - total) | [optional] 
**PaidAt** | Pointer to **string** | paid_at is the timestamp when this invoice was paid | [optional] 
**PaymentStatus** | Pointer to [**TypesPaymentStatus**](TypesPaymentStatus.md) |  | [optional] 
**PeriodEnd** | Pointer to **string** | period_end is the end date of the billing period | [optional] 
**PeriodStart** | Pointer to **string** | period_start is the start date of the billing period | [optional] 
**Status** | Pointer to **string** | status represents the current status of this invoice | [optional] 
**Subscription** | Pointer to [**DtoSubscriptionResponse**](DtoSubscriptionResponse.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id is the optional unique identifier of the subscription associated with this invoice | [optional] 
**Subtotal** | Pointer to **string** | subtotal is the amount before taxes and discounts are applied | [optional] 
**Taxes** | Pointer to [**[]DtoTaxAppliedResponse**](DtoTaxAppliedResponse.md) | tax_applied_records contains the tax applied records associated with this invoice | [optional] 
**TenantId** | Pointer to **string** | tenant_id is the unique identifier of the tenant this invoice belongs to | [optional] 
**Total** | Pointer to **string** | total is the total amount of the invoice including taxes and discounts | [optional] 
**TotalDiscount** | Pointer to **string** | total_discount is the total discount amount from coupon applications | [optional] 
**TotalTax** | Pointer to **string** | total_tax is the total tax amount for this invoice | [optional] 
**UpdatedAt** | Pointer to **string** | updated_at is the timestamp when this invoice was last updated | [optional] 
**UpdatedBy** | Pointer to **string** | updated_by is the identifier of the user who last updated this invoice | [optional] 
**Version** | Pointer to **int32** | version is the version number of this invoice | [optional] 
**VoidedAt** | Pointer to **string** | voided_at is the timestamp when this invoice was voided | [optional] 

## Methods

### NewDtoInvoiceResponse

`func NewDtoInvoiceResponse() *DtoInvoiceResponse`

NewDtoInvoiceResponse instantiates a new DtoInvoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoInvoiceResponseWithDefaults

`func NewDtoInvoiceResponseWithDefaults() *DtoInvoiceResponse`

NewDtoInvoiceResponseWithDefaults instantiates a new DtoInvoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDue

`func (o *DtoInvoiceResponse) GetAmountDue() string`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *DtoInvoiceResponse) GetAmountDueOk() (*string, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *DtoInvoiceResponse) SetAmountDue(v string)`

SetAmountDue sets AmountDue field to given value.

### HasAmountDue

`func (o *DtoInvoiceResponse) HasAmountDue() bool`

HasAmountDue returns a boolean if a field has been set.

### GetAmountPaid

`func (o *DtoInvoiceResponse) GetAmountPaid() string`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *DtoInvoiceResponse) GetAmountPaidOk() (*string, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *DtoInvoiceResponse) SetAmountPaid(v string)`

SetAmountPaid sets AmountPaid field to given value.

### HasAmountPaid

`func (o *DtoInvoiceResponse) HasAmountPaid() bool`

HasAmountPaid returns a boolean if a field has been set.

### GetAmountRemaining

`func (o *DtoInvoiceResponse) GetAmountRemaining() string`

GetAmountRemaining returns the AmountRemaining field if non-nil, zero value otherwise.

### GetAmountRemainingOk

`func (o *DtoInvoiceResponse) GetAmountRemainingOk() (*string, bool)`

GetAmountRemainingOk returns a tuple with the AmountRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRemaining

`func (o *DtoInvoiceResponse) SetAmountRemaining(v string)`

SetAmountRemaining sets AmountRemaining field to given value.

### HasAmountRemaining

`func (o *DtoInvoiceResponse) HasAmountRemaining() bool`

HasAmountRemaining returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *DtoInvoiceResponse) GetBillingPeriod() string`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *DtoInvoiceResponse) GetBillingPeriodOk() (*string, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *DtoInvoiceResponse) SetBillingPeriod(v string)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *DtoInvoiceResponse) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetBillingReason

`func (o *DtoInvoiceResponse) GetBillingReason() string`

GetBillingReason returns the BillingReason field if non-nil, zero value otherwise.

### GetBillingReasonOk

`func (o *DtoInvoiceResponse) GetBillingReasonOk() (*string, bool)`

GetBillingReasonOk returns a tuple with the BillingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingReason

`func (o *DtoInvoiceResponse) SetBillingReason(v string)`

SetBillingReason sets BillingReason field to given value.

### HasBillingReason

`func (o *DtoInvoiceResponse) HasBillingReason() bool`

HasBillingReason returns a boolean if a field has been set.

### GetBillingSequence

`func (o *DtoInvoiceResponse) GetBillingSequence() int32`

GetBillingSequence returns the BillingSequence field if non-nil, zero value otherwise.

### GetBillingSequenceOk

`func (o *DtoInvoiceResponse) GetBillingSequenceOk() (*int32, bool)`

GetBillingSequenceOk returns a tuple with the BillingSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSequence

`func (o *DtoInvoiceResponse) SetBillingSequence(v int32)`

SetBillingSequence sets BillingSequence field to given value.

### HasBillingSequence

`func (o *DtoInvoiceResponse) HasBillingSequence() bool`

HasBillingSequence returns a boolean if a field has been set.

### GetCouponApplications

`func (o *DtoInvoiceResponse) GetCouponApplications() []DtoCouponApplicationResponse`

GetCouponApplications returns the CouponApplications field if non-nil, zero value otherwise.

### GetCouponApplicationsOk

`func (o *DtoInvoiceResponse) GetCouponApplicationsOk() (*[]DtoCouponApplicationResponse, bool)`

GetCouponApplicationsOk returns a tuple with the CouponApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponApplications

`func (o *DtoInvoiceResponse) SetCouponApplications(v []DtoCouponApplicationResponse)`

SetCouponApplications sets CouponApplications field to given value.

### HasCouponApplications

`func (o *DtoInvoiceResponse) HasCouponApplications() bool`

HasCouponApplications returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoInvoiceResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoInvoiceResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoInvoiceResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoInvoiceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoInvoiceResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoInvoiceResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoInvoiceResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoInvoiceResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoInvoiceResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoInvoiceResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoInvoiceResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoInvoiceResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomer

`func (o *DtoInvoiceResponse) GetCustomer() DtoCustomerResponse`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *DtoInvoiceResponse) GetCustomerOk() (*DtoCustomerResponse, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *DtoInvoiceResponse) SetCustomer(v DtoCustomerResponse)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *DtoInvoiceResponse) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoInvoiceResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoInvoiceResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoInvoiceResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoInvoiceResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *DtoInvoiceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoInvoiceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoInvoiceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoInvoiceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *DtoInvoiceResponse) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *DtoInvoiceResponse) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *DtoInvoiceResponse) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *DtoInvoiceResponse) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetFinalizedAt

`func (o *DtoInvoiceResponse) GetFinalizedAt() string`

GetFinalizedAt returns the FinalizedAt field if non-nil, zero value otherwise.

### GetFinalizedAtOk

`func (o *DtoInvoiceResponse) GetFinalizedAtOk() (*string, bool)`

GetFinalizedAtOk returns a tuple with the FinalizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizedAt

`func (o *DtoInvoiceResponse) SetFinalizedAt(v string)`

SetFinalizedAt sets FinalizedAt field to given value.

### HasFinalizedAt

`func (o *DtoInvoiceResponse) HasFinalizedAt() bool`

HasFinalizedAt returns a boolean if a field has been set.

### GetId

`func (o *DtoInvoiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoInvoiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoInvoiceResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoInvoiceResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *DtoInvoiceResponse) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoInvoiceResponse) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoInvoiceResponse) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *DtoInvoiceResponse) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *DtoInvoiceResponse) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *DtoInvoiceResponse) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *DtoInvoiceResponse) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *DtoInvoiceResponse) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoicePdfUrl

`func (o *DtoInvoiceResponse) GetInvoicePdfUrl() string`

GetInvoicePdfUrl returns the InvoicePdfUrl field if non-nil, zero value otherwise.

### GetInvoicePdfUrlOk

`func (o *DtoInvoiceResponse) GetInvoicePdfUrlOk() (*string, bool)`

GetInvoicePdfUrlOk returns a tuple with the InvoicePdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdfUrl

`func (o *DtoInvoiceResponse) SetInvoicePdfUrl(v string)`

SetInvoicePdfUrl sets InvoicePdfUrl field to given value.

### HasInvoicePdfUrl

`func (o *DtoInvoiceResponse) HasInvoicePdfUrl() bool`

HasInvoicePdfUrl returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *DtoInvoiceResponse) GetInvoiceStatus() TypesInvoiceStatus`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *DtoInvoiceResponse) GetInvoiceStatusOk() (*TypesInvoiceStatus, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *DtoInvoiceResponse) SetInvoiceStatus(v TypesInvoiceStatus)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *DtoInvoiceResponse) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetInvoiceType

`func (o *DtoInvoiceResponse) GetInvoiceType() TypesInvoiceType`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *DtoInvoiceResponse) GetInvoiceTypeOk() (*TypesInvoiceType, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *DtoInvoiceResponse) SetInvoiceType(v TypesInvoiceType)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *DtoInvoiceResponse) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetLineItems

`func (o *DtoInvoiceResponse) GetLineItems() []DtoInvoiceLineItemResponse`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *DtoInvoiceResponse) GetLineItemsOk() (*[]DtoInvoiceLineItemResponse, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *DtoInvoiceResponse) SetLineItems(v []DtoInvoiceLineItemResponse)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *DtoInvoiceResponse) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoInvoiceResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoInvoiceResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoInvoiceResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoInvoiceResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOverpaidAmount

`func (o *DtoInvoiceResponse) GetOverpaidAmount() string`

GetOverpaidAmount returns the OverpaidAmount field if non-nil, zero value otherwise.

### GetOverpaidAmountOk

`func (o *DtoInvoiceResponse) GetOverpaidAmountOk() (*string, bool)`

GetOverpaidAmountOk returns a tuple with the OverpaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverpaidAmount

`func (o *DtoInvoiceResponse) SetOverpaidAmount(v string)`

SetOverpaidAmount sets OverpaidAmount field to given value.

### HasOverpaidAmount

`func (o *DtoInvoiceResponse) HasOverpaidAmount() bool`

HasOverpaidAmount returns a boolean if a field has been set.

### GetPaidAt

`func (o *DtoInvoiceResponse) GetPaidAt() string`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *DtoInvoiceResponse) GetPaidAtOk() (*string, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *DtoInvoiceResponse) SetPaidAt(v string)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *DtoInvoiceResponse) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *DtoInvoiceResponse) GetPaymentStatus() TypesPaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *DtoInvoiceResponse) GetPaymentStatusOk() (*TypesPaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *DtoInvoiceResponse) SetPaymentStatus(v TypesPaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *DtoInvoiceResponse) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *DtoInvoiceResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *DtoInvoiceResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *DtoInvoiceResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *DtoInvoiceResponse) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *DtoInvoiceResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *DtoInvoiceResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *DtoInvoiceResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *DtoInvoiceResponse) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetStatus

`func (o *DtoInvoiceResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoInvoiceResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoInvoiceResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoInvoiceResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscription

`func (o *DtoInvoiceResponse) GetSubscription() DtoSubscriptionResponse`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *DtoInvoiceResponse) GetSubscriptionOk() (*DtoSubscriptionResponse, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *DtoInvoiceResponse) SetSubscription(v DtoSubscriptionResponse)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *DtoInvoiceResponse) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoInvoiceResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoInvoiceResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoInvoiceResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoInvoiceResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetSubtotal

`func (o *DtoInvoiceResponse) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *DtoInvoiceResponse) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *DtoInvoiceResponse) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *DtoInvoiceResponse) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetTaxes

`func (o *DtoInvoiceResponse) GetTaxes() []DtoTaxAppliedResponse`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *DtoInvoiceResponse) GetTaxesOk() (*[]DtoTaxAppliedResponse, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *DtoInvoiceResponse) SetTaxes(v []DtoTaxAppliedResponse)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *DtoInvoiceResponse) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoInvoiceResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoInvoiceResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoInvoiceResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoInvoiceResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTotal

`func (o *DtoInvoiceResponse) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DtoInvoiceResponse) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DtoInvoiceResponse) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DtoInvoiceResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *DtoInvoiceResponse) GetTotalDiscount() string`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *DtoInvoiceResponse) GetTotalDiscountOk() (*string, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *DtoInvoiceResponse) SetTotalDiscount(v string)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *DtoInvoiceResponse) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### GetTotalTax

`func (o *DtoInvoiceResponse) GetTotalTax() string`

GetTotalTax returns the TotalTax field if non-nil, zero value otherwise.

### GetTotalTaxOk

`func (o *DtoInvoiceResponse) GetTotalTaxOk() (*string, bool)`

GetTotalTaxOk returns a tuple with the TotalTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTax

`func (o *DtoInvoiceResponse) SetTotalTax(v string)`

SetTotalTax sets TotalTax field to given value.

### HasTotalTax

`func (o *DtoInvoiceResponse) HasTotalTax() bool`

HasTotalTax returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoInvoiceResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoInvoiceResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoInvoiceResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoInvoiceResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoInvoiceResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoInvoiceResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoInvoiceResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoInvoiceResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVersion

`func (o *DtoInvoiceResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DtoInvoiceResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DtoInvoiceResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DtoInvoiceResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVoidedAt

`func (o *DtoInvoiceResponse) GetVoidedAt() string`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *DtoInvoiceResponse) GetVoidedAtOk() (*string, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *DtoInvoiceResponse) SetVoidedAt(v string)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *DtoInvoiceResponse) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


