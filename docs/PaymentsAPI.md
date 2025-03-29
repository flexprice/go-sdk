# \PaymentsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentsGet**](PaymentsAPI.md#PaymentsGet) | **Get** /payments | List payments
[**PaymentsIdDelete**](PaymentsAPI.md#PaymentsIdDelete) | **Delete** /payments/{id} | Delete a payment
[**PaymentsIdGet**](PaymentsAPI.md#PaymentsIdGet) | **Get** /payments/{id} | Get a payment by ID
[**PaymentsIdProcessPost**](PaymentsAPI.md#PaymentsIdProcessPost) | **Post** /payments/{id}/process | Process a payment
[**PaymentsIdPut**](PaymentsAPI.md#PaymentsIdPut) | **Put** /payments/{id} | Update a payment
[**PaymentsPost**](PaymentsAPI.md#PaymentsPost) | **Post** /payments | Create a new payment



## PaymentsGet

> DtoListPaymentsResponse PaymentsGet(ctx).Currency(currency).DestinationId(destinationId).DestinationType(destinationType).EndTime(endTime).Expand(expand).Limit(limit).Offset(offset).Order(order).PaymentGateway(paymentGateway).PaymentIds(paymentIds).PaymentMethodType(paymentMethodType).PaymentStatus(paymentStatus).Sort(sort).StartTime(startTime).Status(status).Execute()

List payments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	currency := "currency_example" // string |  (optional)
	destinationId := "destinationId_example" // string |  (optional)
	destinationType := "destinationType_example" // string |  (optional)
	endTime := "endTime_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	paymentGateway := "paymentGateway_example" // string |  (optional)
	paymentIds := []string{"Inner_example"} // []string |  (optional)
	paymentMethodType := "paymentMethodType_example" // string |  (optional)
	paymentStatus := "paymentStatus_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsGet(context.Background()).Currency(currency).DestinationId(destinationId).DestinationType(destinationType).EndTime(endTime).Expand(expand).Limit(limit).Offset(offset).Order(order).PaymentGateway(paymentGateway).PaymentIds(paymentIds).PaymentMethodType(paymentMethodType).PaymentStatus(paymentStatus).Sort(sort).StartTime(startTime).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsGet`: DtoListPaymentsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** |  | 
 **destinationId** | **string** |  | 
 **destinationType** | **string** |  | 
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **paymentGateway** | **string** |  | 
 **paymentIds** | **[]string** |  | 
 **paymentMethodType** | **string** |  | 
 **paymentStatus** | **string** |  | 
 **sort** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**DtoListPaymentsResponse**](DtoListPaymentsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsIdDelete

> map[string]map[string]interface{} PaymentsIdDelete(ctx, id).Execute()

Delete a payment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	id := "id_example" // string | Payment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsIdDelete`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsIdDeleteRequest struct via the builder pattern


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


## PaymentsIdGet

> DtoPaymentResponse PaymentsIdGet(ctx, id).Execute()

Get a payment by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	id := "id_example" // string | Payment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsIdGet`: DtoPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoPaymentResponse**](DtoPaymentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsIdProcessPost

> DtoPaymentResponse PaymentsIdProcessPost(ctx, id).Execute()

Process a payment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	id := "id_example" // string | Payment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsIdProcessPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsIdProcessPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsIdProcessPost`: DtoPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsIdProcessPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsIdProcessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoPaymentResponse**](DtoPaymentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsIdPut

> DtoPaymentResponse PaymentsIdPut(ctx, id).Payment(payment).Execute()

Update a payment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	id := "id_example" // string | Payment ID
	payment := *openapiclient.NewDtoUpdatePaymentRequest() // DtoUpdatePaymentRequest | Payment configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsIdPut(context.Background(), id).Payment(payment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsIdPut`: DtoPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Payment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payment** | [**DtoUpdatePaymentRequest**](DtoUpdatePaymentRequest.md) | Payment configuration | 

### Return type

[**DtoPaymentResponse**](DtoPaymentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsPost

> DtoPaymentResponse PaymentsPost(ctx).Payment(payment).Execute()

Create a new payment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/flexprice/go-sdk/flexpriceclient"
)

func main() {
	payment := *openapiclient.NewDtoCreatePaymentRequest(float32(123), "Currency_example", "DestinationId_example", openapiclient.types.PaymentDestinationType("INVOICE"), openapiclient.types.PaymentMethodType("CARD")) // DtoCreatePaymentRequest | Payment configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.PaymentsPost(context.Background()).Payment(payment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.PaymentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentsPost`: DtoPaymentResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.PaymentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payment** | [**DtoCreatePaymentRequest**](DtoCreatePaymentRequest.md) | Payment configuration | 

### Return type

[**DtoPaymentResponse**](DtoPaymentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

