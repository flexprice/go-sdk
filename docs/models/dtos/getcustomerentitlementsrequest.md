# GetCustomerEntitlementsRequest


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ID`                                                              | `string`                                                          | :heavy_check_mark:                                                | Customer ID                                                       |
| `RefreshGrants`                                                   | `*bool`                                                           | :heavy_minus_sign:                                                | Evaluate entitlement grants before returning their usage (slower) |