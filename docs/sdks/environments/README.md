# Environments

## Overview

### Available Operations

* [CreateEnvironment](#createenvironment) - Create environment
* [UpdateEnvironment](#updateenvironment) - Update environment

## CreateEnvironment

Use when setting up a new environment (e.g. production, staging) for the tenant. Ideal for separating billing or config per environment.

### Example Usage

<!-- UsageSnippet language="go" operationID="createEnvironment" method="post" path="/environments" -->
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

    res, err := s.Environments.CreateEnvironment(ctx, types.DtoCreateEnvironmentRequest{
        Name: "<value>",
        Type: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoEnvironmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [types.DtoCreateEnvironmentRequest](../../models/types/dtocreateenvironmentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*dtos.CreateEnvironmentResponse](../../models/dtos/createenvironmentresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateEnvironment

Use when changing environment name or settings (e.g. renaming or updating metadata).

### Example Usage

<!-- UsageSnippet language="go" operationID="updateEnvironment" method="put" path="/environments/{id}" -->
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

    res, err := s.Environments.UpdateEnvironment(ctx, "<id>", types.DtoUpdateEnvironmentRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoEnvironmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | Environment ID                                                                         |
| `body`                                                                                 | [types.DtoUpdateEnvironmentRequest](../../models/types/dtoupdateenvironmentrequest.md) | :heavy_check_mark:                                                                     | Environment                                                                            |
| `opts`                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*dtos.UpdateEnvironmentResponse](../../models/dtos/updateenvironmentresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |