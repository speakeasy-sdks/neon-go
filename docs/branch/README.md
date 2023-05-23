# Branch

## Overview

These methods allow you to create and manage branches in your Neon project. For related information, see [Manage branches](https://neon.tech/docs/manage/branches).

### Available Operations

* [CreateProjectBranch](#createprojectbranch) - Create a branch
* [CreateProjectBranchDatabase](#createprojectbranchdatabase) - Create a database
* [CreateProjectBranchRole](#createprojectbranchrole) - Create a role
* [DeleteProjectBranch](#deleteprojectbranch) - Delete a branch
* [DeleteProjectBranchDatabase](#deleteprojectbranchdatabase) - Delete a database
* [DeleteProjectBranchRole](#deleteprojectbranchrole) - Delete a role
* [GetProjectBranch](#getprojectbranch) - Get branch details
* [GetProjectBranchDatabase](#getprojectbranchdatabase) - Get database details
* [GetProjectBranchRole](#getprojectbranchrole) - Get role details
* [GetProjectBranchRolePassword](#getprojectbranchrolepassword) - Get role password
* [ListProjectBranchDatabases](#listprojectbranchdatabases) - Get a list of databases
* [ListProjectBranchEndpoints](#listprojectbranchendpoints) - Get a list of branch endpoints
* [ListProjectBranchRoles](#listprojectbranchroles) - Get a list of roles
* [ListProjectBranches](#listprojectbranches) - Get a list of branches
* [ResetProjectBranchRolePassword](#resetprojectbranchrolepassword) - Reset the role password
* [SetPrimaryProjectBranch](#setprimaryprojectbranch) - Set the branch as the primary branch of a project
* [UpdateProjectBranch](#updateprojectbranch) - Update a branch
* [UpdateProjectBranchDatabase](#updateprojectbranchdatabase) - Update a database

## CreateProjectBranch

Creates a branch in the specified project.
You can obtain a `project_id` by listing the projects for your Neon account.

This method does not require a request body, but you can specify one to create an endpoint for the branch or to select a non-default parent branch.
The default behavior is to create a branch from the project's root branch (`main`) with no endpoint, and the branch name is auto-generated.
For related information, see [Manage branches](https://neon.tech/docs/manage/branches/).


### Example Usage

```go
package main

import(
	"context"
	"log"
	"neon"
	"neon/pkg/models/operations"
	"neon/pkg/models/shared"
	"neon/pkg/types"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Branch.CreateProjectBranch(ctx, operations.CreateProjectBranchRequest{
        BranchCreateRequest: &shared.BranchCreateRequest{
            Branch: &shared.BranchCreateRequestBranch{
                Name: sdk.String("Stuart Stiedemann"),
                ParentID: sdk.String("vel"),
                ParentLsn: sdk.String("error"),
                ParentTimestamp: types.MustTimeFromString("2022-03-26T09:37:56.283Z"),
            },
            Endpoints: []shared.BranchCreateRequestEndpointOptions{
                shared.BranchCreateRequestEndpointOptions{
                    AutoscalingLimitMaxCu: sdk.Float64(2975.34),
                    AutoscalingLimitMinCu: sdk.Float64(8917.73),
                    Provisioner: shared.ProvisionerK8sPod.ToPointer(),
                    SuspendTimeoutSeconds: sdk.Int64(963663),
                    Type: shared.EndpointTypeReadOnly,
                },
                shared.BranchCreateRequestEndpointOptions{
                    AutoscalingLimitMaxCu: sdk.Float64(3834.41),
                    AutoscalingLimitMinCu: sdk.Float64(4776.65),
                    Provisioner: shared.ProvisionerK8sNeonvm.ToPointer(),
                    SuspendTimeoutSeconds: sdk.Int64(812169),
                    Type: shared.EndpointTypeReadWrite,
                },
            },
        },
        ProjectID: "iusto",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateProjectBranch201ApplicationJSONObject != nil {
        // handle response
    }
}
```

## CreateProjectBranchDatabase

Creates a database in the specified branch.
A branch can have multiple databases.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.
For related information, see [Manage databases](https://neon.tech/docs/manage/databases/).


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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Branch.CreateProjectBranchDatabase(ctx, operations.CreateProjectBranchDatabaseRequest{
        DatabaseCreateRequest: shared.DatabaseCreateRequest{
            Database: shared.DatabaseCreateRequestDatabase{
                Name: "Charlie Walsh II",
                OwnerName: "veritatis",
            },
        },
        BranchID: "deserunt",
        ProjectID: "perferendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatabaseOperations != nil {
        // handle response
    }
}
```

## CreateProjectBranchRole

Creates a role in the specified branch.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.
In Neon, the terms "role" and "user" are synonymous.
For related information, see [Manage users](https://neon.tech/docs/manage/users/).

Connections established to the active read-write endpoint will be dropped.
If the read-write endpoint is idle, the endpoint becomes active for a short period of time and is suspended afterward.


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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Branch.CreateProjectBranchRole(ctx, operations.CreateProjectBranchRoleRequest{
        RoleCreateRequest: shared.RoleCreateRequest{
            Role: shared.RoleCreateRequestRole{
                Name: "Estelle Will",
            },
        },
        BranchID: "at",
        ProjectID: "at",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RoleOperations != nil {
        // handle response
    }
}
```

## DeleteProjectBranch

Deletes the specified branch from a project, and places
all endpoints into an idle state, breaking existing client connections.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain a `branch_id` by listing the project's branches.
For related information, see [Manage branches](https://neon.tech/docs/manage/branches/).

When a successful response status is received, the endpoints are still active,
and the branch is not yet deleted from storage.
The deletion occurs after all operations finish.
You cannot delete a branch if it is the only remaining branch in the project.
A project must have at least one branch.


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
    res, err := s.Branch.DeleteProjectBranch(ctx, operations.DeleteProjectBranchRequest{
        BranchID: "maiores",
        ProjectID: "molestiae",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BranchOperations != nil {
        // handle response
    }
}
```

## DeleteProjectBranchDatabase

Deletes the specified database from the branch.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` and `database_name` by listing branch's databases.
For related information, see [Manage databases](https://neon.tech/docs/manage/databases/).


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
    res, err := s.Branch.DeleteProjectBranchDatabase(ctx, operations.DeleteProjectBranchDatabaseRequest{
        BranchID: "quod",
        DatabaseName: "quod",
        ProjectID: "esse",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatabaseOperations != nil {
        // handle response
    }
}
```

## DeleteProjectBranchRole

Deletes the specified role from the branch.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.
You can obtain the `role_name` by listing the roles for a branch.
In Neon, the terms "role" and "user" are synonymous.
For related information, see [Managing users](https://neon.tech/docs/manage/users/).


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
    res, err := s.Branch.DeleteProjectBranchRole(ctx, operations.DeleteProjectBranchRoleRequest{
        BranchID: "totam",
        ProjectID: "porro",
        RoleName: "dolorum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RoleOperations != nil {
        // handle response
    }
}
```

## GetProjectBranch

Retrieves information about the specified branch.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain a `branch_id` by listing the project's branches.
A `branch_id` value has a `br-` prefix.

Each Neon project has a root branch named `main`.
A project may contain child branches that were branched from `main` or from another branch.
A parent branch is identified by a `parent_id` value, which is the `id` of the parent branch.
For related information, see [Manage branches](https://neon.tech/docs/manage/branches/).


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
    res, err := s.Branch.GetProjectBranch(ctx, operations.GetProjectBranchRequest{
        BranchID: "dicta",
        ProjectID: "nam",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BranchResponse != nil {
        // handle response
    }
}
```

## GetProjectBranchDatabase

Retrieves information about the specified database.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` and `database_name` by listing branch's databases.
For related information, see [Manage databases](https://neon.tech/docs/manage/databases/).


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
    res, err := s.Branch.GetProjectBranchDatabase(ctx, operations.GetProjectBranchDatabaseRequest{
        BranchID: "officia",
        DatabaseName: "occaecati",
        ProjectID: "fugit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatabaseResponse != nil {
        // handle response
    }
}
```

## GetProjectBranchRole

Retrieves details about the specified role.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.
You can obtain the `role_name` by listing the roles for a branch.
In Neon, the terms "role" and "user" are synonymous.
For related information, see [Managing users](https://neon.tech/docs/manage/users/).


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
    res, err := s.Branch.GetProjectBranchRole(ctx, operations.GetProjectBranchRoleRequest{
        BranchID: "deleniti",
        ProjectID: "hic",
        RoleName: "optio",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RoleResponse != nil {
        // handle response
    }
}
```

## GetProjectBranchRolePassword

Retrieves password of the specified role if possible.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.
You can obtain the `role_name` by listing the roles for a branch.
In Neon, the terms "role" and "user" are synonymous.
For related information, see [Managing users](https://neon.tech/docs/manage/users/).


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
    res, err := s.Branch.GetProjectBranchRolePassword(ctx, operations.GetProjectBranchRolePasswordRequest{
        BranchID: "totam",
        ProjectID: "beatae",
        RoleName: "commodi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RolePasswordResponse != nil {
        // handle response
    }
}
```

## ListProjectBranchDatabases

Retrieves a list of databases for the specified branch.
A branch can have multiple databases.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.
For related information, see [Manage databases](https://neon.tech/docs/manage/databases/).


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
    res, err := s.Branch.ListProjectBranchDatabases(ctx, operations.ListProjectBranchDatabasesRequest{
        BranchID: "molestiae",
        ProjectID: "modi",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatabasesResponse != nil {
        // handle response
    }
}
```

## ListProjectBranchEndpoints

Retrieves a list of endpoints for the specified branch.
Currently, Neon permits only one endpoint per branch.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.


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
    res, err := s.Branch.ListProjectBranchEndpoints(ctx, operations.ListProjectBranchEndpointsRequest{
        BranchID: "qui",
        ProjectID: "impedit",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EndpointsResponse != nil {
        // handle response
    }
}
```

## ListProjectBranchRoles

Retrieves a list of roles from the specified branch.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.
In Neon, the terms "role" and "user" are synonymous.
For related information, see [Manage users](https://neon.tech/docs/manage/users/).


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
    res, err := s.Branch.ListProjectBranchRoles(ctx, operations.ListProjectBranchRolesRequest{
        BranchID: "cum",
        ProjectID: "esse",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RolesResponse != nil {
        // handle response
    }
}
```

## ListProjectBranches

Retrieves a list of branches for the specified project.
You can obtain a `project_id` by listing the projects for your Neon account.

Each Neon project has a root branch named `main`.
A `branch_id` value has a `br-` prefix.
A project may contain child branches that were branched from `main` or from another branch.
A parent branch is identified by the `parent_id` value, which is the `id` of the parent branch.
For related information, see [Manage branches](https://neon.tech/docs/manage/branches/).


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
    res, err := s.Branch.ListProjectBranches(ctx, operations.ListProjectBranchesRequest{
        ProjectID: "ipsum",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BranchesResponse != nil {
        // handle response
    }
}
```

## ResetProjectBranchRolePassword

Resets the password for the specified role.
Returns a new password and operations. The new password is ready to use when the last operation finishes.
The old password remains valid until last operation finishes.
Connections to the read-write endpoint are dropped. If idle,
the read-write endpoint becomes active for a short period of time.

You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.
You can obtain the `role_name` by listing the roles for a branch.
In Neon, the terms "role" and "user" are synonymous.
For related information, see [Managing users](https://neon.tech/docs/manage/users/).


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
    res, err := s.Branch.ResetProjectBranchRolePassword(ctx, operations.ResetProjectBranchRolePasswordRequest{
        BranchID: "excepturi",
        ProjectID: "aspernatur",
        RoleName: "perferendis",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RoleOperations != nil {
        // handle response
    }
}
```

## SetPrimaryProjectBranch

The primary mark is automatically removed from the previous primary branch.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.
For more information, see [Manage branches](https://neon.tech/docs/manage/branches/).


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
    res, err := s.Branch.SetPrimaryProjectBranch(ctx, operations.SetPrimaryProjectBranchRequest{
        BranchID: "ad",
        ProjectID: "natus",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BranchOperations != nil {
        // handle response
    }
}
```

## UpdateProjectBranch

Updates the specified branch. Only changing the branch name is supported.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` by listing the project's branches.
For more information, see [Manage branches](https://neon.tech/docs/manage/branches/).


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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Branch.UpdateProjectBranch(ctx, operations.UpdateProjectBranchRequest{
        BranchUpdateRequest: shared.BranchUpdateRequest{
            Branch: shared.BranchUpdateRequestBranch{
                Name: sdk.String("Sheryl Fadel"),
            },
        },
        BranchID: "hic",
        ProjectID: "saepe",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BranchOperations != nil {
        // handle response
    }
}
```

## UpdateProjectBranchDatabase

Updates the specified database in the branch.
You can obtain a `project_id` by listing the projects for your Neon account.
You can obtain the `branch_id` and `database_name` by listing the branch's databases.
For related information, see [Manage databases](https://neon.tech/docs/manage/databases/).


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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Branch.UpdateProjectBranchDatabase(ctx, operations.UpdateProjectBranchDatabaseRequest{
        DatabaseUpdateRequest: shared.DatabaseUpdateRequest{
            Database: shared.DatabaseUpdateRequestDatabase{
                Name: sdk.String("Harvey Hessel"),
                OwnerName: sdk.String("saepe"),
            },
        },
        BranchID: "quidem",
        DatabaseName: "architecto",
        ProjectID: "ipsa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DatabaseOperations != nil {
        // handle response
    }
}
```
