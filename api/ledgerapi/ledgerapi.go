// Package ledgerapi is the wire shape of the ledger service's API: types and paths only, so the
// service and its client share one definition without either importing the other.
//
// Amounts are decimal strings, never JSON numbers. A consumer would read a number as a float.
package ledgerapi

import "time"

const (
	PathCreateProject     = "/v1/create-project"
	PathGetProject        = "/v1/get-project"
	PathListProjects      = "/v1/list-projects"
	PathRenameProject     = "/v1/rename-project"
	PathDeleteProject     = "/v1/delete-project"
	PathLeaveProject      = "/v1/leave-project"
	PathListMembers       = "/v1/list-members"
	PathSetMemberRole     = "/v1/set-member-role"
	PathRemoveMember      = "/v1/remove-member"
	PathTransferOwnership = "/v1/transfer-ownership"
	PathInvite            = "/v1/invite"
	PathListInvitations   = "/v1/list-invitations"
	PathCancelInvitation  = "/v1/cancel-invitation"
	PathListMyInvitations = "/v1/list-my-invitations"
	PathAcceptInvitation  = "/v1/accept-invitation"
	PathDeclineInvitation = "/v1/decline-invitation"

	PathCreateWallet     = "/v1/create-wallet"
	PathRenameWallet     = "/v1/rename-wallet"
	PathDeleteWallet     = "/v1/delete-wallet"
	PathGetWallet        = "/v1/get-wallet"
	PathListWallets      = "/v1/list-wallets"
	PathCreateCategory   = "/v1/create-category"
	PathUpdateCategory   = "/v1/update-category"
	PathDeleteCategory   = "/v1/delete-category"
	PathGetCategory      = "/v1/get-category"
	PathListCategories   = "/v1/list-categories"
	PathCreateExpense    = "/v1/create-expense"
	PathUpdateExpense    = "/v1/update-expense"
	PathDeleteExpense    = "/v1/delete-expense"
	PathGetExpense       = "/v1/get-expense"
	PathListExpenses     = "/v1/list-expenses"
	PathCreateIncome     = "/v1/create-income"
	PathUpdateIncome     = "/v1/update-income"
	PathDeleteIncome     = "/v1/delete-income"
	PathGetIncome        = "/v1/get-income"
	PathListIncomes      = "/v1/list-incomes"
	PathCreateTransfer   = "/v1/create-transfer"
	PathUpdateTransfer   = "/v1/update-transfer"
	PathDeleteTransfer   = "/v1/delete-transfer"
	PathGetTransfer      = "/v1/get-transfer"
	PathListTransfers    = "/v1/list-transfers"
	PathListTransactions = "/v1/list-transactions"
	PathSummarize        = "/v1/summarize"
)

