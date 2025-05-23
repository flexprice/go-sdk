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

// TypesPauseMode the model 'TypesPauseMode'
type TypesPauseMode string

// List of types.PauseMode
const (
	TYPESPAUSEMODE_PauseModeImmediate TypesPauseMode = "immediate"
	TYPESPAUSEMODE_PauseModeScheduled TypesPauseMode = "scheduled"
	TYPESPAUSEMODE_PauseModePeriodEnd TypesPauseMode = "period_end"
)

// All allowed values of TypesPauseMode enum
var AllowedTypesPauseModeEnumValues = []TypesPauseMode{
	"immediate",
	"scheduled",
	"period_end",
}

func (v *TypesPauseMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesPauseMode(value)
	for _, existing := range AllowedTypesPauseModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesPauseMode", value)
}

// NewTypesPauseModeFromValue returns a pointer to a valid TypesPauseMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesPauseModeFromValue(v string) (*TypesPauseMode, error) {
	ev := TypesPauseMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesPauseMode: valid values are %v", v, AllowedTypesPauseModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesPauseMode) IsValid() bool {
	for _, existing := range AllowedTypesPauseModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.PauseMode value
func (v TypesPauseMode) Ptr() *TypesPauseMode {
	return &v
}

type NullableTypesPauseMode struct {
	value *TypesPauseMode
	isSet bool
}

func (v NullableTypesPauseMode) Get() *TypesPauseMode {
	return v.value
}

func (v *NullableTypesPauseMode) Set(val *TypesPauseMode) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesPauseMode) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesPauseMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesPauseMode(val *TypesPauseMode) *NullableTypesPauseMode {
	return &NullableTypesPauseMode{value: val, isSet: true}
}

func (v NullableTypesPauseMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesPauseMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

