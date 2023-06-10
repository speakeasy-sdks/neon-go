# Project

## Overview

These methods allow you to create and manage Neon projects. For related information, see [Manage projects](https://neon.tech/docs/manage/projects).

### Available Operations

* [CreateProject](#createproject) - Create a project
* [DeleteProject](#deleteproject) - Delete a project
* [GetProject](#getproject) - Get project details
* [ListProjects](#listprojects) - Get a list of projects
* [UpdateProject](#updateproject) - Update a project

## CreateProject

Creates a Neon project.
A project is the top-level object in the Neon object hierarchy.
Plan limits define how many projects you can create.
Neon's Free plan permits one project per Neon account.
For more information, see [Manage projects](https://neon.tech/docs/manage/projects/).

You can specify a region and PostgreSQL version in the request body.
Neon currently supports PostgreSQL 14 and 15.
For supported regions and `region_id` values, see [Regions](https://neon.tech/docs/introduction/regions/).


### Example Usage

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
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Project.CreateProject(ctx, shared.ProjectCreateRequest{
        Project: shared.ProjectCreateRequestProject{
            AutoscalingLimitMaxCu: sdk.Float64(3179.83),
            AutoscalingLimitMinCu: sdk.Float64(8804.76),
            Branch: &shared.ProjectCreateRequestProjectBranch{
                DatabaseName: sdk.String("commodi"),
                Name: sdk.String("Eric Emmerich"),
                RoleName: sdk.String("excepturi"),
            },
            DefaultEndpointSettings: map[string]string{
                "modi": "praesentium",
                "rem": "voluptates",
                "quasi": "repudiandae",
                "sint": "veritatis",
            },
            Name: sdk.String("Miss Randall Hamill"),
            PgVersion: sdk.Int64(131797),
            Provisioner: shared.ProvisionerK8sNeonvm.ToPointer(),
            RegionID: sdk.String("distinctio"),
            Settings: &shared.ProjectSettingsData{
                Quota: &shared.ProjectQuota{
                    ActiveTimeSeconds: sdk.Int64(841386),
                    ComputeTimeSeconds: sdk.Int64(289406),
                    DataTransferBytes: sdk.Int64(264730),
                    LogicalSizeBytes: sdk.Int64(183191),
                    WrittenDataBytes: sdk.Int64(397821),
                },
            },
            StorePasswords: sdk.Bool(false),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateProject201ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.ProjectCreateRequest](../../models/shared/projectcreaterequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.CreateProjectResponse](../../models/operations/createprojectresponse.md), error**


## DeleteProject

Deletes the specified project.
You can obtain a `project_id` by listing the projects for your Neon account.
Deleting a project is a permanent action.
Deleting a project also deletes endpoints, branches, databases, and users that belong to the project.


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
    res, err := s.Project.DeleteProject(ctx, "cupiditate")
    if err != nil {
        log.Fatal(err)
    }

    if res.ProjectResponse != nil {
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

**[*operations.DeleteProjectResponse](../../models/operations/deleteprojectresponse.md), error**


## GetProject

Retrieves information about the specified project.
A project is the top-level object in the Neon object hierarchy.
You can obtain a `project_id` by listing the projects for your Neon account.


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
    res, err := s.Project.GetProject(ctx, "quos")
    if err != nil {
        log.Fatal(err)
    }

    if res.ProjectResponse != nil {
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

**[*operations.GetProjectResponse](../../models/operations/getprojectresponse.md), error**


## ListProjects

Retrieves a list of projects for the Neon account.
A project is the top-level object in the Neon object hierarchy.
For more information, see [Manage projects](https://neon.tech/docs/manage/projects/).


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
    res, err := s.Project.ListProjects(ctx, "perferendis", 164940)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListProjects200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `cursor`                                                                               | **string*                                                                              | :heavy_minus_sign:                                                                     | Specify the cursor value from the previous response to get the next batch of projects. |
| `limit`                                                                                | **int64*                                                                               | :heavy_minus_sign:                                                                     | Specify a value from 1 to 100 to limit number of projects in the response.             |


### Response

**[*operations.ListProjectsResponse](../../models/operations/listprojectsresponse.md), error**


## UpdateProject

Updates the specified project.
You can obtain a `project_id` by listing the projects for your Neon account.
Neon permits updating the project name only.


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
    res, err := s.Project.UpdateProject(ctx, shared.ProjectUpdateRequest{
        Project: shared.ProjectUpdateRequestProject{
            DefaultEndpointSettings: map[string]string{
                "ipsam": "alias",
                "fugit": "dolorum",
                "excepturi": "tempora",
                "facilis": "tempore",
            },
            Name: sdk.String("Lucia Kemmer"),
            Settings: &shared.ProjectSettingsData{
                Quota: &shared.ProjectQuota{
                    ActiveTimeSeconds: sdk.Int64(576157),
                    ComputeTimeSeconds: sdk.Int64(396098),
                    DataTransferBytes: sdk.Int64(592042),
                    LogicalSizeBytes: sdk.Int64(896039),
                    WrittenDataBytes: sdk.Int64(572252),
                },
            },
        },
    }, "officia")
    if err != nil {
        log.Fatal(err)
    }

    if res.ProjectOperations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `projectUpdateRequest`                                                     | [shared.ProjectUpdateRequest](../../models/shared/projectupdaterequest.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `projectID`                                                                | *string*                                                                   | :heavy_check_mark:                                                         | The Neon project ID                                                        |


### Response

**[*operations.UpdateProjectResponse](../../models/operations/updateprojectresponse.md), error**

