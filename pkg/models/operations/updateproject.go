// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"neon/pkg/models/shared"
	"net/http"
)

type UpdateProjectRequest struct {
	ProjectUpdateRequest shared.ProjectUpdateRequest `request:"mediaType=application/json"`
	// The Neon project ID
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type UpdateProjectResponse struct {
	ContentType string
	// General Error
	GeneralError *shared.GeneralError
	// Updated the specified project
	ProjectOperations *shared.ProjectOperations
	StatusCode        int
	RawResponse       *http.Response
}
