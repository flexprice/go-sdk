# PaymentMethodType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.PaymentMethodTypeCard

// Open enum: custom values can be created with a direct type cast
custom := types.PaymentMethodType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `PaymentMethodTypeCard`        | CARD                           |
| `PaymentMethodTypeAch`         | ACH                            |
| `PaymentMethodTypeOffline`     | OFFLINE                        |
| `PaymentMethodTypeCredits`     | CREDITS                        |
| `PaymentMethodTypePaymentLink` | PAYMENT_LINK                   |