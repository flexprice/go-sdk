# \ScheduledTasksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksScheduledGet**](ScheduledTasksAPI.md#TasksScheduledGet) | **Get** /tasks/scheduled | List scheduled tasks
[**TasksScheduledIdDelete**](ScheduledTasksAPI.md#TasksScheduledIdDelete) | **Delete** /tasks/scheduled/{id} | Delete a scheduled task
[**TasksScheduledIdGet**](ScheduledTasksAPI.md#TasksScheduledIdGet) | **Get** /tasks/scheduled/{id} | Get a scheduled task
[**TasksScheduledIdPut**](ScheduledTasksAPI.md#TasksScheduledIdPut) | **Put** /tasks/scheduled/{id} | Update a scheduled task
[**TasksScheduledIdRunPost**](ScheduledTasksAPI.md#TasksScheduledIdRunPost) | **Post** /tasks/scheduled/{id}/run | Trigger force run
[**TasksScheduledPost**](ScheduledTasksAPI.md#TasksScheduledPost) | **Post** /tasks/scheduled | Create a scheduled task



## TasksScheduledGet

> DtoListScheduledTasksResponse TasksScheduledGet(ctx).Limit(limit).Offset(offset).ConnectionId(connectionId).EntityType(entityType).Interval(interval).Enabled(enabled).Execute()

List scheduled tasks



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
	limit := int32(56) // int32 | Limit (optional)
	offset := int32(56) // int32 | Offset (optional)
	connectionId := "connectionId_example" // string | Filter by connection ID (optional)
	entityType := "entityType_example" // string | Filter by entity type (optional)
	interval := "interval_example" // string | Filter by interval (optional)
	enabled := true // bool | Filter by enabled status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduledTasksAPI.TasksScheduledGet(context.Background()).Limit(limit).Offset(offset).ConnectionId(connectionId).EntityType(entityType).Interval(interval).Enabled(enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTasksAPI.TasksScheduledGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksScheduledGet`: DtoListScheduledTasksResponse
	fmt.Fprintf(os.Stdout, "Response from `ScheduledTasksAPI.TasksScheduledGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksScheduledGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit | 
 **offset** | **int32** | Offset | 
 **connectionId** | **string** | Filter by connection ID | 
 **entityType** | **string** | Filter by entity type | 
 **interval** | **string** | Filter by interval | 
 **enabled** | **bool** | Filter by enabled status | 

### Return type

[**DtoListScheduledTasksResponse**](DtoListScheduledTasksResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksScheduledIdDelete

> TasksScheduledIdDelete(ctx, id).Execute()

Delete a scheduled task



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
	id := "id_example" // string | Scheduled Task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScheduledTasksAPI.TasksScheduledIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTasksAPI.TasksScheduledIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Scheduled Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksScheduledIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksScheduledIdGet

> DtoScheduledTaskResponse TasksScheduledIdGet(ctx, id).Execute()

Get a scheduled task



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
	id := "id_example" // string | Scheduled Task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduledTasksAPI.TasksScheduledIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTasksAPI.TasksScheduledIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksScheduledIdGet`: DtoScheduledTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ScheduledTasksAPI.TasksScheduledIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Scheduled Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksScheduledIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoScheduledTaskResponse**](DtoScheduledTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksScheduledIdPut

> DtoScheduledTaskResponse TasksScheduledIdPut(ctx, id).ScheduledTask(scheduledTask).Execute()

Update a scheduled task



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
	id := "id_example" // string | Scheduled Task ID
	scheduledTask := *openapiclient.NewDtoUpdateScheduledTaskRequest(false) // DtoUpdateScheduledTaskRequest | Update request (enabled: true/false to pause/resume)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduledTasksAPI.TasksScheduledIdPut(context.Background(), id).ScheduledTask(scheduledTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTasksAPI.TasksScheduledIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksScheduledIdPut`: DtoScheduledTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ScheduledTasksAPI.TasksScheduledIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Scheduled Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksScheduledIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduledTask** | [**DtoUpdateScheduledTaskRequest**](DtoUpdateScheduledTaskRequest.md) | Update request (enabled: true/false to pause/resume) | 

### Return type

[**DtoScheduledTaskResponse**](DtoScheduledTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksScheduledIdRunPost

> DtoTriggerForceRunResponse TasksScheduledIdRunPost(ctx, id).Request(request).Execute()

Trigger force run



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
	id := "id_example" // string | Scheduled Task ID
	request := *openapiclient.NewDtoTriggerForceRunRequest() // DtoTriggerForceRunRequest | Optional start and end time for custom range (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduledTasksAPI.TasksScheduledIdRunPost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTasksAPI.TasksScheduledIdRunPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksScheduledIdRunPost`: DtoTriggerForceRunResponse
	fmt.Fprintf(os.Stdout, "Response from `ScheduledTasksAPI.TasksScheduledIdRunPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Scheduled Task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksScheduledIdRunPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoTriggerForceRunRequest**](DtoTriggerForceRunRequest.md) | Optional start and end time for custom range | 

### Return type

[**DtoTriggerForceRunResponse**](DtoTriggerForceRunResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksScheduledPost

> DtoScheduledTaskResponse TasksScheduledPost(ctx).ScheduledTask(scheduledTask).Execute()

Create a scheduled task



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
	scheduledTask := *openapiclient.NewDtoCreateScheduledTaskRequest("ConnectionId_example", openapiclient.types.ScheduledTaskEntityType("events"), openapiclient.types.ScheduledTaskInterval("custom"), *openapiclient.NewTypesS3JobConfig()) // DtoCreateScheduledTaskRequest | Scheduled Task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScheduledTasksAPI.TasksScheduledPost(context.Background()).ScheduledTask(scheduledTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScheduledTasksAPI.TasksScheduledPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksScheduledPost`: DtoScheduledTaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ScheduledTasksAPI.TasksScheduledPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksScheduledPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduledTask** | [**DtoCreateScheduledTaskRequest**](DtoCreateScheduledTaskRequest.md) | Scheduled Task | 

### Return type

[**DtoScheduledTaskResponse**](DtoScheduledTaskResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

