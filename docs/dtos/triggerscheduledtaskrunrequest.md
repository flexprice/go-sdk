# TriggerScheduledTaskRunRequest


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ID`                                                                      | *string*                                                                  | :heavy_check_mark:                                                        | Scheduled Task ID                                                         |
| `Body`                                                                    | [*types.DtoTriggerForceRunRequest](../types/dtotriggerforcerunrequest.md) | :heavy_minus_sign:                                                        | Optional start and end time for custom range                              |