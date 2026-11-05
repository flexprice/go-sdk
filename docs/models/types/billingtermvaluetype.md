# BillingTermValueType

fixed or remainder; percentage is reserved and rejected as not supported yet

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.BillingTermValueTypeFixed

// Open enum: custom values can be created with a direct type cast
custom := types.BillingTermValueType("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `BillingTermValueTypeFixed`      | fixed                            |
| `BillingTermValueTypePercentage` | percentage                       |
| `BillingTermValueTypeRemainder`  | remainder                        |