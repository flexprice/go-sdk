# DtoSubscriptionChangePreviewResponse

Response showing the financial impact of a subscription plan change


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ChangeType`                                                        | [*types.SubscriptionChangeType](../types/subscriptionchangetype.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `CurrentPlan`                                                       | [*types.DtoPlanSummary](../types/dtoplansummary.md)                 | :heavy_minus_sign:                                                  | N/A                                                                 |
| `EffectiveDate`                                                     | **string*                                                           | :heavy_minus_sign:                                                  | effective_date is when the change would take effect                 |
| `Metadata`                                                          | map[string]*string*                                                 | :heavy_minus_sign:                                                  | metadata from the request                                           |
| `NewBillingCycle`                                                   | [*types.DtoBillingCycleInfo](../types/dtobillingcycleinfo.md)       | :heavy_minus_sign:                                                  | N/A                                                                 |
| `NextInvoicePreview`                                                | [*types.DtoInvoicePreview](../types/dtoinvoicepreview.md)           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `ProrationDetails`                                                  | [*types.DtoProrationDetails](../types/dtoprorationdetails.md)       | :heavy_minus_sign:                                                  | N/A                                                                 |
| `SubscriptionID`                                                    | **string*                                                           | :heavy_minus_sign:                                                  | subscription_id is the ID of the subscription being changed         |
| `TargetPlan`                                                        | [*types.DtoPlanSummary](../types/dtoplansummary.md)                 | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Warnings`                                                          | []*string*                                                          | :heavy_minus_sign:                                                  | warnings contains any warnings about the change                     |