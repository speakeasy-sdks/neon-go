// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"neon/pkg/models/shared"
	"net/http"
)

type GetProjectOperationRequest struct {
	// The operation ID
	OperationID string `pathParam:"style=simple,explode=false,name=operation_id"`
	// The Neon project ID
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type GetProjectOperationResponse struct {
	ContentType string
	// General Error
	GeneralError *shared.GeneralError
	// Returned details for the specified operation
	OperationResponse *shared.OperationResponse
	StatusCode        int
	RawResponse       *http.Response
}