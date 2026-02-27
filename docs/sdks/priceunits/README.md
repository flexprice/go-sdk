# PriceUnits

## Overview

### Available Operations

* [ListPriceUnits](#listpriceunits) - List price units
* [CreatePriceUnit](#createpriceunit) - Create price unit
* [GetPriceUnitByCode](#getpriceunitbycode) - Get price unit by code
* [QueryPriceUnit](#querypriceunit) - Query price units
* [GetPriceUnit](#getpriceunit) - Get price unit
* [UpdatePriceUnit](#updatepriceunit) - Update price unit
* [DeletePriceUnit](#deletepriceunit) - Delete price unit

## ListPriceUnits

Use when listing price units (e.g. in a catalog or when creating prices). Returns a paginated list; supports status, sort, and pagination.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPriceUnits" method="get" path="/prices/units" -->
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

    res, err := s.PriceUnits.ListPriceUnits(ctx, dtos.ListPriceUnitsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListPriceUnitsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `request`                                                         | [dtos.ListPriceUnitsRequest](../../dtos/listpriceunitsrequest.md) | :heavy_check_mark:                                                | The request object to use for the request.                        |
| `opts`                                                            | [][dtos.Option](../../dtos/option.md)                             | :heavy_minus_sign:                                                | The options for this request.                                     |

### Response

**[*dtos.ListPriceUnitsResponse](../../dtos/listpriceunitsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## CreatePriceUnit

Use when defining a new unit of measure for pricing (e.g. GB, API call, seat). Ideal for metered or usage-based prices.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPriceUnit" method="post" path="/prices/units" -->
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

    res, err := s.PriceUnits.CreatePriceUnit(ctx, types.DtoCreatePriceUnitRequest{
        BaseCurrency: "<value>",
        Code: "<value>",
        ConversionRate: "<value>",
        Name: "<value>",
        Symbol: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreatePriceUnitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |
| `request`                                                                   | [types.DtoCreatePriceUnitRequest](../../types/dtocreatepriceunitrequest.md) | :heavy_check_mark:                                                          | The request object to use for the request.                                  |
| `opts`                                                                      | [][dtos.Option](../../dtos/option.md)                                       | :heavy_minus_sign:                                                          | The options for this request.                                               |

### Response

**[*dtos.CreatePriceUnitResponse](../../dtos/createpriceunitresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetPriceUnitByCode

Use when resolving a price unit by code (e.g. from an external catalog or config). Ideal for integrations.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPriceUnitByCode" method="get" path="/prices/units/code/{code}" -->
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

    res, err := s.PriceUnits.GetPriceUnitByCode(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPriceUnitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `code`                                                | *string*                                              | :heavy_check_mark:                                    | Price unit code                                       |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetPriceUnitByCodeResponse](../../dtos/getpriceunitbycoderesponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## QueryPriceUnit

Use when searching or listing price units (e.g. admin catalog). Returns a paginated list; supports filtering and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryPriceUnit" method="post" path="/prices/units/search" -->
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

    res, err := s.PriceUnits.QueryPriceUnit(ctx, types.PriceUnitFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListPriceUnitsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                               | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ctx`                                                   | [context.Context](https://pkg.go.dev/context#Context)   | :heavy_check_mark:                                      | The context to use for the request.                     |
| `request`                                               | [types.PriceUnitFilter](../../types/priceunitfilter.md) | :heavy_check_mark:                                      | The request object to use for the request.              |
| `opts`                                                  | [][dtos.Option](../../dtos/option.md)                   | :heavy_minus_sign:                                      | The options for this request.                           |

### Response

**[*dtos.QueryPriceUnitResponse](../../dtos/querypriceunitresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetPriceUnit

Use when you need to load a single price unit (e.g. for display or when creating a price).

### Example Usage

<!-- UsageSnippet language="go" operationID="getPriceUnit" method="get" path="/prices/units/{id}" -->
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

    res, err := s.PriceUnits.GetPriceUnit(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPriceUnitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Price unit ID                                         |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetPriceUnitResponse](../../dtos/getpriceunitresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdatePriceUnit

Use when renaming or updating metadata for a price unit. Code is immutable once created.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePriceUnit" method="put" path="/prices/units/{id}" -->
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

    res, err := s.PriceUnits.UpdatePriceUnit(ctx, "<id>", types.DtoUpdatePriceUnitRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoPriceUnitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |
| `id`                                                                        | *string*                                                                    | :heavy_check_mark:                                                          | Price unit ID                                                               |
| `body`                                                                      | [types.DtoUpdatePriceUnitRequest](../../types/dtoupdatepriceunitrequest.md) | :heavy_check_mark:                                                          | Price unit details to update                                                |
| `opts`                                                                      | [][dtos.Option](../../dtos/option.md)                                       | :heavy_minus_sign:                                                          | The options for this request.                                               |

### Response

**[*dtos.UpdatePriceUnitResponse](../../dtos/updatepriceunitresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## DeletePriceUnit

Use when removing a price unit that is no longer needed. Fails if any price references this unit.

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePriceUnit" method="delete" path="/prices/units/{id}" -->
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

    res, err := s.PriceUnits.DeletePriceUnit(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Price unit ID                                         |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeletePriceUnitResponse](../../dtos/deletepriceunitresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |