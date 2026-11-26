# ~~SubscriptionResponseV2PaymentTerms~~

Deprecated: use payment_terms_type and payment_terms_days. "<days> NET" (days 1–365), DUE_ON_RECEIPT, END_OF_MONTH or END_OF_NEXT_MONTH.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.SubscriptionResponseV2PaymentTermsFifteenNet

// Open enum: custom values can be created with a direct type cast
custom := types.SubscriptionResponseV2PaymentTerms("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `SubscriptionResponseV2PaymentTermsFifteenNet`     | 15 NET                                             |
| `SubscriptionResponseV2PaymentTermsThirtyNet`      | 30 NET                                             |
| `SubscriptionResponseV2PaymentTermsFortyFiveNet`   | 45 NET                                             |
| `SubscriptionResponseV2PaymentTermsSixtyNet`       | 60 NET                                             |
| `SubscriptionResponseV2PaymentTermsSeventyFiveNet` | 75 NET                                             |
| `SubscriptionResponseV2PaymentTermsNinetyNet`      | 90 NET                                             |
| `SubscriptionResponseV2PaymentTermsDueOnReceipt`   | DUE_ON_RECEIPT                                     |
| `SubscriptionResponseV2PaymentTermsEndOfMonth`     | END_OF_MONTH                                       |
| `SubscriptionResponseV2PaymentTermsEndOfNextMonth` | END_OF_NEXT_MONTH                                  |