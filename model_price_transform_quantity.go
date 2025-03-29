/*
FlexPrice API

FlexPrice API Service

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flexpriceclient

import (
	"encoding/json"
)

// checks if the PriceTransformQuantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceTransformQuantity{}

// PriceTransformQuantity struct for PriceTransformQuantity
type PriceTransformQuantity struct {
	// Divide quantity by this number
	DivideBy *int32 `json:"divide_by,omitempty"`
	// up or down
	Round *string `json:"round,omitempty"`
}

// NewPriceTransformQuantity instantiates a new PriceTransformQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceTransformQuantity() *PriceTransformQuantity {
	this := PriceTransformQuantity{}
	return &this
}

// NewPriceTransformQuantityWithDefaults instantiates a new PriceTransformQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceTransformQuantityWithDefaults() *PriceTransformQuantity {
	this := PriceTransformQuantity{}
	return &this
}

// GetDivideBy returns the DivideBy field value if set, zero value otherwise.
func (o *PriceTransformQuantity) GetDivideBy() int32 {
	if o == nil || IsNil(o.DivideBy) {
		var ret int32
		return ret
	}
	return *o.DivideBy
}

// GetDivideByOk returns a tuple with the DivideBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceTransformQuantity) GetDivideByOk() (*int32, bool) {
	if o == nil || IsNil(o.DivideBy) {
		return nil, false
	}
	return o.DivideBy, true
}

// HasDivideBy returns a boolean if a field has been set.
func (o *PriceTransformQuantity) HasDivideBy() bool {
	if o != nil && !IsNil(o.DivideBy) {
		return true
	}

	return false
}

// SetDivideBy gets a reference to the given int32 and assigns it to the DivideBy field.
func (o *PriceTransformQuantity) SetDivideBy(v int32) {
	o.DivideBy = &v
}

// GetRound returns the Round field value if set, zero value otherwise.
func (o *PriceTransformQuantity) GetRound() string {
	if o == nil || IsNil(o.Round) {
		var ret string
		return ret
	}
	return *o.Round
}

// GetRoundOk returns a tuple with the Round field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceTransformQuantity) GetRoundOk() (*string, bool) {
	if o == nil || IsNil(o.Round) {
		return nil, false
	}
	return o.Round, true
}

// HasRound returns a boolean if a field has been set.
func (o *PriceTransformQuantity) HasRound() bool {
	if o != nil && !IsNil(o.Round) {
		return true
	}

	return false
}

// SetRound gets a reference to the given string and assigns it to the Round field.
func (o *PriceTransformQuantity) SetRound(v string) {
	o.Round = &v
}

func (o PriceTransformQuantity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceTransformQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DivideBy) {
		toSerialize["divide_by"] = o.DivideBy
	}
	if !IsNil(o.Round) {
		toSerialize["round"] = o.Round
	}
	return toSerialize, nil
}

type NullablePriceTransformQuantity struct {
	value *PriceTransformQuantity
	isSet bool
}

func (v NullablePriceTransformQuantity) Get() *PriceTransformQuantity {
	return v.value
}

func (v *NullablePriceTransformQuantity) Set(val *PriceTransformQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceTransformQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceTransformQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceTransformQuantity(val *PriceTransformQuantity) *NullablePriceTransformQuantity {
	return &NullablePriceTransformQuantity{value: val, isSet: true}
}

func (v NullablePriceTransformQuantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceTransformQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


