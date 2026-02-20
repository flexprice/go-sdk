# DtoSubscriptionScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanBeCancelled** | Pointer to **bool** | can_be_cancelled indicates if the schedule can be cancelled | [optional] 
**CancelledAt** | Pointer to **string** | cancelled_at is when the schedule was cancelled | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | configuration contains type-specific configuration (e.g., target_plan_id for plan changes) | [optional] 
**CreatedAt** | Pointer to **string** | created_at timestamp | [optional] 
**DaysUntilExecution** | Pointer to **int32** | days_until_execution is the number of days until execution | [optional] 
**ErrorMessage** | Pointer to **string** | error_message contains the error if execution failed | [optional] 
**ExecutedAt** | Pointer to **string** | executed_at is when the schedule was executed | [optional] 
**ExecutionResult** | Pointer to **map[string]interface{}** | execution_result contains type-specific execution result | [optional] 
**Id** | Pointer to **string** | id of the schedule | [optional] 
**Metadata** | Pointer to **map[string]string** | metadata from the schedule | [optional] 
**ScheduleType** | Pointer to [**TypesSubscriptionScheduleChangeType**](TypesSubscriptionScheduleChangeType.md) |  | [optional] 
**ScheduledAt** | Pointer to **string** | scheduled_at is when the schedule will execute | [optional] 
**Status** | Pointer to [**TypesScheduleStatus**](TypesScheduleStatus.md) |  | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id is the ID of the subscription | [optional] 
**UpdatedAt** | Pointer to **string** | updated_at timestamp | [optional] 

## Methods

### NewDtoSubscriptionScheduleResponse

`func NewDtoSubscriptionScheduleResponse() *DtoSubscriptionScheduleResponse`

NewDtoSubscriptionScheduleResponse instantiates a new DtoSubscriptionScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoSubscriptionScheduleResponseWithDefaults

`func NewDtoSubscriptionScheduleResponseWithDefaults() *DtoSubscriptionScheduleResponse`

NewDtoSubscriptionScheduleResponseWithDefaults instantiates a new DtoSubscriptionScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanBeCancelled

`func (o *DtoSubscriptionScheduleResponse) GetCanBeCancelled() bool`

GetCanBeCancelled returns the CanBeCancelled field if non-nil, zero value otherwise.

### GetCanBeCancelledOk

`func (o *DtoSubscriptionScheduleResponse) GetCanBeCancelledOk() (*bool, bool)`

GetCanBeCancelledOk returns a tuple with the CanBeCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeCancelled

`func (o *DtoSubscriptionScheduleResponse) SetCanBeCancelled(v bool)`

SetCanBeCancelled sets CanBeCancelled field to given value.

### HasCanBeCancelled

`func (o *DtoSubscriptionScheduleResponse) HasCanBeCancelled() bool`

HasCanBeCancelled returns a boolean if a field has been set.

### GetCancelledAt

`func (o *DtoSubscriptionScheduleResponse) GetCancelledAt() string`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *DtoSubscriptionScheduleResponse) GetCancelledAtOk() (*string, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *DtoSubscriptionScheduleResponse) SetCancelledAt(v string)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *DtoSubscriptionScheduleResponse) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### GetConfiguration

`func (o *DtoSubscriptionScheduleResponse) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DtoSubscriptionScheduleResponse) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DtoSubscriptionScheduleResponse) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DtoSubscriptionScheduleResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoSubscriptionScheduleResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoSubscriptionScheduleResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoSubscriptionScheduleResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoSubscriptionScheduleResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDaysUntilExecution

`func (o *DtoSubscriptionScheduleResponse) GetDaysUntilExecution() int32`

GetDaysUntilExecution returns the DaysUntilExecution field if non-nil, zero value otherwise.

### GetDaysUntilExecutionOk

`func (o *DtoSubscriptionScheduleResponse) GetDaysUntilExecutionOk() (*int32, bool)`

GetDaysUntilExecutionOk returns a tuple with the DaysUntilExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUntilExecution

`func (o *DtoSubscriptionScheduleResponse) SetDaysUntilExecution(v int32)`

SetDaysUntilExecution sets DaysUntilExecution field to given value.

### HasDaysUntilExecution

`func (o *DtoSubscriptionScheduleResponse) HasDaysUntilExecution() bool`

HasDaysUntilExecution returns a boolean if a field has been set.

### GetErrorMessage

`func (o *DtoSubscriptionScheduleResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *DtoSubscriptionScheduleResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *DtoSubscriptionScheduleResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *DtoSubscriptionScheduleResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetExecutedAt

`func (o *DtoSubscriptionScheduleResponse) GetExecutedAt() string`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *DtoSubscriptionScheduleResponse) GetExecutedAtOk() (*string, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *DtoSubscriptionScheduleResponse) SetExecutedAt(v string)`

SetExecutedAt sets ExecutedAt field to given value.

### HasExecutedAt

`func (o *DtoSubscriptionScheduleResponse) HasExecutedAt() bool`

HasExecutedAt returns a boolean if a field has been set.

### GetExecutionResult

`func (o *DtoSubscriptionScheduleResponse) GetExecutionResult() map[string]interface{}`

GetExecutionResult returns the ExecutionResult field if non-nil, zero value otherwise.

### GetExecutionResultOk

`func (o *DtoSubscriptionScheduleResponse) GetExecutionResultOk() (*map[string]interface{}, bool)`

GetExecutionResultOk returns a tuple with the ExecutionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionResult

`func (o *DtoSubscriptionScheduleResponse) SetExecutionResult(v map[string]interface{})`

SetExecutionResult sets ExecutionResult field to given value.

### HasExecutionResult

`func (o *DtoSubscriptionScheduleResponse) HasExecutionResult() bool`

HasExecutionResult returns a boolean if a field has been set.

### GetId

`func (o *DtoSubscriptionScheduleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoSubscriptionScheduleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoSubscriptionScheduleResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoSubscriptionScheduleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoSubscriptionScheduleResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoSubscriptionScheduleResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoSubscriptionScheduleResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoSubscriptionScheduleResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetScheduleType

`func (o *DtoSubscriptionScheduleResponse) GetScheduleType() TypesSubscriptionScheduleChangeType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *DtoSubscriptionScheduleResponse) GetScheduleTypeOk() (*TypesSubscriptionScheduleChangeType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *DtoSubscriptionScheduleResponse) SetScheduleType(v TypesSubscriptionScheduleChangeType)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *DtoSubscriptionScheduleResponse) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduledAt

`func (o *DtoSubscriptionScheduleResponse) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *DtoSubscriptionScheduleResponse) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *DtoSubscriptionScheduleResponse) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *DtoSubscriptionScheduleResponse) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetStatus

`func (o *DtoSubscriptionScheduleResponse) GetStatus() TypesScheduleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoSubscriptionScheduleResponse) GetStatusOk() (*TypesScheduleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoSubscriptionScheduleResponse) SetStatus(v TypesScheduleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoSubscriptionScheduleResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DtoSubscriptionScheduleResponse) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DtoSubscriptionScheduleResponse) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DtoSubscriptionScheduleResponse) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DtoSubscriptionScheduleResponse) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoSubscriptionScheduleResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoSubscriptionScheduleResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoSubscriptionScheduleResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoSubscriptionScheduleResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


