# DtoCreateCouponRequest


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `AmountOff`                                      | **string*                                        | :heavy_minus_sign:                               | N/A                                              |
| `Cadence`                                        | [types.CouponCadence](../types/couponcadence.md) | :heavy_check_mark:                               | N/A                                              |
| `Currency`                                       | **string*                                        | :heavy_minus_sign:                               | N/A                                              |
| `DurationInPeriods`                              | **int64*                                         | :heavy_minus_sign:                               | N/A                                              |
| `MaxRedemptions`                                 | **int64*                                         | :heavy_minus_sign:                               | N/A                                              |
| `Metadata`                                       | map[string]*string*                              | :heavy_minus_sign:                               | N/A                                              |
| `Name`                                           | *string*                                         | :heavy_check_mark:                               | N/A                                              |
| `PercentageOff`                                  | **string*                                        | :heavy_minus_sign:                               | N/A                                              |
| `RedeemAfter`                                    | **string*                                        | :heavy_minus_sign:                               | N/A                                              |
| `RedeemBefore`                                   | **string*                                        | :heavy_minus_sign:                               | N/A                                              |
| `Rules`                                          | map[string]*any*                                 | :heavy_minus_sign:                               | N/A                                              |
| `Type`                                           | [types.CouponType](../types/coupontype.md)       | :heavy_check_mark:                               | N/A                                              |