# Status

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.StatusPublished

// Open enum: custom values can be created with a direct type cast
custom := types.Status("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `StatusPublished` | published         |
| `StatusDeleted`   | deleted           |
| `StatusArchived`  | archived          |