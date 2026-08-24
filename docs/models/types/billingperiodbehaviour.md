# BillingPeriodBehaviour

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.BillingPeriodBehaviourUnchanged

// Open enum: custom values can be created with a direct type cast
custom := types.BillingPeriodBehaviour("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `BillingPeriodBehaviourUnchanged`      | unchanged                              |
| `BillingPeriodBehaviourAnchorAtEffect` | anchor_at_effect                       |
| `BillingPeriodBehaviourAnchorAtConfig` | anchor_at_config                       |