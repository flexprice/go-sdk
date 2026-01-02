# \PriceUnitsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricesUnitsCodeCodeGet**](PriceUnitsAPI.md#PricesUnitsCodeCodeGet) | **Get** /prices/units/code/{code} | Get a price unit by code
[**PricesUnitsGet**](PriceUnitsAPI.md#PricesUnitsGet) | **Get** /prices/units | List price units
[**PricesUnitsIdDelete**](PriceUnitsAPI.md#PricesUnitsIdDelete) | **Delete** /prices/units/{id} | Delete a price unit
[**PricesUnitsIdGet**](PriceUnitsAPI.md#PricesUnitsIdGet) | **Get** /prices/units/{id} | Get a price unit by ID
[**PricesUnitsIdPut**](PriceUnitsAPI.md#PricesUnitsIdPut) | **Put** /prices/units/{id} | Update a price unit
[**PricesUnitsPost**](PriceUnitsAPI.md#PricesUnitsPost) | **Post** /prices/units | Create a new price unit
[**PricesUnitsSearchPost**](PriceUnitsAPI.md#PricesUnitsSearchPost) | **Post** /prices/units/search | List price units by filter



## PricesUnitsCodeCodeGet

> DtoPriceUnitResponse PricesUnitsCodeCodeGet(ctx, code).Execute()

Get a price unit by code



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
	code := "code_example" // string | Price unit code

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceUnitsAPI.PricesUnitsCodeCodeGet(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceUnitsAPI.PricesUnitsCodeCodeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesUnitsCodeCodeGet`: DtoPriceUnitResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceUnitsAPI.PricesUnitsCodeCodeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Price unit code | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricesUnitsCodeCodeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoPriceUnitResponse**](DtoPriceUnitResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesUnitsGet

> DtoListPriceUnitsResponse PricesUnitsGet(ctx).Status(status).Limit(limit).Offset(offset).Sort(sort).Order(order).Execute()

List price units



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
	status := "status_example" // string | Filter by status (optional)
	limit := int32(56) // int32 | Limit number of results (optional)
	offset := int32(56) // int32 | Offset for pagination (optional)
	sort := "sort_example" // string | Sort field (optional)
	order := "order_example" // string | Sort order (asc/desc) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceUnitsAPI.PricesUnitsGet(context.Background()).Status(status).Limit(limit).Offset(offset).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceUnitsAPI.PricesUnitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesUnitsGet`: DtoListPriceUnitsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceUnitsAPI.PricesUnitsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricesUnitsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter by status | 
 **limit** | **int32** | Limit number of results | 
 **offset** | **int32** | Offset for pagination | 
 **sort** | **string** | Sort field | 
 **order** | **string** | Sort order (asc/desc) | 

### Return type

[**DtoListPriceUnitsResponse**](DtoListPriceUnitsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesUnitsIdDelete

> DtoSuccessResponse PricesUnitsIdDelete(ctx, id).Execute()

Delete a price unit



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
	id := "id_example" // string | Price unit ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceUnitsAPI.PricesUnitsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceUnitsAPI.PricesUnitsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesUnitsIdDelete`: DtoSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceUnitsAPI.PricesUnitsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price unit ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricesUnitsIdDeleteRequest struct via the builder pattern


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


## PricesUnitsIdGet

> DtoPriceUnitResponse PricesUnitsIdGet(ctx, id).Execute()

Get a price unit by ID



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
	id := "id_example" // string | Price unit ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceUnitsAPI.PricesUnitsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceUnitsAPI.PricesUnitsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesUnitsIdGet`: DtoPriceUnitResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceUnitsAPI.PricesUnitsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price unit ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricesUnitsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoPriceUnitResponse**](DtoPriceUnitResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesUnitsIdPut

> DtoPriceUnitResponse PricesUnitsIdPut(ctx, id).Body(body).Execute()

Update a price unit



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
	id := "id_example" // string | Price unit ID
	body := *openapiclient.NewDtoUpdatePriceUnitRequest() // DtoUpdatePriceUnitRequest | Price unit details to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceUnitsAPI.PricesUnitsIdPut(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceUnitsAPI.PricesUnitsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesUnitsIdPut`: DtoPriceUnitResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceUnitsAPI.PricesUnitsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Price unit ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricesUnitsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DtoUpdatePriceUnitRequest**](DtoUpdatePriceUnitRequest.md) | Price unit details to update | 

### Return type

[**DtoPriceUnitResponse**](DtoPriceUnitResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesUnitsPost

> DtoCreatePriceUnitResponse PricesUnitsPost(ctx).Body(body).Execute()

Create a new price unit



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
	body := *openapiclient.NewDtoCreatePriceUnitRequest("BaseCurrency_example", "Code_example", "ConversionRate_example", "Name_example", "Symbol_example") // DtoCreatePriceUnitRequest | Price unit details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceUnitsAPI.PricesUnitsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceUnitsAPI.PricesUnitsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesUnitsPost`: DtoCreatePriceUnitResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceUnitsAPI.PricesUnitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricesUnitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DtoCreatePriceUnitRequest**](DtoCreatePriceUnitRequest.md) | Price unit details | 

### Return type

[**DtoCreatePriceUnitResponse**](DtoCreatePriceUnitResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricesUnitsSearchPost

> DtoListPriceUnitsResponse PricesUnitsSearchPost(ctx).Filter(filter).Execute()

List price units by filter



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
	filter := *openapiclient.NewTypesPriceUnitFilter() // TypesPriceUnitFilter | Filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceUnitsAPI.PricesUnitsSearchPost(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceUnitsAPI.PricesUnitsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricesUnitsSearchPost`: DtoListPriceUnitsResponse
	fmt.Fprintf(os.Stdout, "Response from `PriceUnitsAPI.PricesUnitsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPricesUnitsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**TypesPriceUnitFilter**](TypesPriceUnitFilter.md) | Filter | 

### Return type

[**DtoListPriceUnitsResponse**](DtoListPriceUnitsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

