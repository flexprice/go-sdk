# DtoCreateUserRequest


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `Email`                                          | `*string`                                        | :heavy_minus_sign:                               | Required when type is "user"                     |
| `Roles`                                          | []`string`                                       | :heavy_minus_sign:                               | Required when type is "service_account"          |
| `Type`                                           | [types.UserType](../../models/types/usertype.md) | :heavy_check_mark:                               | N/A                                              |