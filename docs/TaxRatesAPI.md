# \TaxRatesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaxesRatesGet**](TaxRatesAPI.md#TaxesRatesGet) | **Get** /taxes/rates | Get tax rates
[**TaxesRatesIdDelete**](TaxRatesAPI.md#TaxesRatesIdDelete) | **Delete** /taxes/rates/{id} | Delete a tax rate
[**TaxesRatesIdGet**](TaxRatesAPI.md#TaxesRatesIdGet) | **Get** /taxes/rates/{id} | Get a tax rate
[**TaxesRatesIdPut**](TaxRatesAPI.md#TaxesRatesIdPut) | **Put** /taxes/rates/{id} | Update a tax rate
[**TaxesRatesPost**](TaxRatesAPI.md#TaxesRatesPost) | **Post** /taxes/rates | Create a tax rate



## TaxesRatesGet

> []DtoTaxRateResponse TaxesRatesGet(ctx).EndTime(endTime).Expand(expand).Limit(limit).Offset(offset).Order(order).Scope(scope).StartTime(startTime).Status(status).TaxrateCodes(taxrateCodes).TaxrateIds(taxrateIds).Execute()

Get tax rates



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
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	scope := "scope_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	taxrateCodes := []string{"Inner_example"} // []string |  (optional)
	taxrateIds := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxRatesAPI.TaxesRatesGet(context.Background()).EndTime(endTime).Expand(expand).Limit(limit).Offset(offset).Order(order).Scope(scope).StartTime(startTime).Status(status).TaxrateCodes(taxrateCodes).TaxrateIds(taxrateIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRatesAPI.TaxesRatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxesRatesGet`: []DtoTaxRateResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxRatesAPI.TaxesRatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxesRatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **scope** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 
 **taxrateCodes** | **[]string** |  | 
 **taxrateIds** | **[]string** |  | 

### Return type

[**[]DtoTaxRateResponse**](DtoTaxRateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxesRatesIdDelete

> TaxesRatesIdDelete(ctx, id).Execute()

Delete a tax rate



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
	id := "id_example" // string | Tax rate ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaxRatesAPI.TaxesRatesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRatesAPI.TaxesRatesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tax rate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxesRatesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxesRatesIdGet

> DtoTaxRateResponse TaxesRatesIdGet(ctx, id).Execute()

Get a tax rate



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
	id := "id_example" // string | Tax rate ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxRatesAPI.TaxesRatesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRatesAPI.TaxesRatesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxesRatesIdGet`: DtoTaxRateResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxRatesAPI.TaxesRatesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tax rate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxesRatesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoTaxRateResponse**](DtoTaxRateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxesRatesIdPut

> DtoTaxRateResponse TaxesRatesIdPut(ctx, id).TaxRate(taxRate).Execute()

Update a tax rate



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
	id := "id_example" // string | Tax rate ID
	taxRate := *openapiclient.NewDtoUpdateTaxRateRequest() // DtoUpdateTaxRateRequest | Tax rate to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxRatesAPI.TaxesRatesIdPut(context.Background(), id).TaxRate(taxRate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRatesAPI.TaxesRatesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxesRatesIdPut`: DtoTaxRateResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxRatesAPI.TaxesRatesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tax rate ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxesRatesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxRate** | [**DtoUpdateTaxRateRequest**](DtoUpdateTaxRateRequest.md) | Tax rate to update | 

### Return type

[**DtoTaxRateResponse**](DtoTaxRateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxesRatesPost

> DtoTaxRateResponse TaxesRatesPost(ctx).TaxRate(taxRate).Execute()

Create a tax rate



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
	taxRate := *openapiclient.NewDtoCreateTaxRateRequest("Code_example", "Name_example") // DtoCreateTaxRateRequest | Tax rate to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxRatesAPI.TaxesRatesPost(context.Background()).TaxRate(taxRate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRatesAPI.TaxesRatesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxesRatesPost`: DtoTaxRateResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxRatesAPI.TaxesRatesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxesRatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxRate** | [**DtoCreateTaxRateRequest**](DtoCreateTaxRateRequest.md) | Tax rate to create | 

### Return type

[**DtoTaxRateResponse**](DtoTaxRateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

