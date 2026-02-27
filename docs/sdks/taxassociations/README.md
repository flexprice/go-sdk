# TaxAssociations

## Overview

### Available Operations

* [ListTaxAssociations](#listtaxassociations) - List tax associations
* [CreateTaxAssociation](#createtaxassociation) - Create Tax Association
* [GetTaxAssociation](#gettaxassociation) - Get Tax Association
* [UpdateTaxAssociation](#updatetaxassociation) - Update tax association
* [DeleteTaxAssociation](#deletetaxassociation) - Delete tax association

## ListTaxAssociations

Use when listing tax associations (e.g. tax config or audit). Returns list with optional filtering.

### Example Usage

<!-- UsageSnippet language="go" operationID="listTaxAssociations" method="get" path="/taxes/associations" -->
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
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.TaxAssociations.ListTaxAssociations(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListTaxAssociationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `entityType`                                          | **string*                                             | :heavy_minus_sign:                                    | Entity Type                                           |
| `entityID`                                            | **string*                                             | :heavy_minus_sign:                                    | Entity ID                                             |
| `externalCustomerID`                                  | **string*                                             | :heavy_minus_sign:                                    | External Customer ID                                  |
| `taxRateID`                                           | **string*                                             | :heavy_minus_sign:                                    | Tax Rate ID                                           |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.ListTaxAssociationsResponse](../../dtos/listtaxassociationsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## CreateTaxAssociation

Use when linking a tax rate to an entity (e.g. customer, product, or region) so that rate applies on invoices.

### Example Usage

<!-- UsageSnippet language="go" operationID="createTaxAssociation" method="post" path="/taxes/associations" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.TaxAssociations.CreateTaxAssociation(ctx, types.DtoCreateTaxAssociationRequest{
        TaxRateCode: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxAssociationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `request`                                                                             | [types.DtoCreateTaxAssociationRequest](../../types/dtocreatetaxassociationrequest.md) | :heavy_check_mark:                                                                    | The request object to use for the request.                                            |
| `opts`                                                                                | [][dtos.Option](../../dtos/option.md)                                                 | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

### Response

**[*dtos.CreateTaxAssociationResponse](../../dtos/createtaxassociationresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetTaxAssociation

Use when you need to load a single tax association (e.g. for display or editing).

### Example Usage

<!-- UsageSnippet language="go" operationID="getTaxAssociation" method="get" path="/taxes/associations/{id}" -->
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
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.TaxAssociations.GetTaxAssociation(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxAssociationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Tax Config ID                                         |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetTaxAssociationResponse](../../dtos/gettaxassociationresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateTaxAssociation

Use when changing a tax association (e.g. switch rate or entity). Request body contains the fields to update.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTaxAssociation" method="put" path="/taxes/associations/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.TaxAssociations.UpdateTaxAssociation(ctx, "<id>", types.DtoTaxAssociationUpdateRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxAssociationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `id`                                                                                  | *string*                                                                              | :heavy_check_mark:                                                                    | Tax Config ID                                                                         |
| `body`                                                                                | [types.DtoTaxAssociationUpdateRequest](../../types/dtotaxassociationupdaterequest.md) | :heavy_check_mark:                                                                    | Tax Config Request                                                                    |
| `opts`                                                                                | [][dtos.Option](../../dtos/option.md)                                                 | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

### Response

**[*dtos.UpdateTaxAssociationResponse](../../dtos/updatetaxassociationresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## DeleteTaxAssociation

Use when removing a tax association (e.g. entity no longer subject to that rate).

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteTaxAssociation" method="delete" path="/taxes/associations/{id}" -->
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
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.TaxAssociations.DeleteTaxAssociation(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxAssociationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Tax Config ID                                         |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeleteTaxAssociationResponse](../../dtos/deletetaxassociationresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |