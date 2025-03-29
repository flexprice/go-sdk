# FlexPrice Go SDK

This is the Go client library for the FlexPrice API.

## Installation

```bash
go get github.com/flexprice/go-sdk
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    flexpriceclient "github.com/flexprice/go-sdk"
)

func main() {
    cfg := flexpriceclient.NewConfiguration()
    cfg.Host = "https://api.flexprice.io"
    
    // Add your API key to the default header
    cfg.DefaultHeader["X-API-Key"] = "YOUR_API_KEY"
    
    client := flexpriceclient.NewAPIClient(cfg)
    
    // Example: Get customer by ID
    customer, resp, err := client.CustomersAPI.CustomersIdGet(context.Background(), "customer-123").Execute()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Customer: %s\n", customer.GetName())
}
```

## Features

- Complete API coverage
- Type-safe client
- Detailed documentation
- Error handling

## Documentation

For detailed API documentation, refer to the code comments and the official FlexPrice API documentation. 