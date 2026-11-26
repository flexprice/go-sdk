# ~~SubscriptionResponsePaymentTerms~~

Deprecated: use payment_terms_type and payment_terms_days. "<days> NET" (days 1–365), DUE_ON_RECEIPT, END_OF_MONTH or END_OF_NEXT_MONTH.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.SubscriptionResponsePaymentTermsFifteenNet

// Open enum: custom values can be created with a direct type cast
custom := types.SubscriptionResponsePaymentTerms("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `SubscriptionResponsePaymentTermsFifteenNet`     | 15 NET                                           |
| `SubscriptionResponsePaymentTermsThirtyNet`      | 30 NET                                           |
| `SubscriptionResponsePaymentTermsFortyFiveNet`   | 45 NET                                           |
| `SubscriptionResponsePaymentTermsSixtyNet`       | 60 NET                                           |
| `SubscriptionResponsePaymentTermsSeventyFiveNet` | 75 NET                                           |
| `SubscriptionResponsePaymentTermsNinetyNet`      | 90 NET                                           |
| `SubscriptionResponsePaymentTermsDueOnReceipt`   | DUE_ON_RECEIPT                                   |
| `SubscriptionResponsePaymentTermsEndOfMonth`     | END_OF_MONTH                                     |
| `SubscriptionResponsePaymentTermsEndOfNextMonth` | END_OF_NEXT_MONTH                                |