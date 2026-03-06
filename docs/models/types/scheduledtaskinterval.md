# ScheduledTaskInterval

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.ScheduledTaskIntervalFifteenMin

// Open enum: custom values can be created with a direct type cast
custom := types.ScheduledTaskInterval("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ScheduledTaskIntervalFifteenMin` | 15MIN                             |
| `ScheduledTaskIntervalCustom`     | custom                            |
| `ScheduledTaskIntervalHourly`     | hourly                            |
| `ScheduledTaskIntervalDaily`      | daily                             |