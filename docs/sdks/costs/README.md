# Costs

## Overview

### Available Operations

* [CreateCostsheet](#createcostsheet) - Create costsheet
* [GetActiveCostsheet](#getactivecostsheet) - Get active costsheet
* [GetDetailedCostAnalytics](#getdetailedcostanalytics) - Get combined revenue and cost analytics
* [GetDetailedCostAnalyticsV2](#getdetailedcostanalyticsv2) - Get combined revenue and cost analytics (V2)
* [QueryCostsheet](#querycostsheet) - Query costsheets
* [GetCostsheet](#getcostsheet) - Get costsheet
* [UpdateCostsheet](#updatecostsheet) - Update costsheet
* [DeleteCostsheet](#deletecostsheet) - Delete costsheet

## CreateCostsheet

Use when setting up a new pricing configuration (e.g. a new product or region). Costsheets group prices and define the default for the environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCostsheet" method="post" path="/costs" -->
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

    res, err := s.Costs.CreateCostsheet(ctx, types.DtoCreateCostsheetRequest{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreateCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [types.DtoCreateCostsheetRequest](../../models/types/dtocreatecostsheetrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.CreateCostsheetResponse](../../models/dtos/createcostsheetresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 409                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetActiveCostsheet

Use when you need the tenant's default pricing configuration (e.g. for checkout or plan display). Returns the active costsheet for the environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="getActiveCostsheet" method="get" path="/costs/active" -->
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

    res, err := s.Costs.GetActiveCostsheet(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCostsheetResponse != nil {
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

**[*dtos.GetActiveCostsheetResponse](../../models/dtos/getactivecostsheetresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 404                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetDetailedCostAnalytics

Use when building dashboards or reports that need revenue vs cost, ROI, and margin over a time period (e.g. finance views or executive summaries).

### Example Usage

<!-- UsageSnippet language="go" operationID="getDetailedCostAnalytics" method="post" path="/costs/analytics" -->
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

    res, err := s.Costs.GetDetailedCostAnalytics(ctx, types.DtoGetCostAnalyticsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetDetailedCostAnalyticsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [types.DtoGetCostAnalyticsRequest](../../models/types/dtogetcostanalyticsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.GetDetailedCostAnalyticsResponse](../../models/dtos/getdetailedcostanalyticsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetDetailedCostAnalyticsV2

Use when you need the same revenue/cost/ROI analytics but computed from the costsheet usage-tracking pipeline (e.g. for consistency with usage-based cost data).

### Example Usage

<!-- UsageSnippet language="go" operationID="getDetailedCostAnalyticsV2" method="post" path="/costs/analytics-v2" -->
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

    res, err := s.Costs.GetDetailedCostAnalyticsV2(ctx, types.DtoGetCostAnalyticsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetDetailedCostAnalyticsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [types.DtoGetCostAnalyticsRequest](../../models/types/dtogetcostanalyticsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.GetDetailedCostAnalyticsV2Response](../../models/dtos/getdetailedcostanalyticsv2response.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## QueryCostsheet

Use when listing or searching costsheets (e.g. admin catalog). Returns a paginated list; supports filtering and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryCostsheet" method="post" path="/costs/search" -->
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

    res, err := s.Costs.QueryCostsheet(ctx, types.CostsheetFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [types.CostsheetFilter](../../models/types/costsheetfilter.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][dtos.Option](../../models/dtos/option.md)                   | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*dtos.QueryCostsheetResponse](../../models/dtos/querycostsheetresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetCostsheet

Use when you need to load a single costsheet (e.g. for editing or display). Supports optional expand for related prices.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCostsheet" method="get" path="/costs/{id}" -->
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

    res, err := s.Costs.GetCostsheet(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `id`                                                      | *string*                                                  | :heavy_check_mark:                                        | Costsheet ID                                              |
| `expand`                                                  | **string*                                                 | :heavy_minus_sign:                                        | Comma-separated list of fields to expand (e.g., 'prices') |
| `opts`                                                    | [][dtos.Option](../../models/dtos/option.md)              | :heavy_minus_sign:                                        | The options for this request.                             |

### Response

**[*dtos.GetCostsheetResponse](../../models/dtos/getcostsheetresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateCostsheet

Use when changing costsheet name or metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCostsheet" method="put" path="/costs/{id}" -->
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

    res, err := s.Costs.UpdateCostsheet(ctx, "<id>", types.DtoUpdateCostsheetRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoUpdateCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `id`                                                                               | *string*                                                                           | :heavy_check_mark:                                                                 | Costsheet ID                                                                       |
| `body`                                                                             | [types.DtoUpdateCostsheetRequest](../../models/types/dtoupdatecostsheetrequest.md) | :heavy_check_mark:                                                                 | Costsheet configuration                                                            |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.UpdateCostsheetResponse](../../models/dtos/updatecostsheetresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404, 409              | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## DeleteCostsheet

Use when retiring a costsheet (e.g. end-of-life product). Soft-deletes; status set to deleted.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCostsheet" method="delete" path="/costs/{id}" -->
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

    res, err := s.Costs.DeleteCostsheet(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoDeleteCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Costsheet ID                                          |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeleteCostsheetResponse](../../models/dtos/deletecostsheetresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |