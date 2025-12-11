# \WalletsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomersIdWalletsGet**](WalletsAPI.md#CustomersIdWalletsGet) | **Get** /customers/{id}/wallets | Get wallets by customer ID
[**CustomersWalletsGet**](WalletsAPI.md#CustomersWalletsGet) | **Get** /customers/wallets | Get Customer Wallets
[**WalletsIdBalanceRealTimeGet**](WalletsAPI.md#WalletsIdBalanceRealTimeGet) | **Get** /wallets/{id}/balance/real-time | Get wallet balance
[**WalletsIdDebitPost**](WalletsAPI.md#WalletsIdDebitPost) | **Post** /wallets/{id}/debit | Debit a wallet
[**WalletsIdGet**](WalletsAPI.md#WalletsIdGet) | **Get** /wallets/{id} | Get wallet by ID
[**WalletsIdPut**](WalletsAPI.md#WalletsIdPut) | **Put** /wallets/{id} | Update a wallet
[**WalletsIdTerminatePost**](WalletsAPI.md#WalletsIdTerminatePost) | **Post** /wallets/{id}/terminate | Terminate a wallet
[**WalletsIdTopUpPost**](WalletsAPI.md#WalletsIdTopUpPost) | **Post** /wallets/{id}/top-up | Top up wallet
[**WalletsIdTransactionsGet**](WalletsAPI.md#WalletsIdTransactionsGet) | **Get** /wallets/{id}/transactions | Get wallet transactions
[**WalletsPost**](WalletsAPI.md#WalletsPost) | **Post** /wallets | Create a new wallet



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


## WalletsIdDebitPost

> DtoWalletResponse WalletsIdDebitPost(ctx, id).Request(request).Execute()

Debit a wallet



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
	request := *openapiclient.NewDtoManualBalanceDebitRequest("IdempotencyKey_example", openapiclient.types.TransactionReason("INVOICE_PAYMENT")) // DtoManualBalanceDebitRequest | Debit wallet request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsIdDebitPost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.WalletsIdDebitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletsIdDebitPost`: DtoWalletResponse
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.WalletsIdDebitPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Wallet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWalletsIdDebitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoManualBalanceDebitRequest**](DtoManualBalanceDebitRequest.md) | Debit wallet request | 

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

> DtoListWalletTransactionsResponse WalletsIdTransactionsGet(ctx, id).CreditsAvailableGt(creditsAvailableGt).EndTime(endTime).Expand(expand).ExpiryDateAfter(expiryDateAfter).ExpiryDateBefore(expiryDateBefore).Id2(id2).Limit(limit).Offset(offset).Order(order).Priority(priority).ReferenceId(referenceId).ReferenceType(referenceType).Sort(sort).StartTime(startTime).Status(status).TransactionReason(transactionReason).TransactionStatus(transactionStatus).Type_(type_).Execute()

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
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	transactionReason := "transactionReason_example" // string |  (optional)
	transactionStatus := "transactionStatus_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.WalletsIdTransactionsGet(context.Background(), id).CreditsAvailableGt(creditsAvailableGt).EndTime(endTime).Expand(expand).ExpiryDateAfter(expiryDateAfter).ExpiryDateBefore(expiryDateBefore).Id2(id2).Limit(limit).Offset(offset).Order(order).Priority(priority).ReferenceId(referenceId).ReferenceType(referenceType).Sort(sort).StartTime(startTime).Status(status).TransactionReason(transactionReason).TransactionStatus(transactionStatus).Type_(type_).Execute()
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
 **sort** | **string** |  | 
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

