# ReminderStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.ReminderStatusPending

// Open enum: custom values can be created with a direct type cast
custom := types.ReminderStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ReminderStatusPending`  | pending                  |
| `ReminderStatusApproved` | approved                 |
| `ReminderStatusHeld`     | held                     |
| `ReminderStatusSkipped`  | skipped                  |