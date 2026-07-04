# CheckoutStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.CheckoutStatusInitiated

// Open enum: custom values can be created with a direct type cast
custom := types.CheckoutStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `CheckoutStatusInitiated` | initiated                 |
| `CheckoutStatusPending`   | pending                   |
| `CheckoutStatusCompleted` | completed                 |
| `CheckoutStatusFailed`    | failed                    |
| `CheckoutStatusExpired`   | expired                   |