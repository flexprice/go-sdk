# CancelSubscriptionScheduleRequest


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ScheduleID`                                                                      | `string`                                                                          | :heavy_check_mark:                                                                | Schedule ID (optional if using request body)                                      |
| `Body`                                                                            | [*types.DtoCancelScheduleRequest](../../models/types/dtocancelschedulerequest.md) | :heavy_minus_sign:                                                                | Cancel request (optional if using path parameter)                                 |