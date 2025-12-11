# \WebhooksAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhooksChargebeeTenantIdEnvironmentIdPost**](WebhooksAPI.md#WebhooksChargebeeTenantIdEnvironmentIdPost) | **Post** /webhooks/chargebee/{tenant_id}/{environment_id} | Handle Chargebee webhook events
[**WebhooksHubspotTenantIdEnvironmentIdPost**](WebhooksAPI.md#WebhooksHubspotTenantIdEnvironmentIdPost) | **Post** /webhooks/hubspot/{tenant_id}/{environment_id} | Handle HubSpot webhook events
[**WebhooksQuickbooksTenantIdEnvironmentIdPost**](WebhooksAPI.md#WebhooksQuickbooksTenantIdEnvironmentIdPost) | **Post** /webhooks/quickbooks/{tenant_id}/{environment_id} | Handle QuickBooks webhook events
[**WebhooksRazorpayTenantIdEnvironmentIdPost**](WebhooksAPI.md#WebhooksRazorpayTenantIdEnvironmentIdPost) | **Post** /webhooks/razorpay/{tenant_id}/{environment_id} | Handle Razorpay webhook events
[**WebhooksStripeTenantIdEnvironmentIdPost**](WebhooksAPI.md#WebhooksStripeTenantIdEnvironmentIdPost) | **Post** /webhooks/stripe/{tenant_id}/{environment_id} | Handle Stripe webhook events



## WebhooksChargebeeTenantIdEnvironmentIdPost

> map[string]interface{} WebhooksChargebeeTenantIdEnvironmentIdPost(ctx, tenantId, environmentId).Authorization(authorization).Execute()

Handle Chargebee webhook events



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
	tenantId := "tenantId_example" // string | Tenant ID
	environmentId := "environmentId_example" // string | Environment ID
	authorization := "authorization_example" // string | Basic Auth credentials (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksChargebeeTenantIdEnvironmentIdPost(context.Background(), tenantId, environmentId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksChargebeeTenantIdEnvironmentIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksChargebeeTenantIdEnvironmentIdPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksChargebeeTenantIdEnvironmentIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant ID | 
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksChargebeeTenantIdEnvironmentIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | Basic Auth credentials | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksHubspotTenantIdEnvironmentIdPost

> map[string]interface{} WebhooksHubspotTenantIdEnvironmentIdPost(ctx, tenantId, environmentId).XHubSpotSignatureV3(xHubSpotSignatureV3).Execute()

Handle HubSpot webhook events



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
	tenantId := "tenantId_example" // string | Tenant ID
	environmentId := "environmentId_example" // string | Environment ID
	xHubSpotSignatureV3 := "xHubSpotSignatureV3_example" // string | HubSpot webhook signature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksHubspotTenantIdEnvironmentIdPost(context.Background(), tenantId, environmentId).XHubSpotSignatureV3(xHubSpotSignatureV3).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksHubspotTenantIdEnvironmentIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksHubspotTenantIdEnvironmentIdPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksHubspotTenantIdEnvironmentIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant ID | 
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksHubspotTenantIdEnvironmentIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xHubSpotSignatureV3** | **string** | HubSpot webhook signature | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksQuickbooksTenantIdEnvironmentIdPost

> map[string]interface{} WebhooksQuickbooksTenantIdEnvironmentIdPost(ctx, tenantId, environmentId).IntuitSignature(intuitSignature).Execute()

Handle QuickBooks webhook events



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
	tenantId := "tenantId_example" // string | Tenant ID
	environmentId := "environmentId_example" // string | Environment ID
	intuitSignature := "intuitSignature_example" // string | QuickBooks webhook signature (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksQuickbooksTenantIdEnvironmentIdPost(context.Background(), tenantId, environmentId).IntuitSignature(intuitSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksQuickbooksTenantIdEnvironmentIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksQuickbooksTenantIdEnvironmentIdPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksQuickbooksTenantIdEnvironmentIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant ID | 
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksQuickbooksTenantIdEnvironmentIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **intuitSignature** | **string** | QuickBooks webhook signature | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksRazorpayTenantIdEnvironmentIdPost

> map[string]interface{} WebhooksRazorpayTenantIdEnvironmentIdPost(ctx, tenantId, environmentId).XRazorpaySignature(xRazorpaySignature).Execute()

Handle Razorpay webhook events



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
	tenantId := "tenantId_example" // string | Tenant ID
	environmentId := "environmentId_example" // string | Environment ID
	xRazorpaySignature := "xRazorpaySignature_example" // string | Razorpay webhook signature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksRazorpayTenantIdEnvironmentIdPost(context.Background(), tenantId, environmentId).XRazorpaySignature(xRazorpaySignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksRazorpayTenantIdEnvironmentIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksRazorpayTenantIdEnvironmentIdPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksRazorpayTenantIdEnvironmentIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant ID | 
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksRazorpayTenantIdEnvironmentIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRazorpaySignature** | **string** | Razorpay webhook signature | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhooksStripeTenantIdEnvironmentIdPost

> map[string]interface{} WebhooksStripeTenantIdEnvironmentIdPost(ctx, tenantId, environmentId).StripeSignature(stripeSignature).Execute()

Handle Stripe webhook events



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
	tenantId := "tenantId_example" // string | Tenant ID
	environmentId := "environmentId_example" // string | Environment ID
	stripeSignature := "stripeSignature_example" // string | Stripe webhook signature

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.WebhooksStripeTenantIdEnvironmentIdPost(context.Background(), tenantId, environmentId).StripeSignature(stripeSignature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.WebhooksStripeTenantIdEnvironmentIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WebhooksStripeTenantIdEnvironmentIdPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.WebhooksStripeTenantIdEnvironmentIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant ID | 
**environmentId** | **string** | Environment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksStripeTenantIdEnvironmentIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stripeSignature** | **string** | Stripe webhook signature | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

