# Tenants

## Overview

### Available Operations

* [GetTenantBillingUsage](#gettenantbillingusage) - Get billing usage for the current tenant
* [UpdateTenant](#updatetenant) - Update a tenant
* [GetTenantByID](#gettenantbyid) - Get tenant by ID

## GetTenantBillingUsage

Use when showing the current tenant's billing usage (e.g. admin billing page or usage caps). Returns subscription and usage for the tenant.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTenantBillingUsage" method="get" path="/tenant/billing" -->
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

    res, err := s.Tenants.GetTenantBillingUsage(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTenantBillingUsage != nil {
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

**[*dtos.GetTenantBillingUsageResponse](../../models/dtos/gettenantbillingusageresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateTenant

Use when changing tenant details (e.g. name or billing info). Request body contains the fields to update.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTenant" method="put" path="/tenants/update" -->
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

    res, err := s.Tenants.UpdateTenant(ctx, types.DtoUpdateTenantRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTenantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [types.DtoUpdateTenantRequest](../../models/types/dtoupdatetenantrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.UpdateTenantResponse](../../models/dtos/updatetenantresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetTenantByID

Get tenant by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="getTenantById" method="get" path="/tenants/{id}" -->
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

    res, err := s.Tenants.GetTenantByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTenantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Tenant ID                                             |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetTenantByIDResponse](../../models/dtos/gettenantbyidresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 404                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |