# SubscriptionChangeType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.SubscriptionChangeTypeUpgrade

// Open enum: custom values can be created with a direct type cast
custom := types.SubscriptionChangeType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `SubscriptionChangeTypeUpgrade`   | upgrade                           |
| `SubscriptionChangeTypeDowngrade` | downgrade                         |
| `SubscriptionChangeTypeLateral`   | lateral                           |