# AccountingPeriods

## Overview

### Available Operations

* [QueryAccountingPeriod](#queryaccountingperiod) - Query accounting periods

## QueryAccountingPeriod

Use when listing the monthly accounting periods revenue recognition posts to, by period status or a period start range. Returns a paginated list.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryAccountingPeriod" method="post" path="/accounting_periods/search" -->
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

    res, err := s.AccountingPeriods.QueryAccountingPeriod(ctx, types.AccountingperiodFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAccountingPeriodsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [types.AccountingperiodFilter](../../models/types/accountingperiodfilter.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.QueryAccountingPeriodResponse](../../models/dtos/queryaccountingperiodresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |