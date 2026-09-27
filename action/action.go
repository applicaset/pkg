// Package action holds the permission strings more than one binary names. The rest stay in web.
// The identity service checks who may reach its registration form and must not import the site.
package action

const (
	// UserCreate is the permission to add an account.
	UserCreate = "user.create"
	// AnyUser is the resource pattern for a question about accounts in general.
	AnyUser = "urn:auth:user:*"
)
