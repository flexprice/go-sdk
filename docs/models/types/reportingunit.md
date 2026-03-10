# ReportingUnit


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ConversionRate`                                                             | `*float64`                                                                   | :heavy_minus_sign:                                                           | Multiplier: reporting_unit_value = unit_value * conversion_rate; must be > 0 |
| `UnitPlural`                                                                 | `*string`                                                                    | :heavy_minus_sign:                                                           | Display unit label, plural (e.g. "seconds")                                  |
| `UnitSingular`                                                               | `*string`                                                                    | :heavy_minus_sign:                                                           | Display unit label, singular (e.g. "second")                                 |