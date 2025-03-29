# \PlansAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlansGet**](PlansAPI.md#PlansGet) | **Get** /plans | Get plans
[**PlansIdDelete**](PlansAPI.md#PlansIdDelete) | **Delete** /plans/{id} | Delete a plan
[**PlansIdGet**](PlansAPI.md#PlansIdGet) | **Get** /plans/{id} | Get a plan
[**PlansIdPut**](PlansAPI.md#PlansIdPut) | **Put** /plans/{id} | Update a plan
[**PlansPost**](PlansAPI.md#PlansPost) | **Post** /plans | Create a new plan



## PlansGet

> DtoListPlansResponse PlansGet(ctx).EndTime(endTime).Expand(expand).Limit(limit).Offset(offset).Order(order).PlanIds(planIds).Sort(sort).StartTime(startTime).Status(status).Execute()

Get plans



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
	endTime := "endTime_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	planIds := []string{"Inner_example"} // []string |  (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.PlansGet(context.Background()).EndTime(endTime).Expand(expand).Limit(limit).Offset(offset).Order(order).PlanIds(planIds).Sort(sort).StartTime(startTime).Status(status).Execute()
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
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **planIds** | **[]string** |  | 
 **sort** | **string** |  | 
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

> map[string]map[string]interface{} PlansIdDelete(ctx, id).Execute()

Delete a plan



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
	id := "id_example" // string | Plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.PlansIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PlansIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlansIdDelete`: map[string]map[string]interface{}
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

**map[string]map[string]interface{}**

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
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
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
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
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
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
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

