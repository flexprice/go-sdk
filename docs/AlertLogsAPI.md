# \AlertLogsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertSearchPost**](AlertLogsAPI.md#AlertSearchPost) | **Post** /alert/search | List alert logs by filter



## AlertSearchPost

> DtoListAlertLogsResponse AlertSearchPost(ctx).Filter(filter).Execute()

List alert logs by filter



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
	filter := *openapiclient.NewTypesAlertLogFilter() // TypesAlertLogFilter | Filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertLogsAPI.AlertSearchPost(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertLogsAPI.AlertSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertSearchPost`: DtoListAlertLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertLogsAPI.AlertSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**TypesAlertLogFilter**](TypesAlertLogFilter.md) | Filter | 

### Return type

[**DtoListAlertLogsResponse**](DtoListAlertLogsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

