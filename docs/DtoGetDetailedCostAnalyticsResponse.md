# DtoGetDetailedCostAnalyticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostAnalytics** | Pointer to [**[]DtoCostAnalyticItem**](DtoCostAnalyticItem.md) | Cost analytics array (flattened from nested structure) | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Margin** | Pointer to **string** | Revenue - Cost | [optional] 
**MarginPercent** | Pointer to **string** | (Margin / Revenue) * 100 | [optional] 
**Roi** | Pointer to **string** | (Revenue - Cost) / Cost | [optional] 
**RoiPercent** | Pointer to **string** | ROI * 100 | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**TotalCost** | Pointer to **string** |  | [optional] 
**TotalRevenue** | Pointer to **string** | Derived metrics | [optional] 

## Methods

### NewDtoGetDetailedCostAnalyticsResponse

`func NewDtoGetDetailedCostAnalyticsResponse() *DtoGetDetailedCostAnalyticsResponse`

NewDtoGetDetailedCostAnalyticsResponse instantiates a new DtoGetDetailedCostAnalyticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetDetailedCostAnalyticsResponseWithDefaults

`func NewDtoGetDetailedCostAnalyticsResponseWithDefaults() *DtoGetDetailedCostAnalyticsResponse`

NewDtoGetDetailedCostAnalyticsResponseWithDefaults instantiates a new DtoGetDetailedCostAnalyticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostAnalytics

`func (o *DtoGetDetailedCostAnalyticsResponse) GetCostAnalytics() []DtoCostAnalyticItem`

GetCostAnalytics returns the CostAnalytics field if non-nil, zero value otherwise.

### GetCostAnalyticsOk

`func (o *DtoGetDetailedCostAnalyticsResponse) GetCostAnalyticsOk() (*[]DtoCostAnalyticItem, bool)`

GetCostAnalyticsOk returns a tuple with the CostAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAnalytics

`func (o *DtoGetDetailedCostAnalyticsResponse) SetCostAnalytics(v []DtoCostAnalyticItem)`

SetCostAnalytics sets CostAnalytics field to given value.

### HasCostAnalytics

`func (o *DtoGetDetailedCostAnalyticsResponse) HasCostAnalytics() bool`

HasCostAnalytics returns a boolean if a field has been set.

### GetCurrency

`func (o *DtoGetDetailedCostAnalyticsResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DtoGetDetailedCostAnalyticsResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DtoGetDetailedCostAnalyticsResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DtoGetDetailedCostAnalyticsResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEndTime

`func (o *DtoGetDetailedCostAnalyticsResponse) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DtoGetDetailedCostAnalyticsResponse) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DtoGetDetailedCostAnalyticsResponse) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DtoGetDetailedCostAnalyticsResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMargin

`func (o *DtoGetDetailedCostAnalyticsResponse) GetMargin() string`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *DtoGetDetailedCostAnalyticsResponse) GetMarginOk() (*string, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *DtoGetDetailedCostAnalyticsResponse) SetMargin(v string)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *DtoGetDetailedCostAnalyticsResponse) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetMarginPercent

`func (o *DtoGetDetailedCostAnalyticsResponse) GetMarginPercent() string`

GetMarginPercent returns the MarginPercent field if non-nil, zero value otherwise.

### GetMarginPercentOk

`func (o *DtoGetDetailedCostAnalyticsResponse) GetMarginPercentOk() (*string, bool)`

GetMarginPercentOk returns a tuple with the MarginPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginPercent

`func (o *DtoGetDetailedCostAnalyticsResponse) SetMarginPercent(v string)`

SetMarginPercent sets MarginPercent field to given value.

### HasMarginPercent

`func (o *DtoGetDetailedCostAnalyticsResponse) HasMarginPercent() bool`

HasMarginPercent returns a boolean if a field has been set.

### GetRoi

`func (o *DtoGetDetailedCostAnalyticsResponse) GetRoi() string`

GetRoi returns the Roi field if non-nil, zero value otherwise.

### GetRoiOk

`func (o *DtoGetDetailedCostAnalyticsResponse) GetRoiOk() (*string, bool)`

GetRoiOk returns a tuple with the Roi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoi

`func (o *DtoGetDetailedCostAnalyticsResponse) SetRoi(v string)`

SetRoi sets Roi field to given value.

### HasRoi

`func (o *DtoGetDetailedCostAnalyticsResponse) HasRoi() bool`

HasRoi returns a boolean if a field has been set.

### GetRoiPercent

`func (o *DtoGetDetailedCostAnalyticsResponse) GetRoiPercent() string`

GetRoiPercent returns the RoiPercent field if non-nil, zero value otherwise.

### GetRoiPercentOk

`func (o *DtoGetDetailedCostAnalyticsResponse) GetRoiPercentOk() (*string, bool)`

GetRoiPercentOk returns a tuple with the RoiPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoiPercent

`func (o *DtoGetDetailedCostAnalyticsResponse) SetRoiPercent(v string)`

SetRoiPercent sets RoiPercent field to given value.

### HasRoiPercent

`func (o *DtoGetDetailedCostAnalyticsResponse) HasRoiPercent() bool`

HasRoiPercent returns a boolean if a field has been set.

### GetStartTime

`func (o *DtoGetDetailedCostAnalyticsResponse) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DtoGetDetailedCostAnalyticsResponse) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DtoGetDetailedCostAnalyticsResponse) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DtoGetDetailedCostAnalyticsResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTotalCost

`func (o *DtoGetDetailedCostAnalyticsResponse) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *DtoGetDetailedCostAnalyticsResponse) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *DtoGetDetailedCostAnalyticsResponse) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *DtoGetDetailedCostAnalyticsResponse) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalRevenue

`func (o *DtoGetDetailedCostAnalyticsResponse) GetTotalRevenue() string`

GetTotalRevenue returns the TotalRevenue field if non-nil, zero value otherwise.

### GetTotalRevenueOk

`func (o *DtoGetDetailedCostAnalyticsResponse) GetTotalRevenueOk() (*string, bool)`

GetTotalRevenueOk returns a tuple with the TotalRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRevenue

`func (o *DtoGetDetailedCostAnalyticsResponse) SetTotalRevenue(v string)`

SetTotalRevenue sets TotalRevenue field to given value.

### HasTotalRevenue

`func (o *DtoGetDetailedCostAnalyticsResponse) HasTotalRevenue() bool`

HasTotalRevenue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


