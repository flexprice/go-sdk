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

// checks if the DtoGetEventsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoGetEventsRequest{}

// DtoGetEventsRequest struct for DtoGetEventsRequest
type DtoGetEventsRequest struct {
	// End time of the events to be fetched in ISO 8601 format Defaults to now if not provided
	EndTime *string `json:"end_time,omitempty"`
	// Event ID is the idempotency key for the event
	EventId *string `json:"event_id,omitempty"`
	// Event name / Unique identifier for the event in your system
	EventName *string `json:"event_name,omitempty"`
	// Customer ID in your system that was sent with the event
	ExternalCustomerId *string `json:"external_customer_id,omitempty"`
	// First key to iterate over the events
	IterFirstKey *string `json:"iter_first_key,omitempty"`
	// Last key to iterate over the events
	IterLastKey *string `json:"iter_last_key,omitempty"`
	// Offset to fetch the events and is set to 0 by default
	Offset *int32 `json:"offset,omitempty"`
	// Order by condition. Allowed values (case sensitive): asc, desc (default: desc)
	Order *string `json:"order,omitempty"`
	// Page size to fetch the events and is set to 50 by default
	PageSize *int32 `json:"page_size,omitempty"`
	// Property filters to filter the events by the keys in `properties` field of the event
	PropertyFilters *map[string][]string `json:"property_filters,omitempty"`
	// Sort by the field. Allowed values (case sensitive): timestamp, event_name (default: timestamp)
	Sort *string `json:"sort,omitempty"`
	// Source to filter the events by the source
	Source *string `json:"source,omitempty"`
	// Start time of the events to be fetched in ISO 8601 format Defaults to last 7 days from now if not provided
	StartTime *string `json:"start_time,omitempty"`
}

// NewDtoGetEventsRequest instantiates a new DtoGetEventsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoGetEventsRequest() *DtoGetEventsRequest {
	this := DtoGetEventsRequest{}
	return &this
}

// NewDtoGetEventsRequestWithDefaults instantiates a new DtoGetEventsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoGetEventsRequestWithDefaults() *DtoGetEventsRequest {
	this := DtoGetEventsRequest{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *DtoGetEventsRequest) SetEndTime(v string) {
	o.EndTime = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *DtoGetEventsRequest) SetEventId(v string) {
	o.EventId = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *DtoGetEventsRequest) SetEventName(v string) {
	o.EventName = &v
}

// GetExternalCustomerId returns the ExternalCustomerId field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetExternalCustomerId() string {
	if o == nil || IsNil(o.ExternalCustomerId) {
		var ret string
		return ret
	}
	return *o.ExternalCustomerId
}

// GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetExternalCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalCustomerId) {
		return nil, false
	}
	return o.ExternalCustomerId, true
}

// HasExternalCustomerId returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasExternalCustomerId() bool {
	if o != nil && !IsNil(o.ExternalCustomerId) {
		return true
	}

	return false
}

// SetExternalCustomerId gets a reference to the given string and assigns it to the ExternalCustomerId field.
func (o *DtoGetEventsRequest) SetExternalCustomerId(v string) {
	o.ExternalCustomerId = &v
}

// GetIterFirstKey returns the IterFirstKey field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetIterFirstKey() string {
	if o == nil || IsNil(o.IterFirstKey) {
		var ret string
		return ret
	}
	return *o.IterFirstKey
}

// GetIterFirstKeyOk returns a tuple with the IterFirstKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetIterFirstKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IterFirstKey) {
		return nil, false
	}
	return o.IterFirstKey, true
}

// HasIterFirstKey returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasIterFirstKey() bool {
	if o != nil && !IsNil(o.IterFirstKey) {
		return true
	}

	return false
}

// SetIterFirstKey gets a reference to the given string and assigns it to the IterFirstKey field.
func (o *DtoGetEventsRequest) SetIterFirstKey(v string) {
	o.IterFirstKey = &v
}

// GetIterLastKey returns the IterLastKey field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetIterLastKey() string {
	if o == nil || IsNil(o.IterLastKey) {
		var ret string
		return ret
	}
	return *o.IterLastKey
}

// GetIterLastKeyOk returns a tuple with the IterLastKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetIterLastKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IterLastKey) {
		return nil, false
	}
	return o.IterLastKey, true
}

// HasIterLastKey returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasIterLastKey() bool {
	if o != nil && !IsNil(o.IterLastKey) {
		return true
	}

	return false
}

// SetIterLastKey gets a reference to the given string and assigns it to the IterLastKey field.
func (o *DtoGetEventsRequest) SetIterLastKey(v string) {
	o.IterLastKey = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *DtoGetEventsRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *DtoGetEventsRequest) SetOrder(v string) {
	o.Order = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *DtoGetEventsRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPropertyFilters returns the PropertyFilters field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetPropertyFilters() map[string][]string {
	if o == nil || IsNil(o.PropertyFilters) {
		var ret map[string][]string
		return ret
	}
	return *o.PropertyFilters
}

// GetPropertyFiltersOk returns a tuple with the PropertyFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetPropertyFiltersOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.PropertyFilters) {
		return nil, false
	}
	return o.PropertyFilters, true
}

// HasPropertyFilters returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasPropertyFilters() bool {
	if o != nil && !IsNil(o.PropertyFilters) {
		return true
	}

	return false
}

// SetPropertyFilters gets a reference to the given map[string][]string and assigns it to the PropertyFilters field.
func (o *DtoGetEventsRequest) SetPropertyFilters(v map[string][]string) {
	o.PropertyFilters = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetSort() string {
	if o == nil || IsNil(o.Sort) {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetSortOk() (*string, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *DtoGetEventsRequest) SetSort(v string) {
	o.Sort = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *DtoGetEventsRequest) SetSource(v string) {
	o.Source = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DtoGetEventsRequest) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoGetEventsRequest) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DtoGetEventsRequest) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *DtoGetEventsRequest) SetStartTime(v string) {
	o.StartTime = &v
}

func (o DtoGetEventsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoGetEventsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.EventId) {
		toSerialize["event_id"] = o.EventId
	}
	if !IsNil(o.EventName) {
		toSerialize["event_name"] = o.EventName
	}
	if !IsNil(o.ExternalCustomerId) {
		toSerialize["external_customer_id"] = o.ExternalCustomerId
	}
	if !IsNil(o.IterFirstKey) {
		toSerialize["iter_first_key"] = o.IterFirstKey
	}
	if !IsNil(o.IterLastKey) {
		toSerialize["iter_last_key"] = o.IterLastKey
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.PropertyFilters) {
		toSerialize["property_filters"] = o.PropertyFilters
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	return toSerialize, nil
}

type NullableDtoGetEventsRequest struct {
	value *DtoGetEventsRequest
	isSet bool
}

func (v NullableDtoGetEventsRequest) Get() *DtoGetEventsRequest {
	return v.value
}

func (v *NullableDtoGetEventsRequest) Set(val *DtoGetEventsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoGetEventsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoGetEventsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoGetEventsRequest(val *DtoGetEventsRequest) *NullableDtoGetEventsRequest {
	return &NullableDtoGetEventsRequest{value: val, isSet: true}
}

func (v NullableDtoGetEventsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoGetEventsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


