# Checkout

## Overview

### Available Operations

* [CreateCheckoutSession](#createcheckoutsession) - Create checkout session
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
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Checkout.CreateCheckoutSession(ctx, types.CreateCheckoutSessionRequest{
        Action: types.CheckoutActionPayInvoice,
        CustomerExternalID: "<id>",
        PaymentProvider: types.CheckoutPaymentProviderRazorpay,
    }, dtos.CreateCheckoutSessionSecurity{
        Option1: &dtos.CreateCheckoutSessionSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
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
| `security`                                                                               | [dtos.CreateCheckoutSessionSecurity](../../models/dtos/createcheckoutsessionsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*dtos.CreateCheckoutSessionResponse](../../models/dtos/createcheckoutsessionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 409             | application/json     |
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
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Checkout.GetCheckoutSession(ctx, dtos.GetCheckoutSessionSecurity{
        Option1: &dtos.GetCheckoutSessionSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckoutSessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `security`                                                                         | [dtos.GetCheckoutSessionSecurity](../../models/dtos/getcheckoutsessionsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |
| `id`                                                                               | `string`                                                                           | :heavy_check_mark:                                                                 | Checkout session ID                                                                |
| `opts`                                                                             | [][dtos.Option](../../models/dtos/option.md)                                       | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*dtos.GetCheckoutSessionResponse](../../models/dtos/getcheckoutsessionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 404                  | application/json     |
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
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Checkout.DeleteCheckoutSession(ctx, dtos.DeleteCheckoutSessionSecurity{
        Option1: &dtos.DeleteCheckoutSessionSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `security`                                                                               | [dtos.DeleteCheckoutSessionSecurity](../../models/dtos/deletecheckoutsessionsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |
| `id`                                                                                     | `string`                                                                                 | :heavy_check_mark:                                                                       | Checkout session ID                                                                      |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

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
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Checkout.CancelCheckoutSession(ctx, dtos.CancelCheckoutSessionSecurity{
        Option1: &dtos.CancelCheckoutSessionSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
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
| `security`                                                                               | [dtos.CancelCheckoutSessionSecurity](../../models/dtos/cancelcheckoutsessionsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |
| `id`                                                                                     | `string`                                                                                 | :heavy_check_mark:                                                                       | Checkout session ID                                                                      |
| `opts`                                                                                   | [][dtos.Option](../../models/dtos/option.md)                                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*dtos.CancelCheckoutSessionResponse](../../models/dtos/cancelcheckoutsessionresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |