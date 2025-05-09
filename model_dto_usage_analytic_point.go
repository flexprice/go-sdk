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

// checks if the DtoUsageAnalyticPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoUsageAnalyticPoint{}

// DtoUsageAnalyticPoint struct for DtoUsageAnalyticPoint
type DtoUsageAnalyticPoint struct {
	Cost *float32 `json:"cost,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Usage *float32 `json:"usage,omitempty"`
}

// NewDtoUsageAnalyticPoint instantiates a new DtoUsageAnalyticPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoUsageAnalyticPoint() *DtoUsageAnalyticPoint {
	this := DtoUsageAnalyticPoint{}
	return &this
}

// NewDtoUsageAnalyticPointWithDefaults instantiates a new DtoUsageAnalyticPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoUsageAnalyticPointWithDefaults() *DtoUsageAnalyticPoint {
	this := DtoUsageAnalyticPoint{}
	return &this
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *DtoUsageAnalyticPoint) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoUsageAnalyticPoint) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *DtoUsageAnalyticPoint) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *DtoUsageAnalyticPoint) SetCost(v float32) {
	o.Cost = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DtoUsageAnalyticPoint) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoUsageAnalyticPoint) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DtoUsageAnalyticPoint) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *DtoUsageAnalyticPoint) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *DtoUsageAnalyticPoint) GetUsage() float32 {
	if o == nil || IsNil(o.Usage) {
		var ret float32
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoUsageAnalyticPoint) GetUsageOk() (*float32, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *DtoUsageAnalyticPoint) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given float32 and assigns it to the Usage field.
func (o *DtoUsageAnalyticPoint) SetUsage(v float32) {
	o.Usage = &v
}

func (o DtoUsageAnalyticPoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoUsageAnalyticPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableDtoUsageAnalyticPoint struct {
	value *DtoUsageAnalyticPoint
	isSet bool
}

func (v NullableDtoUsageAnalyticPoint) Get() *DtoUsageAnalyticPoint {
	return v.value
}

func (v *NullableDtoUsageAnalyticPoint) Set(val *DtoUsageAnalyticPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoUsageAnalyticPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoUsageAnalyticPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoUsageAnalyticPoint(val *DtoUsageAnalyticPoint) *NullableDtoUsageAnalyticPoint {
	return &NullableDtoUsageAnalyticPoint{value: val, isSet: true}
}

func (v NullableDtoUsageAnalyticPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoUsageAnalyticPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


