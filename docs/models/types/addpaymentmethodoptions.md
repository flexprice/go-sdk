# AddPaymentMethodOptions


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `PaymentMethodTypes`                                                                | []`string`                                                                          | :heavy_minus_sign:                                                                  | PaymentMethodTypes are the methods it sets up, e.g. ["enach"] for Razorpay.         |
| `SetsDefault`                                                                       | `*bool`                                                                             | :heavy_minus_sign:                                                                  | SetsDefault reports whether it can make the saved method the default (Stripe only). |