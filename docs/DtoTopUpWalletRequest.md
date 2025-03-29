# DtoTopUpWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** | amount is the number of credits to add to the wallet | 
**Description** | Pointer to **string** | description to add any specific details about the transaction | [optional] 
**ExpiryDate** | Pointer to **int32** | expiry_date YYYYMMDD format in UTC timezone (optional to set nil means no expiry) for ex 20250101 means the credits will expire on 2025-01-01 00:00:00 UTC hence they will be available for use until 2024-12-31 23:59:59 UTC | [optional] 
**GenerateInvoice** | Pointer to **bool** | generate_invoice when true, an invoice will be generated for the transaction | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PurchasedCredits** | Pointer to **bool** | purchased_credits when true, the credits are added as purchased credits | [optional] 
**ReferenceId** | Pointer to **string** | reference_id is the ID of the reference ex payment ID, invoice ID, request ID | [optional] 
**ReferenceType** | Pointer to **string** | reference_type is the type of the reference ex payment, invoice, request | [optional] 

## Methods

### NewDtoTopUpWalletRequest

`func NewDtoTopUpWalletRequest(amount float32, ) *DtoTopUpWalletRequest`

NewDtoTopUpWalletRequest instantiates a new DtoTopUpWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTopUpWalletRequestWithDefaults

`func NewDtoTopUpWalletRequestWithDefaults() *DtoTopUpWalletRequest`

NewDtoTopUpWalletRequestWithDefaults instantiates a new DtoTopUpWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoTopUpWalletRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoTopUpWalletRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoTopUpWalletRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *DtoTopUpWalletRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoTopUpWalletRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoTopUpWalletRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoTopUpWalletRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiryDate

`func (o *DtoTopUpWalletRequest) GetExpiryDate() int32`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *DtoTopUpWalletRequest) GetExpiryDateOk() (*int32, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *DtoTopUpWalletRequest) SetExpiryDate(v int32)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *DtoTopUpWalletRequest) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetGenerateInvoice

`func (o *DtoTopUpWalletRequest) GetGenerateInvoice() bool`

GetGenerateInvoice returns the GenerateInvoice field if non-nil, zero value otherwise.

### GetGenerateInvoiceOk

`func (o *DtoTopUpWalletRequest) GetGenerateInvoiceOk() (*bool, bool)`

GetGenerateInvoiceOk returns a tuple with the GenerateInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateInvoice

`func (o *DtoTopUpWalletRequest) SetGenerateInvoice(v bool)`

SetGenerateInvoice sets GenerateInvoice field to given value.

### HasGenerateInvoice

`func (o *DtoTopUpWalletRequest) HasGenerateInvoice() bool`

HasGenerateInvoice returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoTopUpWalletRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoTopUpWalletRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoTopUpWalletRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoTopUpWalletRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPurchasedCredits

`func (o *DtoTopUpWalletRequest) GetPurchasedCredits() bool`

GetPurchasedCredits returns the PurchasedCredits field if non-nil, zero value otherwise.

### GetPurchasedCreditsOk

`func (o *DtoTopUpWalletRequest) GetPurchasedCreditsOk() (*bool, bool)`

GetPurchasedCreditsOk returns a tuple with the PurchasedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedCredits

`func (o *DtoTopUpWalletRequest) SetPurchasedCredits(v bool)`

SetPurchasedCredits sets PurchasedCredits field to given value.

### HasPurchasedCredits

`func (o *DtoTopUpWalletRequest) HasPurchasedCredits() bool`

HasPurchasedCredits returns a boolean if a field has been set.

### GetReferenceId

`func (o *DtoTopUpWalletRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *DtoTopUpWalletRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *DtoTopUpWalletRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *DtoTopUpWalletRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetReferenceType

`func (o *DtoTopUpWalletRequest) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *DtoTopUpWalletRequest) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *DtoTopUpWalletRequest) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *DtoTopUpWalletRequest) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


