# Integrations

## Overview

### Available Operations

* [LinkIntegrationMapping](#linkintegrationmapping) - Link integration mapping

## LinkIntegrationMapping

Link a FlexPrice entity to provider entity with provider-specific side effects.

### Example Usage

<!-- UsageSnippet language="go" operationID="linkIntegrationMapping" method="post" path="/integrations/mappings/link" -->
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

    res, err := s.Integrations.LinkIntegrationMapping(ctx, types.DtoLinkIntegrationMappingRequest{
        EntityID: "<id>",
        EntityType: types.IntegrationEntityTypeItemPrice,
        ProviderEntityID: "<id>",
        ProviderType: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoLinkIntegrationMappingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [types.DtoLinkIntegrationMappingRequest](../../models/types/dtolinkintegrationmappingrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][dtos.Option](../../models/dtos/option.md)                                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*dtos.LinkIntegrationMappingResponse](../../models/dtos/linkintegrationmappingresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |