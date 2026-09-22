# ColumnType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.ColumnTypeString

// Open enum: custom values can be created with a direct type cast
custom := types.ColumnType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `ColumnTypeString`   | string               |
| `ColumnTypeDecimal`  | decimal              |
| `ColumnTypeDatetime` | datetime             |