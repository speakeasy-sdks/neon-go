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

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `cursor`                                                                               | **string*                                                                              | :heavy_minus_sign:                                                                     | Specify the cursor value from the previous response to get the next batch of projects. |
| `limit`                                                                                | **int64*                                                                               | :heavy_minus_sign:                                                                     | Specify a value from 1 to 1000 to limit number of projects in the response.            |


### Response

**[*operations.ListProjectsConsumptionResponse](../../models/operations/listprojectsconsumptionresponse.md), error**

