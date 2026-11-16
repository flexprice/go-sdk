# BillingEntities

## Overview

### Available Operations

* [ListBillingEntities](#listbillingentities) - List billing entities
* [CreateBillingEntity](#createbillingentity) - Create a billing entity
* [EnableBillingEntities](#enablebillingentities) - Enable billing entities
* [GetBillingEntitiesStatus](#getbillingentitiesstatus) - Get billing entities status
* [GetBillingEntity](#getbillingentity) - Get a billing entity
* [UpdateBillingEntity](#updatebillingentity) - Update a billing entity

## ListBillingEntities

List the billing entities of this environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="listBillingEntities" method="get" path="/billing-entities" -->
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

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.BillingEntities.ListBillingEntities(ctx, dtos.ListBillingEntitiesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListBillingEntitiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [dtos.ListBillingEntitiesRequest](../../models/dtos/listbillingentitiesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.ListBillingEntitiesResponse](../../models/dtos/listbillingentitiesresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CreateBillingEntity

Create a legal entity that bills customers in this environment. Requires billing entities to be enabled.

### Example Usage

<!-- UsageSnippet language="go" operationID="createBillingEntity" method="post" path="/billing-entities" -->
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

    res, err := s.BillingEntities.CreateBillingEntity(ctx, types.CreateBillingEntityRequest{
        InvoiceNumberPrefix: "<value>",
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [types.CreateBillingEntityRequest](../../models/types/createbillingentityrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.CreateBillingEntityResponse](../../models/dtos/createbillingentityresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 409             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## EnableBillingEntities

Start enabling billing entities in the environment. The environment gets a default entity and its existing customers are assigned to it. Irreversible.

### Example Usage

<!-- UsageSnippet language="go" operationID="enableBillingEntities" method="post" path="/billing-entities/enable" -->
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

    res, err := s.BillingEntities.EnableBillingEntities(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntityModeResponse != nil {
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

**[*dtos.EnableBillingEntitiesResponse](../../models/dtos/enablebillingentitiesresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 403             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetBillingEntitiesStatus

Whether billing entities are off, enabling or enabled in the environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="getBillingEntitiesStatus" method="get" path="/billing-entities/status" -->
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

    res, err := s.BillingEntities.GetBillingEntitiesStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntityModeResponse != nil {
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

**[*dtos.GetBillingEntitiesStatusResponse](../../models/dtos/getbillingentitiesstatusresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetBillingEntity

Load a single billing entity by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getBillingEntity" method="get" path="/billing-entities/{id}" -->
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

    res, err := s.BillingEntities.GetBillingEntity(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Billing entity ID                                     |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetBillingEntityResponse](../../models/dtos/getbillingentityresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateBillingEntity

Edit a billing entity. Drafts and future invoices use the new details; finalized invoices keep theirs.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateBillingEntity" method="put" path="/billing-entities/{id}" -->
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

    res, err := s.BillingEntities.UpdateBillingEntity(ctx, "<id>", types.UpdateBillingEntityRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingEntityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | `string`                                                                             | :heavy_check_mark:                                                                   | Billing entity ID                                                                    |
| `body`                                                                               | [types.UpdateBillingEntityRequest](../../models/types/updatebillingentityrequest.md) | :heavy_check_mark:                                                                   | Fields to change                                                                     |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.UpdateBillingEntityResponse](../../models/dtos/updatebillingentityresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404, 409        | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |