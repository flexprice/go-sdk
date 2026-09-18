# Features

## Overview

### Available Operations

* [CreateFeature](#createfeature) - Create feature
* [QueryFeature](#queryfeature) - Query features
* [UpdateFeature](#updatefeature) - Update feature
* [DeleteFeature](#deletefeature) - Delete feature
* [CloneFeature](#clonefeature) - Clone a feature

## CreateFeature

Use when defining a new feature or capability to gate or meter (e.g. feature flags or usage-based limits). Ideal for boolean or usage features.

### Example Usage

<!-- UsageSnippet language="go" operationID="createFeature" method="post" path="/features" -->
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

    res, err := s.Features.CreateFeature(ctx, types.CreateFeatureRequest{
        Meter: &types.CreateMeterRequest{
            Aggregation: types.MeterAggregation{},
            EventName: "api_request",
            Name: "API Usage Meter",
            ResetUsage: types.ResetUsageNever,
        },
        Name: "<value>",
        Type: types.FeatureTypeMetered,
    }, dtos.CreateFeatureSecurity{
        Option1: &dtos.CreateFeatureSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FeatureResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [types.CreateFeatureRequest](../../models/types/createfeaturerequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `security`                                                               | [dtos.CreateFeatureSecurity](../../models/dtos/createfeaturesecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |
| `opts`                                                                   | [][dtos.Option](../../models/dtos/option.md)                             | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*dtos.CreateFeatureResponse](../../models/dtos/createfeatureresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryFeature

Use when listing or searching features (e.g. catalog or entitlement setup). Returns a paginated list; supports filtering and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryFeature" method="post" path="/features/search" -->
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

    res, err := s.Features.QueryFeature(ctx, types.FeatureFilter{}, dtos.QueryFeatureSecurity{
        Option1: &dtos.QueryFeatureSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListFeaturesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [types.FeatureFilter](../../models/types/featurefilter.md)             | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `security`                                                             | [dtos.QueryFeatureSecurity](../../models/dtos/queryfeaturesecurity.md) | :heavy_check_mark:                                                     | The security requirements to use for the request.                      |
| `opts`                                                                 | [][dtos.Option](../../models/dtos/option.md)                           | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*dtos.QueryFeatureResponse](../../models/dtos/queryfeatureresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## UpdateFeature

Use when changing feature definition (e.g. name, type, or meter). Request body contains the fields to update.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateFeature" method="put" path="/features/{id}" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Features.UpdateFeature(ctx, dtos.UpdateFeatureSecurity{
        Option1: &dtos.UpdateFeatureSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", types.UpdateFeatureRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.FeatureResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `security`                                                               | [dtos.UpdateFeatureSecurity](../../models/dtos/updatefeaturesecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |
| `id`                                                                     | `string`                                                                 | :heavy_check_mark:                                                       | Feature ID                                                               |
| `body`                                                                   | [types.UpdateFeatureRequest](../../models/types/updatefeaturerequest.md) | :heavy_check_mark:                                                       | Feature update data                                                      |
| `opts`                                                                   | [][dtos.Option](../../models/dtos/option.md)                             | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*dtos.UpdateFeatureResponse](../../models/dtos/updatefeatureresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeleteFeature

Use when retiring a feature (e.g. deprecated capability). Returns 200 with success message.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteFeature" method="delete" path="/features/{id}" -->
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

    res, err := s.Features.DeleteFeature(ctx, dtos.DeleteFeatureSecurity{
        Option1: &dtos.DeleteFeatureSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SuccessResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `security`                                                               | [dtos.DeleteFeatureSecurity](../../models/dtos/deletefeaturesecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |
| `id`                                                                     | `string`                                                                 | :heavy_check_mark:                                                       | Feature ID                                                               |
| `opts`                                                                   | [][dtos.Option](../../models/dtos/option.md)                             | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*dtos.DeleteFeatureResponse](../../models/dtos/deletefeatureresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## CloneFeature

Clone an existing feature

### Example Usage

<!-- UsageSnippet language="go" operationID="cloneFeature" method="post" path="/features/{id}/clone" -->
```go
package main

import(
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := flexprice.New()

    res, err := s.Features.CloneFeature(ctx, dtos.CloneFeatureSecurity{
        Option1: &dtos.CloneFeatureSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>", types.CloneFeatureRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.FeatureResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `security`                                                             | [dtos.CloneFeatureSecurity](../../models/dtos/clonefeaturesecurity.md) | :heavy_check_mark:                                                     | The security requirements to use for the request.                      |
| `id`                                                                   | `string`                                                               | :heavy_check_mark:                                                     | Source Feature ID                                                      |
| `body`                                                                 | [types.CloneFeatureRequest](../../models/types/clonefeaturerequest.md) | :heavy_check_mark:                                                     | Clone configuration                                                    |
| `opts`                                                                 | [][dtos.Option](../../models/dtos/option.md)                           | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*dtos.CloneFeatureResponse](../../models/dtos/clonefeatureresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404, 409        | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |