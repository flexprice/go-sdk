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

// checks if the DtoWalletTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtoWalletTransactionResponse{}

// DtoWalletTransactionResponse struct for DtoWalletTransactionResponse
type DtoWalletTransactionResponse struct {
	Amount *float32 `json:"amount,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CreditAmount *float32 `json:"credit_amount,omitempty"`
	CreditBalanceAfter *float32 `json:"credit_balance_after,omitempty"`
	CreditBalanceBefore *float32 `json:"credit_balance_before,omitempty"`
	CreditsAvailable *float32 `json:"credits_available,omitempty"`
	Description *string `json:"description,omitempty"`
	ExpiryDate *string `json:"expiry_date,omitempty"`
	Id *string `json:"id,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	ReferenceId *string `json:"reference_id,omitempty"`
	ReferenceType *TypesWalletTxReferenceType `json:"reference_type,omitempty"`
	TransactionReason *TypesTransactionReason `json:"transaction_reason,omitempty"`
	TransactionStatus *TypesTransactionStatus `json:"transaction_status,omitempty"`
	Type *string `json:"type,omitempty"`
	WalletId *string `json:"wallet_id,omitempty"`
}

// NewDtoWalletTransactionResponse instantiates a new DtoWalletTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtoWalletTransactionResponse() *DtoWalletTransactionResponse {
	this := DtoWalletTransactionResponse{}
	return &this
}

// NewDtoWalletTransactionResponseWithDefaults instantiates a new DtoWalletTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtoWalletTransactionResponseWithDefaults() *DtoWalletTransactionResponse {
	this := DtoWalletTransactionResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *DtoWalletTransactionResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DtoWalletTransactionResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreditAmount returns the CreditAmount field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetCreditAmount() float32 {
	if o == nil || IsNil(o.CreditAmount) {
		var ret float32
		return ret
	}
	return *o.CreditAmount
}

// GetCreditAmountOk returns a tuple with the CreditAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetCreditAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CreditAmount) {
		return nil, false
	}
	return o.CreditAmount, true
}

// HasCreditAmount returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasCreditAmount() bool {
	if o != nil && !IsNil(o.CreditAmount) {
		return true
	}

	return false
}

// SetCreditAmount gets a reference to the given float32 and assigns it to the CreditAmount field.
func (o *DtoWalletTransactionResponse) SetCreditAmount(v float32) {
	o.CreditAmount = &v
}

// GetCreditBalanceAfter returns the CreditBalanceAfter field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetCreditBalanceAfter() float32 {
	if o == nil || IsNil(o.CreditBalanceAfter) {
		var ret float32
		return ret
	}
	return *o.CreditBalanceAfter
}

// GetCreditBalanceAfterOk returns a tuple with the CreditBalanceAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetCreditBalanceAfterOk() (*float32, bool) {
	if o == nil || IsNil(o.CreditBalanceAfter) {
		return nil, false
	}
	return o.CreditBalanceAfter, true
}

// HasCreditBalanceAfter returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasCreditBalanceAfter() bool {
	if o != nil && !IsNil(o.CreditBalanceAfter) {
		return true
	}

	return false
}

// SetCreditBalanceAfter gets a reference to the given float32 and assigns it to the CreditBalanceAfter field.
func (o *DtoWalletTransactionResponse) SetCreditBalanceAfter(v float32) {
	o.CreditBalanceAfter = &v
}

// GetCreditBalanceBefore returns the CreditBalanceBefore field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetCreditBalanceBefore() float32 {
	if o == nil || IsNil(o.CreditBalanceBefore) {
		var ret float32
		return ret
	}
	return *o.CreditBalanceBefore
}

// GetCreditBalanceBeforeOk returns a tuple with the CreditBalanceBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetCreditBalanceBeforeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreditBalanceBefore) {
		return nil, false
	}
	return o.CreditBalanceBefore, true
}

// HasCreditBalanceBefore returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasCreditBalanceBefore() bool {
	if o != nil && !IsNil(o.CreditBalanceBefore) {
		return true
	}

	return false
}

// SetCreditBalanceBefore gets a reference to the given float32 and assigns it to the CreditBalanceBefore field.
func (o *DtoWalletTransactionResponse) SetCreditBalanceBefore(v float32) {
	o.CreditBalanceBefore = &v
}

// GetCreditsAvailable returns the CreditsAvailable field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetCreditsAvailable() float32 {
	if o == nil || IsNil(o.CreditsAvailable) {
		var ret float32
		return ret
	}
	return *o.CreditsAvailable
}

// GetCreditsAvailableOk returns a tuple with the CreditsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetCreditsAvailableOk() (*float32, bool) {
	if o == nil || IsNil(o.CreditsAvailable) {
		return nil, false
	}
	return o.CreditsAvailable, true
}

// HasCreditsAvailable returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasCreditsAvailable() bool {
	if o != nil && !IsNil(o.CreditsAvailable) {
		return true
	}

	return false
}

// SetCreditsAvailable gets a reference to the given float32 and assigns it to the CreditsAvailable field.
func (o *DtoWalletTransactionResponse) SetCreditsAvailable(v float32) {
	o.CreditsAvailable = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DtoWalletTransactionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetExpiryDate() string {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetExpiryDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *DtoWalletTransactionResponse) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DtoWalletTransactionResponse) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DtoWalletTransactionResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetReferenceId() string {
	if o == nil || IsNil(o.ReferenceId) {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceId) {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasReferenceId() bool {
	if o != nil && !IsNil(o.ReferenceId) {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *DtoWalletTransactionResponse) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetReferenceType returns the ReferenceType field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetReferenceType() TypesWalletTxReferenceType {
	if o == nil || IsNil(o.ReferenceType) {
		var ret TypesWalletTxReferenceType
		return ret
	}
	return *o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetReferenceTypeOk() (*TypesWalletTxReferenceType, bool) {
	if o == nil || IsNil(o.ReferenceType) {
		return nil, false
	}
	return o.ReferenceType, true
}

// HasReferenceType returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasReferenceType() bool {
	if o != nil && !IsNil(o.ReferenceType) {
		return true
	}

	return false
}

// SetReferenceType gets a reference to the given TypesWalletTxReferenceType and assigns it to the ReferenceType field.
func (o *DtoWalletTransactionResponse) SetReferenceType(v TypesWalletTxReferenceType) {
	o.ReferenceType = &v
}

// GetTransactionReason returns the TransactionReason field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetTransactionReason() TypesTransactionReason {
	if o == nil || IsNil(o.TransactionReason) {
		var ret TypesTransactionReason
		return ret
	}
	return *o.TransactionReason
}

// GetTransactionReasonOk returns a tuple with the TransactionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetTransactionReasonOk() (*TypesTransactionReason, bool) {
	if o == nil || IsNil(o.TransactionReason) {
		return nil, false
	}
	return o.TransactionReason, true
}

// HasTransactionReason returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasTransactionReason() bool {
	if o != nil && !IsNil(o.TransactionReason) {
		return true
	}

	return false
}

// SetTransactionReason gets a reference to the given TypesTransactionReason and assigns it to the TransactionReason field.
func (o *DtoWalletTransactionResponse) SetTransactionReason(v TypesTransactionReason) {
	o.TransactionReason = &v
}

// GetTransactionStatus returns the TransactionStatus field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetTransactionStatus() TypesTransactionStatus {
	if o == nil || IsNil(o.TransactionStatus) {
		var ret TypesTransactionStatus
		return ret
	}
	return *o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetTransactionStatusOk() (*TypesTransactionStatus, bool) {
	if o == nil || IsNil(o.TransactionStatus) {
		return nil, false
	}
	return o.TransactionStatus, true
}

// HasTransactionStatus returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasTransactionStatus() bool {
	if o != nil && !IsNil(o.TransactionStatus) {
		return true
	}

	return false
}

// SetTransactionStatus gets a reference to the given TypesTransactionStatus and assigns it to the TransactionStatus field.
func (o *DtoWalletTransactionResponse) SetTransactionStatus(v TypesTransactionStatus) {
	o.TransactionStatus = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DtoWalletTransactionResponse) SetType(v string) {
	o.Type = &v
}

// GetWalletId returns the WalletId field value if set, zero value otherwise.
func (o *DtoWalletTransactionResponse) GetWalletId() string {
	if o == nil || IsNil(o.WalletId) {
		var ret string
		return ret
	}
	return *o.WalletId
}

// GetWalletIdOk returns a tuple with the WalletId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtoWalletTransactionResponse) GetWalletIdOk() (*string, bool) {
	if o == nil || IsNil(o.WalletId) {
		return nil, false
	}
	return o.WalletId, true
}

// HasWalletId returns a boolean if a field has been set.
func (o *DtoWalletTransactionResponse) HasWalletId() bool {
	if o != nil && !IsNil(o.WalletId) {
		return true
	}

	return false
}

// SetWalletId gets a reference to the given string and assigns it to the WalletId field.
func (o *DtoWalletTransactionResponse) SetWalletId(v string) {
	o.WalletId = &v
}

func (o DtoWalletTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtoWalletTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreditAmount) {
		toSerialize["credit_amount"] = o.CreditAmount
	}
	if !IsNil(o.CreditBalanceAfter) {
		toSerialize["credit_balance_after"] = o.CreditBalanceAfter
	}
	if !IsNil(o.CreditBalanceBefore) {
		toSerialize["credit_balance_before"] = o.CreditBalanceBefore
	}
	if !IsNil(o.CreditsAvailable) {
		toSerialize["credits_available"] = o.CreditsAvailable
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpiryDate) {
		toSerialize["expiry_date"] = o.ExpiryDate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ReferenceId) {
		toSerialize["reference_id"] = o.ReferenceId
	}
	if !IsNil(o.ReferenceType) {
		toSerialize["reference_type"] = o.ReferenceType
	}
	if !IsNil(o.TransactionReason) {
		toSerialize["transaction_reason"] = o.TransactionReason
	}
	if !IsNil(o.TransactionStatus) {
		toSerialize["transaction_status"] = o.TransactionStatus
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.WalletId) {
		toSerialize["wallet_id"] = o.WalletId
	}
	return toSerialize, nil
}

type NullableDtoWalletTransactionResponse struct {
	value *DtoWalletTransactionResponse
	isSet bool
}

func (v NullableDtoWalletTransactionResponse) Get() *DtoWalletTransactionResponse {
	return v.value
}

func (v *NullableDtoWalletTransactionResponse) Set(val *DtoWalletTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDtoWalletTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDtoWalletTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtoWalletTransactionResponse(val *DtoWalletTransactionResponse) *NullableDtoWalletTransactionResponse {
	return &NullableDtoWalletTransactionResponse{value: val, isSet: true}
}

func (v NullableDtoWalletTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtoWalletTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


