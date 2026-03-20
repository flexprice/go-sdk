# InvoiceType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.InvoiceTypeSubscription

// Open enum: custom values can be created with a direct type cast
custom := types.InvoiceType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `InvoiceTypeSubscription` | SUBSCRIPTION              |
| `InvoiceTypeOneOff`       | ONE_OFF                   |
| `InvoiceTypeCredit`       | CREDIT                    |