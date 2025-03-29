# DtoCreateTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | [**TypesEntityType**](TypesEntityType.md) |  | 
**FileName** | Pointer to **string** |  | [optional] 
**FileType** | [**TypesFileType**](TypesFileType.md) |  | 
**FileUrl** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**TaskType** | [**TypesTaskType**](TypesTaskType.md) |  | 

## Methods

### NewDtoCreateTaskRequest

`func NewDtoCreateTaskRequest(entityType TypesEntityType, fileType TypesFileType, fileUrl string, taskType TypesTaskType, ) *DtoCreateTaskRequest`

NewDtoCreateTaskRequest instantiates a new DtoCreateTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtoCreateTaskRequestWithDefaults

`func NewDtoCreateTaskRequestWithDefaults() *DtoCreateTaskRequest`

NewDtoCreateTaskRequestWithDefaults instantiates a new DtoCreateTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *DtoCreateTaskRequest) GetEntityType() TypesEntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DtoCreateTaskRequest) GetEntityTypeOk() (*TypesEntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DtoCreateTaskRequest) SetEntityType(v TypesEntityType)`

SetEntityType sets EntityType field to given value.


### GetFileName

`func (o *DtoCreateTaskRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DtoCreateTaskRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DtoCreateTaskRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DtoCreateTaskRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileType

`func (o *DtoCreateTaskRequest) GetFileType() TypesFileType`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *DtoCreateTaskRequest) GetFileTypeOk() (*TypesFileType, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *DtoCreateTaskRequest) SetFileType(v TypesFileType)`

SetFileType sets FileType field to given value.


### GetFileUrl

`func (o *DtoCreateTaskRequest) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *DtoCreateTaskRequest) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *DtoCreateTaskRequest) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### GetMetadata

`func (o *DtoCreateTaskRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DtoCreateTaskRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DtoCreateTaskRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DtoCreateTaskRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTaskType

`func (o *DtoCreateTaskRequest) GetTaskType() TypesTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *DtoCreateTaskRequest) GetTaskTypeOk() (*TypesTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *DtoCreateTaskRequest) SetTaskType(v TypesTaskType)`

SetTaskType sets TaskType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


