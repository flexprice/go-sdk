# \WalletsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomersIdWalletsGet**](WalletsAPI.md#CustomersIdWalletsGet) | **Get** /customers/{id}/wallets | Get wallets by customer ID
[**CustomersWalletsGet**](WalletsAPI.md#CustomersWalletsGet) | **Get** /customers/wallets | Get Customer Wallets
[**WalletsGet**](WalletsAPI.md#WalletsGet) | **Get** /wallets | List wallets
[**WalletsIdBalanceRealTimeGet**](WalletsAPI.md#WalletsIdBalanceRealTimeGet) | **Get** /wallets/{id}/balance/real-time | Get wallet balance
[**WalletsIdGet**](WalletsAPI.md#WalletsIdGet) | **Get** /wallets/{id} | Get wallet by ID
[**WalletsIdPut**](WalletsAPI.md#WalletsIdPut) | **Put** /wallets/{id} | Update a wallet
[**WalletsIdTerminatePost**](WalletsAPI.md#WalletsIdTerminatePost) | **Post** /wallets/{id}/terminate | Terminate a wallet
[**WalletsIdTopUpPost**](WalletsAPI.md#WalletsIdTopUpPost) | **Post** /wallets/{id}/top-up | Top up wallet
[**WalletsIdTransactionsGet**](WalletsAPI.md#WalletsIdTransactionsGet) | **Get** /wallets/{id}/transactions | Get wallet transactions
[**WalletsPost**](WalletsAPI.md#WalletsPost) | **Post** /wallets | Create a new wallet
[**WalletsSearchPost**](WalletsAPI.md#WalletsSearchPost) | **Post** /wallets/search | List wallets by filter
[**WalletsTransactionsSearchPost**](WalletsAPI.md#WalletsTransactionsSearchPost) | **Post** /wallets/transactions/search | List wallet transactions by filter



## CustomersIdWalletsGet

> []DtoWalletResponse CustomersIdWalletsGet(ctx, id).Execute()

Get wallets by customer ID



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
	resp, r, err := apiClient.WalletsAPI.CustomersIdWalletsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CustomersIdWalletsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersIdWalletsGet`: []DtoWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CustomersIdWalletsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersIdWalletsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DtoWalletResponse**](DtoWalletResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomersWalletsGet

> []DtoWalletResponse CustomersWalletsGet(ctx).Id(id).IncludeRealTimeBalance(includeRealTimeBalance).LookupKey(lookupKey).Execute()

Get Customer Wallets



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
	id := "id_example" // string |  (optional)
	includeRealTimeBalance := true // bool |  (optional) (default to false)
	lookupKey := "lookupKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.CustomersWalletsGet(context.Background()).Id(id).IncludeRealTimeBalance(includeRealTimeBalance).LookupKey(lookupKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CustomersWalletsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersWalletsGet`: []DtoWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CustomersWalletsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomersWalletsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **includeRealTimeBalance** | **bool** |  | [default to false]
 **lookupKey** | **string** |  | 

### Return type

[**[]DtoWalletResponse**](DtoWalletResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsGet

> TypesListResponseDtoWalletResponse WalletsGet(ctx).AlertEnabled(alertEnabled).Expand(expand).Limit(limit).Offset(offset).Order(order).Sort(sort).Status(status).WalletIds(walletIds).Execute()

List wallets



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
	alertEnabled := true // bool |  (optional)
	expand := "expand_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	walletIds := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsGet(context.Background()).AlertEnabled(alertEnabled).Expand(expand).Limit(limit).Offset(offset).Order(order).Sort(sort).Status(status).WalletIds(walletIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsGet`: TypesListResponseDtoWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertEnabled** | **bool** |  | 
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **sort** | **string** |  | 
 **status** | **string** |  | 
 **walletIds** | **[]string** |  | 

### Return type

[**TypesListResponseDtoWalletResponse**](TypesListResponseDtoWalletResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsIdBalanceRealTimeGet

> DtoWalletBalanceResponse WalletsIdBalanceRealTimeGet(ctx, id).Execute()

Get wallet balance



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
	id := "id_example" // string | Wallet ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsIdBalanceRealTimeGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsIdBalanceRealTimeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsIdBalanceRealTimeGet`: DtoWalletBalanceResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsIdBalanceRealTimeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWalletsIdBalanceRealTimeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoWalletBalanceResponse**](DtoWalletBalanceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsIdGet

> DtoWalletResponse WalletsIdGet(ctx, id).Execute()

Get wallet by ID



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
	id := "id_example" // string | Wallet ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsIdGet`: DtoWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWalletsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoWalletResponse**](DtoWalletResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsIdPut

> DtoWalletResponse WalletsIdPut(ctx, id).Request(request).Execute()

Update a wallet



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
	id := "id_example" // string | Wallet ID
	request := *openapiclient.NewDtoUpdateWalletRequest() // DtoUpdateWalletRequest | Update wallet request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsIdPut(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsIdPut`: DtoWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWalletsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoUpdateWalletRequest**](DtoUpdateWalletRequest.md) | Update wallet request | 

### Return type

[**DtoWalletResponse**](DtoWalletResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsIdTerminatePost

> DtoWalletResponse WalletsIdTerminatePost(ctx, id).Execute()

Terminate a wallet



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
	id := "id_example" // string | Wallet ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsIdTerminatePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsIdTerminatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsIdTerminatePost`: DtoWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsIdTerminatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWalletsIdTerminatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoWalletResponse**](DtoWalletResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsIdTopUpPost

> DtoTopUpWalletResponse WalletsIdTopUpPost(ctx, id).Request(request).Execute()

Top up wallet



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
	id := "id_example" // string | Wallet ID
	request := *openapiclient.NewDtoTopUpWalletRequest(openapiclient.types.TransactionReason("INVOICE_PAYMENT")) // DtoTopUpWalletRequest | Top up request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsIdTopUpPost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsIdTopUpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsIdTopUpPost`: DtoTopUpWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsIdTopUpPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWalletsIdTopUpPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoTopUpWalletRequest**](DtoTopUpWalletRequest.md) | Top up request | 

### Return type

[**DtoTopUpWalletResponse**](DtoTopUpWalletResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsIdTransactionsGet

> DtoListWalletTransactionsResponse WalletsIdTransactionsGet(ctx, id).CreatedBy(createdBy).CreditsAvailableGt(creditsAvailableGt).EndTime(endTime).Expand(expand).ExpiryDateAfter(expiryDateAfter).ExpiryDateBefore(expiryDateBefore).Id2(id2).Limit(limit).Offset(offset).Order(order).Priority(priority).ReferenceId(referenceId).ReferenceType(referenceType).StartTime(startTime).Status(status).TransactionReason(transactionReason).TransactionStatus(transactionStatus).Type_(type_).Execute()

Get wallet transactions



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
	id := "id_example" // string | Wallet ID
	createdBy := "createdBy_example" // string |  (optional)
	creditsAvailableGt := float32(8.14) // float32 |  (optional)
	endTime := "endTime_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	expiryDateAfter := "expiryDateAfter_example" // string |  (optional)
	expiryDateBefore := "expiryDateBefore_example" // string |  (optional)
	id2 := "id_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	priority := int32(56) // int32 |  (optional)
	referenceId := "referenceId_example" // string |  (optional)
	referenceType := "referenceType_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	transactionReason := "transactionReason_example" // string |  (optional)
	transactionStatus := "transactionStatus_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsIdTransactionsGet(context.Background(), id).CreatedBy(createdBy).CreditsAvailableGt(creditsAvailableGt).EndTime(endTime).Expand(expand).ExpiryDateAfter(expiryDateAfter).ExpiryDateBefore(expiryDateBefore).Id2(id2).Limit(limit).Offset(offset).Order(order).Priority(priority).ReferenceId(referenceId).ReferenceType(referenceType).StartTime(startTime).Status(status).TransactionReason(transactionReason).TransactionStatus(transactionStatus).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsIdTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsIdTransactionsGet`: DtoListWalletTransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsIdTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWalletsIdTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBy** | **string** |  | 
 **creditsAvailableGt** | **float32** |  | 
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **expiryDateAfter** | **string** |  | 
 **expiryDateBefore** | **string** |  | 
 **id2** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **priority** | **int32** |  | 
 **referenceId** | **string** |  | 
 **referenceType** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 
 **transactionReason** | **string** |  | 
 **transactionStatus** | **string** |  | 
 **type_** | **string** |  | 

### Return type

[**DtoListWalletTransactionsResponse**](DtoListWalletTransactionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsPost

> DtoWalletResponse WalletsPost(ctx).Request(request).Execute()

Create a new wallet



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
	request := *openapiclient.NewDtoCreateWalletRequest("Currency_example") // DtoCreateWalletRequest | Create wallet request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsPost`: DtoWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoCreateWalletRequest**](DtoCreateWalletRequest.md) | Create wallet request | 

### Return type

[**DtoWalletResponse**](DtoWalletResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsSearchPost

> TypesListResponseDtoWalletResponse WalletsSearchPost(ctx).Filter(filter).Execute()

List wallets by filter



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
	filter := *openapiclient.NewTypesWalletFilter() // TypesWalletFilter | Filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsSearchPost(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsSearchPost`: TypesListResponseDtoWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**TypesWalletFilter**](TypesWalletFilter.md) | Filter | 

### Return type

[**TypesListResponseDtoWalletResponse**](TypesListResponseDtoWalletResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletsTransactionsSearchPost

> DtoListWalletTransactionsResponse WalletsTransactionsSearchPost(ctx).Expand(expand).Filter(filter).Execute()

List wallet transactions by filter



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
	expand := "expand_example" // string | Expand fields (e.g., customer,created_by_user,wallet) (optional)
	filter := *openapiclient.NewTypesWalletTransactionFilter() // TypesWalletTransactionFilter | Filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsTransactionsSearchPost(context.Background()).Expand(expand).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsTransactionsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsTransactionsSearchPost`: DtoListWalletTransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsTransactionsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletsTransactionsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | Expand fields (e.g., customer,created_by_user,wallet) | 
 **filter** | [**TypesWalletTransactionFilter**](TypesWalletTransactionFilter.md) | Filter | 

### Return type

[**DtoListWalletTransactionsResponse**](DtoListWalletTransactionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

