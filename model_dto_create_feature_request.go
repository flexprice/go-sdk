/*
FlexPrice API

FlexPrice API Service

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flexpriceclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DtoCreateFeatureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoCreateFeatureRequest{}

// DtoCreateFeatureRequest struct for DtoCreateFeatureRequest
type DtoCreateFeatureRequest struct {
	Description *string `json:"description,omitempty"`
	LookupKey *string `json:"lookup_key,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	MeterId *string `json:"meter_id,omitempty"`
	Name string `json:"name"`
	Type TypesFeatureType `json:"type"`
	UnitPlural *string `json:"unit_plural,omitempty"`
	UnitSingular *string `json:"unit_singular,omitempty"`
}

type _DtoCreateFeatureRequest DtoCreateFeatureRequest

// NewDtoCreateFeatureRequest instantiates a new DtoCreateFeatureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoCreateFeatureRequest(name string, type_ TypesFeatureType) *DtoCreateFeatureRequest {
	this := DtoCreateFeatureRequest{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewDtoCreateFeatureRequestWithDefaults instantiates a new DtoCreateFeatureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoCreateFeatureRequestWithDefaults() *DtoCreateFeatureRequest {
	this := DtoCreateFeatureRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DtoCreateFeatureRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreateFeatureRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DtoCreateFeatureRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DtoCreateFeatureRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLookupKey returns the LookupKey field value if set, zero value otherwise.
func (o *DtoCreateFeatureRequest) GetLookupKey() string {
	if o == nil || IsNil(o.LookupKey) {
		var ret string
		return ret
	}
	return *o.LookupKey
}

// GetLookupKeyOk returns a tuple with the LookupKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreateFeatureRequest) GetLookupKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LookupKey) {
		return nil, false
	}
	return o.LookupKey, true
}

// HasLookupKey returns a boolean if a field has been set.
func (o *DtoCreateFeatureRequest) HasLookupKey() bool {
	if o != nil && !IsNil(o.LookupKey) {
		return true
	}

	return false
}

// SetLookupKey gets a reference to the given string and assigns it to the LookupKey field.
func (o *DtoCreateFeatureRequest) SetLookupKey(v string) {
	o.LookupKey = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DtoCreateFeatureRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreateFeatureRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DtoCreateFeatureRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DtoCreateFeatureRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMeterId returns the MeterId field value if set, zero value otherwise.
func (o *DtoCreateFeatureRequest) GetMeterId() string {
	if o == nil || IsNil(o.MeterId) {
		var ret string
		return ret
	}
	return *o.MeterId
}

// GetMeterIdOk returns a tuple with the MeterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreateFeatureRequest) GetMeterIdOk() (*string, bool) {
	if o == nil || IsNil(o.MeterId) {
		return nil, false
	}
	return o.MeterId, true
}

// HasMeterId returns a boolean if a field has been set.
func (o *DtoCreateFeatureRequest) HasMeterId() bool {
	if o != nil && !IsNil(o.MeterId) {
		return true
	}

	return false
}

// SetMeterId gets a reference to the given string and assigns it to the MeterId field.
func (o *DtoCreateFeatureRequest) SetMeterId(v string) {
	o.MeterId = &v
}

// GetName returns the Name field value
func (o *DtoCreateFeatureRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DtoCreateFeatureRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DtoCreateFeatureRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *DtoCreateFeatureRequest) GetType() TypesFeatureType {
	if o == nil {
		var ret TypesFeatureType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DtoCreateFeatureRequest) GetTypeOk() (*TypesFeatureType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DtoCreateFeatureRequest) SetType(v TypesFeatureType) {
	o.Type = v
}

// GetUnitPlural returns the UnitPlural field value if set, zero value otherwise.
func (o *DtoCreateFeatureRequest) GetUnitPlural() string {
	if o == nil || IsNil(o.UnitPlural) {
		var ret string
		return ret
	}
	return *o.UnitPlural
}

// GetUnitPluralOk returns a tuple with the UnitPlural field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreateFeatureRequest) GetUnitPluralOk() (*string, bool) {
	if o == nil || IsNil(o.UnitPlural) {
		return nil, false
	}
	return o.UnitPlural, true
}

// HasUnitPlural returns a boolean if a field has been set.
func (o *DtoCreateFeatureRequest) HasUnitPlural() bool {
	if o != nil && !IsNil(o.UnitPlural) {
		return true
	}

	return false
}

// SetUnitPlural gets a reference to the given string and assigns it to the UnitPlural field.
func (o *DtoCreateFeatureRequest) SetUnitPlural(v string) {
	o.UnitPlural = &v
}

// GetUnitSingular returns the UnitSingular field value if set, zero value otherwise.
func (o *DtoCreateFeatureRequest) GetUnitSingular() string {
	if o == nil || IsNil(o.UnitSingular) {
		var ret string
		return ret
	}
	return *o.UnitSingular
}

// GetUnitSingularOk returns a tuple with the UnitSingular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreateFeatureRequest) GetUnitSingularOk() (*string, bool) {
	if o == nil || IsNil(o.UnitSingular) {
		return nil, false
	}
	return o.UnitSingular, true
}

// HasUnitSingular returns a boolean if a field has been set.
func (o *DtoCreateFeatureRequest) HasUnitSingular() bool {
	if o != nil && !IsNil(o.UnitSingular) {
		return true
	}

	return false
}

// SetUnitSingular gets a reference to the given string and assigns it to the UnitSingular field.
func (o *DtoCreateFeatureRequest) SetUnitSingular(v string) {
	o.UnitSingular = &v
}

func (o DtoCreateFeatureRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoCreateFeatureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LookupKey) {
		toSerialize["lookup_key"] = o.LookupKey
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MeterId) {
		toSerialize["meter_id"] = o.MeterId
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.UnitPlural) {
		toSerialize["unit_plural"] = o.UnitPlural
	}
	if !IsNil(o.UnitSingular) {
		toSerialize["unit_singular"] = o.UnitSingular
	}
	return toSerialize, nil
}

func (o *DtoCreateFeatureRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDtoCreateFeatureRequest := _DtoCreateFeatureRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDtoCreateFeatureRequest)

	if err != nil {
		return err
	}

	*o = DtoCreateFeatureRequest(varDtoCreateFeatureRequest)

	return err
}

type NullableDtoCreateFeatureRequest struct {
	value *DtoCreateFeatureRequest
	isSet bool
}

func (v NullableDtoCreateFeatureRequest) Get() *DtoCreateFeatureRequest {
	return v.value
}

func (v *NullableDtoCreateFeatureRequest) Set(val *DtoCreateFeatureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoCreateFeatureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoCreateFeatureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoCreateFeatureRequest(val *DtoCreateFeatureRequest) *NullableDtoCreateFeatureRequest {
	return &NullableDtoCreateFeatureRequest{value: val, isSet: true}
}

func (v NullableDtoCreateFeatureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoCreateFeatureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


