# ZohoPaymentGateway

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.ZohoPaymentGatewayPaypal

// Open enum: custom values can be created with a direct type cast
custom := types.ZohoPaymentGateway("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ZohoPaymentGatewayPaypal`       | paypal                           |
| `ZohoPaymentGatewayStripe`       | stripe                           |
| `ZohoPaymentGatewayAuthorizeNet` | authorize_net                    |
| `ZohoPaymentGatewayPayflowPro`   | payflow_pro                      |
| `ZohoPaymentGatewayTwocheckout`  | 2checkout                        |
| `ZohoPaymentGatewayBraintree`    | braintree                        |
| `ZohoPaymentGatewayRazorpay`     | razorpay                         |