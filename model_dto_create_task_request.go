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

// checks if the DtoCreateTaskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoCreateTaskRequest{}

// DtoCreateTaskRequest struct for DtoCreateTaskRequest
type DtoCreateTaskRequest struct {
	EntityType TypesEntityType `json:"entity_type"`
	FileName *string `json:"file_name,omitempty"`
	FileType TypesFileType `json:"file_type"`
	FileUrl string `json:"file_url"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	TaskType TypesTaskType `json:"task_type"`
}

type _DtoCreateTaskRequest DtoCreateTaskRequest

// NewDtoCreateTaskRequest instantiates a new DtoCreateTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoCreateTaskRequest(entityType TypesEntityType, fileType TypesFileType, fileUrl string, taskType TypesTaskType) *DtoCreateTaskRequest {
	this := DtoCreateTaskRequest{}
	this.EntityType = entityType
	this.FileType = fileType
	this.FileUrl = fileUrl
	this.TaskType = taskType
	return &this
}

// NewDtoCreateTaskRequestWithDefaults instantiates a new DtoCreateTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoCreateTaskRequestWithDefaults() *DtoCreateTaskRequest {
	this := DtoCreateTaskRequest{}
	return &this
}

// GetEntityType returns the EntityType field value
func (o *DtoCreateTaskRequest) GetEntityType() TypesEntityType {
	if o == nil {
		var ret TypesEntityType
		return ret
	}

	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *DtoCreateTaskRequest) GetEntityTypeOk() (*TypesEntityType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value
func (o *DtoCreateTaskRequest) SetEntityType(v TypesEntityType) {
	o.EntityType = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *DtoCreateTaskRequest) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreateTaskRequest) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *DtoCreateTaskRequest) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *DtoCreateTaskRequest) SetFileName(v string) {
	o.FileName = &v
}

// GetFileType returns the FileType field value
func (o *DtoCreateTaskRequest) GetFileType() TypesFileType {
	if o == nil {
		var ret TypesFileType
		return ret
	}

	return o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value
// and a boolean to check if the value has been set.
func (o *DtoCreateTaskRequest) GetFileTypeOk() (*TypesFileType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileType, true
}

// SetFileType sets field value
func (o *DtoCreateTaskRequest) SetFileType(v TypesFileType) {
	o.FileType = v
}

// GetFileUrl returns the FileUrl field value
func (o *DtoCreateTaskRequest) GetFileUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileUrl
}

// GetFileUrlOk returns a tuple with the FileUrl field value
// and a boolean to check if the value has been set.
func (o *DtoCreateTaskRequest) GetFileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileUrl, true
}

// SetFileUrl sets field value
func (o *DtoCreateTaskRequest) SetFileUrl(v string) {
	o.FileUrl = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DtoCreateTaskRequest) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoCreateTaskRequest) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DtoCreateTaskRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *DtoCreateTaskRequest) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetTaskType returns the TaskType field value
func (o *DtoCreateTaskRequest) GetTaskType() TypesTaskType {
	if o == nil {
		var ret TypesTaskType
		return ret
	}

	return o.TaskType
}

// GetTaskTypeOk returns a tuple with the TaskType field value
// and a boolean to check if the value has been set.
func (o *DtoCreateTaskRequest) GetTaskTypeOk() (*TypesTaskType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskType, true
}

// SetTaskType sets field value
func (o *DtoCreateTaskRequest) SetTaskType(v TypesTaskType) {
	o.TaskType = v
}

func (o DtoCreateTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoCreateTaskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entity_type"] = o.EntityType
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	toSerialize["file_type"] = o.FileType
	toSerialize["file_url"] = o.FileUrl
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["task_type"] = o.TaskType
	return toSerialize, nil
}

func (o *DtoCreateTaskRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entity_type",
		"file_type",
		"file_url",
		"task_type",
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

	varDtoCreateTaskRequest := _DtoCreateTaskRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDtoCreateTaskRequest)

	if err != nil {
		return err
	}

	*o = DtoCreateTaskRequest(varDtoCreateTaskRequest)

	return err
}

type NullableDtoCreateTaskRequest struct {
	value *DtoCreateTaskRequest
	isSet bool
}

func (v NullableDtoCreateTaskRequest) Get() *DtoCreateTaskRequest {
	return v.value
}

func (v *NullableDtoCreateTaskRequest) Set(val *DtoCreateTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoCreateTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoCreateTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoCreateTaskRequest(val *DtoCreateTaskRequest) *NullableDtoCreateTaskRequest {
	return &NullableDtoCreateTaskRequest{value: val, isSet: true}
}

func (v NullableDtoCreateTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoCreateTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


