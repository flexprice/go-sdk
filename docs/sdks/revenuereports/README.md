# RevenueReports

## Overview

### Available Operations

* [GetRevenueRollforward](#getrevenuerollforward) - Get revenue rollforward
* [GetRevenueWaterfall](#getrevenuewaterfall) - Get revenue waterfall

## GetRevenueRollforward

Use when explaining how deferred revenue, unbilled receivable or another ledger account moved in ended months: opening balance, movements by category (invoiced, earned, adjustments, payments, refunds) and closing balance, optionally per customer, contract or performance obligation.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRevenueRollforward" method="post" path="/revenue_reports/rollforward" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/types"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.RevenueReports.GetRevenueRollforward(ctx, types.RollforwardRequest{
        Currency: "Moldovan Leu",
        PeriodStart: types.MustTimeFromString("2025-06-02T22:16:06.570Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RollforwardResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [types.RollforwardRequest](../../models/types/rollforwardrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.GetRevenueRollforwardResponse](../../models/dtos/getrevenuerollforwardresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetRevenueWaterfall

Use when showing revenue per customer, contract or performance obligation and month: actual (posted), pending (earned, not yet posted) and scheduled (invoiced, not yet earned), with optional month-end deferred and unbilled balances. Covers at most 36 months.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRevenueWaterfall" method="post" path="/revenue_reports/waterfall" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/types"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.RevenueReports.GetRevenueWaterfall(ctx, types.WaterfallRequest{
        Currency: "Dalasi",
        PeriodEnd: types.MustTimeFromString("2025-08-28T20:02:24.445Z"),
        PeriodStart: types.MustTimeFromString("2024-03-13T15:39:09.277Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WaterfallResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [types.WaterfallRequest](../../models/types/waterfallrequest.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][dtos.Option](../../models/dtos/option.md)                     | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*dtos.GetRevenueWaterfallResponse](../../models/dtos/getrevenuewaterfallresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |