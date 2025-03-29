# \PricesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricesGet**](PricesAPI.md#PricesGet) | **Get** /prices | Get prices
[**PricesIdDelete**](PricesAPI.md#PricesIdDelete) | **Delete** /prices/{id} | Delete a price
[**PricesIdGet**](PricesAPI.md#PricesIdGet) | **Get** /prices/{id} | Get a price by ID
[**PricesIdPut**](PricesAPI.md#PricesIdPut) | **Put** /prices/{id} | Update a price
[**PricesPost**](PricesAPI.md#PricesPost) | **Post** /prices | Create a new price



## PricesGet

> DtoListPricesResponse PricesGet(ctx).EndTime(endTime).Expand(expand).Limit(limit).Offset(offset).Order(order).PlanIds(planIds).PriceIds(priceIds).Sort(sort).StartTime(startTime).Status(status).Execute()

Get prices



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
	priceIds := []string{"Inner_example"} // []string |  (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.PricesGet(context.Background()).EndTime(endTime).Expand(expand).Limit(limit).Offset(offset).Order(order).PlanIds(planIds).PriceIds(priceIds).Sort(sort).StartTime(startTime).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.PricesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesGet`: DtoListPricesResponse
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.PricesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **planIds** | **[]string** |  | 
 **priceIds** | **[]string** |  | 
 **sort** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**DtoListPricesResponse**](DtoListPricesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesIdDelete

> map[string]map[string]interface{} PricesIdDelete(ctx, id).Execute()

Delete a price



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
	id := "id_example" // string | Price ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.PricesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.PricesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesIdDelete`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.PricesIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricesIdDeleteRequest struct via the builder pattern


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


## PricesIdGet

> DtoPriceResponse PricesIdGet(ctx, id).Execute()

Get a price by ID



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
	id := "id_example" // string | Price ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.PricesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.PricesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesIdGet`: DtoPriceResponse
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.PricesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoPriceResponse**](DtoPriceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesIdPut

> DtoPriceResponse PricesIdPut(ctx, id).Price(price).Execute()

Update a price



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
	id := "id_example" // string | Price ID
	price := *openapiclient.NewDtoUpdatePriceRequest() // DtoUpdatePriceRequest | Price configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.PricesIdPut(context.Background(), id).Price(price).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.PricesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesIdPut`: DtoPriceResponse
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.PricesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **price** | [**DtoUpdatePriceRequest**](DtoUpdatePriceRequest.md) | Price configuration | 

### Return type

[**DtoPriceResponse**](DtoPriceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesPost

> DtoPriceResponse PricesPost(ctx).Price(price).Execute()

Create a new price



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
	price := *openapiclient.NewDtoCreatePriceRequest(openapiclient.types.BillingCadence("RECURRING"), openapiclient.types.BillingModel("FLAT_FEE"), openapiclient.types.BillingPeriod("MONTHLY"), int32(123), "Currency_example", openapiclient.types.InvoiceCadence("ARREAR"), openapiclient.types.PriceType("USAGE")) // DtoCreatePriceRequest | Price configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricesAPI.PricesPost(context.Background()).Price(price).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricesAPI.PricesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesPost`: DtoPriceResponse
	fmt.Fprintf(os.Stdout, "Response from `PricesAPI.PricesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **price** | [**DtoCreatePriceRequest**](DtoCreatePriceRequest.md) | Price configuration | 

### Return type

[**DtoPriceResponse**](DtoPriceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

