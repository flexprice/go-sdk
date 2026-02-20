# DtoGetEventByIDResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebugTracker** | Pointer to [**DtoDebugTracker**](DtoDebugTracker.md) |  | [optional] 
**Event** | Pointer to [**DtoEvent**](DtoEvent.md) |  | [optional] 
**ProcessedEvents** | Pointer to [**[]DtoFeatureUsageInfo**](DtoFeatureUsageInfo.md) |  | [optional] 
**Status** | Pointer to [**TypesEventProcessingStatusType**](TypesEventProcessingStatusType.md) |  | [optional] 

## Methods

### NewDtoGetEventByIDResponse

`func NewDtoGetEventByIDResponse() *DtoGetEventByIDResponse`

NewDtoGetEventByIDResponse instantiates a new DtoGetEventByIDResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetEventByIDResponseWithDefaults

`func NewDtoGetEventByIDResponseWithDefaults() *DtoGetEventByIDResponse`

NewDtoGetEventByIDResponseWithDefaults instantiates a new DtoGetEventByIDResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebugTracker

`func (o *DtoGetEventByIDResponse) GetDebugTracker() DtoDebugTracker`

GetDebugTracker returns the DebugTracker field if non-nil, zero value otherwise.

### GetDebugTrackerOk

`func (o *DtoGetEventByIDResponse) GetDebugTrackerOk() (*DtoDebugTracker, bool)`

GetDebugTrackerOk returns a tuple with the DebugTracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugTracker

`func (o *DtoGetEventByIDResponse) SetDebugTracker(v DtoDebugTracker)`

SetDebugTracker sets DebugTracker field to given value.

### HasDebugTracker

`func (o *DtoGetEventByIDResponse) HasDebugTracker() bool`

HasDebugTracker returns a boolean if a field has been set.

### GetEvent

`func (o *DtoGetEventByIDResponse) GetEvent() DtoEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *DtoGetEventByIDResponse) GetEventOk() (*DtoEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *DtoGetEventByIDResponse) SetEvent(v DtoEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *DtoGetEventByIDResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetProcessedEvents

`func (o *DtoGetEventByIDResponse) GetProcessedEvents() []DtoFeatureUsageInfo`

GetProcessedEvents returns the ProcessedEvents field if non-nil, zero value otherwise.

### GetProcessedEventsOk

`func (o *DtoGetEventByIDResponse) GetProcessedEventsOk() (*[]DtoFeatureUsageInfo, bool)`

GetProcessedEventsOk returns a tuple with the ProcessedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedEvents

`func (o *DtoGetEventByIDResponse) SetProcessedEvents(v []DtoFeatureUsageInfo)`

SetProcessedEvents sets ProcessedEvents field to given value.

### HasProcessedEvents

`func (o *DtoGetEventByIDResponse) HasProcessedEvents() bool`

HasProcessedEvents returns a boolean if a field has been set.

### GetStatus

`func (o *DtoGetEventByIDResponse) GetStatus() TypesEventProcessingStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoGetEventByIDResponse) GetStatusOk() (*TypesEventProcessingStatusType, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoGetEventByIDResponse) SetStatus(v TypesEventProcessingStatusType)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoGetEventByIDResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


