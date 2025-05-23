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

// checks if the SubscriptionSubscriptionLineItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionSubscriptionLineItem{}

// SubscriptionSubscriptionLineItem struct for SubscriptionSubscriptionLineItem
type SubscriptionSubscriptionLineItem struct {
	BillingPeriod *TypesBillingPeriod `json:"billing_period,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
	EnvironmentId *string `json:"environment_id,omitempty"`
	Id *string `json:"id,omitempty"`
	InvoiceCadence *TypesInvoiceCadence `json:"invoice_cadence,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	MeterDisplayName *string `json:"meter_display_name,omitempty"`
	MeterId *string `json:"meter_id,omitempty"`
	PlanDisplayName *string `json:"plan_display_name,omitempty"`
	PlanId *string `json:"plan_id,omitempty"`
	PriceId *string `json:"price_id,omitempty"`
	PriceType *TypesPriceType `json:"price_type,omitempty"`
	Quantity *float32 `json:"quantity,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	Status *TypesStatus `json:"status,omitempty"`
	SubscriptionId *string `json:"subscription_id,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	TrialPeriod *int32 `json:"trial_period,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// NewSubscriptionSubscriptionLineItem instantiates a new SubscriptionSubscriptionLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionSubscriptionLineItem() *SubscriptionSubscriptionLineItem {
	this := SubscriptionSubscriptionLineItem{}
	return &this
}

// NewSubscriptionSubscriptionLineItemWithDefaults instantiates a new SubscriptionSubscriptionLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionSubscriptionLineItemWithDefaults() *SubscriptionSubscriptionLineItem {
	this := SubscriptionSubscriptionLineItem{}
	return &this
}

// GetBillingPeriod returns the BillingPeriod field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetBillingPeriod() TypesBillingPeriod {
	if o == nil || IsNil(o.BillingPeriod) {
		var ret TypesBillingPeriod
		return ret
	}
	return *o.BillingPeriod
}

// GetBillingPeriodOk returns a tuple with the BillingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetBillingPeriodOk() (*TypesBillingPeriod, bool) {
	if o == nil || IsNil(o.BillingPeriod) {
		return nil, false
	}
	return o.BillingPeriod, true
}

// HasBillingPeriod returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasBillingPeriod() bool {
	if o != nil && !IsNil(o.BillingPeriod) {
		return true
	}

	return false
}

// SetBillingPeriod gets a reference to the given TypesBillingPeriod and assigns it to the BillingPeriod field.
func (o *SubscriptionSubscriptionLineItem) SetBillingPeriod(v TypesBillingPeriod) {
	o.BillingPeriod = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SubscriptionSubscriptionLineItem) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SubscriptionSubscriptionLineItem) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *SubscriptionSubscriptionLineItem) SetCurrency(v string) {
	o.Currency = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *SubscriptionSubscriptionLineItem) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SubscriptionSubscriptionLineItem) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SubscriptionSubscriptionLineItem) SetEndDate(v string) {
	o.EndDate = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *SubscriptionSubscriptionLineItem) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionSubscriptionLineItem) SetId(v string) {
	o.Id = &v
}

// GetInvoiceCadence returns the InvoiceCadence field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetInvoiceCadence() TypesInvoiceCadence {
	if o == nil || IsNil(o.InvoiceCadence) {
		var ret TypesInvoiceCadence
		return ret
	}
	return *o.InvoiceCadence
}

// GetInvoiceCadenceOk returns a tuple with the InvoiceCadence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetInvoiceCadenceOk() (*TypesInvoiceCadence, bool) {
	if o == nil || IsNil(o.InvoiceCadence) {
		return nil, false
	}
	return o.InvoiceCadence, true
}

// HasInvoiceCadence returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasInvoiceCadence() bool {
	if o != nil && !IsNil(o.InvoiceCadence) {
		return true
	}

	return false
}

// SetInvoiceCadence gets a reference to the given TypesInvoiceCadence and assigns it to the InvoiceCadence field.
func (o *SubscriptionSubscriptionLineItem) SetInvoiceCadence(v TypesInvoiceCadence) {
	o.InvoiceCadence = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *SubscriptionSubscriptionLineItem) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMeterDisplayName returns the MeterDisplayName field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetMeterDisplayName() string {
	if o == nil || IsNil(o.MeterDisplayName) {
		var ret string
		return ret
	}
	return *o.MeterDisplayName
}

// GetMeterDisplayNameOk returns a tuple with the MeterDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetMeterDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.MeterDisplayName) {
		return nil, false
	}
	return o.MeterDisplayName, true
}

// HasMeterDisplayName returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasMeterDisplayName() bool {
	if o != nil && !IsNil(o.MeterDisplayName) {
		return true
	}

	return false
}

// SetMeterDisplayName gets a reference to the given string and assigns it to the MeterDisplayName field.
func (o *SubscriptionSubscriptionLineItem) SetMeterDisplayName(v string) {
	o.MeterDisplayName = &v
}

// GetMeterId returns the MeterId field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetMeterId() string {
	if o == nil || IsNil(o.MeterId) {
		var ret string
		return ret
	}
	return *o.MeterId
}

// GetMeterIdOk returns a tuple with the MeterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetMeterIdOk() (*string, bool) {
	if o == nil || IsNil(o.MeterId) {
		return nil, false
	}
	return o.MeterId, true
}

// HasMeterId returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasMeterId() bool {
	if o != nil && !IsNil(o.MeterId) {
		return true
	}

	return false
}

// SetMeterId gets a reference to the given string and assigns it to the MeterId field.
func (o *SubscriptionSubscriptionLineItem) SetMeterId(v string) {
	o.MeterId = &v
}

