# \GroupsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsIdDelete**](GroupsAPI.md#GroupsIdDelete) | **Delete** /groups/{id} | Delete a group
[**GroupsIdGet**](GroupsAPI.md#GroupsIdGet) | **Get** /groups/{id} | Get a group
[**GroupsPost**](GroupsAPI.md#GroupsPost) | **Post** /groups | Create a group
[**GroupsSearchPost**](GroupsAPI.md#GroupsSearchPost) | **Post** /groups/search | Get groups



## GroupsIdDelete

> GroupsIdDelete(ctx, id).Execute()

Delete a group



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
	id := "id_example" // string | Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsIdDeleteRequest struct via the builder pattern


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


## GroupsIdGet

> DtoGroupResponse GroupsIdGet(ctx, id).Execute()

Get a group



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
	id := "id_example" // string | Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsIdGet`: DtoGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoGroupResponse**](DtoGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsPost

> DtoGroupResponse GroupsPost(ctx).Group(group).Execute()

Create a group



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
	group := *openapiclient.NewDtoCreateGroupRequest("EntityType_example", "LookupKey_example", "Name_example") // DtoCreateGroupRequest | Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsPost(context.Background()).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsPost`: DtoGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**DtoCreateGroupRequest**](DtoCreateGroupRequest.md) | Group | 

### Return type

[**DtoGroupResponse**](DtoGroupResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsSearchPost

> DtoListGroupsResponse GroupsSearchPost(ctx).EntityType(entityType).Name(name).LookupKey(lookupKey).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Execute()

Get groups



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
	entityType := "entityType_example" // string | Filter by entity type (e.g., 'price') (optional)
	name := "name_example" // string | Filter by group name (contains search) (optional)
	lookupKey := "lookupKey_example" // string | Filter by lookup key (exact match) (optional)
	limit := int32(56) // int32 | Number of items to return (default: 20) (optional)
	offset := int32(56) // int32 | Number of items to skip (default: 0) (optional)
	sortBy := "sortBy_example" // string | Field to sort by (name, created_at, updated_at) (optional)
	sortOrder := "sortOrder_example" // string | Sort order (asc, desc) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsSearchPost(context.Background()).EntityType(entityType).Name(name).LookupKey(lookupKey).Limit(limit).Offset(offset).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsSearchPost`: DtoListGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityType** | **string** | Filter by entity type (e.g., &#39;price&#39;) | 
 **name** | **string** | Filter by group name (contains search) | 
 **lookupKey** | **string** | Filter by lookup key (exact match) | 
 **limit** | **int32** | Number of items to return (default: 20) | 
 **offset** | **int32** | Number of items to skip (default: 0) | 
 **sortBy** | **string** | Field to sort by (name, created_at, updated_at) | 
 **sortOrder** | **string** | Sort order (asc, desc) | 

### Return type

[**DtoListGroupsResponse**](DtoListGroupsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

