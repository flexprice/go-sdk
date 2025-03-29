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

// checks if the DtoListPricesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoListPricesResponse{}

// DtoListPricesResponse struct for DtoListPricesResponse
type DtoListPricesResponse struct {
	Items []DtoPriceResponse `json:"items,omitempty"`
	Pagination *TypesPaginationResponse `json:"pagination,omitempty"`
}

// NewDtoListPricesResponse instantiates a new DtoListPricesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoListPricesResponse() *DtoListPricesResponse {
	this := DtoListPricesResponse{}
	return &this
}

// NewDtoListPricesResponseWithDefaults instantiates a new DtoListPricesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoListPricesResponseWithDefaults() *DtoListPricesResponse {
	this := DtoListPricesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DtoListPricesResponse) GetItems() []DtoPriceResponse {
	if o == nil || IsNil(o.Items) {
		var ret []DtoPriceResponse
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoListPricesResponse) GetItemsOk() ([]DtoPriceResponse, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DtoListPricesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DtoPriceResponse and assigns it to the Items field.
func (o *DtoListPricesResponse) SetItems(v []DtoPriceResponse) {
	o.Items = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *DtoListPricesResponse) GetPagination() TypesPaginationResponse {
	if o == nil || IsNil(o.Pagination) {
		var ret TypesPaginationResponse
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoListPricesResponse) GetPaginationOk() (*TypesPaginationResponse, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *DtoListPricesResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given TypesPaginationResponse and assigns it to the Pagination field.
func (o *DtoListPricesResponse) SetPagination(v TypesPaginationResponse) {
	o.Pagination = &v
}

func (o DtoListPricesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoListPricesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableDtoListPricesResponse struct {
	value *DtoListPricesResponse
	isSet bool
}

func (v NullableDtoListPricesResponse) Get() *DtoListPricesResponse {
	return v.value
}

func (v *NullableDtoListPricesResponse) Set(val *DtoListPricesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoListPricesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoListPricesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoListPricesResponse(val *DtoListPricesResponse) *NullableDtoListPricesResponse {
	return &NullableDtoListPricesResponse{value: val, isSet: true}
}

func (v NullableDtoListPricesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoListPricesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


