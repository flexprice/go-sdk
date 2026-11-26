# ~~CreateSubscriptionRequestPaymentTerms~~

Deprecated: use payment_terms_type and payment_terms_days. "<days> NET" (days 1–365), DUE_ON_RECEIPT, END_OF_MONTH or END_OF_NEXT_MONTH.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

## Example Usage

```go
import (
	"github.com/flexprice/go-sdk/v2/models/types"
)

value := types.CreateSubscriptionRequestPaymentTermsFifteenNet
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `CreateSubscriptionRequestPaymentTermsFifteenNet`     | 15 NET                                                |
| `CreateSubscriptionRequestPaymentTermsThirtyNet`      | 30 NET                                                |
| `CreateSubscriptionRequestPaymentTermsFortyFiveNet`   | 45 NET                                                |
| `CreateSubscriptionRequestPaymentTermsSixtyNet`       | 60 NET                                                |
| `CreateSubscriptionRequestPaymentTermsSeventyFiveNet` | 75 NET                                                |
| `CreateSubscriptionRequestPaymentTermsNinetyNet`      | 90 NET                                                |
| `CreateSubscriptionRequestPaymentTermsDueOnReceipt`   | DUE_ON_RECEIPT                                        |
| `CreateSubscriptionRequestPaymentTermsEndOfMonth`     | END_OF_MONTH                                          |
| `CreateSubscriptionRequestPaymentTermsEndOfNextMonth` | END_OF_NEXT_MONTH                                     |