# DtoWalletTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreditAmount** | Pointer to **float32** |  | [optional] 
**CreditBalanceAfter** | Pointer to **float32** |  | [optional] 
**CreditBalanceBefore** | Pointer to **float32** |  | [optional] 
**CreditsAvailable** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExpiryDate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ReferenceType** | Pointer to [**TypesWalletTxReferenceType**](TypesWalletTxReferenceType.md) |  | [optional] 
**TransactionReason** | Pointer to [**TypesTransactionReason**](TypesTransactionReason.md) |  | [optional] 
**TransactionStatus** | Pointer to [**TypesTransactionStatus**](TypesTransactionStatus.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**WalletId** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoWalletTransactionResponse

`func NewDtoWalletTransactionResponse() *DtoWalletTransactionResponse`

NewDtoWalletTransactionResponse instantiates a new DtoWalletTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoWalletTransactionResponseWithDefaults

`func NewDtoWalletTransactionResponseWithDefaults() *DtoWalletTransactionResponse`

NewDtoWalletTransactionResponseWithDefaults instantiates a new DtoWalletTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DtoWalletTransactionResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoWalletTransactionResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoWalletTransactionResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DtoWalletTransactionResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoWalletTransactionResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoWalletTransactionResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoWalletTransactionResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoWalletTransactionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditAmount

`func (o *DtoWalletTransactionResponse) GetCreditAmount() float32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *DtoWalletTransactionResponse) GetCreditAmountOk() (*float32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *DtoWalletTransactionResponse) SetCreditAmount(v float32)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *DtoWalletTransactionResponse) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetCreditBalanceAfter

`func (o *DtoWalletTransactionResponse) GetCreditBalanceAfter() float32`

GetCreditBalanceAfter returns the CreditBalanceAfter field if non-nil, zero value otherwise.

### GetCreditBalanceAfterOk

`func (o *DtoWalletTransactionResponse) GetCreditBalanceAfterOk() (*float32, bool)`

GetCreditBalanceAfterOk returns a tuple with the CreditBalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalanceAfter

`func (o *DtoWalletTransactionResponse) SetCreditBalanceAfter(v float32)`

SetCreditBalanceAfter sets CreditBalanceAfter field to given value.

### HasCreditBalanceAfter

`func (o *DtoWalletTransactionResponse) HasCreditBalanceAfter() bool`

HasCreditBalanceAfter returns a boolean if a field has been set.

### GetCreditBalanceBefore

`func (o *DtoWalletTransactionResponse) GetCreditBalanceBefore() float32`

GetCreditBalanceBefore returns the CreditBalanceBefore field if non-nil, zero value otherwise.

### GetCreditBalanceBeforeOk

`func (o *DtoWalletTransactionResponse) GetCreditBalanceBeforeOk() (*float32, bool)`

GetCreditBalanceBeforeOk returns a tuple with the CreditBalanceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalanceBefore

`func (o *DtoWalletTransactionResponse) SetCreditBalanceBefore(v float32)`

SetCreditBalanceBefore sets CreditBalanceBefore field to given value.

### HasCreditBalanceBefore

`func (o *DtoWalletTransactionResponse) HasCreditBalanceBefore() bool`

HasCreditBalanceBefore returns a boolean if a field has been set.

### GetCreditsAvailable

`func (o *DtoWalletTransactionResponse) GetCreditsAvailable() float32`

GetCreditsAvailable returns the CreditsAvailable field if non-nil, zero value otherwise.

### GetCreditsAvailableOk

`func (o *DtoWalletTransactionResponse) GetCreditsAvailableOk() (*float32, bool)`

GetCreditsAvailableOk returns a tuple with the CreditsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsAvailable

`func (o *DtoWalletTransactionResponse) SetCreditsAvailable(v float32)`

SetCreditsAvailable sets CreditsAvailable field to given value.

### HasCreditsAvailable

`func (o *DtoWalletTransactionResponse) HasCreditsAvailable() bool`

HasCreditsAvailable returns a boolean if a field has been set.

### GetDescription

`func (o *DtoWalletTransactionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DtoWalletTransactionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DtoWalletTransactionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DtoWalletTransactionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiryDate

`func (o *DtoWalletTransactionResponse) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *DtoWalletTransactionResponse) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *DtoWalletTransactionResponse) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *DtoWalletTransactionResponse) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetId

`func (o *DtoWalletTransactionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoWalletTransactionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoWalletTransactionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoWalletTransactionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoWalletTransactionResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoWalletTransactionResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoWalletTransactionResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoWalletTransactionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReferenceId

`func (o *DtoWalletTransactionResponse) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *DtoWalletTransactionResponse) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *DtoWalletTransactionResponse) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *DtoWalletTransactionResponse) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetReferenceType

`func (o *DtoWalletTransactionResponse) GetReferenceType() TypesWalletTxReferenceType`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *DtoWalletTransactionResponse) GetReferenceTypeOk() (*TypesWalletTxReferenceType, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *DtoWalletTransactionResponse) SetReferenceType(v TypesWalletTxReferenceType)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *DtoWalletTransactionResponse) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetTransactionReason

`func (o *DtoWalletTransactionResponse) GetTransactionReason() TypesTransactionReason`

GetTransactionReason returns the TransactionReason field if non-nil, zero value otherwise.

### GetTransactionReasonOk

`func (o *DtoWalletTransactionResponse) GetTransactionReasonOk() (*TypesTransactionReason, bool)`

GetTransactionReasonOk returns a tuple with the TransactionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReason

`func (o *DtoWalletTransactionResponse) SetTransactionReason(v TypesTransactionReason)`

SetTransactionReason sets TransactionReason field to given value.

### HasTransactionReason

`func (o *DtoWalletTransactionResponse) HasTransactionReason() bool`

HasTransactionReason returns a boolean if a field has been set.

### GetTransactionStatus

`func (o *DtoWalletTransactionResponse) GetTransactionStatus() TypesTransactionStatus`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *DtoWalletTransactionResponse) GetTransactionStatusOk() (*TypesTransactionStatus, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *DtoWalletTransactionResponse) SetTransactionStatus(v TypesTransactionStatus)`

SetTransactionStatus sets TransactionStatus field to given value.

### HasTransactionStatus

`func (o *DtoWalletTransactionResponse) HasTransactionStatus() bool`

HasTransactionStatus returns a boolean if a field has been set.

### GetType

`func (o *DtoWalletTransactionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoWalletTransactionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoWalletTransactionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DtoWalletTransactionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWalletId

`func (o *DtoWalletTransactionResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *DtoWalletTransactionResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *DtoWalletTransactionResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *DtoWalletTransactionResponse) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


