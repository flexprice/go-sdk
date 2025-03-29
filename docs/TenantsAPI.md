# \TenantsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenantBillingGet**](TenantsAPI.md#TenantBillingGet) | **Get** /tenant/billing | Get billing usage for the current tenant
[**TenantsIdGet**](TenantsAPI.md#TenantsIdGet) | **Get** /tenants/{id} | Get tenant by ID
[**TenantsPost**](TenantsAPI.md#TenantsPost) | **Post** /tenants | Create a new tenant
[**TenantsUpdatePut**](TenantsAPI.md#TenantsUpdatePut) | **Put** /tenants/update | Update a tenant



## TenantBillingGet

> DtoTenantBillingUsage TenantBillingGet(ctx).Execute()

Get billing usage for the current tenant



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.TenantBillingGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.TenantBillingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantBillingGet`: DtoTenantBillingUsage
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.TenantBillingGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTenantBillingGetRequest struct via the builder pattern


### Return type

[**DtoTenantBillingUsage**](DtoTenantBillingUsage.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsIdGet

> DtoTenantResponse TenantsIdGet(ctx, id).Execute()

Get tenant by ID



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
	id := "id_example" // string | Tenant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.TenantsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.TenantsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantsIdGet`: DtoTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.TenantsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tenant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoTenantResponse**](DtoTenantResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsPost

> DtoTenantResponse TenantsPost(ctx).Request(request).Execute()

Create a new tenant



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
	request := *openapiclient.NewDtoCreateTenantRequest("Name_example") // DtoCreateTenantRequest | Create tenant request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.TenantsPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.TenantsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantsPost`: DtoTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.TenantsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoCreateTenantRequest**](DtoCreateTenantRequest.md) | Create tenant request | 

### Return type

[**DtoTenantResponse**](DtoTenantResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantsUpdatePut

> DtoTenantResponse TenantsUpdatePut(ctx).Request(request).Execute()

Update a tenant



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
	request := *openapiclient.NewDtoUpdateTenantRequest() // DtoUpdateTenantRequest | Update tenant request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.TenantsUpdatePut(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.TenantsUpdatePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TenantsUpdatePut`: DtoTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.TenantsUpdatePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantsUpdatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoUpdateTenantRequest**](DtoUpdateTenantRequest.md) | Update tenant request | 

### Return type

[**DtoTenantResponse**](DtoTenantResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

