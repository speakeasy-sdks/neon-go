// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EndpointCreateRequestEndpoint struct {
	AutoscalingLimitMaxCu *float64 `json:"autoscaling_limit_max_cu,omitempty"`
	AutoscalingLimitMinCu *float64 `json:"autoscaling_limit_min_cu,omitempty"`
	// The ID of the branch the compute endpoint will be associated with
	//
	BranchID string `json:"branch_id"`
	// Whether to restrict connections to the compute endpoint
	//
	Disabled *bool `json:"disabled,omitempty"`
	// NOT YET IMPLEMENTED. Whether to permit passwordless access to the compute endpoint.
	//
	PasswordlessAccess *bool `json:"passwordless_access,omitempty"`
	// Whether to enable connection pooling for the compute endpoint
	//
	PoolerEnabled *bool `json:"pooler_enabled,omitempty"`
	// The connection pooler mode. Neon supports PgBouncer in `transaction` mode only.
	//
	PoolerMode *EndpointPoolerMode `json:"pooler_mode,omitempty"`
	// The Neon compute provisioner.
	//
	Provisioner *Provisioner `json:"provisioner,omitempty"`
	// The region where the compute endpoint will be created. Only the project's `region_id` is permitted.
	//
	RegionID *string `json:"region_id,omitempty"`
	// A collection of settings for a compute endpoint
	Settings *EndpointSettingsData `json:"settings,omitempty"`
	// Duration of inactivity in seconds after which endpoint will be
	// automatically suspended. Value `0` means use global default,
	// `-1` means never suspend. Maximum value is 1 week in seconds.
	//
	SuspendTimeoutSeconds *int64 `json:"suspend_timeout_seconds,omitempty"`
	// The compute endpoint type. Either `read_write` or `read_only`.
	// The `read_only` compute endpoint type is not yet supported.
	//
	Type EndpointType `json:"type"`
}

type EndpointCreateRequest struct {
	Endpoint EndpointCreateRequestEndpoint `json:"endpoint"`
}