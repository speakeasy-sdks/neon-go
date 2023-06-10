# Database


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `BranchID`                                                 | *string*                                                   | :heavy_check_mark:                                         | The ID of the branch to which the database belongs<br/>    |
| `CreatedAt`                                                | [time.Time](https://pkg.go.dev/time#Time)                  | :heavy_check_mark:                                         | A timestamp indicating when the database was created<br/>  |
| `ID`                                                       | *int64*                                                    | :heavy_check_mark:                                         | The database ID<br/>                                       |
| `Name`                                                     | *string*                                                   | :heavy_check_mark:                                         | The database name<br/>                                     |
| `OwnerName`                                                | *string*                                                   | :heavy_check_mark:                                         | The name of role that owns the database<br/>               |
| `UpdatedAt`                                                | [time.Time](https://pkg.go.dev/time#Time)                  | :heavy_check_mark:                                         | A timestamp indicating when the database was last updated<br/> |