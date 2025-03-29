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

// checks if the DtoListSubscriptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoListSubscriptionsResponse{}

// DtoListSubscriptionsResponse struct for DtoListSubscriptionsResponse
type DtoListSubscriptionsResponse struct {
	Items []DtoSubscriptionResponse `json:"items,omitempty"`
	Pagination *TypesPaginationResponse `json:"pagination,omitempty"`
}

// NewDtoListSubscriptionsResponse instantiates a new DtoListSubscriptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoListSubscriptionsResponse() *DtoListSubscriptionsResponse {
	this := DtoListSubscriptionsResponse{}
	return &this
}

// NewDtoListSubscriptionsResponseWithDefaults instantiates a new DtoListSubscriptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoListSubscriptionsResponseWithDefaults() *DtoListSubscriptionsResponse {
	this := DtoListSubscriptionsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DtoListSubscriptionsResponse) GetItems() []DtoSubscriptionResponse {
	if o == nil || IsNil(o.Items) {
		var ret []DtoSubscriptionResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoListSubscriptionsResponse) GetItemsOk() ([]DtoSubscriptionResponse, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DtoListSubscriptionsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DtoSubscriptionResponse and assigns it to the Items field.
func (o *DtoListSubscriptionsResponse) SetItems(v []DtoSubscriptionResponse) {
	o.Items = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *DtoListSubscriptionsResponse) GetPagination() TypesPaginationResponse {
	if o == nil || IsNil(o.Pagination) {
		var ret TypesPaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoListSubscriptionsResponse) GetPaginationOk() (*TypesPaginationResponse, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *DtoListSubscriptionsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given TypesPaginationResponse and assigns it to the Pagination field.
func (o *DtoListSubscriptionsResponse) SetPagination(v TypesPaginationResponse) {
	o.Pagination = &v
}

func (o DtoListSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoListSubscriptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableDtoListSubscriptionsResponse struct {
	value *DtoListSubscriptionsResponse
	isSet bool
}

func (v NullableDtoListSubscriptionsResponse) Get() *DtoListSubscriptionsResponse {
	return v.value
}

func (v *NullableDtoListSubscriptionsResponse) Set(val *DtoListSubscriptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoListSubscriptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoListSubscriptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoListSubscriptionsResponse(val *DtoListSubscriptionsResponse) *NullableDtoListSubscriptionsResponse {
	return &NullableDtoListSubscriptionsResponse{value: val, isSet: true}
}

func (v NullableDtoListSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoListSubscriptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


