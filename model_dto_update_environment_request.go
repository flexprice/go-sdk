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

// checks if the DtoUpdateEnvironmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoUpdateEnvironmentRequest{}

// DtoUpdateEnvironmentRequest struct for DtoUpdateEnvironmentRequest
type DtoUpdateEnvironmentRequest struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewDtoUpdateEnvironmentRequest instantiates a new DtoUpdateEnvironmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoUpdateEnvironmentRequest() *DtoUpdateEnvironmentRequest {
	this := DtoUpdateEnvironmentRequest{}
	return &this
}

// NewDtoUpdateEnvironmentRequestWithDefaults instantiates a new DtoUpdateEnvironmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoUpdateEnvironmentRequestWithDefaults() *DtoUpdateEnvironmentRequest {
	this := DtoUpdateEnvironmentRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DtoUpdateEnvironmentRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoUpdateEnvironmentRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DtoUpdateEnvironmentRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DtoUpdateEnvironmentRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DtoUpdateEnvironmentRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoUpdateEnvironmentRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DtoUpdateEnvironmentRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DtoUpdateEnvironmentRequest) SetType(v string) {
	o.Type = &v
}

func (o DtoUpdateEnvironmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoUpdateEnvironmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDtoUpdateEnvironmentRequest struct {
	value *DtoUpdateEnvironmentRequest
	isSet bool
}

func (v NullableDtoUpdateEnvironmentRequest) Get() *DtoUpdateEnvironmentRequest {
	return v.value
}

func (v *NullableDtoUpdateEnvironmentRequest) Set(val *DtoUpdateEnvironmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoUpdateEnvironmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoUpdateEnvironmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoUpdateEnvironmentRequest(val *DtoUpdateEnvironmentRequest) *NullableDtoUpdateEnvironmentRequest {
	return &NullableDtoUpdateEnvironmentRequest{value: val, isSet: true}
}

func (v NullableDtoUpdateEnvironmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoUpdateEnvironmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


