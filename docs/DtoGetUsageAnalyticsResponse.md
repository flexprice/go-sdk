# DtoGetUsageAnalyticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]DtoUsageAnalyticItem**](DtoUsageAnalyticItem.md) |  | [optional] 
**TotalCost** | Pointer to **float32** |  | [optional] 

## Methods

### NewDtoGetUsageAnalyticsResponse

`func NewDtoGetUsageAnalyticsResponse() *DtoGetUsageAnalyticsResponse`

NewDtoGetUsageAnalyticsResponse instantiates a new DtoGetUsageAnalyticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetUsageAnalyticsResponseWithDefaults

`func NewDtoGetUsageAnalyticsResponseWithDefaults() *DtoGetUsageAnalyticsResponse`

NewDtoGetUsageAnalyticsResponseWithDefaults instantiates a new DtoGetUsageAnalyticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *DtoGetUsageAnalyticsResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoGetUsageAnalyticsResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoGetUsageAnalyticsResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoGetUsageAnalyticsResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetItems

`func (o *DtoGetUsageAnalyticsResponse) GetItems() []DtoUsageAnalyticItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DtoGetUsageAnalyticsResponse) GetItemsOk() (*[]DtoUsageAnalyticItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DtoGetUsageAnalyticsResponse) SetItems(v []DtoUsageAnalyticItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *DtoGetUsageAnalyticsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalCost

`func (o *DtoGetUsageAnalyticsResponse) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *DtoGetUsageAnalyticsResponse) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *DtoGetUsageAnalyticsResponse) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *DtoGetUsageAnalyticsResponse) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


