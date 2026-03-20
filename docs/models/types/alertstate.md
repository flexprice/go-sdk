# AlertState

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.AlertStateOk

// Open enum: custom values can be created with a direct type cast
custom := types.AlertState("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `AlertStateOk`      | ok                  |
| `AlertStateInfo`    | info                |
| `AlertStateWarning` | warning             |
| `AlertStateInAlarm` | in_alarm            |