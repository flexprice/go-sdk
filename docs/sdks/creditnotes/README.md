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
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.CreditNotes.CreateCreditNote(ctx, types.CreateCreditNoteRequest{
        InvoiceID: "<id>",
        Reason: types.CreditNoteReasonFraudulent,
    }, dtos.CreateCreditNoteSecurity{
        Option1: &dtos.CreateCreditNoteSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [types.CreateCreditNoteRequest](../../models/types/createcreditnoterequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [dtos.CreateCreditNoteSecurity](../../models/dtos/createcreditnotesecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |
| `opts`                                                                         | [][dtos.Option](../../models/dtos/option.md)                                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*dtos.CreateCreditNoteResponse](../../models/dtos/createcreditnoteresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 401, 403, 404   | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCreditNote

Use when you need to load a single credit note (e.g. for display or reconciliation).

### Example Usage

<!-- UsageSnippet language="go" operationID="getCreditNote" method="get" path="/creditnotes/{id}" -->
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

    res, err := s.CreditNotes.GetCreditNote(ctx, dtos.GetCreditNoteSecurity{
        Option1: &dtos.GetCreditNoteSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `security`                                                               | [dtos.GetCreditNoteSecurity](../../models/dtos/getcreditnotesecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |
| `id`                                                                     | `string`                                                                 | :heavy_check_mark:                                                       | Credit note ID                                                           |
| `opts`                                                                   | [][dtos.Option](../../models/dtos/option.md)                             | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*dtos.GetCreditNoteResponse](../../models/dtos/getcreditnoteresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ProcessCreditNote

Use when locking a draft credit note and applying the credit (e.g. after approval). Once finalized, applied per billing provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="processCreditNote" method="post" path="/creditnotes/{id}/finalize" -->
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

    res, err := s.CreditNotes.ProcessCreditNote(ctx, dtos.ProcessCreditNoteSecurity{
        Option1: &dtos.ProcessCreditNoteSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `security`                                                                          | [dtos.ProcessCreditNoteSecurity](../../models/dtos/processcreditnotesecurity.md)    | :heavy_check_mark:                                                                  | The security requirements to use for the request.                                   |
| `id`                                                                                | `string`                                                                            | :heavy_check_mark:                                                                  | Credit note ID                                                                      |
| `body`                                                                              | [*types.FinalizeCreditNoteRequest](../../models/types/finalizecreditnoterequest.md) | :heavy_minus_sign:                                                                  | Finalize options                                                                    |
| `opts`                                                                              | [][dtos.Option](../../models/dtos/option.md)                                        | :heavy_minus_sign:                                                                  | The options for this request.                                                       |

### Response

**[*dtos.ProcessCreditNoteResponse](../../models/dtos/processcreditnoteresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 401, 403, 404   | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## VoidCreditNote

Use when cancelling a draft credit note (e.g. created by mistake). Only draft credit notes can be voided.

### Example Usage

<!-- UsageSnippet language="go" operationID="voidCreditNote" method="post" path="/creditnotes/{id}/void" -->
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

    res, err := s.CreditNotes.VoidCreditNote(ctx, dtos.VoidCreditNoteSecurity{
        Option1: &dtos.VoidCreditNoteSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `security`                                                                 | [dtos.VoidCreditNoteSecurity](../../models/dtos/voidcreditnotesecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |
| `id`                                                                       | `string`                                                                   | :heavy_check_mark:                                                         | Credit note ID                                                             |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.VoidCreditNoteResponse](../../models/dtos/voidcreditnoteresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 401, 403, 404   | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |