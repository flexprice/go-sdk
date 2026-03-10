# Coupons

## Overview

### Available Operations

* [CreateCoupon](#createcoupon) - Create coupon
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
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Coupons.CreateCoupon(ctx, types.DtoCreateCouponRequest{
        Cadence: types.CouponCadenceRepeated,
        Name: "<value>",
        Type: types.CouponTypePercentage,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCouponResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [types.DtoCreateCouponRequest](../../models/types/dtocreatecouponrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.CreateCouponResponse](../../models/dtos/createcouponresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 401, 403, 404         | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## QueryCoupon

Use when listing or searching coupons (e.g. promo management). Returns a paginated list; supports filtering and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryCoupon" method="post" path="/coupons/search" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Coupons.QueryCoupon(ctx, types.CouponFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListCouponsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [types.CouponFilter](../../models/types/couponfilter.md) | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][dtos.Option](../../models/dtos/option.md)             | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*dtos.QueryCouponResponse](../../models/dtos/querycouponresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetCoupon

Use when you need to load a single coupon (e.g. for display or to validate a code).

### Example Usage

<!-- UsageSnippet language="go" operationID="getCoupon" method="get" path="/coupons/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Coupons.GetCoupon(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCouponResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Coupon ID                                             |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetCouponResponse](../../models/dtos/getcouponresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateCoupon

Use when changing coupon config (e.g. value, validity, or usage limits).

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCoupon" method="put" path="/coupons/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Coupons.UpdateCoupon(ctx, "<id>", types.DtoUpdateCouponRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCouponResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `id`                                                                         | `string`                                                                     | :heavy_check_mark:                                                           | Coupon ID                                                                    |
| `body`                                                                       | [types.DtoUpdateCouponRequest](../../models/types/dtoupdatecouponrequest.md) | :heavy_check_mark:                                                           | Coupon update request                                                        |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.UpdateCouponResponse](../../models/dtos/updatecouponresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 401, 403, 404         | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## DeleteCoupon

Use when retiring a coupon (e.g. campaign ended). Returns 200 with success message.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCoupon" method="delete" path="/coupons/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Coupons.DeleteCoupon(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Coupon ID                                             |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeleteCouponResponse](../../models/dtos/deletecouponresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 401, 403, 404         | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |