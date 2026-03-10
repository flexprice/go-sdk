# DtoSubscriptionUsageByMetersResponse


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Amount`                                              | `*float64`                                            | :heavy_minus_sign:                                    | N/A                                                   |
| `Currency`                                            | `*string`                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `DisplayAmount`                                       | `*string`                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `FilterValues`                                        | map[string][]`string`                                 | :heavy_minus_sign:                                    | N/A                                                   |
| `IsOverage`                                           | `*bool`                                               | :heavy_minus_sign:                                    | Whether this charge is at overage rate                |
| `MeterDisplayName`                                    | `*string`                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `MeterID`                                             | `*string`                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `OverageFactor`                                       | `*float64`                                            | :heavy_minus_sign:                                    | Factor applied to this charge if in overage           |
| `Price`                                               | [*types.PricePrice](../../models/types/priceprice.md) | :heavy_minus_sign:                                    | N/A                                                   |
| `Quantity`                                            | `*float64`                                            | :heavy_minus_sign:                                    | N/A                                                   |
| `SubscriptionLineItemID`                              | `*string`                                             | :heavy_minus_sign:                                    | For feature_usage: direct match by sub_line_item_id   |