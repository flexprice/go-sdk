# \EventsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsAnalyticsPost**](EventsAPI.md#EventsAnalyticsPost) | **Post** /events/analytics | Get usage analytics
[**EventsBulkPost**](EventsAPI.md#EventsBulkPost) | **Post** /events/bulk | Bulk Ingest events
[**EventsHuggingfaceInferencePost**](EventsAPI.md#EventsHuggingfaceInferencePost) | **Post** /events/huggingface-inference | Get hugging face inference data
[**EventsMonitoringGet**](EventsAPI.md#EventsMonitoringGet) | **Get** /events/monitoring | Get monitoring data
[**EventsPost**](EventsAPI.md#EventsPost) | **Post** /events | Ingest event
[**EventsQueryPost**](EventsAPI.md#EventsQueryPost) | **Post** /events/query | List raw events
[**EventsUsageMeterPost**](EventsAPI.md#EventsUsageMeterPost) | **Post** /events/usage/meter | Get usage by meter
[**EventsUsagePost**](EventsAPI.md#EventsUsagePost) | **Post** /events/usage | Get usage statistics



## EventsAnalyticsPost

> DtoGetUsageAnalyticsResponse EventsAnalyticsPost(ctx).Request(request).Execute()

Get usage analytics



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
	request := *openapiclient.NewDtoGetUsageAnalyticsRequest("ExternalCustomerId_example") // DtoGetUsageAnalyticsRequest | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsAnalyticsPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsAnalyticsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsAnalyticsPost`: DtoGetUsageAnalyticsResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsAnalyticsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsAnalyticsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoGetUsageAnalyticsRequest**](DtoGetUsageAnalyticsRequest.md) | Request body | 

### Return type

[**DtoGetUsageAnalyticsResponse**](DtoGetUsageAnalyticsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## EventsHuggingfaceInferencePost

> DtoGetHuggingFaceBillingDataResponse EventsHuggingfaceInferencePost(ctx).Execute()

Get hugging face inference data



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsHuggingfaceInferencePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsHuggingfaceInferencePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsHuggingfaceInferencePost`: DtoGetHuggingFaceBillingDataResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsHuggingfaceInferencePost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEventsHuggingfaceInferencePostRequest struct via the builder pattern


### Return type

[**DtoGetHuggingFaceBillingDataResponse**](DtoGetHuggingFaceBillingDataResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsMonitoringGet

> DtoGetMonitoringDataResponse EventsMonitoringGet(ctx).WindowSize(windowSize).Execute()

Get monitoring data



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
	windowSize := "windowSize_example" // string | Window size for time series data (e.g., 'HOUR', 'DAY') - optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.EventsMonitoringGet(context.Background()).WindowSize(windowSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsMonitoringGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EventsMonitoringGet`: DtoGetMonitoringDataResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsMonitoringGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsMonitoringGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **windowSize** | **string** | Window size for time series data (e.g., &#39;HOUR&#39;, &#39;DAY&#39;) - optional | 

### Return type

[**DtoGetMonitoringDataResponse**](DtoGetMonitoringDataResponse.md)

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

