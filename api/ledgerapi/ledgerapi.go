// Package ledgerapi is the wire shape of the ledger service's API: types and paths only, so the
// service and its client share one definition without either importing the other.
//
// Amounts are decimal strings, never JSON numbers, which a consumer would read as floats.
package ledgerapi

import "time"

const (
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
	OwnerRef  string    `json:"owner_ref"`
	Name      string    `json:"name"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Category struct {
	Ref       string    `json:"ref"`
	ID        string    `json:"id"`
	OwnerRef  string    `json:"owner_ref"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Expense struct {
	Ref        string            `json:"ref"`
	ID         string            `json:"id"`
	OwnerRef   string            `json:"owner_ref"`
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
	OwnerRef   string            `json:"owner_ref"`
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
	OwnerRef       string    `json:"owner_ref"`
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
	OwnerRef   string    `json:"owner_ref"`
	WalletID   string    `json:"wallet_id"`
	Amount     string    `json:"amount"`
	OccurredAt time.Time `json:"occurred_at"`
	ExpenseID  string    `json:"expense_id,omitempty"`
	IncomeID   string    `json:"income_id,omitempty"`
	TransferID string    `json:"transfer_id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

// OwnedRequest names one record of one owner.
type OwnedRequest struct {
	OwnerRef string `json:"owner_ref"`
	ID       string `json:"id"`
}

type OwnerRequest struct {
	OwnerRef string `json:"owner_ref"`
}

type RenameRequest struct {
	OwnerRef string `json:"owner_ref"`
	ID       string `json:"id"`
	Name     string `json:"name"`
}

// ListRequest narrows a listing to one owner, and optionally one wallet and the period [From, To).
type ListRequest struct {
	OwnerRef string    `json:"owner_ref"`
	WalletID string    `json:"wallet_id,omitempty"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
}

type CreateWalletRequest struct {
	OwnerRef string `json:"owner_ref"`
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
	ID       string `json:"id,omitempty"`
	OwnerRef string `json:"owner_ref"`
	Name     string `json:"name"`
	Color    string `json:"color"`
}

type CategoryResponse struct {
	Category Category `json:"category"`
}

type ListCategoriesResponse struct {
	Categories []Category `json:"categories"`
}

// ExpenseRequest creates an expense, or updates the one named by ID.
type ExpenseRequest struct {
	ID         string            `json:"id,omitempty"`
	OwnerRef   string            `json:"owner_ref"`
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
	ID         string            `json:"id,omitempty"`
	OwnerRef   string            `json:"owner_ref"`
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
	ID           string `json:"id,omitempty"`
	OwnerRef     string `json:"owner_ref"`
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
	OwnerRef string    `json:"owner_ref"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
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

// Empty is the request or response of an operation that carries nothing.
type Empty struct{}
