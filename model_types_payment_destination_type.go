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

// TypesPaymentDestinationType the model 'TypesPaymentDestinationType'
type TypesPaymentDestinationType string

// List of types.PaymentDestinationType
const (
	TYPESPAYMENTDESTINATIONTYPE_PaymentDestinationTypeInvoice TypesPaymentDestinationType = "INVOICE"
)

// All allowed values of TypesPaymentDestinationType enum
var AllowedTypesPaymentDestinationTypeEnumValues = []TypesPaymentDestinationType{
	"INVOICE",
}

func (v *TypesPaymentDestinationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesPaymentDestinationType(value)
	for _, existing := range AllowedTypesPaymentDestinationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesPaymentDestinationType", value)
}

// NewTypesPaymentDestinationTypeFromValue returns a pointer to a valid TypesPaymentDestinationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesPaymentDestinationTypeFromValue(v string) (*TypesPaymentDestinationType, error) {
	ev := TypesPaymentDestinationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesPaymentDestinationType: valid values are %v", v, AllowedTypesPaymentDestinationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesPaymentDestinationType) IsValid() bool {
	for _, existing := range AllowedTypesPaymentDestinationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.PaymentDestinationType value
func (v TypesPaymentDestinationType) Ptr() *TypesPaymentDestinationType {
	return &v
}

type NullableTypesPaymentDestinationType struct {
	value *TypesPaymentDestinationType
	isSet bool
}

func (v NullableTypesPaymentDestinationType) Get() *TypesPaymentDestinationType {
	return v.value
}

func (v *NullableTypesPaymentDestinationType) Set(val *TypesPaymentDestinationType) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesPaymentDestinationType) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesPaymentDestinationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesPaymentDestinationType(val *TypesPaymentDestinationType) *NullableTypesPaymentDestinationType {
	return &NullableTypesPaymentDestinationType{value: val, isSet: true}
}

func (v NullableTypesPaymentDestinationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesPaymentDestinationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

