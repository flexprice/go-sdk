# AlertType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.AlertTypeLowOngoingBalance

// Open enum: custom values can be created with a direct type cast
custom := types.AlertType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `AlertTypeLowOngoingBalance`    | low_ongoing_balance             |
| `AlertTypeLowCreditBalance`     | low_credit_balance              |
| `AlertTypeFeatureWalletBalance` | feature_wallet_balance          |