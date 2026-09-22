# Grain

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.GrainHour

// Open enum: custom values can be created with a direct type cast
custom := types.Grain("custom_value")
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `GrainHour`  | hour         |
| `GrainDay`   | day          |
| `GrainWeek`  | week         |
| `GrainMonth` | month        |