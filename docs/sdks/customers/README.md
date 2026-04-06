# Customers

## Overview

### Available Operations

* [UpdateCustomer](#updatecustomer) - Update customer
* [CreateCustomer](#createcustomer) - Create customer
* [GetCustomerByExternalID](#getcustomerbyexternalid) - Get customer by external ID
* [QueryCustomer](#querycustomer) - Query customers
* [GetCustomerUsageSummary](#getcustomerusagesummary) - Get customer usage summary
* [GetCustomer](#getcustomer) - Get customer
* [DeleteCustomer](#deletecustomer) - Delete customer
* [GetCustomerEntitlements](#getcustomerentitlements) - Get customer entitlements
* [GetCustomerUpcomingGrants](#getcustomerupcominggrants) - Get upcoming credit grant applications

## UpdateCustomer

Use when updating customer details (e.g. name, email, or metadata). Identify by id or external_customer_id.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCustomer" method="put" path="/customers" -->
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

    res, err := s.Customers.UpdateCustomer(ctx, types.UpdateCustomerRequest{}, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Customer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `body`                                                                     | [types.UpdateCustomerRequest](../../models/types/updatecustomerrequest.md) | :heavy_check_mark:                                                         | Customer                                                                   |
| `id`                                                                       | `*string`                                                                  | :heavy_minus_sign:                                                         | Customer ID                                                                |
| `externalCustomerID`                                                       | `*string`                                                                  | :heavy_minus_sign:                                                         | Customer External ID                                                       |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.UpdateCustomerResponse](../../models/dtos/updatecustomerresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CreateCustomer

Use when onboarding a new billing customer (e.g. sign-up or CRM sync). Ideal for linking via external_customer_id to your app's user id.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCustomer" method="post" path="/customers" -->
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

    res, err := s.Customers.CreateCustomer(ctx, types.CreateCustomerRequest{
        ExternalID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Customer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [types.CreateCustomerRequest](../../models/types/createcustomerrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.CreateCustomerResponse](../../models/dtos/createcustomerresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCustomerByExternalID

Use when resolving a customer by your app's id (e.g. from your user table). Ideal for integrations that key by external id.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomerByExternalId" method="get" path="/customers/external/{external_id}" -->
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

    res, err := s.Customers.GetCustomerByExternalID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Customer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `externalID`                                          | `string`                                              | :heavy_check_mark:                                    | Customer External ID                                  |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetCustomerByExternalIDResponse](../../models/dtos/getcustomerbyexternalidresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryCustomer

Use when listing or searching customers (e.g. admin CRM or reporting). Returns a paginated list; supports filtering and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryCustomer" method="post" path="/customers/search" -->
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

    res, err := s.Customers.QueryCustomer(ctx, types.CustomerFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListCustomersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [types.CustomerFilter](../../models/types/customerfilter.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][dtos.Option](../../models/dtos/option.md)                 | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*dtos.QueryCustomerResponse](../../models/dtos/querycustomerresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCustomerUsageSummary

Use when showing a customer's usage (e.g. portal or overage alerts). Identify by customer_id or customer_lookup_key; supports filters.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomerUsageSummary" method="get" path="/customers/usage" -->
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

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Customers.GetCustomerUsageSummary(ctx, dtos.GetCustomerUsageSummaryRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerUsageSummaryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [dtos.GetCustomerUsageSummaryRequest](../../models/dtos/getcustomerusagesummaryrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][dtos.Option](../../models/dtos/option.md)                                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*dtos.GetCustomerUsageSummaryResponse](../../models/dtos/getcustomerusagesummaryresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCustomer

Use when you need to load a single customer (e.g. for a billing portal or to attach a subscription).

### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomer" method="get" path="/customers/{id}" -->
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

    res, err := s.Customers.GetCustomer(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Customer != nil {
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

**[*dtos.GetCustomerResponse](../../models/dtos/getcustomerresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeleteCustomer

Use when removing a customer (e.g. GDPR or churn). Returns 204 No Content on success.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCustomer" method="delete" path="/customers/{id}" -->
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

    res, err := s.Customers.DeleteCustomer(ctx, "<id>")
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
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Customer ID                                           |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.DeleteCustomerResponse](../../models/dtos/deletecustomerresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCustomerEntitlements

Use when checking what a customer can access (e.g. feature gating or usage limits). Supports optional filters (feature_ids, subscription_ids).

### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomerEntitlements" method="get" path="/customers/{id}/entitlements" -->
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

    res, err := s.Customers.GetCustomerEntitlements(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerEntitlementsResponse != nil {
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

**[*dtos.GetCustomerEntitlementsResponse](../../models/dtos/getcustomerentitlementsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetCustomerUpcomingGrants

Use when showing upcoming or pending credits for a customer (e.g. in a portal or for forecasting).

### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomerUpcomingGrants" method="get" path="/customers/{id}/grants/upcoming" -->
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

    res, err := s.Customers.GetCustomerUpcomingGrants(ctx, "<id>")
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
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Customer ID                                           |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetCustomerUpcomingGrantsResponse](../../models/dtos/getcustomerupcominggrantsresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |