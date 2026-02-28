# Events

## Overview

### Available Operations

* [IngestEvent](#ingestevent) - Ingest event
* [GetUsageAnalytics](#getusageanalytics) - Get usage analytics
* [IngestEventsBulk](#ingesteventsbulk) - Bulk ingest events
* [GetHuggingfaceInferenceData](#gethuggingfaceinferencedata) - Get Hugging Face inference data
* [ListRawEvents](#listrawevents) - List raw events
* [GetUsageStatistics](#getusagestatistics) - Get usage statistics
* [GetUsageByMeter](#getusagebymeter) - Get usage by meter
* [GetEvent](#getevent) - Get event

## IngestEvent

Use when sending a single usage event from your app (e.g. one API call or one GB stored). Events are processed asynchronously; returns 202 with event_id.

### Example Usage

<!-- UsageSnippet language="go" operationID="ingestEvent" method="post" path="/events" -->
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
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Events.IngestEvent(ctx, types.DtoIngestEventRequest{
        CustomerID: flexprice.Pointer("customer456"),
        EventID: flexprice.Pointer("event123"),
        EventName: "api_request",
        ExternalCustomerID: "customer456",
        Properties: map[string]string{
            ""response_status"": "200}",
            "{"request_size"": "100",
        },
        Source: flexprice.Pointer("api"),
        Timestamp: flexprice.Pointer("2024-03-20T15:04:05Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [types.DtoIngestEventRequest](../../models/types/dtoingesteventrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.IngestEventResponse](../../models/dtos/ingesteventresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetUsageAnalytics

Use when building analytics views (e.g. usage by feature or customer over time). Supports filtering, grouping, and time-series breakdown.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUsageAnalytics" method="post" path="/events/analytics" -->
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
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Events.GetUsageAnalytics(ctx, types.DtoGetUsageAnalyticsRequest{
        ExternalCustomerID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetUsageAnalyticsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [types.DtoGetUsageAnalyticsRequest](../../models/types/dtogetusageanalyticsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*dtos.GetUsageAnalyticsResponse](../../models/dtos/getusageanalyticsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## IngestEventsBulk

Use when batching usage events (e.g. backfill or high-volume ingestion). More efficient than single event calls; returns 202 when accepted.

### Example Usage

<!-- UsageSnippet language="go" operationID="ingestEventsBulk" method="post" path="/events/bulk" -->
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
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Events.IngestEventsBulk(ctx, types.DtoBulkIngestEventRequest{
        Events: []types.DtoIngestEventRequest{},
    })
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
| `request`                                                                          | [types.DtoBulkIngestEventRequest](../../models/types/dtobulkingesteventrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.IngestEventsBulkResponse](../../models/dtos/ingesteventsbulkresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetHuggingfaceInferenceData

Use when fetching Hugging Face inference usage or billing data (e.g. for HF-specific reporting or reconciliation).

### Example Usage

<!-- UsageSnippet language="go" operationID="getHuggingfaceInferenceData" method="post" path="/events/huggingface-inference" -->
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

    res, err := s.Events.GetHuggingfaceInferenceData(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetHuggingFaceBillingDataResponse != nil {
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

**[*dtos.GetHuggingfaceInferenceDataResponse](../../models/dtos/gethuggingfaceinferencedataresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## ListRawEvents

Use when debugging ingestion or exporting raw event data (e.g. support or audit). Returns a paginated list; supports time range and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="listRawEvents" method="post" path="/events/query" -->
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
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Events.ListRawEvents(ctx, types.DtoGetEventsRequest{
        EndTime: flexprice.Pointer("2024-12-09T00:00:00Z"),
        Order: flexprice.Pointer("desc"),
        Sort: flexprice.Pointer("timestamp"),
        StartTime: flexprice.Pointer("2024-11-09T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetEventsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [types.DtoGetEventsRequest](../../models/types/dtogeteventsrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][dtos.Option](../../models/dtos/option.md)                           | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*dtos.ListRawEventsResponse](../../models/dtos/listraweventsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetUsageStatistics

Use when building usage reports or dashboards across events. Supports filters and grouping; defaults to last 7 days if no range provided.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUsageStatistics" method="post" path="/events/usage" -->
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
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Events.GetUsageStatistics(ctx, types.DtoGetUsageRequest{
        AggregationType: types.AggregationTypeCountUnique,
        BillingAnchor: flexprice.Pointer("2024-03-05T14:30:45.123456789Z"),
        CustomerID: flexprice.Pointer("customer456"),
        EndTime: flexprice.Pointer("2024-03-20T00:00:00Z"),
        EventName: "api_request",
        ExternalCustomerID: flexprice.Pointer("customer456"),
        PropertyName: flexprice.Pointer("request_size"),
        StartTime: flexprice.Pointer("2024-03-13T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetUsageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [types.DtoGetUsageRequest](../../models/types/dtogetusagerequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.GetUsageStatisticsResponse](../../models/dtos/getusagestatisticsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetUsageByMeter

Use when showing usage for a specific meter (e.g. dashboard or overage check). Supports time range, filters, and grouping by customer or subscription.

### Example Usage

<!-- UsageSnippet language="go" operationID="getUsageByMeter" method="post" path="/events/usage/meter" -->
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
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Events.GetUsageByMeter(ctx, types.DtoGetUsageByMeterRequest{
        BillingAnchor: flexprice.Pointer("2024-03-05T14:30:45Z"),
        CustomerID: flexprice.Pointer("customer456"),
        EndTime: flexprice.Pointer("2024-12-09T00:00:00Z"),
        ExternalCustomerID: flexprice.Pointer("user_5"),
        MeterID: "123",
        StartTime: flexprice.Pointer("2024-11-09T00:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetUsageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [types.DtoGetUsageByMeterRequest](../../models/types/dtogetusagebymeterrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.GetUsageByMeterResponse](../../models/dtos/getusagebymeterresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetEvent

Use when debugging a specific event (e.g. why it failed or how it was aggregated). Includes processing status and debug info.

### Example Usage

<!-- UsageSnippet language="go" operationID="getEvent" method="get" path="/events/{id}" -->
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

    res, err := s.Events.GetEvent(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetEventByIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Event ID                                              |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetEventResponse](../../models/dtos/geteventresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 404                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |