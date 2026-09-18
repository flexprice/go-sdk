# Rbac

## Overview

### Available Operations

* [ListRbacRoles](#listrbacroles) - List all RBAC roles
* [GetRbacRole](#getrbacrole) - Get a specific RBAC role

## ListRbacRoles

Use when building role pickers or permission UIs. Returns all roles with permissions and descriptions.

### Example Usage

<!-- UsageSnippet language="go" operationID="listRbacRoles" method="get" path="/rbac/roles" -->
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

    res, err := s.Rbac.ListRbacRoles(ctx, dtos.ListRbacRolesSecurity{
        Option1: &dtos.ListRbacRolesSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `security`                                                               | [dtos.ListRbacRolesSecurity](../../models/dtos/listrbacrolessecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |
| `userType`                                                               | [*dtos.UserType](../../models/dtos/usertype.md)                          | :heavy_minus_sign:                                                       | Filter by user type                                                      |
| `opts`                                                                   | [][dtos.Option](../../models/dtos/option.md)                             | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*dtos.ListRbacRolesResponse](../../models/dtos/listrbacrolesresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |

## GetRbacRole

Use when you need to show or edit a single role (e.g. role detail page). Includes permissions, name, and description.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRbacRole" method="get" path="/rbac/roles/{id}" -->
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

    res, err := s.Rbac.GetRbacRole(ctx, dtos.GetRbacRoleSecurity{
        Option1: &dtos.GetRbacRoleSecurityOption1{
            APIKeyAuth: "<YOUR_API_KEY_HERE>",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `security`                                                           | [dtos.GetRbacRoleSecurity](../../models/dtos/getrbacrolesecurity.md) | :heavy_check_mark:                                                   | The security requirements to use for the request.                    |
| `id`                                                                 | `string`                                                             | :heavy_check_mark:                                                   | Role ID                                                              |
| `opts`                                                               | [][dtos.Option](../../models/dtos/option.md)                         | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*dtos.GetRbacRoleResponse](../../models/dtos/getrbacroleresponse.md), error**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| errors.APIError | 4XX, 5XX        | \*/\*           |