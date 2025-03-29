# \IntegrationsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretsIntegrationsIdDelete**](IntegrationsAPI.md#SecretsIntegrationsIdDelete) | **Delete** /secrets/integrations/{id} | Delete an integration
[**SecretsIntegrationsLinkedGet**](IntegrationsAPI.md#SecretsIntegrationsLinkedGet) | **Get** /secrets/integrations/linked | List linked integrations
[**SecretsIntegrationsProviderGet**](IntegrationsAPI.md#SecretsIntegrationsProviderGet) | **Get** /secrets/integrations/{provider} | Get integration details
[**SecretsIntegrationsProviderPost**](IntegrationsAPI.md#SecretsIntegrationsProviderPost) | **Post** /secrets/integrations/{provider} | Create or update an integration



## SecretsIntegrationsIdDelete

> SecretsIntegrationsIdDelete(ctx, id).Execute()

Delete an integration



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
	id := "id_example" // string | Integration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsAPI.SecretsIntegrationsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.SecretsIntegrationsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Integration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsIntegrationsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsIntegrationsLinkedGet

> DtoLinkedIntegrationsResponse SecretsIntegrationsLinkedGet(ctx).Execute()

List linked integrations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.SecretsIntegrationsLinkedGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.SecretsIntegrationsLinkedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretsIntegrationsLinkedGet`: DtoLinkedIntegrationsResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.SecretsIntegrationsLinkedGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsIntegrationsLinkedGetRequest struct via the builder pattern


### Return type

[**DtoLinkedIntegrationsResponse**](DtoLinkedIntegrationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsIntegrationsProviderGet

> DtoSecretResponse SecretsIntegrationsProviderGet(ctx, provider).Execute()

Get integration details



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
	provider := "provider_example" // string | Integration provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.SecretsIntegrationsProviderGet(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.SecretsIntegrationsProviderGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretsIntegrationsProviderGet`: DtoSecretResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.SecretsIntegrationsProviderGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Integration provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsIntegrationsProviderGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoSecretResponse**](DtoSecretResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsIntegrationsProviderPost

> DtoSecretResponse SecretsIntegrationsProviderPost(ctx, provider).Request(request).Execute()

Create or update an integration



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
	provider := "provider_example" // string | Integration provider
	request := *openapiclient.NewDtoCreateIntegrationRequest(map[string]string{"key": "Inner_example"}, "Name_example", openapiclient.types.SecretProvider("flexprice")) // DtoCreateIntegrationRequest | Integration creation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.SecretsIntegrationsProviderPost(context.Background(), provider).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.SecretsIntegrationsProviderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretsIntegrationsProviderPost`: DtoSecretResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.SecretsIntegrationsProviderPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Integration provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsIntegrationsProviderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoCreateIntegrationRequest**](DtoCreateIntegrationRequest.md) | Integration creation request | 

### Return type

[**DtoSecretResponse**](DtoSecretResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

