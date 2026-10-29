# Commitments

## Overview

### Available Operations

* [CreateCommitment](#createcommitment) - Create commitment
* [PreviewCommitment](#previewcommitment) - Preview commitment
* [QueryCommitment](#querycommitment) - Query commitments
* [GetCommitment](#getcommitment) - Get commitment
* [UpdateCommitment](#updatecommitment) - Update commitment
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

## TerminateCommitment

Ends a commitment early, effective at the given time or now.

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