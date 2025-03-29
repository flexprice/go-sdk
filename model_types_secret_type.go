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

// TypesSecretType the model 'TypesSecretType'
type TypesSecretType string

// List of types.SecretType
const (
	TYPESSECRETTYPE_SecretTypePrivateKey TypesSecretType = "private_key"
	TYPESSECRETTYPE_SecretTypePublishableKey TypesSecretType = "publishable_key"
	TYPESSECRETTYPE_SecretTypeIntegration TypesSecretType = "integration"
)

// All allowed values of TypesSecretType enum
var AllowedTypesSecretTypeEnumValues = []TypesSecretType{
	"private_key",
	"publishable_key",
	"integration",
}

func (v *TypesSecretType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesSecretType(value)
	for _, existing := range AllowedTypesSecretTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesSecretType", value)
}

// NewTypesSecretTypeFromValue returns a pointer to a valid TypesSecretType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesSecretTypeFromValue(v string) (*TypesSecretType, error) {
	ev := TypesSecretType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesSecretType: valid values are %v", v, AllowedTypesSecretTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesSecretType) IsValid() bool {
	for _, existing := range AllowedTypesSecretTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.SecretType value
func (v TypesSecretType) Ptr() *TypesSecretType {
	return &v
}

type NullableTypesSecretType struct {
	value *TypesSecretType
	isSet bool
}

func (v NullableTypesSecretType) Get() *TypesSecretType {
	return v.value
}

func (v *NullableTypesSecretType) Set(val *TypesSecretType) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesSecretType) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesSecretType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesSecretType(val *TypesSecretType) *NullableTypesSecretType {
	return &NullableTypesSecretType{value: val, isSet: true}
}

func (v NullableTypesSecretType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesSecretType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

