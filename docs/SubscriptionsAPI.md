# \SubscriptionsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsAddonDelete**](SubscriptionsAPI.md#SubscriptionsAddonDelete) | **Delete** /subscriptions/addon | Remove addon from subscription
[**SubscriptionsAddonPost**](SubscriptionsAPI.md#SubscriptionsAddonPost) | **Post** /subscriptions/addon | Add addon to subscription
[**SubscriptionsGet**](SubscriptionsAPI.md#SubscriptionsGet) | **Get** /subscriptions | List subscriptions
[**SubscriptionsIdActivatePost**](SubscriptionsAPI.md#SubscriptionsIdActivatePost) | **Post** /subscriptions/{id}/activate | Activate draft subscription
[**SubscriptionsIdAddonsAssociationsGet**](SubscriptionsAPI.md#SubscriptionsIdAddonsAssociationsGet) | **Get** /subscriptions/{id}/addons/associations | Get active addon associations
[**SubscriptionsIdCancelPost**](SubscriptionsAPI.md#SubscriptionsIdCancelPost) | **Post** /subscriptions/{id}/cancel | Cancel subscription
[**SubscriptionsIdChangeExecutePost**](SubscriptionsAPI.md#SubscriptionsIdChangeExecutePost) | **Post** /subscriptions/{id}/change/execute | Execute subscription plan change
[**SubscriptionsIdChangePreviewPost**](SubscriptionsAPI.md#SubscriptionsIdChangePreviewPost) | **Post** /subscriptions/{id}/change/preview | Preview subscription plan change
[**SubscriptionsIdEntitlementsGet**](SubscriptionsAPI.md#SubscriptionsIdEntitlementsGet) | **Get** /subscriptions/{id}/entitlements | Get subscription entitlements
[**SubscriptionsIdGet**](SubscriptionsAPI.md#SubscriptionsIdGet) | **Get** /subscriptions/{id} | Get subscription
[**SubscriptionsIdGrantsUpcomingGet**](SubscriptionsAPI.md#SubscriptionsIdGrantsUpcomingGet) | **Get** /subscriptions/{id}/grants/upcoming | Get upcoming credit grant applications
[**SubscriptionsIdLineitemsPost**](SubscriptionsAPI.md#SubscriptionsIdLineitemsPost) | **Post** /subscriptions/{id}/lineitems | Create subscription line item
[**SubscriptionsIdPausePost**](SubscriptionsAPI.md#SubscriptionsIdPausePost) | **Post** /subscriptions/{id}/pause | Pause a subscription
[**SubscriptionsIdPausesGet**](SubscriptionsAPI.md#SubscriptionsIdPausesGet) | **Get** /subscriptions/{id}/pauses | List all pauses for a subscription
[**SubscriptionsIdPut**](SubscriptionsAPI.md#SubscriptionsIdPut) | **Put** /subscriptions/{id} | Update subscription
[**SubscriptionsIdResumePost**](SubscriptionsAPI.md#SubscriptionsIdResumePost) | **Post** /subscriptions/{id}/resume | Resume a paused subscription
[**SubscriptionsIdV2Get**](SubscriptionsAPI.md#SubscriptionsIdV2Get) | **Get** /subscriptions/{id}/v2 | Get subscription V2
[**SubscriptionsLineitemsIdDelete**](SubscriptionsAPI.md#SubscriptionsLineitemsIdDelete) | **Delete** /subscriptions/lineitems/{id} | Delete subscription line item
[**SubscriptionsLineitemsIdPut**](SubscriptionsAPI.md#SubscriptionsLineitemsIdPut) | **Put** /subscriptions/lineitems/{id} | Update subscription line item
[**SubscriptionsPost**](SubscriptionsAPI.md#SubscriptionsPost) | **Post** /subscriptions | Create subscription
[**SubscriptionsSearchPost**](SubscriptionsAPI.md#SubscriptionsSearchPost) | **Post** /subscriptions/search | List subscriptions by filter
[**SubscriptionsUsagePost**](SubscriptionsAPI.md#SubscriptionsUsagePost) | **Post** /subscriptions/usage | Get usage by subscription
[**V1SubscriptionSchedulesGet**](SubscriptionsAPI.md#V1SubscriptionSchedulesGet) | **Get** /v1/subscription-schedules | List all subscription schedules
[**V1SubscriptionSchedulesIdGet**](SubscriptionsAPI.md#V1SubscriptionSchedulesIdGet) | **Get** /v1/subscription-schedules/{id} | Get subscription schedule
[**V1SubscriptionsSchedulesScheduleIdCancelPost**](SubscriptionsAPI.md#V1SubscriptionsSchedulesScheduleIdCancelPost) | **Post** /v1/subscriptions/schedules/{schedule_id}/cancel | Cancel subscription schedule
[**V1SubscriptionsSubscriptionIdSchedulesGet**](SubscriptionsAPI.md#V1SubscriptionsSubscriptionIdSchedulesGet) | **Get** /v1/subscriptions/{subscription_id}/schedules | List subscription schedules



## SubscriptionsAddonDelete

> DtoSuccessResponse SubscriptionsAddonDelete(ctx).Request(request).Execute()

Remove addon from subscription



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
	request := *openapiclient.NewDtoRemoveAddonRequest("AddonAssociationId_example") // DtoRemoveAddonRequest | Remove Addon Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsAddonDelete(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsAddonDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsAddonDelete`: DtoSuccessResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsAddonDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsAddonDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoRemoveAddonRequest**](DtoRemoveAddonRequest.md) | Remove Addon Request | 

### Return type

[**DtoSuccessResponse**](DtoSuccessResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsAddonPost

> DtoAddonAssociationResponse SubscriptionsAddonPost(ctx).Request(request).Execute()

Add addon to subscription



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
	request := *openapiclient.NewDtoAddAddonRequest("AddonId_example", "SubscriptionId_example") // DtoAddAddonRequest | Add Addon Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsAddonPost(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsAddonPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsAddonPost`: DtoAddonAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsAddonPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsAddonPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**DtoAddAddonRequest**](DtoAddAddonRequest.md) | Add Addon Request | 

### Return type

[**DtoAddonAssociationResponse**](DtoAddonAssociationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGet

> DtoListSubscriptionsResponse SubscriptionsGet(ctx).ActiveAt(activeAt).BillingCadence(billingCadence).BillingPeriod(billingPeriod).CustomerId(customerId).EndTime(endTime).Expand(expand).ExternalCustomerId(externalCustomerId).InvoicingCustomerIds(invoicingCustomerIds).Limit(limit).Offset(offset).Order(order).ParentSubscriptionIds(parentSubscriptionIds).PlanId(planId).StartTime(startTime).Status(status).SubscriptionIds(subscriptionIds).SubscriptionStatus(subscriptionStatus).WithLineItems(withLineItems).Execute()

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
	externalCustomerId := "externalCustomerId_example" // string | ExternalCustomerID filters by external customer ID (optional)
	invoicingCustomerIds := []string{"Inner_example"} // []string | InvoicingCustomerIDs filters by invoicing customer ID (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	parentSubscriptionIds := []string{"Inner_example"} // []string | ParentSubscriptionIDs filters by parent subscription IDs (optional)
	planId := "planId_example" // string | PlanID filters by plan ID (optional)
	startTime := "startTime_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	subscriptionIds := []string{"Inner_example"} // []string |  (optional)
	subscriptionStatus := []string{"SubscriptionStatus_example"} // []string | SubscriptionStatus filters by subscription status (optional)
	withLineItems := true // bool | WithLineItems includes line items in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsGet(context.Background()).ActiveAt(activeAt).BillingCadence(billingCadence).BillingPeriod(billingPeriod).CustomerId(customerId).EndTime(endTime).Expand(expand).ExternalCustomerId(externalCustomerId).InvoicingCustomerIds(invoicingCustomerIds).Limit(limit).Offset(offset).Order(order).ParentSubscriptionIds(parentSubscriptionIds).PlanId(planId).StartTime(startTime).Status(status).SubscriptionIds(subscriptionIds).SubscriptionStatus(subscriptionStatus).WithLineItems(withLineItems).Execute()
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
 **externalCustomerId** | **string** | ExternalCustomerID filters by external customer ID | 
 **invoicingCustomerIds** | **[]string** | InvoicingCustomerIDs filters by invoicing customer ID | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **parentSubscriptionIds** | **[]string** | ParentSubscriptionIDs filters by parent subscription IDs | 
 **planId** | **string** | PlanID filters by plan ID | 
 **startTime** | **string** |  | 
 **status** | **string** |  | 
 **subscriptionIds** | **[]string** |  | 
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


## SubscriptionsIdActivatePost

> DtoSubscriptionResponse SubscriptionsIdActivatePost(ctx, id).Request(request).Execute()

Activate draft subscription



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
	request := *openapiclient.NewDtoActivateDraftSubscriptionRequest("StartDate_example") // DtoActivateDraftSubscriptionRequest | Activate Draft Subscription Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdActivatePost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdActivatePost`: DtoSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdActivatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdActivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoActivateDraftSubscriptionRequest**](DtoActivateDraftSubscriptionRequest.md) | Activate Draft Subscription Request | 

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


## SubscriptionsIdAddonsAssociationsGet

> []DtoAddonAssociationResponse SubscriptionsIdAddonsAssociationsGet(ctx, id).Execute()

Get active addon associations



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
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdAddonsAssociationsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdAddonsAssociationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdAddonsAssociationsGet`: []DtoAddonAssociationResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdAddonsAssociationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdAddonsAssociationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DtoAddonAssociationResponse**](DtoAddonAssociationResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIdCancelPost

> DtoCancelSubscriptionResponse SubscriptionsIdCancelPost(ctx, id).Request(request).Execute()

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
	request := *openapiclient.NewDtoCancelSubscriptionRequest(openapiclient.types.CancellationType("immediate")) // DtoCancelSubscriptionRequest | Cancel Subscription Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdCancelPost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdCancelPost`: DtoCancelSubscriptionResponse
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

 **request** | [**DtoCancelSubscriptionRequest**](DtoCancelSubscriptionRequest.md) | Cancel Subscription Request | 

### Return type

[**DtoCancelSubscriptionResponse**](DtoCancelSubscriptionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIdChangeExecutePost

> DtoSubscriptionChangeExecuteResponse SubscriptionsIdChangeExecutePost(ctx, id).Request(request).Execute()

Execute subscription plan change



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
	request := *openapiclient.NewDtoSubscriptionChangeRequest(openapiclient.types.BillingCadence("RECURRING"), openapiclient.types.BillingCycle("anniversary"), openapiclient.types.BillingPeriod("MONTHLY"), openapiclient.types.ProrationBehavior("create_prorations"), "TargetPlanId_example") // DtoSubscriptionChangeRequest | Subscription change request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdChangeExecutePost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdChangeExecutePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdChangeExecutePost`: DtoSubscriptionChangeExecuteResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdChangeExecutePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdChangeExecutePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoSubscriptionChangeRequest**](DtoSubscriptionChangeRequest.md) | Subscription change request | 

### Return type

[**DtoSubscriptionChangeExecuteResponse**](DtoSubscriptionChangeExecuteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIdChangePreviewPost

> DtoSubscriptionChangePreviewResponse SubscriptionsIdChangePreviewPost(ctx, id).Request(request).Execute()

Preview subscription plan change



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
	request := *openapiclient.NewDtoSubscriptionChangeRequest(openapiclient.types.BillingCadence("RECURRING"), openapiclient.types.BillingCycle("anniversary"), openapiclient.types.BillingPeriod("MONTHLY"), openapiclient.types.ProrationBehavior("create_prorations"), "TargetPlanId_example") // DtoSubscriptionChangeRequest | Subscription change preview request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdChangePreviewPost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdChangePreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdChangePreviewPost`: DtoSubscriptionChangePreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdChangePreviewPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdChangePreviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoSubscriptionChangeRequest**](DtoSubscriptionChangeRequest.md) | Subscription change preview request | 

### Return type

[**DtoSubscriptionChangePreviewResponse**](DtoSubscriptionChangePreviewResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIdEntitlementsGet

> DtoSubscriptionEntitlementsResponse SubscriptionsIdEntitlementsGet(ctx, id).FeatureIds(featureIds).Execute()

Get subscription entitlements



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
	featureIds := []string{"Inner_example"} // []string | Feature IDs to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdEntitlementsGet(context.Background(), id).FeatureIds(featureIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdEntitlementsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdEntitlementsGet`: DtoSubscriptionEntitlementsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdEntitlementsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdEntitlementsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **featureIds** | **[]string** | Feature IDs to filter by | 

### Return type

[**DtoSubscriptionEntitlementsResponse**](DtoSubscriptionEntitlementsResponse.md)

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


## SubscriptionsIdGrantsUpcomingGet

> DtoListCreditGrantApplicationsResponse SubscriptionsIdGrantsUpcomingGet(ctx, id).Execute()

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
	id := "id_example" // string | Subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdGrantsUpcomingGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdGrantsUpcomingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdGrantsUpcomingGet`: DtoListCreditGrantApplicationsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdGrantsUpcomingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdGrantsUpcomingGetRequest struct via the builder pattern


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


## SubscriptionsIdLineitemsPost

> DtoSubscriptionLineItemResponse SubscriptionsIdLineitemsPost(ctx, id).Request(request).Execute()

Create subscription line item



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
	request := *openapiclient.NewDtoCreateSubscriptionLineItemRequest() // DtoCreateSubscriptionLineItemRequest | Create Line Item Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdLineitemsPost(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdLineitemsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdLineitemsPost`: DtoSubscriptionLineItemResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdLineitemsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdLineitemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoCreateSubscriptionLineItemRequest**](DtoCreateSubscriptionLineItemRequest.md) | Create Line Item Request | 

### Return type

[**DtoSubscriptionLineItemResponse**](DtoSubscriptionLineItemResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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

[ApiKeyAuth](../README.md#ApiKeyAuth)

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


## SubscriptionsIdPut

> DtoSubscriptionResponse SubscriptionsIdPut(ctx, id).Request(request).Execute()

Update subscription



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
	request := *openapiclient.NewDtoUpdateSubscriptionRequest() // DtoUpdateSubscriptionRequest | Update Subscription Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdPut(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdPut`: DtoSubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoUpdateSubscriptionRequest**](DtoUpdateSubscriptionRequest.md) | Update Subscription Request | 

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

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsIdV2Get

> DtoSubscriptionResponseV2 SubscriptionsIdV2Get(ctx, id).Expand(expand).Execute()

Get subscription V2



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
	expand := "expand_example" // string | Comma-separated list of fields to expand (e.g., 'subscription_line_items,prices,plan') (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsIdV2Get(context.Background(), id).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsIdV2Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsIdV2Get`: DtoSubscriptionResponseV2
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsIdV2Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsIdV2GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Comma-separated list of fields to expand (e.g., &#39;subscription_line_items,prices,plan&#39;) | 

### Return type

[**DtoSubscriptionResponseV2**](DtoSubscriptionResponseV2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsLineitemsIdDelete

> DtoSubscriptionLineItemResponse SubscriptionsLineitemsIdDelete(ctx, id).Request(request).Execute()

Delete subscription line item



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
	id := "id_example" // string | Line Item ID
	request := *openapiclient.NewDtoDeleteSubscriptionLineItemRequest() // DtoDeleteSubscriptionLineItemRequest | Delete Line Item Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsLineitemsIdDelete(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsLineitemsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsLineitemsIdDelete`: DtoSubscriptionLineItemResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsLineitemsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Line Item ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsLineitemsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoDeleteSubscriptionLineItemRequest**](DtoDeleteSubscriptionLineItemRequest.md) | Delete Line Item Request | 

### Return type

[**DtoSubscriptionLineItemResponse**](DtoSubscriptionLineItemResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsLineitemsIdPut

> DtoSubscriptionLineItemResponse SubscriptionsLineitemsIdPut(ctx, id).Request(request).Execute()

Update subscription line item



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
	id := "id_example" // string | Line Item ID
	request := *openapiclient.NewDtoUpdateSubscriptionLineItemRequest() // DtoUpdateSubscriptionLineItemRequest | Update Line Item Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsLineitemsIdPut(context.Background(), id).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsLineitemsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsLineitemsIdPut`: DtoSubscriptionLineItemResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsLineitemsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Line Item ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsLineitemsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoUpdateSubscriptionLineItemRequest**](DtoUpdateSubscriptionLineItemRequest.md) | Update Line Item Request | 

### Return type

[**DtoSubscriptionLineItemResponse**](DtoSubscriptionLineItemResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

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
	subscription := *openapiclient.NewDtoCreateSubscriptionRequest(openapiclient.types.BillingCadence("RECURRING"), openapiclient.types.BillingPeriod("MONTHLY"), "Currency_example", "PlanId_example") // DtoCreateSubscriptionRequest | Subscription Request

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


## SubscriptionsSearchPost

> DtoListSubscriptionsResponse SubscriptionsSearchPost(ctx).Filter(filter).Execute()

List subscriptions by filter



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
	filter := *openapiclient.NewTypesSubscriptionFilter() // TypesSubscriptionFilter | Filter

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsSearchPost(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsSearchPost`: DtoListSubscriptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**TypesSubscriptionFilter**](TypesSubscriptionFilter.md) | Filter | 

### Return type

[**DtoListSubscriptionsResponse**](DtoListSubscriptionsResponse.md)

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


## V1SubscriptionSchedulesGet

> DtoGetPendingSchedulesResponse V1SubscriptionSchedulesGet(ctx).PendingOnly(pendingOnly).SubscriptionId(subscriptionId).Limit(limit).Offset(offset).Execute()

List all subscription schedules



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
	pendingOnly := true // bool | Filter to pending schedules only (optional)
	subscriptionId := "subscriptionId_example" // string | Filter by subscription ID (optional)
	limit := int32(56) // int32 | Limit results (optional)
	offset := int32(56) // int32 | Offset for pagination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.V1SubscriptionSchedulesGet(context.Background()).PendingOnly(pendingOnly).SubscriptionId(subscriptionId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.V1SubscriptionSchedulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SubscriptionSchedulesGet`: DtoGetPendingSchedulesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.V1SubscriptionSchedulesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SubscriptionSchedulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pendingOnly** | **bool** | Filter to pending schedules only | 
 **subscriptionId** | **string** | Filter by subscription ID | 
 **limit** | **int32** | Limit results | 
 **offset** | **int32** | Offset for pagination | 

### Return type

[**DtoGetPendingSchedulesResponse**](DtoGetPendingSchedulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SubscriptionSchedulesIdGet

> DtoSubscriptionScheduleResponse V1SubscriptionSchedulesIdGet(ctx, id).Execute()

Get subscription schedule



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
	id := "id_example" // string | Schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.V1SubscriptionSchedulesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.V1SubscriptionSchedulesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SubscriptionSchedulesIdGet`: DtoSubscriptionScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.V1SubscriptionSchedulesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SubscriptionSchedulesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoSubscriptionScheduleResponse**](DtoSubscriptionScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SubscriptionsSchedulesScheduleIdCancelPost

> DtoCancelScheduleResponse V1SubscriptionsSchedulesScheduleIdCancelPost(ctx, scheduleId).Request(request).Execute()

Cancel subscription schedule



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
	scheduleId := "scheduleId_example" // string | Schedule ID (optional if using request body)
	request := *openapiclient.NewDtoCancelScheduleRequest() // DtoCancelScheduleRequest | Cancel request (optional if using path parameter) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.V1SubscriptionsSchedulesScheduleIdCancelPost(context.Background(), scheduleId).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.V1SubscriptionsSchedulesScheduleIdCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SubscriptionsSchedulesScheduleIdCancelPost`: DtoCancelScheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.V1SubscriptionsSchedulesScheduleIdCancelPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | Schedule ID (optional if using request body) | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SubscriptionsSchedulesScheduleIdCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**DtoCancelScheduleRequest**](DtoCancelScheduleRequest.md) | Cancel request (optional if using path parameter) | 

### Return type

[**DtoCancelScheduleResponse**](DtoCancelScheduleResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SubscriptionsSubscriptionIdSchedulesGet

> DtoGetPendingSchedulesResponse V1SubscriptionsSubscriptionIdSchedulesGet(ctx, subscriptionId).Execute()

List subscription schedules



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
	subscriptionId := "subscriptionId_example" // string | Subscription ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.V1SubscriptionsSubscriptionIdSchedulesGet(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.V1SubscriptionsSubscriptionIdSchedulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SubscriptionsSubscriptionIdSchedulesGet`: DtoGetPendingSchedulesResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.V1SubscriptionsSubscriptionIdSchedulesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1SubscriptionsSubscriptionIdSchedulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoGetPendingSchedulesResponse**](DtoGetPendingSchedulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

