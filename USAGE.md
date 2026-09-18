<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/dtos"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
	ctx := context.Background()

	s := flexprice.New()

	res, err := s.Addons.CreateAddon(ctx, types.CreateAddonRequest{
		LookupKey: "<value>",
		Name:      "<value>",
	}, dtos.CreateAddonSecurity{
		Option1: &dtos.CreateAddonSecurityOption1{
			APIKeyAuth: "<YOUR_API_KEY_HERE>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateAddonResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->