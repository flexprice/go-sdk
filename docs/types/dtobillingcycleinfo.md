# DtoBillingCycleInfo


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `BillingAnchor`                                     | **string*                                           | :heavy_minus_sign:                                  | billing_anchor is the new billing anchor            |
| `BillingCadence`                                    | [*types.BillingCadence](../types/billingcadence.md) | :heavy_minus_sign:                                  | N/A                                                 |
| `BillingPeriod`                                     | [*types.BillingPeriod](../types/billingperiod.md)   | :heavy_minus_sign:                                  | N/A                                                 |
| `BillingPeriodCount`                                | **int64*                                            | :heavy_minus_sign:                                  | billing_period_count is the billing period count    |
| `PeriodEnd`                                         | **string*                                           | :heavy_minus_sign:                                  | period_end is the end of the new billing period     |
| `PeriodStart`                                       | **string*                                           | :heavy_minus_sign:                                  | period_start is the start of the new billing period |