# FXRates

## Overview

### Available Operations

* [CreateFXRate](#createfxrate) - Create an FX rate
* [QueryFXRates](#queryfxrates) - Query FX rates
* [GetFXRate](#getfxrate) - Get an FX rate
* [UpdateFXRate](#updatefxrate) - Update an FX rate
* [DeleteFXRate](#deletefxrate) - Delete an FX rate

## CreateFXRate

Configure a fixed exchange rate at tenant, customer or subscription scope. Overrides need a tenant rate for the same pair.

### Example Usage

<!-- UsageSnippet language="go" operationID="createFXRate" method="post" path="/forex" -->
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

    res, err := s.FXRates.CreateFXRate(ctx, types.CreateFXRateRequest{
        FromCurrency: "<value>",
        Scope: types.FXRateScopeTenant,
        ToCurrency: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FXRateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [types.CreateFXRateRequest](../../models/types/createfxraterequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][dtos.Option](../../models/dtos/option.md)                           | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*dtos.CreateFXRateResponse](../../models/dtos/createfxrateresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryFXRates

Filter FX rates via a request body (POST used for a complex query, but read-only).

### Example Usage

<!-- UsageSnippet language="go" operationID="queryFXRates" method="post" path="/forex/query" -->
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

    res, err := s.FXRates.QueryFXRates(ctx, types.FXRateFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListFXRatesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [types.FXRateFilter](../../models/types/fxratefilter.md) | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][dtos.Option](../../models/dtos/option.md)             | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*dtos.QueryFXRatesResponse](../../models/dtos/queryfxratesresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetFXRate

Load a single FX rate by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getFXRate" method="get" path="/forex/{id}" -->
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

    res, err := s.FXRates.GetFXRate(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.FXRateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | FX rate ID                                            |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetFXRateResponse](../../models/dtos/getfxrateresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateFXRate

Update a rate's value, validity window (overrides only) or metadata. Scope, scope_id and the currency pair are immutable.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateFXRate" method="put" path="/forex/{id}" -->
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

    res, err := s.FXRates.UpdateFXRate(ctx, "<id>", types.UpdateFXRateRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.FXRateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `id`                                                                   | `string`                                                               | :heavy_check_mark:                                                     | FX rate ID                                                             |
| `body`                                                                 | [types.UpdateFXRateRequest](../../models/types/updatefxraterequest.md) | :heavy_check_mark:                                                     | FX rate fields to update                                               |
| `opts`                                                                 | [][dtos.Option](../../models/dtos/option.md)                           | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*dtos.UpdateFXRateResponse](../../models/dtos/updatefxrateresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeleteFXRate

Archive a customer or subscription override. Tenant rates cannot be deleted.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteFXRate" method="delete" path="/forex/{id}" -->
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

    res, err := s.FXRates.DeleteFXRate(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | FX rate ID                                            |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeleteFXRateResponse](../../models/dtos/deletefxrateresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |