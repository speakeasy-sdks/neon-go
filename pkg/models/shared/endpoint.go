// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type Endpoint struct {
	AutoscalingLimitMaxCu float64 `json:"autoscaling_limit_max_cu"`
	AutoscalingLimitMinCu float64 `json:"autoscaling_limit_min_cu"`
	// The ID of the branch that the compute endpoint is associated with
	//
	BranchID string `json:"branch_id"`
	// A timestamp indicating when the compute endpoint was created
	//
	CreatedAt time.Time `json:"created_at"`
	// The compute endpoint creation source
	//
	CreationSource string `json:"creation_source"`
	// The state of the compute endpoint
	//
	CurrentState EndpointState `json:"current_state"`
	// Whether to restrict connections to the compute endpoint
	//
	Disabled bool `json:"disabled"`
	// The hostname of the compute endpoint. This is the hostname specified when connecting to a Neon database.
	//
	Host string `json:"host"`
	// The compute endpoint ID. Compute endpoint IDs have an `ep-` prefix. For example: `ep-little-smoke-851426`
	//
	ID string `json:"id"`
	// A timestamp indicating when the compute endpoint was last active
	//
	LastActive *time.Time `json:"last_active,omitempty"`
	// Whether to permit passwordless access to the compute endpoint
	//
	PasswordlessAccess bool `json:"passwordless_access"`
	// The state of the compute endpoint
	//
	PendingState *EndpointState `json:"pending_state,omitempty"`
	// Whether connection pooling is enabled for the compute endpoint
	//
	PoolerEnabled bool `json:"pooler_enabled"`
	// The connection pooler mode. Neon supports PgBouncer in `transaction` mode only.
	//
	PoolerMode EndpointPoolerMode `json:"pooler_mode"`
	// The ID of the project to which the compute endpoint belongs
	//
	ProjectID string `json:"project_id"`
	// The Neon compute provisioner.
	//
	Provisioner Provisioner `json:"provisioner"`
	// DEPRECATED. Use the "host" property instead.
	//
	ProxyHost string `json:"proxy_host"`
	// The region identifier
	//
	RegionID string `json:"region_id"`
	// A collection of settings for a compute endpoint
	Settings EndpointSettingsData `json:"settings"`
	// Duration of inactivity in seconds after which endpoint will be
	// automatically suspended. Value `0` means use global default,
	// `-1` means never suspend. Maximum value is 1 week in seconds.
	//
	SuspendTimeoutSeconds int64 `json:"suspend_timeout_seconds"`
	// The compute endpoint type. Either `read_write` or `read_only`.
	// The `read_only` compute endpoint type is not yet supported.
	//
	Type EndpointType `json:"type"`
	// A timestamp indicating when the compute endpoint was last updated
	//
	UpdatedAt time.Time `json:"updated_at"`
}