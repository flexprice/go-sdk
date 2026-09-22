# EntityCreationStatus

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.EntityCreationStatusCreated

// Open enum: custom values can be created with a direct type cast
custom := types.EntityCreationStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `EntityCreationStatusCreated`             | created                                   |
| `EntityCreationStatusSuperseded`          | superseded                                |
| `EntityCreationStatusFailedAlreadyExists` | failed_already_exists                     |