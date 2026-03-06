# DebugTrackerStatus

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.DebugTrackerStatusUnprocessed

// Open enum: custom values can be created with a direct type cast
custom := types.DebugTrackerStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `DebugTrackerStatusUnprocessed` | unprocessed                     |
| `DebugTrackerStatusNotFound`    | not_found                       |
| `DebugTrackerStatusFound`       | found                           |
| `DebugTrackerStatusError`       | error                           |