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

// checks if the DtoUpdateTaskStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoUpdateTaskStatusRequest{}

// DtoUpdateTaskStatusRequest struct for DtoUpdateTaskStatusRequest
type DtoUpdateTaskStatusRequest struct {
	TaskStatus TypesTaskStatus `json:"task_status"`
}

type _DtoUpdateTaskStatusRequest DtoUpdateTaskStatusRequest

// NewDtoUpdateTaskStatusRequest instantiates a new DtoUpdateTaskStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoUpdateTaskStatusRequest(taskStatus TypesTaskStatus) *DtoUpdateTaskStatusRequest {
	this := DtoUpdateTaskStatusRequest{}
	this.TaskStatus = taskStatus
	return &this
}

// NewDtoUpdateTaskStatusRequestWithDefaults instantiates a new DtoUpdateTaskStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoUpdateTaskStatusRequestWithDefaults() *DtoUpdateTaskStatusRequest {
	this := DtoUpdateTaskStatusRequest{}
	return &this
}

// GetTaskStatus returns the TaskStatus field value
func (o *DtoUpdateTaskStatusRequest) GetTaskStatus() TypesTaskStatus {
	if o == nil {
		var ret TypesTaskStatus
		return ret
	}

	return o.TaskStatus
}

// GetTaskStatusOk returns a tuple with the TaskStatus field value
// and a boolean to check if the value has been set.
func (o *DtoUpdateTaskStatusRequest) GetTaskStatusOk() (*TypesTaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskStatus, true
}

// SetTaskStatus sets field value
func (o *DtoUpdateTaskStatusRequest) SetTaskStatus(v TypesTaskStatus) {
	o.TaskStatus = v
}

func (o DtoUpdateTaskStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoUpdateTaskStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["task_status"] = o.TaskStatus
	return toSerialize, nil
}

func (o *DtoUpdateTaskStatusRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"task_status",
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

	varDtoUpdateTaskStatusRequest := _DtoUpdateTaskStatusRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDtoUpdateTaskStatusRequest)

	if err != nil {
		return err
	}

	*o = DtoUpdateTaskStatusRequest(varDtoUpdateTaskStatusRequest)

	return err
}

type NullableDtoUpdateTaskStatusRequest struct {
	value *DtoUpdateTaskStatusRequest
	isSet bool
}

func (v NullableDtoUpdateTaskStatusRequest) Get() *DtoUpdateTaskStatusRequest {
	return v.value
}

func (v *NullableDtoUpdateTaskStatusRequest) Set(val *DtoUpdateTaskStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoUpdateTaskStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoUpdateTaskStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoUpdateTaskStatusRequest(val *DtoUpdateTaskStatusRequest) *NullableDtoUpdateTaskStatusRequest {
	return &NullableDtoUpdateTaskStatusRequest{value: val, isSet: true}
}

func (v NullableDtoUpdateTaskStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoUpdateTaskStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


