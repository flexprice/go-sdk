# Invoices

## Overview

### Available Operations

* [GetCustomerInvoiceSummary](#getcustomerinvoicesummary) - Get customer invoice summary
* [CreateInvoice](#createinvoice) - Create one-off invoice
* [GetInvoicePreview](#getinvoicepreview) - Get invoice preview
* [QueryInvoice](#queryinvoice) - Query invoices
* [GetInvoice](#getinvoice) - Get invoice
* [UpdateInvoice](#updateinvoice) - Update invoice
* [TriggerInvoiceCommsWebhook](#triggerinvoicecommswebhook) - Trigger invoice communication webhook
* [FinalizeInvoice](#finalizeinvoice) - Finalize invoice
* [UpdateInvoicePaymentStatus](#updateinvoicepaymentstatus) - Update invoice payment status
* [AttemptInvoicePayment](#attemptinvoicepayment) - Attempt invoice payment
* [GetInvoicePdf](#getinvoicepdf) - Get invoice PDF
* [RecalculateInvoice](#recalculateinvoice) - Recalculate invoice (voided invoice)
* [RecalculateInvoiceV2](#recalculateinvoicev2) - Recalculate draft invoice (v2)
* [VoidInvoice](#voidinvoice) - Void invoice

## GetCustomerInvoiceSummary

Use when showing a customer's invoice overview (e.g. billing portal or balance summary). Includes totals and multi-currency breakdown.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomerInvoiceSummary" method="get" path="/customers/{id}/invoices/summary" -->
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

    res, err := s.Invoices.GetCustomerInvoiceSummary(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerMultiCurrencyInvoiceSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Customer ID                                           |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetCustomerInvoiceSummaryResponse](../../models/dtos/getcustomerinvoicesummaryresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CreateInvoice

Use when creating a manual or one-off invoice (e.g. custom charge or non-recurring billing). Invoice is created in draft; finalize when ready.

### Example Usage

<!-- UsageSnippet language="go" operationID="createInvoice" method="post" path="/invoices" -->
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

    res, err := s.Invoices.CreateInvoice(ctx, types.CreateInvoiceRequest{
        AmountDue: "<value>",
        Currency: "Surinam Dollar",
        CustomerID: "<id>",
        Subtotal: "<value>",
        Total: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Invoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [types.CreateInvoiceRequest](../../models/types/createinvoicerequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][dtos.Option](../../models/dtos/option.md)                             | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*dtos.CreateInvoiceResponse](../../models/dtos/createinvoiceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetInvoicePreview

Use when showing a customer what they will be charged (e.g. preview before checkout or plan change). No invoice is created.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInvoicePreview" method="post" path="/invoices/preview" -->
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

    res, err := s.Invoices.GetInvoicePreview(ctx, types.GetPreviewInvoiceRequest{
        SubscriptionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Invoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [types.GetPreviewInvoiceRequest](../../models/types/getpreviewinvoicerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][dtos.Option](../../models/dtos/option.md)                                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*dtos.GetInvoicePreviewResponse](../../models/dtos/getinvoicepreviewresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryInvoice

Use when listing or searching invoices (e.g. admin view or customer history). Returns a paginated list; supports filtering by customer, status, date range.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryInvoice" method="post" path="/invoices/search" -->
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

    res, err := s.Invoices.QueryInvoice(ctx, types.InvoiceFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListInvoicesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [types.InvoiceFilter](../../models/types/invoicefilter.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |
| `opts`                                                     | [][dtos.Option](../../models/dtos/option.md)               | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*dtos.QueryInvoiceResponse](../../models/dtos/queryinvoiceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetInvoice

Use when loading an invoice for display or editing (e.g. portal or reconciliation). Supports group_by for usage breakdown and force_runtime_recalculation.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInvoice" method="get" path="/invoices/{id}" -->
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

    res, err := s.Invoices.GetInvoice(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Invoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `id`                                                                                    | `string`                                                                                | :heavy_check_mark:                                                                      | Invoice ID                                                                              |
| `expandBySource`                                                                        | `*bool`                                                                                 | :heavy_minus_sign:                                                                      | Include source-level price breakdown for usage line items (legacy)                      |
| `groupBy`                                                                               | []`string`                                                                              | :heavy_minus_sign:                                                                      | Group usage breakdown by specified fields (e.g., source, feature_id, properties.org_id) |
| `opts`                                                                                  | [][dtos.Option](../../models/dtos/option.md)                                            | :heavy_minus_sign:                                                                      | The options for this request.                                                           |

### Response

**[*dtos.GetInvoiceResponse](../../models/dtos/getinvoiceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateInvoice

Use when updating invoice metadata or due date (e.g. PDF URL, net terms). For paid invoices only safe fields can be updated.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateInvoice" method="put" path="/invoices/{id}" -->
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

    res, err := s.Invoices.UpdateInvoice(ctx, "<id>", types.UpdateInvoiceRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Invoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `id`                                                                     | `string`                                                                 | :heavy_check_mark:                                                       | Invoice ID                                                               |
| `body`                                                                   | [types.UpdateInvoiceRequest](../../models/types/updateinvoicerequest.md) | :heavy_check_mark:                                                       | Invoice Update Request                                                   |
| `opts`                                                                   | [][dtos.Option](../../models/dtos/option.md)                             | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*dtos.UpdateInvoiceResponse](../../models/dtos/updateinvoiceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## TriggerInvoiceCommsWebhook

Use when sending an invoice to the customer (e.g. trigger email or Slack). Payload includes full invoice details for your integration.

### Example Usage

<!-- UsageSnippet language="go" operationID="triggerInvoiceCommsWebhook" method="post" path="/invoices/{id}/comms/trigger" -->
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

    res, err := s.Invoices.TriggerInvoiceCommsWebhook(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
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

**[*dtos.TriggerInvoiceCommsWebhookResponse](../../models/dtos/triggerinvoicecommswebhookresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## FinalizeInvoice

Use when locking an invoice for payment (e.g. after review). Once finalized, line items are locked; invoice can be paid or voided.

### Example Usage

<!-- UsageSnippet language="go" operationID="finalizeInvoice" method="post" path="/invoices/{id}/finalize" -->
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

    res, err := s.Invoices.FinalizeInvoice(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
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

**[*dtos.FinalizeInvoiceResponse](../../models/dtos/finalizeinvoiceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateInvoicePaymentStatus

Use when reconciling payment status from an external gateway or manual entry (e.g. mark paid after bank confirmation).

### Example Usage

<!-- UsageSnippet language="go" operationID="updateInvoicePaymentStatus" method="put" path="/invoices/{id}/payment" -->
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

    res, err := s.Invoices.UpdateInvoicePaymentStatus(ctx, "<id>", types.UpdatePaymentStatusRequest{
        PaymentStatus: types.PaymentStatusInitiated,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Invoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | `string`                                                                             | :heavy_check_mark:                                                                   | Invoice ID                                                                           |
| `body`                                                                               | [types.UpdatePaymentStatusRequest](../../models/types/updatepaymentstatusrequest.md) | :heavy_check_mark:                                                                   | Payment Status Update Request                                                        |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.UpdateInvoicePaymentStatusResponse](../../models/dtos/updateinvoicepaymentstatusresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## AttemptInvoicePayment

Use when paying an invoice with the customer's wallet balance (e.g. prepaid credits or balance applied at checkout).

### Example Usage

<!-- UsageSnippet language="go" operationID="attemptInvoicePayment" method="post" path="/invoices/{id}/payment/attempt" -->
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

    res, err := s.Invoices.AttemptInvoicePayment(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
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

**[*dtos.AttemptInvoicePaymentResponse](../../models/dtos/attemptinvoicepaymentresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetInvoicePdf

Use when delivering an invoice PDF to the customer (e.g. email attachment or download). Use url=true for a presigned URL instead of binary. Use force_generate=true to regenerate and re-upload the PDF even if one already exists in S3.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInvoicePdf" method="get" path="/invoices/{id}/pdf" -->
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

    res, err := s.Invoices.GetInvoicePdf(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseStream != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                             | Type                                                                                                                                                                  | Required                                                                                                                                                              | Description                                                                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                 | :heavy_check_mark:                                                                                                                                                    | The context to use for the request.                                                                                                                                   |
| `id`                                                                                                                                                                  | `string`                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                    | Invoice ID                                                                                                                                                            |
| `url_`                                                                                                                                                                | `*bool`                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                    | Return presigned URL from s3 instead of PDF                                                                                                                           |
| `forceGenerate`                                                                                                                                                       | `*bool`                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                    | Force regeneration of the PDF even if one already exists in S3 (default: false). Note: force_generate has no effect if invoice_pdf_url is already set on the invoice. |
| `opts`                                                                                                                                                                | [][dtos.Option](../../models/dtos/option.md)                                                                                                                          | :heavy_minus_sign:                                                                                                                                                    | The options for this request.                                                                                                                                         |

### Response

**[*dtos.GetInvoicePdfResponse](../../models/dtos/getinvoicepdfresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## RecalculateInvoice

Starts an async workflow that creates a fresh replacement invoice for a voided SUBSCRIPTION invoice (same billing period). Returns workflow_id and run_id; poll workflow status or GET the new invoice via recalculated_invoice_id after completion.

### Example Usage

<!-- UsageSnippet language="go" operationID="recalculateInvoice" method="post" path="/invoices/{id}/recalculate" -->
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

    res, err := s.Invoices.RecalculateInvoice(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ModelsTemporalWorkflowResult != nil {
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

**[*dtos.RecalculateInvoiceResponse](../../models/dtos/recalculateinvoiceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## RecalculateInvoiceV2

Recalculates a draft SUBSCRIPTION invoice in-place (replaces line items, reapplies credits/coupons/taxes). Use when subscription or usage data changed before finalizing.

### Example Usage

<!-- UsageSnippet language="go" operationID="recalculateInvoiceV2" method="post" path="/invoices/{id}/recalculate-v2" -->
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

    res, err := s.Invoices.RecalculateInvoiceV2(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Invoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |
| `id`                                                                | `string`                                                            | :heavy_check_mark:                                                  | Invoice ID                                                          |
| `finalize`                                                          | `*bool`                                                             | :heavy_minus_sign:                                                  | Whether to finalize the invoice after recalculation (default: true) |
| `opts`                                                              | [][dtos.Option](../../models/dtos/option.md)                        | :heavy_minus_sign:                                                  | The options for this request.                                       |

### Response

**[*dtos.RecalculateInvoiceV2Response](../../models/dtos/recalculateinvoicev2response.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## VoidInvoice

Use when cancelling an invoice (e.g. order cancelled or duplicate). Only unpaid invoices can be voided.

### Example Usage

<!-- UsageSnippet language="go" operationID="voidInvoice" method="post" path="/invoices/{id}/void" -->
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

    res, err := s.Invoices.VoidInvoice(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
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

**[*dtos.VoidInvoiceResponse](../../models/dtos/voidinvoiceresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |