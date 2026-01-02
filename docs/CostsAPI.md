# \CostsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CostsActiveGet**](CostsAPI.md#CostsActiveGet) | **Get** /costs/active | Get active costsheet for tenant
[**CostsAnalyticsPost**](CostsAPI.md#CostsAnalyticsPost) | **Post** /costs/analytics | Get combined revenue and cost analytics
[**CostsAnalyticsV2Post**](CostsAPI.md#CostsAnalyticsV2Post) | **Post** /costs/analytics-v2 | Get combined revenue and cost analytics
[**CostsIdDelete**](CostsAPI.md#CostsIdDelete) | **Delete** /costs/{id} | Delete a costsheet
[**CostsIdGet**](CostsAPI.md#CostsIdGet) | **Get** /costs/{id} | Get a costsheet by ID
[**CostsIdPut**](CostsAPI.md#CostsIdPut) | **Put** /costs/{id} | Update a costsheet
[**CostsPost**](CostsAPI.md#CostsPost) | **Post** /costs | Create a new costsheet
[**CostsSearchPost**](CostsAPI.md#CostsSearchPost) | **Post** /costs/search | List costsheets by filter



## CostsActiveGet

> DtoCostsheetResponse CostsActiveGet(ctx).Execute()

Get active costsheet for tenant



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
	resp, r, err := apiClient.CostsAPI.CostsActiveGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostsAPI.CostsActiveGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostsActiveGet`: DtoCostsheetResponse
	fmt.Fprintf(os.Stdout, "Response from `CostsAPI.CostsActiveGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCostsActiveGetRequest struct via the builder pattern


### Return type

[**DtoCostsheetResponse**](DtoCostsheetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostsAnalyticsPost

> DtoGetDetailedCostAnalyticsResponse CostsAnalyticsPost(ctx).Request(request).Execute()

Get combined revenue and cost analytics



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
	request := *openapiclient.NewDtoGetCostAnalyticsRequest() // DtoGetCostAnalyticsRequest | Combined analytics request (start_time/end_time optional - defaults to last 7 days)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostsAPI.CostsAnalyticsPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostsAPI.CostsAnalyticsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostsAnalyticsPost`: DtoGetDetailedCostAnalyticsResponse
	fmt.Fprintf(os.Stdout, "Response from `CostsAPI.CostsAnalyticsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostsAnalyticsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoGetCostAnalyticsRequest**](DtoGetCostAnalyticsRequest.md) | Combined analytics request (start_time/end_time optional - defaults to last 7 days) | 

### Return type

[**DtoGetDetailedCostAnalyticsResponse**](DtoGetDetailedCostAnalyticsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostsAnalyticsV2Post

> DtoGetDetailedCostAnalyticsResponse CostsAnalyticsV2Post(ctx).Request(request).Execute()

Get combined revenue and cost analytics



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
	request := *openapiclient.NewDtoGetCostAnalyticsRequest() // DtoGetCostAnalyticsRequest | Combined analytics request (start_time/end_time optional - defaults to last 7 days)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostsAPI.CostsAnalyticsV2Post(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostsAPI.CostsAnalyticsV2Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostsAnalyticsV2Post`: DtoGetDetailedCostAnalyticsResponse
	fmt.Fprintf(os.Stdout, "Response from `CostsAPI.CostsAnalyticsV2Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostsAnalyticsV2PostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoGetCostAnalyticsRequest**](DtoGetCostAnalyticsRequest.md) | Combined analytics request (start_time/end_time optional - defaults to last 7 days) | 

### Return type

[**DtoGetDetailedCostAnalyticsResponse**](DtoGetDetailedCostAnalyticsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostsIdDelete

> DtoDeleteCostsheetResponse CostsIdDelete(ctx, id).Execute()

Delete a costsheet



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
	id := "id_example" // string | Costsheet ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostsAPI.CostsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostsAPI.CostsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostsIdDelete`: DtoDeleteCostsheetResponse
	fmt.Fprintf(os.Stdout, "Response from `CostsAPI.CostsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Costsheet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCostsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoDeleteCostsheetResponse**](DtoDeleteCostsheetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostsIdGet

> DtoGetCostsheetResponse CostsIdGet(ctx, id).Expand(expand).Execute()

Get a costsheet by ID



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
	id := "id_example" // string | Costsheet ID
	expand := "expand_example" // string | Comma-separated list of fields to expand (e.g., 'prices') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostsAPI.CostsIdGet(context.Background(), id).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostsAPI.CostsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostsIdGet`: DtoGetCostsheetResponse
	fmt.Fprintf(os.Stdout, "Response from `CostsAPI.CostsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Costsheet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCostsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Comma-separated list of fields to expand (e.g., &#39;prices&#39;) | 

### Return type

[**DtoGetCostsheetResponse**](DtoGetCostsheetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostsIdPut

> DtoUpdateCostsheetResponse CostsIdPut(ctx, id).Costsheet(costsheet).Execute()

Update a costsheet



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
	id := "id_example" // string | Costsheet ID
	costsheet := *openapiclient.NewDtoUpdateCostsheetRequest() // DtoUpdateCostsheetRequest | Costsheet configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostsAPI.CostsIdPut(context.Background(), id).Costsheet(costsheet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostsAPI.CostsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostsIdPut`: DtoUpdateCostsheetResponse
	fmt.Fprintf(os.Stdout, "Response from `CostsAPI.CostsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Costsheet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCostsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **costsheet** | [**DtoUpdateCostsheetRequest**](DtoUpdateCostsheetRequest.md) | Costsheet configuration | 

### Return type

[**DtoUpdateCostsheetResponse**](DtoUpdateCostsheetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostsPost

> DtoCreateCostsheetResponse CostsPost(ctx).Costsheet(costsheet).Execute()

Create a new costsheet



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
	costsheet := *openapiclient.NewDtoCreateCostsheetRequest("Name_example") // DtoCreateCostsheetRequest | Costsheet configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostsAPI.CostsPost(context.Background()).Costsheet(costsheet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostsAPI.CostsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostsPost`: DtoCreateCostsheetResponse
	fmt.Fprintf(os.Stdout, "Response from `CostsAPI.CostsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **costsheet** | [**DtoCreateCostsheetRequest**](DtoCreateCostsheetRequest.md) | Costsheet configuration | 

### Return type

[**DtoCreateCostsheetResponse**](DtoCreateCostsheetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CostsSearchPost

> DtoListCostsheetResponse CostsSearchPost(ctx).Filter(filter).Execute()

List costsheets by filter



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
	filter := *openapiclient.NewCostsheetFilter() // CostsheetFilter | Filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CostsAPI.CostsSearchPost(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CostsAPI.CostsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CostsSearchPost`: DtoListCostsheetResponse
	fmt.Fprintf(os.Stdout, "Response from `CostsAPI.CostsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCostsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**CostsheetFilter**](CostsheetFilter.md) | Filter | 

### Return type

[**DtoListCostsheetResponse**](DtoListCostsheetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

