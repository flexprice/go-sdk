# Wallets

## Overview

### Available Operations

* [GetCustomerWallets](#getcustomerwallets) - Get Customer Wallets
* [GetWalletsByCustomerID](#getwalletsbycustomerid) - Get wallets by customer ID
* [CreateWallet](#createwallet) - Create a new wallet
* [QueryWallet](#querywallet) - Query wallets
* [QueryWalletTransaction](#querywallettransaction) - Query wallet transactions
* [GetWallet](#getwallet) - Get wallet
* [UpdateWallet](#updatewallet) - Update a wallet
* [GetWalletBalance](#getwalletbalance) - Get wallet balance
* [TerminateWallet](#terminatewallet) - Terminate a wallet
* [TopUpWallet](#topupwallet) - Top up wallet
* [GetWalletTransactions](#getwallettransactions) - Get wallet transactions

## GetCustomerWallets

Use when resolving wallets by external customer id or lookup key (e.g. from your app's user id). Supports optional real-time balance and expand.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCustomerWallets" method="get" path="/customers/wallets" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.GetCustomerWallets(ctx, dtos.GetCustomerWalletsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoWalletResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [dtos.GetCustomerWalletsRequest](../../models/dtos/getcustomerwalletsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][dtos.Option](../../models/dtos/option.md)                                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*dtos.GetCustomerWalletsResponse](../../models/dtos/getcustomerwalletsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetWalletsByCustomerID

Use when showing a customer's wallets (e.g. balance overview by currency or in a billing portal). Supports optional expand for balance breakdown.

### Example Usage

<!-- UsageSnippet language="go" operationID="getWalletsByCustomerId" method="get" path="/customers/{id}/wallets" -->
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
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.GetWalletsByCustomerID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoWalletResponses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Customer ID                                           |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetWalletsByCustomerIDResponse](../../models/dtos/getwalletsbycustomeridresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## CreateWallet

Use when giving a customer a prepaid or credit balance (e.g. prepaid plans or promotional credits).

### Example Usage

<!-- UsageSnippet language="go" operationID="createWallet" method="post" path="/wallets" -->
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
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.CreateWallet(ctx, types.DtoCreateWalletRequest{
        Currency: "Seychelles Rupee",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoWalletResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [types.DtoCreateWalletRequest](../../models/types/dtocreatewalletrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.CreateWalletResponse](../../models/dtos/createwalletresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## QueryWallet

Use when listing or searching wallets (e.g. admin view or reporting). Returns a paginated list; supports filtering by customer and status.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryWallet" method="post" path="/wallets/search" -->
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
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.QueryWallet(ctx, types.WalletFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResponseDtoWalletResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [types.WalletFilter](../../models/types/walletfilter.md) | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][dtos.Option](../../models/dtos/option.md)             | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*dtos.QueryWalletResponse](../../models/dtos/querywalletresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## QueryWalletTransaction

Use when searching or reporting on wallet transactions (e.g. cross-wallet history or reconciliation). Returns a paginated list; supports filtering by wallet, customer, type, date range.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryWalletTransaction" method="post" path="/wallets/transactions/search" -->
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
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.QueryWalletTransaction(ctx, types.WalletTransactionFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListWalletTransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [types.WalletTransactionFilter](../../models/types/wallettransactionfilter.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][dtos.Option](../../models/dtos/option.md)                                   | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*dtos.QueryWalletTransactionResponse](../../models/dtos/querywallettransactionresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400                        | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetWallet

Use when you need to load a single wallet (e.g. for a balance or settings view).

### Example Usage

<!-- UsageSnippet language="go" operationID="getWallet" method="get" path="/wallets/{id}" -->
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
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.GetWallet(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoWalletResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Wallet ID                                             |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetWalletResponse](../../models/dtos/getwalletresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## UpdateWallet

Use when changing wallet settings (e.g. enabling or updating auto top-up thresholds).

### Example Usage

<!-- UsageSnippet language="go" operationID="updateWallet" method="put" path="/wallets/{id}" -->
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
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.UpdateWallet(ctx, "<id>", types.DtoUpdateWalletRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoWalletResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `id`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | Wallet ID                                                                    |
| `body`                                                                       | [types.DtoUpdateWalletRequest](../../models/types/dtoupdatewalletrequest.md) | :heavy_check_mark:                                                           | Update wallet request                                                        |
| `opts`                                                                       | [][dtos.Option](../../models/dtos/option.md)                                 | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*dtos.UpdateWalletResponse](../../models/dtos/updatewalletresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetWalletBalance

Use when displaying or checking current wallet balance (e.g. before charging or in a portal). Supports optional expand for credits breakdown and from_cache.

### Example Usage

<!-- UsageSnippet language="go" operationID="getWalletBalance" method="get" path="/wallets/{id}/balance/real-time" -->
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
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.GetWalletBalance(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoWalletBalanceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Wallet ID                                             |
| `expand`                                              | **string*                                             | :heavy_minus_sign:                                    | Expand fields (e.g., credits_available_breakdown)     |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.GetWalletBalanceResponse](../../models/dtos/getwalletbalanceresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## TerminateWallet

Use when closing a customer wallet (e.g. churn or migration). Closes the wallet and applies remaining balance per policy (refund or forfeit).

### Example Usage

<!-- UsageSnippet language="go" operationID="terminateWallet" method="post" path="/wallets/{id}/terminate" -->
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
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.TerminateWallet(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoWalletResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | Wallet ID                                             |
| `opts`                                                | [][dtos.Option](../../models/dtos/option.md)          | :heavy_minus_sign:                                    | The options for this request.                         |

### Response

**[*dtos.TerminateWalletResponse](../../models/dtos/terminatewalletresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## TopUpWallet

Use when adding funds to a wallet (e.g. top-up, refund, or manual credit). Supports optional idempotency via reference.

### Example Usage

<!-- UsageSnippet language="go" operationID="topUpWallet" method="post" path="/wallets/{id}/top-up" -->
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
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.TopUpWallet(ctx, "<id>", types.DtoTopUpWalletRequest{
        TransactionReason: types.TransactionReasonCreditAdjustment,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoTopUpWalletResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `id`                                                                       | *string*                                                                   | :heavy_check_mark:                                                         | Wallet ID                                                                  |
| `body`                                                                     | [types.DtoTopUpWalletRequest](../../models/types/dtotopupwalletrequest.md) | :heavy_check_mark:                                                         | Top up request                                                             |
| `opts`                                                                     | [][dtos.Option](../../models/dtos/option.md)                               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*dtos.TopUpWalletResponse](../../models/dtos/topupwalletresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |

## GetWalletTransactions

Use when showing transaction history for a wallet (e.g. credit/debit log or audit). Returns a paginated list; supports limit, offset, and filters.

### Example Usage

<!-- UsageSnippet language="go" operationID="getWalletTransactions" method="get" path="/wallets/{id}/transactions" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/models/dtos"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New(
        flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Wallets.GetWalletTransactions(ctx, dtos.GetWalletTransactionsRequest{
        IDPathParameter: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DtoListWalletTransactionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [dtos.GetWalletTransactionsRequest](../../models/dtos/getwallettransactionsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                           | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*dtos.GetWalletTransactionsResponse](../../models/dtos/getwallettransactionsresponse.md), error**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| errors.ErrorsErrorResponse | 400, 404                   | application/json           |
| errors.ErrorsErrorResponse | 500                        | application/json           |
| errors.APIError            | 4XX, 5XX                   | \*/\*                      |