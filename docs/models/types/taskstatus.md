# TaskStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.TaskStatusPending

// Open enum: custom values can be created with a direct type cast
custom := types.TaskStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `TaskStatusPending`    | PENDING                |
| `TaskStatusProcessing` | PROCESSING             |
| `TaskStatusCompleted`  | COMPLETED              |
| `TaskStatusFailed`     | FAILED                 |