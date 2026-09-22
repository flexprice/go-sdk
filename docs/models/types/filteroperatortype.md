# FilterOperatorType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.FilterOperatorTypeEq

// Open enum: custom values can be created with a direct type cast
custom := types.FilterOperatorType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `FilterOperatorTypeEq`          | eq                              |
| `FilterOperatorTypeContains`    | contains                        |
| `FilterOperatorTypeNotContains` | not_contains                    |
| `FilterOperatorTypeGt`          | gt                              |
| `FilterOperatorTypeLt`          | lt                              |
| `FilterOperatorTypeGte`         | gte                             |
| `FilterOperatorTypeIn`          | in                              |
| `FilterOperatorTypeNotIn`       | not_in                          |
| `FilterOperatorTypeBefore`      | before                          |
| `FilterOperatorTypeAfter`       | after                           |