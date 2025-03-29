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

// checks if the DtoEntitlementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoEntitlementResponse{}

// DtoEntitlementResponse struct for DtoEntitlementResponse
type DtoEntitlementResponse struct {
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	EnvironmentId *string `json:"environment_id,omitempty"`
	Feature *DtoFeatureResponse `json:"feature,omitempty"`
	FeatureId *string `json:"feature_id,omitempty"`
	FeatureType *TypesFeatureType `json:"feature_type,omitempty"`
	Id *string `json:"id,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	IsSoftLimit *bool `json:"is_soft_limit,omitempty"`
	Plan *DtoPlanResponse `json:"plan,omitempty"`
	PlanId *string `json:"plan_id,omitempty"`
	StaticValue *string `json:"static_value,omitempty"`
	Status *TypesStatus `json:"status,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	UsageLimit *int32 `json:"usage_limit,omitempty"`
	UsageResetPeriod *TypesBillingPeriod `json:"usage_reset_period,omitempty"`
}

// NewDtoEntitlementResponse instantiates a new DtoEntitlementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoEntitlementResponse() *DtoEntitlementResponse {
	this := DtoEntitlementResponse{}
	return &this
}

// NewDtoEntitlementResponseWithDefaults instantiates a new DtoEntitlementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoEntitlementResponseWithDefaults() *DtoEntitlementResponse {
	this := DtoEntitlementResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DtoEntitlementResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DtoEntitlementResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *DtoEntitlementResponse) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetFeature() DtoFeatureResponse {
	if o == nil || IsNil(o.Feature) {
		var ret DtoFeatureResponse
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetFeatureOk() (*DtoFeatureResponse, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given DtoFeatureResponse and assigns it to the Feature field.
func (o *DtoEntitlementResponse) SetFeature(v DtoFeatureResponse) {
	o.Feature = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId) {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetFeatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *DtoEntitlementResponse) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetFeatureType() TypesFeatureType {
	if o == nil || IsNil(o.FeatureType) {
		var ret TypesFeatureType
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetFeatureTypeOk() (*TypesFeatureType, bool) {
	if o == nil || IsNil(o.FeatureType) {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasFeatureType() bool {
	if o != nil && !IsNil(o.FeatureType) {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given TypesFeatureType and assigns it to the FeatureType field.
func (o *DtoEntitlementResponse) SetFeatureType(v TypesFeatureType) {
	o.FeatureType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DtoEntitlementResponse) SetId(v string) {
	o.Id = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *DtoEntitlementResponse) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsSoftLimit returns the IsSoftLimit field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetIsSoftLimit() bool {
	if o == nil || IsNil(o.IsSoftLimit) {
		var ret bool
		return ret
	}
	return *o.IsSoftLimit
}

// GetIsSoftLimitOk returns a tuple with the IsSoftLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetIsSoftLimitOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSoftLimit) {
		return nil, false
	}
	return o.IsSoftLimit, true
}

// HasIsSoftLimit returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasIsSoftLimit() bool {
	if o != nil && !IsNil(o.IsSoftLimit) {
		return true
	}

	return false
}

// SetIsSoftLimit gets a reference to the given bool and assigns it to the IsSoftLimit field.
func (o *DtoEntitlementResponse) SetIsSoftLimit(v bool) {
	o.IsSoftLimit = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetPlan() DtoPlanResponse {
	if o == nil || IsNil(o.Plan) {
		var ret DtoPlanResponse
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetPlanOk() (*DtoPlanResponse, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given DtoPlanResponse and assigns it to the Plan field.
func (o *DtoEntitlementResponse) SetPlan(v DtoPlanResponse) {
	o.Plan = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *DtoEntitlementResponse) SetPlanId(v string) {
	o.PlanId = &v
}

// GetStaticValue returns the StaticValue field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetStaticValue() string {
	if o == nil || IsNil(o.StaticValue) {
		var ret string
		return ret
	}
	return *o.StaticValue
}

// GetStaticValueOk returns a tuple with the StaticValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetStaticValueOk() (*string, bool) {
	if o == nil || IsNil(o.StaticValue) {
		return nil, false
	}
	return o.StaticValue, true
}

// HasStaticValue returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasStaticValue() bool {
	if o != nil && !IsNil(o.StaticValue) {
		return true
	}

	return false
}

// SetStaticValue gets a reference to the given string and assigns it to the StaticValue field.
func (o *DtoEntitlementResponse) SetStaticValue(v string) {
	o.StaticValue = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetStatus() TypesStatus {
	if o == nil || IsNil(o.Status) {
		var ret TypesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetStatusOk() (*TypesStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TypesStatus and assigns it to the Status field.
func (o *DtoEntitlementResponse) SetStatus(v TypesStatus) {
	o.Status = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *DtoEntitlementResponse) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DtoEntitlementResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *DtoEntitlementResponse) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetUsageLimit() int32 {
	if o == nil || IsNil(o.UsageLimit) {
		var ret int32
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetUsageLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.UsageLimit) {
		return nil, false
	}
	return o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasUsageLimit() bool {
	if o != nil && !IsNil(o.UsageLimit) {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.
func (o *DtoEntitlementResponse) SetUsageLimit(v int32) {
	o.UsageLimit = &v
}

// GetUsageResetPeriod returns the UsageResetPeriod field value if set, zero value otherwise.
func (o *DtoEntitlementResponse) GetUsageResetPeriod() TypesBillingPeriod {
	if o == nil || IsNil(o.UsageResetPeriod) {
		var ret TypesBillingPeriod
		return ret
	}
	return *o.UsageResetPeriod
}

// GetUsageResetPeriodOk returns a tuple with the UsageResetPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoEntitlementResponse) GetUsageResetPeriodOk() (*TypesBillingPeriod, bool) {
	if o == nil || IsNil(o.UsageResetPeriod) {
		return nil, false
	}
	return o.UsageResetPeriod, true
}

// HasUsageResetPeriod returns a boolean if a field has been set.
func (o *DtoEntitlementResponse) HasUsageResetPeriod() bool {
	if o != nil && !IsNil(o.UsageResetPeriod) {
		return true
	}

	return false
}

// SetUsageResetPeriod gets a reference to the given TypesBillingPeriod and assigns it to the UsageResetPeriod field.
func (o *DtoEntitlementResponse) SetUsageResetPeriod(v TypesBillingPeriod) {
	o.UsageResetPeriod = &v
}

func (o DtoEntitlementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoEntitlementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}
	if !IsNil(o.FeatureId) {
		toSerialize["feature_id"] = o.FeatureId
	}
	if !IsNil(o.FeatureType) {
		toSerialize["feature_type"] = o.FeatureType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if !IsNil(o.IsSoftLimit) {
		toSerialize["is_soft_limit"] = o.IsSoftLimit
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.StaticValue) {
		toSerialize["static_value"] = o.StaticValue
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
	if !IsNil(o.UsageLimit) {
		toSerialize["usage_limit"] = o.UsageLimit
	}
	if !IsNil(o.UsageResetPeriod) {
		toSerialize["usage_reset_period"] = o.UsageResetPeriod
	}
	return toSerialize, nil
}

type NullableDtoEntitlementResponse struct {
	value *DtoEntitlementResponse
	isSet bool
}

func (v NullableDtoEntitlementResponse) Get() *DtoEntitlementResponse {
	return v.value
}

func (v *NullableDtoEntitlementResponse) Set(val *DtoEntitlementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoEntitlementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoEntitlementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoEntitlementResponse(val *DtoEntitlementResponse) *NullableDtoEntitlementResponse {
	return &NullableDtoEntitlementResponse{value: val, isSet: true}
}

func (v NullableDtoEntitlementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoEntitlementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


