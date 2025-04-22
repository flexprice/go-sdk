# DtoTopUpWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | amount is the amount in the currency of the wallet to be added NOTE: this is not the number of credits to add, but the amount in the currency amount &#x3D; credits_to_add * conversion_rate if both amount and credits_to_add are provided, amount will be ignored ex if the wallet has a conversion_rate of 2 then adding an amount of 10 USD in the wallet wil add 5 credits in the wallet | [optional] 
**CreditsToAdd** | Pointer to **float32** | credits_to_add is the number of credits to add to the wallet | [optional] 
**Description** | Pointer to **string** | description to add any specific details about the transaction | [optional] 
**ExpiryDateUtc** | Pointer to **string** | expiry_date_utc is the expiry date in UTC timezone ex 2025-01-01 00:00:00 UTC | [optional] 
**IdempotencyKey** | **string** | idempotency_key is a unique key for the transaction | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**TransactionReason** | [**TypesTransactionReason**](TypesTransactionReason.md) |  | 

## Methods

### NewDtoTopUpWalletRequest

`func NewDtoTopUpWalletRequest(idempotencyKey string, transactionReason TypesTransactionReason, ) *DtoTopUpWalletRequest`

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

### HasAmount

`func (o *DtoTopUpWalletRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreditsToAdd

`func (o *DtoTopUpWalletRequest) GetCreditsToAdd() float32`

GetCreditsToAdd returns the CreditsToAdd field if non-nil, zero value otherwise.

### GetCreditsToAddOk

`func (o *DtoTopUpWalletRequest) GetCreditsToAddOk() (*float32, bool)`

GetCreditsToAddOk returns a tuple with the CreditsToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsToAdd

`func (o *DtoTopUpWalletRequest) SetCreditsToAdd(v float32)`

SetCreditsToAdd sets CreditsToAdd field to given value.

### HasCreditsToAdd

`func (o *DtoTopUpWalletRequest) HasCreditsToAdd() bool`

HasCreditsToAdd returns a boolean if a field has been set.

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

### GetExpiryDateUtc

`func (o *DtoTopUpWalletRequest) GetExpiryDateUtc() string`

GetExpiryDateUtc returns the ExpiryDateUtc field if non-nil, zero value otherwise.

### GetExpiryDateUtcOk

`func (o *DtoTopUpWalletRequest) GetExpiryDateUtcOk() (*string, bool)`

GetExpiryDateUtcOk returns a tuple with the ExpiryDateUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateUtc

`func (o *DtoTopUpWalletRequest) SetExpiryDateUtc(v string)`

SetExpiryDateUtc sets ExpiryDateUtc field to given value.

### HasExpiryDateUtc

`func (o *DtoTopUpWalletRequest) HasExpiryDateUtc() bool`

HasExpiryDateUtc returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *DtoTopUpWalletRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoTopUpWalletRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoTopUpWalletRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


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

### GetTransactionReason

`func (o *DtoTopUpWalletRequest) GetTransactionReason() TypesTransactionReason`

GetTransactionReason returns the TransactionReason field if non-nil, zero value otherwise.

### GetTransactionReasonOk

`func (o *DtoTopUpWalletRequest) GetTransactionReasonOk() (*TypesTransactionReason, bool)`

GetTransactionReasonOk returns a tuple with the TransactionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReason

`func (o *DtoTopUpWalletRequest) SetTransactionReason(v TypesTransactionReason)`

SetTransactionReason sets TransactionReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


