# \SubscriptionsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsGet**](SubscriptionsAPI.md#SubscriptionsGet) | **Get** /subscriptions | List subscriptions
[**SubscriptionsIdCancelPost**](SubscriptionsAPI.md#SubscriptionsIdCancelPost) | **Post** /subscriptions/{id}/cancel | Cancel subscription
[**SubscriptionsIdGet**](SubscriptionsAPI.md#SubscriptionsIdGet) | **Get** /subscriptions/{id} | Get subscription
[**SubscriptionsIdPausePost**](SubscriptionsAPI.md#SubscriptionsIdPausePost) | **Post** /subscriptions/{id}/pause | Pause a subscription
[**SubscriptionsIdPausesGet**](SubscriptionsAPI.md#SubscriptionsIdPausesGet) | **Get** /subscriptions/{id}/pauses | List all pauses for a subscription
[**SubscriptionsIdResumePost**](SubscriptionsAPI.md#SubscriptionsIdResumePost) | **Post** /subscriptions/{id}/resume | Resume a paused subscription
[**SubscriptionsPost**](SubscriptionsAPI.md#SubscriptionsPost) | **Post** /subscriptions | Create subscription
[**SubscriptionsUsagePost**](SubscriptionsAPI.md#SubscriptionsUsagePost) | **Post** /subscriptions/usage | Get usage by subscription



## SubscriptionsGet

> DtoListSubscriptionsResponse SubscriptionsGet(ctx).ActiveAt(activeAt).BillingCadence(billingCadence).BillingPeriod(billingPeriod).CustomerId(customerId).EndTime(endTime).Expand(expand).IncludeCanceled(includeCanceled).Limit(limit).Offset(offset).Order(order).PlanId(planId).Sort(sort).StartTime(startTime).Status(status).SubscriptionStatus(subscriptionStatus).WithLineItems(withLineItems).Execute()

List subscriptions



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
	activeAt := "activeAt_example" // string | ActiveAt filters subscriptions that are active at the given time (optional)
	billingCadence := []string{"BillingCadence_example"} // []string | BillingCadence filters by billing cadence (optional)
	billingPeriod := []string{"BillingPeriod_example"} // []string | BillingPeriod filters by billing period (optional)
	customerId := "customerId_example" // string | CustomerID filters by customer ID (optional)
	endTime := "endTime_example" // string |  (optional)
	expand := "expand_example" // string |  (optional)
	includeCanceled := true // bool | IncludeCanceled includes canceled subscriptions if true (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	planId := "planId_example" // string | PlanID filters by plan ID (optional)
	sort := "sort_example" // string |  (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	subscriptionStatus := []string{"SubscriptionStatus_example"} // []string | SubscriptionStatus filters by subscription status (optional)
	withLineItems := true // bool | WithLineItems includes line items in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsGet(context.Background()).ActiveAt(activeAt).BillingCadence(billingCadence).BillingPeriod(billingPeriod).CustomerId(customerId).EndTime(endTime).Expand(expand).IncludeCanceled(includeCanceled).Limit(limit).Offset(offset).Order(order).PlanId(planId).Sort(sort).StartTime(startTime).Status(status).SubscriptionStatus(subscriptionStatus).WithLineItems(withLineItems).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsGet`: DtoListSubscriptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeAt** | **string** | ActiveAt filters subscriptions that are active at the given time | 
 **billingCadence** | **[]string** | BillingCadence filters by billing cadence | 
 **billingPeriod** | **[]string** | BillingPeriod filters by billing period | 
 **customerId** | **string** | CustomerID filters by customer ID | 
 **endTime** | **string** |  | 
 **expand** | **string** |  | 
 **includeCanceled** | **bool** | IncludeCanceled includes canceled subscriptions if true | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **planId** | **string** | PlanID filters by plan ID | 
 **sort** | **string** |  | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 
 **subscriptionStatus** | **[]string** | SubscriptionStatus filters by subscription status | 
 **withLineItems** | **bool** | WithLineItems includes line items in the response | 

### Return type

[**DtoListSubscriptionsResponse**](DtoListSubscriptionsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIdCancelPost

> map[string]map[string]interface{} SubscriptionsIdCancelPost(ctx, id).CancelAtPeriodEnd(cancelAtPeriodEnd).Execute()

Cancel subscription



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
	id := "id_example" // string | Subscription ID
	cancelAtPeriodEnd := true // bool | Cancel at period end (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdCancelPost(context.Background(), id).CancelAtPeriodEnd(cancelAtPeriodEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdCancelPost`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelAtPeriodEnd** | **bool** | Cancel at period end | 

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


## SubscriptionsIdGet

> DtoSubscriptionResponse SubscriptionsIdGet(ctx, id).Execute()

Get subscription



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
	id := "id_example" // string | Subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdGet`: DtoSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoSubscriptionResponse**](DtoSubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIdPausePost

> DtoSubscriptionPauseResponse SubscriptionsIdPausePost(ctx, id).Request(request).Execute()

Pause a subscription



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
	id := "id_example" // string | Subscription ID
	request := *openapiclient.NewDtoPauseSubscriptionRequest(openapiclient.types.PauseMode("immediate")) // DtoPauseSubscriptionRequest | Pause subscription request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdPausePost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdPausePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdPausePost`: DtoSubscriptionPauseResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdPausePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdPausePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoPauseSubscriptionRequest**](DtoPauseSubscriptionRequest.md) | Pause subscription request | 

### Return type

[**DtoSubscriptionPauseResponse**](DtoSubscriptionPauseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIdPausesGet

> []DtoListSubscriptionPausesResponse SubscriptionsIdPausesGet(ctx, id).Execute()

List all pauses for a subscription



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
	id := "id_example" // string | Subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdPausesGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdPausesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdPausesGet`: []DtoListSubscriptionPausesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdPausesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdPausesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DtoListSubscriptionPausesResponse**](DtoListSubscriptionPausesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIdResumePost

> DtoSubscriptionPauseResponse SubscriptionsIdResumePost(ctx, id).Request(request).Execute()

Resume a paused subscription



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
	id := "id_example" // string | Subscription ID
	request := *openapiclient.NewDtoResumeSubscriptionRequest(openapiclient.types.ResumeMode("immediate")) // DtoResumeSubscriptionRequest | Resume subscription request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdResumePost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdResumePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdResumePost`: DtoSubscriptionPauseResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdResumePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdResumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoResumeSubscriptionRequest**](DtoResumeSubscriptionRequest.md) | Resume subscription request | 

### Return type

[**DtoSubscriptionPauseResponse**](DtoSubscriptionPauseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsPost

> DtoSubscriptionResponse SubscriptionsPost(ctx).Subscription(subscription).Execute()

Create subscription



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
	subscription := *openapiclient.NewDtoCreateSubscriptionRequest(openapiclient.types.BillingCadence("RECURRING"), openapiclient.types.BillingPeriod("MONTHLY"), int32(123), "Currency_example", "CustomerId_example", "PlanId_example", "StartDate_example") // DtoCreateSubscriptionRequest | Subscription Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsPost(context.Background()).Subscription(subscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsPost`: DtoSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription** | [**DtoCreateSubscriptionRequest**](DtoCreateSubscriptionRequest.md) | Subscription Request | 

### Return type

[**DtoSubscriptionResponse**](DtoSubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsUsagePost

> DtoGetUsageBySubscriptionResponse SubscriptionsUsagePost(ctx).Request(request).Execute()

Get usage by subscription



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
	request := *openapiclient.NewDtoGetUsageBySubscriptionRequest("123") // DtoGetUsageBySubscriptionRequest | Usage request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsUsagePost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsUsagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsUsagePost`: DtoGetUsageBySubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsUsagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsUsagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoGetUsageBySubscriptionRequest**](DtoGetUsageBySubscriptionRequest.md) | Usage request | 

### Return type

[**DtoGetUsageBySubscriptionResponse**](DtoGetUsageBySubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

