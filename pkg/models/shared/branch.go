// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type Branch struct {
	ActiveTimeSeconds  int64 `json:"active_time_seconds"`
	ComputeTimeSeconds int64 `json:"compute_time_seconds"`
	// CPU seconds used by all the endpoints of the branch, including deleted ones.
	// This value is reset at the beginning of each billing period.
	// Examples:
	// 1. A branch that uses 1 CPU for 1 second is equal to `cpu_used_sec=1`.
	// 2. A branch that uses 2 CPUs simultaneously for 1 second is equal to `cpu_used_sec=2`.
	//
	CPUUsedSec int64 `json:"cpu_used_sec"`
	// A timestamp indicating when the branch was created
	//
	CreatedAt time.Time `json:"created_at"`
	// The branch creation source
	//
	CreationSource string `json:"creation_source"`
	// The branch state
	CurrentState      BranchState `json:"current_state"`
	DataTransferBytes int64       `json:"data_transfer_bytes"`
	// The branch ID. This value is generated when a branch is created. A `branch_id` value has a `br` prefix. For example: `br-small-term-683261`.
	//
	ID string `json:"id"`
	// The logical size of the branch, in bytes
	//
	LogicalSize *int64 `json:"logical_size,omitempty"`
	// The branch name
	//
	Name string `json:"name"`
	// The `branch_id` of the parent branch
	//
	ParentID *string `json:"parent_id,omitempty"`
	// The Log Sequence Number (LSN) on the parent branch from which this branch was created
	//
	ParentLsn *string `json:"parent_lsn,omitempty"`
	// The point in time on the parent branch from which this branch was created
	//
	ParentTimestamp *time.Time `json:"parent_timestamp,omitempty"`
	// The branch state
	PendingState *BranchState `json:"pending_state,omitempty"`
	// Whether the branch is the project's primary branch
	//
	Primary bool `json:"primary"`
	// The ID of the project to which the branch belongs
	//
	ProjectID string `json:"project_id"`
	// A timestamp indicating when the branch was last updated
	//
	UpdatedAt        time.Time `json:"updated_at"`
	WrittenDataBytes int64     `json:"written_data_bytes"`
}
