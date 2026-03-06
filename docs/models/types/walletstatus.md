# WalletStatus

## Example Usage

```go
import (
	"github.com/flexprice/flexprice-go/v2/models/types"
)

value := types.WalletStatusActive

// Open enum: custom values can be created with a direct type cast
custom := types.WalletStatus("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `WalletStatusActive` | active               |
| `WalletStatusFrozen` | frozen               |
| `WalletStatusClosed` | closed               |