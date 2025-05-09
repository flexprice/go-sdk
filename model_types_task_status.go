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

// TypesTaskStatus the model 'TypesTaskStatus'
type TypesTaskStatus string

// List of types.TaskStatus
const (
	TYPESTASKSTATUS_TaskStatusPending TypesTaskStatus = "PENDING"
	TYPESTASKSTATUS_TaskStatusProcessing TypesTaskStatus = "PROCESSING"
	TYPESTASKSTATUS_TaskStatusCompleted TypesTaskStatus = "COMPLETED"
	TYPESTASKSTATUS_TaskStatusFailed TypesTaskStatus = "FAILED"
)

// All allowed values of TypesTaskStatus enum
var AllowedTypesTaskStatusEnumValues = []TypesTaskStatus{
	"PENDING",
	"PROCESSING",
	"COMPLETED",
	"FAILED",
}

func (v *TypesTaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesTaskStatus(value)
	for _, existing := range AllowedTypesTaskStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesTaskStatus", value)
}

// NewTypesTaskStatusFromValue returns a pointer to a valid TypesTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesTaskStatusFromValue(v string) (*TypesTaskStatus, error) {
	ev := TypesTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesTaskStatus: valid values are %v", v, AllowedTypesTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesTaskStatus) IsValid() bool {
	for _, existing := range AllowedTypesTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.TaskStatus value
func (v TypesTaskStatus) Ptr() *TypesTaskStatus {
	return &v
}

type NullableTypesTaskStatus struct {
	value *TypesTaskStatus
	isSet bool
}

func (v NullableTypesTaskStatus) Get() *TypesTaskStatus {
	return v.value
}

func (v *NullableTypesTaskStatus) Set(val *TypesTaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesTaskStatus(val *TypesTaskStatus) *NullableTypesTaskStatus {
	return &NullableTypesTaskStatus{value: val, isSet: true}
}

func (v NullableTypesTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

