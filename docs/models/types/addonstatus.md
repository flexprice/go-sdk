# AddonStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.AddonStatusActive

// Open enum: custom values can be created with a direct type cast
custom := types.AddonStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `AddonStatusActive`    | active                 |
| `AddonStatusCancelled` | cancelled              |
| `AddonStatusPending`   | pending                |