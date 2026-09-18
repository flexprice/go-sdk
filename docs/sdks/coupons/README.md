# Coupons

## Overview

### Available Operations

* [CreateCoupon](#createcoupon) - Create coupon
* [GetCouponByCode](#getcouponbycode) - Get coupon by code
* [QueryCoupon](#querycoupon) - Query coupons
* [GetCoupon](#getcoupon) - Get coupon
* [UpdateCoupon](#updatecoupon) - Update coupon
* [DeleteCoupon](#deletecoupon) - Delete coupon

## CreateCoupon

Use when creating a discount (e.g. promo code or referral). Ideal for percent or fixed value, with optional validity and usage limits.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCoupon" method="post" path="/coupons" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Coupons.CreateCoupon(ctx, types.CreateCouponRequest{
        Cadence: types.CouponCadenceRepeated,
        Name: "<value>",
        Type: types.CouponTypePercentage,
    }, dtos.CreateCouponSecurity{
        Option1: &dtos.CreateCouponSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CouponResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [types.CreateCouponRequest](../../models/types/createcouponrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `security`                                                             | [dtos.CreateCouponSecurity](../../models/dtos/createcouponsecurity.md) | :heavy_check_mark:                                                     | The security requirements to use for the request.                      |
| `opts`                                                                 | [][dtos.Option](../../models/dtos/option.md)                           | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*dtos.CreateCouponResponse](../../models/dtos/createcouponresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 401, 403, 404   | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCouponByCode

Use when resolving a coupon by promo code (e.g. checkout or validation).

### Example Usage

<!-- UsageSnippet language="go" operationID="getCouponByCode" method="get" path="/coupons/code/{code}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Coupons.GetCouponByCode(ctx, dtos.GetCouponByCodeSecurity{
        Option1: &dtos.GetCouponByCodeSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CouponResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `security`                                                                   | [dtos.GetCouponByCodeSecurity](../../models/dtos/getcouponbycodesecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |
| `code`                                                                       | `string`                                                                     | :heavy_check_mark:                                                           | Coupon code                                                                  |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.GetCouponByCodeResponse](../../models/dtos/getcouponbycoderesponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryCoupon

Use when listing or searching coupons (e.g. promo management). Returns a paginated list; supports filtering and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryCoupon" method="post" path="/coupons/search" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Coupons.QueryCoupon(ctx, types.CouponFilter{}, dtos.QueryCouponSecurity{
        Option1: &dtos.QueryCouponSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCouponsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [types.CouponFilter](../../models/types/couponfilter.md)             | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `security`                                                           | [dtos.QueryCouponSecurity](../../models/dtos/querycouponsecurity.md) | :heavy_check_mark:                                                   | The security requirements to use for the request.                    |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.QueryCouponResponse](../../models/dtos/querycouponresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCoupon

Use when you need to load a single coupon (e.g. for display or to validate a code).

### Example Usage

<!-- UsageSnippet language="go" operationID="getCoupon" method="get" path="/coupons/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Coupons.GetCoupon(ctx, dtos.GetCouponSecurity{
        Option1: &dtos.GetCouponSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CouponResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `security`                                                       | [dtos.GetCouponSecurity](../../models/dtos/getcouponsecurity.md) | :heavy_check_mark:                                               | The security requirements to use for the request.                |
| `id`                                                             | `string`                                                         | :heavy_check_mark:                                               | Coupon ID                                                        |
| `opts`                                                           | [][dtos.Option](../../models/dtos/option.md)                     | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*dtos.GetCouponResponse](../../models/dtos/getcouponresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateCoupon

Use when changing coupon config (e.g. value, validity, or usage limits).

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCoupon" method="put" path="/coupons/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Coupons.UpdateCoupon(ctx, dtos.UpdateCouponSecurity{
        Option1: &dtos.UpdateCouponSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", types.UpdateCouponRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.CouponResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `security`                                                             | [dtos.UpdateCouponSecurity](../../models/dtos/updatecouponsecurity.md) | :heavy_check_mark:                                                     | The security requirements to use for the request.                      |
| `id`                                                                   | `string`                                                               | :heavy_check_mark:                                                     | Coupon ID                                                              |
| `body`                                                                 | [types.UpdateCouponRequest](../../models/types/updatecouponrequest.md) | :heavy_check_mark:                                                     | Coupon update request                                                  |
| `opts`                                                                 | [][dtos.Option](../../models/dtos/option.md)                           | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*dtos.UpdateCouponResponse](../../models/dtos/updatecouponresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 401, 403, 404   | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeleteCoupon

Use when retiring a coupon (e.g. campaign ended). Returns 200 with success message.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCoupon" method="delete" path="/coupons/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Coupons.DeleteCoupon(ctx, dtos.DeleteCouponSecurity{
        Option1: &dtos.DeleteCouponSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `security`                                                             | [dtos.DeleteCouponSecurity](../../models/dtos/deletecouponsecurity.md) | :heavy_check_mark:                                                     | The security requirements to use for the request.                      |
| `id`                                                                   | `string`                                                               | :heavy_check_mark:                                                     | Coupon ID                                                              |
| `opts`                                                                 | [][dtos.Option](../../models/dtos/option.md)                           | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*dtos.DeleteCouponResponse](../../models/dtos/deletecouponresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 401, 403, 404   | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |