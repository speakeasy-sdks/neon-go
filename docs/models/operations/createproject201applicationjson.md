# CreateProject201ApplicationJSON

Created a project.
The project includes a connection URI with a database, password, and role.
At least one non-protected role is created with a password.
Wait until the operations are finished before attempting to connect to a project database.



## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Branch`                                                               | [shared.Branch](../../models/shared/branch.md)                         | :heavy_check_mark:                                                     | N/A                                                                    |
| `ConnectionUris`                                                       | [][shared.ConnectionDetails](../../models/shared/connectiondetails.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `Databases`                                                            | [][shared.Database](../../models/shared/database.md)                   | :heavy_check_mark:                                                     | N/A                                                                    |
| `Endpoints`                                                            | [][shared.Endpoint](../../models/shared/endpoint.md)                   | :heavy_check_mark:                                                     | N/A                                                                    |
| `Operations`                                                           | [][shared.Operation](../../models/shared/operation.md)                 | :heavy_check_mark:                                                     | N/A                                                                    |
| `Project`                                                              | [shared.Project](../../models/shared/project.md)                       | :heavy_check_mark:                                                     | N/A                                                                    |
| `Roles`                                                                | [][shared.Role](../../models/shared/role.md)                           | :heavy_check_mark:                                                     | N/A                                                                    |