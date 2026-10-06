# FXRateScope

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.FXRateScopeTenant

// Open enum: custom values can be created with a direct type cast
custom := types.FXRateScope("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `FXRateScopeTenant`       | tenant                    |
| `FXRateScopeCustomer`     | customer                  |
| `FXRateScopeSubscription` | subscription              |