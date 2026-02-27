# Integrations

## Overview

### Available Operations

* [GetIntegration](#getintegration) - Get integration details
* [CreateOrUpdateIntegration](#createorupdateintegration) - Create or update an integration
* [ListLinkedIntegrations](#listlinkedintegrations) - List linked integrations
* [DeleteIntegration](#deleteintegration) - Delete an integration

## GetIntegration

Use when you need to check or display integration config (e.g. which provider is linked). Sensitive values may be redacted.

### Example Usage

<!-- UsageSnippet language="go" operationID="getIntegration" method="get" path="/secrets/integrations/by-provider/{provider}" -->
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

    res, err := s.Integrations.GetIntegration(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSecretResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `provider`                                            | *string*                                              | :heavy_check_mark:                                    | Integration provider                                  |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetIntegrationResponse](../../dtos/getintegrationresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 404                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## CreateOrUpdateIntegration

Use when storing or updating credentials for an external integration (e.g. Stripe, HubSpot). Secrets are encrypted at rest.

### Example Usage

<!-- UsageSnippet language="go" operationID="createOrUpdateIntegration" method="post" path="/secrets/integrations/create/{provider}" -->
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

    res, err := s.Integrations.CreateOrUpdateIntegration(ctx, "<value>", types.DtoCreateIntegrationRequest{
        Credentials: map[string]string{
            "key": "<value>",
        },
        Name: "<value>",
        Provider: types.SecretProviderChargebee,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSecretResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |
| `provider`                                                                      | *string*                                                                        | :heavy_check_mark:                                                              | Integration provider                                                            |
| `body`                                                                          | [types.DtoCreateIntegrationRequest](../../types/dtocreateintegrationrequest.md) | :heavy_check_mark:                                                              | Integration creation request                                                    |
| `opts`                                                                          | [][dtos.Option](../../dtos/option.md)                                           | :heavy_minus_sign:                                                              | The options for this request.                                                   |

### Response

**[*dtos.CreateOrUpdateIntegrationResponse](../../dtos/createorupdateintegrationresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## ListLinkedIntegrations

Use when showing which integrations are connected (e.g. settings page). Returns providers that have valid linked credentials.

### Example Usage

<!-- UsageSnippet language="go" operationID="listLinkedIntegrations" method="get" path="/secrets/integrations/linked" -->
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

    res, err := s.Integrations.ListLinkedIntegrations(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoLinkedIntegrationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.ListLinkedIntegrationsResponse](../../dtos/listlinkedintegrationsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## DeleteIntegration

Use when disconnecting an integration (e.g. switching provider or removing OAuth). Deletes stored credentials.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteIntegration" method="delete" path="/secrets/integrations/{id}" -->
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

    res, err := s.Integrations.DeleteIntegration(ctx, "<id>")
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Integration ID                                        |
| `opts`                                                | [][dtos.Option](../../dtos/option.md)                 | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeleteIntegrationResponse](../../dtos/deleteintegrationresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 404                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |