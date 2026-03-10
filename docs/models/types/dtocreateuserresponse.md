# DtoCreateUserResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Email`                                                             | `*string`                                                           | :heavy_minus_sign:                                                  | Empty for service accounts                                          |
| `ID`                                                                | `*string`                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Password`                                                          | `*string`                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Roles`                                                             | []`string`                                                          | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Tenant`                                                            | [*types.DtoTenantResponse](../../models/types/dtotenantresponse.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Type`                                                              | [*types.UserType](../../models/types/usertype.md)                   | :heavy_minus_sign:                                                  | N/A                                                                 |