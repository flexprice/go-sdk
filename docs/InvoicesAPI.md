# \InvoicesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomersIdInvoicesSummaryGet**](InvoicesAPI.md#CustomersIdInvoicesSummaryGet) | **Get** /customers/{id}/invoices/summary | Get a customer invoice summary
[**InvoicesGet**](InvoicesAPI.md#InvoicesGet) | **Get** /invoices | List invoices
[**InvoicesIdFinalizePost**](InvoicesAPI.md#InvoicesIdFinalizePost) | **Post** /invoices/{id}/finalize | Finalize an invoice
[**InvoicesIdGet**](InvoicesAPI.md#InvoicesIdGet) | **Get** /invoices/{id} | Get an invoice by ID
[**InvoicesIdPaymentAttemptPost**](InvoicesAPI.md#InvoicesIdPaymentAttemptPost) | **Post** /invoices/{id}/payment/attempt | Attempt payment for an invoice
[**InvoicesIdPaymentPut**](InvoicesAPI.md#InvoicesIdPaymentPut) | **Put** /invoices/{id}/payment | Update invoice payment status
[**InvoicesIdPdfGet**](InvoicesAPI.md#InvoicesIdPdfGet) | **Get** /invoices/{id}/pdf | Get PDF for an invoice
[**InvoicesIdVoidPost**](InvoicesAPI.md#InvoicesIdVoidPost) | **Post** /invoices/{id}/void | Void an invoice
[**InvoicesPost**](InvoicesAPI.md#InvoicesPost) | **Post** /invoices | Create a new invoice
[**InvoicesPreviewPost**](InvoicesAPI.md#InvoicesPreviewPost) | **Post** /invoices/preview | Get a preview invoice



## CustomersIdInvoicesSummaryGet

> DtoCustomerMultiCurrencyInvoiceSummary CustomersIdInvoicesSummaryGet(ctx, id).Execute()

Get a customer invoice summary



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
	resp, r, err := apiClient.InvoicesAPI.CustomersIdInvoicesSummaryGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CustomersIdInvoicesSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomersIdInvoicesSummaryGet`: DtoCustomerMultiCurrencyInvoiceSummary
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CustomersIdInvoicesSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Customer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomersIdInvoicesSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoCustomerMultiCurrencyInvoiceSummary**](DtoCustomerMultiCurrencyInvoiceSummary.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesGet

> DtoListInvoicesResponse InvoicesGet(ctx).AmountDueGt(amountDueGt).AmountRemainingGt(amountRemainingGt).CustomerId(customerId).EndTime(endTime).Expand(expand).InvoiceIds(invoiceIds).InvoiceStatus(invoiceStatus).InvoiceType(invoiceType).Limit(limit).Offset(offset).Order(order).PaymentStatus(paymentStatus).Sort(sort).StartTime(startTime).Status(status).SubscriptionId(subscriptionId).Execute()

List invoices



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
	amountDueGt := float32(8.14) // float32 |  (optional)
	amountRemainingGt := float32(8.14) // float32 |  (optional)
	customerId := "customerId_example" // string |  (optional)
	endTime := "endTime_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	invoiceIds := []string{"Inner_example"} // []string |  (optional)
	invoiceStatus := []string{"InvoiceStatus_example"} // []string |  (optional)
	invoiceType := "invoiceType_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	paymentStatus := []string{"PaymentStatus_example"} // []string |  (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	subscriptionId := "subscriptionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesGet(context.Background()).AmountDueGt(amountDueGt).AmountRemainingGt(amountRemainingGt).CustomerId(customerId).EndTime(endTime).Expand(expand).InvoiceIds(invoiceIds).InvoiceStatus(invoiceStatus).InvoiceType(invoiceType).Limit(limit).Offset(offset).Order(order).PaymentStatus(paymentStatus).Sort(sort).StartTime(startTime).Status(status).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesGet`: DtoListInvoicesResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amountDueGt** | **float32** |  | 
 **amountRemainingGt** | **float32** |  | 
 **customerId** | **string** |  | 
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **invoiceIds** | **[]string** |  | 
 **invoiceStatus** | **[]string** |  | 
 **invoiceType** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **paymentStatus** | **[]string** |  | 
 **sort** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 
 **subscriptionId** | **string** |  | 

### Return type

[**DtoListInvoicesResponse**](DtoListInvoicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesIdFinalizePost

> map[string]map[string]interface{} InvoicesIdFinalizePost(ctx, id).Execute()

Finalize an invoice



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
	id := "id_example" // string | Invoice ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesIdFinalizePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesIdFinalizePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesIdFinalizePost`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesIdFinalizePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesIdFinalizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesIdGet

> DtoInvoiceResponse InvoicesIdGet(ctx, id).Execute()

Get an invoice by ID



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
	id := "id_example" // string | Invoice ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesIdGet`: DtoInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoInvoiceResponse**](DtoInvoiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesIdPaymentAttemptPost

> map[string]map[string]interface{} InvoicesIdPaymentAttemptPost(ctx, id).Execute()

Attempt payment for an invoice



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
	id := "id_example" // string | Invoice ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesIdPaymentAttemptPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesIdPaymentAttemptPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesIdPaymentAttemptPost`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesIdPaymentAttemptPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesIdPaymentAttemptPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesIdPaymentPut

> DtoInvoiceResponse InvoicesIdPaymentPut(ctx, id).Request(request).Execute()

Update invoice payment status



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
	id := "id_example" // string | Invoice ID
	request := *openapiclient.NewDtoUpdatePaymentStatusRequest(openapiclient.types.PaymentStatus("PENDING")) // DtoUpdatePaymentStatusRequest | Payment Status Update Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesIdPaymentPut(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesIdPaymentPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesIdPaymentPut`: DtoInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesIdPaymentPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesIdPaymentPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoUpdatePaymentStatusRequest**](DtoUpdatePaymentStatusRequest.md) | Payment Status Update Request | 

### Return type

[**DtoInvoiceResponse**](DtoInvoiceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesIdPdfGet

> *os.File InvoicesIdPdfGet(ctx, id).Url(url).Execute()

Get PDF for an invoice



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
	id := "id_example" // string | Invoice ID
	url := true // bool | Return presigned URL from s3 instead of PDF (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesIdPdfGet(context.Background(), id).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesIdPdfGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesIdPdfGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesIdPdfGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesIdPdfGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **url** | **bool** | Return presigned URL from s3 instead of PDF | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesIdVoidPost

> map[string]map[string]interface{} InvoicesIdVoidPost(ctx, id).Execute()

Void an invoice



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
	id := "id_example" // string | Invoice ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesIdVoidPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesIdVoidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesIdVoidPost`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesIdVoidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Invoice ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesIdVoidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesPost

> DtoInvoiceResponse InvoicesPost(ctx).Invoice(invoice).Execute()

Create a new invoice



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
	invoice := *openapiclient.NewDtoCreateInvoiceRequest(float32(123), "Currency_example", "CustomerId_example") // DtoCreateInvoiceRequest | Invoice details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesPost(context.Background()).Invoice(invoice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesPost`: DtoInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoice** | [**DtoCreateInvoiceRequest**](DtoCreateInvoiceRequest.md) | Invoice details | 

### Return type

[**DtoInvoiceResponse**](DtoInvoiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvoicesPreviewPost

> DtoInvoiceResponse InvoicesPreviewPost(ctx).Request(request).Execute()

Get a preview invoice



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
	request := *openapiclient.NewDtoGetPreviewInvoiceRequest("SubscriptionId_example") // DtoGetPreviewInvoiceRequest | Preview Invoice Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.InvoicesPreviewPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.InvoicesPreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoicesPreviewPost`: DtoInvoiceResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.InvoicesPreviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoicesPreviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoGetPreviewInvoiceRequest**](DtoGetPreviewInvoiceRequest.md) | Preview Invoice Request | 

### Return type

[**DtoInvoiceResponse**](DtoInvoiceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

