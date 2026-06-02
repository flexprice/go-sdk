# Integrations

## Overview

### Available Operations

* [GetIntegrationConfig](#getintegrationconfig) - Get integration configurations
* [LinkIntegrationMapping](#linkintegrationmapping) - Link integration mapping
* [GetEntityIntegrationMappings](#getentityintegrationmappings) - Get entity integration mappings

## GetIntegrationConfig

Returns the base capabilities and current sync configuration for all connected providers.

### Example Usage

<!-- UsageSnippet language="go" operationID="getIntegrationConfig" method="get" path="/integrations/config" -->
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

    res, err := s.Integrations.GetIntegrationConfig(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationConfigResponse != nil {
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

**[*dtos.GetIntegrationConfigResponse](../../models/dtos/getintegrationconfigresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## LinkIntegrationMapping

Link a FlexPrice entity to provider entity with provider-specific side effects.

### Example Usage

<!-- UsageSnippet language="go" operationID="linkIntegrationMapping" method="post" path="/integrations/link" -->
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

    res, err := s.Integrations.LinkIntegrationMapping(ctx, types.LinkIntegrationMappingRequest{
        EntityID: "<id>",
        EntityType: types.IntegrationEntityTypeItemPrice,
        ProviderEntityID: "<id>",
        ProviderType: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LinkIntegrationMappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [types.LinkIntegrationMappingRequest](../../models/types/linkintegrationmappingrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][dtos.Option](../../models/dtos/option.md)                                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*dtos.LinkIntegrationMappingResponse](../../models/dtos/linkintegrationmappingresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetEntityIntegrationMappings

Get integration mappings for a specific entity by entity type and entity ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getEntityIntegrationMappings" method="get" path="/integrations/mappings" -->
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

    res, err := s.Integrations.GetEntityIntegrationMappings(ctx, "<value>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ListEntityIntegrationMappingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |
| `entityType`                                                                                              | `string`                                                                                                  | :heavy_check_mark:                                                                                        | Entity type (customer, plan, invoice, subscription, payment, credit_note, addon, item, item_price, price) |
| `entityID`                                                                                                | `string`                                                                                                  | :heavy_check_mark:                                                                                        | Entity ID                                                                                                 |
| `opts`                                                                                                    | [][dtos.Option](../../models/dtos/option.md)                                                              | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |

### Response

**[*dtos.GetEntityIntegrationMappingsResponse](../../models/dtos/getentityintegrationmappingsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |