# PaymentStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.PaymentStatusInitiated

// Open enum: custom values can be created with a direct type cast
custom := types.PaymentStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `PaymentStatusInitiated`         | INITIATED                        |
| `PaymentStatusPending`           | PENDING                          |
| `PaymentStatusProcessing`        | PROCESSING                       |
| `PaymentStatusSucceeded`         | SUCCEEDED                        |
| `PaymentStatusOverpaid`          | OVERPAID                         |
| `PaymentStatusFailed`            | FAILED                           |
| `PaymentStatusRefunded`          | REFUNDED                         |
| `PaymentStatusPartiallyRefunded` | PARTIALLY_REFUNDED               |