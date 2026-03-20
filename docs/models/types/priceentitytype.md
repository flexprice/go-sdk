# PriceEntityType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.PriceEntityTypePlan

// Open enum: custom values can be created with a direct type cast
custom := types.PriceEntityType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `PriceEntityTypePlan`         | PLAN                          |
| `PriceEntityTypeSubscription` | SUBSCRIPTION                  |
| `PriceEntityTypeAddon`        | ADDON                         |
| `PriceEntityTypePrice`        | PRICE                         |
| `PriceEntityTypeCostsheet`    | COSTSHEET                     |