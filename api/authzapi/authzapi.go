// Package authzapi is the wire shape of the authorization service's API: types and paths only, so
// the service and its clients share one definition without any of them importing another.
package authzapi

const (
	PathCan           = "/v1/can"
	PathGrant         = "/v1/grant"
	PathAssignRole    = "/v1/assign-role"
	PathRevokeRole    = "/v1/revoke-role"
	PathSubjectRoles  = "/v1/subject-roles"
	PathListRoles     = "/v1/roles"
	PathDefineRole    = "/v1/define-role"
	PathPurgeResource = "/v1/purge-resource"
	PathPurgeSubject  = "/v1/purge-subject"
)

// CanRequest asks for the subject alone, or also for the groups it belongs to.
type CanRequest struct {
	Subject  string   `json:"subject"`
	Groups   []string `json:"groups,omitempty"`
	Action   string   `json:"action"`
	Resource string   `json:"resource"`
}

type CanResponse struct {
	Allowed bool `json:"allowed"`
}

type GrantRequest struct {
	Subject  string   `json:"subject"`
	Actions  []string `json:"actions"`
	Resource string   `json:"resource"`
}

type RoleRequest struct {
	Subject string `json:"subject"`
	Role    string `json:"role"`
}

type SubjectRequest struct {
	Subject string `json:"subject"`
}

type ResourceRequest struct {
	Resource string `json:"resource"`
}

// Permission is an action over a resource. Either may end in a wildcard.
type Permission struct {
	Action   string `json:"action"`
	Resource string `json:"resource"`
}

// DefineRoleRequest creates a role or replaces its description and permissions.
type DefineRoleRequest struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions"`
}

// RolesResponse carries role names only. No caller reads a description.
type RolesResponse struct {
	Roles []string `json:"roles"`
}

// Empty is the request or response of an operation that carries nothing.
type Empty struct{}
