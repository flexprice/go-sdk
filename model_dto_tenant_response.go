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

// checks if the DtoTenantResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoTenantResponse{}

// DtoTenantResponse struct for DtoTenantResponse
type DtoTenantResponse struct {
	BillingDetails *DtoTenantBillingDetails `json:"billing_details,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewDtoTenantResponse instantiates a new DtoTenantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoTenantResponse() *DtoTenantResponse {
	this := DtoTenantResponse{}
	return &this
}

// NewDtoTenantResponseWithDefaults instantiates a new DtoTenantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoTenantResponseWithDefaults() *DtoTenantResponse {
	this := DtoTenantResponse{}
	return &this
}

// GetBillingDetails returns the BillingDetails field value if set, zero value otherwise.
func (o *DtoTenantResponse) GetBillingDetails() DtoTenantBillingDetails {
	if o == nil || IsNil(o.BillingDetails) {
		var ret DtoTenantBillingDetails
		return ret
	}
	return *o.BillingDetails
}

// GetBillingDetailsOk returns a tuple with the BillingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTenantResponse) GetBillingDetailsOk() (*DtoTenantBillingDetails, bool) {
	if o == nil || IsNil(o.BillingDetails) {
		return nil, false
	}
	return o.BillingDetails, true
}

// HasBillingDetails returns a boolean if a field has been set.
func (o *DtoTenantResponse) HasBillingDetails() bool {
	if o != nil && !IsNil(o.BillingDetails) {
		return true
	}

	return false
}

// SetBillingDetails gets a reference to the given DtoTenantBillingDetails and assigns it to the BillingDetails field.
func (o *DtoTenantResponse) SetBillingDetails(v DtoTenantBillingDetails) {
	o.BillingDetails = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DtoTenantResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTenantResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DtoTenantResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DtoTenantResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DtoTenantResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTenantResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DtoTenantResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DtoTenantResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DtoTenantResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTenantResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DtoTenantResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DtoTenantResponse) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DtoTenantResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTenantResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DtoTenantResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DtoTenantResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DtoTenantResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTenantResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DtoTenantResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DtoTenantResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o DtoTenantResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoTenantResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingDetails) {
		toSerialize["billing_details"] = o.BillingDetails
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableDtoTenantResponse struct {
	value *DtoTenantResponse
	isSet bool
}

func (v NullableDtoTenantResponse) Get() *DtoTenantResponse {
	return v.value
}

func (v *NullableDtoTenantResponse) Set(val *DtoTenantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoTenantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoTenantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoTenantResponse(val *DtoTenantResponse) *NullableDtoTenantResponse {
	return &NullableDtoTenantResponse{value: val, isSet: true}
}

func (v NullableDtoTenantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoTenantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


