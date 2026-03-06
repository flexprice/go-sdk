# CreditGrantExpiryType

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.CreditGrantExpiryTypeNever

// Open enum: custom values can be created with a direct type cast
custom := types.CreditGrantExpiryType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `CreditGrantExpiryTypeNever`        | NEVER                               |
| `CreditGrantExpiryTypeDuration`     | DURATION                            |
| `CreditGrantExpiryTypeBillingCycle` | BILLING_CYCLE                       |