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

// TypesAutoTopupTrigger the model 'TypesAutoTopupTrigger'
type TypesAutoTopupTrigger string

// List of types.AutoTopupTrigger
const (
	TYPESAUTOTOPUPTRIGGER_AutoTopupTriggerDisabled TypesAutoTopupTrigger = "disabled"
	TYPESAUTOTOPUPTRIGGER_AutoTopupTriggerBalanceBelowThreshold TypesAutoTopupTrigger = "balance_below_threshold"
)

// All allowed values of TypesAutoTopupTrigger enum
var AllowedTypesAutoTopupTriggerEnumValues = []TypesAutoTopupTrigger{
	"disabled",
	"balance_below_threshold",
}

func (v *TypesAutoTopupTrigger) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesAutoTopupTrigger(value)
	for _, existing := range AllowedTypesAutoTopupTriggerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesAutoTopupTrigger", value)
}

// NewTypesAutoTopupTriggerFromValue returns a pointer to a valid TypesAutoTopupTrigger
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesAutoTopupTriggerFromValue(v string) (*TypesAutoTopupTrigger, error) {
	ev := TypesAutoTopupTrigger(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesAutoTopupTrigger: valid values are %v", v, AllowedTypesAutoTopupTriggerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesAutoTopupTrigger) IsValid() bool {
	for _, existing := range AllowedTypesAutoTopupTriggerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.AutoTopupTrigger value
func (v TypesAutoTopupTrigger) Ptr() *TypesAutoTopupTrigger {
	return &v
}

type NullableTypesAutoTopupTrigger struct {
	value *TypesAutoTopupTrigger
	isSet bool
}

func (v NullableTypesAutoTopupTrigger) Get() *TypesAutoTopupTrigger {
	return v.value
}

func (v *NullableTypesAutoTopupTrigger) Set(val *TypesAutoTopupTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesAutoTopupTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesAutoTopupTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesAutoTopupTrigger(val *TypesAutoTopupTrigger) *NullableTypesAutoTopupTrigger {
	return &NullableTypesAutoTopupTrigger{value: val, isSet: true}
}

func (v NullableTypesAutoTopupTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesAutoTopupTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

