# \SecretsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretsApiKeysGet**](SecretsAPI.md#SecretsApiKeysGet) | **Get** /secrets/api/keys | List API keys
[**SecretsApiKeysIdDelete**](SecretsAPI.md#SecretsApiKeysIdDelete) | **Delete** /secrets/api/keys/{id} | Delete an API key
[**SecretsApiKeysPost**](SecretsAPI.md#SecretsApiKeysPost) | **Post** /secrets/api/keys | Create a new API key



## SecretsApiKeysGet

> DtoListSecretsResponse SecretsApiKeysGet(ctx).Limit(limit).Offset(offset).Status(status).Execute()

List API keys



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
	limit := int32(56) // int32 | Limit (optional)
	offset := int32(56) // int32 | Offset (optional)
	status := "status_example" // string | Status (published/archived) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretsApiKeysGet(context.Background()).Limit(limit).Offset(offset).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretsApiKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretsApiKeysGet`: DtoListSecretsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretsApiKeysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretsApiKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit | 
 **offset** | **int32** | Offset | 
 **status** | **string** | Status (published/archived) | 

### Return type

[**DtoListSecretsResponse**](DtoListSecretsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecretsApiKeysIdDelete

> SecretsApiKeysIdDelete(ctx, id).Execute()

Delete an API key



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
	id := "id_example" // string | API key ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsAPI.SecretsApiKeysIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretsApiKeysIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | API key ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecretsApiKeysIdDeleteRequest struct via the builder pattern


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


## SecretsApiKeysPost

> DtoCreateAPIKeyResponse SecretsApiKeysPost(ctx).Request(request).Execute()

Create a new API key



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
	request := *openapiclient.NewDtoCreateAPIKeyRequest("Name_example", openapiclient.types.SecretType("private_key")) // DtoCreateAPIKeyRequest | API key creation request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.SecretsApiKeysPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.SecretsApiKeysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecretsApiKeysPost`: DtoCreateAPIKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.SecretsApiKeysPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSecretsApiKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoCreateAPIKeyRequest**](DtoCreateAPIKeyRequest.md) | API key creation request | 

### Return type

[**DtoCreateAPIKeyResponse**](DtoCreateAPIKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

