# DtoManualBalanceDebitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credits** | Pointer to **string** | credits is the number of credits to debit from the wallet | [optional] 
**Description** | Pointer to **string** | description to add any specific details about the transaction | [optional] 
**IdempotencyKey** | **string** | idempotency_key is a unique key for the transaction | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**TransactionReason** | [**TypesTransactionReason**](TypesTransactionReason.md) |  | 

## Methods

### NewDtoManualBalanceDebitRequest

`func NewDtoManualBalanceDebitRequest(idempotencyKey string, transactionReason TypesTransactionReason, ) *DtoManualBalanceDebitRequest`

NewDtoManualBalanceDebitRequest instantiates a new DtoManualBalanceDebitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoManualBalanceDebitRequestWithDefaults

`func NewDtoManualBalanceDebitRequestWithDefaults() *DtoManualBalanceDebitRequest`

NewDtoManualBalanceDebitRequestWithDefaults instantiates a new DtoManualBalanceDebitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredits

`func (o *DtoManualBalanceDebitRequest) GetCredits() string`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *DtoManualBalanceDebitRequest) GetCreditsOk() (*string, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *DtoManualBalanceDebitRequest) SetCredits(v string)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *DtoManualBalanceDebitRequest) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetDescription

`func (o *DtoManualBalanceDebitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoManualBalanceDebitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoManualBalanceDebitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoManualBalanceDebitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *DtoManualBalanceDebitRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoManualBalanceDebitRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoManualBalanceDebitRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetMetadata

`func (o *DtoManualBalanceDebitRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoManualBalanceDebitRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoManualBalanceDebitRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoManualBalanceDebitRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTransactionReason

`func (o *DtoManualBalanceDebitRequest) GetTransactionReason() TypesTransactionReason`

GetTransactionReason returns the TransactionReason field if non-nil, zero value otherwise.

### GetTransactionReasonOk

`func (o *DtoManualBalanceDebitRequest) GetTransactionReasonOk() (*TypesTransactionReason, bool)`

GetTransactionReasonOk returns a tuple with the TransactionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReason

`func (o *DtoManualBalanceDebitRequest) SetTransactionReason(v TypesTransactionReason)`

SetTransactionReason sets TransactionReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


