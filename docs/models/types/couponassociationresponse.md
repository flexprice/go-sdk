# CouponAssociationResponse


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Coupon`                                      | [*types.Coupon](../../models/types/coupon.md) | :heavy_minus_sign:                            | N/A                                           |
| `CouponID`                                    | `*string`                                     | :heavy_minus_sign:                            | N/A                                           |
| `CreatedAt`                                   | [*time.Time](https://pkg.go.dev/time#Time)    | :heavy_minus_sign:                            | N/A                                           |
| `CreatedBy`                                   | `*string`                                     | :heavy_minus_sign:                            | N/A                                           |
| `EndDate`                                     | [*time.Time](https://pkg.go.dev/time#Time)    | :heavy_minus_sign:                            | Optional                                      |
| `EnvironmentID`                               | `*string`                                     | :heavy_minus_sign:                            | N/A                                           |
| `ID`                                          | `*string`                                     | :heavy_minus_sign:                            | N/A                                           |
| `Metadata`                                    | map[string]`string`                           | :heavy_minus_sign:                            | N/A                                           |
| `StartDate`                                   | [*time.Time](https://pkg.go.dev/time#Time)    | :heavy_minus_sign:                            | N/A                                           |
| `Status`                                      | [*types.Status](../../models/types/status.md) | :heavy_minus_sign:                            | N/A                                           |
| `SubscriptionID`                              | `*string`                                     | :heavy_minus_sign:                            | Mandatory                                     |
| `SubscriptionLineItemID`                      | `*string`                                     | :heavy_minus_sign:                            | Optional                                      |
| `SubscriptionPhaseID`                         | `*string`                                     | :heavy_minus_sign:                            | Optional                                      |
| `TenantID`                                    | `*string`                                     | :heavy_minus_sign:                            | N/A                                           |
| `UpdatedAt`                                   | [*time.Time](https://pkg.go.dev/time#Time)    | :heavy_minus_sign:                            | N/A                                           |
| `UpdatedBy`                                   | `*string`                                     | :heavy_minus_sign:                            | N/A                                           |