# WebhookEvents

## Overview

### Available Operations

* [PostWebhookEventsCreditNoteCreated](#postwebhookeventscreditnotecreated) - credit_note.created
* [PostWebhookEventsCreditNoteUpdated](#postwebhookeventscreditnoteupdated) - credit_note.updated
* [PostWebhookEventsCustomerCreated](#postwebhookeventscustomercreated) - customer.created
* [PostWebhookEventsCustomerDeleted](#postwebhookeventscustomerdeleted) - customer.deleted
* [PostWebhookEventsCustomerUpdated](#postwebhookeventscustomerupdated) - customer.updated
* [PostWebhookEventsEntitlementCreated](#postwebhookeventsentitlementcreated) - entitlement.created
* [PostWebhookEventsEntitlementDeleted](#postwebhookeventsentitlementdeleted) - entitlement.deleted
* [PostWebhookEventsEntitlementUpdated](#postwebhookeventsentitlementupdated) - entitlement.updated
* [PostWebhookEventsFeatureCreated](#postwebhookeventsfeaturecreated) - feature.created
* [PostWebhookEventsFeatureDeleted](#postwebhookeventsfeaturedeleted) - feature.deleted
* [PostWebhookEventsFeatureUpdated](#postwebhookeventsfeatureupdated) - feature.updated
* [PostWebhookEventsFeatureWalletBalanceAlert](#postwebhookeventsfeaturewalletbalancealert) - feature.wallet_balance.alert
* [PostWebhookEventsInvoiceCommunicationTriggered](#postwebhookeventsinvoicecommunicationtriggered) - invoice.communication.triggered
* [PostWebhookEventsInvoicePaymentOverdue](#postwebhookeventsinvoicepaymentoverdue) - invoice.payment.overdue
* [PostWebhookEventsInvoiceUpdate](#postwebhookeventsinvoiceupdate) - invoice.update
* [PostWebhookEventsInvoiceUpdateFinalized](#postwebhookeventsinvoiceupdatefinalized) - invoice.update.finalized
* [PostWebhookEventsInvoiceUpdatePayment](#postwebhookeventsinvoiceupdatepayment) - invoice.update.payment
* [PostWebhookEventsInvoiceUpdateVoided](#postwebhookeventsinvoiceupdatevoided) - invoice.update.voided
* [PostWebhookEventsPaymentCreated](#postwebhookeventspaymentcreated) - payment.created
* [PostWebhookEventsPaymentFailed](#postwebhookeventspaymentfailed) - payment.failed
* [PostWebhookEventsPaymentPending](#postwebhookeventspaymentpending) - payment.pending
* [PostWebhookEventsPaymentSuccess](#postwebhookeventspaymentsuccess) - payment.success
* [PostWebhookEventsPaymentUpdated](#postwebhookeventspaymentupdated) - payment.updated
* [PostWebhookEventsSubscriptionActivated](#postwebhookeventssubscriptionactivated) - subscription.activated
* [PostWebhookEventsSubscriptionCancelled](#postwebhookeventssubscriptioncancelled) - subscription.cancelled
* [PostWebhookEventsSubscriptionCreated](#postwebhookeventssubscriptioncreated) - subscription.created
* [PostWebhookEventsSubscriptionDraftCreated](#postwebhookeventssubscriptiondraftcreated) - subscription.draft.created
* [PostWebhookEventsSubscriptionPaused](#postwebhookeventssubscriptionpaused) - subscription.paused
* [PostWebhookEventsSubscriptionPhaseCreated](#postwebhookeventssubscriptionphasecreated) - subscription.phase.created
* [PostWebhookEventsSubscriptionPhaseDeleted](#postwebhookeventssubscriptionphasedeleted) - subscription.phase.deleted
* [PostWebhookEventsSubscriptionPhaseUpdated](#postwebhookeventssubscriptionphaseupdated) - subscription.phase.updated
* [PostWebhookEventsSubscriptionRenewalDue](#postwebhookeventssubscriptionrenewaldue) - subscription.renewal.due
* [PostWebhookEventsSubscriptionResumed](#postwebhookeventssubscriptionresumed) - subscription.resumed
* [PostWebhookEventsSubscriptionUpdated](#postwebhookeventssubscriptionupdated) - subscription.updated
* [PostWebhookEventsWalletCreated](#postwebhookeventswalletcreated) - wallet.created
* [PostWebhookEventsWalletCreditBalanceDropped](#postwebhookeventswalletcreditbalancedropped) - wallet.credit_balance.dropped
* [PostWebhookEventsWalletCreditBalanceRecovered](#postwebhookeventswalletcreditbalancerecovered) - wallet.credit_balance.recovered
* [PostWebhookEventsWalletOngoingBalanceDropped](#postwebhookeventswalletongoingbalancedropped) - wallet.ongoing_balance.dropped
* [PostWebhookEventsWalletOngoingBalanceRecovered](#postwebhookeventswalletongoingbalancerecovered) - wallet.ongoing_balance.recovered
* [PostWebhookEventsWalletTerminated](#postwebhookeventswalletterminated) - wallet.terminated
* [PostWebhookEventsWalletTransactionCreated](#postwebhookeventswallettransactioncreated) - wallet.transaction.created
* [PostWebhookEventsWalletUpdated](#postwebhookeventswalletupdated) - wallet.updated

## PostWebhookEventsCreditNoteCreated

Fired when a new credit note is created. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/credit_note.created" method="post" path="/webhook-events/credit_note.created" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsCreditNoteCreated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoCreditNoteWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsCreditNoteCreatedResponse](../../models/dtos/postwebhookeventscreditnotecreatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsCreditNoteUpdated

Fired when a credit note is updated. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/credit_note.updated" method="post" path="/webhook-events/credit_note.updated" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsCreditNoteUpdated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoCreditNoteWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsCreditNoteUpdatedResponse](../../models/dtos/postwebhookeventscreditnoteupdatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsCustomerCreated

Fired when a new customer is created. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/customer.created" method="post" path="/webhook-events/customer.created" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsCustomerCreated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoCustomerWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsCustomerCreatedResponse](../../models/dtos/postwebhookeventscustomercreatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsCustomerDeleted

Fired when a customer is deleted. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/customer.deleted" method="post" path="/webhook-events/customer.deleted" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsCustomerDeleted(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoCustomerWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsCustomerDeletedResponse](../../models/dtos/postwebhookeventscustomerdeletedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsCustomerUpdated

Fired when a customer is updated. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/customer.updated" method="post" path="/webhook-events/customer.updated" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsCustomerUpdated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoCustomerWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsCustomerUpdatedResponse](../../models/dtos/postwebhookeventscustomerupdatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsEntitlementCreated

Fired when a new entitlement is created. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/entitlement.created" method="post" path="/webhook-events/entitlement.created" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsEntitlementCreated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoEntitlementWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsEntitlementCreatedResponse](../../models/dtos/postwebhookeventsentitlementcreatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsEntitlementDeleted

Fired when an entitlement is deleted. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/entitlement.deleted" method="post" path="/webhook-events/entitlement.deleted" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsEntitlementDeleted(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoEntitlementWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsEntitlementDeletedResponse](../../models/dtos/postwebhookeventsentitlementdeletedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsEntitlementUpdated

Fired when an entitlement is updated. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/entitlement.updated" method="post" path="/webhook-events/entitlement.updated" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsEntitlementUpdated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoEntitlementWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsEntitlementUpdatedResponse](../../models/dtos/postwebhookeventsentitlementupdatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsFeatureCreated

Fired when a new feature is created. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/feature.created" method="post" path="/webhook-events/feature.created" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsFeatureCreated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoFeatureWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsFeatureCreatedResponse](../../models/dtos/postwebhookeventsfeaturecreatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsFeatureDeleted

Fired when a feature is deleted. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/feature.deleted" method="post" path="/webhook-events/feature.deleted" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsFeatureDeleted(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoFeatureWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsFeatureDeletedResponse](../../models/dtos/postwebhookeventsfeaturedeletedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsFeatureUpdated

Fired when a feature is updated. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/feature.updated" method="post" path="/webhook-events/feature.updated" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsFeatureUpdated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoFeatureWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsFeatureUpdatedResponse](../../models/dtos/postwebhookeventsfeatureupdatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsFeatureWalletBalanceAlert

Fired when a feature's associated wallet balance crosses an alert threshold. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/feature.wallet_balance.alert" method="post" path="/webhook-events/feature.wallet_balance.alert" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsFeatureWalletBalanceAlert(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoAlertWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsFeatureWalletBalanceAlertResponse](../../models/dtos/postwebhookeventsfeaturewalletbalancealertresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsInvoiceCommunicationTriggered

Fired when an invoice communication (e.g. email notification) is triggered. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/invoice.communication.triggered" method="post" path="/webhook-events/invoice.communication.triggered" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsInvoiceCommunicationTriggered(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoCommunicationWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsInvoiceCommunicationTriggeredResponse](../../models/dtos/postwebhookeventsinvoicecommunicationtriggeredresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsInvoicePaymentOverdue

Fired when an invoice payment is overdue past the due date. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/invoice.payment.overdue" method="post" path="/webhook-events/invoice.payment.overdue" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsInvoicePaymentOverdue(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoInvoiceWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsInvoicePaymentOverdueResponse](../../models/dtos/postwebhookeventsinvoicepaymentoverdueresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsInvoiceUpdate

Fired when an invoice is updated. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/invoice.update" method="post" path="/webhook-events/invoice.update" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsInvoiceUpdate(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoInvoiceWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsInvoiceUpdateResponse](../../models/dtos/postwebhookeventsinvoiceupdateresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsInvoiceUpdateFinalized

Fired when an invoice is finalized and locked for payment. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/invoice.update.finalized" method="post" path="/webhook-events/invoice.update.finalized" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsInvoiceUpdateFinalized(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoInvoiceWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsInvoiceUpdateFinalizedResponse](../../models/dtos/postwebhookeventsinvoiceupdatefinalizedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsInvoiceUpdatePayment

Fired when an invoice payment status changes. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/invoice.update.payment" method="post" path="/webhook-events/invoice.update.payment" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsInvoiceUpdatePayment(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoInvoiceWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsInvoiceUpdatePaymentResponse](../../models/dtos/postwebhookeventsinvoiceupdatepaymentresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsInvoiceUpdateVoided

Fired when an invoice is voided (e.g. order cancelled or duplicate). Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/invoice.update.voided" method="post" path="/webhook-events/invoice.update.voided" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsInvoiceUpdateVoided(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoInvoiceWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsInvoiceUpdateVoidedResponse](../../models/dtos/postwebhookeventsinvoiceupdatevoidedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsPaymentCreated

Fired when a new payment is created. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/payment.created" method="post" path="/webhook-events/payment.created" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsPaymentCreated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoPaymentWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsPaymentCreatedResponse](../../models/dtos/postwebhookeventspaymentcreatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsPaymentFailed

Fired when a payment fails. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/payment.failed" method="post" path="/webhook-events/payment.failed" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsPaymentFailed(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoPaymentWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsPaymentFailedResponse](../../models/dtos/postwebhookeventspaymentfailedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsPaymentPending

Fired when a payment is pending processing. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/payment.pending" method="post" path="/webhook-events/payment.pending" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsPaymentPending(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoPaymentWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsPaymentPendingResponse](../../models/dtos/postwebhookeventspaymentpendingresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsPaymentSuccess

Fired when a payment succeeds. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/payment.success" method="post" path="/webhook-events/payment.success" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsPaymentSuccess(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoPaymentWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsPaymentSuccessResponse](../../models/dtos/postwebhookeventspaymentsuccessresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsPaymentUpdated

Fired when a payment is updated. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/payment.updated" method="post" path="/webhook-events/payment.updated" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsPaymentUpdated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoPaymentWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsPaymentUpdatedResponse](../../models/dtos/postwebhookeventspaymentupdatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionActivated

Fired when a draft subscription is activated. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.activated" method="post" path="/webhook-events/subscription.activated" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionActivated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionActivatedResponse](../../models/dtos/postwebhookeventssubscriptionactivatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionCancelled

Fired when a subscription is cancelled (immediately or end-of-period). Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.cancelled" method="post" path="/webhook-events/subscription.cancelled" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionCancelled(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionCancelledResponse](../../models/dtos/postwebhookeventssubscriptioncancelledresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionCreated

Fired when a new subscription is created. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.created" method="post" path="/webhook-events/subscription.created" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionCreated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionCreatedResponse](../../models/dtos/postwebhookeventssubscriptioncreatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionDraftCreated

Fired when a new draft subscription is created (not yet activated). Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.draft.created" method="post" path="/webhook-events/subscription.draft.created" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionDraftCreated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionDraftCreatedResponse](../../models/dtos/postwebhookeventssubscriptiondraftcreatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionPaused

Fired when a subscription is paused. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.paused" method="post" path="/webhook-events/subscription.paused" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionPaused(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionPausedResponse](../../models/dtos/postwebhookeventssubscriptionpausedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionPhaseCreated

Fired when a new subscription phase is created. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.phase.created" method="post" path="/webhook-events/subscription.phase.created" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionPhaseCreated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionPhaseWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionPhaseCreatedResponse](../../models/dtos/postwebhookeventssubscriptionphasecreatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionPhaseDeleted

Fired when a subscription phase is deleted. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.phase.deleted" method="post" path="/webhook-events/subscription.phase.deleted" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionPhaseDeleted(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionPhaseWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionPhaseDeletedResponse](../../models/dtos/postwebhookeventssubscriptionphasedeletedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionPhaseUpdated

Fired when a subscription phase is updated. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.phase.updated" method="post" path="/webhook-events/subscription.phase.updated" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionPhaseUpdated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionPhaseWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionPhaseUpdatedResponse](../../models/dtos/postwebhookeventssubscriptionphaseupdatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionRenewalDue

Fired when a subscription renewal is upcoming (cron-driven). Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.renewal.due" method="post" path="/webhook-events/subscription.renewal.due" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionRenewalDue(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionRenewalDueResponse](../../models/dtos/postwebhookeventssubscriptionrenewaldueresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionResumed

Fired when a paused subscription is resumed. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.resumed" method="post" path="/webhook-events/subscription.resumed" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionResumed(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionResumedResponse](../../models/dtos/postwebhookeventssubscriptionresumedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsSubscriptionUpdated

Fired when a subscription is updated (e.g. quantity, billing anchor, or metadata changes). Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/subscription.updated" method="post" path="/webhook-events/subscription.updated" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsSubscriptionUpdated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoSubscriptionWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsSubscriptionUpdatedResponse](../../models/dtos/postwebhookeventssubscriptionupdatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsWalletCreated

Fired when a new wallet is created. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/wallet.created" method="post" path="/webhook-events/wallet.created" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsWalletCreated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoWalletWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsWalletCreatedResponse](../../models/dtos/postwebhookeventswalletcreatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsWalletCreditBalanceDropped

Fired when a wallet's credit balance drops below an alert threshold. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/wallet.credit_balance.dropped" method="post" path="/webhook-events/wallet.credit_balance.dropped" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsWalletCreditBalanceDropped(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoWalletWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsWalletCreditBalanceDroppedResponse](../../models/dtos/postwebhookeventswalletcreditbalancedroppedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsWalletCreditBalanceRecovered

Fired when a wallet's credit balance recovers above an alert threshold. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/wallet.credit_balance.recovered" method="post" path="/webhook-events/wallet.credit_balance.recovered" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsWalletCreditBalanceRecovered(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoWalletWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsWalletCreditBalanceRecoveredResponse](../../models/dtos/postwebhookeventswalletcreditbalancerecoveredresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsWalletOngoingBalanceDropped

Fired when a wallet's ongoing balance drops below an alert threshold. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/wallet.ongoing_balance.dropped" method="post" path="/webhook-events/wallet.ongoing_balance.dropped" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsWalletOngoingBalanceDropped(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoWalletWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsWalletOngoingBalanceDroppedResponse](../../models/dtos/postwebhookeventswalletongoingbalancedroppedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsWalletOngoingBalanceRecovered

Fired when a wallet's ongoing balance recovers above an alert threshold. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/wallet.ongoing_balance.recovered" method="post" path="/webhook-events/wallet.ongoing_balance.recovered" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsWalletOngoingBalanceRecovered(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoWalletWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsWalletOngoingBalanceRecoveredResponse](../../models/dtos/postwebhookeventswalletongoingbalancerecoveredresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsWalletTerminated

Fired when a wallet is terminated. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/wallet.terminated" method="post" path="/webhook-events/wallet.terminated" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsWalletTerminated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoWalletWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsWalletTerminatedResponse](../../models/dtos/postwebhookeventswalletterminatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsWalletTransactionCreated

Fired when a new wallet transaction is created (top-up, deduction, etc.). Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/wallet.transaction.created" method="post" path="/webhook-events/wallet.transaction.created" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsWalletTransactionCreated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoTransactionWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsWalletTransactionCreatedResponse](../../models/dtos/postwebhookeventswallettransactioncreatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## PostWebhookEventsWalletUpdated

Fired when a wallet is updated. Doc-only for parsing.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/webhook-events/wallet.updated" method="post" path="/webhook-events/wallet.updated" -->
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

    res, err := s.WebhookEvents.PostWebhookEventsWalletUpdated(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookDtoWalletWebhookPayload != nil {
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

**[*dtos.PostWebhookEventsWalletUpdatedResponse](../../models/dtos/postwebhookeventswalletupdatedresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |