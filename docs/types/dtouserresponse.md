# DtoUserResponse


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Email`                                                   | **string*                                                 | :heavy_minus_sign:                                        | Empty for service accounts                                |
| `ID`                                                      | **string*                                                 | :heavy_minus_sign:                                        | N/A                                                       |
| `Roles`                                                   | []*string*                                                | :heavy_minus_sign:                                        | N/A                                                       |
| `Tenant`                                                  | [*types.DtoTenantResponse](../types/dtotenantresponse.md) | :heavy_minus_sign:                                        | N/A                                                       |
| `Type`                                                    | [*types.UserType](../types/usertype.md)                   | :heavy_minus_sign:                                        | N/A                                                       |