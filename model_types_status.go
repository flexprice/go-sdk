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

// TypesStatus the model 'TypesStatus'
type TypesStatus string

// List of types.Status
const (
	TYPESSTATUS_StatusPublished TypesStatus = "published"
	TYPESSTATUS_StatusDeleted TypesStatus = "deleted"
	TYPESSTATUS_StatusArchived TypesStatus = "archived"
)

// All allowed values of TypesStatus enum
var AllowedTypesStatusEnumValues = []TypesStatus{
	"published",
	"deleted",
	"archived",
}

func (v *TypesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesStatus(value)
	for _, existing := range AllowedTypesStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesStatus", value)
}

// NewTypesStatusFromValue returns a pointer to a valid TypesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesStatusFromValue(v string) (*TypesStatus, error) {
	ev := TypesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesStatus: valid values are %v", v, AllowedTypesStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesStatus) IsValid() bool {
	for _, existing := range AllowedTypesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.Status value
func (v TypesStatus) Ptr() *TypesStatus {
	return &v
}

type NullableTypesStatus struct {
	value *TypesStatus
	isSet bool
}

func (v NullableTypesStatus) Get() *TypesStatus {
	return v.value
}

func (v *NullableTypesStatus) Set(val *TypesStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesStatus(val *TypesStatus) *NullableTypesStatus {
	return &NullableTypesStatus{value: val, isSet: true}
}

func (v NullableTypesStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

