# DtoCancelSubscriptionResponse


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `CancellationType`                                           | [*types.CancellationType](../types/cancellationtype.md)      | :heavy_minus_sign:                                           | N/A                                                          |
| `EffectiveDate`                                              | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |
| `Message`                                                    | **string*                                                    | :heavy_minus_sign:                                           | Response metadata                                            |
| `ProcessedAt`                                                | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |
| `ProrationDetails`                                           | [][types.DtoProrationDetail](../types/dtoprorationdetail.md) | :heavy_minus_sign:                                           | N/A                                                          |
| `ProrationInvoice`                                           | [*types.DtoInvoiceResponse](../types/dtoinvoiceresponse.md)  | :heavy_minus_sign:                                           | N/A                                                          |
| `Reason`                                                     | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |
| `Status`                                                     | [*types.SubscriptionStatus](../types/subscriptionstatus.md)  | :heavy_minus_sign:                                           | N/A                                                          |
| `SubscriptionID`                                             | **string*                                                    | :heavy_minus_sign:                                           | Basic cancellation info                                      |
| `TotalCreditAmount`                                          | **string*                                                    | :heavy_minus_sign:                                           | N/A                                                          |