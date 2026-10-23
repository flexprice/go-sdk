# Tenants

## Overview

### Available Operations

* [UpdateTenant](#updatetenant) - Update a tenant
* [GetTenantByID](#gettenantbyid) - Get tenant by ID

## UpdateTenant

Use when changing tenant details (e.g. name or billing info). Request body contains the fields to update.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTenant" method="put" path="/tenants/update" -->
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

    res, err := s.Tenants.UpdateTenant(ctx, types.UpdateTenantRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TenantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [types.UpdateTenantRequest](../../models/types/updatetenantrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][dtos.Option](../../models/dtos/option.md)                           | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*dtos.UpdateTenantResponse](../../models/dtos/updatetenantresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetTenantByID

Get tenant by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="getTenantById" method="get" path="/tenants/{id}" -->
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

    res, err := s.Tenants.GetTenantByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.TenantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Tenant ID                                             |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetTenantByIDResponse](../../models/dtos/gettenantbyidresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |