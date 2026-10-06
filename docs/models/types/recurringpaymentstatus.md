# RecurringPaymentStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.RecurringPaymentStatusPending

// Open enum: custom values can be created with a direct type cast
custom := types.RecurringPaymentStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `RecurringPaymentStatusPending`   | PENDING                           |
| `RecurringPaymentStatusActive`    | ACTIVE                            |
| `RecurringPaymentStatusPaused`    | PAUSED                            |
| `RecurringPaymentStatusRejected`  | REJECTED                          |
| `RecurringPaymentStatusCancelled` | CANCELLED                         |
| `RecurringPaymentStatusExpired`   | EXPIRED                           |
| `RecurringPaymentStatusUnknown`   | UNKNOWN                           |