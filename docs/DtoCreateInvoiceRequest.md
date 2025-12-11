# DtoCreateInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDue** | **string** | amount_due is the total amount that needs to be paid for this invoice | 
**AmountPaid** | Pointer to **string** | amount_paid is the amount that has been paid towards this invoice | [optional] 
**BillingPeriod** | Pointer to **string** | billing_period is the period this invoice covers (e.g., \&quot;monthly\&quot;, \&quot;yearly\&quot;) | [optional] 
**BillingReason** | Pointer to [**TypesInvoiceBillingReason**](TypesInvoiceBillingReason.md) |  | [optional] 
**Coupons** | Pointer to **[]string** | coupons | [optional] 
**Currency** | **string** | currency is the three-letter ISO currency code (e.g., USD, EUR) for the invoice | 
**CustomerId** | **string** | customer_id is the unique identifier of the customer this invoice belongs to | 
**Description** | Pointer to **string** | description is an optional text description of the invoice | [optional] 
**DueDate** | Pointer to **string** | due_date is the date by which payment is expected | [optional] 
**EnvironmentId** | Pointer to **string** | environment_id is the unique identifier of the environment this invoice belongs to | [optional] 
**IdempotencyKey** | Pointer to **string** | idempotency_key is an optional key used to prevent duplicate invoice creation | [optional] 
**InvoiceCoupons** | Pointer to [**[]DtoInvoiceCoupon**](DtoInvoiceCoupon.md) | Invoice Coupns | [optional] 
**InvoiceNumber** | Pointer to **string** | invoice_number is an optional human-readable identifier for the invoice | [optional] 
**InvoicePdfUrl** | Pointer to **string** | invoice_pdf_url is the URL where customers can download the PDF version of this invoice | [optional] 
**InvoiceStatus** | Pointer to [**TypesInvoiceStatus**](TypesInvoiceStatus.md) |  | [optional] 
**InvoiceType** | Pointer to [**TypesInvoiceType**](TypesInvoiceType.md) |  | [optional] 
**LineItemCoupons** | Pointer to [**[]DtoInvoiceLineItemCoupon**](DtoInvoiceLineItemCoupon.md) | Invoice Line Item Coupons | [optional] 
**LineItems** | Pointer to [**[]DtoCreateInvoiceLineItemRequest**](DtoCreateInvoiceLineItemRequest.md) | line_items contains the individual items that make up this invoice | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PaymentStatus** | Pointer to [**TypesPaymentStatus**](TypesPaymentStatus.md) |  | [optional] 
**PeriodEnd** | Pointer to **string** | period_end is the end date of the billing period | [optional] 
**PeriodStart** | Pointer to **string** | period_start is the start date of the billing period | [optional] 
**PreparedTaxRates** | Pointer to [**[]DtoTaxRateResponse**](DtoTaxRateResponse.md) | prepared_tax_rates contains the tax rates pre-resolved by the caller (e.g., billing service) These are applied at invoice level by the invoice service without further resolution | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id is the optional unique identifier of the subscription associated with this invoice | [optional] 
**Subtotal** | **string** | subtotal is the amount before taxes and discounts are applied | 
**TaxRateOverrides** | Pointer to [**[]DtoTaxRateOverride**](DtoTaxRateOverride.md) | tax_rate_overrides is the tax rate overrides to be applied to the invoice | [optional] 
**TaxRates** | Pointer to **[]string** | tax_rates | [optional] 
**Total** | **string** | total is the total amount of the invoice including taxes and discounts | 

## Methods

### NewDtoCreateInvoiceRequest

`func NewDtoCreateInvoiceRequest(amountDue string, currency string, customerId string, subtotal string, total string, ) *DtoCreateInvoiceRequest`

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

`func (o *DtoCreateInvoiceRequest) GetAmountDue() string`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *DtoCreateInvoiceRequest) GetAmountDueOk() (*string, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *DtoCreateInvoiceRequest) SetAmountDue(v string)`

SetAmountDue sets AmountDue field to given value.


### GetAmountPaid

`func (o *DtoCreateInvoiceRequest) GetAmountPaid() string`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *DtoCreateInvoiceRequest) GetAmountPaidOk() (*string, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *DtoCreateInvoiceRequest) SetAmountPaid(v string)`

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

### GetCoupons

`func (o *DtoCreateInvoiceRequest) GetCoupons() []string`

GetCoupons returns the Coupons field if non-nil, zero value otherwise.

### GetCouponsOk

`func (o *DtoCreateInvoiceRequest) GetCouponsOk() (*[]string, bool)`

GetCouponsOk returns a tuple with the Coupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupons

`func (o *DtoCreateInvoiceRequest) SetCoupons(v []string)`

SetCoupons sets Coupons field to given value.

### HasCoupons

`func (o *DtoCreateInvoiceRequest) HasCoupons() bool`

HasCoupons returns a boolean if a field has been set.

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

### GetInvoiceCoupons

`func (o *DtoCreateInvoiceRequest) GetInvoiceCoupons() []DtoInvoiceCoupon`

GetInvoiceCoupons returns the InvoiceCoupons field if non-nil, zero value otherwise.

### GetInvoiceCouponsOk

