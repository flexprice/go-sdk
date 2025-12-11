# \AddonsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddonsGet**](AddonsAPI.md#AddonsGet) | **Get** /addons | List addons
[**AddonsIdDelete**](AddonsAPI.md#AddonsIdDelete) | **Delete** /addons/{id} | Delete addon
[**AddonsIdGet**](AddonsAPI.md#AddonsIdGet) | **Get** /addons/{id} | Get addon
[**AddonsIdPut**](AddonsAPI.md#AddonsIdPut) | **Put** /addons/{id} | Update addon
[**AddonsLookupLookupKeyGet**](AddonsAPI.md#AddonsLookupLookupKeyGet) | **Get** /addons/lookup/{lookup_key} | Get addon by lookup key
[**AddonsPost**](AddonsAPI.md#AddonsPost) | **Post** /addons | Create addon
[**AddonsSearchPost**](AddonsAPI.md#AddonsSearchPost) | **Post** /addons/search | List addons by filter



## AddonsGet

> DtoListAddonsResponse AddonsGet(ctx).AddonIds(addonIds).AddonType(addonType).EndTime(endTime).Expand(expand).Limit(limit).LookupKeys(lookupKeys).Offset(offset).Order(order).StartTime(startTime).Status(status).Execute()

List addons



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
	addonIds := []string{"Inner_example"} // []string |  (optional)
	addonType := "addonType_example" // string |  (optional)
	endTime := "endTime_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	lookupKeys := []string{"Inner_example"} // []string |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsGet(context.Background()).AddonIds(addonIds).AddonType(addonType).EndTime(endTime).Expand(expand).Limit(limit).LookupKeys(lookupKeys).Offset(offset).Order(order).StartTime(startTime).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsGet`: DtoListAddonsResponse
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddonsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addonIds** | **[]string** |  | 
 **addonType** | **string** |  | 
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **lookupKeys** | **[]string** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**DtoListAddonsResponse**](DtoListAddonsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsIdDelete

> map[string]map[string]interface{} AddonsIdDelete(ctx, id).Execute()

Delete addon



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
	id := "id_example" // string | Addon ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsIdDelete`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Addon ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsIdDeleteRequest struct via the builder pattern


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


## AddonsIdGet

> DtoAddonResponse AddonsIdGet(ctx, id).Execute()

Get addon



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
	id := "id_example" // string | Addon ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsIdGet`: DtoAddonResponse
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Addon ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoAddonResponse**](DtoAddonResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsIdPut

> DtoAddonResponse AddonsIdPut(ctx, id).Addon(addon).Execute()

Update addon



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
	id := "id_example" // string | Addon ID
	addon := *openapiclient.NewDtoUpdateAddonRequest() // DtoUpdateAddonRequest | Update Addon Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsIdPut(context.Background(), id).Addon(addon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsIdPut`: DtoAddonResponse
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Addon ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addon** | [**DtoUpdateAddonRequest**](DtoUpdateAddonRequest.md) | Update Addon Request | 

### Return type

[**DtoAddonResponse**](DtoAddonResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsLookupLookupKeyGet

> DtoAddonResponse AddonsLookupLookupKeyGet(ctx, lookupKey).Execute()

Get addon by lookup key



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
	lookupKey := "lookupKey_example" // string | Addon Lookup Key

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsLookupLookupKeyGet(context.Background(), lookupKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsLookupLookupKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsLookupLookupKeyGet`: DtoAddonResponse
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsLookupLookupKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookupKey** | **string** | Addon Lookup Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonsLookupLookupKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoAddonResponse**](DtoAddonResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsPost

> DtoCreateAddonResponse AddonsPost(ctx).Addon(addon).Execute()

Create addon



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
	addon := *openapiclient.NewDtoCreateAddonRequest("LookupKey_example", "Name_example", openapiclient.types.AddonType("onetime")) // DtoCreateAddonRequest | Addon Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsPost(context.Background()).Addon(addon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsPost`: DtoCreateAddonResponse
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddonsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addon** | [**DtoCreateAddonRequest**](DtoCreateAddonRequest.md) | Addon Request | 

### Return type

[**DtoCreateAddonResponse**](DtoCreateAddonResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonsSearchPost

> DtoListAddonsResponse AddonsSearchPost(ctx).Filter(filter).Execute()

List addons by filter



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
	filter := *openapiclient.NewTypesAddonFilter() // TypesAddonFilter | Filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddonsAPI.AddonsSearchPost(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddonsAPI.AddonsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonsSearchPost`: DtoListAddonsResponse
	fmt.Fprintf(os.Stdout, "Response from `AddonsAPI.AddonsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddonsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**TypesAddonFilter**](TypesAddonFilter.md) | Filter | 

### Return type

[**DtoListAddonsResponse**](DtoListAddonsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

