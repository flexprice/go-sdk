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

// TypesWalletType the model 'TypesWalletType'
type TypesWalletType string

// List of types.WalletType
const (
	TYPESWALLETTYPE_WalletTypePromotional TypesWalletType = "PROMOTIONAL"
	TYPESWALLETTYPE_WalletTypePrePaid TypesWalletType = "PRE_PAID"
)

// All allowed values of TypesWalletType enum
var AllowedTypesWalletTypeEnumValues = []TypesWalletType{
	"PROMOTIONAL",
	"PRE_PAID",
}

func (v *TypesWalletType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesWalletType(value)
	for _, existing := range AllowedTypesWalletTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesWalletType", value)
}

// NewTypesWalletTypeFromValue returns a pointer to a valid TypesWalletType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesWalletTypeFromValue(v string) (*TypesWalletType, error) {
	ev := TypesWalletType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesWalletType: valid values are %v", v, AllowedTypesWalletTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesWalletType) IsValid() bool {
	for _, existing := range AllowedTypesWalletTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.WalletType value
func (v TypesWalletType) Ptr() *TypesWalletType {
	return &v
}

type NullableTypesWalletType struct {
	value *TypesWalletType
	isSet bool
}

func (v NullableTypesWalletType) Get() *TypesWalletType {
	return v.value
}

func (v *NullableTypesWalletType) Set(val *TypesWalletType) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesWalletType) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesWalletType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesWalletType(val *TypesWalletType) *NullableTypesWalletType {
	return &NullableTypesWalletType{value: val, isSet: true}
}

func (v NullableTypesWalletType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesWalletType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

