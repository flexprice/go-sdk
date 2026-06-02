# SubModifyTrialEndRequest


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Action`                                                                         | [types.TrialEndAction](../../models/types/trialendaction.md)                     | :heavy_check_mark:                                                               | N/A                                                                              |
| `NewTrialEnd`                                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                       | :heavy_minus_sign:                                                               | NewTrialEnd is the new trial end date. Required when action is "scheduled_date". |