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

// checks if the PriceJSONBTransformQuantity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceJSONBTransformQuantity{}

// PriceJSONBTransformQuantity struct for PriceJSONBTransformQuantity
type PriceJSONBTransformQuantity struct {
	// Divide quantity by this number
	DivideBy *int32 `json:"divide_by,omitempty"`
	// up or down
	Round *string `json:"round,omitempty"`
}

// NewPriceJSONBTransformQuantity instantiates a new PriceJSONBTransformQuantity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceJSONBTransformQuantity() *PriceJSONBTransformQuantity {
	this := PriceJSONBTransformQuantity{}
	return &this
}

// NewPriceJSONBTransformQuantityWithDefaults instantiates a new PriceJSONBTransformQuantity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceJSONBTransformQuantityWithDefaults() *PriceJSONBTransformQuantity {
	this := PriceJSONBTransformQuantity{}
	return &this
}

// GetDivideBy returns the DivideBy field value if set, zero value otherwise.
func (o *PriceJSONBTransformQuantity) GetDivideBy() int32 {
	if o == nil || IsNil(o.DivideBy) {
		var ret int32
		return ret
	}
	return *o.DivideBy
}

// GetDivideByOk returns a tuple with the DivideBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceJSONBTransformQuantity) GetDivideByOk() (*int32, bool) {
	if o == nil || IsNil(o.DivideBy) {
		return nil, false
	}
	return o.DivideBy, true
}

// HasDivideBy returns a boolean if a field has been set.
func (o *PriceJSONBTransformQuantity) HasDivideBy() bool {
	if o != nil && !IsNil(o.DivideBy) {
		return true
	}

	return false
}

// SetDivideBy gets a reference to the given int32 and assigns it to the DivideBy field.
func (o *PriceJSONBTransformQuantity) SetDivideBy(v int32) {
	o.DivideBy = &v
}

// GetRound returns the Round field value if set, zero value otherwise.
func (o *PriceJSONBTransformQuantity) GetRound() string {
	if o == nil || IsNil(o.Round) {
		var ret string
		return ret
	}
	return *o.Round
}

// GetRoundOk returns a tuple with the Round field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceJSONBTransformQuantity) GetRoundOk() (*string, bool) {
	if o == nil || IsNil(o.Round) {
		return nil, false
	}
	return o.Round, true
}

// HasRound returns a boolean if a field has been set.
func (o *PriceJSONBTransformQuantity) HasRound() bool {
	if o != nil && !IsNil(o.Round) {
		return true
	}

	return false
}

// SetRound gets a reference to the given string and assigns it to the Round field.
func (o *PriceJSONBTransformQuantity) SetRound(v string) {
	o.Round = &v
}

func (o PriceJSONBTransformQuantity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceJSONBTransformQuantity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DivideBy) {
		toSerialize["divide_by"] = o.DivideBy
	}
	if !IsNil(o.Round) {
		toSerialize["round"] = o.Round
	}
	return toSerialize, nil
}

type NullablePriceJSONBTransformQuantity struct {
	value *PriceJSONBTransformQuantity
	isSet bool
}

func (v NullablePriceJSONBTransformQuantity) Get() *PriceJSONBTransformQuantity {
	return v.value
}

func (v *NullablePriceJSONBTransformQuantity) Set(val *PriceJSONBTransformQuantity) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceJSONBTransformQuantity) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceJSONBTransformQuantity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceJSONBTransformQuantity(val *PriceJSONBTransformQuantity) *NullablePriceJSONBTransformQuantity {
	return &NullablePriceJSONBTransformQuantity{value: val, isSet: true}
}

func (v NullablePriceJSONBTransformQuantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceJSONBTransformQuantity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


