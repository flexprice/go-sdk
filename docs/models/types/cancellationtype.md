# CancellationType

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.CancellationTypeImmediate

// Open enum: custom values can be created with a direct type cast
custom := types.CancellationType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `CancellationTypeImmediate`   | immediate                     |
| `CancellationTypeEndOfPeriod` | end_of_period                 |