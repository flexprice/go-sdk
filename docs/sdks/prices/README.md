# Prices

## Overview

### Available Operations

* [CreatePrice](#createprice) - Create price
* [CreatePricesBulk](#createpricesbulk) - Create prices in bulk
* [GetPriceByLookupKey](#getpricebylookupkey) - Get price by lookup key
* [QueryPrice](#queryprice) - Query prices
* [GetPrice](#getprice) - Get price
* [UpdatePrice](#updateprice) - Update price
* [DeletePrice](#deleteprice) - Delete price

## CreatePrice

Use when adding a new price to a plan or catalog (e.g. per-seat, flat, or metered). Ideal for both simple and usage-based pricing.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPrice" method="post" path="/prices" -->
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

    res, err := s.Prices.CreatePrice(ctx, types.CreatePriceRequest{
        BillingModel: types.BillingModelFlatFee,
        BillingPeriod: types.BillingPeriodDaily,
        Currency: "Dong",
        EntityID: "<id>",
        EntityType: types.PriceEntityTypePrice,
        InvoiceCadence: types.InvoiceCadenceAdvance,
        PriceUnitType: types.PriceUnitTypeFiat,
        Type: types.PriceTypeFixed,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PriceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [types.CreatePriceRequest](../../models/types/createpricerequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.CreatePriceResponse](../../models/dtos/createpriceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CreatePricesBulk

Use when creating many prices at once (e.g. importing a catalog or setting up a plan with multiple tiers).

### Example Usage

<!-- UsageSnippet language="go" operationID="createPricesBulk" method="post" path="/prices/bulk" -->
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

    res, err := s.Prices.CreatePricesBulk(ctx, types.CreateBulkPriceRequest{
        Items: []types.CreatePriceRequest{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateBulkPriceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [types.CreateBulkPriceRequest](../../models/types/createbulkpricerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.CreatePricesBulkResponse](../../models/dtos/createpricesbulkresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetPriceByLookupKey

Use when resolving a price by external id (e.g. from your catalog or CMS). Ideal for integrations.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPriceByLookupKey" method="get" path="/prices/lookup/{lookup_key}" -->
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

    res, err := s.Prices.GetPriceByLookupKey(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.PriceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `lookupKey`                                           | `string`                                              | :heavy_check_mark:                                    | Lookup key                                            |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetPriceByLookupKeyResponse](../../models/dtos/getpricebylookupkeyresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryPrice

Use when listing or searching prices (e.g. plan builder or catalog). Returns a paginated list; supports filtering and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryPrice" method="post" path="/prices/search" -->
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

    res, err := s.Prices.QueryPrice(ctx, types.PriceFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPricesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [types.PriceFilter](../../models/types/pricefilter.md) | :heavy_check_mark:                                     | The request object to use for the request.             |
| `opts`                                                 | [][dtos.Option](../../models/dtos/option.md)           | :heavy_minus_sign:                                     | The options for this request.                          |

### Response

**[*dtos.QueryPriceResponse](../../models/dtos/querypriceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetPrice

Use when you need to load a single price (e.g. for display or editing). Response includes expanded meter and price unit when applicable.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPrice" method="get" path="/prices/{id}" -->
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

    res, err := s.Prices.GetPrice(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.PriceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Price ID                                              |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetPriceResponse](../../models/dtos/getpriceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdatePrice

Use when changing price configuration (e.g. amount, billing scheme, or metadata).

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePrice" method="put" path="/prices/{id}" -->
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

    res, err := s.Prices.UpdatePrice(ctx, "<id>", types.UpdatePriceRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.PriceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `id`                                                                 | `string`                                                             | :heavy_check_mark:                                                   | Price ID                                                             |
| `body`                                                               | [types.UpdatePriceRequest](../../models/types/updatepricerequest.md) | :heavy_check_mark:                                                   | Price configuration                                                  |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.UpdatePriceResponse](../../models/dtos/updatepriceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeletePrice

Use when retiring a price (e.g. end-of-life or replacement). Optional effective date or cascade for subscriptions.

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePrice" method="delete" path="/prices/{id}" -->
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

    res, err := s.Prices.DeletePrice(ctx, "<id>", types.DeletePriceRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `id`                                                                 | `string`                                                             | :heavy_check_mark:                                                   | Price ID                                                             |
| `body`                                                               | [types.DeletePriceRequest](../../models/types/deletepricerequest.md) | :heavy_check_mark:                                                   | Delete Price Request                                                 |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.DeletePriceResponse](../../models/dtos/deletepriceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |