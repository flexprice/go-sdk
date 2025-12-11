# DtoGetUsageByMeterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAnchor** | Pointer to **string** | BillingAnchor enables custom monthly billing periods for meter usage aggregation.  Usage guidelines: - Only effective when WindowSize &#x3D; \&quot;MONTH\&quot; - For other window sizes (DAY, HOUR, WEEK), this field is ignored - When nil, uses standard calendar months (1st to 1st) - When provided, creates custom monthly periods (e.g., 5th to 5th)  Common use cases: - Subscription billing periods that don&#39;t align with calendar months - Customer-specific billing cycles (e.g., signed up on 15th) - Multi-tenant systems with different billing anchor dates  Example: If BillingAnchor &#x3D; \&quot;2024-03-05T14:30:45Z\&quot; and WindowSize &#x3D; \&quot;MONTH\&quot;:   - March period: 2024-03-05 14:30:45 to 2024-04-05 14:30:45   - April period: 2024-04-05 14:30:45 to 2024-05-05 14:30:45 | [optional] 
**BucketSize** | Pointer to [**TypesWindowSize**](TypesWindowSize.md) |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**ExternalCustomerId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **map[string][]string** |  | [optional] 
**MeterId** | **string** |  | 
**StartTime** | Pointer to **string** |  | [optional] 
**WindowSize** | Pointer to [**TypesWindowSize**](TypesWindowSize.md) |  | [optional] 

## Methods

### NewDtoGetUsageByMeterRequest

`func NewDtoGetUsageByMeterRequest(meterId string, ) *DtoGetUsageByMeterRequest`

NewDtoGetUsageByMeterRequest instantiates a new DtoGetUsageByMeterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetUsageByMeterRequestWithDefaults

`func NewDtoGetUsageByMeterRequestWithDefaults() *DtoGetUsageByMeterRequest`

NewDtoGetUsageByMeterRequestWithDefaults instantiates a new DtoGetUsageByMeterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingAnchor

`func (o *DtoGetUsageByMeterRequest) GetBillingAnchor() string`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *DtoGetUsageByMeterRequest) GetBillingAnchorOk() (*string, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *DtoGetUsageByMeterRequest) SetBillingAnchor(v string)`

SetBillingAnchor sets BillingAnchor field to given value.

### HasBillingAnchor

`func (o *DtoGetUsageByMeterRequest) HasBillingAnchor() bool`

HasBillingAnchor returns a boolean if a field has been set.

### GetBucketSize

`func (o *DtoGetUsageByMeterRequest) GetBucketSize() TypesWindowSize`

GetBucketSize returns the BucketSize field if non-nil, zero value otherwise.

### GetBucketSizeOk

`func (o *DtoGetUsageByMeterRequest) GetBucketSizeOk() (*TypesWindowSize, bool)`

GetBucketSizeOk returns a tuple with the BucketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSize

`func (o *DtoGetUsageByMeterRequest) SetBucketSize(v TypesWindowSize)`

SetBucketSize sets BucketSize field to given value.

### HasBucketSize

`func (o *DtoGetUsageByMeterRequest) HasBucketSize() bool`

HasBucketSize returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoGetUsageByMeterRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoGetUsageByMeterRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoGetUsageByMeterRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoGetUsageByMeterRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEndTime

`func (o *DtoGetUsageByMeterRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DtoGetUsageByMeterRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DtoGetUsageByMeterRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DtoGetUsageByMeterRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExternalCustomerId

`func (o *DtoGetUsageByMeterRequest) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *DtoGetUsageByMeterRequest) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *DtoGetUsageByMeterRequest) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *DtoGetUsageByMeterRequest) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetFilters

`func (o *DtoGetUsageByMeterRequest) GetFilters() map[string][]string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DtoGetUsageByMeterRequest) GetFiltersOk() (*map[string][]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DtoGetUsageByMeterRequest) SetFilters(v map[string][]string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DtoGetUsageByMeterRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMeterId

`func (o *DtoGetUsageByMeterRequest) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *DtoGetUsageByMeterRequest) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *DtoGetUsageByMeterRequest) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.


### GetStartTime

`func (o *DtoGetUsageByMeterRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DtoGetUsageByMeterRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DtoGetUsageByMeterRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DtoGetUsageByMeterRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetWindowSize

`func (o *DtoGetUsageByMeterRequest) GetWindowSize() TypesWindowSize`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *DtoGetUsageByMeterRequest) GetWindowSizeOk() (*TypesWindowSize, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *DtoGetUsageByMeterRequest) SetWindowSize(v TypesWindowSize)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *DtoGetUsageByMeterRequest) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