type Wallet struct {
	Ref       string    `json:"ref"`
	ID        string    `json:"id"`
	ProjectID string    `json:"project_id"`
	CreatedBy string    `json:"created_by"`
	Name      string    `json:"name"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Category struct {
	Ref       string    `json:"ref"`
	ID        string    `json:"id"`
	ProjectID string    `json:"project_id"`
	CreatedBy string    `json:"created_by"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Expense struct {
	Ref        string            `json:"ref"`
	ID         string            `json:"id"`
	ProjectID  string            `json:"project_id"`
	CreatedBy  string            `json:"created_by"`
	WalletID   string            `json:"wallet_id"`
	Amount     string            `json:"amount"`
	CategoryID string            `json:"category_id,omitempty"`
	RefundOf   string            `json:"refund_of,omitempty"`
	Note       string            `json:"note"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

type Income struct {
	Ref        string            `json:"ref"`
	ID         string            `json:"id"`
	ProjectID  string            `json:"project_id"`
	CreatedBy  string            `json:"created_by"`
	WalletID   string            `json:"wallet_id"`
	Amount     string            `json:"amount"`
	Note       string            `json:"note"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

type Transfer struct {
	Ref            string    `json:"ref"`
	ID             string    `json:"id"`
	ProjectID      string    `json:"project_id"`
	CreatedBy      string    `json:"created_by"`
	FromWalletID   string    `json:"from_wallet_id"`
	ToWalletID     string    `json:"to_wallet_id"`
	Amount         string    `json:"amount"`
	ReceivedAmount string    `json:"received_amount"`
	Note           string    `json:"note"`
	OccurredAt     time.Time `json:"occurred_at"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Transaction is one movement of money in a wallet, read only. Exactly one of ExpenseID, IncomeID
// and TransferID is set.
type Transaction struct {
	Ref        string    `json:"ref"`
	ID         string    `json:"id"`
	ProjectID  string    `json:"project_id"`
	WalletID   string    `json:"wallet_id"`
	Amount     string    `json:"amount"`
	OccurredAt time.Time `json:"occurred_at"`
	ExpenseID  string    `json:"expense_id,omitempty"`
	IncomeID   string    `json:"income_id,omitempty"`
	TransferID string    `json:"transfer_id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

// Access names who asks and the project they ask about. Every request about a project carries it.
type Access struct {
	ActorRef  string `json:"actor_ref"`
	ProjectID string `json:"project_id"`
}

// RecordRequest names one record, or one invitation, of a project.
type RecordRequest struct {
	Access
	ID string `json:"id"`
}

type ProjectRequest struct {
	Access
}

type RenameRequest struct {
	Access
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListRequest narrows a listing to one project, and optionally one wallet and the period [From, To).
type ListRequest struct {
	Access
	WalletID string    `json:"wallet_id,omitempty"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
}

type CreateWalletRequest struct {
	Access
	Name     string `json:"name"`
	Currency string `json:"currency"`
}

type WalletResponse struct {
	Wallet Wallet `json:"wallet"`
}

type ListWalletsResponse struct {
	Wallets []Wallet `json:"wallets"`
}

// CategoryRequest creates a category, or updates the one named by ID.
type CategoryRequest struct {
	ID string `json:"id,omitempty"`
	Access
	Name  string `json:"name"`
	Color string `json:"color"`
}

type CategoryResponse struct {
	Category Category `json:"category"`
}

type ListCategoriesResponse struct {
	Categories []Category `json:"categories"`
}

// ExpenseRequest creates an expense, or updates the one named by ID.
type ExpenseRequest struct {
	ID string `json:"id,omitempty"`
	Access
	WalletID   string            `json:"wallet_id"`
	Amount     string            `json:"amount"`
	CategoryID string            `json:"category_id,omitempty"`
	RefundOf   string            `json:"refund_of,omitempty"`
	Note       string            `json:"note"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

type ExpenseResponse struct {
	Expense Expense `json:"expense"`
}

type ListExpensesResponse struct {
	Expenses []Expense `json:"expenses"`
}

// IncomeRequest creates income, or updates the one named by ID.
type IncomeRequest struct {
	ID string `json:"id,omitempty"`
	Access
	WalletID   string            `json:"wallet_id"`
	Amount     string            `json:"amount"`
	Note       string            `json:"note"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

type IncomeResponse struct {
	Income Income `json:"income"`
}

type ListIncomesResponse struct {
	Incomes []Income `json:"incomes"`
}

// TransferRequest creates a transfer, or updates the one named by ID.
type TransferRequest struct {
	ID string `json:"id,omitempty"`
	Access
	FromWalletID string `json:"from_wallet_id"`
	ToWalletID   string `json:"to_wallet_id"`
	Amount       string `json:"amount"`
	// ReceivedAmount may be empty when both wallets share a currency.
	ReceivedAmount string    `json:"received_amount,omitempty"`
	Note           string    `json:"note"`
	OccurredAt     time.Time `json:"occurred_at"`
}

type TransferResponse struct {
	Transfer Transfer `json:"transfer"`
}

type ListTransfersResponse struct {
	Transfers []Transfer `json:"transfers"`
}

type ListTransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type SummarizeRequest struct {
	Access
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type WalletBalance struct {
	Wallet  Wallet `json:"wallet"`
	Balance string `json:"balance"`
}

type CategoryShare struct {
	CategoryID string `json:"category_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Color      string `json:"color,omitempty"`
	Amount     string `json:"amount"`
	Percent    string `json:"percent"`
}

type Spending struct {
	Currency   string          `json:"currency"`
	Total      string          `json:"total"`
	Categories []CategoryShare `json:"categories"`
}

type SummaryResponse struct {
	Wallets  []WalletBalance `json:"wallets"`
	Spending []Spending      `json:"spending"`
}

type Project struct {
	Ref       string    `json:"ref"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Membership is a project with the role the asking user holds in it: owner, editor or viewer.
type Membership struct {
	Project Project `json:"project"`
	Role    string  `json:"role"`
}

type Member struct {
	ProjectID string    `json:"project_id"`
	UserRef   string    `json:"user_ref"`
	Role      string    `json:"role"`
	JoinedAt  time.Time `json:"joined_at"`
}

type Invitation struct {
	Ref        string    `json:"ref"`
	ID         string    `json:"id"`
	ProjectID  string    `json:"project_id"`
	InviteeRef string    `json:"invitee_ref"`
	Role       string    `json:"role"`
	InvitedBy  string    `json:"invited_by"`
	CreatedAt  time.Time `json:"created_at"`
}

// ProjectInvitation is an invitation as its invitee sees it, with the project it is for.
type ProjectInvitation struct {
	Invitation Invitation `json:"invitation"`
	Project    Project    `json:"project"`
}

// ActorRequest names the user asking, about their own projects or invitations.
type ActorRequest struct {
	ActorRef string `json:"actor_ref"`
}

type CreateProjectRequest struct {
	ActorRef string `json:"actor_ref"`
	Name     string `json:"name"`
}

type RenameProjectRequest struct {
	Access
	Name string `json:"name"`
}

// MemberRequest names a member of the project. Role is set only to change theirs.
type MemberRequest struct {
	Access
	UserRef string `json:"user_ref"`
	Role    string `json:"role,omitempty"`
}

// TransferOwnershipRequest hands the project to the member at UserRef. FormerOwnerRole is what the
// owner keeps: editor or viewer.
type TransferOwnershipRequest struct {
	Access
	UserRef         string `json:"user_ref"`
	FormerOwnerRole string `json:"former_owner_role"`
}

type InviteRequest struct {
	Access
	InviteeRef string `json:"invitee_ref"`
	Role       string `json:"role"`
}

// InviteeRequest is the invitee answering an invitation.
type InviteeRequest struct {
	ActorRef string `json:"actor_ref"`
	ID       string `json:"id"`
}

type ProjectResponse struct {
	Project Project `json:"project"`
}

type MembershipResponse struct {
	Membership Membership `json:"membership"`
}

type ListMembershipsResponse struct {
	Memberships []Membership `json:"memberships"`
}

type MemberResponse struct {
	Member Member `json:"member"`
}

type ListMembersResponse struct {
	Members []Member `json:"members"`
}

type InvitationResponse struct {
	Invitation Invitation `json:"invitation"`
}

type ListInvitationsResponse struct {
	Invitations []Invitation `json:"invitations"`
}

type ListProjectInvitationsResponse struct {
	Invitations []ProjectInvitation `json:"invitations"`
}

// Empty is the request or response of an operation that carries nothing.
type Empty struct{}
