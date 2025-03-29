# \EntitlementsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntitlementsGet**](EntitlementsAPI.md#EntitlementsGet) | **Get** /entitlements | Get entitlements
[**EntitlementsIdDelete**](EntitlementsAPI.md#EntitlementsIdDelete) | **Delete** /entitlements/{id} | Delete an entitlement
[**EntitlementsIdGet**](EntitlementsAPI.md#EntitlementsIdGet) | **Get** /entitlements/{id} | Get an entitlement by ID
[**EntitlementsIdPut**](EntitlementsAPI.md#EntitlementsIdPut) | **Put** /entitlements/{id} | Update an entitlement
[**EntitlementsPost**](EntitlementsAPI.md#EntitlementsPost) | **Post** /entitlements | Create a new entitlement
[**PlansIdEntitlementsGet**](EntitlementsAPI.md#PlansIdEntitlementsGet) | **Get** /plans/{id}/entitlements | Get plan entitlements



## EntitlementsGet

> DtoListEntitlementsResponse EntitlementsGet(ctx).EndTime(endTime).Expand(expand).FeatureIds(featureIds).FeatureType(featureType).IsEnabled(isEnabled).Limit(limit).Offset(offset).Order(order).PlanIds(planIds).Sort(sort).StartTime(startTime).Status(status).Execute()

Get entitlements



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
	featureIds := []string{"Inner_example"} // []string |  (optional)
	featureType := "featureType_example" // string |  (optional)
	isEnabled := true // bool |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	planIds := []string{"Inner_example"} // []string | Specific filters for entitlements (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.EntitlementsGet(context.Background()).EndTime(endTime).Expand(expand).FeatureIds(featureIds).FeatureType(featureType).IsEnabled(isEnabled).Limit(limit).Offset(offset).Order(order).PlanIds(planIds).Sort(sort).StartTime(startTime).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.EntitlementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitlementsGet`: DtoListEntitlementsResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.EntitlementsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **featureIds** | **[]string** |  | 
 **featureType** | **string** |  | 
 **isEnabled** | **bool** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **planIds** | **[]string** | Specific filters for entitlements | 
 **sort** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**DtoListEntitlementsResponse**](DtoListEntitlementsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsIdDelete

> map[string]map[string]interface{} EntitlementsIdDelete(ctx, id).Execute()

Delete an entitlement



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
	id := "id_example" // string | Entitlement ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.EntitlementsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.EntitlementsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitlementsIdDelete`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.EntitlementsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsIdDeleteRequest struct via the builder pattern


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


## EntitlementsIdGet

> DtoEntitlementResponse EntitlementsIdGet(ctx, id).Execute()

Get an entitlement by ID



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
	id := "id_example" // string | Entitlement ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.EntitlementsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.EntitlementsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitlementsIdGet`: DtoEntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.EntitlementsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoEntitlementResponse**](DtoEntitlementResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsIdPut

> DtoEntitlementResponse EntitlementsIdPut(ctx, id).Entitlement(entitlement).Execute()

Update an entitlement



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
	id := "id_example" // string | Entitlement ID
	entitlement := *openapiclient.NewDtoUpdateEntitlementRequest() // DtoUpdateEntitlementRequest | Entitlement configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.EntitlementsIdPut(context.Background(), id).Entitlement(entitlement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.EntitlementsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitlementsIdPut`: DtoEntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.EntitlementsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlement** | [**DtoUpdateEntitlementRequest**](DtoUpdateEntitlementRequest.md) | Entitlement configuration | 

### Return type

[**DtoEntitlementResponse**](DtoEntitlementResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsPost

> DtoEntitlementResponse EntitlementsPost(ctx).Entitlement(entitlement).Execute()

Create a new entitlement



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
	entitlement := *openapiclient.NewDtoCreateEntitlementRequest("FeatureId_example", openapiclient.types.FeatureType("metered")) // DtoCreateEntitlementRequest | Entitlement configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.EntitlementsPost(context.Background()).Entitlement(entitlement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.EntitlementsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitlementsPost`: DtoEntitlementResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.EntitlementsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlement** | [**DtoCreateEntitlementRequest**](DtoCreateEntitlementRequest.md) | Entitlement configuration | 

### Return type

[**DtoEntitlementResponse**](DtoEntitlementResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlansIdEntitlementsGet

> DtoPlanResponse PlansIdEntitlementsGet(ctx, id).Execute()

Get plan entitlements



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
	resp, r, err := apiClient.EntitlementsAPI.PlansIdEntitlementsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.PlansIdEntitlementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlansIdEntitlementsGet`: DtoPlanResponse
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.PlansIdEntitlementsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlansIdEntitlementsGetRequest struct via the builder pattern


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

