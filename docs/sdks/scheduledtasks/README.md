# ScheduledTasks

## Overview

### Available Operations

* [ListScheduledTasks](#listscheduledtasks) - List scheduled tasks
* [CreateScheduledTask](#createscheduledtask) - Create scheduled task
* [ScheduleUpdateBillingPeriod](#scheduleupdatebillingperiod) - Schedule update billing period
* [GetScheduledTask](#getscheduledtask) - Get scheduled task
* [UpdateScheduledTask](#updatescheduledtask) - Update a scheduled task
* [DeleteScheduledTask](#deletescheduledtask) - Delete a scheduled task
* [TriggerScheduledTaskRun](#triggerscheduledtaskrun) - Trigger force run

## ListScheduledTasks

Use when listing or managing scheduled tasks in an admin UI. Returns a list; supports filtering by status, type, and pagination.

### Example Usage

<!-- UsageSnippet language="go" operationID="listScheduledTasks" method="get" path="/tasks/scheduled" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.ScheduledTasks.ListScheduledTasks(ctx, dtos.ListScheduledTasksRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListScheduledTasksResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [dtos.ListScheduledTasksRequest](../../models/dtos/listscheduledtasksrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][dtos.Option](../../models/dtos/option.md)                                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*dtos.ListScheduledTasksResponse](../../models/dtos/listscheduledtasksresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## CreateScheduledTask

Use when setting up recurring data exports or other scheduled jobs. Ideal for report generation or syncing data on a schedule.

### Example Usage

<!-- UsageSnippet language="go" operationID="createScheduledTask" method="post" path="/tasks/scheduled" -->
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

    res, err := s.ScheduledTasks.CreateScheduledTask(ctx, types.DtoCreateScheduledTaskRequest{
        ConnectionID: "<id>",
        EntityType: types.ScheduledTaskEntityTypeCreditTopups,
        Interval: types.ScheduledTaskIntervalHourly,
        JobConfig: types.S3JobConfig{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoScheduledTaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [types.DtoCreateScheduledTaskRequest](../../models/types/dtocreatescheduledtaskrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][dtos.Option](../../models/dtos/option.md)                                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*dtos.CreateScheduledTaskResponse](../../models/dtos/createscheduledtaskresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## ScheduleUpdateBillingPeriod

Use when you need to trigger a billing-period update workflow (e.g. to recalculate or sync billing windows).

### Example Usage

<!-- UsageSnippet language="go" operationID="scheduleUpdateBillingPeriod" method="post" path="/tasks/scheduled/schedule-update-billing-period" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.ScheduledTasks.ScheduleUpdateBillingPeriod(ctx, dtos.ScheduleUpdateBillingPeriodRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [dtos.ScheduleUpdateBillingPeriodRequest](../../models/dtos/scheduleupdatebillingperiodrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][dtos.Option](../../models/dtos/option.md)                                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*dtos.ScheduleUpdateBillingPeriodResponse](../../models/dtos/scheduleupdatebillingperiodresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetScheduledTask

Use when you need to load a single scheduled task (e.g. to show details in a UI or check its configuration).

### Example Usage

<!-- UsageSnippet language="go" operationID="getScheduledTask" method="get" path="/tasks/scheduled/{id}" -->
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

    res, err := s.ScheduledTasks.GetScheduledTask(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoScheduledTaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Scheduled Task ID                                     |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetScheduledTaskResponse](../../models/dtos/getscheduledtaskresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateScheduledTask

Use when pausing or resuming a scheduled task. Only the enabled field can be changed.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateScheduledTask" method="put" path="/tasks/scheduled/{id}" -->
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

    res, err := s.ScheduledTasks.UpdateScheduledTask(ctx, "<id>", types.DtoUpdateScheduledTaskRequest{
        Enabled: false,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoScheduledTaskResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `id`                                                                                       | *string*                                                                                   | :heavy_check_mark:                                                                         | Scheduled Task ID                                                                          |
| `body`                                                                                     | [types.DtoUpdateScheduledTaskRequest](../../models/types/dtoupdatescheduledtaskrequest.md) | :heavy_check_mark:                                                                         | Update request (enabled: true/false to pause/resume)                                       |
| `opts`                                                                                     | [][dtos.Option](../../models/dtos/option.md)                                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*dtos.UpdateScheduledTaskResponse](../../models/dtos/updatescheduledtaskresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## DeleteScheduledTask

Use when removing a scheduled task from the active roster. Archives the task and removes it from the scheduler (soft delete).

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteScheduledTask" method="delete" path="/tasks/scheduled/{id}" -->
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

    res, err := s.ScheduledTasks.DeleteScheduledTask(ctx, "<id>")
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
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Scheduled Task ID                                     |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeleteScheduledTaskResponse](../../models/dtos/deletescheduledtaskresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## TriggerScheduledTaskRun

Use when you need to run a scheduled export immediately (e.g. on-demand report or catch-up). Supports optional custom time range.

### Example Usage

<!-- UsageSnippet language="go" operationID="triggerScheduledTaskRun" method="post" path="/tasks/scheduled/{id}/run" -->
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

    res, err := s.ScheduledTasks.TriggerScheduledTaskRun(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTriggerForceRunResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `id`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | Scheduled Task ID                                                                   |
| `body`                                                                              | [*types.DtoTriggerForceRunRequest](../../models/types/dtotriggerforcerunrequest.md) | :heavy_minus_sign:                                                                  | Optional start and end time for custom range                                        |
| `opts`                                                                              | [][dtos.Option](../../models/dtos/option.md)                                        | :heavy_minus_sign:                                                                  | The options for this request.                                                       |

### Response

**[*dtos.TriggerScheduledTaskRunResponse](../../models/dtos/triggerscheduledtaskrunresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |