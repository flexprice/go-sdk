# EntityType

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.EntityTypeEvents

// Open enum: custom values can be created with a direct type cast
custom := types.EntityType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `EntityTypeEvents`    | EVENTS                |
| `EntityTypePrices`    | PRICES                |
| `EntityTypeCustomers` | CUSTOMERS             |
| `EntityTypeFeatures`  | FEATURES              |