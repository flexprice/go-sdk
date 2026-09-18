# Costs

## Overview

### Available Operations

* [CreateCostsheet](#createcostsheet) - Create costsheet
* [GetActiveCostsheet](#getactivecostsheet) - Get active costsheet
* [GetDetailedCostAnalytics](#getdetailedcostanalytics) - Get combined revenue and cost analytics
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
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Costs.CreateCostsheet(ctx, types.CreateCostsheetRequest{
        Name: "<value>",
    }, dtos.CreateCostsheetSecurity{
        Option1: &dtos.CreateCostsheetSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [types.CreateCostsheetRequest](../../models/types/createcostsheetrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [dtos.CreateCostsheetSecurity](../../models/dtos/createcostsheetsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.CreateCostsheetResponse](../../models/dtos/createcostsheetresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 409             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetActiveCostsheet

Use when you need the tenant's default pricing configuration (e.g. for checkout or plan display). Returns the active costsheet for the environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="getActiveCostsheet" method="get" path="/costs/active" -->
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

    s := flexprice.New()

    res, err := s.Costs.GetActiveCostsheet(ctx, dtos.GetActiveCostsheetSecurity{
        Option1: &dtos.GetActiveCostsheetSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `security`                                                                         | [dtos.GetActiveCostsheetSecurity](../../models/dtos/getactivecostsheetsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.GetActiveCostsheetResponse](../../models/dtos/getactivecostsheetresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetDetailedCostAnalytics

Use when building dashboards or reports that need revenue vs cost, ROI, and margin over a time period (e.g. finance views or executive summaries).

### Example Usage

<!-- UsageSnippet language="go" operationID="getDetailedCostAnalytics" method="post" path="/costs/analytics" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Costs.GetDetailedCostAnalytics(ctx, types.GetCostAnalyticsRequest{}, dtos.GetDetailedCostAnalyticsSecurity{
        Option1: &dtos.GetDetailedCostAnalyticsSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDetailedCostAnalyticsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [types.GetCostAnalyticsRequest](../../models/types/getcostanalyticsrequest.md)                 | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `security`                                                                                     | [dtos.GetDetailedCostAnalyticsSecurity](../../models/dtos/getdetailedcostanalyticssecurity.md) | :heavy_check_mark:                                                                             | The security requirements to use for the request.                                              |
| `opts`                                                                                         | [][dtos.Option](../../models/dtos/option.md)                                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*dtos.GetDetailedCostAnalyticsResponse](../../models/dtos/getdetailedcostanalyticsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryCostsheet

Use when listing or searching costsheets (e.g. admin catalog). Returns a paginated list; supports filtering and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryCostsheet" method="post" path="/costs/search" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Costs.QueryCostsheet(ctx, types.CostsheetFilter{}, dtos.QueryCostsheetSecurity{
        Option1: &dtos.QueryCostsheetSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [types.CostsheetFilter](../../models/types/costsheetfilter.md)             | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [dtos.QueryCostsheetSecurity](../../models/dtos/querycostsheetsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.QueryCostsheetResponse](../../models/dtos/querycostsheetresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCostsheet

Use when you need to load a single costsheet (e.g. for editing or display). Supports optional expand for related prices.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCostsheet" method="get" path="/costs/{id}" -->
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

    s := flexprice.New()

    res, err := s.Costs.GetCostsheet(ctx, dtos.GetCostsheetSecurity{
        Option1: &dtos.GetCostsheetSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `security`                                                             | [dtos.GetCostsheetSecurity](../../models/dtos/getcostsheetsecurity.md) | :heavy_check_mark:                                                     | The security requirements to use for the request.                      |
| `id`                                                                   | `string`                                                               | :heavy_check_mark:                                                     | Costsheet ID                                                           |
| `expand`                                                               | `*string`                                                              | :heavy_minus_sign:                                                     | Comma-separated list of fields to expand (e.g., 'prices')              |
| `opts`                                                                 | [][dtos.Option](../../models/dtos/option.md)                           | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*dtos.GetCostsheetResponse](../../models/dtos/getcostsheetresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateCostsheet

Use when changing costsheet name or metadata.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCostsheet" method="put" path="/costs/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Costs.UpdateCostsheet(ctx, dtos.UpdateCostsheetSecurity{
        Option1: &dtos.UpdateCostsheetSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", types.UpdateCostsheetRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `security`                                                                   | [dtos.UpdateCostsheetSecurity](../../models/dtos/updatecostsheetsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |
| `id`                                                                         | `string`                                                                     | :heavy_check_mark:                                                           | Costsheet ID                                                                 |
| `body`                                                                       | [types.UpdateCostsheetRequest](../../models/types/updatecostsheetrequest.md) | :heavy_check_mark:                                                           | Costsheet configuration                                                      |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.UpdateCostsheetResponse](../../models/dtos/updatecostsheetresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404, 409        | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeleteCostsheet

Use when retiring a costsheet (e.g. end-of-life product). Soft-deletes; status set to deleted.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCostsheet" method="delete" path="/costs/{id}" -->
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

    s := flexprice.New()

    res, err := s.Costs.DeleteCostsheet(ctx, dtos.DeleteCostsheetSecurity{
        Option1: &dtos.DeleteCostsheetSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteCostsheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `security`                                                                   | [dtos.DeleteCostsheetSecurity](../../models/dtos/deletecostsheetsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |
| `id`                                                                         | `string`                                                                     | :heavy_check_mark:                                                           | Costsheet ID                                                                 |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.DeleteCostsheetResponse](../../models/dtos/deletecostsheetresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |