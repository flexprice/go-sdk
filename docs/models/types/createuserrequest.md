# CreateUserRequest


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `Email`                                          | `*string`                                        | :heavy_minus_sign:                               | Required when type is "user"                     |
| `Name`                                           | `*string`                                        | :heavy_minus_sign:                               | Display name; optional for service accounts      |
| `Roles`                                          | []`string`                                       | :heavy_minus_sign:                               | Required when type is "service_account"          |
| `Type`                                           | [types.UserType](../../models/types/usertype.md) | :heavy_check_mark:                               | N/A                                              |