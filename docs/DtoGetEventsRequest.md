# DtoGetEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountTotal** | Pointer to **bool** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **string** |  | [optional] 
**EventName** | **string** |  | 
**ExternalCustomerId** | Pointer to **string** |  | [optional] 
**IterFirstKey** | Pointer to **string** |  | [optional] 
**IterLastKey** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PropertyFilters** | Pointer to **map[string][]string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoGetEventsRequest

`func NewDtoGetEventsRequest(eventName string, ) *DtoGetEventsRequest`

NewDtoGetEventsRequest instantiates a new DtoGetEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetEventsRequestWithDefaults

`func NewDtoGetEventsRequestWithDefaults() *DtoGetEventsRequest`

NewDtoGetEventsRequestWithDefaults instantiates a new DtoGetEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountTotal

`func (o *DtoGetEventsRequest) GetCountTotal() bool`

GetCountTotal returns the CountTotal field if non-nil, zero value otherwise.

### GetCountTotalOk

`func (o *DtoGetEventsRequest) GetCountTotalOk() (*bool, bool)`

GetCountTotalOk returns a tuple with the CountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountTotal

`func (o *DtoGetEventsRequest) SetCountTotal(v bool)`

SetCountTotal sets CountTotal field to given value.

### HasCountTotal

`func (o *DtoGetEventsRequest) HasCountTotal() bool`

HasCountTotal returns a boolean if a field has been set.

### GetEndTime

`func (o *DtoGetEventsRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DtoGetEventsRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DtoGetEventsRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DtoGetEventsRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEventId

`func (o *DtoGetEventsRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *DtoGetEventsRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *DtoGetEventsRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *DtoGetEventsRequest) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventName

`func (o *DtoGetEventsRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DtoGetEventsRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DtoGetEventsRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetExternalCustomerId

`func (o *DtoGetEventsRequest) GetExternalCustomerId() string`

GetExternalCustomerId returns the ExternalCustomerId field if non-nil, zero value otherwise.

### GetExternalCustomerIdOk

`func (o *DtoGetEventsRequest) GetExternalCustomerIdOk() (*string, bool)`

GetExternalCustomerIdOk returns a tuple with the ExternalCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCustomerId

`func (o *DtoGetEventsRequest) SetExternalCustomerId(v string)`

SetExternalCustomerId sets ExternalCustomerId field to given value.

### HasExternalCustomerId

`func (o *DtoGetEventsRequest) HasExternalCustomerId() bool`

HasExternalCustomerId returns a boolean if a field has been set.

### GetIterFirstKey

`func (o *DtoGetEventsRequest) GetIterFirstKey() string`

GetIterFirstKey returns the IterFirstKey field if non-nil, zero value otherwise.

### GetIterFirstKeyOk

`func (o *DtoGetEventsRequest) GetIterFirstKeyOk() (*string, bool)`

GetIterFirstKeyOk returns a tuple with the IterFirstKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterFirstKey

`func (o *DtoGetEventsRequest) SetIterFirstKey(v string)`

SetIterFirstKey sets IterFirstKey field to given value.

### HasIterFirstKey

`func (o *DtoGetEventsRequest) HasIterFirstKey() bool`

HasIterFirstKey returns a boolean if a field has been set.

### GetIterLastKey

`func (o *DtoGetEventsRequest) GetIterLastKey() string`

GetIterLastKey returns the IterLastKey field if non-nil, zero value otherwise.

### GetIterLastKeyOk

`func (o *DtoGetEventsRequest) GetIterLastKeyOk() (*string, bool)`

GetIterLastKeyOk returns a tuple with the IterLastKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterLastKey

`func (o *DtoGetEventsRequest) SetIterLastKey(v string)`

SetIterLastKey sets IterLastKey field to given value.

### HasIterLastKey

`func (o *DtoGetEventsRequest) HasIterLastKey() bool`

HasIterLastKey returns a boolean if a field has been set.

### GetOffset

`func (o *DtoGetEventsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DtoGetEventsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DtoGetEventsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DtoGetEventsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPageSize

`func (o *DtoGetEventsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DtoGetEventsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DtoGetEventsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DtoGetEventsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPropertyFilters

`func (o *DtoGetEventsRequest) GetPropertyFilters() map[string][]string`

GetPropertyFilters returns the PropertyFilters field if non-nil, zero value otherwise.

### GetPropertyFiltersOk

`func (o *DtoGetEventsRequest) GetPropertyFiltersOk() (*map[string][]string, bool)`

GetPropertyFiltersOk returns a tuple with the PropertyFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyFilters

`func (o *DtoGetEventsRequest) SetPropertyFilters(v map[string][]string)`

SetPropertyFilters sets PropertyFilters field to given value.

### HasPropertyFilters

`func (o *DtoGetEventsRequest) HasPropertyFilters() bool`

HasPropertyFilters returns a boolean if a field has been set.

### GetStartTime

`func (o *DtoGetEventsRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DtoGetEventsRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DtoGetEventsRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DtoGetEventsRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


