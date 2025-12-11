# \CreditGrantsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditgrantsGet**](CreditGrantsAPI.md#CreditgrantsGet) | **Get** /creditgrants | Get credit grants
[**CreditgrantsIdDelete**](CreditGrantsAPI.md#CreditgrantsIdDelete) | **Delete** /creditgrants/{id} | Delete a credit grant
[**CreditgrantsIdGet**](CreditGrantsAPI.md#CreditgrantsIdGet) | **Get** /creditgrants/{id} | Get a credit grant by ID
[**CreditgrantsIdPut**](CreditGrantsAPI.md#CreditgrantsIdPut) | **Put** /creditgrants/{id} | Update a credit grant
[**CreditgrantsPost**](CreditGrantsAPI.md#CreditgrantsPost) | **Post** /creditgrants | Create a new credit grant
[**PlansIdCreditgrantsGet**](CreditGrantsAPI.md#PlansIdCreditgrantsGet) | **Get** /plans/{id}/creditgrants | Get plan credit grants



## CreditgrantsGet

> DtoListCreditGrantsResponse CreditgrantsGet(ctx).EndTime(endTime).Expand(expand).Limit(limit).Offset(offset).Order(order).PlanIds(planIds).Scope(scope).Sort(sort).StartTime(startTime).Status(status).SubscriptionIds(subscriptionIds).Execute()

Get credit grants



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
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	planIds := []string{"Inner_example"} // []string | Specific filters for credit grants (optional)
	scope := "scope_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	subscriptionIds := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditGrantsAPI.CreditgrantsGet(context.Background()).EndTime(endTime).Expand(expand).Limit(limit).Offset(offset).Order(order).PlanIds(planIds).Scope(scope).Sort(sort).StartTime(startTime).Status(status).SubscriptionIds(subscriptionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditGrantsAPI.CreditgrantsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditgrantsGet`: DtoListCreditGrantsResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditGrantsAPI.CreditgrantsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditgrantsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **planIds** | **[]string** | Specific filters for credit grants | 
 **scope** | **string** |  | 
 **sort** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 
 **subscriptionIds** | **[]string** |  | 

### Return type

[**DtoListCreditGrantsResponse**](DtoListCreditGrantsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditgrantsIdDelete

> map[string]map[string]interface{} CreditgrantsIdDelete(ctx, id).Execute()

Delete a credit grant



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
	id := "id_example" // string | Credit Grant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditGrantsAPI.CreditgrantsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditGrantsAPI.CreditgrantsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditgrantsIdDelete`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CreditGrantsAPI.CreditgrantsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credit Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditgrantsIdDeleteRequest struct via the builder pattern


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


## CreditgrantsIdGet

> DtoCreditGrantResponse CreditgrantsIdGet(ctx, id).Execute()

Get a credit grant by ID



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
	id := "id_example" // string | Credit Grant ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditGrantsAPI.CreditgrantsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditGrantsAPI.CreditgrantsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditgrantsIdGet`: DtoCreditGrantResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditGrantsAPI.CreditgrantsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credit Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditgrantsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoCreditGrantResponse**](DtoCreditGrantResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditgrantsIdPut

> DtoCreditGrantResponse CreditgrantsIdPut(ctx, id).CreditGrant(creditGrant).Execute()

Update a credit grant



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
	id := "id_example" // string | Credit Grant ID
	creditGrant := *openapiclient.NewDtoUpdateCreditGrantRequest() // DtoUpdateCreditGrantRequest | Credit Grant configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditGrantsAPI.CreditgrantsIdPut(context.Background(), id).CreditGrant(creditGrant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditGrantsAPI.CreditgrantsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditgrantsIdPut`: DtoCreditGrantResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditGrantsAPI.CreditgrantsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credit Grant ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditgrantsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creditGrant** | [**DtoUpdateCreditGrantRequest**](DtoUpdateCreditGrantRequest.md) | Credit Grant configuration | 

### Return type

[**DtoCreditGrantResponse**](DtoCreditGrantResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditgrantsPost

> DtoCreditGrantResponse CreditgrantsPost(ctx).CreditGrant(creditGrant).Execute()

Create a new credit grant



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
	creditGrant := *openapiclient.NewDtoCreateCreditGrantRequest(openapiclient.types.CreditGrantCadence("ONETIME"), "Credits_example", "Name_example", openapiclient.types.CreditGrantScope("PLAN")) // DtoCreateCreditGrantRequest | Credit Grant configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditGrantsAPI.CreditgrantsPost(context.Background()).CreditGrant(creditGrant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditGrantsAPI.CreditgrantsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditgrantsPost`: DtoCreditGrantResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditGrantsAPI.CreditgrantsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditgrantsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditGrant** | [**DtoCreateCreditGrantRequest**](DtoCreateCreditGrantRequest.md) | Credit Grant configuration | 

### Return type

[**DtoCreditGrantResponse**](DtoCreditGrantResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlansIdCreditgrantsGet

> DtoListCreditGrantsResponse PlansIdCreditgrantsGet(ctx, id).Execute()

Get plan credit grants



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
	resp, r, err := apiClient.CreditGrantsAPI.PlansIdCreditgrantsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditGrantsAPI.PlansIdCreditgrantsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlansIdCreditgrantsGet`: DtoListCreditGrantsResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditGrantsAPI.PlansIdCreditgrantsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlansIdCreditgrantsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoListCreditGrantsResponse**](DtoListCreditGrantsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

