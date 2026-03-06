# FeatureType

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.FeatureTypeMetered

// Open enum: custom values can be created with a direct type cast
custom := types.FeatureType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `FeatureTypeMetered` | metered              |
| `FeatureTypeBoolean` | boolean              |
| `FeatureTypeStatic`  | static               |