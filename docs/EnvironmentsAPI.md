# \EnvironmentsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsGet**](EnvironmentsAPI.md#EnvironmentsGet) | **Get** /environments | Get environments
[**EnvironmentsIdGet**](EnvironmentsAPI.md#EnvironmentsIdGet) | **Get** /environments/{id} | Get an environment
[**EnvironmentsIdPut**](EnvironmentsAPI.md#EnvironmentsIdPut) | **Put** /environments/{id} | Update an environment
[**EnvironmentsPost**](EnvironmentsAPI.md#EnvironmentsPost) | **Post** /environments | Create an environment



## EnvironmentsGet

> DtoListEnvironmentsResponse EnvironmentsGet(ctx).Expand(expand).Limit(limit).Offset(offset).Order(order).Sort(sort).Status(status).Execute()

Get environments



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
	expand := "expand_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.EnvironmentsGet(context.Background()).Expand(expand).Limit(limit).Offset(offset).Order(order).Sort(sort).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.EnvironmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsGet`: DtoListEnvironmentsResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.EnvironmentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **sort** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**DtoListEnvironmentsResponse**](DtoListEnvironmentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsIdGet

> DtoEnvironmentResponse EnvironmentsIdGet(ctx, id).Execute()

Get an environment



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
	id := "id_example" // string | Environment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.EnvironmentsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.EnvironmentsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsIdGet`: DtoEnvironmentResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.EnvironmentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoEnvironmentResponse**](DtoEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsIdPut

> DtoEnvironmentResponse EnvironmentsIdPut(ctx, id).Environment(environment).Execute()

Update an environment



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
	id := "id_example" // string | Environment ID
	environment := *openapiclient.NewDtoUpdateEnvironmentRequest() // DtoUpdateEnvironmentRequest | Environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.EnvironmentsIdPut(context.Background(), id).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.EnvironmentsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsIdPut`: DtoEnvironmentResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.EnvironmentsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | [**DtoUpdateEnvironmentRequest**](DtoUpdateEnvironmentRequest.md) | Environment | 

### Return type

[**DtoEnvironmentResponse**](DtoEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsPost

> DtoEnvironmentResponse EnvironmentsPost(ctx).Environment(environment).Execute()

Create an environment



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
	environment := *openapiclient.NewDtoCreateEnvironmentRequest("Name_example", "Type_example") // DtoCreateEnvironmentRequest | Environment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.EnvironmentsPost(context.Background()).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.EnvironmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsPost`: DtoEnvironmentResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.EnvironmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | [**DtoCreateEnvironmentRequest**](DtoCreateEnvironmentRequest.md) | Environment | 

### Return type

[**DtoEnvironmentResponse**](DtoEnvironmentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

