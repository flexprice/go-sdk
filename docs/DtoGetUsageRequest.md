# DtoGetUsageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationType** | [**TypesAggregationType**](TypesAggregationType.md) |  | 
**BillingAnchor** | Pointer to **string** | BillingAnchor enables custom monthly billing periods for usage aggregation.  When to use: - WindowSize &#x3D; \&quot;MONTH\&quot; AND you need custom monthly periods (not calendar months) - Subscription billing that doesn&#39;t align with calendar months - Example: Customer signed up on 15th, so billing periods are 15th to 15th  When NOT to use: - WindowSize !&#x3D; \&quot;MONTH\&quot; (ignored for DAY, HOUR, WEEK, etc.) - Standard calendar-based billing (1st to 1st of each month)  Example values: - \&quot;2024-03-05T14:30:45.123456789Z\&quot; (5th of each month at 2:30:45 PM) - \&quot;2024-01-15T00:00:00Z\&quot; (15th of each month at midnight) - \&quot;2024-02-29T12:00:00Z\&quot; (29th of each month at noon - handles leap years) | [optional] 
**BucketSize** | Pointer to [**TypesWindowSize**](TypesWindowSize.md) |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**EventName** | **string** |  | 
**ExternalCustomerId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to **map[string][]string** |  | [optional] 
**Multiplier** | Pointer to **string** |  | [optional] 
**PropertyName** | Pointer to **string** | will be empty/ignored in case of COUNT | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**WindowSize** | Pointer to [**TypesWindowSize**](TypesWindowSize.md) |  | [optional] 

## Methods

### NewDtoGetUsageRequest

`func NewDtoGetUsageRequest(aggregationType TypesAggregationType, eventName string, ) *DtoGetUsageRequest`

NewDtoGetUsageRequest instantiates a new DtoGetUsageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetUsageRequestWithDefaults

`func NewDtoGetUsageRequestWithDefaults() *DtoGetUsageRequest`

NewDtoGetUsageRequestWithDefaults instantiates a new DtoGetUsageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationType

`func (o *DtoGetUsageRequest) GetAggregationType() TypesAggregationType`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *DtoGetUsageRequest) GetAggregationTypeOk() (*TypesAggregationType, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *DtoGetUsageRequest) SetAggregationType(v TypesAggregationType)`

SetAggregationType sets AggregationType field to given value.


### GetBillingAnchor

`func (o *DtoGetUsageRequest) GetBillingAnchor() string`

GetBillingAnchor returns the BillingAnchor field if non-nil, zero value otherwise.

### GetBillingAnchorOk

`func (o *DtoGetUsageRequest) GetBillingAnchorOk() (*string, bool)`

GetBillingAnchorOk returns a tuple with the BillingAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchor

`func (o *DtoGetUsageRequest) SetBillingAnchor(v string)`

SetBillingAnchor sets BillingAnchor field to given value.

### HasBillingAnchor

`func (o *DtoGetUsageRequest) HasBillingAnchor() bool`

HasBillingAnchor returns a boolean if a field has been set.

### GetBucketSize

`func (o *DtoGetUsageRequest) GetBucketSize() TypesWindowSize`

GetBucketSize returns the BucketSize field if non-nil, zero value otherwise.

### GetBucketSizeOk

`func (o *DtoGetUsageRequest) GetBucketSizeOk() (*TypesWindowSize, bool)`

GetBucketSizeOk returns a tuple with the BucketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSize

`func (o *DtoGetUsageRequest) SetBucketSize(v TypesWindowSize)`

SetBucketSize sets BucketSize field to given value.

### HasBucketSize

`func (o *DtoGetUsageRequest) HasBucketSize() bool`

HasBucketSize returns a boolean if a field has been set.

### GetCustomerId

`func (o *DtoGetUsageRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DtoGetUsageRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DtoGetUsageRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DtoGetUsageRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetEndTime

`func (o *DtoGetUsageRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DtoGetUsageRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DtoGetUsageRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DtoGetUsageRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEventName

`func (o *DtoGetUsageRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DtoGetUsageRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DtoGetUsageRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetExternalCustomerId

`func (o *DtoGetUsageRequest) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *DtoGetUsageRequest) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *DtoGetUsageRequest) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *DtoGetUsageRequest) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetFilters

`func (o *DtoGetUsageRequest) GetFilters() map[string][]string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DtoGetUsageRequest) GetFiltersOk() (*map[string][]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DtoGetUsageRequest) SetFilters(v map[string][]string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DtoGetUsageRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMultiplier

`func (o *DtoGetUsageRequest) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *DtoGetUsageRequest) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *DtoGetUsageRequest) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *DtoGetUsageRequest) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetPropertyName

`func (o *DtoGetUsageRequest) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *DtoGetUsageRequest) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *DtoGetUsageRequest) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *DtoGetUsageRequest) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetStartTime

`func (o *DtoGetUsageRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DtoGetUsageRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DtoGetUsageRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DtoGetUsageRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetWindowSize

`func (o *DtoGetUsageRequest) GetWindowSize() TypesWindowSize`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *DtoGetUsageRequest) GetWindowSizeOk() (*TypesWindowSize, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *DtoGetUsageRequest) SetWindowSize(v TypesWindowSize)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *DtoGetUsageRequest) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


