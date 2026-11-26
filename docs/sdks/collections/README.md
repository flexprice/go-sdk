# Collections

## Overview

### Available Operations

* [GetCollectionsBoard](#getcollectionsboard) - Get collections board
* [GetCollectionsBoardColumn](#getcollectionsboardcolumn) - Get a page of a collections board column
* [ListReminderDrafts](#listreminderdrafts) - List reminder drafts
* [BulkUpdateReminderDrafts](#bulkupdatereminderdrafts) - Update reminders in bulk
* [UpdateReminderDraft](#updatereminderdraft) - Update reminder
* [GetReminderDraft](#getreminderdraft) - Get reminder draft

## GetCollectionsBoard

Use when you need every open invoice by where it stands: sent but not due, overdue (one card per invoice, split by reminder status), and paid in the last 7 days. Counts and per-currency totals cover whole columns and follow the q and currency filters; by default no cards are returned (limit 0): each column carries count and totals, and has_more is true when the column has cards to fetch with GET /collections/board/columns/{key}. With limit above 0 the first page of cards is returned too, and has_more says more cards follow. available_currencies lists the currencies of open invoices and of the invoices the Paid column can show, before q and currency apply.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCollectionsBoard" method="get" path="/collections/board" -->
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

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Collections.GetCollectionsBoard(ctx, dtos.GetCollectionsBoardRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionsBoardResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [dtos.GetCollectionsBoardRequest](../../models/dtos/getcollectionsboardrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.GetCollectionsBoardResponse](../../models/dtos/getcollectionsboardresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 403             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCollectionsBoardColumn

Use when you need the cards of one board column, one page at a time, after the board call. Pages are numbered from 1 and ordered stably (not due by due date, overdue and held by remaining balance, chasing by newest approval first, paid by newest payment first); send the as_of of the board response, and the same from and to, with every page so the overdue boundary, the date range and the 7-day window do not move while you scroll. A row that changes between page requests can repeat or be skipped, so reset pages when the board changes. Page times limit may not exceed 5000.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCollectionsBoardColumn" method="get" path="/collections/board/columns/{key}" -->
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

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Collections.GetCollectionsBoardColumn(ctx, dtos.GetCollectionsBoardColumnRequest{
        Key: "<key>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CollectionsColumnPageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [dtos.GetCollectionsBoardColumnRequest](../../models/dtos/getcollectionsboardcolumnrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][dtos.Option](../../models/dtos/option.md)                                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*dtos.GetCollectionsBoardColumnResponse](../../models/dtos/getcollectionsboardcolumnresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 403, 404        | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ListReminderDrafts

Use when you need the payment reminder of each overdue invoice with its tone, text and status. A reminder nobody acted on is rendered when read and stored nowhere; only edits and decisions are stored. The id of a reminder is its invoice id. The list is paged and nothing is sent. Order: status=approved lists the newest approval first (ties by invoice id); every other list puts the largest remaining balance first (ties by invoice id). Approvals made after as_of are left out of the list and the summary. from and to narrow the list and the summary by due date (invoices without a due date are left out when either is given); The summary (counts per status, totals per currency) follows from, to, currency and q like the board, and ignores status and customer_id so the status tabs count every status; available_currencies lists the currencies of open invoices and of invoices paid in the last 7 days, regardless of filters.

### Example Usage

<!-- UsageSnippet language="go" operationID="listReminderDrafts" method="get" path="/collections/reminders" -->
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

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Collections.ListReminderDrafts(ctx, dtos.ListReminderDraftsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListRemindersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [dtos.ListReminderDraftsRequest](../../models/dtos/listreminderdraftsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][dtos.Option](../../models/dtos/option.md)                                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*dtos.ListReminderDraftsResponse](../../models/dtos/listreminderdraftsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 403             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## BulkUpdateReminderDrafts

Use when you want to apply one status (and optionally tone) change to up to 200 invoices at once, with the same rules as updating one reminder. Subject and body are not allowed. All or nothing: if any invoice cannot take the change (unknown or not overdue, a change not legal for its status, someone changed it meanwhile, or, where emails are on, an approval for a customer without an email address), the request fails and no invoice changes. The error details name the invoice. Approving queues each email for the dispatcher.

### Example Usage

<!-- UsageSnippet language="go" operationID="bulkUpdateReminderDrafts" method="put" path="/collections/reminders/bulk-update" -->
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

    res, err := s.Collections.BulkUpdateReminderDrafts(ctx, types.BulkUpdateRemindersRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.BulkUpdateRemindersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [types.BulkUpdateRemindersRequest](../../models/types/bulkupdateremindersrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.BulkUpdateReminderDraftsResponse](../../models/dtos/bulkupdatereminderdraftsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 403, 404, 409   | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateReminderDraft

Use when you need to approve, hold, skip or reopen (status pending) a reminder, edit its tone, subject or body, or both at once. The request decides: status is pending, approved, held or skipped (omit it to keep the current one); status_reason is required for held and optional for skipped. A tone change rewrites subject and body from that tone's template, then a subject or body in the same request is applied on top. Allowed moves: pending to approved, held or skipped; held to pending, approved or skipped; approved or skipped to pending. Mail can only be edited while pending or held. Where emails are on for the environment, approving queues the email and fails with 400 when the customer has no email address (customer_email is empty and no contact person lists invoice_payment_reminder in email_types) or the tenant has no reply-to email; where emails are off, approving only records the decision.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateReminderDraft" method="put" path="/collections/reminders/update/{id}" -->
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

    res, err := s.Collections.UpdateReminderDraft(ctx, "<id>", types.UpdateReminderRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.InvoiceReminderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `id`                                                                       | `string`                                                                   | :heavy_check_mark:                                                         | Invoice ID                                                                 |
| `body`                                                                     | [types.UpdateReminderRequest](../../models/types/updatereminderrequest.md) | :heavy_check_mark:                                                         | Target status and/or mail changes                                          |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.UpdateReminderDraftResponse](../../models/dtos/updatereminderdraftresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 403, 404, 409   | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetReminderDraft

Use when you need the reminder of one overdue invoice. The id is the invoice id; 404 when the invoice is not open and overdue.

### Example Usage

<!-- UsageSnippet language="go" operationID="getReminderDraft" method="get" path="/collections/reminders/{id}" -->
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

    res, err := s.Collections.GetReminderDraft(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.InvoiceReminderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Invoice ID                                            |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetReminderDraftResponse](../../models/dtos/getreminderdraftresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 403, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |