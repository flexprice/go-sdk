# SelfBilling

## Overview

### Available Operations

* [StartSelfBillingCheckout](#startselfbillingcheckout) - Start self-billing checkout
* [PollSelfBillingCheckout](#pollselfbillingcheckout) - Poll self-billing checkout
* [CancelSelfBillingCheckout](#cancelselfbillingcheckout) - Cancel self-billing checkout
* [CreateSelfBillingPortalSession](#createselfbillingportalsession) - Create self-billing portal session
* [GetSelfBillingStatus](#getselfbillingstatus) - Get self-billing status
* [GetSelfBillingSummary](#getselfbillingsummary) - Get self-billing summary
* [GetTenantBillingSummary](#gettenantbillingsummary) - Get tenant billing summary

## StartSelfBillingCheckout

Saves the billing details and opens the checkout that ends the caller tenant's trial.

### Example Usage

<!-- UsageSnippet language="go" operationID="startSelfBillingCheckout" method="post" path="/portal/billing/checkout" -->
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

    res, err := s.SelfBilling.StartSelfBillingCheckout(ctx, types.SelfBillingCheckoutRequest{
        AddressCity: "<value>",
        AddressCountry: "<value>",
        AddressLine1: "856 Hammes Creek",
        AddressPostalCode: "<value>",
        Email: "Pamela.Rogahn3@gmail.com",
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SelfBillingCheckoutResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [types.SelfBillingCheckoutRequest](../../models/types/selfbillingcheckoutrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][dtos.Option](../../models/dtos/option.md)                                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*dtos.StartSelfBillingCheckoutResponse](../../models/dtos/startselfbillingcheckoutresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404, 409        | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## PollSelfBillingCheckout

Reports the caller tenant's checkout and activates the tenant once it is paid.

### Example Usage

<!-- UsageSnippet language="go" operationID="pollSelfBillingCheckout" method="get" path="/portal/billing/checkout/{id}" -->
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

    res, err := s.SelfBilling.PollSelfBillingCheckout(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SelfBillingCheckoutResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Checkout session ID                                   |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.PollSelfBillingCheckoutResponse](../../models/dtos/pollselfbillingcheckoutresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CancelSelfBillingCheckout

Cancels the caller tenant's open checkout.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelSelfBillingCheckout" method="post" path="/portal/billing/checkout/{id}/cancel" -->
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

    res, err := s.SelfBilling.CancelSelfBillingCheckout(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Checkout session ID                                   |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.CancelSelfBillingCheckoutResponse](../../models/dtos/cancelselfbillingcheckoutresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |

## CreateSelfBillingPortalSession

Opens a short-lived customer portal session for the caller tenant.

### Example Usage

<!-- UsageSnippet language="go" operationID="createSelfBillingPortalSession" method="post" path="/portal/billing/portal-session" -->
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

    res, err := s.SelfBilling.CreateSelfBillingPortalSession(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.SelfBillingPortalResponse != nil {
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

**[*dtos.CreateSelfBillingPortalSessionResponse](../../models/dtos/createselfbillingportalsessionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetSelfBillingStatus

Reports whether self-billing is enabled and the caller tenant's billing status.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSelfBillingStatus" method="get" path="/portal/billing/status" -->
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

    res, err := s.SelfBilling.GetSelfBillingStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.SelfBillingStatusResponse != nil {
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

**[*dtos.GetSelfBillingStatusResponse](../../models/dtos/getselfbillingstatusresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetSelfBillingSummary

Returns billing status, recent invoices, usage and a plan with charges. Charges use expanded subscription line-item prices, including overrides; without a current subscription they use the configured self-serve plan. Pricing terms are not an invoice total.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSelfBillingSummary" method="get" path="/portal/billing/summary" -->
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

    res, err := s.SelfBilling.GetSelfBillingSummary(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.SelfBillingSummaryResponse != nil {
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

**[*dtos.GetSelfBillingSummaryResponse](../../models/dtos/getselfbillingsummaryresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetTenantBillingSummary

Returns the caller tenant's billing summary; same as /portal/billing/summary.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTenantBillingSummary" method="get" path="/tenants/billing" -->
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

    res, err := s.SelfBilling.GetTenantBillingSummary(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.SelfBillingSummaryResponse != nil {
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

**[*dtos.GetTenantBillingSummaryResponse](../../models/dtos/gettenantbillingsummaryresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |