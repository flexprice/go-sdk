/*
FlexPrice API

FlexPrice API Service

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flexprice

import (
	"encoding/json"
	"fmt"
)

// TypesInvoiceBillingReason the model 'TypesInvoiceBillingReason'
type TypesInvoiceBillingReason string

// List of types.InvoiceBillingReason
const (
	TYPESINVOICEBILLINGREASON_InvoiceBillingReasonSubscriptionCreate TypesInvoiceBillingReason = "SUBSCRIPTION_CREATE"
	TYPESINVOICEBILLINGREASON_InvoiceBillingReasonSubscriptionCycle TypesInvoiceBillingReason = "SUBSCRIPTION_CYCLE"
	TYPESINVOICEBILLINGREASON_InvoiceBillingReasonSubscriptionUpdate TypesInvoiceBillingReason = "SUBSCRIPTION_UPDATE"
	TYPESINVOICEBILLINGREASON_InvoiceBillingReasonManual TypesInvoiceBillingReason = "MANUAL"
)

// All allowed values of TypesInvoiceBillingReason enum
var AllowedTypesInvoiceBillingReasonEnumValues = []TypesInvoiceBillingReason{
	"SUBSCRIPTION_CREATE",
	"SUBSCRIPTION_CYCLE",
	"SUBSCRIPTION_UPDATE",
	"MANUAL",
}

func (v *TypesInvoiceBillingReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesInvoiceBillingReason(value)
	for _, existing := range AllowedTypesInvoiceBillingReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesInvoiceBillingReason", value)
}

// NewTypesInvoiceBillingReasonFromValue returns a pointer to a valid TypesInvoiceBillingReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesInvoiceBillingReasonFromValue(v string) (*TypesInvoiceBillingReason, error) {
	ev := TypesInvoiceBillingReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesInvoiceBillingReason: valid values are %v", v, AllowedTypesInvoiceBillingReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesInvoiceBillingReason) IsValid() bool {
	for _, existing := range AllowedTypesInvoiceBillingReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.InvoiceBillingReason value
func (v TypesInvoiceBillingReason) Ptr() *TypesInvoiceBillingReason {
	return &v
}

type NullableTypesInvoiceBillingReason struct {
	value *TypesInvoiceBillingReason
	isSet bool
}

func (v NullableTypesInvoiceBillingReason) Get() *TypesInvoiceBillingReason {
	return v.value
}

func (v *NullableTypesInvoiceBillingReason) Set(val *TypesInvoiceBillingReason) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesInvoiceBillingReason) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesInvoiceBillingReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesInvoiceBillingReason(val *TypesInvoiceBillingReason) *NullableTypesInvoiceBillingReason {
	return &NullableTypesInvoiceBillingReason{value: val, isSet: true}
}

func (v NullableTypesInvoiceBillingReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesInvoiceBillingReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