// GetPlanDisplayName returns the PlanDisplayName field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetPlanDisplayName() string {
	if o == nil || IsNil(o.PlanDisplayName) {
		var ret string
		return ret
	}
	return *o.PlanDisplayName
}

// GetPlanDisplayNameOk returns a tuple with the PlanDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetPlanDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlanDisplayName) {
		return nil, false
	}
	return o.PlanDisplayName, true
}

// HasPlanDisplayName returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasPlanDisplayName() bool {
	if o != nil && !IsNil(o.PlanDisplayName) {
		return true
	}

	return false
}

// SetPlanDisplayName gets a reference to the given string and assigns it to the PlanDisplayName field.
func (o *SubscriptionSubscriptionLineItem) SetPlanDisplayName(v string) {
	o.PlanDisplayName = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SubscriptionSubscriptionLineItem) SetPlanId(v string) {
	o.PlanId = &v
}

// GetPriceId returns the PriceId field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetPriceId() string {
	if o == nil || IsNil(o.PriceId) {
		var ret string
		return ret
	}
	return *o.PriceId
}

// GetPriceIdOk returns a tuple with the PriceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetPriceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PriceId) {
		return nil, false
	}
	return o.PriceId, true
}

// HasPriceId returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasPriceId() bool {
	if o != nil && !IsNil(o.PriceId) {
		return true
	}

	return false
}

// SetPriceId gets a reference to the given string and assigns it to the PriceId field.
func (o *SubscriptionSubscriptionLineItem) SetPriceId(v string) {
	o.PriceId = &v
}

// GetPriceType returns the PriceType field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetPriceType() TypesPriceType {
	if o == nil || IsNil(o.PriceType) {
		var ret TypesPriceType
		return ret
	}
	return *o.PriceType
}

// GetPriceTypeOk returns a tuple with the PriceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetPriceTypeOk() (*TypesPriceType, bool) {
	if o == nil || IsNil(o.PriceType) {
		return nil, false
	}
	return o.PriceType, true
}

// HasPriceType returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasPriceType() bool {
	if o != nil && !IsNil(o.PriceType) {
		return true
	}

	return false
}

// SetPriceType gets a reference to the given TypesPriceType and assigns it to the PriceType field.
func (o *SubscriptionSubscriptionLineItem) SetPriceType(v TypesPriceType) {
	o.PriceType = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *SubscriptionSubscriptionLineItem) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SubscriptionSubscriptionLineItem) SetStartDate(v string) {
	o.StartDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetStatus() TypesStatus {
	if o == nil || IsNil(o.Status) {
		var ret TypesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetStatusOk() (*TypesStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TypesStatus and assigns it to the Status field.
func (o *SubscriptionSubscriptionLineItem) SetStatus(v TypesStatus) {
	o.Status = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *SubscriptionSubscriptionLineItem) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *SubscriptionSubscriptionLineItem) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTrialPeriod returns the TrialPeriod field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetTrialPeriod() int32 {
	if o == nil || IsNil(o.TrialPeriod) {
		var ret int32
		return ret
	}
	return *o.TrialPeriod
}

// GetTrialPeriodOk returns a tuple with the TrialPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetTrialPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialPeriod) {
		return nil, false
	}
	return o.TrialPeriod, true
}

// HasTrialPeriod returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasTrialPeriod() bool {
	if o != nil && !IsNil(o.TrialPeriod) {
		return true
	}

	return false
}

// SetTrialPeriod gets a reference to the given int32 and assigns it to the TrialPeriod field.
func (o *SubscriptionSubscriptionLineItem) SetTrialPeriod(v int32) {
	o.TrialPeriod = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *SubscriptionSubscriptionLineItem) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *SubscriptionSubscriptionLineItem) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubscriptionLineItem) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SubscriptionSubscriptionLineItem) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *SubscriptionSubscriptionLineItem) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o SubscriptionSubscriptionLineItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionSubscriptionLineItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingPeriod) {
		toSerialize["billing_period"] = o.BillingPeriod
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customer_id"] = o.CustomerId
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InvoiceCadence) {
		toSerialize["invoice_cadence"] = o.InvoiceCadence
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MeterDisplayName) {
		toSerialize["meter_display_name"] = o.MeterDisplayName
	}
	if !IsNil(o.MeterId) {
		toSerialize["meter_id"] = o.MeterId
	}
	if !IsNil(o.PlanDisplayName) {
		toSerialize["plan_display_name"] = o.PlanDisplayName
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.PriceId) {
		toSerialize["price_id"] = o.PriceId
	}
	if !IsNil(o.PriceType) {
		toSerialize["price_type"] = o.PriceType
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !IsNil(o.TrialPeriod) {
		toSerialize["trial_period"] = o.TrialPeriod
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	return toSerialize, nil
}

type NullableSubscriptionSubscriptionLineItem struct {
	value *SubscriptionSubscriptionLineItem
	isSet bool
}

func (v NullableSubscriptionSubscriptionLineItem) Get() *SubscriptionSubscriptionLineItem {
	return v.value
}

func (v *NullableSubscriptionSubscriptionLineItem) Set(val *SubscriptionSubscriptionLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionSubscriptionLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionSubscriptionLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionSubscriptionLineItem(val *SubscriptionSubscriptionLineItem) *NullableSubscriptionSubscriptionLineItem {
	return &NullableSubscriptionSubscriptionLineItem{value: val, isSet: true}
}

func (v NullableSubscriptionSubscriptionLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionSubscriptionLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


