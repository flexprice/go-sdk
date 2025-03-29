# DtoListTasksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DtoTaskResponse**](DtoTaskResponse.md) |  | [optional] 
**Pagination** | Pointer to [**TypesPaginationResponse**](TypesPaginationResponse.md) |  | [optional] 

## Methods

### NewDtoListTasksResponse

`func NewDtoListTasksResponse() *DtoListTasksResponse`

NewDtoListTasksResponse instantiates a new DtoListTasksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoListTasksResponseWithDefaults

`func NewDtoListTasksResponseWithDefaults() *DtoListTasksResponse`

NewDtoListTasksResponseWithDefaults instantiates a new DtoListTasksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DtoListTasksResponse) GetItems() []DtoTaskResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DtoListTasksResponse) GetItemsOk() (*[]DtoTaskResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DtoListTasksResponse) SetItems(v []DtoTaskResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *DtoListTasksResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPagination

`func (o *DtoListTasksResponse) GetPagination() TypesPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DtoListTasksResponse) GetPaginationOk() (*TypesPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DtoListTasksResponse) SetPagination(v TypesPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DtoListTasksResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


