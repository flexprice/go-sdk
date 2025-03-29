/*
FlexPrice API

FlexPrice API Service

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flexpriceclient

import (
	"encoding/json"
	"fmt"
)

// TypesInvoiceCadence the model 'TypesInvoiceCadence'
type TypesInvoiceCadence string

// List of types.InvoiceCadence
const (
	TYPESINVOICECADENCE_InvoiceCadenceArrear TypesInvoiceCadence = "ARREAR"
	TYPESINVOICECADENCE_InvoiceCadenceAdvance TypesInvoiceCadence = "ADVANCE"
)

// All allowed values of TypesInvoiceCadence enum
var AllowedTypesInvoiceCadenceEnumValues = []TypesInvoiceCadence{
	"ARREAR",
	"ADVANCE",
}

func (v *TypesInvoiceCadence) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesInvoiceCadence(value)
	for _, existing := range AllowedTypesInvoiceCadenceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesInvoiceCadence", value)
}

// NewTypesInvoiceCadenceFromValue returns a pointer to a valid TypesInvoiceCadence
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesInvoiceCadenceFromValue(v string) (*TypesInvoiceCadence, error) {
	ev := TypesInvoiceCadence(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesInvoiceCadence: valid values are %v", v, AllowedTypesInvoiceCadenceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesInvoiceCadence) IsValid() bool {
	for _, existing := range AllowedTypesInvoiceCadenceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.InvoiceCadence value
func (v TypesInvoiceCadence) Ptr() *TypesInvoiceCadence {
	return &v
}

type NullableTypesInvoiceCadence struct {
	value *TypesInvoiceCadence
	isSet bool
}

func (v NullableTypesInvoiceCadence) Get() *TypesInvoiceCadence {
	return v.value
}

func (v *NullableTypesInvoiceCadence) Set(val *TypesInvoiceCadence) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesInvoiceCadence) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesInvoiceCadence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesInvoiceCadence(val *TypesInvoiceCadence) *NullableTypesInvoiceCadence {
	return &NullableTypesInvoiceCadence{value: val, isSet: true}
}

func (v NullableTypesInvoiceCadence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesInvoiceCadence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

