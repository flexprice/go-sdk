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

// TypesEntityType the model 'TypesEntityType'
type TypesEntityType string

// List of types.EntityType
const (
	TYPESENTITYTYPE_EntityTypeEvents TypesEntityType = "EVENTS"
	TYPESENTITYTYPE_EntityTypePrices TypesEntityType = "PRICES"
	TYPESENTITYTYPE_EntityTypeCustomers TypesEntityType = "CUSTOMERS"
)

// All allowed values of TypesEntityType enum
var AllowedTypesEntityTypeEnumValues = []TypesEntityType{
	"EVENTS",
	"PRICES",
	"CUSTOMERS",
}

func (v *TypesEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesEntityType(value)
	for _, existing := range AllowedTypesEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesEntityType", value)
}

// NewTypesEntityTypeFromValue returns a pointer to a valid TypesEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesEntityTypeFromValue(v string) (*TypesEntityType, error) {
	ev := TypesEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesEntityType: valid values are %v", v, AllowedTypesEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesEntityType) IsValid() bool {
	for _, existing := range AllowedTypesEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.EntityType value
func (v TypesEntityType) Ptr() *TypesEntityType {
	return &v
}

type NullableTypesEntityType struct {
	value *TypesEntityType
	isSet bool
}

func (v NullableTypesEntityType) Get() *TypesEntityType {
	return v.value
}

func (v *NullableTypesEntityType) Set(val *TypesEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesEntityType(val *TypesEntityType) *NullableTypesEntityType {
	return &NullableTypesEntityType{value: val, isSet: true}
}

func (v NullableTypesEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

