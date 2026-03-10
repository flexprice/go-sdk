# CreditGrants

## Overview

### Available Operations

* [CreateCreditGrant](#createcreditgrant) - Create credit grant
* [GetCreditGrant](#getcreditgrant) - Get credit grant
* [UpdateCreditGrant](#updatecreditgrant) - Update credit grant
* [DeleteCreditGrant](#deletecreditgrant) - Delete credit grant
* [GetPlanCreditGrants](#getplancreditgrants) - Get plan credit grants

## CreateCreditGrant

Use when giving a customer or plan credits (e.g. prepaid balance or promotional credits). Scope can be plan or subscription; supports start/end dates.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCreditGrant" method="post" path="/creditgrants" -->
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

    res, err := s.CreditGrants.CreateCreditGrant(ctx, types.DtoCreateCreditGrantRequest{
        Cadence: types.CreditGrantCadenceOnetime,
        Credits: "<value>",
        Name: "<value>",
        Scope: types.CreditGrantScopePlan,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreditGrantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [types.DtoCreateCreditGrantRequest](../../models/types/dtocreatecreditgrantrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*dtos.CreateCreditGrantResponse](../../models/dtos/createcreditgrantresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetCreditGrant

Use when you need to load a single credit grant (e.g. for display or to check balance).

### Example Usage

<!-- UsageSnippet language="go" operationID="getCreditGrant" method="get" path="/creditgrants/{id}" -->
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

    res, err := s.CreditGrants.GetCreditGrant(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreditGrantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Credit Grant ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetCreditGrantResponse](../../models/dtos/getcreditgrantresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateCreditGrant

Use when changing a credit grant (e.g. amount or end date). Request body contains the fields to update.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCreditGrant" method="put" path="/creditgrants/{id}" -->
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

    res, err := s.CreditGrants.UpdateCreditGrant(ctx, "<id>", types.DtoUpdateCreditGrantRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreditGrantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | `string`                                                                               | :heavy_check_mark:                                                                     | Credit Grant ID                                                                        |
| `body`                                                                                 | [types.DtoUpdateCreditGrantRequest](../../models/types/dtoupdatecreditgrantrequest.md) | :heavy_check_mark:                                                                     | Credit Grant configuration                                                             |
| `opts`                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*dtos.UpdateCreditGrantResponse](../../models/dtos/updatecreditgrantresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## DeleteCreditGrant

Use when removing or ending a credit grant (e.g. revoke promo or close prepaid). Plan-scoped grants are archived; subscription-scoped supports optional effective_date in body.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCreditGrant" method="delete" path="/creditgrants/{id}" -->
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

    res, err := s.CreditGrants.DeleteCreditGrant(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `id`                                                                                    | `string`                                                                                | :heavy_check_mark:                                                                      | Credit Grant ID                                                                         |
| `body`                                                                                  | [*types.DtoDeleteCreditGrantRequest](../../models/types/dtodeletecreditgrantrequest.md) | :heavy_minus_sign:                                                                      | Optional: effective_date for subscription-scoped grants                                 |
| `opts`                                                                                  | [][dtos.Option](../../models/dtos/option.md)                                            | :heavy_minus_sign:                                                                      | The options for this request.                                                           |

### Response

**[*dtos.DeleteCreditGrantResponse](../../models/dtos/deletecreditgrantresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetPlanCreditGrants

Use when listing credits attached to a plan (e.g. included prepaid or promo credits).

### Example Usage

<!-- UsageSnippet language="go" operationID="getPlanCreditGrants" method="get" path="/plans/{id}/creditgrants" -->
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

    res, err := s.CreditGrants.GetPlanCreditGrants(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListCreditGrantsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Plan ID                                               |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetPlanCreditGrantsResponse](../../models/dtos/getplancreditgrantsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |