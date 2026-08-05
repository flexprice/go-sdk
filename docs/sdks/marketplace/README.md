# Marketplace

## Overview

### Available Operations

* [PostMarketplaceAgreements](#postmarketplaceagreements) - Register an AWS Marketplace agreement

## PostMarketplaceAgreements

Registers an AWS Marketplace buyer agreement against an existing Flexprice subscription, upserting plan/subscription/customer integration mappings in one call.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/marketplace/agreements" method="post" path="/marketplace/agreements" -->
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

    res, err := s.Marketplace.PostMarketplaceAgreements(ctx, types.RegisterMarketplaceAgreementRequest{
        CustomerID: "<id>",
        PlanID: "<id>",
        Provider: types.SecretProviderTabs,
        SubscriptionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.RegisterMarketplaceAgreementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [types.RegisterMarketplaceAgreementRequest](../../models/types/registermarketplaceagreementrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][dtos.Option](../../models/dtos/option.md)                                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*dtos.PostMarketplaceAgreementsResponse](../../models/dtos/postmarketplaceagreementsresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |