// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"neon/pkg/models/shared"
	"net/http"
)

type GetProjectBranchRequest struct {
	// The branch ID
	BranchID string `pathParam:"style=simple,explode=false,name=branch_id"`
	// The Neon project ID
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type GetProjectBranchResponse struct {
	// Returned information about the specified branch
	BranchResponse *shared.BranchResponse
	ContentType    string
	// General Error
	GeneralError *shared.GeneralError
	StatusCode   int
	RawResponse  *http.Response
}
