# \TaxAssociationsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaxesAssociationsGet**](TaxAssociationsAPI.md#TaxesAssociationsGet) | **Get** /taxes/associations | List tax associations
[**TaxesAssociationsIdDelete**](TaxAssociationsAPI.md#TaxesAssociationsIdDelete) | **Delete** /taxes/associations/{id} | Delete tax association
[**TaxesAssociationsIdGet**](TaxAssociationsAPI.md#TaxesAssociationsIdGet) | **Get** /taxes/associations/{id} | Get Tax Association
[**TaxesAssociationsIdPut**](TaxAssociationsAPI.md#TaxesAssociationsIdPut) | **Put** /taxes/associations/{id} | Update tax association
[**TaxesAssociationsPost**](TaxAssociationsAPI.md#TaxesAssociationsPost) | **Post** /taxes/associations | Create Tax Association



## TaxesAssociationsGet

> DtoListTaxAssociationsResponse TaxesAssociationsGet(ctx).TaxAssociation(taxAssociation).Execute()

List tax associations



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
	taxAssociation := *openapiclient.NewTypesTaxAssociationFilter() // TypesTaxAssociationFilter | Tax Association Filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAssociationsAPI.TaxesAssociationsGet(context.Background()).TaxAssociation(taxAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAssociationsAPI.TaxesAssociationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxesAssociationsGet`: DtoListTaxAssociationsResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxAssociationsAPI.TaxesAssociationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxesAssociationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxAssociation** | [**TypesTaxAssociationFilter**](TypesTaxAssociationFilter.md) | Tax Association Filter | 

### Return type

[**DtoListTaxAssociationsResponse**](DtoListTaxAssociationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxesAssociationsIdDelete

> DtoTaxAssociationResponse TaxesAssociationsIdDelete(ctx, id).Execute()

Delete tax association



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
	id := "id_example" // string | Tax Config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAssociationsAPI.TaxesAssociationsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAssociationsAPI.TaxesAssociationsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxesAssociationsIdDelete`: DtoTaxAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxAssociationsAPI.TaxesAssociationsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tax Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxesAssociationsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoTaxAssociationResponse**](DtoTaxAssociationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxesAssociationsIdGet

> DtoTaxAssociationResponse TaxesAssociationsIdGet(ctx, id).Execute()

Get Tax Association



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
	id := "id_example" // string | Tax Config ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAssociationsAPI.TaxesAssociationsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAssociationsAPI.TaxesAssociationsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxesAssociationsIdGet`: DtoTaxAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxAssociationsAPI.TaxesAssociationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tax Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxesAssociationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoTaxAssociationResponse**](DtoTaxAssociationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxesAssociationsIdPut

> DtoTaxAssociationResponse TaxesAssociationsIdPut(ctx, id).TaxConfig(taxConfig).Execute()

Update tax association



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
	id := "id_example" // string | Tax Config ID
	taxConfig := *openapiclient.NewDtoTaxAssociationUpdateRequest() // DtoTaxAssociationUpdateRequest | Tax Config Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAssociationsAPI.TaxesAssociationsIdPut(context.Background(), id).TaxConfig(taxConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAssociationsAPI.TaxesAssociationsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxesAssociationsIdPut`: DtoTaxAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxAssociationsAPI.TaxesAssociationsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tax Config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxesAssociationsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxConfig** | [**DtoTaxAssociationUpdateRequest**](DtoTaxAssociationUpdateRequest.md) | Tax Config Request | 

### Return type

[**DtoTaxAssociationResponse**](DtoTaxAssociationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxesAssociationsPost

> DtoTaxAssociationResponse TaxesAssociationsPost(ctx).TaxConfig(taxConfig).Execute()

Create Tax Association



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
	taxConfig := *openapiclient.NewDtoCreateTaxAssociationRequest("EntityId_example", openapiclient.types.TaxRateEntityType("customer"), "TaxRateCode_example") // DtoCreateTaxAssociationRequest | Tax Config Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxAssociationsAPI.TaxesAssociationsPost(context.Background()).TaxConfig(taxConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxAssociationsAPI.TaxesAssociationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TaxesAssociationsPost`: DtoTaxAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `TaxAssociationsAPI.TaxesAssociationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxesAssociationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxConfig** | [**DtoCreateTaxAssociationRequest**](DtoCreateTaxAssociationRequest.md) | Tax Config Request | 

### Return type

[**DtoTaxAssociationResponse**](DtoTaxAssociationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