`func (o *DtoCreateInvoiceRequest) GetInvoiceCouponsOk() (*[]DtoInvoiceCoupon, bool)`

GetInvoiceCouponsOk returns a tuple with the InvoiceCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCoupons

`func (o *DtoCreateInvoiceRequest) SetInvoiceCoupons(v []DtoInvoiceCoupon)`

SetInvoiceCoupons sets InvoiceCoupons field to given value.

### HasInvoiceCoupons

`func (o *DtoCreateInvoiceRequest) HasInvoiceCoupons() bool`

HasInvoiceCoupons returns a boolean if a field has been set.

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

### GetInvoicePdfUrl

`func (o *DtoCreateInvoiceRequest) GetInvoicePdfUrl() string`

GetInvoicePdfUrl returns the InvoicePdfUrl field if non-nil, zero value otherwise.

### GetInvoicePdfUrlOk

`func (o *DtoCreateInvoiceRequest) GetInvoicePdfUrlOk() (*string, bool)`

GetInvoicePdfUrlOk returns a tuple with the InvoicePdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdfUrl

`func (o *DtoCreateInvoiceRequest) SetInvoicePdfUrl(v string)`

SetInvoicePdfUrl sets InvoicePdfUrl field to given value.

### HasInvoicePdfUrl

`func (o *DtoCreateInvoiceRequest) HasInvoicePdfUrl() bool`

HasInvoicePdfUrl returns a boolean if a field has been set.

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

### GetLineItemCoupons

`func (o *DtoCreateInvoiceRequest) GetLineItemCoupons() []DtoInvoiceLineItemCoupon`

GetLineItemCoupons returns the LineItemCoupons field if non-nil, zero value otherwise.

### GetLineItemCouponsOk

`func (o *DtoCreateInvoiceRequest) GetLineItemCouponsOk() (*[]DtoInvoiceLineItemCoupon, bool)`

GetLineItemCouponsOk returns a tuple with the LineItemCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItemCoupons

`func (o *DtoCreateInvoiceRequest) SetLineItemCoupons(v []DtoInvoiceLineItemCoupon)`

SetLineItemCoupons sets LineItemCoupons field to given value.

### HasLineItemCoupons

`func (o *DtoCreateInvoiceRequest) HasLineItemCoupons() bool`

HasLineItemCoupons returns a boolean if a field has been set.

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

### GetPreparedTaxRates

`func (o *DtoCreateInvoiceRequest) GetPreparedTaxRates() []DtoTaxRateResponse`

GetPreparedTaxRates returns the PreparedTaxRates field if non-nil, zero value otherwise.

### GetPreparedTaxRatesOk

`func (o *DtoCreateInvoiceRequest) GetPreparedTaxRatesOk() (*[]DtoTaxRateResponse, bool)`

GetPreparedTaxRatesOk returns a tuple with the PreparedTaxRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparedTaxRates

`func (o *DtoCreateInvoiceRequest) SetPreparedTaxRates(v []DtoTaxRateResponse)`

SetPreparedTaxRates sets PreparedTaxRates field to given value.

### HasPreparedTaxRates

`func (o *DtoCreateInvoiceRequest) HasPreparedTaxRates() bool`

HasPreparedTaxRates returns a boolean if a field has been set.

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

### GetSubtotal

`func (o *DtoCreateInvoiceRequest) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *DtoCreateInvoiceRequest) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *DtoCreateInvoiceRequest) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.


### GetTaxRateOverrides

`func (o *DtoCreateInvoiceRequest) GetTaxRateOverrides() []DtoTaxRateOverride`

GetTaxRateOverrides returns the TaxRateOverrides field if non-nil, zero value otherwise.

### GetTaxRateOverridesOk

`func (o *DtoCreateInvoiceRequest) GetTaxRateOverridesOk() (*[]DtoTaxRateOverride, bool)`

GetTaxRateOverridesOk returns a tuple with the TaxRateOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRateOverrides

`func (o *DtoCreateInvoiceRequest) SetTaxRateOverrides(v []DtoTaxRateOverride)`

SetTaxRateOverrides sets TaxRateOverrides field to given value.

### HasTaxRateOverrides

`func (o *DtoCreateInvoiceRequest) HasTaxRateOverrides() bool`

HasTaxRateOverrides returns a boolean if a field has been set.

### GetTaxRates

`func (o *DtoCreateInvoiceRequest) GetTaxRates() []string`

GetTaxRates returns the TaxRates field if non-nil, zero value otherwise.

### GetTaxRatesOk

`func (o *DtoCreateInvoiceRequest) GetTaxRatesOk() (*[]string, bool)`

GetTaxRatesOk returns a tuple with the TaxRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRates

`func (o *DtoCreateInvoiceRequest) SetTaxRates(v []string)`

SetTaxRates sets TaxRates field to given value.

### HasTaxRates

`func (o *DtoCreateInvoiceRequest) HasTaxRates() bool`

HasTaxRates returns a boolean if a field has been set.

### GetTotal

`func (o *DtoCreateInvoiceRequest) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DtoCreateInvoiceRequest) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DtoCreateInvoiceRequest) SetTotal(v string)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


