// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"neon/pkg/models/shared"
	"net/http"
)

type SuspendProjectEndpointRequest struct {
	// The endpoint ID
	EndpointID string `pathParam:"style=simple,explode=false,name=endpoint_id"`
	// The Neon project ID
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type SuspendProjectEndpointResponse struct {
	ContentType string
	// Suspended the specified endpoint
	EndpointOperations *shared.EndpointOperations
	// General Error
	GeneralError *shared.GeneralError
	StatusCode   int
	RawResponse  *http.Response
}
