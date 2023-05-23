// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"neon/pkg/models/shared"
	"net/http"
)

type CreateProjectBranchDatabaseRequest struct {
	DatabaseCreateRequest shared.DatabaseCreateRequest `request:"mediaType=application/json"`
	// The branch ID
	BranchID string `pathParam:"style=simple,explode=false,name=branch_id"`
	// The Neon project ID
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type CreateProjectBranchDatabaseResponse struct {
	ContentType string
	// Created a database in the specified branch
	DatabaseOperations *shared.DatabaseOperations
	// General Error
	GeneralError *shared.GeneralError
	StatusCode   int
	RawResponse  *http.Response
}