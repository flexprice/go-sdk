# DtoGetPendingSchedulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | count is the number of pending schedules | [optional] 
**Schedules** | Pointer to [**[]DtoSubscriptionScheduleResponse**](DtoSubscriptionScheduleResponse.md) | schedules is the list of pending schedules | [optional] 

## Methods

### NewDtoGetPendingSchedulesResponse

`func NewDtoGetPendingSchedulesResponse() *DtoGetPendingSchedulesResponse`

NewDtoGetPendingSchedulesResponse instantiates a new DtoGetPendingSchedulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoGetPendingSchedulesResponseWithDefaults

`func NewDtoGetPendingSchedulesResponseWithDefaults() *DtoGetPendingSchedulesResponse`

NewDtoGetPendingSchedulesResponseWithDefaults instantiates a new DtoGetPendingSchedulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DtoGetPendingSchedulesResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DtoGetPendingSchedulesResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DtoGetPendingSchedulesResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DtoGetPendingSchedulesResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSchedules

`func (o *DtoGetPendingSchedulesResponse) GetSchedules() []DtoSubscriptionScheduleResponse`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *DtoGetPendingSchedulesResponse) GetSchedulesOk() (*[]DtoSubscriptionScheduleResponse, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *DtoGetPendingSchedulesResponse) SetSchedules(v []DtoSubscriptionScheduleResponse)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *DtoGetPendingSchedulesResponse) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


