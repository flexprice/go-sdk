# TransactionStatus

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.TransactionStatusPending

// Open enum: custom values can be created with a direct type cast
custom := types.TransactionStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `TransactionStatusPending`   | pending                      |
| `TransactionStatusCompleted` | completed                    |
| `TransactionStatusFailed`    | failed                       |