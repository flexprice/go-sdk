# Analytics

## Overview

### Available Operations

* [QueryAnalytics](#queryanalytics) - Run an ad-hoc analytics query
* [ListAnalyticsReports](#listanalyticsreports) - List financial reports
* [RunAnalyticsReport](#runanalyticsreport) - Run a financial report
* [GetRevenueAnalytics](#getrevenueanalytics) - Query revenue analytics
* [ListAnalyticsViews](#listanalyticsviews) - List analytics views
* [CreateAnalyticsView](#createanalyticsview) - Create an analytics view
* [GetAnalyticsView](#getanalyticsview) - Get an analytics view
* [UpdateAnalyticsView](#updateanalyticsview) - Update an analytics view
* [DeleteAnalyticsView](#deleteanalyticsview) - Delete an analytics view
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
                types.MetricEventsVolume,
            },
            Shape: types.ShapeScalar,
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

## ListAnalyticsReports

Returns the catalog of reports with their allowed shapes, dimensions and grains.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAnalyticsReports" method="get" path="/analytics/reports" -->
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

    res, err := s.Analytics.ListAnalyticsReports(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ReportCatalogResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.ListAnalyticsReportsResponse](../../models/dtos/listanalyticsreportsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 403                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## RunAnalyticsReport

Runs the report identified by key. range is an absolute YYYY-MM-DD..YYYY-MM-DD (end exclusive) or last_N_days. compare is previous_period or previous_year.

### Example Usage

<!-- UsageSnippet language="go" operationID="runAnalyticsReport" method="post" path="/analytics/reports/{key}/query" -->
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

    res, err := s.Analytics.RunAnalyticsReport(ctx, "<key>", types.ReportQueryRequest{
        Range: "2026-09-01..2026-10-01",
        Shape: types.ReportQueryRequestShapeBreakdown,
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `key`                                                                | `string`                                                             | :heavy_check_mark:                                                   | Report key                                                           |
| `body`                                                               | [types.ReportQueryRequest](../../models/types/reportqueryrequest.md) | :heavy_check_mark:                                                   | Report query request                                                 |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.RunAnalyticsReportResponse](../../models/dtos/runanalyticsreportresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 403, 404        | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetRevenueAnalytics

Aggregates revenue_facts by the requested dimensions at day/period/total granularity. allocation_policy places whole-period charges on their booked day (billed) or spreads them across the period (amortized); include_adjustments breaks out true-up/overage/revert amounts as labeled rows. Requires the tenant's revenue analytics setting.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRevenueAnalytics" method="post" path="/analytics/revenue" -->
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

    res, err := s.Analytics.GetRevenueAnalytics(ctx, types.RevenueAnalyticsRequest{
        StartTime: types.MustTimeFromString("2026-09-19T08:36:56.747Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RevenueAnalyticsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [types.RevenueAnalyticsRequest](../../models/types/revenueanalyticsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][dtos.Option](../../models/dtos/option.md)                                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*dtos.GetRevenueAnalyticsResponse](../../models/dtos/getrevenueanalyticsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 403             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ListAnalyticsViews

Lists saved views, optionally filtered by family.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAnalyticsViews" method="get" path="/analytics/views" -->
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

    res, err := s.Analytics.ListAnalyticsViews(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListViewsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `family`                                              | `*string`                                             | :heavy_minus_sign:                                    | Filter by view family                                 |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.ListAnalyticsViewsResponse](../../models/dtos/listanalyticsviewsresponse.md), error**

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
                types.MetricNetBilling,
            },
            Shape: types.ShapeScalar,
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

## GetAnalyticsView

Returns a saved view by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAnalyticsView" method="get" path="/analytics/views/{id}" -->
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

    res, err := s.Analytics.GetAnalyticsView(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ViewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | View ID                                               |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetAnalyticsViewResponse](../../models/dtos/getanalyticsviewresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateAnalyticsView

Replaces the name and definition of a saved view.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAnalyticsView" method="put" path="/analytics/views/{id}" -->
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

    res, err := s.Analytics.UpdateAnalyticsView(ctx, "<id>", types.CreateViewRequest{
        Definition: types.AnalyticsViewDefinition{
            Metrics: []types.Metric{},
            Shape: types.ShapeTimeseries,
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
| `id`                                                               | `string`                                                           | :heavy_check_mark:                                                 | View ID                                                            |
| `body`                                                             | [types.CreateViewRequest](../../models/types/createviewrequest.md) | :heavy_check_mark:                                                 | View request                                                       |
| `opts`                                                             | [][dtos.Option](../../models/dtos/option.md)                       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*dtos.UpdateAnalyticsViewResponse](../../models/dtos/updateanalyticsviewresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeleteAnalyticsView

Deletes a saved view by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAnalyticsView" method="delete" path="/analytics/views/{id}" -->
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

    res, err := s.Analytics.DeleteAnalyticsView(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | View ID                                               |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeleteAnalyticsViewResponse](../../models/dtos/deleteanalyticsviewresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
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