# CouponAssociations

## Overview

### Available Operations

* [ListCouponAssociations](#listcouponassociations) - List coupon associations
* [GetCouponAssociation](#getcouponassociation) - Get coupon association

## ListCouponAssociations

List coupon associations with optional filters. Coupon associations are created and removed via the subscription modify API.

### Example Usage

<!-- UsageSnippet language="go" operationID="listCouponAssociations" method="get" path="/coupons/associations" -->
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

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.CouponAssociations.ListCouponAssociations(ctx, dtos.ListCouponAssociationsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCouponAssociationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [dtos.ListCouponAssociationsRequest](../../models/dtos/listcouponassociationsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*dtos.ListCouponAssociationsResponse](../../models/dtos/listcouponassociationsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCouponAssociation

Get a single coupon association by ID. Coupon associations are created and removed via the subscription modify API.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCouponAssociation" method="get" path="/coupons/associations/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.CouponAssociations.GetCouponAssociation(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CouponAssociationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Coupon Association ID                                 |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetCouponAssociationResponse](../../models/dtos/getcouponassociationresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |