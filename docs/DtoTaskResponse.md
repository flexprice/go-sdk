# DtoTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to [**TypesEntityType**](TypesEntityType.md) |  | [optional] 
**ErrorSummary** | Pointer to **string** |  | [optional] 
**FailedAt** | Pointer to **string** |  | [optional] 
**FailedRecords** | Pointer to **int32** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**FileType** | Pointer to [**TypesFileType**](TypesFileType.md) |  | [optional] 
**FileUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ProcessedRecords** | Pointer to **int32** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SuccessfulRecords** | Pointer to **int32** |  | [optional] 
**TaskStatus** | Pointer to [**TypesTaskStatus**](TypesTaskStatus.md) |  | [optional] 
**TaskType** | Pointer to [**TypesTaskType**](TypesTaskType.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**TotalRecords** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewDtoTaskResponse

`func NewDtoTaskResponse() *DtoTaskResponse`

NewDtoTaskResponse instantiates a new DtoTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoTaskResponseWithDefaults

`func NewDtoTaskResponseWithDefaults() *DtoTaskResponse`

NewDtoTaskResponseWithDefaults instantiates a new DtoTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *DtoTaskResponse) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *DtoTaskResponse) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *DtoTaskResponse) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *DtoTaskResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DtoTaskResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DtoTaskResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DtoTaskResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DtoTaskResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *DtoTaskResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DtoTaskResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DtoTaskResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DtoTaskResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEntityType

`func (o *DtoTaskResponse) GetEntityType() TypesEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoTaskResponse) GetEntityTypeOk() (*TypesEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoTaskResponse) SetEntityType(v TypesEntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DtoTaskResponse) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetErrorSummary

`func (o *DtoTaskResponse) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *DtoTaskResponse) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *DtoTaskResponse) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *DtoTaskResponse) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.

### GetFailedAt

`func (o *DtoTaskResponse) GetFailedAt() string`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *DtoTaskResponse) GetFailedAtOk() (*string, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *DtoTaskResponse) SetFailedAt(v string)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *DtoTaskResponse) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetFailedRecords

`func (o *DtoTaskResponse) GetFailedRecords() int32`

GetFailedRecords returns the FailedRecords field if non-nil, zero value otherwise.

### GetFailedRecordsOk

`func (o *DtoTaskResponse) GetFailedRecordsOk() (*int32, bool)`

GetFailedRecordsOk returns a tuple with the FailedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRecords

`func (o *DtoTaskResponse) SetFailedRecords(v int32)`

SetFailedRecords sets FailedRecords field to given value.

### HasFailedRecords

`func (o *DtoTaskResponse) HasFailedRecords() bool`

HasFailedRecords returns a boolean if a field has been set.

### GetFileName

`func (o *DtoTaskResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DtoTaskResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DtoTaskResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DtoTaskResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileType

`func (o *DtoTaskResponse) GetFileType() TypesFileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *DtoTaskResponse) GetFileTypeOk() (*TypesFileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *DtoTaskResponse) SetFileType(v TypesFileType)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *DtoTaskResponse) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFileUrl

`func (o *DtoTaskResponse) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *DtoTaskResponse) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *DtoTaskResponse) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *DtoTaskResponse) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetId

`func (o *DtoTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DtoTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DtoTaskResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DtoTaskResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DtoTaskResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoTaskResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoTaskResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoTaskResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProcessedRecords

`func (o *DtoTaskResponse) GetProcessedRecords() int32`

GetProcessedRecords returns the ProcessedRecords field if non-nil, zero value otherwise.

### GetProcessedRecordsOk

`func (o *DtoTaskResponse) GetProcessedRecordsOk() (*int32, bool)`

GetProcessedRecordsOk returns a tuple with the ProcessedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedRecords

`func (o *DtoTaskResponse) SetProcessedRecords(v int32)`

SetProcessedRecords sets ProcessedRecords field to given value.

### HasProcessedRecords

`func (o *DtoTaskResponse) HasProcessedRecords() bool`

HasProcessedRecords returns a boolean if a field has been set.

### GetStartedAt

`func (o *DtoTaskResponse) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DtoTaskResponse) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DtoTaskResponse) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DtoTaskResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DtoTaskResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtoTaskResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtoTaskResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtoTaskResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessfulRecords

`func (o *DtoTaskResponse) GetSuccessfulRecords() int32`

GetSuccessfulRecords returns the SuccessfulRecords field if non-nil, zero value otherwise.

### GetSuccessfulRecordsOk

`func (o *DtoTaskResponse) GetSuccessfulRecordsOk() (*int32, bool)`

GetSuccessfulRecordsOk returns a tuple with the SuccessfulRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulRecords

`func (o *DtoTaskResponse) SetSuccessfulRecords(v int32)`

SetSuccessfulRecords sets SuccessfulRecords field to given value.

### HasSuccessfulRecords

`func (o *DtoTaskResponse) HasSuccessfulRecords() bool`

HasSuccessfulRecords returns a boolean if a field has been set.

### GetTaskStatus

`func (o *DtoTaskResponse) GetTaskStatus() TypesTaskStatus`

GetTaskStatus returns the TaskStatus field if non-nil, zero value otherwise.

### GetTaskStatusOk

`func (o *DtoTaskResponse) GetTaskStatusOk() (*TypesTaskStatus, bool)`

GetTaskStatusOk returns a tuple with the TaskStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStatus

`func (o *DtoTaskResponse) SetTaskStatus(v TypesTaskStatus)`

SetTaskStatus sets TaskStatus field to given value.

### HasTaskStatus

`func (o *DtoTaskResponse) HasTaskStatus() bool`

HasTaskStatus returns a boolean if a field has been set.

### GetTaskType

`func (o *DtoTaskResponse) GetTaskType() TypesTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *DtoTaskResponse) GetTaskTypeOk() (*TypesTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *DtoTaskResponse) SetTaskType(v TypesTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *DtoTaskResponse) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetTenantId

`func (o *DtoTaskResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DtoTaskResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DtoTaskResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DtoTaskResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTotalRecords

`func (o *DtoTaskResponse) GetTotalRecords() int32`

GetTotalRecords returns the TotalRecords field if non-nil, zero value otherwise.

### GetTotalRecordsOk

`func (o *DtoTaskResponse) GetTotalRecordsOk() (*int32, bool)`

GetTotalRecordsOk returns a tuple with the TotalRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecords

`func (o *DtoTaskResponse) SetTotalRecords(v int32)`

SetTotalRecords sets TotalRecords field to given value.

### HasTotalRecords

`func (o *DtoTaskResponse) HasTotalRecords() bool`

HasTotalRecords returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DtoTaskResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DtoTaskResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DtoTaskResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DtoTaskResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DtoTaskResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DtoTaskResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DtoTaskResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DtoTaskResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


