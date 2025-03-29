# \MetersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetersGet**](MetersAPI.md#MetersGet) | **Get** /meters | List meters
[**MetersIdDelete**](MetersAPI.md#MetersIdDelete) | **Delete** /meters/{id} | Delete meter
[**MetersIdDisablePost**](MetersAPI.md#MetersIdDisablePost) | **Post** /meters/{id}/disable | Disable meter [TODO: Deprecate]
[**MetersIdGet**](MetersAPI.md#MetersIdGet) | **Get** /meters/{id} | Get meter
[**MetersIdPut**](MetersAPI.md#MetersIdPut) | **Put** /meters/{id} | Update meter
[**MetersPost**](MetersAPI.md#MetersPost) | **Post** /meters | Create meter



## MetersGet

> []DtoMeterResponse MetersGet(ctx).EndTime(endTime).EventName(eventName).Expand(expand).Limit(limit).MeterIds(meterIds).Offset(offset).Order(order).Sort(sort).StartTime(startTime).Status(status).Execute()

List meters



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
	eventName := "eventName_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	meterIds := []string{"Inner_example"} // []string |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersGet(context.Background()).EndTime(endTime).EventName(eventName).Expand(expand).Limit(limit).MeterIds(meterIds).Offset(offset).Order(order).Sort(sort).StartTime(startTime).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersGet`: []DtoMeterResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **string** |  | 
 **eventName** | **string** |  | 
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **meterIds** | **[]string** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **sort** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**[]DtoMeterResponse**](DtoMeterResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersIdDelete

> map[string]string MetersIdDelete(ctx, id).Execute()

Delete meter



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
	id := "id_example" // string | Meter ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersIdDelete`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Meter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetersIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersIdDisablePost

> map[string]string MetersIdDisablePost(ctx, id).Execute()

Disable meter [TODO: Deprecate]



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
	id := "id_example" // string | Meter ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersIdDisablePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersIdDisablePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersIdDisablePost`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersIdDisablePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Meter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetersIdDisablePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersIdGet

> DtoMeterResponse MetersIdGet(ctx, id).Execute()

Get meter



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
	id := "id_example" // string | Meter ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersIdGet`: DtoMeterResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Meter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoMeterResponse**](DtoMeterResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersIdPut

> DtoMeterResponse MetersIdPut(ctx, id).Meter(meter).Execute()

Update meter



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
	id := "id_example" // string | Meter ID
	meter := *openapiclient.NewDtoUpdateMeterRequest() // DtoUpdateMeterRequest | Meter configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersIdPut(context.Background(), id).Meter(meter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersIdPut`: DtoMeterResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Meter ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetersIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **meter** | [**DtoUpdateMeterRequest**](DtoUpdateMeterRequest.md) | Meter configuration | 

### Return type

[**DtoMeterResponse**](DtoMeterResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersPost

> DtoMeterResponse MetersPost(ctx).Meter(meter).Execute()

Create meter



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
	meter := *openapiclient.NewDtoCreateMeterRequest(*openapiclient.NewMeterAggregation(), "api_request", "API Usage Meter", openapiclient.types.ResetUsage("BILLING_PERIOD")) // DtoCreateMeterRequest | Meter configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.MetersPost(context.Background()).Meter(meter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetersPost`: DtoMeterResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.MetersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **meter** | [**DtoCreateMeterRequest**](DtoCreateMeterRequest.md) | Meter configuration | 

### Return type

[**DtoMeterResponse**](DtoMeterResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

