# ScheduleStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.ScheduleStatusPending

// Open enum: custom values can be created with a direct type cast
custom := types.ScheduleStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ScheduleStatusPending`   | pending                   |
| `ScheduleStatusExecuting` | executing                 |
| `ScheduleStatusExecuted`  | executed                  |
| `ScheduleStatusCancelled` | cancelled                 |
| `ScheduleStatusFailed`    | failed                    |