# CreditNotes

## Overview

### Available Operations

* [CreateCreditNote](#createcreditnote) - Create credit note
* [GetCreditNote](#getcreditnote) - Get credit note
* [ProcessCreditNote](#processcreditnote) - Finalize credit note
* [VoidCreditNote](#voidcreditnote) - Void credit note

## CreateCreditNote

Use when issuing a refund or adjustment (e.g. customer dispute or proration). Links to an invoice; create as draft then finalize.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCreditNote" method="post" path="/creditnotes" -->
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

    res, err := s.CreditNotes.CreateCreditNote(ctx, types.DtoCreateCreditNoteRequest{
        InvoiceID: "<id>",
        Reason: types.CreditNoteReasonFraudulent,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [types.DtoCreateCreditNoteRequest](../../models/types/dtocreatecreditnoterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.CreateCreditNoteResponse](../../models/dtos/createcreditnoteresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 401, 403, 404         | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetCreditNote

Use when you need to load a single credit note (e.g. for display or reconciliation).

### Example Usage

<!-- UsageSnippet language="go" operationID="getCreditNote" method="get" path="/creditnotes/{id}" -->
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

    res, err := s.CreditNotes.GetCreditNote(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Credit note ID                                        |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetCreditNoteResponse](../../models/dtos/getcreditnoteresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## ProcessCreditNote

Use when locking a draft credit note and applying the credit (e.g. after approval). Once finalized, applied per billing provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="processCreditNote" method="post" path="/creditnotes/{id}/finalize" -->
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

    res, err := s.CreditNotes.ProcessCreditNote(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Credit note ID                                        |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.ProcessCreditNoteResponse](../../models/dtos/processcreditnoteresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 401, 403, 404         | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## VoidCreditNote

Use when cancelling a draft credit note (e.g. created by mistake). Only draft credit notes can be voided.

### Example Usage

<!-- UsageSnippet language="go" operationID="voidCreditNote" method="post" path="/creditnotes/{id}/void" -->
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

    res, err := s.CreditNotes.VoidCreditNote(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Credit note ID                                        |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.VoidCreditNoteResponse](../../models/dtos/voidcreditnoteresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 401, 403, 404         | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |