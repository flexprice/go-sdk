# Commitments

## Overview

### Available Operations

* [CreateCommitment](#createcommitment) - Create commitment
* [PreviewCommitment](#previewcommitment) - Preview commitment
* [QueryCommitment](#querycommitment) - Query commitments
* [GetCommitment](#getcommitment) - Get commitment
* [UpdateCommitment](#updatecommitment) - Update commitment
* [GetCommitmentBillingTerms](#getcommitmentbillingterms) - Get billing terms
* [ReplaceCommitmentBillingTerms](#replacecommitmentbillingterms) - Edit billing terms
* [PreviewCommitmentBillingTerms](#previewcommitmentbillingterms) - Preview billing terms edit
* [ListCommitmentInstalments](#listcommitmentinstalments) - List instalments
* [TriggerCommitmentInstalment](#triggercommitmentinstalment) - Trigger instalment
* [TerminateCommitment](#terminatecommitment) - Terminate commitment

## CreateCommitment

Creates a commitment on a subscription: a minimum spend or quantity over a term, with optional committed and overage rates.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCommitment" method="post" path="/commitments" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Commitments.CreateCommitment(ctx, types.CreateCommitmentRequest{
        Measurement: types.CommitmentMeasurementBillingPeriod,
        Name: "<value>",
        Scope: types.CommitmentScopeInput{
            EntityType: types.CommitmentScopeTypeGroup,
        },
        StartDate: types.MustTimeFromString("2025-03-05T22:53:12.661Z"),
        Type: types.CommitmentTypeQuantity,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommitmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [types.CreateCommitmentRequest](../../models/types/createcommitmentrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][dtos.Option](../../models/dtos/option.md)                                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*dtos.CreateCommitmentResponse](../../models/dtos/createcommitmentresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## PreviewCommitment

Validates a commitment and returns its computed value without creating it.

### Example Usage

<!-- UsageSnippet language="go" operationID="previewCommitment" method="post" path="/commitments/preview" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"github.com/flexprice/go-sdk/v2/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Commitments.PreviewCommitment(ctx, types.CreateCommitmentRequest{
        Measurement: types.CommitmentMeasurementBillingPeriod,
        Name: "<value>",
        Scope: types.CommitmentScopeInput{
            EntityType: types.CommitmentScopeTypeMeter,
        },
        StartDate: types.MustTimeFromString("2024-11-23T11:41:25.905Z"),
        Type: types.CommitmentTypeAmount,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CommitmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [types.CreateCommitmentRequest](../../models/types/createcommitmentrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][dtos.Option](../../models/dtos/option.md)                                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*dtos.PreviewCommitmentResponse](../../models/dtos/previewcommitmentresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryCommitment

Lists commitments with filtering and sorting. Returns a paginated list.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryCommitment" method="post" path="/commitments/search" -->
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

    res, err := s.Commitments.QueryCommitment(ctx, types.CommitmentFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCommitmentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [types.CommitmentFilter](../../models/types/commitmentfilter.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][dtos.Option](../../models/dtos/option.md)                     | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*dtos.QueryCommitmentResponse](../../models/dtos/querycommitmentresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCommitment

Loads a single commitment by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommitment" method="get" path="/commitments/{id}" -->
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

    res, err := s.Commitments.GetCommitment(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CommitmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Commitment ID                                         |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetCommitmentResponse](../../models/dtos/getcommitmentresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateCommitment

Updates the descriptive fields of a commitment. Economic terms are immutable.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCommitment" method="put" path="/commitments/{id}" -->
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

    res, err := s.Commitments.UpdateCommitment(ctx, "<id>", types.UpdateCommitmentRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.CommitmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | `string`                                                                       | :heavy_check_mark:                                                             | Commitment ID                                                                  |
| `body`                                                                         | [types.UpdateCommitmentRequest](../../models/types/updatecommitmentrequest.md) | :heavy_check_mark:                                                             | Commitment update                                                              |
| `opts`                                                                         | [][dtos.Option](../../models/dtos/option.md)                                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*dtos.UpdateCommitmentResponse](../../models/dtos/updatecommitmentresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCommitmentBillingTerms

Returns the commitment's current billing terms with its scheduled and invoiced instalments in billing order. Invoiced instalments may come from terms an edit replaced. Cancelled instalments are listed by the instalments endpoint.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCommitmentBillingTerms" method="get" path="/commitments/{id}/billing-terms" -->
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

    res, err := s.Commitments.GetCommitmentBillingTerms(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingTermsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Commitment ID                                         |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetCommitmentBillingTermsResponse](../../models/dtos/getcommitmentbillingtermsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ReplaceCommitmentBillingTerms

Replaces the whole set of billing terms of an active commitment. Its scheduled instalments are cancelled and the new terms generate instalments for what the invoiced ones leave of the commitment total; invoiced instalments and their invoices are kept. A term without bill_on or start_date starts at the billing period holding now, or at the commitment start once its term has ended. Returns the new terms with the scheduled and invoiced instalments in billing order.

### Example Usage

<!-- UsageSnippet language="go" operationID="replaceCommitmentBillingTerms" method="put" path="/commitments/{id}/billing-terms" -->
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

    res, err := s.Commitments.ReplaceCommitmentBillingTerms(ctx, "<id>", types.ReplaceBillingTermsRequest{
        BillingTerms: []types.BillingTermInput{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingTermsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | `string`                                                                             | :heavy_check_mark:                                                                   | Commitment ID                                                                        |
| `body`                                                                               | [types.ReplaceBillingTermsRequest](../../models/types/replacebillingtermsrequest.md) | :heavy_check_mark:                                                                   | Billing terms                                                                        |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.ReplaceCommitmentBillingTermsResponse](../../models/dtos/replacecommitmentbillingtermsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## PreviewCommitmentBillingTerms

Returns what editing the billing terms would produce, without saving anything: the new terms with the invoiced and new instalments in billing order.

### Example Usage

<!-- UsageSnippet language="go" operationID="previewCommitmentBillingTerms" method="post" path="/commitments/{id}/billing-terms/preview" -->
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

    res, err := s.Commitments.PreviewCommitmentBillingTerms(ctx, "<id>", types.ReplaceBillingTermsRequest{
        BillingTerms: []types.BillingTermInput{
            types.BillingTermInput{
                Kind: types.BillingTermKindOnce,
                ValueType: types.BillingTermValueTypeFixed,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingTermsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | `string`                                                                             | :heavy_check_mark:                                                                   | Commitment ID                                                                        |
| `body`                                                                               | [types.ReplaceBillingTermsRequest](../../models/types/replacebillingtermsrequest.md) | :heavy_check_mark:                                                                   | Billing terms                                                                        |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.PreviewCommitmentBillingTermsResponse](../../models/dtos/previewcommitmentbillingtermsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ListCommitmentInstalments

Lists the instalments of the commitment in billing order. Without a status filter every instalment is listed, cancelled ones included.

### Example Usage

<!-- UsageSnippet language="go" operationID="listCommitmentInstalments" method="get" path="/commitments/{id}/instalments" -->
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

    res, err := s.Commitments.ListCommitmentInstalments(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListInstalmentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |
| `id`                                                                | `string`                                                            | :heavy_check_mark:                                                  | Commitment ID                                                       |
| `status`                                                            | []`string`                                                          | :heavy_minus_sign:                                                  | Only instalments in these statuses (scheduled, invoiced, cancelled) |
| `opts`                                                              | [][dtos.Option](../../models/dtos/option.md)                        | :heavy_minus_sign:                                                  | The options for this request.                                       |

### Response

**[*dtos.ListCommitmentInstalmentsResponse](../../models/dtos/listcommitmentinstalmentsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## TriggerCommitmentInstalment

Bills a manual instalment of an active commitment on bill_on, now by default, which must fall within the commitment term. The instalment goes on the first subscription invoice at or after that date, or on its own invoice once due when its instalment_invoicing, the term's or else the commitment's, is separate_invoice. Either way it waits while the subscription is trialing or paused. Until it is triggered, a manual instalment waits at the end of the term, and bill_on can only move earlier; edit the billing terms to bill it later.

### Example Usage

<!-- UsageSnippet language="go" operationID="triggerCommitmentInstalment" method="post" path="/commitments/{id}/instalments/{instalment_id}/trigger" -->
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

    res, err := s.Commitments.TriggerCommitmentInstalment(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingtermInstalment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `id`                                                                              | `string`                                                                          | :heavy_check_mark:                                                                | Commitment ID                                                                     |
| `instalmentID`                                                                    | `string`                                                                          | :heavy_check_mark:                                                                | Instalment ID                                                                     |
| `body`                                                                            | [*types.TriggerInstalmentRequest](../../models/types/triggerinstalmentrequest.md) | :heavy_minus_sign:                                                                | Trigger request                                                                   |
| `opts`                                                                            | [][dtos.Option](../../models/dtos/option.md)                                      | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*dtos.TriggerCommitmentInstalmentResponse](../../models/dtos/triggercommitmentinstalmentresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## TerminateCommitment

Ends a commitment immediately. effective_at defaults to now and cannot be in the future. Under billing terms, its scheduled instalments are cancelled, and those already due by effective_at are billed together as one settlement; later instalments are not billed and nothing is refunded. The settlement and a fixed cancellation fee, when given, are billed together on one one-off invoice.

### Example Usage

<!-- UsageSnippet language="go" operationID="terminateCommitment" method="post" path="/commitments/{id}/terminate" -->
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

    res, err := s.Commitments.TerminateCommitment(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CommitmentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `id`                                                                                  | `string`                                                                              | :heavy_check_mark:                                                                    | Commitment ID                                                                         |
| `body`                                                                                | [*types.TerminateCommitmentRequest](../../models/types/terminatecommitmentrequest.md) | :heavy_minus_sign:                                                                    | Termination request                                                                   |
| `opts`                                                                                | [][dtos.Option](../../models/dtos/option.md)                                          | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

### Response

**[*dtos.TerminateCommitmentResponse](../../models/dtos/terminatecommitmentresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |