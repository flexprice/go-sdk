# PaymentMethodStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.PaymentMethodStatusActive

// Open enum: custom values can be created with a direct type cast
custom := types.PaymentMethodStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `PaymentMethodStatusActive`   | ACTIVE                        |
| `PaymentMethodStatusInactive` | INACTIVE                      |
| `PaymentMethodStatusExpired`  | EXPIRED                       |