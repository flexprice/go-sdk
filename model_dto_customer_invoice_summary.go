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

// checks if the DtoCustomerInvoiceSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoCustomerInvoiceSummary{}

// DtoCustomerInvoiceSummary struct for DtoCustomerInvoiceSummary
type DtoCustomerInvoiceSummary struct {
	Currency *string `json:"currency,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	OverdueInvoiceCount *int32 `json:"overdue_invoice_count,omitempty"`
	TotalInvoiceCount *int32 `json:"total_invoice_count,omitempty"`
	TotalOverdueAmount *float32 `json:"total_overdue_amount,omitempty"`
	TotalRevenueAmount *float32 `json:"total_revenue_amount,omitempty"`
	TotalUnpaidAmount *float32 `json:"total_unpaid_amount,omitempty"`
	UnpaidFixedCharges *float32 `json:"unpaid_fixed_charges,omitempty"`
	UnpaidInvoiceCount *int32 `json:"unpaid_invoice_count,omitempty"`
	UnpaidUsageCharges *float32 `json:"unpaid_usage_charges,omitempty"`
}

// NewDtoCustomerInvoiceSummary instantiates a new DtoCustomerInvoiceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoCustomerInvoiceSummary() *DtoCustomerInvoiceSummary {
	this := DtoCustomerInvoiceSummary{}
	return &this
}

// NewDtoCustomerInvoiceSummaryWithDefaults instantiates a new DtoCustomerInvoiceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoCustomerInvoiceSummaryWithDefaults() *DtoCustomerInvoiceSummary {
	this := DtoCustomerInvoiceSummary{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *DtoCustomerInvoiceSummary) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCustomerInvoiceSummary) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *DtoCustomerInvoiceSummary) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *DtoCustomerInvoiceSummary) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *DtoCustomerInvoiceSummary) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCustomerInvoiceSummary) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *DtoCustomerInvoiceSummary) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *DtoCustomerInvoiceSummary) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetOverdueInvoiceCount returns the OverdueInvoiceCount field value if set, zero value otherwise.
func (o *DtoCustomerInvoiceSummary) GetOverdueInvoiceCount() int32 {
	if o == nil || IsNil(o.OverdueInvoiceCount) {
		var ret int32
		return ret
	}
	return *o.OverdueInvoiceCount
}

// GetOverdueInvoiceCountOk returns a tuple with the OverdueInvoiceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCustomerInvoiceSummary) GetOverdueInvoiceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OverdueInvoiceCount) {
		return nil, false
	}
	return o.OverdueInvoiceCount, true
}

// HasOverdueInvoiceCount returns a boolean if a field has been set.
func (o *DtoCustomerInvoiceSummary) HasOverdueInvoiceCount() bool {
	if o != nil && !IsNil(o.OverdueInvoiceCount) {
		return true
	}

	return false
}

// SetOverdueInvoiceCount gets a reference to the given int32 and assigns it to the OverdueInvoiceCount field.
func (o *DtoCustomerInvoiceSummary) SetOverdueInvoiceCount(v int32) {
	o.OverdueInvoiceCount = &v
}

// GetTotalInvoiceCount returns the TotalInvoiceCount field value if set, zero value otherwise.
func (o *DtoCustomerInvoiceSummary) GetTotalInvoiceCount() int32 {
	if o == nil || IsNil(o.TotalInvoiceCount) {
		var ret int32
		return ret
	}
	return *o.TotalInvoiceCount
}

// GetTotalInvoiceCountOk returns a tuple with the TotalInvoiceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCustomerInvoiceSummary) GetTotalInvoiceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalInvoiceCount) {
		return nil, false
	}
	return o.TotalInvoiceCount, true
}

// HasTotalInvoiceCount returns a boolean if a field has been set.
func (o *DtoCustomerInvoiceSummary) HasTotalInvoiceCount() bool {
	if o != nil && !IsNil(o.TotalInvoiceCount) {
		return true
	}

	return false
}

// SetTotalInvoiceCount gets a reference to the given int32 and assigns it to the TotalInvoiceCount field.
func (o *DtoCustomerInvoiceSummary) SetTotalInvoiceCount(v int32) {
	o.TotalInvoiceCount = &v
}

// GetTotalOverdueAmount returns the TotalOverdueAmount field value if set, zero value otherwise.
func (o *DtoCustomerInvoiceSummary) GetTotalOverdueAmount() float32 {
	if o == nil || IsNil(o.TotalOverdueAmount) {
		var ret float32
		return ret
	}
	return *o.TotalOverdueAmount
}

// GetTotalOverdueAmountOk returns a tuple with the TotalOverdueAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCustomerInvoiceSummary) GetTotalOverdueAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalOverdueAmount) {
		return nil, false
	}
	return o.TotalOverdueAmount, true
}

// HasTotalOverdueAmount returns a boolean if a field has been set.
func (o *DtoCustomerInvoiceSummary) HasTotalOverdueAmount() bool {
	if o != nil && !IsNil(o.TotalOverdueAmount) {
		return true
	}

	return false
}

// SetTotalOverdueAmount gets a reference to the given float32 and assigns it to the TotalOverdueAmount field.
func (o *DtoCustomerInvoiceSummary) SetTotalOverdueAmount(v float32) {
	o.TotalOverdueAmount = &v
}

// GetTotalRevenueAmount returns the TotalRevenueAmount field value if set, zero value otherwise.
func (o *DtoCustomerInvoiceSummary) GetTotalRevenueAmount() float32 {
	if o == nil || IsNil(o.TotalRevenueAmount) {
		var ret float32
		return ret
	}
	return *o.TotalRevenueAmount
}

// GetTotalRevenueAmountOk returns a tuple with the TotalRevenueAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCustomerInvoiceSummary) GetTotalRevenueAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalRevenueAmount) {
		return nil, false
	}
	return o.TotalRevenueAmount, true
}

// HasTotalRevenueAmount returns a boolean if a field has been set.
func (o *DtoCustomerInvoiceSummary) HasTotalRevenueAmount() bool {
	if o != nil && !IsNil(o.TotalRevenueAmount) {
		return true
	}

	return false
}

// SetTotalRevenueAmount gets a reference to the given float32 and assigns it to the TotalRevenueAmount field.
func (o *DtoCustomerInvoiceSummary) SetTotalRevenueAmount(v float32) {
	o.TotalRevenueAmount = &v
}

// GetTotalUnpaidAmount returns the TotalUnpaidAmount field value if set, zero value otherwise.
func (o *DtoCustomerInvoiceSummary) GetTotalUnpaidAmount() float32 {
	if o == nil || IsNil(o.TotalUnpaidAmount) {
		var ret float32
		return ret
	}
	return *o.TotalUnpaidAmount
}

// GetTotalUnpaidAmountOk returns a tuple with the TotalUnpaidAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCustomerInvoiceSummary) GetTotalUnpaidAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalUnpaidAmount) {
		return nil, false
	}
	return o.TotalUnpaidAmount, true
}

// HasTotalUnpaidAmount returns a boolean if a field has been set.
func (o *DtoCustomerInvoiceSummary) HasTotalUnpaidAmount() bool {
	if o != nil && !IsNil(o.TotalUnpaidAmount) {
		return true
	}

	return false
}

// SetTotalUnpaidAmount gets a reference to the given float32 and assigns it to the TotalUnpaidAmount field.
func (o *DtoCustomerInvoiceSummary) SetTotalUnpaidAmount(v float32) {
	o.TotalUnpaidAmount = &v
}

// GetUnpaidFixedCharges returns the UnpaidFixedCharges field value if set, zero value otherwise.
func (o *DtoCustomerInvoiceSummary) GetUnpaidFixedCharges() float32 {
	if o == nil || IsNil(o.UnpaidFixedCharges) {
		var ret float32
		return ret
	}
	return *o.UnpaidFixedCharges
}

// GetUnpaidFixedChargesOk returns a tuple with the UnpaidFixedCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCustomerInvoiceSummary) GetUnpaidFixedChargesOk() (*float32, bool) {
	if o == nil || IsNil(o.UnpaidFixedCharges) {
		return nil, false
	}
	return o.UnpaidFixedCharges, true
}

// HasUnpaidFixedCharges returns a boolean if a field has been set.
func (o *DtoCustomerInvoiceSummary) HasUnpaidFixedCharges() bool {
	if o != nil && !IsNil(o.UnpaidFixedCharges) {
		return true
	}

	return false
}

// SetUnpaidFixedCharges gets a reference to the given float32 and assigns it to the UnpaidFixedCharges field.
func (o *DtoCustomerInvoiceSummary) SetUnpaidFixedCharges(v float32) {
	o.UnpaidFixedCharges = &v
}

// GetUnpaidInvoiceCount returns the UnpaidInvoiceCount field value if set, zero value otherwise.
func (o *DtoCustomerInvoiceSummary) GetUnpaidInvoiceCount() int32 {
	if o == nil || IsNil(o.UnpaidInvoiceCount) {
		var ret int32
		return ret
	}
	return *o.UnpaidInvoiceCount
}

// GetUnpaidInvoiceCountOk returns a tuple with the UnpaidInvoiceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCustomerInvoiceSummary) GetUnpaidInvoiceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UnpaidInvoiceCount) {
		return nil, false
	}
	return o.UnpaidInvoiceCount, true
}

// HasUnpaidInvoiceCount returns a boolean if a field has been set.
func (o *DtoCustomerInvoiceSummary) HasUnpaidInvoiceCount() bool {
	if o != nil && !IsNil(o.UnpaidInvoiceCount) {
		return true
	}

	return false
}

// SetUnpaidInvoiceCount gets a reference to the given int32 and assigns it to the UnpaidInvoiceCount field.
func (o *DtoCustomerInvoiceSummary) SetUnpaidInvoiceCount(v int32) {
	o.UnpaidInvoiceCount = &v
}

// GetUnpaidUsageCharges returns the UnpaidUsageCharges field value if set, zero value otherwise.
func (o *DtoCustomerInvoiceSummary) GetUnpaidUsageCharges() float32 {
	if o == nil || IsNil(o.UnpaidUsageCharges) {
		var ret float32
		return ret
	}
	return *o.UnpaidUsageCharges
}

// GetUnpaidUsageChargesOk returns a tuple with the UnpaidUsageCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCustomerInvoiceSummary) GetUnpaidUsageChargesOk() (*float32, bool) {
	if o == nil || IsNil(o.UnpaidUsageCharges) {
		return nil, false
	}
	return o.UnpaidUsageCharges, true
}

// HasUnpaidUsageCharges returns a boolean if a field has been set.
func (o *DtoCustomerInvoiceSummary) HasUnpaidUsageCharges() bool {
	if o != nil && !IsNil(o.UnpaidUsageCharges) {
		return true
	}

	return false
}

// SetUnpaidUsageCharges gets a reference to the given float32 and assigns it to the UnpaidUsageCharges field.
func (o *DtoCustomerInvoiceSummary) SetUnpaidUsageCharges(v float32) {
	o.UnpaidUsageCharges = &v
}

func (o DtoCustomerInvoiceSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoCustomerInvoiceSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.OverdueInvoiceCount) {
		toSerialize["overdue_invoice_count"] = o.OverdueInvoiceCount
	}
	if !IsNil(o.TotalInvoiceCount) {
		toSerialize["total_invoice_count"] = o.TotalInvoiceCount
	}
	if !IsNil(o.TotalOverdueAmount) {
		toSerialize["total_overdue_amount"] = o.TotalOverdueAmount
	}
	if !IsNil(o.TotalRevenueAmount) {
		toSerialize["total_revenue_amount"] = o.TotalRevenueAmount
	}
	if !IsNil(o.TotalUnpaidAmount) {
		toSerialize["total_unpaid_amount"] = o.TotalUnpaidAmount
	}
	if !IsNil(o.UnpaidFixedCharges) {
		toSerialize["unpaid_fixed_charges"] = o.UnpaidFixedCharges
	}
	if !IsNil(o.UnpaidInvoiceCount) {
		toSerialize["unpaid_invoice_count"] = o.UnpaidInvoiceCount
	}
	if !IsNil(o.UnpaidUsageCharges) {
		toSerialize["unpaid_usage_charges"] = o.UnpaidUsageCharges
	}
	return toSerialize, nil
}

type NullableDtoCustomerInvoiceSummary struct {
	value *DtoCustomerInvoiceSummary
	isSet bool
}

func (v NullableDtoCustomerInvoiceSummary) Get() *DtoCustomerInvoiceSummary {
	return v.value
}

func (v *NullableDtoCustomerInvoiceSummary) Set(val *DtoCustomerInvoiceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoCustomerInvoiceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoCustomerInvoiceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoCustomerInvoiceSummary(val *DtoCustomerInvoiceSummary) *NullableDtoCustomerInvoiceSummary {
	return &NullableDtoCustomerInvoiceSummary{value: val, isSet: true}
}

func (v NullableDtoCustomerInvoiceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoCustomerInvoiceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


