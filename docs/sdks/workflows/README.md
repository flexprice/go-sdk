# Workflows

## Overview

### Available Operations

* [QueryWorkflow](#queryworkflow) - Query workflows

## QueryWorkflow

Use when listing or auditing workflow runs (e.g. ops dashboard or debugging). Returns a paginated list; supports filtering by workflow type and status.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryWorkflow" method="post" path="/workflows/search" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Workflows.QueryWorkflow(ctx, types.WorkflowExecutionFilter{}, dtos.QueryWorkflowSecurity{
        Option1: &dtos.QueryWorkflowSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListWorkflowsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [types.WorkflowExecutionFilter](../../models/types/workflowexecutionfilter.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [dtos.QueryWorkflowSecurity](../../models/dtos/queryworkflowsecurity.md)       | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |
| `opts`                                                                         | [][dtos.Option](../../models/dtos/option.md)                                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*dtos.QueryWorkflowResponse](../../models/dtos/queryworkflowresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |