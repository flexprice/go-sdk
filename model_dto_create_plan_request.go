/*
FlexPrice API

FlexPrice API Service

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package flexprice

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DtoCreatePlanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoCreatePlanRequest{}

// DtoCreatePlanRequest struct for DtoCreatePlanRequest
type DtoCreatePlanRequest struct {
	Description *string `json:"description,omitempty"`
	Entitlements []DtoCreatePlanEntitlementRequest `json:"entitlements,omitempty"`
	LookupKey *string `json:"lookup_key,omitempty"`
	Name string `json:"name"`
	Prices []DtoCreatePlanPriceRequest `json:"prices,omitempty"`
}

type _DtoCreatePlanRequest DtoCreatePlanRequest

// NewDtoCreatePlanRequest instantiates a new DtoCreatePlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoCreatePlanRequest(name string) *DtoCreatePlanRequest {
	this := DtoCreatePlanRequest{}
	this.Name = name
	return &this
}

// NewDtoCreatePlanRequestWithDefaults instantiates a new DtoCreatePlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoCreatePlanRequestWithDefaults() *DtoCreatePlanRequest {
	this := DtoCreatePlanRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DtoCreatePlanRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreatePlanRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DtoCreatePlanRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DtoCreatePlanRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *DtoCreatePlanRequest) GetEntitlements() []DtoCreatePlanEntitlementRequest {
	if o == nil || IsNil(o.Entitlements) {
		var ret []DtoCreatePlanEntitlementRequest
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreatePlanRequest) GetEntitlementsOk() ([]DtoCreatePlanEntitlementRequest, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *DtoCreatePlanRequest) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []DtoCreatePlanEntitlementRequest and assigns it to the Entitlements field.
func (o *DtoCreatePlanRequest) SetEntitlements(v []DtoCreatePlanEntitlementRequest) {
	o.Entitlements = v
}

// GetLookupKey returns the LookupKey field value if set, zero value otherwise.
func (o *DtoCreatePlanRequest) GetLookupKey() string {
	if o == nil || IsNil(o.LookupKey) {
		var ret string
		return ret
	}
	return *o.LookupKey
}

// GetLookupKeyOk returns a tuple with the LookupKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreatePlanRequest) GetLookupKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LookupKey) {
		return nil, false
	}
	return o.LookupKey, true
}

// HasLookupKey returns a boolean if a field has been set.
func (o *DtoCreatePlanRequest) HasLookupKey() bool {
	if o != nil && !IsNil(o.LookupKey) {
		return true
	}

	return false
}

// SetLookupKey gets a reference to the given string and assigns it to the LookupKey field.
func (o *DtoCreatePlanRequest) SetLookupKey(v string) {
	o.LookupKey = &v
}

// GetName returns the Name field value
func (o *DtoCreatePlanRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DtoCreatePlanRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DtoCreatePlanRequest) SetName(v string) {
	o.Name = v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *DtoCreatePlanRequest) GetPrices() []DtoCreatePlanPriceRequest {
	if o == nil || IsNil(o.Prices) {
		var ret []DtoCreatePlanPriceRequest
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreatePlanRequest) GetPricesOk() ([]DtoCreatePlanPriceRequest, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *DtoCreatePlanRequest) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []DtoCreatePlanPriceRequest and assigns it to the Prices field.
func (o *DtoCreatePlanRequest) SetPrices(v []DtoCreatePlanPriceRequest) {
	o.Prices = v
}

func (o DtoCreatePlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoCreatePlanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.LookupKey) {
		toSerialize["lookup_key"] = o.LookupKey
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	return toSerialize, nil
}

func (o *DtoCreatePlanRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varDtoCreatePlanRequest := _DtoCreatePlanRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDtoCreatePlanRequest)

	if err != nil {
		return err
	}

	*o = DtoCreatePlanRequest(varDtoCreatePlanRequest)

	return err
}

type NullableDtoCreatePlanRequest struct {
	value *DtoCreatePlanRequest
	isSet bool
}

func (v NullableDtoCreatePlanRequest) Get() *DtoCreatePlanRequest {
	return v.value
}

func (v *NullableDtoCreatePlanRequest) Set(val *DtoCreatePlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoCreatePlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoCreatePlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoCreatePlanRequest(val *DtoCreatePlanRequest) *NullableDtoCreatePlanRequest {
	return &NullableDtoCreatePlanRequest{value: val, isSet: true}
}

func (v NullableDtoCreatePlanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoCreatePlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


