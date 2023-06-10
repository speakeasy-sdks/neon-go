<div align="center">
    <img src="https://github.com/speakeasy-sdks/neon-go/assets/6267663/c272d080-d133-43a9-a142-8c9508923854" width="300">
    <h1>Neon Go SDK</h1>
   <p>Fully managed serverless PostgreSQL</p>
   <a href="https://neon.tech/docs/introduction"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/neon-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/neon-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/neon-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/neon-go?sort=semver&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/neon-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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
            BearerAuth: "",
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [APIKey](docs/sdks/apikey/README.md)

* [CreateAPIKey](docs/sdks/apikey/README.md#createapikey) - Create an API key
* [ListAPIKeys](docs/sdks/apikey/README.md#listapikeys) - Get a list of API keys
* [RevokeAPIKey](docs/sdks/apikey/README.md#revokeapikey) - Revoke an API key

### [Branch](docs/sdks/branch/README.md)

* [CreateProjectBranch](docs/sdks/branch/README.md#createprojectbranch) - Create a branch
* [CreateProjectBranchDatabase](docs/sdks/branch/README.md#createprojectbranchdatabase) - Create a database
* [CreateProjectBranchRole](docs/sdks/branch/README.md#createprojectbranchrole) - Create a role
* [DeleteProjectBranch](docs/sdks/branch/README.md#deleteprojectbranch) - Delete a branch
* [DeleteProjectBranchDatabase](docs/sdks/branch/README.md#deleteprojectbranchdatabase) - Delete a database
* [DeleteProjectBranchRole](docs/sdks/branch/README.md#deleteprojectbranchrole) - Delete a role
* [GetProjectBranch](docs/sdks/branch/README.md#getprojectbranch) - Get branch details
* [GetProjectBranchDatabase](docs/sdks/branch/README.md#getprojectbranchdatabase) - Get database details
* [GetProjectBranchRole](docs/sdks/branch/README.md#getprojectbranchrole) - Get role details
* [GetProjectBranchRolePassword](docs/sdks/branch/README.md#getprojectbranchrolepassword) - Get role password
* [ListProjectBranchDatabases](docs/sdks/branch/README.md#listprojectbranchdatabases) - Get a list of databases
* [ListProjectBranchEndpoints](docs/sdks/branch/README.md#listprojectbranchendpoints) - Get a list of branch endpoints
* [ListProjectBranchRoles](docs/sdks/branch/README.md#listprojectbranchroles) - Get a list of roles
* [ListProjectBranches](docs/sdks/branch/README.md#listprojectbranches) - Get a list of branches
* [ResetProjectBranchRolePassword](docs/sdks/branch/README.md#resetprojectbranchrolepassword) - Reset the role password
* [SetPrimaryProjectBranch](docs/sdks/branch/README.md#setprimaryprojectbranch) - Set the branch as the primary branch of a project
* [UpdateProjectBranch](docs/sdks/branch/README.md#updateprojectbranch) - Update a branch
* [UpdateProjectBranchDatabase](docs/sdks/branch/README.md#updateprojectbranchdatabase) - Update a database

### [Endpoint](docs/sdks/endpoint/README.md)

* [CreateProjectEndpoint](docs/sdks/endpoint/README.md#createprojectendpoint) - Create an endpoint
* [DeleteProjectEndpoint](docs/sdks/endpoint/README.md#deleteprojectendpoint) - Delete an endpoint
* [GetProjectEndpoint](docs/sdks/endpoint/README.md#getprojectendpoint) - Get an endpoint
* [ListProjectEndpoints](docs/sdks/endpoint/README.md#listprojectendpoints) - Get a list of endpoints
* [StartProjectEndpoint](docs/sdks/endpoint/README.md#startprojectendpoint) - Start an endpoint
* [SuspendProjectEndpoint](docs/sdks/endpoint/README.md#suspendprojectendpoint) - Suspend an endpoint
* [UpdateProjectEndpoint](docs/sdks/endpoint/README.md#updateprojectendpoint) - Update an endpoint

### [Operation](docs/sdks/operation/README.md)

* [GetProjectOperation](docs/sdks/operation/README.md#getprojectoperation) - Get operation details
* [ListProjectOperations](docs/sdks/operation/README.md#listprojectoperations) - Get a list of operations

### [Preview](docs/sdks/preview/README.md)

* [ListProjectsConsumption](docs/sdks/preview/README.md#listprojectsconsumption) - Get a list of projects consumption

### [Project](docs/sdks/project/README.md)

* [CreateProject](docs/sdks/project/README.md#createproject) - Create a project
* [DeleteProject](docs/sdks/project/README.md#deleteproject) - Delete a project
* [GetProject](docs/sdks/project/README.md#getproject) - Get project details
* [ListProjects](docs/sdks/project/README.md#listprojects) - Get a list of projects
* [UpdateProject](docs/sdks/project/README.md#updateproject) - Update a project
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
