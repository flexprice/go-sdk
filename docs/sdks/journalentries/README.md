# JournalEntries

## Overview

### Available Operations

* [QueryJournalEntry](#queryjournalentry) - Query journal entries
* [GetJournalEntry](#getjournalentry) - Get journal entry

## QueryJournalEntry

Use when listing journal entries by period, event type, account, customer, contract, performance obligation or source. Returns a paginated list with lines.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryJournalEntry" method="post" path="/journal_entries/search" -->
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

    res, err := s.JournalEntries.QueryJournalEntry(ctx, types.JournalFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListJournalEntriesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [types.JournalFilter](../../models/types/journalfilter.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |
| `opts`                                                     | [][dtos.Option](../../models/dtos/option.md)               | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*dtos.QueryJournalEntryResponse](../../models/dtos/queryjournalentryresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetJournalEntry

Use when you need one revenue recognition journal entry with its debit and credit lines. The journal is derived from billing and recognition and is read-only.

### Example Usage

<!-- UsageSnippet language="go" operationID="getJournalEntry" method="get" path="/journal_entries/{id}" -->
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

    res, err := s.JournalEntries.GetJournalEntry(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.JournalEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Journal entry ID                                      |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetJournalEntryResponse](../../models/dtos/getjournalentryresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |