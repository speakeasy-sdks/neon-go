// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ProjectQuota - Per-project consumption quota. If quota is exceeded, then all active computes
// will be automatically suspended and it won't be possible to start them neither
// with API method call or with incoming proxy connections. The only exception is
// `logical_size_bytes`, which is applied on per-branch basis, i.e. only compute
// on the branch with `logical_size` exceeding quota will be suspended.
//
// Quotas are enforced based on per-project consumption metrics with the same names,
// which are reset at the end of each billing period (first day of the month).
// Logical size is also an exception here, as it represents the total size of data
// stored in some branch, so it's not reset.
//
// Zero quota value or empty mean 'unlimited'.
type ProjectQuota struct {
	// The total amount of wall-clock time allowed to be spent by project's compute endpoints.
	//
	ActiveTimeSeconds *int64 `json:"active_time_seconds,omitempty"`
	// The total amount of CPU seconds allowed to be spent by project's compute endpoints.
	//
	ComputeTimeSeconds *int64 `json:"compute_time_seconds,omitempty"`
	// Total amount of data transferred from all project's branches using proxy.
	//
	DataTransferBytes *int64 `json:"data_transfer_bytes,omitempty"`
	// Limit on the logical size of every project's branch.
	//
	LogicalSizeBytes *int64 `json:"logical_size_bytes,omitempty"`
	// Total amount of data written to all project's branches.
	//
	WrittenDataBytes *int64 `json:"written_data_bytes,omitempty"`
}