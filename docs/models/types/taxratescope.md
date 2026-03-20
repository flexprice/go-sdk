# TaxRateScope

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.TaxRateScopeInternal

// Open enum: custom values can be created with a direct type cast
custom := types.TaxRateScope("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `TaxRateScopeInternal` | INTERNAL               |
| `TaxRateScopeExternal` | EXTERNAL               |
| `TaxRateScopeOnetime`  | ONETIME                |