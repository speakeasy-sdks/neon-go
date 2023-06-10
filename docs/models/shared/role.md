# Role


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `BranchID`                                             | *string*                                               | :heavy_check_mark:                                     | The ID of the branch to which the role belongs<br/>    |
| `CreatedAt`                                            | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | A timestamp indicating when the role was created<br/>  |
| `Name`                                                 | *string*                                               | :heavy_check_mark:                                     | The role name<br/>                                     |
| `Password`                                             | **string*                                              | :heavy_minus_sign:                                     | The role password<br/>                                 |
| `Protected`                                            | **bool*                                                | :heavy_minus_sign:                                     | Whether or not the role is system-protected<br/>       |
| `UpdatedAt`                                            | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | A timestamp indicating when the role was last updated<br/> |