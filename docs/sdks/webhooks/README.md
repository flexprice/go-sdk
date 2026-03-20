# Webhooks

## Overview

### Available Operations

* [HandleChargebeeWebhook](#handlechargebeewebhook) - Handle Chargebee webhook events
* [HandleHubspotWebhook](#handlehubspotwebhook) - Handle HubSpot webhook events
* [HandleMoyasarWebhook](#handlemoyasarwebhook) - Handle Moyasar webhook events
* [HandleNomodWebhook](#handlenomodwebhook) - Handle Nomod webhook events
* [HandlePaddleWebhook](#handlepaddlewebhook) - Handle Paddle webhook events
* [HandleQuickbooksWebhook](#handlequickbookswebhook) - Handle QuickBooks webhook events
* [HandleRazorpayWebhook](#handlerazorpaywebhook) - Handle Razorpay webhook events
* [HandleStripeWebhook](#handlestripewebhook) - Handle Stripe webhook events

## HandleChargebeeWebhook

Use as the Chargebee webhook endpoint URL. Receives payment and subscription events from Chargebee to sync status into FlexPrice.

### Example Usage

<!-- UsageSnippet language="go" operationID="handleChargebeeWebhook" method="post" path="/webhooks/chargebee/{tenant_id}/{environment_id}" -->
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

    res, err := s.Webhooks.HandleChargebeeWebhook(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `tenantID`                                            | `string`                                              | :heavy_check_mark:                                    | Tenant ID                                             |
| `environmentID`                                       | `string`                                              | :heavy_check_mark:                                    | Environment ID                                        |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.HandleChargebeeWebhookResponse](../../models/dtos/handlechargebeewebhookresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## HandleHubspotWebhook

Use as the HubSpot webhook endpoint URL. Receives deal and customer events (e.g. deal closed won) to create or update customers in FlexPrice.

### Example Usage

<!-- UsageSnippet language="go" operationID="handleHubspotWebhook" method="post" path="/webhooks/hubspot/{tenant_id}/{environment_id}" -->
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

    res, err := s.Webhooks.HandleHubspotWebhook(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `tenantID`                                            | `string`                                              | :heavy_check_mark:                                    | Tenant ID                                             |
| `environmentID`                                       | `string`                                              | :heavy_check_mark:                                    | Environment ID                                        |
| `xHubSpotSignatureV3`                                 | `string`                                              | :heavy_check_mark:                                    | HubSpot webhook signature                             |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.HandleHubspotWebhookResponse](../../models/dtos/handlehubspotwebhookresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## HandleMoyasarWebhook

Use as the Moyasar webhook endpoint URL. Receives payment events from Moyasar to update payment status in FlexPrice.

### Example Usage

<!-- UsageSnippet language="go" operationID="handleMoyasarWebhook" method="post" path="/webhooks/moyasar/{tenant_id}/{environment_id}" -->
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

    res, err := s.Webhooks.HandleMoyasarWebhook(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `tenantID`                                            | `string`                                              | :heavy_check_mark:                                    | Tenant ID                                             |
| `environmentID`                                       | `string`                                              | :heavy_check_mark:                                    | Environment ID                                        |
| `xMoyasarSignature`                                   | `*string`                                             | :heavy_minus_sign:                                    | Moyasar webhook signature                             |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.HandleMoyasarWebhookResponse](../../models/dtos/handlemoyasarwebhookresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## HandleNomodWebhook

Use as the Nomod webhook endpoint URL. Receives payment and invoice events from Nomod to keep FlexPrice in sync.

### Example Usage

<!-- UsageSnippet language="go" operationID="handleNomodWebhook" method="post" path="/webhooks/nomod/{tenant_id}/{environment_id}" -->
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

    res, err := s.Webhooks.HandleNomodWebhook(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `tenantID`                                            | `string`                                              | :heavy_check_mark:                                    | Tenant ID                                             |
| `environmentID`                                       | `string`                                              | :heavy_check_mark:                                    | Environment ID                                        |
| `xAPIKey`                                             | `*string`                                             | :heavy_minus_sign:                                    | Nomod webhook secret (if configured)                  |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.HandleNomodWebhookResponse](../../models/dtos/handlenomodwebhookresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## HandlePaddleWebhook

Use as the Paddle webhook endpoint URL. Receives transaction.completed events from Paddle to update payment status in FlexPrice.

### Example Usage

<!-- UsageSnippet language="go" operationID="handlePaddleWebhook" method="post" path="/webhooks/paddle/{tenant_id}/{environment_id}" -->
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

    res, err := s.Webhooks.HandlePaddleWebhook(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `tenantID`                                            | `string`                                              | :heavy_check_mark:                                    | Tenant ID                                             |
| `environmentID`                                       | `string`                                              | :heavy_check_mark:                                    | Environment ID                                        |
| `paddleSignature`                                     | `string`                                              | :heavy_check_mark:                                    | Paddle webhook signature                              |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.HandlePaddleWebhookResponse](../../models/dtos/handlepaddlewebhookresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## HandleQuickbooksWebhook

Use as the QuickBooks webhook endpoint URL. Receives payment events from QuickBooks to sync payment status into FlexPrice.

### Example Usage

<!-- UsageSnippet language="go" operationID="handleQuickbooksWebhook" method="post" path="/webhooks/quickbooks/{tenant_id}/{environment_id}" -->
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

    res, err := s.Webhooks.HandleQuickbooksWebhook(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `tenantID`                                            | `string`                                              | :heavy_check_mark:                                    | Tenant ID                                             |
| `environmentID`                                       | `string`                                              | :heavy_check_mark:                                    | Environment ID                                        |
| `intuitSignature`                                     | `*string`                                             | :heavy_minus_sign:                                    | QuickBooks webhook signature                          |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.HandleQuickbooksWebhookResponse](../../models/dtos/handlequickbookswebhookresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## HandleRazorpayWebhook

Use as the Razorpay webhook endpoint URL. Receives payment capture and failure events to update invoice or payment status in FlexPrice.

### Example Usage

<!-- UsageSnippet language="go" operationID="handleRazorpayWebhook" method="post" path="/webhooks/razorpay/{tenant_id}/{environment_id}" -->
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

    res, err := s.Webhooks.HandleRazorpayWebhook(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `tenantID`                                            | `string`                                              | :heavy_check_mark:                                    | Tenant ID                                             |
| `environmentID`                                       | `string`                                              | :heavy_check_mark:                                    | Environment ID                                        |
| `xRazorpaySignature`                                  | `string`                                              | :heavy_check_mark:                                    | Razorpay webhook signature                            |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.HandleRazorpayWebhookResponse](../../models/dtos/handlerazorpaywebhookresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## HandleStripeWebhook

Use as the Stripe webhook endpoint URL. Receives payment and customer events from Stripe to keep FlexPrice in sync (e.g. payment succeeded, customer created).

### Example Usage

<!-- UsageSnippet language="go" operationID="handleStripeWebhook" method="post" path="/webhooks/stripe/{tenant_id}/{environment_id}" -->
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

    res, err := s.Webhooks.HandleStripeWebhook(ctx, "<id>", "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `tenantID`                                            | `string`                                              | :heavy_check_mark:                                    | Tenant ID                                             |
| `environmentID`                                       | `string`                                              | :heavy_check_mark:                                    | Environment ID                                        |
| `stripeSignature`                                     | `string`                                              | :heavy_check_mark:                                    | Stripe webhook signature                              |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.HandleStripeWebhookResponse](../../models/dtos/handlestripewebhookresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |