# Groups

## Overview

### Available Operations

* [CreateGroup](#creategroup) - Create group
* [QueryGroup](#querygroup) - Query groups
* [GetGroup](#getgroup) - Get group
* [DeleteGroup](#deletegroup) - Delete group

## CreateGroup

Use when organizing entities into a group (e.g. for filtering prices or plans by product line or region).

### Example Usage

<!-- UsageSnippet language="go" operationID="createGroup" method="post" path="/groups" -->
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

    res, err := s.Groups.CreateGroup(ctx, types.CreateGroupRequest{
        EntityType: "<value>",
        LookupKey: "<value>",
        Name: "<value>",
    }, dtos.CreateGroupSecurity{
        Option1: &dtos.CreateGroupSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GroupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [types.CreateGroupRequest](../../models/types/creategrouprequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `security`                                                           | [dtos.CreateGroupSecurity](../../models/dtos/creategroupsecurity.md) | :heavy_check_mark:                                                   | The security requirements to use for the request.                    |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.CreateGroupResponse](../../models/dtos/creategroupresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## QueryGroup

Use when listing or searching groups (e.g. admin catalog). Returns a paginated list; supports filtering and sorting.

### Example Usage

<!-- UsageSnippet language="go" operationID="queryGroup" method="post" path="/groups/search" -->
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

    res, err := s.Groups.QueryGroup(ctx, types.GroupFilter{}, dtos.QueryGroupSecurity{
        Option1: &dtos.QueryGroupSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListGroupsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [types.GroupFilter](../../models/types/groupfilter.md)             | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `security`                                                         | [dtos.QueryGroupSecurity](../../models/dtos/querygroupsecurity.md) | :heavy_check_mark:                                                 | The security requirements to use for the request.                  |
| `opts`                                                             | [][dtos.Option](../../models/dtos/option.md)                       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*dtos.QueryGroupResponse](../../models/dtos/querygroupresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetGroup

Use when you need to load a single group (e.g. for display or to assign entities).

### Example Usage

<!-- UsageSnippet language="go" operationID="getGroup" method="get" path="/groups/{id}" -->
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

    res, err := s.Groups.GetGroup(ctx, dtos.GetGroupSecurity{
        Option1: &dtos.GetGroupSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.GroupResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `security`                                                     | [dtos.GetGroupSecurity](../../models/dtos/getgroupsecurity.md) | :heavy_check_mark:                                             | The security requirements to use for the request.              |
| `id`                                                           | `string`                                                       | :heavy_check_mark:                                             | Group ID                                                       |
| `opts`                                                         | [][dtos.Option](../../models/dtos/option.md)                   | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*dtos.GetGroupResponse](../../models/dtos/getgroupresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## DeleteGroup

Use when removing a group and clearing its entity associations (e.g. retiring a product line). Returns 204 or 200 on success.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteGroup" method="delete" path="/groups/{id}" -->
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

    res, err := s.Groups.DeleteGroup(ctx, dtos.DeleteGroupSecurity{
        Option1: &dtos.DeleteGroupSecurityOption1{
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `security`                                                           | [dtos.DeleteGroupSecurity](../../models/dtos/deletegroupsecurity.md) | :heavy_check_mark:                                                   | The security requirements to use for the request.                    |
| `id`                                                                 | `string`                                                             | :heavy_check_mark:                                                   | Group ID                                                             |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.DeleteGroupResponse](../../models/dtos/deletegroupresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400, 404             | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |