// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"neon/pkg/models/shared"
	"net/http"
)

type CreateProjectEndpointRequest struct {
	EndpointCreateRequest shared.EndpointCreateRequest `request:"mediaType=application/json"`
	// The Neon project ID
	ProjectID string `pathParam:"style=simple,explode=false,name=project_id"`
}

type CreateProjectEndpointResponse struct {
	ContentType string
	// Created an endpoint
	EndpointOperations *shared.EndpointOperations
	// General Error
	GeneralError *shared.GeneralError
	StatusCode   int
	RawResponse  *http.Response
}
