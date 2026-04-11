# MeterUsage

## Overview

### Available Operations

* [GetMeterUsageAnalytics](#getmeterusageanalytics) - Get meter usage analytics
* [QueryMeterUsage](#querymeterusage) - Query meter usage

## GetMeterUsageAnalytics

Query aggregated usage from meter_usage table for multiple meters

### Example Usage

<!-- UsageSnippet language="go" operationID="getMeterUsageAnalytics" method="post" path="/meter-usage/analytics" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.MeterUsage.GetMeterUsageAnalytics(ctx, types.MeterUsageAnalyticsRequest{
        AggregationType: types.AggregationTypeCount,
        EndTime: types.MustTimeFromString("2024-02-01T00:00:00Z"),
        ExternalCustomerID: "cust_123",
        MeterIds: []string{
            "mtr_abc",
            "mtr_def",
        },
        StartTime: types.MustTimeFromString("2024-01-01T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MeterUsageAnalyticsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [types.MeterUsageAnalyticsRequest](../../models/types/meterusageanalyticsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.GetMeterUsageAnalyticsResponse](../../models/dtos/getmeterusageanalyticsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryMeterUsage

Query aggregated usage from meter_usage table for a single meter with optional time-window bucketing

### Example Usage

<!-- UsageSnippet language="go" operationID="queryMeterUsage" method="post" path="/meter-usage/query" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.MeterUsage.QueryMeterUsage(ctx, types.MeterUsageQueryRequest{
        AggregationType: types.AggregationTypeWeightedSum,
        BillingAnchor: types.MustNewTimeFromString("2024-01-15T00:00:00Z"),
        EndTime: types.MustTimeFromString("2024-02-01T00:00:00Z"),
        ExternalCustomerID: "cust_123",
        MeterID: "mtr_abc",
        StartTime: types.MustTimeFromString("2024-01-01T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MeterUsageQueryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [types.MeterUsageQueryRequest](../../models/types/meterusagequeryrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.QueryMeterUsageResponse](../../models/dtos/querymeterusageresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |