# PerformanceObligations

## Overview

### Available Operations

* [QueryPerformanceObligation](#queryperformanceobligation) - Query performance obligations
* [GetPerformanceObligation](#getperformanceobligation) - Get performance obligation

## QueryPerformanceObligation

Use when searching performance obligations by contract, customer, status or recognition method. Returns a paginated list.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryPerformanceObligation" method="post" path="/performance_obligations/search" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.PerformanceObligations.QueryPerformanceObligation(ctx, types.PobFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPerformanceObligationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [types.PobFilter](../../models/types/pobfilter.md)    | :heavy_check_mark:                                    | The request object to use for the request.            |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.QueryPerformanceObligationResponse](../../models/dtos/queryperformanceobligationresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetPerformanceObligation

Use when you need to load a single performance obligation with the subscription line items it covers. Performance obligations are derived from subscriptions and are read-only.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPerformanceObligation" method="get" path="/performance_obligations/{id}" -->
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

    res, err := s.PerformanceObligations.GetPerformanceObligation(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.PerformanceObligationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Performance obligation ID                             |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetPerformanceObligationResponse](../../models/dtos/getperformanceobligationresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |