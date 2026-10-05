// Package discussapi is the wire shape of the discussion service: types and paths only, so the
// service, its client and its callers share one definition without importing each other.
package discussapi

import "time"

// Only sibling services on the private network reach these paths. The service trusts its caller to
// have decided who may read and write on a resource. It decides only who may change a comment.
const (
	PathListComments  = "/v1/list-comments"
	PathAddComment    = "/v1/add-comment"
	PathEditComment   = "/v1/edit-comment"
	PathDeleteComment = "/v1/delete-comment"
	PathPurgeResource = "/v1/purge-resource"
)

// Comment is one entry of a resource's thread. A deleted comment keeps its place with an empty
// body, so its replies keep their parent.
type Comment struct {
	ID          string     `json:"id"`
	ResourceRef string     `json:"resource_ref"`
	ParentID    string     `json:"parent_id"`
	AuthorRef   string     `json:"author_ref"`
	Body        string     `json:"body"`
	Revision    int        `json:"revision"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	EditedAt    *time.Time `json:"edited_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type ResourceRequest struct {
	ResourceRef string `json:"resource_ref"`
}

type CommentsResponse struct {
	Comments []Comment `json:"comments"`
}

type AddCommentRequest struct {
	ActorRef    string `json:"actor_ref"`
	ResourceRef string `json:"resource_ref"`
	ParentID    string `json:"parent_id"`
	Body        string `json:"body"`
}

// EditCommentRequest saves Body over Revision. ResourceRef must be the comment's own, so a caller
// that authorized one resource cannot reach a comment on another.
type EditCommentRequest struct {
	ActorRef    string `json:"actor_ref"`
	ResourceRef string `json:"resource_ref"`
	CommentID   string `json:"comment_id"`
	Revision    int    `json:"revision"`
	Body        string `json:"body"`
}

type DeleteCommentRequest struct {
	ActorRef    string `json:"actor_ref"`
	ResourceRef string `json:"resource_ref"`
	CommentID   string `json:"comment_id"`
}

type CommentResponse struct {
	Comment Comment `json:"comment"`
}

// Empty is the response of an operation that carries nothing.
type Empty struct{}
