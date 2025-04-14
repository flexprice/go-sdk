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

// checks if the DtoPaymentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoPaymentResponse{}

// DtoPaymentResponse struct for DtoPaymentResponse
type DtoPaymentResponse struct {
	Amount *float32 `json:"amount,omitempty"`
	Attempts []DtoPaymentAttemptResponse `json:"attempts,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DestinationId *string `json:"destination_id,omitempty"`
	DestinationType *TypesPaymentDestinationType `json:"destination_type,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	FailedAt *string `json:"failed_at,omitempty"`
	Id *string `json:"id,omitempty"`
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	InvoiceNumber *string `json:"invoice_number,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	PaymentMethodId *string `json:"payment_method_id,omitempty"`
	PaymentMethodType *TypesPaymentMethodType `json:"payment_method_type,omitempty"`
	PaymentStatus *TypesPaymentStatus `json:"payment_status,omitempty"`
	RefundedAt *string `json:"refunded_at,omitempty"`
	SucceededAt *string `json:"succeeded_at,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	TrackAttempts *bool `json:"track_attempts,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// NewDtoPaymentResponse instantiates a new DtoPaymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoPaymentResponse() *DtoPaymentResponse {
	this := DtoPaymentResponse{}
	return &this
}

// NewDtoPaymentResponseWithDefaults instantiates a new DtoPaymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoPaymentResponseWithDefaults() *DtoPaymentResponse {
	this := DtoPaymentResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *DtoPaymentResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetAttempts returns the Attempts field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetAttempts() []DtoPaymentAttemptResponse {
	if o == nil || IsNil(o.Attempts) {
		var ret []DtoPaymentAttemptResponse
		return ret
	}
	return o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetAttemptsOk() ([]DtoPaymentAttemptResponse, bool) {
	if o == nil || IsNil(o.Attempts) {
		return nil, false
	}
	return o.Attempts, true
}

// HasAttempts returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasAttempts() bool {
	if o != nil && !IsNil(o.Attempts) {
		return true
	}

	return false
}

// SetAttempts gets a reference to the given []DtoPaymentAttemptResponse and assigns it to the Attempts field.
func (o *DtoPaymentResponse) SetAttempts(v []DtoPaymentAttemptResponse) {
	o.Attempts = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DtoPaymentResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DtoPaymentResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *DtoPaymentResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetDestinationId() string {
	if o == nil || IsNil(o.DestinationId) {
		var ret string
		return ret
	}
	return *o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetDestinationIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationId) {
		return nil, false
	}
	return o.DestinationId, true
}

// HasDestinationId returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasDestinationId() bool {
	if o != nil && !IsNil(o.DestinationId) {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given string and assigns it to the DestinationId field.
func (o *DtoPaymentResponse) SetDestinationId(v string) {
	o.DestinationId = &v
}

// GetDestinationType returns the DestinationType field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetDestinationType() TypesPaymentDestinationType {
	if o == nil || IsNil(o.DestinationType) {
		var ret TypesPaymentDestinationType
		return ret
	}
	return *o.DestinationType
}

// GetDestinationTypeOk returns a tuple with the DestinationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetDestinationTypeOk() (*TypesPaymentDestinationType, bool) {
	if o == nil || IsNil(o.DestinationType) {
		return nil, false
	}
	return o.DestinationType, true
}

// HasDestinationType returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasDestinationType() bool {
	if o != nil && !IsNil(o.DestinationType) {
		return true
	}

	return false
}

// SetDestinationType gets a reference to the given TypesPaymentDestinationType and assigns it to the DestinationType field.
func (o *DtoPaymentResponse) SetDestinationType(v TypesPaymentDestinationType) {
	o.DestinationType = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *DtoPaymentResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetFailedAt returns the FailedAt field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetFailedAt() string {
	if o == nil || IsNil(o.FailedAt) {
		var ret string
		return ret
	}
	return *o.FailedAt
}

// GetFailedAtOk returns a tuple with the FailedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetFailedAtOk() (*string, bool) {
	if o == nil || IsNil(o.FailedAt) {
		return nil, false
	}
	return o.FailedAt, true
}

// HasFailedAt returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasFailedAt() bool {
	if o != nil && !IsNil(o.FailedAt) {
		return true
	}

	return false
}

// SetFailedAt gets a reference to the given string and assigns it to the FailedAt field.
func (o *DtoPaymentResponse) SetFailedAt(v string) {
	o.FailedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DtoPaymentResponse) SetId(v string) {
	o.Id = &v
}

// GetIdempotencyKey returns the IdempotencyKey field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetIdempotencyKey() string {
	if o == nil || IsNil(o.IdempotencyKey) {
		var ret string
		return ret
	}
	return *o.IdempotencyKey
}

// GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetIdempotencyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IdempotencyKey) {
		return nil, false
	}
	return o.IdempotencyKey, true
}

// HasIdempotencyKey returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasIdempotencyKey() bool {
	if o != nil && !IsNil(o.IdempotencyKey) {
		return true
	}

	return false
}

