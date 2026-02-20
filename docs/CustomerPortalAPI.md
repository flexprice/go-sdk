# \CustomerPortalAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortalExternalIdGet**](CustomerPortalAPI.md#PortalExternalIdGet) | **Get** /portal/{external_id} | Create a customer portal session



## PortalExternalIdGet

> DtoPortalSessionResponse PortalExternalIdGet(ctx, externalId).Execute()

Create a customer portal session



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
	externalId := "externalId_example" // string | Customer External ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerPortalAPI.PortalExternalIdGet(context.Background(), externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerPortalAPI.PortalExternalIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortalExternalIdGet`: DtoPortalSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerPortalAPI.PortalExternalIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** | Customer External ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPortalExternalIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoPortalSessionResponse**](DtoPortalSessionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

