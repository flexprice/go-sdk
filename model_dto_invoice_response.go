/*
FlexPrice API

FlexPrice API Service

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flexprice

import (
	"encoding/json"
)

// checks if the DtoInvoiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoInvoiceResponse{}

// DtoInvoiceResponse struct for DtoInvoiceResponse
type DtoInvoiceResponse struct {
	AmountDue *float32 `json:"amount_due,omitempty"`
	AmountPaid *float32 `json:"amount_paid,omitempty"`
	AmountRemaining *float32 `json:"amount_remaining,omitempty"`
	BillingPeriod *string `json:"billing_period,omitempty"`
	BillingReason *string `json:"billing_reason,omitempty"`
	BillingSequence *int32 `json:"billing_sequence,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Customer *DtoCustomerResponse `json:"customer,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	Description *string `json:"description,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	FinalizedAt *string `json:"finalized_at,omitempty"`
	Id *string `json:"id,omitempty"`
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	InvoiceNumber *string `json:"invoice_number,omitempty"`
	InvoicePdfUrl *string `json:"invoice_pdf_url,omitempty"`
	InvoiceStatus *TypesInvoiceStatus `json:"invoice_status,omitempty"`
	InvoiceType *TypesInvoiceType `json:"invoice_type,omitempty"`
	LineItems []DtoInvoiceLineItemResponse `json:"line_items,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	PaidAt *string `json:"paid_at,omitempty"`
	PaymentStatus *TypesPaymentStatus `json:"payment_status,omitempty"`
	PeriodEnd *string `json:"period_end,omitempty"`
	PeriodStart *string `json:"period_start,omitempty"`
	Status *string `json:"status,omitempty"`
	Subscription *DtoSubscriptionResponse `json:"subscription,omitempty"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	Version *int32 `json:"version,omitempty"`
	VoidedAt *string `json:"voided_at,omitempty"`
}

// NewDtoInvoiceResponse instantiates a new DtoInvoiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoInvoiceResponse() *DtoInvoiceResponse {
	this := DtoInvoiceResponse{}
	return &this
}

// NewDtoInvoiceResponseWithDefaults instantiates a new DtoInvoiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoInvoiceResponseWithDefaults() *DtoInvoiceResponse {
	this := DtoInvoiceResponse{}
	return &this
}

// GetAmountDue returns the AmountDue field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetAmountDue() float32 {
	if o == nil || IsNil(o.AmountDue) {
		var ret float32
		return ret
	}
	return *o.AmountDue
}

// GetAmountDueOk returns a tuple with the AmountDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetAmountDueOk() (*float32, bool) {
	if o == nil || IsNil(o.AmountDue) {
		return nil, false
	}
	return o.AmountDue, true
}

// HasAmountDue returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasAmountDue() bool {
	if o != nil && !IsNil(o.AmountDue) {
		return true
	}

	return false
}

// SetAmountDue gets a reference to the given float32 and assigns it to the AmountDue field.
func (o *DtoInvoiceResponse) SetAmountDue(v float32) {
	o.AmountDue = &v
}

// GetAmountPaid returns the AmountPaid field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetAmountPaid() float32 {
	if o == nil || IsNil(o.AmountPaid) {
		var ret float32
		return ret
	}
	return *o.AmountPaid
}

// GetAmountPaidOk returns a tuple with the AmountPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetAmountPaidOk() (*float32, bool) {
	if o == nil || IsNil(o.AmountPaid) {
		return nil, false
	}
	return o.AmountPaid, true
}

// HasAmountPaid returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasAmountPaid() bool {
	if o != nil && !IsNil(o.AmountPaid) {
		return true
	}

	return false
}

// SetAmountPaid gets a reference to the given float32 and assigns it to the AmountPaid field.
func (o *DtoInvoiceResponse) SetAmountPaid(v float32) {
	o.AmountPaid = &v
}

// GetAmountRemaining returns the AmountRemaining field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetAmountRemaining() float32 {
	if o == nil || IsNil(o.AmountRemaining) {
		var ret float32
		return ret
	}
	return *o.AmountRemaining
}

// GetAmountRemainingOk returns a tuple with the AmountRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetAmountRemainingOk() (*float32, bool) {
	if o == nil || IsNil(o.AmountRemaining) {
		return nil, false
	}
	return o.AmountRemaining, true
}

// HasAmountRemaining returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasAmountRemaining() bool {
	if o != nil && !IsNil(o.AmountRemaining) {
		return true
	}

	return false
}

// SetAmountRemaining gets a reference to the given float32 and assigns it to the AmountRemaining field.
func (o *DtoInvoiceResponse) SetAmountRemaining(v float32) {
	o.AmountRemaining = &v
}

// GetBillingPeriod returns the BillingPeriod field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetBillingPeriod() string {
	if o == nil || IsNil(o.BillingPeriod) {
		var ret string
		return ret
	}
	return *o.BillingPeriod
}

// GetBillingPeriodOk returns a tuple with the BillingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetBillingPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.BillingPeriod) {
		return nil, false
	}
	return o.BillingPeriod, true
}

// HasBillingPeriod returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasBillingPeriod() bool {
	if o != nil && !IsNil(o.BillingPeriod) {
		return true
	}

	return false
}

// SetBillingPeriod gets a reference to the given string and assigns it to the BillingPeriod field.
func (o *DtoInvoiceResponse) SetBillingPeriod(v string) {
	o.BillingPeriod = &v
}

// GetBillingReason returns the BillingReason field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetBillingReason() string {
	if o == nil || IsNil(o.BillingReason) {
		var ret string
		return ret
	}
	return *o.BillingReason
}

// GetBillingReasonOk returns a tuple with the BillingReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetBillingReasonOk() (*string, bool) {
	if o == nil || IsNil(o.BillingReason) {
		return nil, false
	}
	return o.BillingReason, true
}

// HasBillingReason returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasBillingReason() bool {
	if o != nil && !IsNil(o.BillingReason) {
		return true
	}

	return false
}

// SetBillingReason gets a reference to the given string and assigns it to the BillingReason field.
func (o *DtoInvoiceResponse) SetBillingReason(v string) {
	o.BillingReason = &v
}

// GetBillingSequence returns the BillingSequence field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetBillingSequence() int32 {
	if o == nil || IsNil(o.BillingSequence) {
		var ret int32
		return ret
	}
	return *o.BillingSequence
}

// GetBillingSequenceOk returns a tuple with the BillingSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetBillingSequenceOk() (*int32, bool) {
	if o == nil || IsNil(o.BillingSequence) {
		return nil, false
	}
	return o.BillingSequence, true
}

// HasBillingSequence returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasBillingSequence() bool {
	if o != nil && !IsNil(o.BillingSequence) {
		return true
	}

	return false
}

// SetBillingSequence gets a reference to the given int32 and assigns it to the BillingSequence field.
func (o *DtoInvoiceResponse) SetBillingSequence(v int32) {
	o.BillingSequence = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DtoInvoiceResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DtoInvoiceResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *DtoInvoiceResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetCustomer() DtoCustomerResponse {
	if o == nil || IsNil(o.Customer) {
		var ret DtoCustomerResponse
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetCustomerOk() (*DtoCustomerResponse, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given DtoCustomerResponse and assigns it to the Customer field.
func (o *DtoInvoiceResponse) SetCustomer(v DtoCustomerResponse) {
	o.Customer = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *DtoInvoiceResponse) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DtoInvoiceResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *DtoInvoiceResponse) SetDueDate(v string) {
	o.DueDate = &v
}

// GetFinalizedAt returns the FinalizedAt field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetFinalizedAt() string {
	if o == nil || IsNil(o.FinalizedAt) {
		var ret string
		return ret
	}
	return *o.FinalizedAt
}

// GetFinalizedAtOk returns a tuple with the FinalizedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetFinalizedAtOk() (*string, bool) {
	if o == nil || IsNil(o.FinalizedAt) {
		return nil, false
	}
	return o.FinalizedAt, true
}

// HasFinalizedAt returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasFinalizedAt() bool {
	if o != nil && !IsNil(o.FinalizedAt) {
		return true
	}

	return false
}

// SetFinalizedAt gets a reference to the given string and assigns it to the FinalizedAt field.
func (o *DtoInvoiceResponse) SetFinalizedAt(v string) {
	o.FinalizedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DtoInvoiceResponse) SetId(v string) {
	o.Id = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetIdempotencyKey() string {
	if o == nil || IsNil(o.IdempotencyKey) {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotencyKey) {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasIdempotencyKey() bool {
	if o != nil && !IsNil(o.IdempotencyKey) {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *DtoInvoiceResponse) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *DtoInvoiceResponse) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetInvoicePdfUrl returns the InvoicePdfUrl field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetInvoicePdfUrl() string {
	if o == nil || IsNil(o.InvoicePdfUrl) {
		var ret string
		return ret
	}
	return *o.InvoicePdfUrl
}

// GetInvoicePdfUrlOk returns a tuple with the InvoicePdfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetInvoicePdfUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InvoicePdfUrl) {
		return nil, false
	}
	return o.InvoicePdfUrl, true
}

// HasInvoicePdfUrl returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasInvoicePdfUrl() bool {
	if o != nil && !IsNil(o.InvoicePdfUrl) {
		return true
	}

	return false
}

// SetInvoicePdfUrl gets a reference to the given string and assigns it to the InvoicePdfUrl field.
func (o *DtoInvoiceResponse) SetInvoicePdfUrl(v string) {
	o.InvoicePdfUrl = &v
}

// GetInvoiceStatus returns the InvoiceStatus field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetInvoiceStatus() TypesInvoiceStatus {
	if o == nil || IsNil(o.InvoiceStatus) {
		var ret TypesInvoiceStatus
		return ret
	}
	return *o.InvoiceStatus
}

// GetInvoiceStatusOk returns a tuple with the InvoiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetInvoiceStatusOk() (*TypesInvoiceStatus, bool) {
	if o == nil || IsNil(o.InvoiceStatus) {
		return nil, false
	}
	return o.InvoiceStatus, true
}

// HasInvoiceStatus returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasInvoiceStatus() bool {
	if o != nil && !IsNil(o.InvoiceStatus) {
		return true
	}

	return false
}

// SetInvoiceStatus gets a reference to the given TypesInvoiceStatus and assigns it to the InvoiceStatus field.
func (o *DtoInvoiceResponse) SetInvoiceStatus(v TypesInvoiceStatus) {
	o.InvoiceStatus = &v
}

// GetInvoiceType returns the InvoiceType field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetInvoiceType() TypesInvoiceType {
	if o == nil || IsNil(o.InvoiceType) {
		var ret TypesInvoiceType
		return ret
	}
	return *o.InvoiceType
}

// GetInvoiceTypeOk returns a tuple with the InvoiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetInvoiceTypeOk() (*TypesInvoiceType, bool) {
	if o == nil || IsNil(o.InvoiceType) {
		return nil, false
	}
	return o.InvoiceType, true
}

// HasInvoiceType returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasInvoiceType() bool {
	if o != nil && !IsNil(o.InvoiceType) {
		return true
	}

	return false
}

// SetInvoiceType gets a reference to the given TypesInvoiceType and assigns it to the InvoiceType field.
func (o *DtoInvoiceResponse) SetInvoiceType(v TypesInvoiceType) {
	o.InvoiceType = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetLineItems() []DtoInvoiceLineItemResponse {
	if o == nil || IsNil(o.LineItems) {
		var ret []DtoInvoiceLineItemResponse
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetLineItemsOk() ([]DtoInvoiceLineItemResponse, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []DtoInvoiceLineItemResponse and assigns it to the LineItems field.
func (o *DtoInvoiceResponse) SetLineItems(v []DtoInvoiceLineItemResponse) {
	o.LineItems = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DtoInvoiceResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetPaidAt returns the PaidAt field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetPaidAt() string {
	if o == nil || IsNil(o.PaidAt) {
		var ret string
		return ret
	}
	return *o.PaidAt
}

// GetPaidAtOk returns a tuple with the PaidAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetPaidAtOk() (*string, bool) {
	if o == nil || IsNil(o.PaidAt) {
		return nil, false
	}
	return o.PaidAt, true
}

// HasPaidAt returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasPaidAt() bool {
	if o != nil && !IsNil(o.PaidAt) {
		return true
	}

	return false
}

// SetPaidAt gets a reference to the given string and assigns it to the PaidAt field.
func (o *DtoInvoiceResponse) SetPaidAt(v string) {
	o.PaidAt = &v
}

// GetPaymentStatus returns the PaymentStatus field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetPaymentStatus() TypesPaymentStatus {
	if o == nil || IsNil(o.PaymentStatus) {
		var ret TypesPaymentStatus
		return ret
	}
	return *o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetPaymentStatusOk() (*TypesPaymentStatus, bool) {
	if o == nil || IsNil(o.PaymentStatus) {
		return nil, false
	}
	return o.PaymentStatus, true
}

// HasPaymentStatus returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasPaymentStatus() bool {
	if o != nil && !IsNil(o.PaymentStatus) {
		return true
	}

	return false
}

// SetPaymentStatus gets a reference to the given TypesPaymentStatus and assigns it to the PaymentStatus field.
func (o *DtoInvoiceResponse) SetPaymentStatus(v TypesPaymentStatus) {
	o.PaymentStatus = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetPeriodEnd() string {
	if o == nil || IsNil(o.PeriodEnd) {
		var ret string
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetPeriodEndOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodEnd) {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasPeriodEnd() bool {
	if o != nil && !IsNil(o.PeriodEnd) {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given string and assigns it to the PeriodEnd field.
func (o *DtoInvoiceResponse) SetPeriodEnd(v string) {
	o.PeriodEnd = &v
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetPeriodStart() string {
	if o == nil || IsNil(o.PeriodStart) {
		var ret string
		return ret
	}
	return *o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetPeriodStartOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodStart) {
		return nil, false
	}
	return o.PeriodStart, true
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasPeriodStart() bool {
	if o != nil && !IsNil(o.PeriodStart) {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given string and assigns it to the PeriodStart field.
func (o *DtoInvoiceResponse) SetPeriodStart(v string) {
	o.PeriodStart = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DtoInvoiceResponse) SetStatus(v string) {
	o.Status = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetSubscription() DtoSubscriptionResponse {
	if o == nil || IsNil(o.Subscription) {
		var ret DtoSubscriptionResponse
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetSubscriptionOk() (*DtoSubscriptionResponse, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given DtoSubscriptionResponse and assigns it to the Subscription field.
func (o *DtoInvoiceResponse) SetSubscription(v DtoSubscriptionResponse) {
	o.Subscription = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *DtoInvoiceResponse) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *DtoInvoiceResponse) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DtoInvoiceResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *DtoInvoiceResponse) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DtoInvoiceResponse) SetVersion(v int32) {
	o.Version = &v
}

// GetVoidedAt returns the VoidedAt field value if set, zero value otherwise.
func (o *DtoInvoiceResponse) GetVoidedAt() string {
	if o == nil || IsNil(o.VoidedAt) {
		var ret string
		return ret
	}
	return *o.VoidedAt
}

// GetVoidedAtOk returns a tuple with the VoidedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoInvoiceResponse) GetVoidedAtOk() (*string, bool) {
	if o == nil || IsNil(o.VoidedAt) {
		return nil, false
	}
	return o.VoidedAt, true
}

// HasVoidedAt returns a boolean if a field has been set.
func (o *DtoInvoiceResponse) HasVoidedAt() bool {
	if o != nil && !IsNil(o.VoidedAt) {
		return true
	}

	return false
}

// SetVoidedAt gets a reference to the given string and assigns it to the VoidedAt field.
func (o *DtoInvoiceResponse) SetVoidedAt(v string) {
	o.VoidedAt = &v
}

func (o DtoInvoiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoInvoiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmountDue) {
		toSerialize["amount_due"] = o.AmountDue
	}
	if !IsNil(o.AmountPaid) {
		toSerialize["amount_paid"] = o.AmountPaid
	}
	if !IsNil(o.AmountRemaining) {
		toSerialize["amount_remaining"] = o.AmountRemaining
	}
	if !IsNil(o.BillingPeriod) {
		toSerialize["billing_period"] = o.BillingPeriod
	}
	if !IsNil(o.BillingReason) {
		toSerialize["billing_reason"] = o.BillingReason
	}
	if !IsNil(o.BillingSequence) {
		toSerialize["billing_sequence"] = o.BillingSequence
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if !IsNil(o.FinalizedAt) {
		toSerialize["finalized_at"] = o.FinalizedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IdempotencyKey) {
		toSerialize["idempotency_key"] = o.IdempotencyKey
	}
	if !IsNil(o.InvoiceNumber) {
		toSerialize["invoice_number"] = o.InvoiceNumber
	}
	if !IsNil(o.InvoicePdfUrl) {
		toSerialize["invoice_pdf_url"] = o.InvoicePdfUrl
	}
	if !IsNil(o.InvoiceStatus) {
		toSerialize["invoice_status"] = o.InvoiceStatus
	}
	if !IsNil(o.InvoiceType) {
		toSerialize["invoice_type"] = o.InvoiceType
	}
	if !IsNil(o.LineItems) {
		toSerialize["line_items"] = o.LineItems
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PaidAt) {
		toSerialize["paid_at"] = o.PaidAt
	}
	if !IsNil(o.PaymentStatus) {
		toSerialize["payment_status"] = o.PaymentStatus
	}
	if !IsNil(o.PeriodEnd) {
		toSerialize["period_end"] = o.PeriodEnd
	}
	if !IsNil(o.PeriodStart) {
		toSerialize["period_start"] = o.PeriodStart
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VoidedAt) {
		toSerialize["voided_at"] = o.VoidedAt
	}
	return toSerialize, nil
}

type NullableDtoInvoiceResponse struct {
	value *DtoInvoiceResponse
	isSet bool
}

func (v NullableDtoInvoiceResponse) Get() *DtoInvoiceResponse {
	return v.value
}

func (v *NullableDtoInvoiceResponse) Set(val *DtoInvoiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoInvoiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoInvoiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoInvoiceResponse(val *DtoInvoiceResponse) *NullableDtoInvoiceResponse {
	return &NullableDtoInvoiceResponse{value: val, isSet: true}
}

func (v NullableDtoInvoiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoInvoiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


