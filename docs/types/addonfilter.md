# AddonFilter


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `AddonIds`                                                | []*string*                                                | :heavy_minus_sign:                                        | N/A                                                       |
| `AddonType`                                               | [*types.AddonType](../types/addontype.md)                 | :heavy_minus_sign:                                        | N/A                                                       |
| `EndTime`                                                 | **string*                                                 | :heavy_minus_sign:                                        | N/A                                                       |
| `Expand`                                                  | **string*                                                 | :heavy_minus_sign:                                        | N/A                                                       |
| `Filters`                                                 | [][types.FilterCondition](../types/filtercondition.md)    | :heavy_minus_sign:                                        | filters allows complex filtering based on multiple fields |
| `Limit`                                                   | **int64*                                                  | :heavy_minus_sign:                                        | N/A                                                       |
| `LookupKeys`                                              | []*string*                                                | :heavy_minus_sign:                                        | N/A                                                       |
| `Offset`                                                  | **int64*                                                  | :heavy_minus_sign:                                        | N/A                                                       |
| `Order`                                                   | [*types.AddonFilterOrder](../types/addonfilterorder.md)   | :heavy_minus_sign:                                        | N/A                                                       |
| `Sort`                                                    | [][types.SortCondition](../types/sortcondition.md)        | :heavy_minus_sign:                                        | N/A                                                       |
| `StartTime`                                               | **string*                                                 | :heavy_minus_sign:                                        | N/A                                                       |
| `Status`                                                  | [*types.Status](../types/status.md)                       | :heavy_minus_sign:                                        | N/A                                                       |