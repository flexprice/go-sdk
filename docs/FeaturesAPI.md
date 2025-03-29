# \FeaturesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FeaturesGet**](FeaturesAPI.md#FeaturesGet) | **Get** /features | List features
[**FeaturesIdDelete**](FeaturesAPI.md#FeaturesIdDelete) | **Delete** /features/{id} | Delete a feature
[**FeaturesIdGet**](FeaturesAPI.md#FeaturesIdGet) | **Get** /features/{id} | Get a feature by ID
[**FeaturesIdPut**](FeaturesAPI.md#FeaturesIdPut) | **Put** /features/{id} | Update a feature
[**FeaturesPost**](FeaturesAPI.md#FeaturesPost) | **Post** /features | Create a new feature



## FeaturesGet

> DtoListFeaturesResponse FeaturesGet(ctx).EndTime(endTime).Expand(expand).FeatureIds(featureIds).Limit(limit).LookupKey(lookupKey).MeterIds(meterIds).Offset(offset).Order(order).Sort(sort).StartTime(startTime).Status(status).Execute()

List features



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
	featureIds := []string{"Inner_example"} // []string | Feature specific filters (optional)
	limit := int32(56) // int32 |  (optional)
	lookupKey := "lookupKey_example" // string |  (optional)
	meterIds := []string{"Inner_example"} // []string |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.FeaturesGet(context.Background()).EndTime(endTime).Expand(expand).FeatureIds(featureIds).Limit(limit).LookupKey(lookupKey).MeterIds(meterIds).Offset(offset).Order(order).Sort(sort).StartTime(startTime).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.FeaturesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeaturesGet`: DtoListFeaturesResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.FeaturesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeaturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **featureIds** | **[]string** | Feature specific filters | 
 **limit** | **int32** |  | 
 **lookupKey** | **string** |  | 
 **meterIds** | **[]string** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **sort** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**DtoListFeaturesResponse**](DtoListFeaturesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FeaturesIdDelete

> map[string]map[string]interface{} FeaturesIdDelete(ctx, id).Execute()

Delete a feature



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
	id := "id_example" // string | Feature ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.FeaturesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.FeaturesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeaturesIdDelete`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.FeaturesIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Feature ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFeaturesIdDeleteRequest struct via the builder pattern


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


## FeaturesIdGet

> DtoFeatureResponse FeaturesIdGet(ctx, id).Execute()

Get a feature by ID



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
	id := "id_example" // string | Feature ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.FeaturesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.FeaturesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeaturesIdGet`: DtoFeatureResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.FeaturesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Feature ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFeaturesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoFeatureResponse**](DtoFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FeaturesIdPut

> DtoFeatureResponse FeaturesIdPut(ctx, id).Feature(feature).Execute()

Update a feature



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
	id := "id_example" // string | Feature ID
	feature := *openapiclient.NewDtoUpdateFeatureRequest() // DtoUpdateFeatureRequest | Feature update data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.FeaturesIdPut(context.Background(), id).Feature(feature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.FeaturesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeaturesIdPut`: DtoFeatureResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.FeaturesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Feature ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFeaturesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **feature** | [**DtoUpdateFeatureRequest**](DtoUpdateFeatureRequest.md) | Feature update data | 

### Return type

[**DtoFeatureResponse**](DtoFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FeaturesPost

> DtoFeatureResponse FeaturesPost(ctx).Feature(feature).Execute()

Create a new feature



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
	feature := *openapiclient.NewDtoCreateFeatureRequest("Name_example", openapiclient.types.FeatureType("metered")) // DtoCreateFeatureRequest | Feature to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.FeaturesPost(context.Background()).Feature(feature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.FeaturesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeaturesPost`: DtoFeatureResponse
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.FeaturesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeaturesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feature** | [**DtoCreateFeatureRequest**](DtoCreateFeatureRequest.md) | Feature to create | 

### Return type

[**DtoFeatureResponse**](DtoFeatureResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

