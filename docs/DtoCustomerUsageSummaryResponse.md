# DtoCustomerUsageSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | Pointer to **string** |  | [optional] 
**Features** | Pointer to [**[]DtoFeatureUsageSummary**](DtoFeatureUsageSummary.md) |  | [optional] 
**Pagination** | Pointer to [**TypesPaginationResponse**](TypesPaginationResponse.md) |  | [optional] 
**Period** | Pointer to [**DtoBillingPeriodInfo**](DtoBillingPeriodInfo.md) |  | [optional] 

## Methods

### NewDtoCustomerUsageSummaryResponse

`func NewDtoCustomerUsageSummaryResponse() *DtoCustomerUsageSummaryResponse`

NewDtoCustomerUsageSummaryResponse instantiates a new DtoCustomerUsageSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCustomerUsageSummaryResponseWithDefaults

`func NewDtoCustomerUsageSummaryResponseWithDefaults() *DtoCustomerUsageSummaryResponse`

NewDtoCustomerUsageSummaryResponseWithDefaults instantiates a new DtoCustomerUsageSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *DtoCustomerUsageSummaryResponse) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoCustomerUsageSummaryResponse) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoCustomerUsageSummaryResponse) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoCustomerUsageSummaryResponse) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetFeatures

`func (o *DtoCustomerUsageSummaryResponse) GetFeatures() []DtoFeatureUsageSummary`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DtoCustomerUsageSummaryResponse) GetFeaturesOk() (*[]DtoFeatureUsageSummary, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DtoCustomerUsageSummaryResponse) SetFeatures(v []DtoFeatureUsageSummary)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DtoCustomerUsageSummaryResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetPagination

`func (o *DtoCustomerUsageSummaryResponse) GetPagination() TypesPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DtoCustomerUsageSummaryResponse) GetPaginationOk() (*TypesPaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DtoCustomerUsageSummaryResponse) SetPagination(v TypesPaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DtoCustomerUsageSummaryResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetPeriod

`func (o *DtoCustomerUsageSummaryResponse) GetPeriod() DtoBillingPeriodInfo`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DtoCustomerUsageSummaryResponse) GetPeriodOk() (*DtoBillingPeriodInfo, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DtoCustomerUsageSummaryResponse) SetPeriod(v DtoBillingPeriodInfo)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DtoCustomerUsageSummaryResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


