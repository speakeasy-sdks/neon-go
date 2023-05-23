<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"neon"
	"neon/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.APIKey.CreateAPIKey(ctx, shared.APIKeyCreateRequest{
        KeyName: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIKeyCreateResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->