# Environments

## Overview

### Available Operations

* [CloneEnvironment](#cloneenvironment) - Clone an environment

## CloneEnvironment

Clone all published features and plans from the source environment into a target environment. If target_environment_id is provided, entities are cloned into that existing environment. Otherwise a new environment is created from name and type first.

### Example Usage

<!-- UsageSnippet language="go" operationID="cloneEnvironment" method="post" path="/environments/{id}/clone" -->
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

    res, err := s.Environments.CloneEnvironment(ctx, "<id>", types.CloneEnvironmentRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ModelsTemporalWorkflowResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | `string`                                                                       | :heavy_check_mark:                                                             | Source Environment ID                                                          |
| `body`                                                                         | [types.CloneEnvironmentRequest](../../models/types/cloneenvironmentrequest.md) | :heavy_check_mark:                                                             | Clone configuration                                                            |
| `opts`                                                                         | [][dtos.Option](../../models/dtos/option.md)                                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*dtos.CloneEnvironmentResponse](../../models/dtos/cloneenvironmentresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |