# UsageRecords

## Overview

### Available Operations

* [PostUsageRecordsSearch](#postusagerecordssearch) - List usage records

## PostUsageRecordsSearch

Lists usage records. Also accepts filters/sort for a filtered query.

### Example Usage

<!-- UsageSnippet language="go" operationID="post_/usage-records/search" method="post" path="/usage-records/search" -->
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

    res, err := s.UsageRecords.PostUsageRecordsSearch(ctx, types.UsageRecordFilter{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListUsageRecordsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [types.UsageRecordFilter](../../models/types/usagerecordfilter.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][dtos.Option](../../models/dtos/option.md)                       | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*dtos.PostUsageRecordsSearchResponse](../../models/dtos/postusagerecordssearchresponse.md), error**

### Errors

| Error Type           | Status Code          | Content Type         |
| -------------------- | -------------------- | -------------------- |
| errors.ErrorResponse | 400                  | application/json     |
| errors.ErrorResponse | 500                  | application/json     |
| errors.APIError      | 4XX, 5XX             | \*/\*                |