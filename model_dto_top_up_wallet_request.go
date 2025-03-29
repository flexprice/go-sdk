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

// checks if the DtoTopUpWalletRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoTopUpWalletRequest{}

// DtoTopUpWalletRequest struct for DtoTopUpWalletRequest
type DtoTopUpWalletRequest struct {
	// amount is the number of credits to add to the wallet
	Amount float32 `json:"amount"`
	// description to add any specific details about the transaction
	Description *string `json:"description,omitempty"`
	// expiry_date YYYYMMDD format in UTC timezone (optional to set nil means no expiry) for ex 20250101 means the credits will expire on 2025-01-01 00:00:00 UTC hence they will be available for use until 2024-12-31 23:59:59 UTC
	ExpiryDate *int32 `json:"expiry_date,omitempty"`
	// generate_invoice when true, an invoice will be generated for the transaction
	GenerateInvoice *bool `json:"generate_invoice,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	// purchased_credits when true, the credits are added as purchased credits
	PurchasedCredits *bool `json:"purchased_credits,omitempty"`
	// reference_id is the ID of the reference ex payment ID, invoice ID, request ID
	ReferenceId *string `json:"reference_id,omitempty"`
	// reference_type is the type of the reference ex payment, invoice, request
	ReferenceType *string `json:"reference_type,omitempty"`
}

type _DtoTopUpWalletRequest DtoTopUpWalletRequest

// NewDtoTopUpWalletRequest instantiates a new DtoTopUpWalletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoTopUpWalletRequest(amount float32) *DtoTopUpWalletRequest {
	this := DtoTopUpWalletRequest{}
	this.Amount = amount
	return &this
}

// NewDtoTopUpWalletRequestWithDefaults instantiates a new DtoTopUpWalletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoTopUpWalletRequestWithDefaults() *DtoTopUpWalletRequest {
	this := DtoTopUpWalletRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *DtoTopUpWalletRequest) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *DtoTopUpWalletRequest) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *DtoTopUpWalletRequest) SetAmount(v float32) {
	o.Amount = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DtoTopUpWalletRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTopUpWalletRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DtoTopUpWalletRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DtoTopUpWalletRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *DtoTopUpWalletRequest) GetExpiryDate() int32 {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret int32
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTopUpWalletRequest) GetExpiryDateOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *DtoTopUpWalletRequest) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given int32 and assigns it to the ExpiryDate field.
func (o *DtoTopUpWalletRequest) SetExpiryDate(v int32) {
	o.ExpiryDate = &v
}

// GetGenerateInvoice returns the GenerateInvoice field value if set, zero value otherwise.
func (o *DtoTopUpWalletRequest) GetGenerateInvoice() bool {
	if o == nil || IsNil(o.GenerateInvoice) {
		var ret bool
		return ret
	}
	return *o.GenerateInvoice
}

// GetGenerateInvoiceOk returns a tuple with the GenerateInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTopUpWalletRequest) GetGenerateInvoiceOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerateInvoice) {
		return nil, false
	}
	return o.GenerateInvoice, true
}

// HasGenerateInvoice returns a boolean if a field has been set.
func (o *DtoTopUpWalletRequest) HasGenerateInvoice() bool {
	if o != nil && !IsNil(o.GenerateInvoice) {
		return true
	}

	return false
}

// SetGenerateInvoice gets a reference to the given bool and assigns it to the GenerateInvoice field.
func (o *DtoTopUpWalletRequest) SetGenerateInvoice(v bool) {
	o.GenerateInvoice = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DtoTopUpWalletRequest) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTopUpWalletRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DtoTopUpWalletRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DtoTopUpWalletRequest) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetPurchasedCredits returns the PurchasedCredits field value if set, zero value otherwise.
func (o *DtoTopUpWalletRequest) GetPurchasedCredits() bool {
	if o == nil || IsNil(o.PurchasedCredits) {
		var ret bool
		return ret
	}
	return *o.PurchasedCredits
}

// GetPurchasedCreditsOk returns a tuple with the PurchasedCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTopUpWalletRequest) GetPurchasedCreditsOk() (*bool, bool) {
	if o == nil || IsNil(o.PurchasedCredits) {
		return nil, false
	}
	return o.PurchasedCredits, true
}

// HasPurchasedCredits returns a boolean if a field has been set.
func (o *DtoTopUpWalletRequest) HasPurchasedCredits() bool {
	if o != nil && !IsNil(o.PurchasedCredits) {
		return true
	}

	return false
}

// SetPurchasedCredits gets a reference to the given bool and assigns it to the PurchasedCredits field.
func (o *DtoTopUpWalletRequest) SetPurchasedCredits(v bool) {
	o.PurchasedCredits = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *DtoTopUpWalletRequest) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTopUpWalletRequest) GetReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *DtoTopUpWalletRequest) HasReferenceId() bool {
	if o != nil && !IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *DtoTopUpWalletRequest) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetReferenceType returns the ReferenceType field value if set, zero value otherwise.
func (o *DtoTopUpWalletRequest) GetReferenceType() string {
	if o == nil || IsNil(o.ReferenceType) {
		var ret string
		return ret
	}
	return *o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoTopUpWalletRequest) GetReferenceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceType) {
		return nil, false
	}
	return o.ReferenceType, true
}

// HasReferenceType returns a boolean if a field has been set.
func (o *DtoTopUpWalletRequest) HasReferenceType() bool {
	if o != nil && !IsNil(o.ReferenceType) {
		return true
	}

	return false
}

// SetReferenceType gets a reference to the given string and assigns it to the ReferenceType field.
func (o *DtoTopUpWalletRequest) SetReferenceType(v string) {
	o.ReferenceType = &v
}

func (o DtoTopUpWalletRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoTopUpWalletRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpiryDate) {
		toSerialize["expiry_date"] = o.ExpiryDate
	}
	if !IsNil(o.GenerateInvoice) {
		toSerialize["generate_invoice"] = o.GenerateInvoice
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PurchasedCredits) {
		toSerialize["purchased_credits"] = o.PurchasedCredits
	}
	if !IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !IsNil(o.ReferenceType) {
		toSerialize["reference_type"] = o.ReferenceType
	}
	return toSerialize, nil
}

func (o *DtoTopUpWalletRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
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

	varDtoTopUpWalletRequest := _DtoTopUpWalletRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDtoTopUpWalletRequest)

	if err != nil {
		return err
	}

	*o = DtoTopUpWalletRequest(varDtoTopUpWalletRequest)

	return err
}

type NullableDtoTopUpWalletRequest struct {
	value *DtoTopUpWalletRequest
	isSet bool
}

func (v NullableDtoTopUpWalletRequest) Get() *DtoTopUpWalletRequest {
	return v.value
}

func (v *NullableDtoTopUpWalletRequest) Set(val *DtoTopUpWalletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoTopUpWalletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoTopUpWalletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoTopUpWalletRequest(val *DtoTopUpWalletRequest) *NullableDtoTopUpWalletRequest {
	return &NullableDtoTopUpWalletRequest{value: val, isSet: true}
}

func (v NullableDtoTopUpWalletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoTopUpWalletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


