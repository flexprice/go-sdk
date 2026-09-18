# Tasks

## Overview

### Available Operations

* [ListTasks](#listtasks) - List tasks
* [CreateTask](#createtask) - Import a CSV of usage events
* [GetTaskResult](#gettaskresult) - Get task processing result
* [GetTask](#gettask) - Get a task
* [DownloadTaskExport](#downloadtaskexport) - Download task export file
* [UpdateTaskStatus](#updatetaskstatus) - Update task status

## ListTasks

Use when listing or searching async tasks (e.g. admin queue view). Returns list with optional filtering.

### Example Usage

<!-- UsageSnippet language="go" operationID="listTasks" method="get" path="/tasks" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Tasks.ListTasks(ctx, dtos.ListTasksRequest{}, dtos.ListTasksSecurity{
        Option1: &dtos.ListTasksSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListTasksResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [dtos.ListTasksRequest](../../models/dtos/listtasksrequest.md)   | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `security`                                                       | [dtos.ListTasksSecurity](../../models/dtos/listtaskssecurity.md) | :heavy_check_mark:                                               | The security requirements to use for the request.                |
| `opts`                                                           | [][dtos.Option](../../models/dtos/option.md)                     | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*dtos.ListTasksResponse](../../models/dtos/listtasksresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CreateTask

Use to submit a CSV of usage events for async ingestion. The CSV must already have been uploaded to the Flexprice-managed imports bucket (currently via CSV Box) — pass the upload_id and the backend fetches the file from S3 and streams rows into ClickHouse. Returns the task ID and Temporal workflow IDs for polling.

### Example Usage

<!-- UsageSnippet language="go" operationID="createTask" method="post" path="/tasks" -->
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

    res, err := s.Tasks.CreateTask(ctx, types.CreateTaskRequest{
        EntityType: types.EntityTypeFeatures,
        FileProvider: "<value>",
        FileType: types.FileTypeJSON,
        TaskType: types.TaskTypeExport,
        UploadID: "<id>",
    }, dtos.CreateTaskSecurity{
        Option1: &dtos.CreateTaskSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ModelsTemporalWorkflowResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [types.CreateTaskRequest](../../models/types/createtaskrequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `security`                                                         | [dtos.CreateTaskSecurity](../../models/dtos/createtasksecurity.md) | :heavy_check_mark:                                                 | The security requirements to use for the request.                  |
| `opts`                                                             | [][dtos.Option](../../models/dtos/option.md)                       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*dtos.CreateTaskResponse](../../models/dtos/createtaskresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetTaskResult

Use when fetching the outcome of a completed task (e.g. export URL or error message). Call after task status is complete.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTaskResult" method="get" path="/tasks/result" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Tasks.GetTaskResult(ctx, dtos.GetTaskResultSecurity{
        Option1: &dtos.GetTaskResultSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ModelsTemporalWorkflowResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `security`                                                               | [dtos.GetTaskResultSecurity](../../models/dtos/gettaskresultsecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |
| `workflowID`                                                             | `string`                                                                 | :heavy_check_mark:                                                       | Workflow ID                                                              |
| `opts`                                                                   | [][dtos.Option](../../models/dtos/option.md)                             | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*dtos.GetTaskResultResponse](../../models/dtos/gettaskresultresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetTask

Use when checking task status or progress (e.g. polling after create). Returns task by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTask" method="get" path="/tasks/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Tasks.GetTask(ctx, dtos.GetTaskSecurity{
        Option1: &dtos.GetTaskSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `security`                                                   | [dtos.GetTaskSecurity](../../models/dtos/gettasksecurity.md) | :heavy_check_mark:                                           | The security requirements to use for the request.            |
| `id`                                                         | `string`                                                     | :heavy_check_mark:                                           | Task ID                                                      |
| `opts`                                                       | [][dtos.Option](../../models/dtos/option.md)                 | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*dtos.GetTaskResponse](../../models/dtos/gettaskresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DownloadTaskExport

Use when letting a user download an exported file (e.g. report or data export). Returns a presigned URL; supports FlexPrice or customer-owned S3.

### Example Usage

<!-- UsageSnippet language="go" operationID="downloadTaskExport" method="get" path="/tasks/{id}/download" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Tasks.DownloadTaskExport(ctx, dtos.DownloadTaskExportSecurity{
        Option1: &dtos.DownloadTaskExportSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `security`                                                                         | [dtos.DownloadTaskExportSecurity](../../models/dtos/downloadtaskexportsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | Task ID                                                                            |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.DownloadTaskExportResponse](../../models/dtos/downloadtaskexportresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateTaskStatus

Use when updating task status (e.g. marking complete or failed from a worker). Typically called by backend processors.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateTaskStatus" method="put" path="/tasks/{id}/status" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Tasks.UpdateTaskStatus(ctx, dtos.UpdateTaskStatusSecurity{
        Option1: &dtos.UpdateTaskStatusSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", types.UpdateTaskStatusRequest{
        TaskStatus: types.TaskStatusProcessing,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `security`                                                                     | [dtos.UpdateTaskStatusSecurity](../../models/dtos/updatetaskstatussecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |
| `id`                                                                           | `string`                                                                       | :heavy_check_mark:                                                             | Task ID                                                                        |
| `body`                                                                         | [types.UpdateTaskStatusRequest](../../models/types/updatetaskstatusrequest.md) | :heavy_check_mark:                                                             | Status update                                                                  |
| `opts`                                                                         | [][dtos.Option](../../models/dtos/option.md)                                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*dtos.UpdateTaskStatusResponse](../../models/dtos/updatetaskstatusresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |