# GetCustomerEntitlementsByExternalIDRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ExternalID`                                                      | `string`                                                          | :heavy_check_mark:                                                | Customer External ID                                              |
| `RefreshGrants`                                                   | `*bool`                                                           | :heavy_minus_sign:                                                | Evaluate entitlement grants before returning their usage (slower) |