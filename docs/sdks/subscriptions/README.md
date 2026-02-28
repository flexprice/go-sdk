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

    res, err := s.Subscriptions.CreateSubscription(ctx, types.DtoCreateSubscriptionRequest{
        BillingCadence: types.BillingCadenceOnetime,
        BillingPeriod: types.BillingPeriodDaily,
        Currency: "New Leu",
        PlanID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [types.DtoCreateSubscriptionRequest](../../models/types/dtocreatesubscriptionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*dtos.CreateSubscriptionResponse](../../models/dtos/createsubscriptionresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## AddSubscriptionAddon

Use when adding an optional product or add-on to an existing subscription (e.g. extra storage or support tier).

### Example Usage

<!-- UsageSnippet language="go" operationID="addSubscriptionAddon" method="post" path="/subscriptions/addon" -->
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

    res, err := s.Subscriptions.AddSubscriptionAddon(ctx, types.DtoAddAddonRequest{
        AddonID: "<id>",
        SubscriptionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoAddonAssociationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [types.DtoAddAddonRequest](../../models/types/dtoaddaddonrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.AddSubscriptionAddonResponse](../../models/dtos/addsubscriptionaddonresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## RemoveSubscriptionAddon

Use when removing an add-on from a subscription (e.g. downgrade or opt-out).

### Example Usage

<!-- UsageSnippet language="go" operationID="removeSubscriptionAddon" method="delete" path="/subscriptions/addon" -->
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

    res, err := s.Subscriptions.RemoveSubscriptionAddon(ctx, types.DtoRemoveAddonRequest{
        AddonAssociationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [types.DtoRemoveAddonRequest](../../models/types/dtoremoveaddonrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.RemoveSubscriptionAddonResponse](../../models/dtos/removesubscriptionaddonresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateSubscriptionLineItem

Use when changing a subscription line item (e.g. quantity or price). Implemented by ending the current line and creating a new one for clean billing.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSubscriptionLineItem" method="put" path="/subscriptions/lineitems/{id}" -->
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

    res, err := s.Subscriptions.UpdateSubscriptionLineItem(ctx, "<id>", types.DtoUpdateSubscriptionLineItemRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionLineItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `id`                                                                                                     | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Line Item ID                                                                                             |
| `body`                                                                                                   | [types.DtoUpdateSubscriptionLineItemRequest](../../models/types/dtoupdatesubscriptionlineitemrequest.md) | :heavy_check_mark:                                                                                       | Update Line Item Request                                                                                 |
| `opts`                                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*dtos.UpdateSubscriptionLineItemResponse](../../models/dtos/updatesubscriptionlineitemresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## DeleteSubscriptionLineItem

Use when removing a charge or seat from a subscription (e.g. downgrade). Line item ends; retained for history but no longer billed.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSubscriptionLineItem" method="delete" path="/subscriptions/lineitems/{id}" -->
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

    res, err := s.Subscriptions.DeleteSubscriptionLineItem(ctx, "<id>", types.DtoDeleteSubscriptionLineItemRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionLineItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `id`                                                                                                     | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Line Item ID                                                                                             |
| `body`                                                                                                   | [types.DtoDeleteSubscriptionLineItemRequest](../../models/types/dtodeletesubscriptionlineitemrequest.md) | :heavy_check_mark:                                                                                       | Delete Line Item Request                                                                                 |
| `opts`                                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*dtos.DeleteSubscriptionLineItemResponse](../../models/dtos/deletesubscriptionlineitemresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## QuerySubscription

Use when listing or searching subscriptions (e.g. admin view or customer subscription list). Returns a paginated list; supports filtering by customer, plan, status.

### Example Usage

<!-- UsageSnippet language="go" operationID="querySubscription" method="post" path="/subscriptions/search" -->
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

    res, err := s.Subscriptions.QuerySubscription(ctx, types.SubscriptionFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListSubscriptionsResponse != nil {
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

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetSubscriptionUsage

Use when showing usage for a subscription (e.g. in a portal or for overage checks). Supports time range and filters.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionUsage" method="post" path="/subscriptions/usage" -->
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

    res, err := s.Subscriptions.GetSubscriptionUsage(ctx, types.DtoGetUsageBySubscriptionRequest{
        EndTime: flexprice.Pointer("2024-03-20T00:00:00Z"),
        LifetimeUsage: flexprice.Pointer(false),
        StartTime: flexprice.Pointer("2024-03-13T00:00:00Z"),
        SubscriptionID: "123",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetUsageBySubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [types.DtoGetUsageBySubscriptionRequest](../../models/types/dtogetusagebysubscriptionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][dtos.Option](../../models/dtos/option.md)                                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*dtos.GetSubscriptionUsageResponse](../../models/dtos/getsubscriptionusageresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetSubscription

Use when you need to load a single subscription (e.g. for a billing portal or to check status).

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscription" method="get" path="/subscriptions/{id}" -->
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

    res, err := s.Subscriptions.GetSubscription(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetSubscriptionResponse](../../models/dtos/getsubscriptionresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateSubscription

Use when changing subscription details (e.g. quantity, billing anchor, or parent). Supports partial update; send "" to clear parent_subscription_id.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSubscription" method="put" path="/subscriptions/{id}" -->
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

    res, err := s.Subscriptions.UpdateSubscription(ctx, "<id>", types.DtoUpdateSubscriptionRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Subscription ID                                                                          |
| `body`                                                                                   | [types.DtoUpdateSubscriptionRequest](../../models/types/dtoupdatesubscriptionrequest.md) | :heavy_check_mark:                                                                       | Update Subscription Request                                                              |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*dtos.UpdateSubscriptionResponse](../../models/dtos/updatesubscriptionresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## ActivateSubscription

Use when turning a draft subscription live (e.g. after collecting payment or completing setup). Once activated, billing and entitlements apply.

### Example Usage

<!-- UsageSnippet language="go" operationID="activateSubscription" method="post" path="/subscriptions/{id}/activate" -->
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

    res, err := s.Subscriptions.ActivateSubscription(ctx, "<id>", types.DtoActivateDraftSubscriptionRequest{
        StartDate: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `id`                                                                                                   | *string*                                                                                               | :heavy_check_mark:                                                                                     | Subscription ID                                                                                        |
| `body`                                                                                                 | [types.DtoActivateDraftSubscriptionRequest](../../models/types/dtoactivatedraftsubscriptionrequest.md) | :heavy_check_mark:                                                                                     | Activate Draft Subscription Request                                                                    |
| `opts`                                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*dtos.ActivateSubscriptionResponse](../../models/dtos/activatesubscriptionresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetSubscriptionAddonAssociations

Use when listing which add-ons are currently attached to a subscription (e.g. for display or editing).

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionAddonAssociations" method="get" path="/subscriptions/{id}/addons/associations" -->
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

    res, err := s.Subscriptions.GetSubscriptionAddonAssociations(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoAddonAssociationResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetSubscriptionAddonAssociationsResponse](../../models/dtos/getsubscriptionaddonassociationsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## CancelSubscription

Use when a customer churns or downgrades. Supports immediate or end-of-period cancellation and proration. Ideal for self-serve or support-driven cancellations.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelSubscription" method="post" path="/subscriptions/{id}/cancel" -->
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

    res, err := s.Subscriptions.CancelSubscription(ctx, "<id>", types.DtoCancelSubscriptionRequest{
        CancellationType: types.CancellationTypeImmediate,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCancelSubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Subscription ID                                                                          |
| `body`                                                                                   | [types.DtoCancelSubscriptionRequest](../../models/types/dtocancelsubscriptionrequest.md) | :heavy_check_mark:                                                                       | Cancel Subscription Request                                                              |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*dtos.CancelSubscriptionResponse](../../models/dtos/cancelsubscriptionresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## ExecuteSubscriptionChange

Use when applying a plan change (e.g. upgrade or downgrade). Executes proration and generates invoice or credit as needed.

### Example Usage

<!-- UsageSnippet language="go" operationID="executeSubscriptionChange" method="post" path="/subscriptions/{id}/change/execute" -->
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

    res, err := s.Subscriptions.ExecuteSubscriptionChange(ctx, "<id>", types.DtoSubscriptionChangeRequest{
        BillingCadence: types.BillingCadenceRecurring,
        BillingCycle: types.BillingCycleAnniversary,
        BillingPeriod: types.BillingPeriodQuarterly,
        ProrationBehavior: types.ProrationBehaviorCreateProrations,
        TargetPlanID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionChangeExecuteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Subscription ID                                                                          |
| `body`                                                                                   | [types.DtoSubscriptionChangeRequest](../../models/types/dtosubscriptionchangerequest.md) | :heavy_check_mark:                                                                       | Subscription change request                                                              |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*dtos.ExecuteSubscriptionChangeResponse](../../models/dtos/executesubscriptionchangeresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## PreviewSubscriptionChange

Use when showing a customer the cost of a plan change before they confirm (e.g. upgrade/downgrade preview with proration).

### Example Usage

<!-- UsageSnippet language="go" operationID="previewSubscriptionChange" method="post" path="/subscriptions/{id}/change/preview" -->
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

    res, err := s.Subscriptions.PreviewSubscriptionChange(ctx, "<id>", types.DtoSubscriptionChangeRequest{
        BillingCadence: types.BillingCadenceRecurring,
        BillingCycle: types.BillingCycleCalendar,
        BillingPeriod: types.BillingPeriodHalfYearly,
        ProrationBehavior: types.ProrationBehaviorCreateProrations,
        TargetPlanID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionChangePreviewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Subscription ID                                                                          |
| `body`                                                                                   | [types.DtoSubscriptionChangeRequest](../../models/types/dtosubscriptionchangerequest.md) | :heavy_check_mark:                                                                       | Subscription change preview request                                                      |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*dtos.PreviewSubscriptionChangeResponse](../../models/dtos/previewsubscriptionchangeresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetSubscriptionEntitlements

Use when checking what features or limits a subscription has (e.g. entitlement checks or feature gating). Optional feature_ids to filter.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionEntitlements" method="get" path="/subscriptions/{id}/entitlements" -->
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

    res, err := s.Subscriptions.GetSubscriptionEntitlements(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `featureIds`                                          | []*string*                                            | :heavy_minus_sign:                                    | Feature IDs to filter by                              |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetSubscriptionEntitlementsResponse](../../models/dtos/getsubscriptionentitlementsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetSubscriptionUpcomingGrants

Use when showing upcoming or pending credits for a subscription (e.g. in a portal or for forecasting).

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionUpcomingGrants" method="get" path="/subscriptions/{id}/grants/upcoming" -->
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

    res, err := s.Subscriptions.GetSubscriptionUpcomingGrants(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListCreditGrantApplicationsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetSubscriptionUpcomingGrantsResponse](../../models/dtos/getsubscriptionupcominggrantsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## CreateSubscriptionLineItem

Use when adding a new charge or seat to a subscription (e.g. extra seat or one-time add). Supports price_id or inline price.

### Example Usage

<!-- UsageSnippet language="go" operationID="createSubscriptionLineItem" method="post" path="/subscriptions/{id}/lineitems" -->
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

    res, err := s.Subscriptions.CreateSubscriptionLineItem(ctx, "<id>", types.DtoCreateSubscriptionLineItemRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionLineItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `id`                                                                                                     | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Subscription ID                                                                                          |
| `body`                                                                                                   | [types.DtoCreateSubscriptionLineItemRequest](../../models/types/dtocreatesubscriptionlineitemrequest.md) | :heavy_check_mark:                                                                                       | Create Line Item Request                                                                                 |
| `opts`                                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*dtos.CreateSubscriptionLineItemResponse](../../models/dtos/createsubscriptionlineitemresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## PauseSubscription

Use when temporarily stopping a subscription (e.g. customer hold or seasonal pause). Billing and access pause; resume when ready.

### Example Usage

<!-- UsageSnippet language="go" operationID="pauseSubscription" method="post" path="/subscriptions/{id}/pause" -->
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

    res, err := s.Subscriptions.PauseSubscription(ctx, "<id>", types.DtoPauseSubscriptionRequest{
        PauseMode: types.PauseModeScheduled,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionPauseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | Subscription ID                                                                        |
| `body`                                                                                 | [types.DtoPauseSubscriptionRequest](../../models/types/dtopausesubscriptionrequest.md) | :heavy_check_mark:                                                                     | Pause subscription request                                                             |
| `opts`                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*dtos.PauseSubscriptionResponse](../../models/dtos/pausesubscriptionresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## ListSubscriptionPauses

Use when showing pause history for a subscription (e.g. support or audit). Returns all past and future pauses.

### Example Usage

<!-- UsageSnippet language="go" operationID="listSubscriptionPauses" method="get" path="/subscriptions/{id}/pauses" -->
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

    res, err := s.Subscriptions.ListSubscriptionPauses(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListSubscriptionPausesResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.ListSubscriptionPausesResponse](../../models/dtos/listsubscriptionpausesresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## ResumeSubscription

Use when reactivating a paused subscription (e.g. end of hold). Billing and access resume from the resume date.

### Example Usage

<!-- UsageSnippet language="go" operationID="resumeSubscription" method="post" path="/subscriptions/{id}/resume" -->
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

    res, err := s.Subscriptions.ResumeSubscription(ctx, "<id>", types.DtoResumeSubscriptionRequest{
        ResumeMode: types.ResumeModeAuto,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionPauseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `id`                                                                                     | *string*                                                                                 | :heavy_check_mark:                                                                       | Subscription ID                                                                          |
| `body`                                                                                   | [types.DtoResumeSubscriptionRequest](../../models/types/dtoresumesubscriptionrequest.md) | :heavy_check_mark:                                                                       | Resume subscription request                                                              |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*dtos.ResumeSubscriptionResponse](../../models/dtos/resumesubscriptionresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetSubscriptionV2

Use when you need a subscription with related data (line items, prices, plan). Supports expand for detailed payloads without extra round-trips.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSubscriptionV2" method="get" path="/subscriptions/{id}/v2" -->
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

    res, err := s.Subscriptions.GetSubscriptionV2(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionResponseV2 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | Subscription ID                                                                        |
| `expand`                                                                               | **string*                                                                              | :heavy_minus_sign:                                                                     | Comma-separated list of fields to expand (e.g., 'subscription_line_items,prices,plan') |
| `opts`                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*dtos.GetSubscriptionV2Response](../../models/dtos/getsubscriptionv2response.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## ListAllSubscriptionSchedules

Use when listing or searching scheduled changes across subscriptions (e.g. admin view). Returns schedules with optional filtering.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAllSubscriptionSchedules" method="get" path="/v1/subscription-schedules" -->
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

    res, err := s.Subscriptions.ListAllSubscriptionSchedules(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetPendingSchedulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `pendingOnly`                                         | **bool*                                               | :heavy_minus_sign:                                    | Filter to pending schedules only                      |
| `subscriptionID`                                      | **string*                                             | :heavy_minus_sign:                                    | Filter by subscription ID                             |
| `limit`                                               | **int64*                                              | :heavy_minus_sign:                                    | Limit results                                         |
| `offset`                                              | **int64*                                              | :heavy_minus_sign:                                    | Offset for pagination                                 |
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
	flexprice "github.com/flexprice/flexprice-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.GetSubscriptionSchedule(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoSubscriptionScheduleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Schedule ID                                           |
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
	flexprice "github.com/flexprice/flexprice-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.CancelSubscriptionSchedule(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoCancelScheduleResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `scheduleID`                                                                      | *string*                                                                          | :heavy_check_mark:                                                                | Schedule ID (optional if using request body)                                      |
| `body`                                                                            | [*types.DtoCancelScheduleRequest](../../models/types/dtocancelschedulerequest.md) | :heavy_minus_sign:                                                                | Cancel request (optional if using path parameter)                                 |
| `opts`                                                                            | [][dtos.Option](../../models/dtos/option.md)                                      | :heavy_minus_sign:                                                                | The options for this request.                                                     |

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
	flexprice "github.com/flexprice/flexprice-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        "https://api.example.com",
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Subscriptions.ListSubscriptionSchedules(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoGetPendingSchedulesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `subscriptionID`                                      | *string*                                              | :heavy_check_mark:                                    | Subscription ID                                       |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.ListSubscriptionSchedulesResponse](../../models/dtos/listsubscriptionschedulesresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |