# Analytics

## Overview

### Available Operations

* [QueryAnalytics](#queryanalytics) - Run an ad-hoc analytics query
* [CreateAnalyticsView](#createanalyticsview) - Create an analytics view
* [QueryAnalyticsView](#queryanalyticsview) - Query an analytics view

## QueryAnalytics

Resolves the given view definition against the supplied variables and executes it.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryAnalytics" method="post" path="/analytics/query" -->
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

    res, err := s.Analytics.QueryAnalytics(ctx, types.AnalyticsQueryRequest{
        Definition: types.AnalyticsViewDefinition{
            Metrics: []types.Metric{
                types.MetricUsageQuantity,
            },
            Shape: types.ShapeBreakdown,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsQueryResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [types.AnalyticsQueryRequest](../../models/types/analyticsqueryrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.QueryAnalyticsResponse](../../models/dtos/queryanalyticsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CreateAnalyticsView

Persists a named view definition that can later be queried by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAnalyticsView" method="post" path="/analytics/views" -->
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

    res, err := s.Analytics.CreateAnalyticsView(ctx, types.CreateViewRequest{
        Definition: types.AnalyticsViewDefinition{
            Metrics: []types.Metric{
                types.MetricEventCount,
            },
            Shape: types.ShapeBreakdown,
        },
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ViewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [types.CreateViewRequest](../../models/types/createviewrequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][dtos.Option](../../models/dtos/option.md)                       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*dtos.CreateAnalyticsViewResponse](../../models/dtos/createanalyticsviewresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryAnalyticsView

Resolves the view's definition against the supplied variables and executes it.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryAnalyticsView" method="post" path="/analytics/views/{id}/query" -->
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

    res, err := s.Analytics.QueryAnalyticsView(ctx, "<id>", types.ViewQueryRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyticsQueryResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `id`                                                             | `string`                                                         | :heavy_check_mark:                                               | View ID                                                          |
| `body`                                                           | [types.ViewQueryRequest](../../models/types/viewqueryrequest.md) | :heavy_check_mark:                                               | View query request                                               |
| `opts`                                                           | [][dtos.Option](../../models/dtos/option.md)                     | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*dtos.QueryAnalyticsViewResponse](../../models/dtos/queryanalyticsviewresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |