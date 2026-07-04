# GetUsageAnalyticsResponse


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Currency`                                                             | `*string`                                                              | :heavy_minus_sign:                                                     | N/A                                                                    |
| `CustomAnalytics`                                                      | [][types.CustomAnalyticItem](../../models/types/customanalyticitem.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Items`                                                                | [][types.UsageAnalyticItem](../../models/types/usageanalyticitem.md)   | :heavy_minus_sign:                                                     | N/A                                                                    |
| `Subtotal`                                                             | `*string`                                                              | :heavy_minus_sign:                                                     | N/A                                                                    |
| `TotalCost`                                                            | `*string`                                                              | :heavy_minus_sign:                                                     | TotalCost is the final cost after discount (Subtotal - TotalDiscount)  |
| `TotalDiscount`                                                        | `*string`                                                              | :heavy_minus_sign:                                                     | N/A                                                                    |