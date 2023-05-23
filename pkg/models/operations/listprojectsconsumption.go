// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"neon/pkg/models/shared"
	"net/http"
)

type ListProjectsConsumptionRequest struct {
	// Specify the cursor value from the previous response to get the next batch of projects.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Specify a value from 1 to 1000 to limit number of projects in the response.
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
}

// ListProjectsConsumption200ApplicationJSON - Returned a list of per-project consumption for the Neon account
type ListProjectsConsumption200ApplicationJSON struct {
	// Cursor based pagination is used. The user must pass the cursor as is to the backend.
	// For more information about cursor based pagination, see
	// https://learn.microsoft.com/en-us/ef/core/querying/pagination#keyset-pagination
	//
	Pagination *shared.Pagination          `json:"pagination,omitempty"`
	Projects   []shared.ProjectConsumption `json:"projects"`
}

type ListProjectsConsumptionResponse struct {
	ContentType string
	// General Error
	GeneralError *shared.GeneralError
	StatusCode   int
	RawResponse  *http.Response
	// Returned a list of per-project consumption for the Neon account
	ListProjectsConsumption200ApplicationJSONObject *ListProjectsConsumption200ApplicationJSON
}
