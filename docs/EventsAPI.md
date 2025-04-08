# \EventsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsBulkPost**](EventsAPI.md#EventsBulkPost) | **Post** /events/bulk | Bulk Ingest events
[**EventsPost**](EventsAPI.md#EventsPost) | **Post** /events | Ingest event
[**EventsQueryPost**](EventsAPI.md#EventsQueryPost) | **Post** /events/query | List raw events
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
	openapiclient "github.com/flexprice/go-sdk/flexprice"
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
	openapiclient "github.com/flexprice/go-sdk/flexprice"
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


## EventsQueryPost

> DtoGetEventsResponse EventsQueryPost(ctx).Request(request).Execute()

List raw events



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
	request := *openapiclient.NewDtoGetEventsRequest() // DtoGetEventsRequest | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsQueryPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsQueryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsQueryPost`: DtoGetEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsQueryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoGetEventsRequest**](DtoGetEventsRequest.md) | Request body | 

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
	openapiclient "github.com/flexprice/go-sdk/flexprice"
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
	openapiclient "github.com/flexprice/go-sdk/flexprice"
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

