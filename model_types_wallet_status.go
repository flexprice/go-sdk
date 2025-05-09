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

// TypesWalletStatus the model 'TypesWalletStatus'
type TypesWalletStatus string

// List of types.WalletStatus
const (
	TYPESWALLETSTATUS_WalletStatusActive TypesWalletStatus = "active"
	TYPESWALLETSTATUS_WalletStatusFrozen TypesWalletStatus = "frozen"
	TYPESWALLETSTATUS_WalletStatusClosed TypesWalletStatus = "closed"
)

// All allowed values of TypesWalletStatus enum
var AllowedTypesWalletStatusEnumValues = []TypesWalletStatus{
	"active",
	"frozen",
	"closed",
}

func (v *TypesWalletStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesWalletStatus(value)
	for _, existing := range AllowedTypesWalletStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesWalletStatus", value)
}

// NewTypesWalletStatusFromValue returns a pointer to a valid TypesWalletStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesWalletStatusFromValue(v string) (*TypesWalletStatus, error) {
	ev := TypesWalletStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesWalletStatus: valid values are %v", v, AllowedTypesWalletStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesWalletStatus) IsValid() bool {
	for _, existing := range AllowedTypesWalletStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.WalletStatus value
func (v TypesWalletStatus) Ptr() *TypesWalletStatus {
	return &v
}

type NullableTypesWalletStatus struct {
	value *TypesWalletStatus
	isSet bool
}

func (v NullableTypesWalletStatus) Get() *TypesWalletStatus {
	return v.value
}

func (v *NullableTypesWalletStatus) Set(val *TypesWalletStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesWalletStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesWalletStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesWalletStatus(val *TypesWalletStatus) *NullableTypesWalletStatus {
	return &NullableTypesWalletStatus{value: val, isSet: true}
}

func (v NullableTypesWalletStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesWalletStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

