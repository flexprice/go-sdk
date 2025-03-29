# DtoCreateMeterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | [**MeterAggregation**](MeterAggregation.md) |  | 
**EventName** | **string** |  | 
**Filters** | Pointer to [**[]MeterFilter**](MeterFilter.md) |  | [optional] 
**Name** | **string** |  | 
**ResetUsage** | [**TypesResetUsage**](TypesResetUsage.md) |  | 

## Methods

### NewDtoCreateMeterRequest

`func NewDtoCreateMeterRequest(aggregation MeterAggregation, eventName string, name string, resetUsage TypesResetUsage, ) *DtoCreateMeterRequest`

NewDtoCreateMeterRequest instantiates a new DtoCreateMeterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateMeterRequestWithDefaults

`func NewDtoCreateMeterRequestWithDefaults() *DtoCreateMeterRequest`

NewDtoCreateMeterRequestWithDefaults instantiates a new DtoCreateMeterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *DtoCreateMeterRequest) GetAggregation() MeterAggregation`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *DtoCreateMeterRequest) GetAggregationOk() (*MeterAggregation, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *DtoCreateMeterRequest) SetAggregation(v MeterAggregation)`

SetAggregation sets Aggregation field to given value.


### GetEventName

`func (o *DtoCreateMeterRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DtoCreateMeterRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DtoCreateMeterRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetFilters

`func (o *DtoCreateMeterRequest) GetFilters() []MeterFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DtoCreateMeterRequest) GetFiltersOk() (*[]MeterFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DtoCreateMeterRequest) SetFilters(v []MeterFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DtoCreateMeterRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetName

`func (o *DtoCreateMeterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtoCreateMeterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtoCreateMeterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetResetUsage

`func (o *DtoCreateMeterRequest) GetResetUsage() TypesResetUsage`

GetResetUsage returns the ResetUsage field if non-nil, zero value otherwise.

### GetResetUsageOk

`func (o *DtoCreateMeterRequest) GetResetUsageOk() (*TypesResetUsage, bool)`

GetResetUsageOk returns a tuple with the ResetUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetUsage

`func (o *DtoCreateMeterRequest) SetResetUsage(v TypesResetUsage)`

SetResetUsage sets ResetUsage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


