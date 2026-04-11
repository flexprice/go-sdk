# Subscriptions

## Overview

### Available Operations

* [CreateSubscription](#createsubscription) - Create subscription
* [AddSubscriptionAddon](#addsubscriptionaddon) - Add addon to subscription
* [RemoveSubscriptionAddon](#removesubscriptionaddon) - Remove addon from subscription
* [UpdateSubscriptionLineItem](#updatesubscriptionlineitem) - Update subscription line item
* [DeleteSubscriptionLineItem](#deletesubscriptionlineitem) - Delete subscription line item
* [QuerySubscription](#querysubscription) - Query subscriptions
* [GetSubscriptionUsage](#getsubscriptionusage) - Get usage by subscription
* [GetSubscription](#getsubscription) - Get subscription
* [UpdateSubscription](#updatesubscription) - Update subscription
* [ActivateSubscription](#activatesubscription) - Activate draft subscription
* [GetSubscriptionAddonAssociations](#getsubscriptionaddonassociations) - Get active addon associations
* [CancelSubscription](#cancelsubscription) - Cancel subscription
* [ExecuteSubscriptionChange](#executesubscriptionchange) - Execute subscription plan change
* [PreviewSubscriptionChange](#previewsubscriptionchange) - Preview subscription plan change
* [GetSubscriptionEntitlements](#getsubscriptionentitlements) - Get subscription entitlements
* [GetSubscriptionUpcomingGrants](#getsubscriptionupcominggrants) - Get upcoming credit grant applications
* [CreateSubscriptionLineItem](#createsubscriptionlineitem) - Create subscription line item
* [ExecuteSubscriptionModify](#executesubscriptionmodify) - Execute subscription modification
* [PreviewSubscriptionModify](#previewsubscriptionmodify) - Preview subscription modification
* [PauseSubscription](#pausesubscription) - Pause a subscription
* [ListSubscriptionPauses](#listsubscriptionpauses) - List all pauses for a subscription
* [ResumeSubscription](#resumesubscription) - Resume a paused subscription
* [GetSubscriptionV2](#getsubscriptionv2) - Get subscription (V2)
* [ListAllSubscriptionSchedules](#listallsubscriptionschedules) - List all subscription schedules
* [GetSubscriptionSchedule](#getsubscriptionschedule) - Get subscription schedule
* [CancelSubscriptionSchedule](#cancelsubscriptionschedule) - Cancel subscription schedule
* [ListSubscriptionSchedules](#listsubscriptionschedules) - List subscription schedules

## CreateSubscription

Use when onboarding a customer to a plan or starting a new subscription. Ideal for draft subscriptions (activate later) or active from start.

### Example Usage

<!-- UsageSnippet language="go" operationID="createSubscription" method="post" path="/subscriptions" -->
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

    res, err := s.Subscriptions.CreateSubscription(ctx, types.CreateSubscriptionRequest{
        BillingPeriod: types.BillingPeriodOnetime,
        Currency: "Kwacha",
        PlanID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [types.CreateSubscriptionRequest](../../models/types/createsubscriptionrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.CreateSubscriptionResponse](../../models/dtos/createsubscriptionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## AddSubscriptionAddon

Use when adding an optional product or add-on to an existing subscription (e.g. extra storage or support tier).

### Example Usage

<!-- UsageSnippet language="go" operationID="addSubscriptionAddon" method="post" path="/subscriptions/addon" -->
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

    res, err := s.Subscriptions.AddSubscriptionAddon(ctx, types.AddAddonRequest{
        AddonID: "<id>",
        SubscriptionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AddonAssociationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [types.AddAddonRequest](../../models/types/addaddonrequest.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][dtos.Option](../../models/dtos/option.md)                   | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*dtos.AddSubscriptionAddonResponse](../../models/dtos/addsubscriptionaddonresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## RemoveSubscriptionAddon

Use when removing an add-on from a subscription (e.g. downgrade or opt-out).

### Example Usage

<!-- UsageSnippet language="go" operationID="removeSubscriptionAddon" method="delete" path="/subscriptions/addon" -->
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

    res, err := s.Subscriptions.RemoveSubscriptionAddon(ctx, types.RemoveAddonRequest{
        AddonAssociationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [types.RemoveAddonRequest](../../models/types/removeaddonrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.RemoveSubscriptionAddonResponse](../../models/dtos/removesubscriptionaddonresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateSubscriptionLineItem

Use when changing a subscription line item (e.g. quantity or price). Implemented by ending the current line and creating a new one for clean billing.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSubscriptionLineItem" method="put" path="/subscriptions/lineitems/{id}" -->
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

    res, err := s.Subscriptions.UpdateSubscriptionLineItem(ctx, "<id>", types.UpdateSubscriptionLineItemRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionLineItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | `string`                                                                                           | :heavy_check_mark:                                                                                 | Line Item ID                                                                                       |
| `body`                                                                                             | [types.UpdateSubscriptionLineItemRequest](../../models/types/updatesubscriptionlineitemrequest.md) | :heavy_check_mark:                                                                                 | Update Line Item Request                                                                           |
| `opts`                                                                                             | [][dtos.Option](../../models/dtos/option.md)                                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*dtos.UpdateSubscriptionLineItemResponse](../../models/dtos/updatesubscriptionlineitemresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeleteSubscriptionLineItem

Use when removing a charge or seat from a subscription (e.g. downgrade). Line item ends; retained for history but no longer billed.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSubscriptionLineItem" method="delete" path="/subscriptions/lineitems/{id}" -->
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

    res, err := s.Subscriptions.DeleteSubscriptionLineItem(ctx, "<id>", types.DeleteSubscriptionLineItemRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionLineItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | `string`                                                                                           | :heavy_check_mark:                                                                                 | Line Item ID                                                                                       |
| `body`                                                                                             | [types.DeleteSubscriptionLineItemRequest](../../models/types/deletesubscriptionlineitemrequest.md) | :heavy_check_mark:                                                                                 | Delete Line Item Request                                                                           |
| `opts`                                                                                             | [][dtos.Option](../../models/dtos/option.md)                                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*dtos.DeleteSubscriptionLineItemResponse](../../models/dtos/deletesubscriptionlineitemresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QuerySubscription

Use when listing or searching subscriptions (e.g. admin view or customer subscription list). Returns a paginated list; supports filtering by customer, plan, status.

### Example Usage

<!-- UsageSnippet language="go" operationID="querySubscription" method="post" path="/subscriptions/search" -->
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

    res, err := s.Subscriptions.QuerySubscription(ctx, types.SubscriptionFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSubscriptionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [types.SubscriptionFilter](../../models/types/subscriptionfilter.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.QuerySubscriptionResponse](../../models/dtos/querysubscriptionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetSubscriptionUsage

Use when showing usage for a subscription (e.g. in a portal or for overage checks). Supports time range and filters.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionUsage" method="post" path="/subscriptions/usage" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/types"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.GetSubscriptionUsage(ctx, types.GetUsageBySubscriptionRequest{
        EndTime: types.MustNewTimeFromString("2024-03-20T00:00:00Z"),
        LifetimeUsage: flexprice.Pointer(false),
        StartTime: types.MustNewTimeFromString("2024-03-13T00:00:00Z"),
        SubscriptionID: "123",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetUsageBySubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [types.GetUsageBySubscriptionRequest](../../models/types/getusagebysubscriptionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][dtos.Option](../../models/dtos/option.md)                                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*dtos.GetSubscriptionUsageResponse](../../models/dtos/getsubscriptionusageresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetSubscription

Use when you need to load a single subscription (e.g. for a billing portal or to check status).

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscription" method="get" path="/subscriptions/{id}" -->
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

    res, err := s.Subscriptions.GetSubscription(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetSubscriptionResponse](../../models/dtos/getsubscriptionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateSubscription

Use when changing subscription details (e.g. quantity, billing anchor, or parent). Supports partial update; send "" to clear parent_subscription_id.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSubscription" method="put" path="/subscriptions/{id}" -->
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

    res, err := s.Subscriptions.UpdateSubscription(ctx, "<id>", types.UpdateSubscriptionRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | Subscription ID                                                                    |
| `body`                                                                             | [types.UpdateSubscriptionRequest](../../models/types/updatesubscriptionrequest.md) | :heavy_check_mark:                                                                 | Update Subscription Request                                                        |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.UpdateSubscriptionResponse](../../models/dtos/updatesubscriptionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ActivateSubscription

Use when turning a draft subscription live (e.g. after collecting payment or completing setup). Once activated, billing and entitlements apply.

### Example Usage

<!-- UsageSnippet language="go" operationID="activateSubscription" method="post" path="/subscriptions/{id}/activate" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/types"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.ActivateSubscription(ctx, "<id>", types.ActivateDraftSubscriptionRequest{
        StartDate: types.MustTimeFromString("2026-02-04T05:02:31.632Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `id`                                                                                             | `string`                                                                                         | :heavy_check_mark:                                                                               | Subscription ID                                                                                  |
| `body`                                                                                           | [types.ActivateDraftSubscriptionRequest](../../models/types/activatedraftsubscriptionrequest.md) | :heavy_check_mark:                                                                               | Activate Draft Subscription Request                                                              |
| `opts`                                                                                           | [][dtos.Option](../../models/dtos/option.md)                                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*dtos.ActivateSubscriptionResponse](../../models/dtos/activatesubscriptionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetSubscriptionAddonAssociations

Use when listing which add-ons are currently attached to a subscription (e.g. for display or editing).

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionAddonAssociations" method="get" path="/subscriptions/{id}/addons/associations" -->
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

    res, err := s.Subscriptions.GetSubscriptionAddonAssociations(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.AddonAssociationResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetSubscriptionAddonAssociationsResponse](../../models/dtos/getsubscriptionaddonassociationsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CancelSubscription

Use when a customer churns or downgrades. Supports immediate or end-of-period cancellation and proration. Ideal for self-serve or support-driven cancellations.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelSubscription" method="post" path="/subscriptions/{id}/cancel" -->
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

    res, err := s.Subscriptions.CancelSubscription(ctx, "<id>", types.CancelSubscriptionRequest{
        CancellationType: types.CancellationTypeImmediate,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CancelSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | Subscription ID                                                                    |
| `body`                                                                             | [types.CancelSubscriptionRequest](../../models/types/cancelsubscriptionrequest.md) | :heavy_check_mark:                                                                 | Cancel Subscription Request                                                        |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.CancelSubscriptionResponse](../../models/dtos/cancelsubscriptionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ExecuteSubscriptionChange

Use when applying a plan change (e.g. upgrade or downgrade). Executes proration and generates invoice or credit as needed.

### Example Usage

<!-- UsageSnippet language="go" operationID="executeSubscriptionChange" method="post" path="/subscriptions/{id}/change/execute" -->
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

    res, err := s.Subscriptions.ExecuteSubscriptionChange(ctx, "<id>", types.SubscriptionChangeRequest{
        BillingCadence: types.BillingCadenceRecurring,
        BillingCycle: types.BillingCycleAnniversary,
        BillingPeriod: types.BillingPeriodAnnual,
        ProrationBehavior: types.ProrationBehaviorNone,
        TargetPlanID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionChangeExecuteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | Subscription ID                                                                    |
| `body`                                                                             | [types.SubscriptionChangeRequest](../../models/types/subscriptionchangerequest.md) | :heavy_check_mark:                                                                 | Subscription change request                                                        |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.ExecuteSubscriptionChangeResponse](../../models/dtos/executesubscriptionchangeresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## PreviewSubscriptionChange

Use when showing a customer the cost of a plan change before they confirm (e.g. upgrade/downgrade preview with proration).

### Example Usage

<!-- UsageSnippet language="go" operationID="previewSubscriptionChange" method="post" path="/subscriptions/{id}/change/preview" -->
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

    res, err := s.Subscriptions.PreviewSubscriptionChange(ctx, "<id>", types.SubscriptionChangeRequest{
        BillingCadence: types.BillingCadenceRecurring,
        BillingCycle: types.BillingCycleAnniversary,
        BillingPeriod: types.BillingPeriodOnetime,
        ProrationBehavior: types.ProrationBehaviorNone,
        TargetPlanID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionChangePreviewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | Subscription ID                                                                    |
| `body`                                                                             | [types.SubscriptionChangeRequest](../../models/types/subscriptionchangerequest.md) | :heavy_check_mark:                                                                 | Subscription change preview request                                                |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.PreviewSubscriptionChangeResponse](../../models/dtos/previewsubscriptionchangeresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetSubscriptionEntitlements

Use when checking what features or limits a subscription has (e.g. entitlement checks or feature gating). Optional feature_ids to filter.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionEntitlements" method="get" path="/subscriptions/{id}/entitlements" -->
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

    res, err := s.Subscriptions.GetSubscriptionEntitlements(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `featureIds`                                          | []`string`                                            | :heavy_minus_sign:                                    | Feature IDs to filter by                              |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetSubscriptionEntitlementsResponse](../../models/dtos/getsubscriptionentitlementsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetSubscriptionUpcomingGrants

Use when showing upcoming or pending credits for a subscription (e.g. in a portal or for forecasting).

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionUpcomingGrants" method="get" path="/subscriptions/{id}/grants/upcoming" -->
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

    res, err := s.Subscriptions.GetSubscriptionUpcomingGrants(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCreditGrantApplicationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetSubscriptionUpcomingGrantsResponse](../../models/dtos/getsubscriptionupcominggrantsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CreateSubscriptionLineItem

Use when adding a new charge or seat to a subscription (e.g. extra seat or one-time add). Supports price_id or inline price.

### Example Usage

<!-- UsageSnippet language="go" operationID="createSubscriptionLineItem" method="post" path="/subscriptions/{id}/lineitems" -->
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

    res, err := s.Subscriptions.CreateSubscriptionLineItem(ctx, "<id>", types.CreateSubscriptionLineItemRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionLineItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `id`                                                                                               | `string`                                                                                           | :heavy_check_mark:                                                                                 | Subscription ID                                                                                    |
| `body`                                                                                             | [types.CreateSubscriptionLineItemRequest](../../models/types/createsubscriptionlineitemrequest.md) | :heavy_check_mark:                                                                                 | Create Line Item Request                                                                           |
| `opts`                                                                                             | [][dtos.Option](../../models/dtos/option.md)                                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*dtos.CreateSubscriptionLineItemResponse](../../models/dtos/createsubscriptionlineitemresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ExecuteSubscriptionModify

Execute a mid-cycle subscription modification (inheritance or quantity change).

### Example Usage

<!-- UsageSnippet language="go" operationID="executeSubscriptionModify" method="post" path="/subscriptions/{id}/modify/execute" -->
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

    res, err := s.Subscriptions.ExecuteSubscriptionModify(ctx, "<id>", types.ExecuteSubscriptionModifyRequest{
        Type: types.SubscriptionModifyTypeInheritance,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionModifyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `id`                                                                                             | `string`                                                                                         | :heavy_check_mark:                                                                               | Subscription ID                                                                                  |
| `body`                                                                                           | [types.ExecuteSubscriptionModifyRequest](../../models/types/executesubscriptionmodifyrequest.md) | :heavy_check_mark:                                                                               | Modification request                                                                             |
| `opts`                                                                                           | [][dtos.Option](../../models/dtos/option.md)                                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*dtos.ExecuteSubscriptionModifyResponse](../../models/dtos/executesubscriptionmodifyresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## PreviewSubscriptionModify

Preview the impact of a mid-cycle subscription modification without committing changes.

### Example Usage

<!-- UsageSnippet language="go" operationID="previewSubscriptionModify" method="post" path="/subscriptions/{id}/modify/preview" -->
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

    res, err := s.Subscriptions.PreviewSubscriptionModify(ctx, "<id>", types.ExecuteSubscriptionModifyRequest{
        Type: types.SubscriptionModifyTypeQuantityChange,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionModifyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `id`                                                                                             | `string`                                                                                         | :heavy_check_mark:                                                                               | Subscription ID                                                                                  |
| `body`                                                                                           | [types.ExecuteSubscriptionModifyRequest](../../models/types/executesubscriptionmodifyrequest.md) | :heavy_check_mark:                                                                               | Modification preview request                                                                     |
| `opts`                                                                                           | [][dtos.Option](../../models/dtos/option.md)                                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*dtos.PreviewSubscriptionModifyResponse](../../models/dtos/previewsubscriptionmodifyresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## PauseSubscription

Use when temporarily stopping a subscription (e.g. customer hold or seasonal pause). Billing and access pause; resume when ready.

### Example Usage

<!-- UsageSnippet language="go" operationID="pauseSubscription" method="post" path="/subscriptions/{id}/pause" -->
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

    res, err := s.Subscriptions.PauseSubscription(ctx, "<id>", types.PauseSubscriptionRequest{
        PauseMode: types.PauseModeScheduled,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionPauseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `id`                                                                             | `string`                                                                         | :heavy_check_mark:                                                               | Subscription ID                                                                  |
| `body`                                                                           | [types.PauseSubscriptionRequest](../../models/types/pausesubscriptionrequest.md) | :heavy_check_mark:                                                               | Pause subscription request                                                       |
| `opts`                                                                           | [][dtos.Option](../../models/dtos/option.md)                                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*dtos.PauseSubscriptionResponse](../../models/dtos/pausesubscriptionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ListSubscriptionPauses

Use when showing pause history for a subscription (e.g. support or audit). Returns all past and future pauses.

### Example Usage

<!-- UsageSnippet language="go" operationID="listSubscriptionPauses" method="get" path="/subscriptions/{id}/pauses" -->
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

    res, err := s.Subscriptions.ListSubscriptionPauses(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSubscriptionPausesResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.ListSubscriptionPausesResponse](../../models/dtos/listsubscriptionpausesresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ResumeSubscription

Use when reactivating a paused subscription (e.g. end of hold). Billing and access resume from the resume date.

### Example Usage

<!-- UsageSnippet language="go" operationID="resumeSubscription" method="post" path="/subscriptions/{id}/resume" -->
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

    res, err := s.Subscriptions.ResumeSubscription(ctx, "<id>", types.ResumeSubscriptionRequest{
        ResumeMode: types.ResumeModeAuto,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionPauseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | Subscription ID                                                                    |
| `body`                                                                             | [types.ResumeSubscriptionRequest](../../models/types/resumesubscriptionrequest.md) | :heavy_check_mark:                                                                 | Resume subscription request                                                        |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.ResumeSubscriptionResponse](../../models/dtos/resumesubscriptionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetSubscriptionV2

Use when you need a subscription with related data (line items, prices, plan). Supports expand for detailed payloads without extra round-trips.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionV2" method="get" path="/subscriptions/{id}/v2" -->
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

    res, err := s.Subscriptions.GetSubscriptionV2(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionResponseV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | `string`                                                                               | :heavy_check_mark:                                                                     | Subscription ID                                                                        |
| `expand`                                                                               | `*string`                                                                              | :heavy_minus_sign:                                                                     | Comma-separated list of fields to expand (e.g., 'subscription_line_items,prices,plan') |
| `opts`                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*dtos.GetSubscriptionV2Response](../../models/dtos/getsubscriptionv2response.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## ListAllSubscriptionSchedules

Use when listing or searching scheduled changes across subscriptions (e.g. admin view). Returns schedules with optional filtering.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAllSubscriptionSchedules" method="get" path="/v1/subscription-schedules" -->
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

    res, err := s.Subscriptions.ListAllSubscriptionSchedules(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetPendingSchedulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `pendingOnly`                                         | `*bool`                                               | :heavy_minus_sign:                                    | Filter to pending schedules only                      |
| `subscriptionID`                                      | `*string`                                             | :heavy_minus_sign:                                    | Filter by subscription ID                             |
| `limit`                                               | `*int64`                                              | :heavy_minus_sign:                                    | Limit results                                         |
| `offset`                                              | `*int64`                                              | :heavy_minus_sign:                                    | Offset for pagination                                 |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.ListAllSubscriptionSchedulesResponse](../../models/dtos/listallsubscriptionschedulesresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## GetSubscriptionSchedule

Use when you need to load a single scheduled change (e.g. to show when a plan change or renewal takes effect).

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionSchedule" method="get" path="/v1/subscription-schedules/{id}" -->
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

    res, err := s.Subscriptions.GetSubscriptionSchedule(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionScheduleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Schedule ID                                           |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetSubscriptionScheduleResponse](../../models/dtos/getsubscriptionscheduleresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## CancelSubscriptionSchedule

Use when cancelling a scheduled change (e.g. customer changed mind). Identify by schedule ID in path or by subscription ID + schedule type in body.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelSubscriptionSchedule" method="post" path="/v1/subscriptions/schedules/{schedule_id}/cancel" -->
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

    res, err := s.Subscriptions.CancelSubscriptionSchedule(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CancelScheduleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |
| `scheduleID`                                                                | `string`                                                                    | :heavy_check_mark:                                                          | Schedule ID (optional if using request body)                                |
| `body`                                                                      | [*types.CancelScheduleRequest](../../models/types/cancelschedulerequest.md) | :heavy_minus_sign:                                                          | Cancel request (optional if using path parameter)                           |
| `opts`                                                                      | [][dtos.Option](../../models/dtos/option.md)                                | :heavy_minus_sign:                                                          | The options for this request.                                               |

### Response

**[*dtos.CancelSubscriptionScheduleResponse](../../models/dtos/cancelsubscriptionscheduleresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## ListSubscriptionSchedules

Use when listing scheduled changes for a subscription (e.g. upcoming plan change or renewal). Returns all schedules for that subscription.

### Example Usage

<!-- UsageSnippet language="go" operationID="listSubscriptionSchedules" method="get" path="/v1/subscriptions/{subscription_id}/schedules" -->
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

    res, err := s.Subscriptions.ListSubscriptionSchedules(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetPendingSchedulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `subscriptionID`                                      | `string`                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.ListSubscriptionSchedulesResponse](../../models/dtos/listsubscriptionschedulesresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |