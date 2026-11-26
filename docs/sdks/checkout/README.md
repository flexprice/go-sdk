# Checkout

## Overview

### Available Operations

* [CreateCheckoutSession](#createcheckoutsession) - Create checkout session
* [QueryCheckoutSession](#querycheckoutsession) - Query checkout sessions
* [GetCheckoutSession](#getcheckoutsession) - Get checkout session
* [DeleteCheckoutSession](#deletecheckoutsession) - Delete checkout session
* [CancelCheckoutSession](#cancelcheckoutsession) - Cancel checkout session

## CreateCheckoutSession

Create checkout session

### Example Usage

<!-- UsageSnippet language="go" operationID="createCheckoutSession" method="post" path="/checkout/sessions" -->
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

    res, err := s.Checkout.CreateCheckoutSession(ctx, types.CreateCheckoutSessionRequest{
        Action: types.CheckoutActionPayInvoice,
        CustomerExternalID: "<id>",
        PaymentProvider: types.CheckoutPaymentProviderRazorpay,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckoutSessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [types.CreateCheckoutSessionRequest](../../models/types/createcheckoutsessionrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*dtos.CreateCheckoutSessionResponse](../../models/dtos/createcheckoutsessionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 409             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryCheckoutSession

Use when listing checkout sessions, e.g. a customer's sessions. Expand is not supported here; read one session with GET /checkout/sessions/{id}?expand=invoice,payment for its invoice, payment and attempts.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryCheckoutSession" method="post" path="/checkout/sessions/search" -->
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

    res, err := s.Checkout.QueryCheckoutSession(ctx, types.CheckoutSessionFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCheckoutSessionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [types.CheckoutSessionFilter](../../models/types/checkoutsessionfilter.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.QueryCheckoutSessionResponse](../../models/dtos/querycheckoutsessionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCheckoutSession

Get checkout session

### Example Usage

<!-- UsageSnippet language="go" operationID="getCheckoutSession" method="get" path="/checkout/sessions/{id}" -->
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

    res, err := s.Checkout.GetCheckoutSession(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckoutSessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Checkout session ID                                   |
| `expand`                                              | `*string`                                             | :heavy_minus_sign:                                    | Comma-separated fields to expand: invoice, payment    |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetCheckoutSessionResponse](../../models/dtos/getcheckoutsessionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeleteCheckoutSession

Delete checkout session

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCheckoutSession" method="delete" path="/checkout/sessions/{id}" -->
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

    res, err := s.Checkout.DeleteCheckoutSession(ctx, "<id>")
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

**[*dtos.DeleteCheckoutSessionResponse](../../models/dtos/deletecheckoutsessionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CancelCheckoutSession

Cancel checkout session

### Example Usage

<!-- UsageSnippet language="go" operationID="cancelCheckoutSession" method="post" path="/checkout/sessions/{id}/cancel" -->
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

    res, err := s.Checkout.CancelCheckoutSession(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckoutSessionResponse != nil {
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

**[*dtos.CancelCheckoutSessionResponse](../../models/dtos/cancelcheckoutsessionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |