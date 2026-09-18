# Entitlements

## Overview

### Available Operations

* [GetAddonEntitlements](#getaddonentitlements) - Get addon entitlements
* [CreateEntitlement](#createentitlement) - Create entitlement
* [CreateEntitlementsBulk](#createentitlementsbulk) - Create entitlements in bulk
* [QueryEntitlement](#queryentitlement) - Query entitlements
* [GetEntitlement](#getentitlement) - Get entitlement
* [UpdateEntitlement](#updateentitlement) - Update entitlement
* [DeleteEntitlement](#deleteentitlement) - Delete entitlement
* [GetPlanEntitlements](#getplanentitlements) - Get plan entitlements

## GetAddonEntitlements

Use when checking what features or limits an addon grants (e.g. for display or entitlement logic).

### Example Usage

<!-- UsageSnippet language="go" operationID="getAddonEntitlements" method="get" path="/addons/{id}/entitlements" -->
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

    res, err := s.Entitlements.GetAddonEntitlements(ctx, dtos.GetAddonEntitlementsSecurity{
        Option1: &dtos.GetAddonEntitlementsSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `security`                                                                             | [dtos.GetAddonEntitlementsSecurity](../../models/dtos/getaddonentitlementssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |
| `id`                                                                                   | `string`                                                                               | :heavy_check_mark:                                                                     | Addon ID                                                                               |
| `opts`                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*dtos.GetAddonEntitlementsResponse](../../models/dtos/getaddonentitlementsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CreateEntitlement

Use when attaching a feature (and its limit) to a plan or addon (e.g. "10 seats" or "1000 API calls"). Defines what the plan/addon includes.

### Example Usage

<!-- UsageSnippet language="go" operationID="createEntitlement" method="post" path="/entitlements" -->
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

    res, err := s.Entitlements.CreateEntitlement(ctx, types.CreateEntitlementRequest{
        FeatureID: "<id>",
        FeatureType: types.FeatureTypeMetered,
    }, dtos.CreateEntitlementSecurity{
        Option1: &dtos.CreateEntitlementSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [types.CreateEntitlementRequest](../../models/types/createentitlementrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [dtos.CreateEntitlementSecurity](../../models/dtos/createentitlementsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |
| `opts`                                                                           | [][dtos.Option](../../models/dtos/option.md)                                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*dtos.CreateEntitlementResponse](../../models/dtos/createentitlementresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CreateEntitlementsBulk

Use when attaching many features to a plan or addon at once (e.g. initial plan setup or import). Bulk version of create entitlement.

### Example Usage

<!-- UsageSnippet language="go" operationID="createEntitlementsBulk" method="post" path="/entitlements/bulk" -->
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

    res, err := s.Entitlements.CreateEntitlementsBulk(ctx, types.CreateBulkEntitlementRequest{
        Items: []types.CreateEntitlementRequest{
            types.CreateEntitlementRequest{
                FeatureID: "<id>",
                FeatureType: types.FeatureTypeConfig,
            },
        },
    }, dtos.CreateEntitlementsBulkSecurity{
        Option1: &dtos.CreateEntitlementsBulkSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateBulkEntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [types.CreateBulkEntitlementRequest](../../models/types/createbulkentitlementrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [dtos.CreateEntitlementsBulkSecurity](../../models/dtos/createentitlementsbulksecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |
| `opts`                                                                                     | [][dtos.Option](../../models/dtos/option.md)                                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*dtos.CreateEntitlementsBulkResponse](../../models/dtos/createentitlementsbulkresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryEntitlement

Use when listing or searching entitlements (e.g. plan editor or audit). Returns a paginated list; supports filtering by plan, addon, feature.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryEntitlement" method="post" path="/entitlements/search" -->
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

    res, err := s.Entitlements.QueryEntitlement(ctx, types.EntitlementFilter{}, dtos.QueryEntitlementSecurity{
        Option1: &dtos.QueryEntitlementSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [types.EntitlementFilter](../../models/types/entitlementfilter.md)             | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [dtos.QueryEntitlementSecurity](../../models/dtos/queryentitlementsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |
| `opts`                                                                         | [][dtos.Option](../../models/dtos/option.md)                                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*dtos.QueryEntitlementResponse](../../models/dtos/queryentitlementresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetEntitlement

Use when you need to load a single entitlement (e.g. to display or edit a feature limit).

### Example Usage

<!-- UsageSnippet language="go" operationID="getEntitlement" method="get" path="/entitlements/{id}" -->
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

    res, err := s.Entitlements.GetEntitlement(ctx, dtos.GetEntitlementSecurity{
        Option1: &dtos.GetEntitlementSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `security`                                                                 | [dtos.GetEntitlementSecurity](../../models/dtos/getentitlementsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |
| `id`                                                                       | `string`                                                                   | :heavy_check_mark:                                                         | Entitlement ID                                                             |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.GetEntitlementResponse](../../models/dtos/getentitlementresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateEntitlement

Use when changing an entitlement (e.g. increasing or decreasing a limit). Request body contains the fields to update.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateEntitlement" method="put" path="/entitlements/{id}" -->
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

    res, err := s.Entitlements.UpdateEntitlement(ctx, dtos.UpdateEntitlementSecurity{
        Option1: &dtos.UpdateEntitlementSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", types.UpdateEntitlementRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `security`                                                                       | [dtos.UpdateEntitlementSecurity](../../models/dtos/updateentitlementsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |
| `id`                                                                             | `string`                                                                         | :heavy_check_mark:                                                               | Entitlement ID                                                                   |
| `body`                                                                           | [types.UpdateEntitlementRequest](../../models/types/updateentitlementrequest.md) | :heavy_check_mark:                                                               | Entitlement configuration                                                        |
| `opts`                                                                           | [][dtos.Option](../../models/dtos/option.md)                                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*dtos.UpdateEntitlementResponse](../../models/dtos/updateentitlementresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeleteEntitlement

Use when removing a feature from a plan or addon (e.g. deprecating a capability). Returns 200 with success message.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteEntitlement" method="delete" path="/entitlements/{id}" -->
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

    res, err := s.Entitlements.DeleteEntitlement(ctx, dtos.DeleteEntitlementSecurity{
        Option1: &dtos.DeleteEntitlementSecurityOption1{
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `security`                                                                       | [dtos.DeleteEntitlementSecurity](../../models/dtos/deleteentitlementsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |
| `id`                                                                             | `string`                                                                         | :heavy_check_mark:                                                               | Entitlement ID                                                                   |
| `opts`                                                                           | [][dtos.Option](../../models/dtos/option.md)                                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*dtos.DeleteEntitlementResponse](../../models/dtos/deleteentitlementresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetPlanEntitlements

Use when checking what a plan includes (e.g. feature list or limits for display or gating).

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlanEntitlements" method="get" path="/plans/{id}/entitlements" -->
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

    res, err := s.Entitlements.GetPlanEntitlements(ctx, dtos.GetPlanEntitlementsSecurity{
        Option1: &dtos.GetPlanEntitlementsSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `security`                                                                           | [dtos.GetPlanEntitlementsSecurity](../../models/dtos/getplanentitlementssecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |
| `id`                                                                                 | `string`                                                                             | :heavy_check_mark:                                                                   | Plan ID                                                                              |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.GetPlanEntitlementsResponse](../../models/dtos/getplanentitlementsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |