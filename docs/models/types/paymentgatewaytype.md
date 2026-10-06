# PaymentGatewayType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.PaymentGatewayTypeStripe

// Open enum: custom values can be created with a direct type cast
custom := types.PaymentGatewayType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `PaymentGatewayTypeStripe`    | stripe                        |
| `PaymentGatewayTypeRazorpay`  | razorpay                      |
| `PaymentGatewayTypeNomod`     | nomod                         |
| `PaymentGatewayTypeMoyasar`   | moyasar                       |
| `PaymentGatewayTypePaddle`    | paddle                        |
| `PaymentGatewayTypeWhop`      | whop                          |
| `PaymentGatewayTypeChargebee` | chargebee                     |