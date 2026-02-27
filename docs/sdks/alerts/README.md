# Alerts

## Overview

### Available Operations

* [QueryAlertLog](#queryalertlog) - Query alert logs

## QueryAlertLog

Use when viewing or searching alert history (e.g. support triage or customer-facing alert log). Returns a paginated list; supports filtering by type, customer, subscription.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryAlertLog" method="post" path="/alerts/search" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Alerts.QueryAlertLog(ctx, types.AlertLogFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListAlertLogsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [types.AlertLogFilter](../../types/alertlogfilter.md) | :heavy_check_mark:                                    | The request object to use for the request.            |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.QueryAlertLogResponse](../../dtos/queryalertlogresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |