# SystemEntityType

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.SystemEntityTypeFeature

// Open enum: custom values can be created with a direct type cast
custom := types.SystemEntityType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `SystemEntityTypeFeature`           | feature                             |
| `SystemEntityTypeCustomer`          | customer                            |
| `SystemEntityTypePlan`              | plan                                |
| `SystemEntityTypeSubscription`      | subscription                        |
| `SystemEntityTypeInvoice`           | invoice                             |
| `SystemEntityTypePayment`           | payment                             |
| `SystemEntityTypeCreditNote`        | credit_note                         |
| `SystemEntityTypeRefund`            | refund                              |
| `SystemEntityTypeWallet`            | wallet                              |
| `SystemEntityTypeEntitlement`       | entitlement                         |
| `SystemEntityTypeCheckoutSession`   | checkout_session                    |
| `SystemEntityTypePaymentMethod`     | payment_method                      |
| `SystemEntityTypePrice`             | price                               |
| `SystemEntityTypeEvent`             | event                               |
| `SystemEntityTypeCommitment`        | commitment                          |
| `SystemEntityTypeBillingTerm`       | billing_term                        |
| `SystemEntityTypeBillingInstalment` | billing_instalment                  |