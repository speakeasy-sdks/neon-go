// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RoleCreateRequestRole struct {
	// The role name. Cannot exceed 63 bytes in length.
	//
	Name string `json:"name"`
}

type RoleCreateRequest struct {
	Role RoleCreateRequestRole `json:"role"`
}
