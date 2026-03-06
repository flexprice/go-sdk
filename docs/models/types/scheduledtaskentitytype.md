# ScheduledTaskEntityType

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.ScheduledTaskEntityTypeEvents

// Open enum: custom values can be created with a direct type cast
custom := types.ScheduledTaskEntityType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ScheduledTaskEntityTypeEvents`       | events                                |
| `ScheduledTaskEntityTypeInvoice`      | invoice                               |
| `ScheduledTaskEntityTypeCreditTopups` | credit_topups                         |
| `ScheduledTaskEntityTypeCreditUsage`  | credit_usage                          |