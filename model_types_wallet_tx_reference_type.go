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

// TypesWalletTxReferenceType the model 'TypesWalletTxReferenceType'
type TypesWalletTxReferenceType string

// List of types.WalletTxReferenceType
const (
	TYPESWALLETTXREFERENCETYPE_WalletTxReferenceTypeInvoice TypesWalletTxReferenceType = "INVOICE"
	TYPESWALLETTXREFERENCETYPE_WalletTxReferenceTypePayment TypesWalletTxReferenceType = "PAYMENT"
	TYPESWALLETTXREFERENCETYPE_WalletTxReferenceTypeExternal TypesWalletTxReferenceType = "EXTERNAL"
	TYPESWALLETTXREFERENCETYPE_WalletTxReferenceTypeRequest TypesWalletTxReferenceType = "REQUEST"
)

// All allowed values of TypesWalletTxReferenceType enum
var AllowedTypesWalletTxReferenceTypeEnumValues = []TypesWalletTxReferenceType{
	"INVOICE",
	"PAYMENT",
	"EXTERNAL",
	"REQUEST",
}

func (v *TypesWalletTxReferenceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesWalletTxReferenceType(value)
	for _, existing := range AllowedTypesWalletTxReferenceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesWalletTxReferenceType", value)
}

// NewTypesWalletTxReferenceTypeFromValue returns a pointer to a valid TypesWalletTxReferenceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesWalletTxReferenceTypeFromValue(v string) (*TypesWalletTxReferenceType, error) {
	ev := TypesWalletTxReferenceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesWalletTxReferenceType: valid values are %v", v, AllowedTypesWalletTxReferenceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesWalletTxReferenceType) IsValid() bool {
	for _, existing := range AllowedTypesWalletTxReferenceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.WalletTxReferenceType value
func (v TypesWalletTxReferenceType) Ptr() *TypesWalletTxReferenceType {
	return &v
}

type NullableTypesWalletTxReferenceType struct {
	value *TypesWalletTxReferenceType
	isSet bool
}

func (v NullableTypesWalletTxReferenceType) Get() *TypesWalletTxReferenceType {
	return v.value
}

func (v *NullableTypesWalletTxReferenceType) Set(val *TypesWalletTxReferenceType) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesWalletTxReferenceType) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesWalletTxReferenceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesWalletTxReferenceType(val *TypesWalletTxReferenceType) *NullableTypesWalletTxReferenceType {
	return &NullableTypesWalletTxReferenceType{value: val, isSet: true}
}

func (v NullableTypesWalletTxReferenceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesWalletTxReferenceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

