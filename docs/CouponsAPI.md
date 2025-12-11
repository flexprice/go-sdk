# \CouponsAPI

All URIs are relative to */v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CouponsGet**](CouponsAPI.md#CouponsGet) | **Get** /coupons | List coupons with filtering
[**CouponsIdDelete**](CouponsAPI.md#CouponsIdDelete) | **Delete** /coupons/{id} | Delete a coupon
[**CouponsIdGet**](CouponsAPI.md#CouponsIdGet) | **Get** /coupons/{id} | Get a coupon by ID
[**CouponsIdPut**](CouponsAPI.md#CouponsIdPut) | **Put** /coupons/{id} | Update a coupon
[**CouponsPost**](CouponsAPI.md#CouponsPost) | **Post** /coupons | Create a new coupon



## CouponsGet

> DtoListCouponsResponse CouponsGet(ctx).CouponIds(couponIds).Expand(expand).Limit(limit).Offset(offset).Order(order).Status(status).Execute()

List coupons with filtering



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
	couponIds := []string{"Inner_example"} // []string |  (optional)
	expand := "expand_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)
	offset := int32(56) // int32 |  (optional)
	order := "order_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.CouponsGet(context.Background()).CouponIds(couponIds).Expand(expand).Limit(limit).Offset(offset).Order(order).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.CouponsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CouponsGet`: DtoListCouponsResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.CouponsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCouponsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponIds** | **[]string** |  | 
 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **order** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**DtoListCouponsResponse**](DtoListCouponsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CouponsIdDelete

> map[string]string CouponsIdDelete(ctx, id).Execute()

Delete a coupon



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
	id := "id_example" // string | Coupon ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.CouponsIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.CouponsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CouponsIdDelete`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.CouponsIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Coupon ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCouponsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CouponsIdGet

> DtoCouponResponse CouponsIdGet(ctx, id).Execute()

Get a coupon by ID



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
	id := "id_example" // string | Coupon ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.CouponsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.CouponsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CouponsIdGet`: DtoCouponResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.CouponsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Coupon ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCouponsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DtoCouponResponse**](DtoCouponResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CouponsIdPut

> DtoCouponResponse CouponsIdPut(ctx, id).Coupon(coupon).Execute()

Update a coupon



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
	id := "id_example" // string | Coupon ID
	coupon := *openapiclient.NewDtoUpdateCouponRequest() // DtoUpdateCouponRequest | Coupon update request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.CouponsIdPut(context.Background(), id).Coupon(coupon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.CouponsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CouponsIdPut`: DtoCouponResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.CouponsIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Coupon ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCouponsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **coupon** | [**DtoUpdateCouponRequest**](DtoUpdateCouponRequest.md) | Coupon update request | 

### Return type

[**DtoCouponResponse**](DtoCouponResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CouponsPost

> DtoCouponResponse CouponsPost(ctx).Coupon(coupon).Execute()

Create a new coupon



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
	coupon := *openapiclient.NewDtoCreateCouponRequest(openapiclient.types.CouponCadence("once"), "Name_example", openapiclient.types.CouponType("fixed")) // DtoCreateCouponRequest | Coupon request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.CouponsPost(context.Background()).Coupon(coupon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.CouponsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CouponsPost`: DtoCouponResponse
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.CouponsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCouponsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **coupon** | [**DtoCreateCouponRequest**](DtoCreateCouponRequest.md) | Coupon request | 

### Return type

[**DtoCouponResponse**](DtoCouponResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

