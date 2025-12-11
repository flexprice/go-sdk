# \CreditNotesAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditnotesGet**](CreditNotesAPI.md#CreditnotesGet) | **Get** /creditnotes | List credit notes with filtering
[**CreditnotesIdFinalizePost**](CreditNotesAPI.md#CreditnotesIdFinalizePost) | **Post** /creditnotes/{id}/finalize | Process a draft credit note
[**CreditnotesIdGet**](CreditNotesAPI.md#CreditnotesIdGet) | **Get** /creditnotes/{id} | Get a credit note by ID
[**CreditnotesIdVoidPost**](CreditNotesAPI.md#CreditnotesIdVoidPost) | **Post** /creditnotes/{id}/void | Void a credit note
[**CreditnotesPost**](CreditNotesAPI.md#CreditnotesPost) | **Post** /creditnotes | Create a new credit note



## CreditnotesGet

> DtoListCreditNotesResponse CreditnotesGet(ctx).CreditNoteIds(creditNoteIds).CreditNoteStatus(creditNoteStatus).CreditNoteType(creditNoteType).EndTime(endTime).Expand(expand).InvoiceId(invoiceId).Limit(limit).Offset(offset).Order(order).Sort(sort).StartTime(startTime).Status(status).Execute()

List credit notes with filtering



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
	creditNoteIds := []string{"Inner_example"} // []string |  (optional)
	creditNoteStatus := []string{"CreditNoteStatus_example"} // []string |  (optional)
	creditNoteType := "creditNoteType_example" // string |  (optional)
	endTime := "endTime_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	invoiceId := "invoiceId_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditnotesGet(context.Background()).CreditNoteIds(creditNoteIds).CreditNoteStatus(creditNoteStatus).CreditNoteType(creditNoteType).EndTime(endTime).Expand(expand).InvoiceId(invoiceId).Limit(limit).Offset(offset).Order(order).Sort(sort).StartTime(startTime).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditnotesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditnotesGet`: DtoListCreditNotesResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditnotesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditnotesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditNoteIds** | **[]string** |  | 
 **creditNoteStatus** | **[]string** |  | 
 **creditNoteType** | **string** |  | 
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **invoiceId** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **sort** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**DtoListCreditNotesResponse**](DtoListCreditNotesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditnotesIdFinalizePost

> DtoCreditNoteResponse CreditnotesIdFinalizePost(ctx, id).Execute()

Process a draft credit note



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
	id := "id_example" // string | Credit note ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditnotesIdFinalizePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditnotesIdFinalizePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditnotesIdFinalizePost`: DtoCreditNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditnotesIdFinalizePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credit note ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditnotesIdFinalizePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoCreditNoteResponse**](DtoCreditNoteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditnotesIdGet

> DtoCreditNoteResponse CreditnotesIdGet(ctx, id).Execute()

Get a credit note by ID



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
	id := "id_example" // string | Credit note ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditnotesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditnotesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditnotesIdGet`: DtoCreditNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditnotesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credit note ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditnotesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoCreditNoteResponse**](DtoCreditNoteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditnotesIdVoidPost

> DtoCreditNoteResponse CreditnotesIdVoidPost(ctx, id).Execute()

Void a credit note



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
	id := "id_example" // string | Credit note ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditnotesIdVoidPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditnotesIdVoidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditnotesIdVoidPost`: DtoCreditNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditnotesIdVoidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Credit note ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreditnotesIdVoidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoCreditNoteResponse**](DtoCreditNoteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditnotesPost

> DtoCreditNoteResponse CreditnotesPost(ctx).CreditNote(creditNote).Execute()

Create a new credit note



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
	creditNote := *openapiclient.NewDtoCreateCreditNoteRequest("InvoiceId_example", openapiclient.types.CreditNoteReason("DUPLICATE")) // DtoCreateCreditNoteRequest | Credit note request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditNotesAPI.CreditnotesPost(context.Background()).CreditNote(creditNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditNotesAPI.CreditnotesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditnotesPost`: DtoCreditNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `CreditNotesAPI.CreditnotesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditnotesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **creditNote** | [**DtoCreateCreditNoteRequest**](DtoCreateCreditNoteRequest.md) | Credit note request | 

### Return type

[**DtoCreditNoteResponse**](DtoCreditNoteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

