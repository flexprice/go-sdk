# TriggerScheduledTaskRunRequest


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ID`                                                                          | `string`                                                                      | :heavy_check_mark:                                                            | Scheduled Task ID                                                             |
| `Body`                                                                        | [*types.TriggerForceRunRequest](../../models/types/triggerforcerunrequest.md) | :heavy_minus_sign:                                                            | Optional start and end time for custom range                                  |