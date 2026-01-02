# \PlansAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlansGet**](PlansAPI.md#PlansGet) | **Get** /plans | Get plans
[**PlansIdDelete**](PlansAPI.md#PlansIdDelete) | **Delete** /plans/{id} | Delete a plan
[**PlansIdGet**](PlansAPI.md#PlansIdGet) | **Get** /plans/{id} | Get a plan
[**PlansIdPut**](PlansAPI.md#PlansIdPut) | **Put** /plans/{id} | Update a plan
[**PlansIdSyncSubscriptionsPost**](PlansAPI.md#PlansIdSyncSubscriptionsPost) | **Post** /plans/{id}/sync/subscriptions | Synchronize plan prices
[**PlansPost**](PlansAPI.md#PlansPost) | **Post** /plans | Create a new plan
[**PlansSearchPost**](PlansAPI.md#PlansSearchPost) | **Post** /plans/search | List plans by filter



## PlansGet

> DtoListPlansResponse PlansGet(ctx).EndTime(endTime).Expand(expand).Limit(limit).LookupKey(lookupKey).Offset(offset).Order(order).PlanIds(planIds).StartTime(startTime).Status(status).Execute()

Get plans



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
	endTime := "endTime_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	lookupKey := "lookupKey_example" // string |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	planIds := []string{"Inner_example"} // []string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.PlansGet(context.Background()).EndTime(endTime).Expand(expand).Limit(limit).LookupKey(lookupKey).Offset(offset).Order(order).PlanIds(planIds).StartTime(startTime).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PlansGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlansGet`: DtoListPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.PlansGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlansGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **lookupKey** | **string** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **planIds** | **[]string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**DtoListPlansResponse**](DtoListPlansResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlansIdDelete

> DtoSuccessResponse PlansIdDelete(ctx, id).Execute()

Delete a plan



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
	id := "id_example" // string | Plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.PlansIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PlansIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlansIdDelete`: DtoSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.PlansIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlansIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoSuccessResponse**](DtoSuccessResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlansIdGet

> DtoPlanResponse PlansIdGet(ctx, id).Execute()

Get a plan



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
	id := "id_example" // string | Plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.PlansIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PlansIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlansIdGet`: DtoPlanResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.PlansIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlansIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoPlanResponse**](DtoPlanResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlansIdPut

> DtoPlanResponse PlansIdPut(ctx, id).Plan(plan).Execute()

Update a plan



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
	id := "id_example" // string | Plan ID
	plan := *openapiclient.NewDtoUpdatePlanRequest() // DtoUpdatePlanRequest | Plan update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.PlansIdPut(context.Background(), id).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PlansIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlansIdPut`: DtoPlanResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.PlansIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlansIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **plan** | [**DtoUpdatePlanRequest**](DtoUpdatePlanRequest.md) | Plan update | 

### Return type

[**DtoPlanResponse**](DtoPlanResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlansIdSyncSubscriptionsPost

> ModelsTemporalWorkflowResult PlansIdSyncSubscriptionsPost(ctx, id).Execute()

Synchronize plan prices



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
	id := "id_example" // string | Plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.PlansIdSyncSubscriptionsPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PlansIdSyncSubscriptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlansIdSyncSubscriptionsPost`: ModelsTemporalWorkflowResult
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.PlansIdSyncSubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlansIdSyncSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PlansPost

> DtoPlanResponse PlansPost(ctx).Plan(plan).Execute()

Create a new plan



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
	plan := *openapiclient.NewDtoCreatePlanRequest("Name_example") // DtoCreatePlanRequest | Plan configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.PlansPost(context.Background()).Plan(plan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PlansPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlansPost`: DtoPlanResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.PlansPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlansPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **plan** | [**DtoCreatePlanRequest**](DtoCreatePlanRequest.md) | Plan configuration | 

### Return type

[**DtoPlanResponse**](DtoPlanResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlansSearchPost

> DtoListPlansResponse PlansSearchPost(ctx).Filter(filter).Execute()

List plans by filter



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
	filter := *openapiclient.NewTypesPlanFilter() // TypesPlanFilter | Filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.PlansSearchPost(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PlansSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlansSearchPost`: DtoListPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.PlansSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlansSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**TypesPlanFilter**](TypesPlanFilter.md) | Filter | 

### Return type

[**DtoListPlansResponse**](DtoListPlansResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

