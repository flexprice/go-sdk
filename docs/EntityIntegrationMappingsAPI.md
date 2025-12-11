# \EntityIntegrationMappingsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntityIntegrationMappingsGet**](EntityIntegrationMappingsAPI.md#EntityIntegrationMappingsGet) | **Get** /entity-integration-mappings | List entity integration mappings
[**EntityIntegrationMappingsIdDelete**](EntityIntegrationMappingsAPI.md#EntityIntegrationMappingsIdDelete) | **Delete** /entity-integration-mappings/{id} | Delete entity integration mapping
[**EntityIntegrationMappingsIdGet**](EntityIntegrationMappingsAPI.md#EntityIntegrationMappingsIdGet) | **Get** /entity-integration-mappings/{id} | Get entity integration mapping
[**EntityIntegrationMappingsPost**](EntityIntegrationMappingsAPI.md#EntityIntegrationMappingsPost) | **Post** /entity-integration-mappings | Create entity integration mapping



## EntityIntegrationMappingsGet

> DtoListEntityIntegrationMappingsResponse EntityIntegrationMappingsGet(ctx).EntityId(entityId).EntityType(entityType).ProviderType(providerType).ProviderEntityId(providerEntityId).Limit(limit).Offset(offset).Execute()

List entity integration mappings



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
	entityId := "entityId_example" // string | Filter by FlexPrice entity ID (optional)
	entityType := "entityType_example" // string | Filter by entity type (optional)
	providerType := "providerType_example" // string | Filter by provider type (optional)
	providerEntityId := "providerEntityId_example" // string | Filter by provider entity ID (optional)
	limit := int32(56) // int32 | Number of results to return (default: 20, max: 100) (optional)
	offset := int32(56) // int32 | Pagination offset (default: 0) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntityIntegrationMappingsAPI.EntityIntegrationMappingsGet(context.Background()).EntityId(entityId).EntityType(entityType).ProviderType(providerType).ProviderEntityId(providerEntityId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntityIntegrationMappingsAPI.EntityIntegrationMappingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntityIntegrationMappingsGet`: DtoListEntityIntegrationMappingsResponse
	fmt.Fprintf(os.Stdout, "Response from `EntityIntegrationMappingsAPI.EntityIntegrationMappingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntityIntegrationMappingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityId** | **string** | Filter by FlexPrice entity ID | 
 **entityType** | **string** | Filter by entity type | 
 **providerType** | **string** | Filter by provider type | 
 **providerEntityId** | **string** | Filter by provider entity ID | 
 **limit** | **int32** | Number of results to return (default: 20, max: 100) | 
 **offset** | **int32** | Pagination offset (default: 0) | 

### Return type

[**DtoListEntityIntegrationMappingsResponse**](DtoListEntityIntegrationMappingsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntityIntegrationMappingsIdDelete

> EntityIntegrationMappingsIdDelete(ctx, id).Execute()

Delete entity integration mapping



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
	id := "id_example" // string | Entity integration mapping ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntityIntegrationMappingsAPI.EntityIntegrationMappingsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntityIntegrationMappingsAPI.EntityIntegrationMappingsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Entity integration mapping ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntityIntegrationMappingsIdDeleteRequest struct via the builder pattern


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


## EntityIntegrationMappingsIdGet

> DtoEntityIntegrationMappingResponse EntityIntegrationMappingsIdGet(ctx, id).Execute()

Get entity integration mapping



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
	id := "id_example" // string | Entity integration mapping ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntityIntegrationMappingsAPI.EntityIntegrationMappingsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntityIntegrationMappingsAPI.EntityIntegrationMappingsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntityIntegrationMappingsIdGet`: DtoEntityIntegrationMappingResponse
	fmt.Fprintf(os.Stdout, "Response from `EntityIntegrationMappingsAPI.EntityIntegrationMappingsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Entity integration mapping ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntityIntegrationMappingsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoEntityIntegrationMappingResponse**](DtoEntityIntegrationMappingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntityIntegrationMappingsPost

> DtoEntityIntegrationMappingResponse EntityIntegrationMappingsPost(ctx).EntityIntegrationMapping(entityIntegrationMapping).Execute()

Create entity integration mapping



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
	entityIntegrationMapping := *openapiclient.NewDtoCreateEntityIntegrationMappingRequest("EntityId_example", openapiclient.types.IntegrationEntityType("customer"), "ProviderEntityId_example", "ProviderType_example") // DtoCreateEntityIntegrationMappingRequest | Entity integration mapping data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntityIntegrationMappingsAPI.EntityIntegrationMappingsPost(context.Background()).EntityIntegrationMapping(entityIntegrationMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntityIntegrationMappingsAPI.EntityIntegrationMappingsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntityIntegrationMappingsPost`: DtoEntityIntegrationMappingResponse
	fmt.Fprintf(os.Stdout, "Response from `EntityIntegrationMappingsAPI.EntityIntegrationMappingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntityIntegrationMappingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityIntegrationMapping** | [**DtoCreateEntityIntegrationMappingRequest**](DtoCreateEntityIntegrationMappingRequest.md) | Entity integration mapping data | 

### Return type

[**DtoEntityIntegrationMappingResponse**](DtoEntityIntegrationMappingResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

