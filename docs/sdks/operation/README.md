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
            BearerAuth: "",
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `operationID`                                         | *string*                                              | :heavy_check_mark:                                    | The operation ID                                      |
| `projectID`                                           | *string*                                              | :heavy_check_mark:                                    | The Neon project ID                                   |


### Response

**[*operations.GetProjectOperationResponse](../../models/operations/getprojectoperationresponse.md), error**


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
            BearerAuth: "",
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

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `projectID`                                                                             | *string*                                                                                | :heavy_check_mark:                                                                      | The Neon project ID                                                                     |
| `cursor`                                                                                | **string*                                                                               | :heavy_minus_sign:                                                                      | Specify the cursor value from the previous response to get the next batch of operations |
| `limit`                                                                                 | **int64*                                                                                | :heavy_minus_sign:                                                                      | Specify a value from 1 to 1000 to limit number of operations in the response            |


### Response

**[*operations.ListProjectOperationsResponse](../../models/operations/listprojectoperationsresponse.md), error**

