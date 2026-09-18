# Plans

## Overview

### Available Operations

* [CreatePlan](#createplan) - Create plan
* [QueryPlan](#queryplan) - Query plans
* [GetPlan](#getplan) - Get plan
* [UpdatePlan](#updateplan) - Update plan
* [DeletePlan](#deleteplan) - Delete plan
* [ClonePlan](#cloneplan) - Clone a plan
* [SyncPlanPrices](#syncplanprices) - Synchronize plan prices

## CreatePlan

Use when defining a new pricing plan (e.g. Free, Pro, Enterprise). Attach prices and entitlements; customers subscribe to plans.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPlan" method="post" path="/plans" -->
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

    res, err := s.Plans.CreatePlan(ctx, types.CreatePlanRequest{
        Name: "<value>",
    }, dtos.CreatePlanSecurity{
        Option1: &dtos.CreatePlanSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PlanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [types.CreatePlanRequest](../../models/types/createplanrequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `security`                                                         | [dtos.CreatePlanSecurity](../../models/dtos/createplansecurity.md) | :heavy_check_mark:                                                 | The security requirements to use for the request.                  |
| `opts`                                                             | [][dtos.Option](../../models/dtos/option.md)                       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*dtos.CreatePlanResponse](../../models/dtos/createplanresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryPlan

Use when listing or searching plans (e.g. plan picker or admin catalog). Returns a paginated list; supports filtering and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryPlan" method="post" path="/plans/search" -->
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

    res, err := s.Plans.QueryPlan(ctx, types.PlanFilter{}, dtos.QueryPlanSecurity{
        Option1: &dtos.QueryPlanSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPlansResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [types.PlanFilter](../../models/types/planfilter.md)             | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `security`                                                       | [dtos.QueryPlanSecurity](../../models/dtos/queryplansecurity.md) | :heavy_check_mark:                                               | The security requirements to use for the request.                |
| `opts`                                                           | [][dtos.Option](../../models/dtos/option.md)                     | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*dtos.QueryPlanResponse](../../models/dtos/queryplanresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetPlan

Use when you need to load a single plan (e.g. for display or to create a subscription).

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlan" method="get" path="/plans/{id}" -->
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

    res, err := s.Plans.GetPlan(ctx, dtos.GetPlanSecurity{
        Option1: &dtos.GetPlanSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.PlanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `security`                                                   | [dtos.GetPlanSecurity](../../models/dtos/getplansecurity.md) | :heavy_check_mark:                                           | The security requirements to use for the request.            |
| `id`                                                         | `string`                                                     | :heavy_check_mark:                                           | Plan ID                                                      |
| `opts`                                                       | [][dtos.Option](../../models/dtos/option.md)                 | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*dtos.GetPlanResponse](../../models/dtos/getplanresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdatePlan

Use when changing plan details (e.g. name, interval, or metadata). Partial update supported.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePlan" method="put" path="/plans/{id}" -->
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

    res, err := s.Plans.UpdatePlan(ctx, dtos.UpdatePlanSecurity{
        Option1: &dtos.UpdatePlanSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", types.UpdatePlanRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PlanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `security`                                                         | [dtos.UpdatePlanSecurity](../../models/dtos/updateplansecurity.md) | :heavy_check_mark:                                                 | The security requirements to use for the request.                  |
| `id`                                                               | `string`                                                           | :heavy_check_mark:                                                 | Plan ID                                                            |
| `body`                                                             | [types.UpdatePlanRequest](../../models/types/updateplanrequest.md) | :heavy_check_mark:                                                 | Plan update                                                        |
| `opts`                                                             | [][dtos.Option](../../models/dtos/option.md)                       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*dtos.UpdatePlanResponse](../../models/dtos/updateplanresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeletePlan

Use when retiring a plan (e.g. end-of-life). Existing subscriptions may be affected. Returns 200 with success message.

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePlan" method="delete" path="/plans/{id}" -->
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

    res, err := s.Plans.DeletePlan(ctx, dtos.DeletePlanSecurity{
        Option1: &dtos.DeletePlanSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `security`                                                         | [dtos.DeletePlanSecurity](../../models/dtos/deleteplansecurity.md) | :heavy_check_mark:                                                 | The security requirements to use for the request.                  |
| `id`                                                               | `string`                                                           | :heavy_check_mark:                                                 | Plan ID                                                            |
| `opts`                                                             | [][dtos.Option](../../models/dtos/option.md)                       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*dtos.DeletePlanResponse](../../models/dtos/deleteplanresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ClonePlan

Clone an existing plan, copying its active prices, published entitlements, and published credit grants

### Example Usage

<!-- UsageSnippet language="go" operationID="clonePlan" method="post" path="/plans/{id}/clone" -->
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

    res, err := s.Plans.ClonePlan(ctx, dtos.ClonePlanSecurity{
        Option1: &dtos.ClonePlanSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", types.ClonePlanRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PlanResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `security`                                                       | [dtos.ClonePlanSecurity](../../models/dtos/cloneplansecurity.md) | :heavy_check_mark:                                               | The security requirements to use for the request.                |
| `id`                                                             | `string`                                                         | :heavy_check_mark:                                               | Source Plan ID                                                   |
| `body`                                                           | [types.ClonePlanRequest](../../models/types/cloneplanrequest.md) | :heavy_check_mark:                                               | Clone configuration                                              |
| `opts`                                                           | [][dtos.Option](../../models/dtos/option.md)                     | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*dtos.ClonePlanResponse](../../models/dtos/cloneplanresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404, 409        | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## SyncPlanPrices

Use when you have changed plan prices and need to push them to all active subscriptions (e.g. global price update). Returns workflow ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="syncPlanPrices" method="post" path="/plans/{id}/sync/subscriptions" -->
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

    res, err := s.Plans.SyncPlanPrices(ctx, dtos.SyncPlanPricesSecurity{
        Option1: &dtos.SyncPlanPricesSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ModelsTemporalWorkflowResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `security`                                                                 | [dtos.SyncPlanPricesSecurity](../../models/dtos/syncplanpricessecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |
| `id`                                                                       | `string`                                                                   | :heavy_check_mark:                                                         | Plan ID                                                                    |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.SyncPlanPricesResponse](../../models/dtos/syncplanpricesresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404, 422        | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |