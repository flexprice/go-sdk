# CheckoutAction

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.CheckoutActionCreateSubscription

// Open enum: custom values can be created with a direct type cast
custom := types.CheckoutAction("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `CheckoutActionCreateSubscription` | create_subscription                |
| `CheckoutActionModifySubscription` | modify_subscription                |
| `CheckoutActionWalletTopup`        | wallet_topup                       |
| `CheckoutActionAddAddon`           | add_addon                          |