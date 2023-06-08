# Preview

## Overview

New API methods that are in Beta / Preview state and could be subjected to significant changes in the future.

### Available Operations

* [ListProjectsConsumption](#listprojectsconsumption) - Get a list of projects consumption

## ListProjectsConsumption

Note, this is a preview API and could be subjected to significant changes in the future.
Retrieves a list of per-project consumption for the current billing period.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"neon"
	"neon/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Preview.ListProjectsConsumption(ctx, "dicta", 688661)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListProjectsConsumption200ApplicationJSONObject != nil {
        // handle response
    }
}
```
