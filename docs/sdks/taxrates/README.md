# TaxRates

## Overview

### Available Operations

* [GetTaxRates](#gettaxrates) - Get tax rates
* [CreateTaxRate](#createtaxrate) - Create a tax rate
* [GetTaxRate](#gettaxrate) - Get a tax rate
* [UpdateTaxRate](#updatetaxrate) - Update a tax rate
* [DeleteTaxRate](#deletetaxrate) - Delete a tax rate

## GetTaxRates

Use when listing tax rates (e.g. tax config UI). Returns tax rates with optional filters.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTaxRates" method="get" path="/taxes/rates" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.TaxRates.GetTaxRates(ctx, dtos.GetTaxRatesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxRateResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `request`                                                   | [dtos.GetTaxRatesRequest](../../dtos/gettaxratesrequest.md) | :heavy_check_mark:                                          | The request object to use for the request.                  |
| `opts`                                                      | [][dtos.Option](../../dtos/option.md)                       | :heavy_minus_sign:                                          | The options for this request.                               |

### Response

**[*dtos.GetTaxRatesResponse](../../dtos/gettaxratesresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## CreateTaxRate

Use when defining a new tax rate (e.g. VAT or sales tax) for use in invoices. Attach to customers or products via tax associations.

### Example Usage

<!-- UsageSnippet language="go" operationID="createTaxRate" method="post" path="/taxes/rates" -->
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

    res, err := s.TaxRates.CreateTaxRate(ctx, types.DtoCreateTaxRateRequest{
        Code: "<value>",
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxRateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `request`                                                               | [types.DtoCreateTaxRateRequest](../../types/dtocreatetaxraterequest.md) | :heavy_check_mark:                                                      | The request object to use for the request.                              |
| `opts`                                                                  | [][dtos.Option](../../dtos/option.md)                                   | :heavy_minus_sign:                                                      | The options for this request.                                           |

### Response

**[*dtos.CreateTaxRateResponse](../../dtos/createtaxrateresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetTaxRate

Use when you need to load a single tax rate (e.g. for display or when creating an association).

### Example Usage

<!-- UsageSnippet language="go" operationID="getTaxRate" method="get" path="/taxes/rates/{id}" -->
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

    res, err := s.TaxRates.GetTaxRate(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxRateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Tax rate ID                                           |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetTaxRateResponse](../../dtos/gettaxrateresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateTaxRate

Use when changing a tax rate (e.g. rate value or name). Request body contains the fields to update.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTaxRate" method="put" path="/taxes/rates/{id}" -->
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

    res, err := s.TaxRates.UpdateTaxRate(ctx, "<id>", types.DtoUpdateTaxRateRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTaxRateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `id`                                                                    | *string*                                                                | :heavy_check_mark:                                                      | Tax rate ID                                                             |
| `body`                                                                  | [types.DtoUpdateTaxRateRequest](../../types/dtoupdatetaxraterequest.md) | :heavy_check_mark:                                                      | Tax rate to update                                                      |
| `opts`                                                                  | [][dtos.Option](../../dtos/option.md)                                   | :heavy_minus_sign:                                                      | The options for this request.                                           |

### Response

**[*dtos.UpdateTaxRateResponse](../../dtos/updatetaxrateresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## DeleteTaxRate

Use when retiring a tax rate (e.g. no longer applicable). Fails if still referenced by associations.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteTaxRate" method="delete" path="/taxes/rates/{id}" -->
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

    res, err := s.TaxRates.DeleteTaxRate(ctx, "<id>")
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Tax rate ID                                           |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeleteTaxRateResponse](../../dtos/deletetaxrateresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |