# DtoCreateScheduledTaskRequest


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ConnectionID`                                                       | *string*                                                             | :heavy_check_mark:                                                   | N/A                                                                  |
| `Enabled`                                                            | **bool*                                                              | :heavy_minus_sign:                                                   | N/A                                                                  |
| `EntityType`                                                         | [types.ScheduledTaskEntityType](../types/scheduledtaskentitytype.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `Interval`                                                           | [types.ScheduledTaskInterval](../types/scheduledtaskinterval.md)     | :heavy_check_mark:                                                   | N/A                                                                  |
| `JobConfig`                                                          | [types.S3JobConfig](../types/s3jobconfig.md)                         | :heavy_check_mark:                                                   | N/A                                                                  |