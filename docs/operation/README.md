# Operation

## Overview

These methods allow you to view operation details for your Neon project. For related information, see [Operations](https://neon.tech/docs/manage/operations).

### Available Operations

* [GetProjectOperation](#getprojectoperation) - Get operation details
* [ListProjectOperations](#listprojectoperations) - Get a list of operations

## GetProjectOperation

Retrieves details for the specified operation.
An operation is an action performed on a Neon project resource.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain a `operation_id` by listing operations for the project.


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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Operation.GetProjectOperation(ctx, "aa52c3f5-ad01-49da-9ffe-78f097b0074f", "dicta")
    if err != nil {
        log.Fatal(err)
    }

    if res.OperationResponse != nil {
        // handle response
    }
}
```

## ListProjectOperations

Retrieves a list of operations for the specified Neon project.
You can obtain a `project_id` by listing the projects for your Neon account.
The number of operations returned can be large.
To paginate the response, issue an initial request with a `limit` value.
Then, add the `cursor` value that was returned in the response to the next request.


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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Operation.ListProjectOperations(ctx, "corporis", "dolore", 480894)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListProjectOperations200ApplicationJSONObject != nil {
        // handle response
    }
}
```
