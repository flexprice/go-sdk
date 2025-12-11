# DtoTopUpWalletResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | Pointer to **string** | Invoice ID if an invoice was created (only for PURCHASED_CREDIT_INVOICED) | [optional] 
**Wallet** | Pointer to [**DtoWalletResponse**](DtoWalletResponse.md) |  | [optional] 
**WalletTransaction** | Pointer to [**DtoWalletTransactionResponse**](DtoWalletTransactionResponse.md) |  | [optional] 

## Methods

### NewDtoTopUpWalletResponse

`func NewDtoTopUpWalletResponse() *DtoTopUpWalletResponse`

NewDtoTopUpWalletResponse instantiates a new DtoTopUpWalletResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTopUpWalletResponseWithDefaults

`func NewDtoTopUpWalletResponseWithDefaults() *DtoTopUpWalletResponse`

NewDtoTopUpWalletResponseWithDefaults instantiates a new DtoTopUpWalletResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *DtoTopUpWalletResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *DtoTopUpWalletResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *DtoTopUpWalletResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *DtoTopUpWalletResponse) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetWallet

`func (o *DtoTopUpWalletResponse) GetWallet() DtoWalletResponse`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *DtoTopUpWalletResponse) GetWalletOk() (*DtoWalletResponse, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *DtoTopUpWalletResponse) SetWallet(v DtoWalletResponse)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *DtoTopUpWalletResponse) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetWalletTransaction

`func (o *DtoTopUpWalletResponse) GetWalletTransaction() DtoWalletTransactionResponse`

GetWalletTransaction returns the WalletTransaction field if non-nil, zero value otherwise.

### GetWalletTransactionOk

`func (o *DtoTopUpWalletResponse) GetWalletTransactionOk() (*DtoWalletTransactionResponse, bool)`

GetWalletTransactionOk returns a tuple with the WalletTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletTransaction

`func (o *DtoTopUpWalletResponse) SetWalletTransaction(v DtoWalletTransactionResponse)`

SetWalletTransaction sets WalletTransaction field to given value.

### HasWalletTransaction

`func (o *DtoTopUpWalletResponse) HasWalletTransaction() bool`

HasWalletTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


