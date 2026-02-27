# PriceFilter


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `AllowExpiredPrices`                                    | **bool*                                                 | :heavy_minus_sign:                                      | N/A                                                     |
| `EndTime`                                               | **string*                                               | :heavy_minus_sign:                                      | N/A                                                     |
| `EntityIds`                                             | []*string*                                              | :heavy_minus_sign:                                      | N/A                                                     |
| `EntityType`                                            | [*types.PriceEntityType](../types/priceentitytype.md)   | :heavy_minus_sign:                                      | N/A                                                     |
| `Expand`                                                | **string*                                               | :heavy_minus_sign:                                      | N/A                                                     |
| `Filters`                                               | [][types.FilterCondition](../types/filtercondition.md)  | :heavy_minus_sign:                                      | DSL filters                                             |
| `Limit`                                                 | **int64*                                                | :heavy_minus_sign:                                      | N/A                                                     |
| `MeterIds`                                              | []*string*                                              | :heavy_minus_sign:                                      | N/A                                                     |
| `Offset`                                                | **int64*                                                | :heavy_minus_sign:                                      | N/A                                                     |
| `Order`                                                 | [*types.PriceFilterOrder](../types/pricefilterorder.md) | :heavy_minus_sign:                                      | N/A                                                     |
| `ParentPriceID`                                         | **string*                                               | :heavy_minus_sign:                                      | N/A                                                     |
| `PlanIds`                                               | []*string*                                              | :heavy_minus_sign:                                      | Price override filtering fields                         |
| `PriceIds`                                              | []*string*                                              | :heavy_minus_sign:                                      | N/A                                                     |
| `Sort`                                                  | **string*                                               | :heavy_minus_sign:                                      | N/A                                                     |
| `StartDateLt`                                           | **string*                                               | :heavy_minus_sign:                                      | N/A                                                     |
| `StartTime`                                             | **string*                                               | :heavy_minus_sign:                                      | N/A                                                     |
| `Status`                                                | [*types.Status](../types/status.md)                     | :heavy_minus_sign:                                      | N/A                                                     |
| `SubscriptionID`                                        | **string*                                               | :heavy_minus_sign:                                      | N/A                                                     |