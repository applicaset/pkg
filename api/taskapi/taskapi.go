// Package taskapi is the wire shape of todoset's task service: types and paths only, so the
// service, its client and the site share one definition without importing each other.
package taskapi

import "time"

// Every request names the acting user. The service trusts it: only sibling services on the private
// network reach these paths, and the site has already resolved the session.
const (
	PathRoot          = "/v1/root"
	PathPage          = "/v1/page"
	PathCreate        = "/v1/create"
	PathRename        = "/v1/rename"
	PathDescribe      = "/v1/describe"
	PathSetCompleted  = "/v1/set-completed"
	PathReorder       = "/v1/reorder"
	PathMove          = "/v1/move"
	PathDelete        = "/v1/delete"
	PathUndo          = "/v1/undo"
	PathAddComment    = "/v1/add-comment"
	PathEditComment   = "/v1/edit-comment"
	PathDeleteComment = "/v1/delete-comment"
	PathContact       = "/v1/contact"
	// PathEvents answers with a text/event-stream of changed task ids, until the caller hangs up.
	PathEvents = "/v1/events"
)

type Task struct {
	ID                  string     `json:"id"`
	OwnerID             string     `json:"owner_id"`
	ParentID            string     `json:"parent_id"`
	Title               string     `json:"title"`
	Description         string     `json:"description"`
	Position            int64      `json:"position"`
	Completed           bool       `json:"completed"`
	ChildCount          int        `json:"child_count"`
	CompletedChildCount int        `json:"completed_child_count"`
	TitleRevision       int        `json:"title_revision"`
	DescriptionRevision int        `json:"description_revision"`
	Progress            string     `json:"progress"`
	ProgressPercent     float64    `json:"progress_percent"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at,omitempty"`
}

func (t Task) IsRoot() bool { return t.ParentID == "" }

type Activity struct {
	ID              string     `json:"id"`
	Kind            string     `json:"kind"`
	TaskID          string     `json:"task_id"`
	AuthorID        string     `json:"author_id"`
	Type            string     `json:"type"`
	Count           int        `json:"count"`
	From            string     `json:"from"`
	To              string     `json:"to"`
	Body            string     `json:"body"`
	BodyRevision    int        `json:"body_revision"`
	ParentCommentID string     `json:"parent_comment_id"`
	EditedAt        *time.Time `json:"edited_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func (a Activity) IsComment() bool { return a.Kind == "comment" }

type Page struct {
	Task     Task       `json:"task"`
	Parent   *Task      `json:"parent,omitempty"`
	Children []Task     `json:"children"`
	Activity []Activity `json:"activity"`
	// MoreActivity says entries older than the last one exist.
	MoreActivity bool `json:"more_activity"`
}

type Deletion struct {
	Task        Task      `json:"task"`
	Descendants int       `json:"descendants"`
	UndoUntil   time.Time `json:"undo_until"`
}

type ActorRequest struct {
	ActorID string `json:"actor_id"`
}

type PageRequest struct {
	ActorID        string `json:"actor_id"`
	TaskID         string `json:"task_id"`
	ActivityBefore string `json:"activity_before"`
}

type TaskRequest struct {
	ActorID string `json:"actor_id"`
	TaskID  string `json:"task_id"`
}

type CreateRequest struct {
	ActorID  string `json:"actor_id"`
	ParentID string `json:"parent_id"`
	Title    string `json:"title"`
}

// EditRequest carries the revision the text was typed over, for Rename and Describe.
type EditRequest struct {
	ActorID  string `json:"actor_id"`
	TaskID   string `json:"task_id"`
	Revision int    `json:"revision"`
	Text     string `json:"text"`
}

type SetCompletedRequest struct {
	ActorID   string `json:"actor_id"`
	TaskID    string `json:"task_id"`
	Completed bool   `json:"completed"`
}

type ReorderRequest struct {
	ActorID    string   `json:"actor_id"`
	ParentID   string   `json:"parent_id"`
	OrderedIDs []string `json:"ordered_ids"`
}

type MoveRequest struct {
	ActorID string `json:"actor_id"`
	TaskID  string `json:"task_id"`
	Offset  int    `json:"offset"`
}

type AddCommentRequest struct {
	ActorID         string `json:"actor_id"`
	TaskID          string `json:"task_id"`
	ParentCommentID string `json:"parent_comment_id"`
	Body            string `json:"body"`
}

type EditCommentRequest struct {
	ActorID   string `json:"actor_id"`
	CommentID string `json:"comment_id"`
	Revision  int    `json:"revision"`
	Body      string `json:"body"`
}

type CommentRequest struct {
	ActorID   string `json:"actor_id"`
	CommentID string `json:"comment_id"`
}

type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type TaskResponse struct {
	Task Task `json:"task"`
}

type PageResponse struct {
	Page Page `json:"page"`
}

type DeletionResponse struct {
	Deletion Deletion `json:"deletion"`
}

// Empty answers a void operation, so success and failure decode through the same path.
type Empty struct{}
