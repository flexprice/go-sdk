# \UsersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersMeGet**](UsersAPI.md#UsersMeGet) | **Get** /users/me | Get user info
[**UsersPost**](UsersAPI.md#UsersPost) | **Post** /users | Create service account
[**UsersSearchPost**](UsersAPI.md#UsersSearchPost) | **Post** /users/search | List users with filters



## UsersMeGet

> DtoUserResponse UsersMeGet(ctx).Execute()

Get user info



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
	resp, r, err := apiClient.UsersAPI.UsersMeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersMeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersMeGet`: DtoUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersMeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsersMeGetRequest struct via the builder pattern


### Return type

[**DtoUserResponse**](DtoUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPost

> DtoUserResponse UsersPost(ctx).Request(request).Execute()

Create service account



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
	request := *openapiclient.NewDtoCreateUserRequest([]string{"Roles_example"}, openapiclient.types.UserType("user")) // DtoCreateUserRequest | Create service account request (type must be 'service_account', roles are required)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPost`: DtoUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoCreateUserRequest**](DtoCreateUserRequest.md) | Create service account request (type must be &#39;service_account&#39;, roles are required) | 

### Return type

[**DtoUserResponse**](DtoUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersSearchPost

> DtoListUsersResponse UsersSearchPost(ctx).Filter(filter).Execute()

List users with filters



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
	filter := *openapiclient.NewTypesUserFilter() // TypesUserFilter | Filter parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersSearchPost(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersSearchPost`: DtoListUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**TypesUserFilter**](TypesUserFilter.md) | Filter parameters | 

### Return type

[**DtoListUsersResponse**](DtoListUsersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

