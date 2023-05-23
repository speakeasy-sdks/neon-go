// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"neon/pkg/models/shared"
	"net/http"
)

type DeleteProjectBranchRoleRequest struct {
	// The branch ID
	BranchID string `pathParam:"style=simple,explode=false,name=branch_id"`
	// The Neon project ID
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
	// The role name
	RoleName string `pathParam:"style=simple,explode=false,name=role_name"`
}

type DeleteProjectBranchRoleResponse struct {
	ContentType string
	// General Error
	GeneralError *shared.GeneralError
	// Deleted the specified role from the branch
	RoleOperations *shared.RoleOperations
	StatusCode     int
	RawResponse    *http.Response
}
