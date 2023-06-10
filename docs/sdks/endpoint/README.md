# Endpoint

## Overview

These methods allow you to create and manage compute endpoints in your Neon project. For related information, see [Manage compute endpoints](https://neon.tech/docs/manage/endpoints).

### Available Operations

* [CreateProjectEndpoint](#createprojectendpoint) - Create an endpoint
* [DeleteProjectEndpoint](#deleteprojectendpoint) - Delete an endpoint
* [GetProjectEndpoint](#getprojectendpoint) - Get an endpoint
* [ListProjectEndpoints](#listprojectendpoints) - Get a list of endpoints
* [StartProjectEndpoint](#startprojectendpoint) - Start an endpoint
* [SuspendProjectEndpoint](#suspendprojectendpoint) - Suspend an endpoint
* [UpdateProjectEndpoint](#updateprojectendpoint) - Update an endpoint

## CreateProjectEndpoint

Creates an endpoint for the specified branch.
An endpoint is a Neon compute instance.
There is a maximum of one endpoint per branch.
If the specified branch already has an endpoint, the operation fails.

You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain `branch_id` by listing the project's branches.
A `branch_id` has a `br-` prefix.
Currently, only the `read_write` endpoint type is supported.
For supported regions and `region_id` values, see [Regions](https://neon.tech/docs/introduction/regions/).
For more information about endpoints, see [Manage endpoints](https://neon.tech/docs/manage/endpoints/).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"neon"
	"neon/pkg/models/operations"
	"neon/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Endpoint.CreateProjectEndpoint(ctx, shared.EndpointCreateRequest{
        Endpoint: shared.EndpointCreateRequestEndpoint{
            AutoscalingLimitMaxCu: sdk.Float64(9698.1),
            AutoscalingLimitMinCu: sdk.Float64(6667.67),
            BranchID: "mollitia",
            Disabled: sdk.Bool(false),
            PasswordlessAccess: sdk.Bool(false),
            PoolerEnabled: sdk.Bool(false),
            PoolerMode: shared.EndpointPoolerModeTransaction.ToPointer(),
            Provisioner: shared.ProvisionerK8sNeonvm.ToPointer(),
            RegionID: sdk.String("dolores"),
            Settings: &shared.EndpointSettingsData{
                PgSettings: map[string]string{
                    "corporis": "explicabo",
                },
            },
            SuspendTimeoutSeconds: sdk.Int64(750686),
            Type: shared.EndpointTypeReadOnly,
        },
    }, "omnis")
    if err != nil {
        log.Fatal(err)
    }

    if res.EndpointOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `endpointCreateRequest`                                                      | [shared.EndpointCreateRequest](../../models/shared/endpointcreaterequest.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `projectID`                                                                  | *string*                                                                     | :heavy_check_mark:                                                           | The Neon project ID                                                          |


### Response

**[*operations.CreateProjectEndpointResponse](../../models/operations/createprojectendpointresponse.md), error**


## DeleteProjectEndpoint

Delete the specified endpoint.
An endpoint is a Neon compute instance.
Deleting an endpoint drops existing network connections to the endpoint.
The deletion is completed when last operation in the chain finishes successfully.

You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain an `endpoint_id` by listing your project's endpoints.
An `endpoint_id` has an `ep-` prefix.
For more information about endpoints, see [Manage endpoints](https://neon.tech/docs/manage/endpoints/).


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
    res, err := s.Endpoint.DeleteProjectEndpoint(ctx, "nemo", "minima")
    if err != nil {
        log.Fatal(err)
    }

    if res.EndpointOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `endpointID`                                          | *string*                                              | :heavy_check_mark:                                    | The endpoint ID                                       |
| `projectID`                                           | *string*                                              | :heavy_check_mark:                                    | The Neon project ID                                   |


### Response

**[*operations.DeleteProjectEndpointResponse](../../models/operations/deleteprojectendpointresponse.md), error**


## GetProjectEndpoint

Retrieves information about the specified endpoint.
An endpoint is a Neon compute instance.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain an `endpoint_id` by listing your project's endpoints.
An `endpoint_id` has an `ep-` prefix.
For more information about endpoints, see [Manage endpoints](https://neon.tech/docs/manage/endpoints/).


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
    res, err := s.Endpoint.GetProjectEndpoint(ctx, "excepturi", "accusantium")
    if err != nil {
        log.Fatal(err)
    }

    if res.EndpointResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `endpointID`                                          | *string*                                              | :heavy_check_mark:                                    | The endpoint ID                                       |
| `projectID`                                           | *string*                                              | :heavy_check_mark:                                    | The Neon project ID                                   |


### Response

**[*operations.GetProjectEndpointResponse](../../models/operations/getprojectendpointresponse.md), error**


## ListProjectEndpoints

Retrieves a list of endpoints for the specified project.
An endpoint is a Neon compute instance.
You can obtain a `project_id` by listing the projects for your Neon account.
For more information about endpoints, see [Manage endpoints](https://neon.tech/docs/manage/endpoints/).


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
    res, err := s.Endpoint.ListProjectEndpoints(ctx, "iure")
    if err != nil {
        log.Fatal(err)
    }

    if res.EndpointsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `projectID`                                           | *string*                                              | :heavy_check_mark:                                    | The Neon project ID                                   |


### Response

**[*operations.ListProjectEndpointsResponse](../../models/operations/listprojectendpointsresponse.md), error**


## StartProjectEndpoint

Starts an endpoint. The endpoint is ready to use
after the last operation in chain finishes successfully.

You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain an `endpoint_id` by listing your project's endpoints.
An `endpoint_id` has an `ep-` prefix.
For more information about endpoints, see [Manage endpoints](https://neon.tech/docs/manage/endpoints/).


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
    res, err := s.Endpoint.StartProjectEndpoint(ctx, "culpa", "doloribus")
    if err != nil {
        log.Fatal(err)
    }

    if res.EndpointOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `endpointID`                                          | *string*                                              | :heavy_check_mark:                                    | The endpoint ID                                       |
| `projectID`                                           | *string*                                              | :heavy_check_mark:                                    | The Neon project ID                                   |


### Response

**[*operations.StartProjectEndpointResponse](../../models/operations/startprojectendpointresponse.md), error**


## SuspendProjectEndpoint

Suspend the specified endpoint
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain an `endpoint_id` by listing your project's endpoints.
An `endpoint_id` has an `ep-` prefix.
For more information about endpoints, see [Manage endpoints](https://neon.tech/docs/manage/endpoints/).


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
    res, err := s.Endpoint.SuspendProjectEndpoint(ctx, "sapiente", "architecto")
    if err != nil {
        log.Fatal(err)
    }

    if res.EndpointOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `endpointID`                                          | *string*                                              | :heavy_check_mark:                                    | The endpoint ID                                       |
| `projectID`                                           | *string*                                              | :heavy_check_mark:                                    | The Neon project ID                                   |


### Response

**[*operations.SuspendProjectEndpointResponse](../../models/operations/suspendprojectendpointresponse.md), error**


## UpdateProjectEndpoint

Updates the specified endpoint. Currently, only changing the associated branch is supported.
The branch that you specify cannot have an existing endpoint.

You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain an `endpoint_id` and `branch_id` by listing your project's endpoints.
An `endpoint_id` has an `ep-` prefix. A `branch_id` has a `br-` prefix.
For more information about endpoints, see [Manage endpoints](https://neon.tech/docs/manage/endpoints/).

If the returned list of operations is not empty, the endpoint is not ready to use.
The client must wait for the last operation to finish before using the endpoint.
If the endpoint was idle before the update, the endpoint becomes active for a short period of time,
and the control plane suspends it again after the update.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"neon"
	"neon/pkg/models/operations"
	"neon/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Endpoint.UpdateProjectEndpoint(ctx, shared.EndpointUpdateRequest{
        Endpoint: shared.EndpointUpdateRequestEndpoint{
            AutoscalingLimitMaxCu: sdk.Float64(6527.9),
            AutoscalingLimitMinCu: sdk.Float64(2088.76),
            BranchID: sdk.String("culpa"),
            Disabled: sdk.Bool(false),
            PasswordlessAccess: sdk.Bool(false),
            PoolerEnabled: sdk.Bool(false),
            PoolerMode: shared.EndpointPoolerModeTransaction.ToPointer(),
            Provisioner: shared.ProvisionerK8sPod.ToPointer(),
            Settings: &shared.EndpointSettingsData{
                PgSettings: map[string]string{
                    "mollitia": "occaecati",
                    "numquam": "commodi",
                    "quam": "molestiae",
                    "velit": "error",
                },
            },
            SuspendTimeoutSeconds: sdk.Int64(158969),
        },
    }, "quis", "vitae")
    if err != nil {
        log.Fatal(err)
    }

    if res.EndpointOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `endpointUpdateRequest`                                                      | [shared.EndpointUpdateRequest](../../models/shared/endpointupdaterequest.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `endpointID`                                                                 | *string*                                                                     | :heavy_check_mark:                                                           | The endpoint ID                                                              |
| `projectID`                                                                  | *string*                                                                     | :heavy_check_mark:                                                           | The Neon project ID                                                          |


### Response

**[*operations.UpdateProjectEndpointResponse](../../models/operations/updateprojectendpointresponse.md), error**

