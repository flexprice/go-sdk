# SubscriptionStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.SubscriptionStatusActive

// Open enum: custom values can be created with a direct type cast
custom := types.SubscriptionStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `SubscriptionStatusActive`     | active                         |
| `SubscriptionStatusPaused`     | paused                         |
| `SubscriptionStatusCancelled`  | cancelled                      |
| `SubscriptionStatusIncomplete` | incomplete                     |
| `SubscriptionStatusTrialing`   | trialing                       |
| `SubscriptionStatusDraft`      | draft                          |