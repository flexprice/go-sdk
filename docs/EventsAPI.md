# \EventsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsBulkPost**](EventsAPI.md#EventsBulkPost) | **Post** /events/bulk | Bulk Ingest events
[**EventsGet**](EventsAPI.md#EventsGet) | **Get** /events | Get raw events
[**EventsPost**](EventsAPI.md#EventsPost) | **Post** /events | Ingest event
[**EventsUsageMeterPost**](EventsAPI.md#EventsUsageMeterPost) | **Post** /events/usage/meter | Get usage by meter
[**EventsUsagePost**](EventsAPI.md#EventsUsagePost) | **Post** /events/usage | Get usage statistics



## EventsBulkPost

> map[string]string EventsBulkPost(ctx).Event(event).Execute()

Bulk Ingest events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	event := *openapiclient.NewDtoBulkIngestEventRequest([]openapiclient.DtoIngestEventRequest{*openapiclient.NewDtoIngestEventRequest("api_request", "customer456")}) // DtoBulkIngestEventRequest | Event data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsBulkPost(context.Background()).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsBulkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsBulkPost`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsBulkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsBulkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **event** | [**DtoBulkIngestEventRequest**](DtoBulkIngestEventRequest.md) | Event data | 

### Return type

**map[string]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsGet

> DtoGetEventsResponse EventsGet(ctx).ExternalCustomerId(externalCustomerId).EventName(eventName).StartTime(startTime).EndTime(endTime).IterFirstKey(iterFirstKey).IterLastKey(iterLastKey).PageSize(pageSize).Execute()

Get raw events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	externalCustomerId := "externalCustomerId_example" // string | External Customer ID (optional)
	eventName := "eventName_example" // string | Event Name (optional)
	startTime := "startTime_example" // string | Start Time (RFC3339) (optional)
	endTime := "endTime_example" // string | End Time (RFC3339) (optional)
	iterFirstKey := "iterFirstKey_example" // string | Iter First Key (unix_timestamp_nanoseconds::event_id) (optional)
	iterLastKey := "iterLastKey_example" // string | Iter Last Key (unix_timestamp_nanoseconds::event_id) (optional)
	pageSize := int32(56) // int32 | Page Size (1-50) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsGet(context.Background()).ExternalCustomerId(externalCustomerId).EventName(eventName).StartTime(startTime).EndTime(endTime).IterFirstKey(iterFirstKey).IterLastKey(iterLastKey).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsGet`: DtoGetEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalCustomerId** | **string** | External Customer ID | 
 **eventName** | **string** | Event Name | 
 **startTime** | **string** | Start Time (RFC3339) | 
 **endTime** | **string** | End Time (RFC3339) | 
 **iterFirstKey** | **string** | Iter First Key (unix_timestamp_nanoseconds::event_id) | 
 **iterLastKey** | **string** | Iter Last Key (unix_timestamp_nanoseconds::event_id) | 
 **pageSize** | **int32** | Page Size (1-50) | 

### Return type

[**DtoGetEventsResponse**](DtoGetEventsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsPost

> map[string]string EventsPost(ctx).Event(event).Execute()

Ingest event



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	event := *openapiclient.NewDtoIngestEventRequest("api_request", "customer456") // DtoIngestEventRequest | Event data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsPost(context.Background()).Event(event).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsPost`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **event** | [**DtoIngestEventRequest**](DtoIngestEventRequest.md) | Event data | 

### Return type

**map[string]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsUsageMeterPost

> DtoGetUsageResponse EventsUsageMeterPost(ctx).Request(request).Execute()

Get usage by meter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	request := *openapiclient.NewDtoGetUsageByMeterRequest("123") // DtoGetUsageByMeterRequest | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsUsageMeterPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsUsageMeterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsUsageMeterPost`: DtoGetUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsUsageMeterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsUsageMeterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoGetUsageByMeterRequest**](DtoGetUsageByMeterRequest.md) | Request body | 

### Return type

[**DtoGetUsageResponse**](DtoGetUsageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsUsagePost

> DtoGetUsageResponse EventsUsagePost(ctx).Request(request).Execute()

Get usage statistics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	request := *openapiclient.NewDtoGetUsageRequest(openapiclient.types.AggregationType("COUNT"), "api_request") // DtoGetUsageRequest | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsUsagePost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsUsagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsUsagePost`: DtoGetUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsUsagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsUsagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoGetUsageRequest**](DtoGetUsageRequest.md) | Request body | 

### Return type

[**DtoGetUsageResponse**](DtoGetUsageResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

