<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	flexprice "github.com/flexprice/go-sdk/v2"
	"github.com/flexprice/go-sdk/v2/models/types"
	"log"
)

func main() {
	ctx := context.Background()

	s := flexprice.New(
		flexprice.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Activity.ListActivity(ctx, types.ListActivityRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListActivityResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->