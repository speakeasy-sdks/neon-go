// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type BranchCreateRequestBranch struct {
	// The branch name
	//
	Name *string `json:"name,omitempty"`
	// The `branch_id` of the parent branch
	//
	ParentID *string `json:"parent_id,omitempty"`
	// A Log Sequence Number (LSN) on the parent branch. The branch will be created with data from this LSN.
	//
	ParentLsn *string `json:"parent_lsn,omitempty"`
	// A timestamp identifying a point in time on the parent branch. The branch will be created with data starting from this point in time.
	//
	ParentTimestamp *time.Time `json:"parent_timestamp,omitempty"`
}

type BranchCreateRequest struct {
	Branch    *BranchCreateRequestBranch           `json:"branch,omitempty"`
	Endpoints []BranchCreateRequestEndpointOptions `json:"endpoints,omitempty"`
}