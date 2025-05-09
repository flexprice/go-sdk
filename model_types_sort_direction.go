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

// TypesSortDirection the model 'TypesSortDirection'
type TypesSortDirection string

// List of types.SortDirection
const (
	TYPESSORTDIRECTION_SortDirectionAsc TypesSortDirection = "asc"
	TYPESSORTDIRECTION_SortDirectionDesc TypesSortDirection = "desc"
)

// All allowed values of TypesSortDirection enum
var AllowedTypesSortDirectionEnumValues = []TypesSortDirection{
	"asc",
	"desc",
}

func (v *TypesSortDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesSortDirection(value)
	for _, existing := range AllowedTypesSortDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesSortDirection", value)
}

// NewTypesSortDirectionFromValue returns a pointer to a valid TypesSortDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesSortDirectionFromValue(v string) (*TypesSortDirection, error) {
	ev := TypesSortDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesSortDirection: valid values are %v", v, AllowedTypesSortDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesSortDirection) IsValid() bool {
	for _, existing := range AllowedTypesSortDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.SortDirection value
func (v TypesSortDirection) Ptr() *TypesSortDirection {
	return &v
}

type NullableTypesSortDirection struct {
	value *TypesSortDirection
	isSet bool
}

func (v NullableTypesSortDirection) Get() *TypesSortDirection {
	return v.value
}

func (v *NullableTypesSortDirection) Set(val *TypesSortDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesSortDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesSortDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesSortDirection(val *TypesSortDirection) *NullableTypesSortDirection {
	return &NullableTypesSortDirection{value: val, isSet: true}
}

func (v NullableTypesSortDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesSortDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

