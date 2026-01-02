# DtoWalletTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedByUser** | Pointer to [**DtoUserResponse**](DtoUserResponse.md) |  | [optional] 
**CreditAmount** | Pointer to **string** |  | [optional] 
**CreditBalanceAfter** | Pointer to **string** |  | [optional] 
**CreditBalanceBefore** | Pointer to **string** |  | [optional] 
**CreditsAvailable** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**DtoCustomerResponse**](DtoCustomerResponse.md) |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**ExpiryDate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdempotencyKey** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**ReferenceType** | Pointer to [**TypesWalletTxReferenceType**](TypesWalletTxReferenceType.md) |  | [optional] 
**Status** | Pointer to [**TypesStatus**](TypesStatus.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TransactionReason** | Pointer to [**TypesTransactionReason**](TypesTransactionReason.md) |  | [optional] 
**TransactionStatus** | Pointer to [**TypesTransactionStatus**](TypesTransactionStatus.md) |  | [optional] 
**Type** | Pointer to [**TypesTransactionType**](TypesTransactionType.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Wallet** | Pointer to [**DtoWalletResponse**](DtoWalletResponse.md) |  | [optional] 
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

`func (o *DtoWalletTransactionResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DtoWalletTransactionResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DtoWalletTransactionResponse) SetAmount(v string)`

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

### GetCreatedBy

`func (o *DtoWalletTransactionResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoWalletTransactionResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoWalletTransactionResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoWalletTransactionResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUser

`func (o *DtoWalletTransactionResponse) GetCreatedByUser() DtoUserResponse`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *DtoWalletTransactionResponse) GetCreatedByUserOk() (*DtoUserResponse, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *DtoWalletTransactionResponse) SetCreatedByUser(v DtoUserResponse)`

SetCreatedByUser sets CreatedByUser field to given value.

### HasCreatedByUser

`func (o *DtoWalletTransactionResponse) HasCreatedByUser() bool`

HasCreatedByUser returns a boolean if a field has been set.

### GetCreditAmount

`func (o *DtoWalletTransactionResponse) GetCreditAmount() string`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *DtoWalletTransactionResponse) GetCreditAmountOk() (*string, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *DtoWalletTransactionResponse) SetCreditAmount(v string)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *DtoWalletTransactionResponse) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetCreditBalanceAfter

`func (o *DtoWalletTransactionResponse) GetCreditBalanceAfter() string`

GetCreditBalanceAfter returns the CreditBalanceAfter field if non-nil, zero value otherwise.

### GetCreditBalanceAfterOk

`func (o *DtoWalletTransactionResponse) GetCreditBalanceAfterOk() (*string, bool)`

GetCreditBalanceAfterOk returns a tuple with the CreditBalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalanceAfter

`func (o *DtoWalletTransactionResponse) SetCreditBalanceAfter(v string)`

SetCreditBalanceAfter sets CreditBalanceAfter field to given value.

### HasCreditBalanceAfter

`func (o *DtoWalletTransactionResponse) HasCreditBalanceAfter() bool`

HasCreditBalanceAfter returns a boolean if a field has been set.

### GetCreditBalanceBefore

`func (o *DtoWalletTransactionResponse) GetCreditBalanceBefore() string`

GetCreditBalanceBefore returns the CreditBalanceBefore field if non-nil, zero value otherwise.

### GetCreditBalanceBeforeOk

`func (o *DtoWalletTransactionResponse) GetCreditBalanceBeforeOk() (*string, bool)`

GetCreditBalanceBeforeOk returns a tuple with the CreditBalanceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditBalanceBefore

`func (o *DtoWalletTransactionResponse) SetCreditBalanceBefore(v string)`

SetCreditBalanceBefore sets CreditBalanceBefore field to given value.

### HasCreditBalanceBefore

`func (o *DtoWalletTransactionResponse) HasCreditBalanceBefore() bool`

HasCreditBalanceBefore returns a boolean if a field has been set.

### GetCreditsAvailable

`func (o *DtoWalletTransactionResponse) GetCreditsAvailable() string`

GetCreditsAvailable returns the CreditsAvailable field if non-nil, zero value otherwise.

### GetCreditsAvailableOk

`func (o *DtoWalletTransactionResponse) GetCreditsAvailableOk() (*string, bool)`

GetCreditsAvailableOk returns a tuple with the CreditsAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsAvailable

`func (o *DtoWalletTransactionResponse) SetCreditsAvailable(v string)`

SetCreditsAvailable sets CreditsAvailable field to given value.

### HasCreditsAvailable

`func (o *DtoWalletTransactionResponse) HasCreditsAvailable() bool`

HasCreditsAvailable returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoWalletTransactionResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoWalletTransactionResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoWalletTransactionResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoWalletTransactionResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomer

`func (o *DtoWalletTransactionResponse) GetCustomer() DtoCustomerResponse`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *DtoWalletTransactionResponse) GetCustomerOk() (*DtoCustomerResponse, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *DtoWalletTransactionResponse) SetCustomer(v DtoCustomerResponse)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *DtoWalletTransactionResponse) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoWalletTransactionResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoWalletTransactionResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoWalletTransactionResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoWalletTransactionResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

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

### GetEnvironmentId

`func (o *DtoWalletTransactionResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoWalletTransactionResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoWalletTransactionResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoWalletTransactionResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

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

### GetIdempotencyKey

`func (o *DtoWalletTransactionResponse) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *DtoWalletTransactionResponse) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *DtoWalletTransactionResponse) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *DtoWalletTransactionResponse) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

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

### GetPriority

`func (o *DtoWalletTransactionResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtoWalletTransactionResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtoWalletTransactionResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtoWalletTransactionResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

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

### GetStatus

`func (o *DtoWalletTransactionResponse) GetStatus() TypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoWalletTransactionResponse) GetStatusOk() (*TypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoWalletTransactionResponse) SetStatus(v TypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoWalletTransactionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoWalletTransactionResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoWalletTransactionResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoWalletTransactionResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoWalletTransactionResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

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

`func (o *DtoWalletTransactionResponse) GetType() TypesTransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtoWalletTransactionResponse) GetTypeOk() (*TypesTransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtoWalletTransactionResponse) SetType(v TypesTransactionType)`

SetType sets Type field to given value.

### HasType

`func (o *DtoWalletTransactionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoWalletTransactionResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoWalletTransactionResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoWalletTransactionResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoWalletTransactionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoWalletTransactionResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoWalletTransactionResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoWalletTransactionResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoWalletTransactionResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetWallet

`func (o *DtoWalletTransactionResponse) GetWallet() DtoWalletResponse`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *DtoWalletTransactionResponse) GetWalletOk() (*DtoWalletResponse, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *DtoWalletTransactionResponse) SetWallet(v DtoWalletResponse)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *DtoWalletTransactionResponse) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

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


