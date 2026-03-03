# EntityIntegrationMappings

## Overview

### Available Operations

* [CreateEntityIntegrationMapping](#createentityintegrationmapping) - Create entity integration mapping
* [DeleteEntityIntegrationMapping](#deleteentityintegrationmapping) - Delete entity integration mapping

## CreateEntityIntegrationMapping

Use when linking a FlexPrice entity to an external system (e.g. CRM or payment provider) so you can sync or reconcile by external ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="createEntityIntegrationMapping" method="post" path="/entity-integration-mappings" -->
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

    res, err := s.EntityIntegrationMappings.CreateEntityIntegrationMapping(ctx, types.DtoCreateEntityIntegrationMappingRequest{
        EntityID: "<id>",
        EntityType: types.IntegrationEntityTypeCreditNote,
        ProviderEntityID: "<id>",
        ProviderType: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoEntityIntegrationMappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [types.DtoCreateEntityIntegrationMappingRequest](../../models/types/dtocreateentityintegrationmappingrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][dtos.Option](../../models/dtos/option.md)                                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*dtos.CreateEntityIntegrationMappingResponse](../../models/dtos/createentityintegrationmappingresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 401, 409              | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## DeleteEntityIntegrationMapping

Use when unlinking a FlexPrice entity from an external system or cleaning up stale integration mappings.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteEntityIntegrationMapping" method="delete" path="/entity-integration-mappings/{id}" -->
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

    res, err := s.EntityIntegrationMappings.DeleteEntityIntegrationMapping(ctx, "<id>")
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Entity integration mapping ID                         |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeleteEntityIntegrationMappingResponse](../../models/dtos/deleteentityintegrationmappingresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 401, 404              | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |