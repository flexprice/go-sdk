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

// checks if the DtoPriceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoPriceResponse{}

// DtoPriceResponse struct for DtoPriceResponse
type DtoPriceResponse struct {
	// Amount stored in main currency units (e.g., dollars, not cents) For USD: 12.50 means $12.50
	Amount *float32 `json:"amount,omitempty"`
	BillingCadence *TypesBillingCadence `json:"billing_cadence,omitempty"`
	BillingModel *TypesBillingModel `json:"billing_model,omitempty"`
	BillingPeriod *TypesBillingPeriod `json:"billing_period,omitempty"`
	// BillingPeriodCount is the count of the billing period ex 1, 3, 6, 12
	BillingPeriodCount *int32 `json:"billing_period_count,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	// Currency 3 digit ISO currency code in lowercase ex usd, eur, gbp
	Currency *string `json:"currency,omitempty"`
	// Description of the price
	Description *string `json:"description,omitempty"`
	// DisplayAmount is the formatted amount with currency symbol For USD: $12.50
	DisplayAmount *string `json:"display_amount,omitempty"`
	// EnvironmentID is the environment identifier for the price
	EnvironmentId *string `json:"environment_id,omitempty"`
	// ID uuid identifier for the price
	Id *string `json:"id,omitempty"`
	InvoiceCadence *TypesInvoiceCadence `json:"invoice_cadence,omitempty"`
	// LookupKey used for looking up the price in the database
	LookupKey *string `json:"lookup_key,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Meter *DtoMeterResponse `json:"meter,omitempty"`
	// MeterID is the id of the meter for usage based pricing
	MeterId *string `json:"meter_id,omitempty"`
	// PlanID is the id of the plan for plan based pricing
	PlanId *string `json:"plan_id,omitempty"`
	Status *TypesStatus `json:"status,omitempty"`
	TenantId *string `json:"tenant_id,omitempty"`
	TierMode *TypesBillingTier `json:"tier_mode,omitempty"`
	Tiers []PricePriceTier `json:"tiers,omitempty"`
	TransformQuantity *PriceJSONBTransformQuantity `json:"transform_quantity,omitempty"`
	// TrialPeriod is the number of days for the trial period Note: This is only applicable for recurring prices (BILLING_CADENCE_RECURRING)
	TrialPeriod *int32 `json:"trial_period,omitempty"`
	Type *TypesPriceType `json:"type,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// NewDtoPriceResponse instantiates a new DtoPriceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoPriceResponse() *DtoPriceResponse {
	this := DtoPriceResponse{}
	return &this
}

// NewDtoPriceResponseWithDefaults instantiates a new DtoPriceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoPriceResponseWithDefaults() *DtoPriceResponse {
	this := DtoPriceResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *DtoPriceResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetBillingCadence returns the BillingCadence field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetBillingCadence() TypesBillingCadence {
	if o == nil || IsNil(o.BillingCadence) {
		var ret TypesBillingCadence
		return ret
	}
	return *o.BillingCadence
}

// GetBillingCadenceOk returns a tuple with the BillingCadence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetBillingCadenceOk() (*TypesBillingCadence, bool) {
	if o == nil || IsNil(o.BillingCadence) {
		return nil, false
	}
	return o.BillingCadence, true
}

// HasBillingCadence returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasBillingCadence() bool {
	if o != nil && !IsNil(o.BillingCadence) {
		return true
	}

	return false
}

// SetBillingCadence gets a reference to the given TypesBillingCadence and assigns it to the BillingCadence field.
func (o *DtoPriceResponse) SetBillingCadence(v TypesBillingCadence) {
	o.BillingCadence = &v
}

// GetBillingModel returns the BillingModel field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetBillingModel() TypesBillingModel {
	if o == nil || IsNil(o.BillingModel) {
		var ret TypesBillingModel
		return ret
	}
	return *o.BillingModel
}

// GetBillingModelOk returns a tuple with the BillingModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetBillingModelOk() (*TypesBillingModel, bool) {
	if o == nil || IsNil(o.BillingModel) {
		return nil, false
	}
	return o.BillingModel, true
}

// HasBillingModel returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasBillingModel() bool {
	if o != nil && !IsNil(o.BillingModel) {
		return true
	}

	return false
}

// SetBillingModel gets a reference to the given TypesBillingModel and assigns it to the BillingModel field.
func (o *DtoPriceResponse) SetBillingModel(v TypesBillingModel) {
	o.BillingModel = &v
}

// GetBillingPeriod returns the BillingPeriod field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetBillingPeriod() TypesBillingPeriod {
	if o == nil || IsNil(o.BillingPeriod) {
		var ret TypesBillingPeriod
		return ret
	}
	return *o.BillingPeriod
}

// GetBillingPeriodOk returns a tuple with the BillingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetBillingPeriodOk() (*TypesBillingPeriod, bool) {
	if o == nil || IsNil(o.BillingPeriod) {
		return nil, false
	}
	return o.BillingPeriod, true
}

// HasBillingPeriod returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasBillingPeriod() bool {
	if o != nil && !IsNil(o.BillingPeriod) {
		return true
	}

	return false
}

// SetBillingPeriod gets a reference to the given TypesBillingPeriod and assigns it to the BillingPeriod field.
func (o *DtoPriceResponse) SetBillingPeriod(v TypesBillingPeriod) {
	o.BillingPeriod = &v
}

// GetBillingPeriodCount returns the BillingPeriodCount field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetBillingPeriodCount() int32 {
	if o == nil || IsNil(o.BillingPeriodCount) {
		var ret int32
		return ret
	}
	return *o.BillingPeriodCount
}

// GetBillingPeriodCountOk returns a tuple with the BillingPeriodCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetBillingPeriodCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BillingPeriodCount) {
		return nil, false
	}
	return o.BillingPeriodCount, true
}

// HasBillingPeriodCount returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasBillingPeriodCount() bool {
	if o != nil && !IsNil(o.BillingPeriodCount) {
		return true
	}

	return false
}

// SetBillingPeriodCount gets a reference to the given int32 and assigns it to the BillingPeriodCount field.
func (o *DtoPriceResponse) SetBillingPeriodCount(v int32) {
	o.BillingPeriodCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DtoPriceResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DtoPriceResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *DtoPriceResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DtoPriceResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayAmount returns the DisplayAmount field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetDisplayAmount() string {
	if o == nil || IsNil(o.DisplayAmount) {
		var ret string
		return ret
	}
	return *o.DisplayAmount
}

// GetDisplayAmountOk returns a tuple with the DisplayAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetDisplayAmountOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayAmount) {
		return nil, false
	}
	return o.DisplayAmount, true
}

// HasDisplayAmount returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasDisplayAmount() bool {
	if o != nil && !IsNil(o.DisplayAmount) {
		return true
	}

	return false
}

// SetDisplayAmount gets a reference to the given string and assigns it to the DisplayAmount field.
func (o *DtoPriceResponse) SetDisplayAmount(v string) {
	o.DisplayAmount = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *DtoPriceResponse) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DtoPriceResponse) SetId(v string) {
	o.Id = &v
}

// GetInvoiceCadence returns the InvoiceCadence field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetInvoiceCadence() TypesInvoiceCadence {
	if o == nil || IsNil(o.InvoiceCadence) {
		var ret TypesInvoiceCadence
		return ret
	}
	return *o.InvoiceCadence
}

// GetInvoiceCadenceOk returns a tuple with the InvoiceCadence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetInvoiceCadenceOk() (*TypesInvoiceCadence, bool) {
	if o == nil || IsNil(o.InvoiceCadence) {
		return nil, false
	}
	return o.InvoiceCadence, true
}

// HasInvoiceCadence returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasInvoiceCadence() bool {
	if o != nil && !IsNil(o.InvoiceCadence) {
		return true
	}

	return false
}

// SetInvoiceCadence gets a reference to the given TypesInvoiceCadence and assigns it to the InvoiceCadence field.
func (o *DtoPriceResponse) SetInvoiceCadence(v TypesInvoiceCadence) {
	o.InvoiceCadence = &v
}

// GetLookupKey returns the LookupKey field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetLookupKey() string {
	if o == nil || IsNil(o.LookupKey) {
		var ret string
		return ret
	}
	return *o.LookupKey
}

// GetLookupKeyOk returns a tuple with the LookupKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetLookupKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LookupKey) {
		return nil, false
	}
	return o.LookupKey, true
}

// HasLookupKey returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasLookupKey() bool {
	if o != nil && !IsNil(o.LookupKey) {
		return true
	}

	return false
}

// SetLookupKey gets a reference to the given string and assigns it to the LookupKey field.
func (o *DtoPriceResponse) SetLookupKey(v string) {
	o.LookupKey = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DtoPriceResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMeter returns the Meter field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetMeter() DtoMeterResponse {
	if o == nil || IsNil(o.Meter) {
		var ret DtoMeterResponse
		return ret
	}
	return *o.Meter
}

// GetMeterOk returns a tuple with the Meter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetMeterOk() (*DtoMeterResponse, bool) {
	if o == nil || IsNil(o.Meter) {
		return nil, false
	}
	return o.Meter, true
}

// HasMeter returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasMeter() bool {
	if o != nil && !IsNil(o.Meter) {
		return true
	}

	return false
}

// SetMeter gets a reference to the given DtoMeterResponse and assigns it to the Meter field.
func (o *DtoPriceResponse) SetMeter(v DtoMeterResponse) {
	o.Meter = &v
}

// GetMeterId returns the MeterId field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetMeterId() string {
	if o == nil || IsNil(o.MeterId) {
		var ret string
		return ret
	}
	return *o.MeterId
}

// GetMeterIdOk returns a tuple with the MeterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetMeterIdOk() (*string, bool) {
	if o == nil || IsNil(o.MeterId) {
		return nil, false
	}
	return o.MeterId, true
}

// HasMeterId returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasMeterId() bool {
	if o != nil && !IsNil(o.MeterId) {
		return true
	}

	return false
}

// SetMeterId gets a reference to the given string and assigns it to the MeterId field.
func (o *DtoPriceResponse) SetMeterId(v string) {
	o.MeterId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *DtoPriceResponse) SetPlanId(v string) {
	o.PlanId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetStatus() TypesStatus {
	if o == nil || IsNil(o.Status) {
		var ret TypesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetStatusOk() (*TypesStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TypesStatus and assigns it to the Status field.
func (o *DtoPriceResponse) SetStatus(v TypesStatus) {
	o.Status = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *DtoPriceResponse) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTierMode returns the TierMode field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetTierMode() TypesBillingTier {
	if o == nil || IsNil(o.TierMode) {
		var ret TypesBillingTier
		return ret
	}
	return *o.TierMode
}

// GetTierModeOk returns a tuple with the TierMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetTierModeOk() (*TypesBillingTier, bool) {
	if o == nil || IsNil(o.TierMode) {
		return nil, false
	}
	return o.TierMode, true
}

// HasTierMode returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasTierMode() bool {
	if o != nil && !IsNil(o.TierMode) {
		return true
	}

	return false
}

// SetTierMode gets a reference to the given TypesBillingTier and assigns it to the TierMode field.
func (o *DtoPriceResponse) SetTierMode(v TypesBillingTier) {
	o.TierMode = &v
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetTiers() []PricePriceTier {
	if o == nil || IsNil(o.Tiers) {
		var ret []PricePriceTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetTiersOk() ([]PricePriceTier, bool) {
	if o == nil || IsNil(o.Tiers) {
		return nil, false
	}
	return o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasTiers() bool {
	if o != nil && !IsNil(o.Tiers) {
		return true
	}

	return false
}

// SetTiers gets a reference to the given []PricePriceTier and assigns it to the Tiers field.
func (o *DtoPriceResponse) SetTiers(v []PricePriceTier) {
	o.Tiers = v
}

// GetTransformQuantity returns the TransformQuantity field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetTransformQuantity() PriceJSONBTransformQuantity {
	if o == nil || IsNil(o.TransformQuantity) {
		var ret PriceJSONBTransformQuantity
		return ret
	}
	return *o.TransformQuantity
}

// GetTransformQuantityOk returns a tuple with the TransformQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetTransformQuantityOk() (*PriceJSONBTransformQuantity, bool) {
	if o == nil || IsNil(o.TransformQuantity) {
		return nil, false
	}
	return o.TransformQuantity, true
}

// HasTransformQuantity returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasTransformQuantity() bool {
	if o != nil && !IsNil(o.TransformQuantity) {
		return true
	}

	return false
}

// SetTransformQuantity gets a reference to the given PriceJSONBTransformQuantity and assigns it to the TransformQuantity field.
func (o *DtoPriceResponse) SetTransformQuantity(v PriceJSONBTransformQuantity) {
	o.TransformQuantity = &v
}

// GetTrialPeriod returns the TrialPeriod field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetTrialPeriod() int32 {
	if o == nil || IsNil(o.TrialPeriod) {
		var ret int32
		return ret
	}
	return *o.TrialPeriod
}

// GetTrialPeriodOk returns a tuple with the TrialPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetTrialPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialPeriod) {
		return nil, false
	}
	return o.TrialPeriod, true
}

// HasTrialPeriod returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasTrialPeriod() bool {
	if o != nil && !IsNil(o.TrialPeriod) {
		return true
	}

	return false
}

// SetTrialPeriod gets a reference to the given int32 and assigns it to the TrialPeriod field.
func (o *DtoPriceResponse) SetTrialPeriod(v int32) {
	o.TrialPeriod = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetType() TypesPriceType {
	if o == nil || IsNil(o.Type) {
		var ret TypesPriceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetTypeOk() (*TypesPriceType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TypesPriceType and assigns it to the Type field.
func (o *DtoPriceResponse) SetType(v TypesPriceType) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DtoPriceResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *DtoPriceResponse) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoPriceResponse) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DtoPriceResponse) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *DtoPriceResponse) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

func (o DtoPriceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoPriceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BillingCadence) {
		toSerialize["billing_cadence"] = o.BillingCadence
	}
	if !IsNil(o.BillingModel) {
		toSerialize["billing_model"] = o.BillingModel
	}
	if !IsNil(o.BillingPeriod) {
		toSerialize["billing_period"] = o.BillingPeriod
	}
	if !IsNil(o.BillingPeriodCount) {
		toSerialize["billing_period_count"] = o.BillingPeriodCount
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
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayAmount) {
		toSerialize["display_amount"] = o.DisplayAmount
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
	if !IsNil(o.LookupKey) {
		toSerialize["lookup_key"] = o.LookupKey
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Meter) {
		toSerialize["meter"] = o.Meter
	}
	if !IsNil(o.MeterId) {
		toSerialize["meter_id"] = o.MeterId
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenant_id"] = o.TenantId
	}
	if !IsNil(o.TierMode) {
		toSerialize["tier_mode"] = o.TierMode
	}
	if !IsNil(o.Tiers) {
		toSerialize["tiers"] = o.Tiers
	}
	if !IsNil(o.TransformQuantity) {
		toSerialize["transform_quantity"] = o.TransformQuantity
	}
	if !IsNil(o.TrialPeriod) {
		toSerialize["trial_period"] = o.TrialPeriod
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	return toSerialize, nil
}

type NullableDtoPriceResponse struct {
	value *DtoPriceResponse
	isSet bool
}

func (v NullableDtoPriceResponse) Get() *DtoPriceResponse {
	return v.value
}

func (v *NullableDtoPriceResponse) Set(val *DtoPriceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoPriceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoPriceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoPriceResponse(val *DtoPriceResponse) *NullableDtoPriceResponse {
	return &NullableDtoPriceResponse{value: val, isSet: true}
}

func (v NullableDtoPriceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoPriceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


