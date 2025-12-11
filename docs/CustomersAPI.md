# \CustomersAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomersGet**](CustomersAPI.md#CustomersGet) | **Get** /customers | Get customers
[**CustomersIdDelete**](CustomersAPI.md#CustomersIdDelete) | **Delete** /customers/{id} | Delete a customer
[**CustomersIdEntitlementsGet**](CustomersAPI.md#CustomersIdEntitlementsGet) | **Get** /customers/{id}/entitlements | Get customer entitlements
[**CustomersIdGet**](CustomersAPI.md#CustomersIdGet) | **Get** /customers/{id} | Get a customer
[**CustomersIdGrantsUpcomingGet**](CustomersAPI.md#CustomersIdGrantsUpcomingGet) | **Get** /customers/{id}/grants/upcoming | Get upcoming credit grant applications
[**CustomersIdPut**](CustomersAPI.md#CustomersIdPut) | **Put** /customers/{id} | Update a customer
[**CustomersLookupLookupKeyGet**](CustomersAPI.md#CustomersLookupLookupKeyGet) | **Get** /customers/lookup/{lookup_key} | Get a customer by lookup key
[**CustomersPost**](CustomersAPI.md#CustomersPost) | **Post** /customers | Create a customer
[**CustomersSearchPost**](CustomersAPI.md#CustomersSearchPost) | **Post** /customers/search | List customers by filter
[**CustomersUsageGet**](CustomersAPI.md#CustomersUsageGet) | **Get** /customers/usage | Get customer usage summary



## CustomersGet

> DtoListCustomersResponse CustomersGet(ctx).CustomerIds(customerIds).Email(email).EndTime(endTime).Expand(expand).ExternalId(externalId).ExternalIds(externalIds).Limit(limit).Offset(offset).Order(order).ParentCustomerIds(parentCustomerIds).StartTime(startTime).Status(status).Execute()

Get customers



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
	customerIds := []string{"Inner_example"} // []string |  (optional)
	email := "email_example" // string |  (optional)
	endTime := "endTime_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	externalId := "externalId_example" // string |  (optional)
	externalIds := []string{"Inner_example"} // []string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	parentCustomerIds := []string{"Inner_example"} // []string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersGet(context.Background()).CustomerIds(customerIds).Email(email).EndTime(endTime).Expand(expand).ExternalId(externalId).ExternalIds(externalIds).Limit(limit).Offset(offset).Order(order).ParentCustomerIds(parentCustomerIds).StartTime(startTime).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersGet`: DtoListCustomersResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerIds** | **[]string** |  | 
 **email** | **string** |  | 
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **externalId** | **string** |  | 
 **externalIds** | **[]string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **parentCustomerIds** | **[]string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**DtoListCustomersResponse**](DtoListCustomersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersIdDelete

> CustomersIdDelete(ctx, id).Execute()

Delete a customer



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
	id := "id_example" // string | Customer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomersAPI.CustomersIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersIdDeleteRequest struct via the builder pattern


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


## CustomersIdEntitlementsGet

> DtoCustomerEntitlementsResponse CustomersIdEntitlementsGet(ctx, id).FeatureIds(featureIds).SubscriptionIds(subscriptionIds).Execute()

Get customer entitlements



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
	id := "id_example" // string | Customer ID
	featureIds := []string{"Inner_example"} // []string |  (optional)
	subscriptionIds := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersIdEntitlementsGet(context.Background(), id).FeatureIds(featureIds).SubscriptionIds(subscriptionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersIdEntitlementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersIdEntitlementsGet`: DtoCustomerEntitlementsResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersIdEntitlementsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersIdEntitlementsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureIds** | **[]string** |  | 
 **subscriptionIds** | **[]string** |  | 

### Return type

[**DtoCustomerEntitlementsResponse**](DtoCustomerEntitlementsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersIdGet

> DtoCustomerResponse CustomersIdGet(ctx, id).Execute()

Get a customer



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
	id := "id_example" // string | Customer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersIdGet`: DtoCustomerResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoCustomerResponse**](DtoCustomerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersIdGrantsUpcomingGet

> DtoListCreditGrantApplicationsResponse CustomersIdGrantsUpcomingGet(ctx, id).Execute()

Get upcoming credit grant applications



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
	id := "id_example" // string | Customer ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersIdGrantsUpcomingGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersIdGrantsUpcomingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersIdGrantsUpcomingGet`: DtoListCreditGrantApplicationsResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersIdGrantsUpcomingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersIdGrantsUpcomingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoListCreditGrantApplicationsResponse**](DtoListCreditGrantApplicationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersIdPut

> DtoCustomerResponse CustomersIdPut(ctx, id).Customer(customer).Execute()

Update a customer



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
	id := "id_example" // string | Customer ID
	customer := *openapiclient.NewDtoUpdateCustomerRequest() // DtoUpdateCustomerRequest | Customer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersIdPut(context.Background(), id).Customer(customer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersIdPut`: DtoCustomerResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customer** | [**DtoUpdateCustomerRequest**](DtoUpdateCustomerRequest.md) | Customer | 

### Return type

[**DtoCustomerResponse**](DtoCustomerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersLookupLookupKeyGet

> DtoCustomerResponse CustomersLookupLookupKeyGet(ctx, lookupKey).Execute()

Get a customer by lookup key



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
	lookupKey := "lookupKey_example" // string | Customer Lookup Key (external_id)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersLookupLookupKeyGet(context.Background(), lookupKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersLookupLookupKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersLookupLookupKeyGet`: DtoCustomerResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersLookupLookupKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lookupKey** | **string** | Customer Lookup Key (external_id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersLookupLookupKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoCustomerResponse**](DtoCustomerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersPost

> DtoCustomerResponse CustomersPost(ctx).Customer(customer).Execute()

Create a customer



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
	customer := *openapiclient.NewDtoCreateCustomerRequest("ExternalId_example") // DtoCreateCustomerRequest | Customer

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersPost(context.Background()).Customer(customer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersPost`: DtoCustomerResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customer** | [**DtoCreateCustomerRequest**](DtoCreateCustomerRequest.md) | Customer | 

### Return type

[**DtoCustomerResponse**](DtoCustomerResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersSearchPost

> DtoListCustomersResponse CustomersSearchPost(ctx).Filter(filter).Execute()

List customers by filter



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
	filter := *openapiclient.NewTypesCustomerFilter() // TypesCustomerFilter | Filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersSearchPost(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersSearchPost`: DtoListCustomersResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomersSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**TypesCustomerFilter**](TypesCustomerFilter.md) | Filter | 

### Return type

[**DtoListCustomersResponse**](DtoListCustomersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersUsageGet

> DtoCustomerUsageSummaryResponse CustomersUsageGet(ctx).CustomerId(customerId).CustomerLookupKey(customerLookupKey).FeatureIds(featureIds).FeatureLookupKeys(featureLookupKeys).SubscriptionIds(subscriptionIds).Execute()

Get customer usage summary



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
	customerId := "customerId_example" // string |  (optional)
	customerLookupKey := "customerLookupKey_example" // string |  (optional)
	featureIds := []string{"Inner_example"} // []string |  (optional)
	featureLookupKeys := []string{"Inner_example"} // []string |  (optional)
	subscriptionIds := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CustomersUsageGet(context.Background()).CustomerId(customerId).CustomerLookupKey(customerLookupKey).FeatureIds(featureIds).FeatureLookupKeys(featureLookupKeys).SubscriptionIds(subscriptionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CustomersUsageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersUsageGet`: DtoCustomerUsageSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CustomersUsageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomersUsageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerId** | **string** |  | 
 **customerLookupKey** | **string** |  | 
 **featureIds** | **[]string** |  | 
 **featureLookupKeys** | **[]string** |  | 
 **subscriptionIds** | **[]string** |  | 

### Return type

[**DtoCustomerUsageSummaryResponse**](DtoCustomerUsageSummaryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

