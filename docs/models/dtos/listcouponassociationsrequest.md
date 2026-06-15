# ListCouponAssociationsRequest


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `SubscriptionIds`                         | []`string`                                | :heavy_minus_sign:                        | Filter by subscription IDs (max 100)      |
| `CouponIds`                               | []`string`                                | :heavy_minus_sign:                        | Filter by coupon IDs (max 100)            |
| `ActiveOnly`                              | `*bool`                                   | :heavy_minus_sign:                        | Return only currently active associations |
| `Limit`                                   | `*int64`                                  | :heavy_minus_sign:                        | Page size                                 |
| `Offset`                                  | `*int64`                                  | :heavy_minus_sign:                        | Page offset                               |