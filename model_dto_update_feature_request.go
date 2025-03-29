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

// checks if the DtoUpdateFeatureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoUpdateFeatureRequest{}

// DtoUpdateFeatureRequest struct for DtoUpdateFeatureRequest
type DtoUpdateFeatureRequest struct {
	Description *string `json:"description,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	UnitPlural *string `json:"unit_plural,omitempty"`
	UnitSingular *string `json:"unit_singular,omitempty"`
}

// NewDtoUpdateFeatureRequest instantiates a new DtoUpdateFeatureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoUpdateFeatureRequest() *DtoUpdateFeatureRequest {
	this := DtoUpdateFeatureRequest{}
	return &this
}

// NewDtoUpdateFeatureRequestWithDefaults instantiates a new DtoUpdateFeatureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoUpdateFeatureRequestWithDefaults() *DtoUpdateFeatureRequest {
	this := DtoUpdateFeatureRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DtoUpdateFeatureRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoUpdateFeatureRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DtoUpdateFeatureRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DtoUpdateFeatureRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DtoUpdateFeatureRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoUpdateFeatureRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DtoUpdateFeatureRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DtoUpdateFeatureRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DtoUpdateFeatureRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoUpdateFeatureRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DtoUpdateFeatureRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DtoUpdateFeatureRequest) SetName(v string) {
	o.Name = &v
}

// GetUnitPlural returns the UnitPlural field value if set, zero value otherwise.
func (o *DtoUpdateFeatureRequest) GetUnitPlural() string {
	if o == nil || IsNil(o.UnitPlural) {
		var ret string
		return ret
	}
	return *o.UnitPlural
}

// GetUnitPluralOk returns a tuple with the UnitPlural field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoUpdateFeatureRequest) GetUnitPluralOk() (*string, bool) {
	if o == nil || IsNil(o.UnitPlural) {
		return nil, false
	}
	return o.UnitPlural, true
}

// HasUnitPlural returns a boolean if a field has been set.
func (o *DtoUpdateFeatureRequest) HasUnitPlural() bool {
	if o != nil && !IsNil(o.UnitPlural) {
		return true
	}

	return false
}

// SetUnitPlural gets a reference to the given string and assigns it to the UnitPlural field.
func (o *DtoUpdateFeatureRequest) SetUnitPlural(v string) {
	o.UnitPlural = &v
}

// GetUnitSingular returns the UnitSingular field value if set, zero value otherwise.
func (o *DtoUpdateFeatureRequest) GetUnitSingular() string {
	if o == nil || IsNil(o.UnitSingular) {
		var ret string
		return ret
	}
	return *o.UnitSingular
}

// GetUnitSingularOk returns a tuple with the UnitSingular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoUpdateFeatureRequest) GetUnitSingularOk() (*string, bool) {
	if o == nil || IsNil(o.UnitSingular) {
		return nil, false
	}
	return o.UnitSingular, true
}

// HasUnitSingular returns a boolean if a field has been set.
func (o *DtoUpdateFeatureRequest) HasUnitSingular() bool {
	if o != nil && !IsNil(o.UnitSingular) {
		return true
	}

	return false
}

// SetUnitSingular gets a reference to the given string and assigns it to the UnitSingular field.
func (o *DtoUpdateFeatureRequest) SetUnitSingular(v string) {
	o.UnitSingular = &v
}

func (o DtoUpdateFeatureRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoUpdateFeatureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UnitPlural) {
		toSerialize["unit_plural"] = o.UnitPlural
	}
	if !IsNil(o.UnitSingular) {
		toSerialize["unit_singular"] = o.UnitSingular
	}
	return toSerialize, nil
}

type NullableDtoUpdateFeatureRequest struct {
	value *DtoUpdateFeatureRequest
	isSet bool
}

func (v NullableDtoUpdateFeatureRequest) Get() *DtoUpdateFeatureRequest {
	return v.value
}

func (v *NullableDtoUpdateFeatureRequest) Set(val *DtoUpdateFeatureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoUpdateFeatureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoUpdateFeatureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoUpdateFeatureRequest(val *DtoUpdateFeatureRequest) *NullableDtoUpdateFeatureRequest {
	return &NullableDtoUpdateFeatureRequest{value: val, isSet: true}
}

func (v NullableDtoUpdateFeatureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoUpdateFeatureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


