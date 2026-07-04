# PaymentActionType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.PaymentActionTypeCheckoutURL

// Open enum: custom values can be created with a direct type cast
custom := types.PaymentActionType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `PaymentActionTypeCheckoutURL` | checkout_url                   |
| `PaymentActionTypePaymentLink` | payment_link                   |