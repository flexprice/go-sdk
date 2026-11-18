# IntegrationCapabilityType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.IntegrationCapabilityTypeCheckout

// Open enum: custom values can be created with a direct type cast
custom := types.IntegrationCapabilityType("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `IntegrationCapabilityTypeCheckout`                | checkout                                           |
| `IntegrationCapabilityTypeAutoCharge`              | auto_charge                                        |
| `IntegrationCapabilityTypeSetDefaultMethod`        | set_default_method                                 |
| `IntegrationCapabilityTypePaymentLink`             | payment_link                                       |
| `IntegrationCapabilityTypePaymentMethodManagement` | payment_method_management                          |
| `IntegrationCapabilityTypeInvoiceSync`             | invoice_sync                                       |
| `IntegrationCapabilityTypeListPaymentMethods`      | list_payment_methods                               |
| `IntegrationCapabilityTypeAddPaymentMethod`        | add_payment_method                                 |
| `IntegrationCapabilityTypeDeletePaymentMethod`     | delete_payment_method                              |