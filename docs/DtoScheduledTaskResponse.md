# DtoScheduledTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EntityType** | Pointer to [**TypesScheduledTaskEntityType**](TypesScheduledTaskEntityType.md) |  | [optional] 
**EnvironmentId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Interval** | Pointer to [**TypesScheduledTaskInterval**](TypesScheduledTaskInterval.md) |  | [optional] 
**JobConfig** | Pointer to [**TypesS3JobConfig**](TypesS3JobConfig.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoScheduledTaskResponse

`func NewDtoScheduledTaskResponse() *DtoScheduledTaskResponse`

NewDtoScheduledTaskResponse instantiates a new DtoScheduledTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoScheduledTaskResponseWithDefaults

`func NewDtoScheduledTaskResponseWithDefaults() *DtoScheduledTaskResponse`

NewDtoScheduledTaskResponseWithDefaults instantiates a new DtoScheduledTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *DtoScheduledTaskResponse) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DtoScheduledTaskResponse) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DtoScheduledTaskResponse) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *DtoScheduledTaskResponse) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoScheduledTaskResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoScheduledTaskResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoScheduledTaskResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoScheduledTaskResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *DtoScheduledTaskResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DtoScheduledTaskResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DtoScheduledTaskResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DtoScheduledTaskResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoScheduledTaskResponse) GetEntityType() TypesScheduledTaskEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoScheduledTaskResponse) GetEntityTypeOk() (*TypesScheduledTaskEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoScheduledTaskResponse) SetEntityType(v TypesScheduledTaskEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DtoScheduledTaskResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DtoScheduledTaskResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DtoScheduledTaskResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DtoScheduledTaskResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DtoScheduledTaskResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *DtoScheduledTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoScheduledTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoScheduledTaskResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoScheduledTaskResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *DtoScheduledTaskResponse) GetInterval() TypesScheduledTaskInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DtoScheduledTaskResponse) GetIntervalOk() (*TypesScheduledTaskInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DtoScheduledTaskResponse) SetInterval(v TypesScheduledTaskInterval)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DtoScheduledTaskResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetJobConfig

`func (o *DtoScheduledTaskResponse) GetJobConfig() TypesS3JobConfig`

GetJobConfig returns the JobConfig field if non-nil, zero value otherwise.

### GetJobConfigOk

`func (o *DtoScheduledTaskResponse) GetJobConfigOk() (*TypesS3JobConfig, bool)`

GetJobConfigOk returns a tuple with the JobConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobConfig

`func (o *DtoScheduledTaskResponse) SetJobConfig(v TypesS3JobConfig)`

SetJobConfig sets JobConfig field to given value.

### HasJobConfig

`func (o *DtoScheduledTaskResponse) HasJobConfig() bool`

HasJobConfig returns a boolean if a field has been set.

### GetStatus

`func (o *DtoScheduledTaskResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoScheduledTaskResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoScheduledTaskResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoScheduledTaskResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoScheduledTaskResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoScheduledTaskResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoScheduledTaskResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoScheduledTaskResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoScheduledTaskResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoScheduledTaskResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoScheduledTaskResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoScheduledTaskResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


