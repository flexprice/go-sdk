/*
FlexPrice API

FlexPrice API Service

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flexprice

import (
	"encoding/json"
)

// checks if the MeterAggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeterAggregation{}

// MeterAggregation struct for MeterAggregation
type MeterAggregation struct {
	// Field is the key in $event.properties on which the aggregation is to be applied For ex if the aggregation type is sum for API usage, the field could be \"duration_ms\"
	Field *string `json:"field,omitempty"`
	Type *TypesAggregationType `json:"type,omitempty"`
}

// NewMeterAggregation instantiates a new MeterAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeterAggregation() *MeterAggregation {
	this := MeterAggregation{}
	return &this
}

// NewMeterAggregationWithDefaults instantiates a new MeterAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeterAggregationWithDefaults() *MeterAggregation {
	this := MeterAggregation{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *MeterAggregation) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeterAggregation) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *MeterAggregation) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *MeterAggregation) SetField(v string) {
	o.Field = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MeterAggregation) GetType() TypesAggregationType {
	if o == nil || IsNil(o.Type) {
		var ret TypesAggregationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeterAggregation) GetTypeOk() (*TypesAggregationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MeterAggregation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TypesAggregationType and assigns it to the Type field.
func (o *MeterAggregation) SetType(v TypesAggregationType) {
	o.Type = &v
}

func (o MeterAggregation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeterAggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableMeterAggregation struct {
	value *MeterAggregation
	isSet bool
}

func (v NullableMeterAggregation) Get() *MeterAggregation {
	return v.value
}

func (v *NullableMeterAggregation) Set(val *MeterAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableMeterAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableMeterAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeterAggregation(val *MeterAggregation) *NullableMeterAggregation {
	return &NullableMeterAggregation{value: val, isSet: true}
}

func (v NullableMeterAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeterAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


