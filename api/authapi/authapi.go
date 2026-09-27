// Package authapi is the wire shape of the identity service's API: types and paths only, so the
// service and its client share one definition without either importing the other.
package authapi

// Paths the site calls. Registering, authenticating, session creation, password change and setup
// are left out on purpose: the identity service's own pages handle them.
const (
	PathResolveSession = "/v1/resolve-session"
	PathGetUser        = "/v1/get-user"
	// PathGetUserByUsername tells whether a username exists. That is no secret from a signed-in
	// site: usernames are shown to every member of a shared project.
	PathGetUserByUsername = "/v1/get-user-by-username"
	PathListUsers         = "/v1/list-users"
	PathUpdateProfile     = "/v1/update-profile"
	PathDeleteUser        = "/v1/delete-user"
	PathSetupOpen         = "/v1/setup-open"
)

// User carries no credentials and no timestamps: nothing outside the identity service reads them.
type User struct {
	Ref      string `json:"ref"`
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

// ResolveSessionRequest carries a live session token. Nothing may log it. The token rides in the
// body to stay out of access logs, Referer headers and proxy error pages.
type ResolveSessionRequest struct {
	Token string `json:"token"`
}

type UserResponse struct {
	User User `json:"user"`
}

type GetUserRequest struct {
	UserRef string `json:"user_ref"`
}

type GetUserByUsernameRequest struct {
	Username string `json:"username"`
}

type ListUsersRequest struct {
	Limit int `json:"limit"`
}

type ListUsersResponse struct {
	Users []User `json:"users"`
}

type UpdateProfileRequest struct {
	UserRef  string `json:"user_ref"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type DeleteUserRequest struct {
	UserRef string `json:"user_ref"`
}

type SetupOpenResponse struct {
	Open bool `json:"open"`
}

// Empty answers a void operation, rather than 204, so success and failure decode through the same
// path.
type Empty struct{}
