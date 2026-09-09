# RefundReason

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.RefundReasonDuplicate

// Open enum: custom values can be created with a direct type cast
custom := types.RefundReason("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `RefundReasonDuplicate`           | DUPLICATE                         |
| `RefundReasonFraudulent`          | FRAUDULENT                        |
| `RefundReasonRequestedByCustomer` | REQUESTED_BY_CUSTOMER             |
| `RefundReasonOrderChange`         | ORDER_CHANGE                      |
| `RefundReasonServiceIssue`        | SERVICE_ISSUE                     |
| `RefundReasonOther`               | OTHER                             |