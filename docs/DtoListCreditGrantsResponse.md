# DtoListCreditGrantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DtoCreditGrantResponse**](DtoCreditGrantResponse.md) |  | [optional] 
**Pagination** | Pointer to [**TypesPaginationResponse**](TypesPaginationResponse.md) |  | [optional] 

## Methods

### NewDtoListCreditGrantsResponse

`func NewDtoListCreditGrantsResponse() *DtoListCreditGrantsResponse`

NewDtoListCreditGrantsResponse instantiates a new DtoListCreditGrantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoListCreditGrantsResponseWithDefaults

`func NewDtoListCreditGrantsResponseWithDefaults() *DtoListCreditGrantsResponse`

NewDtoListCreditGrantsResponseWithDefaults instantiates a new DtoListCreditGrantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DtoListCreditGrantsResponse) GetItems() []DtoCreditGrantResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DtoListCreditGrantsResponse) GetItemsOk() (*[]DtoCreditGrantResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DtoListCreditGrantsResponse) SetItems(v []DtoCreditGrantResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *DtoListCreditGrantsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPagination

`func (o *DtoListCreditGrantsResponse) GetPagination() TypesPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DtoListCreditGrantsResponse) GetPaginationOk() (*TypesPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DtoListCreditGrantsResponse) SetPagination(v TypesPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DtoListCreditGrantsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


