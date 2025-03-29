# DtoListWalletTransactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DtoWalletTransactionResponse**](DtoWalletTransactionResponse.md) |  | [optional] 
**Pagination** | Pointer to [**TypesPaginationResponse**](TypesPaginationResponse.md) |  | [optional] 

## Methods

### NewDtoListWalletTransactionsResponse

`func NewDtoListWalletTransactionsResponse() *DtoListWalletTransactionsResponse`

NewDtoListWalletTransactionsResponse instantiates a new DtoListWalletTransactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoListWalletTransactionsResponseWithDefaults

`func NewDtoListWalletTransactionsResponseWithDefaults() *DtoListWalletTransactionsResponse`

NewDtoListWalletTransactionsResponseWithDefaults instantiates a new DtoListWalletTransactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DtoListWalletTransactionsResponse) GetItems() []DtoWalletTransactionResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DtoListWalletTransactionsResponse) GetItemsOk() (*[]DtoWalletTransactionResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DtoListWalletTransactionsResponse) SetItems(v []DtoWalletTransactionResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *DtoListWalletTransactionsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPagination

`func (o *DtoListWalletTransactionsResponse) GetPagination() TypesPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DtoListWalletTransactionsResponse) GetPaginationOk() (*TypesPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DtoListWalletTransactionsResponse) SetPagination(v TypesPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DtoListWalletTransactionsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


