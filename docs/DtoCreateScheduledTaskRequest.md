# DtoCreateScheduledTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**EntityType** | [**TypesScheduledTaskEntityType**](TypesScheduledTaskEntityType.md) |  | 
**Interval** | [**TypesScheduledTaskInterval**](TypesScheduledTaskInterval.md) |  | 
**JobConfig** | [**TypesS3JobConfig**](TypesS3JobConfig.md) |  | 

## Methods

### NewDtoCreateScheduledTaskRequest

`func NewDtoCreateScheduledTaskRequest(connectionId string, entityType TypesScheduledTaskEntityType, interval TypesScheduledTaskInterval, jobConfig TypesS3JobConfig, ) *DtoCreateScheduledTaskRequest`

NewDtoCreateScheduledTaskRequest instantiates a new DtoCreateScheduledTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateScheduledTaskRequestWithDefaults

`func NewDtoCreateScheduledTaskRequestWithDefaults() *DtoCreateScheduledTaskRequest`

NewDtoCreateScheduledTaskRequestWithDefaults instantiates a new DtoCreateScheduledTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *DtoCreateScheduledTaskRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DtoCreateScheduledTaskRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DtoCreateScheduledTaskRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetEnabled

`func (o *DtoCreateScheduledTaskRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DtoCreateScheduledTaskRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DtoCreateScheduledTaskRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DtoCreateScheduledTaskRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoCreateScheduledTaskRequest) GetEntityType() TypesScheduledTaskEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoCreateScheduledTaskRequest) GetEntityTypeOk() (*TypesScheduledTaskEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoCreateScheduledTaskRequest) SetEntityType(v TypesScheduledTaskEntityType)`

SetEntityType sets EntityType field to given value.


### GetInterval

`func (o *DtoCreateScheduledTaskRequest) GetInterval() TypesScheduledTaskInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DtoCreateScheduledTaskRequest) GetIntervalOk() (*TypesScheduledTaskInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DtoCreateScheduledTaskRequest) SetInterval(v TypesScheduledTaskInterval)`

SetInterval sets Interval field to given value.


### GetJobConfig

`func (o *DtoCreateScheduledTaskRequest) GetJobConfig() TypesS3JobConfig`

GetJobConfig returns the JobConfig field if non-nil, zero value otherwise.

### GetJobConfigOk

`func (o *DtoCreateScheduledTaskRequest) GetJobConfigOk() (*TypesS3JobConfig, bool)`

GetJobConfigOk returns a tuple with the JobConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobConfig

`func (o *DtoCreateScheduledTaskRequest) SetJobConfig(v TypesS3JobConfig)`

SetJobConfig sets JobConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


