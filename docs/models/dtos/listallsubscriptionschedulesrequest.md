# ListAllSubscriptionSchedulesRequest


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `PendingOnly`                                                | `*bool`                                                      | :heavy_minus_sign:                                           | Filter to pending schedules only                             |
| `SubscriptionID`                                             | `*string`                                                    | :heavy_minus_sign:                                           | Filter by subscription ID                                    |
| `SubscriptionIds`                                            | []`string`                                                   | :heavy_minus_sign:                                           | Filter by subscription IDs                                   |
| `ScheduleType`                                               | [][dtos.ScheduleType](../../models/dtos/scheduletype.md)     | :heavy_minus_sign:                                           | Filter by schedule type                                      |
| `ScheduleStatus`                                             | [][dtos.ScheduleStatus](../../models/dtos/schedulestatus.md) | :heavy_minus_sign:                                           | Filter by schedule status                                    |
| `Limit`                                                      | `*int64`                                                     | :heavy_minus_sign:                                           | Limit results                                                |
| `Offset`                                                     | `*int64`                                                     | :heavy_minus_sign:                                           | Offset for pagination                                        |