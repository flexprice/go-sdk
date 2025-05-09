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

// checks if the DtoAggregatedFeature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoAggregatedFeature{}

// DtoAggregatedFeature struct for DtoAggregatedFeature
type DtoAggregatedFeature struct {
	Entitlement *DtoAggregatedEntitlement `json:"entitlement,omitempty"`
	Feature *DtoFeatureResponse `json:"feature,omitempty"`
	Sources []DtoEntitlementSource `json:"sources,omitempty"`
}

// NewDtoAggregatedFeature instantiates a new DtoAggregatedFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoAggregatedFeature() *DtoAggregatedFeature {
	this := DtoAggregatedFeature{}
	return &this
}

// NewDtoAggregatedFeatureWithDefaults instantiates a new DtoAggregatedFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoAggregatedFeatureWithDefaults() *DtoAggregatedFeature {
	this := DtoAggregatedFeature{}
	return &this
}

// GetEntitlement returns the Entitlement field value if set, zero value otherwise.
func (o *DtoAggregatedFeature) GetEntitlement() DtoAggregatedEntitlement {
	if o == nil || IsNil(o.Entitlement) {
		var ret DtoAggregatedEntitlement
		return ret
	}
	return *o.Entitlement
}

// GetEntitlementOk returns a tuple with the Entitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoAggregatedFeature) GetEntitlementOk() (*DtoAggregatedEntitlement, bool) {
	if o == nil || IsNil(o.Entitlement) {
		return nil, false
	}
	return o.Entitlement, true
}

// HasEntitlement returns a boolean if a field has been set.
func (o *DtoAggregatedFeature) HasEntitlement() bool {
	if o != nil && !IsNil(o.Entitlement) {
		return true
	}

	return false
}

// SetEntitlement gets a reference to the given DtoAggregatedEntitlement and assigns it to the Entitlement field.
func (o *DtoAggregatedFeature) SetEntitlement(v DtoAggregatedEntitlement) {
	o.Entitlement = &v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *DtoAggregatedFeature) GetFeature() DtoFeatureResponse {
	if o == nil || IsNil(o.Feature) {
		var ret DtoFeatureResponse
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoAggregatedFeature) GetFeatureOk() (*DtoFeatureResponse, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *DtoAggregatedFeature) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given DtoFeatureResponse and assigns it to the Feature field.
func (o *DtoAggregatedFeature) SetFeature(v DtoFeatureResponse) {
	o.Feature = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *DtoAggregatedFeature) GetSources() []DtoEntitlementSource {
	if o == nil || IsNil(o.Sources) {
		var ret []DtoEntitlementSource
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoAggregatedFeature) GetSourcesOk() ([]DtoEntitlementSource, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *DtoAggregatedFeature) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []DtoEntitlementSource and assigns it to the Sources field.
func (o *DtoAggregatedFeature) SetSources(v []DtoEntitlementSource) {
	o.Sources = v
}

func (o DtoAggregatedFeature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoAggregatedFeature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entitlement) {
		toSerialize["entitlement"] = o.Entitlement
	}
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	return toSerialize, nil
}

type NullableDtoAggregatedFeature struct {
	value *DtoAggregatedFeature
	isSet bool
}

func (v NullableDtoAggregatedFeature) Get() *DtoAggregatedFeature {
	return v.value
}

func (v *NullableDtoAggregatedFeature) Set(val *DtoAggregatedFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoAggregatedFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoAggregatedFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoAggregatedFeature(val *DtoAggregatedFeature) *NullableDtoAggregatedFeature {
	return &NullableDtoAggregatedFeature{value: val, isSet: true}
}

func (v NullableDtoAggregatedFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoAggregatedFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


