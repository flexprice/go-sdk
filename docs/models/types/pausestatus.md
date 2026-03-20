# PauseStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.PauseStatusNone

// Open enum: custom values can be created with a direct type cast
custom := types.PauseStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `PauseStatusNone`      | none                   |
| `PauseStatusActive`    | active                 |
| `PauseStatusScheduled` | scheduled              |
| `PauseStatusCompleted` | completed              |
| `PauseStatusCancelled` | cancelled              |