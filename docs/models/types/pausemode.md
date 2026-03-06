# PauseMode

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.PauseModeImmediate

// Open enum: custom values can be created with a direct type cast
custom := types.PauseMode("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `PauseModeImmediate` | immediate            |
| `PauseModeScheduled` | scheduled            |
| `PauseModePeriodEnd` | period_end           |