<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	flexprice "github.com/flexprice/flexprice-go/v2"
	"github.com/flexprice/flexprice-go/v2/types"
	"log"
)

func main() {
	ctx := context.Background()

	s := flexprice.New(
		"https://api.example.com",
		flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Addons.CreateAddon(ctx, types.DtoCreateAddonRequest{
		LookupKey: "<value>",
		Name:      "<value>",
		Type:      types.AddonTypeMultipleInstance,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.DtoCreateAddonResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->