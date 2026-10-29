# Contracts

## Overview

### Available Operations

* [QueryContract](#querycontract) - Query contracts
* [GetContract](#getcontract) - Get contract

## QueryContract

Use when searching revenue recognition contracts by customer, subscription or contract status. Returns a paginated list.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryContract" method="post" path="/contracts/search" -->
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

    res, err := s.Contracts.QueryContract(ctx, types.ContractFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListContractsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [types.ContractFilter](../../models/types/contractfilter.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][dtos.Option](../../models/dtos/option.md)                 | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*dtos.QueryContractResponse](../../models/dtos/querycontractresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetContract

Use when you need to load a single revenue recognition contract with the subscriptions it covers. Contracts are derived from subscriptions and are read-only.

### Example Usage

<!-- UsageSnippet language="go" operationID="getContract" method="get" path="/contracts/{id}" -->
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

    res, err := s.Contracts.GetContract(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ContractResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | `string`                                              | :heavy_check_mark:                                    | Contract ID                                           |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetContractResponse](../../models/dtos/getcontractresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |