# \TasksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksGet**](TasksAPI.md#TasksGet) | **Get** /tasks | List tasks
[**TasksIdGet**](TasksAPI.md#TasksIdGet) | **Get** /tasks/{id} | Get a task
[**TasksIdStatusPut**](TasksAPI.md#TasksIdStatusPut) | **Put** /tasks/{id}/status | Update task status
[**TasksPost**](TasksAPI.md#TasksPost) | **Post** /tasks | Create a new task
[**TasksResultGet**](TasksAPI.md#TasksResultGet) | **Get** /tasks/result | Get task processing result



## TasksGet

> DtoListTasksResponse TasksGet(ctx).CreatedBy(createdBy).EndTime(endTime).EntityType(entityType).Expand(expand).Limit(limit).Offset(offset).Order(order).ScheduledTaskId(scheduledTaskId).Sort(sort).StartTime(startTime).Status(status).TaskStatus(taskStatus).TaskType(taskType).Execute()

List tasks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexprice"
)

func main() {
	createdBy := "createdBy_example" // string |  (optional)
	endTime := "endTime_example" // string |  (optional)
	entityType := "entityType_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	scheduledTaskId := "scheduledTaskId_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	taskStatus := "taskStatus_example" // string |  (optional)
	taskType := "taskType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksGet(context.Background()).CreatedBy(createdBy).EndTime(endTime).EntityType(entityType).Expand(expand).Limit(limit).Offset(offset).Order(order).ScheduledTaskId(scheduledTaskId).Sort(sort).StartTime(startTime).Status(status).TaskStatus(taskStatus).TaskType(taskType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksGet`: DtoListTasksResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdBy** | **string** |  | 
 **endTime** | **string** |  | 
 **entityType** | **string** |  | 
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **scheduledTaskId** | **string** |  | 
 **sort** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 
 **taskStatus** | **string** |  | 
 **taskType** | **string** |  | 

### Return type

[**DtoListTasksResponse**](DtoListTasksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdGet

> DtoTaskResponse TasksIdGet(ctx, id).Execute()

Get a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexprice"
)

func main() {
	id := "id_example" // string | Task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdGet`: DtoTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoTaskResponse**](DtoTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksIdStatusPut

> map[string]map[string]interface{} TasksIdStatusPut(ctx, id).Status(status).Execute()

Update task status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexprice"
)

func main() {
	id := "id_example" // string | Task ID
	status := *openapiclient.NewDtoUpdateTaskStatusRequest(openapiclient.types.TaskStatus("PENDING")) // DtoUpdateTaskStatusRequest | Status update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksIdStatusPut(context.Background(), id).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksIdStatusPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksIdStatusPut`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksIdStatusPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksIdStatusPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**DtoUpdateTaskStatusRequest**](DtoUpdateTaskStatusRequest.md) | Status update | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksPost

> DtoTaskResponse TasksPost(ctx).Task(task).Execute()

Create a new task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexprice"
)

func main() {
	task := *openapiclient.NewDtoCreateTaskRequest(openapiclient.types.EntityType("EVENTS"), openapiclient.types.FileType("CSV"), "FileUrl_example", openapiclient.types.TaskType("IMPORT")) // DtoCreateTaskRequest | Task configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksPost(context.Background()).Task(task).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksPost`: DtoTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **task** | [**DtoCreateTaskRequest**](DtoCreateTaskRequest.md) | Task configuration | 

### Return type

[**DtoTaskResponse**](DtoTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksResultGet

> ModelsTemporalWorkflowResult TasksResultGet(ctx).WorkflowId(workflowId).Execute()

Get task processing result



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexprice"
)

func main() {
	workflowId := "workflowId_example" // string | Workflow ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksResultGet(context.Background()).WorkflowId(workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksResultGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksResultGet`: ModelsTemporalWorkflowResult
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksResultGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksResultGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowId** | **string** | Workflow ID | 

### Return type

[**ModelsTemporalWorkflowResult**](ModelsTemporalWorkflowResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

