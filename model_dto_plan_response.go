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

// checks if the DtoPlanResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoPlanResponse{}

// DtoPlanResponse struct for DtoPlanResponse
type DtoPlanResponse struct {
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Description *string `json:"description,omitempty"`
	Entitlements []DtoEntitlementResponse `json:"entitlements,omitempty"`
	EnvironmentId *string `json:"environment_id,omitempty"`
	Id *string `json:"id,omitempty"`
	LookupKey *string `json:"lookup_key,omitempty"`
	Name *string `json:"name,omitempty"`
	Prices []DtoPriceResponse `json:"prices,omitempty"`
	Status *TypesStatus `json:"status,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// NewDtoPlanResponse instantiates a new DtoPlanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoPlanResponse() *DtoPlanResponse {
	this := DtoPlanResponse{}
	return &this
}

// NewDtoPlanResponseWithDefaults instantiates a new DtoPlanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoPlanResponseWithDefaults() *DtoPlanResponse {
	this := DtoPlanResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DtoPlanResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DtoPlanResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DtoPlanResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetEntitlements() []DtoEntitlementResponse {
	if o == nil || IsNil(o.Entitlements) {
		var ret []DtoEntitlementResponse
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetEntitlementsOk() ([]DtoEntitlementResponse, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []DtoEntitlementResponse and assigns it to the Entitlements field.
func (o *DtoPlanResponse) SetEntitlements(v []DtoEntitlementResponse) {
	o.Entitlements = v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *DtoPlanResponse) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DtoPlanResponse) SetId(v string) {
	o.Id = &v
}

// GetLookupKey returns the LookupKey field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetLookupKey() string {
	if o == nil || IsNil(o.LookupKey) {
		var ret string
		return ret
	}
	return *o.LookupKey
}

// GetLookupKeyOk returns a tuple with the LookupKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetLookupKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LookupKey) {
		return nil, false
	}
	return o.LookupKey, true
}

// HasLookupKey returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasLookupKey() bool {
	if o != nil && !IsNil(o.LookupKey) {
		return true
	}

	return false
}

// SetLookupKey gets a reference to the given string and assigns it to the LookupKey field.
func (o *DtoPlanResponse) SetLookupKey(v string) {
	o.LookupKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DtoPlanResponse) SetName(v string) {
	o.Name = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetPrices() []DtoPriceResponse {
	if o == nil || IsNil(o.Prices) {
		var ret []DtoPriceResponse
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetPricesOk() ([]DtoPriceResponse, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []DtoPriceResponse and assigns it to the Prices field.
func (o *DtoPlanResponse) SetPrices(v []DtoPriceResponse) {
	o.Prices = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetStatus() TypesStatus {
	if o == nil || IsNil(o.Status) {
		var ret TypesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetStatusOk() (*TypesStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TypesStatus and assigns it to the Status field.
func (o *DtoPlanResponse) SetStatus(v TypesStatus) {
	o.Status = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *DtoPlanResponse) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DtoPlanResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *DtoPlanResponse) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPlanResponse) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DtoPlanResponse) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *DtoPlanResponse) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o DtoPlanResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoPlanResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LookupKey) {
		toSerialize["lookup_key"] = o.LookupKey
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	return toSerialize, nil
}

type NullableDtoPlanResponse struct {
	value *DtoPlanResponse
	isSet bool
}

func (v NullableDtoPlanResponse) Get() *DtoPlanResponse {
	return v.value
}

func (v *NullableDtoPlanResponse) Set(val *DtoPlanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoPlanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoPlanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoPlanResponse(val *DtoPlanResponse) *NullableDtoPlanResponse {
	return &NullableDtoPlanResponse{value: val, isSet: true}
}

func (v NullableDtoPlanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoPlanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


