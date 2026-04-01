# ScheduledTaskInterval

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.ScheduledTaskIntervalFifteenMin

// Open enum: custom values can be created with a direct type cast
custom := types.ScheduledTaskInterval("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ScheduledTaskIntervalFifteenMin` | 15MIN                             |
| `ScheduledTaskIntervalThirtyMin`  | 30MIN                             |
| `ScheduledTaskIntervalCustom`     | custom                            |
| `ScheduledTaskIntervalHourly`     | hourly                            |
| `ScheduledTaskIntervalDaily`      | daily                             |