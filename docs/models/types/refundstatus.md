# RefundStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.RefundStatusPending

// Open enum: custom values can be created with a direct type cast
custom := types.RefundStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `RefundStatusPending`    | PENDING                  |
| `RefundStatusProcessing` | PROCESSING               |
| `RefundStatusSucceeded`  | SUCCEEDED                |
| `RefundStatusFailed`     | FAILED                   |
| `RefundStatusCancelled`  | CANCELLED                |