// SetIdempotencyKey gets a reference to the given string and assigns it to the IdempotencyKey field.
func (o *DtoPaymentResponse) SetIdempotencyKey(v string) {
	o.IdempotencyKey = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *DtoPaymentResponse) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DtoPaymentResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *DtoPaymentResponse) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetPaymentMethodType() TypesPaymentMethodType {
	if o == nil || IsNil(o.PaymentMethodType) {
		var ret TypesPaymentMethodType
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetPaymentMethodTypeOk() (*TypesPaymentMethodType, bool) {
	if o == nil || IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasPaymentMethodType() bool {
	if o != nil && !IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given TypesPaymentMethodType and assigns it to the PaymentMethodType field.
func (o *DtoPaymentResponse) SetPaymentMethodType(v TypesPaymentMethodType) {
	o.PaymentMethodType = &v
}

// GetPaymentStatus returns the PaymentStatus field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetPaymentStatus() TypesPaymentStatus {
	if o == nil || IsNil(o.PaymentStatus) {
		var ret TypesPaymentStatus
		return ret
	}
	return *o.PaymentStatus
}

// GetPaymentStatusOk returns a tuple with the PaymentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetPaymentStatusOk() (*TypesPaymentStatus, bool) {
	if o == nil || IsNil(o.PaymentStatus) {
		return nil, false
	}
	return o.PaymentStatus, true
}

// HasPaymentStatus returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasPaymentStatus() bool {
	if o != nil && !IsNil(o.PaymentStatus) {
		return true
	}

	return false
}

// SetPaymentStatus gets a reference to the given TypesPaymentStatus and assigns it to the PaymentStatus field.
func (o *DtoPaymentResponse) SetPaymentStatus(v TypesPaymentStatus) {
	o.PaymentStatus = &v
}

// GetRefundedAt returns the RefundedAt field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetRefundedAt() string {
	if o == nil || IsNil(o.RefundedAt) {
		var ret string
		return ret
	}
	return *o.RefundedAt
}

// GetRefundedAtOk returns a tuple with the RefundedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetRefundedAtOk() (*string, bool) {
	if o == nil || IsNil(o.RefundedAt) {
		return nil, false
	}
	return o.RefundedAt, true
}

// HasRefundedAt returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasRefundedAt() bool {
	if o != nil && !IsNil(o.RefundedAt) {
		return true
	}

	return false
}

// SetRefundedAt gets a reference to the given string and assigns it to the RefundedAt field.
func (o *DtoPaymentResponse) SetRefundedAt(v string) {
	o.RefundedAt = &v
}

// GetSucceededAt returns the SucceededAt field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetSucceededAt() string {
	if o == nil || IsNil(o.SucceededAt) {
		var ret string
		return ret
	}
	return *o.SucceededAt
}

// GetSucceededAtOk returns a tuple with the SucceededAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetSucceededAtOk() (*string, bool) {
	if o == nil || IsNil(o.SucceededAt) {
		return nil, false
	}
	return o.SucceededAt, true
}

// HasSucceededAt returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasSucceededAt() bool {
	if o != nil && !IsNil(o.SucceededAt) {
		return true
	}

	return false
}

// SetSucceededAt gets a reference to the given string and assigns it to the SucceededAt field.
func (o *DtoPaymentResponse) SetSucceededAt(v string) {
	o.SucceededAt = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *DtoPaymentResponse) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTrackAttempts returns the TrackAttempts field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetTrackAttempts() bool {
	if o == nil || IsNil(o.TrackAttempts) {
		var ret bool
		return ret
	}
	return *o.TrackAttempts
}

// GetTrackAttemptsOk returns a tuple with the TrackAttempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetTrackAttemptsOk() (*bool, bool) {
	if o == nil || IsNil(o.TrackAttempts) {
		return nil, false
	}
	return o.TrackAttempts, true
}

// HasTrackAttempts returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasTrackAttempts() bool {
	if o != nil && !IsNil(o.TrackAttempts) {
		return true
	}

	return false
}

// SetTrackAttempts gets a reference to the given bool and assigns it to the TrackAttempts field.
func (o *DtoPaymentResponse) SetTrackAttempts(v bool) {
	o.TrackAttempts = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DtoPaymentResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *DtoPaymentResponse) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPaymentResponse) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DtoPaymentResponse) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *DtoPaymentResponse) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o DtoPaymentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoPaymentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Attempts) {
		toSerialize["attempts"] = o.Attempts
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
	if !IsNil(o.DestinationId) {
		toSerialize["destination_id"] = o.DestinationId
	}
	if !IsNil(o.DestinationType) {
		toSerialize["destination_type"] = o.DestinationType
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if !IsNil(o.FailedAt) {
		toSerialize["failed_at"] = o.FailedAt
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
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PaymentMethodId) {
		toSerialize["payment_method_id"] = o.PaymentMethodId
	}
	if !IsNil(o.PaymentMethodType) {
		toSerialize["payment_method_type"] = o.PaymentMethodType
	}
	if !IsNil(o.PaymentStatus) {
		toSerialize["payment_status"] = o.PaymentStatus
	}
	if !IsNil(o.RefundedAt) {
		toSerialize["refunded_at"] = o.RefundedAt
	}
	if !IsNil(o.SucceededAt) {
		toSerialize["succeeded_at"] = o.SucceededAt
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !IsNil(o.TrackAttempts) {
		toSerialize["track_attempts"] = o.TrackAttempts
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	return toSerialize, nil
}

type NullableDtoPaymentResponse struct {
	value *DtoPaymentResponse
	isSet bool
}

func (v NullableDtoPaymentResponse) Get() *DtoPaymentResponse {
	return v.value
}

func (v *NullableDtoPaymentResponse) Set(val *DtoPaymentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoPaymentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoPaymentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoPaymentResponse(val *DtoPaymentResponse) *NullableDtoPaymentResponse {
	return &NullableDtoPaymentResponse{value: val, isSet: true}
}

func (v NullableDtoPaymentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoPaymentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


