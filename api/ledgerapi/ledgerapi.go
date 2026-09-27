// Package ledgerapi is the wire shape of the ledger service's API: types and paths only, so the
// service and its client share one definition without either importing the other.
//
// Amounts are decimal strings, never JSON numbers, which a consumer would read as floats.
package ledgerapi

import "time"

const (
	PathCreateWallet      = "/v1/create-wallet"
	PathRenameWallet      = "/v1/rename-wallet"
	PathDeleteWallet      = "/v1/delete-wallet"
	PathGetWallet         = "/v1/get-wallet"
	PathListWallets       = "/v1/list-wallets"
	PathCreateCategory    = "/v1/create-category"
	PathRenameCategory    = "/v1/rename-category"
	PathDeleteCategory    = "/v1/delete-category"
	PathGetCategory       = "/v1/get-category"
	PathListCategories    = "/v1/list-categories"
	PathCreateTransaction = "/v1/create-transaction"
	PathUpdateTransaction = "/v1/update-transaction"
	PathDeleteTransaction = "/v1/delete-transaction"
	PathGetTransaction    = "/v1/get-transaction"
	PathListTransactions  = "/v1/list-transactions"
	PathCreateTransfer    = "/v1/create-transfer"
	PathSummarize         = "/v1/summarize"
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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Transaction struct {
	Ref        string            `json:"ref"`
	ID         string            `json:"id"`
	OwnerRef   string            `json:"owner_ref"`
	WalletID   string            `json:"wallet_id"`
	Kind       string            `json:"kind"`
	Amount     string            `json:"amount"`
	CategoryID string            `json:"category_id,omitempty"`
	Note       string            `json:"note"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
	TransferID string            `json:"transfer_id,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
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

type CreateCategoryRequest struct {
	OwnerRef string `json:"owner_ref"`
	Name     string `json:"name"`
}

type CategoryResponse struct {
	Category Category `json:"category"`
}

type ListCategoriesResponse struct {
	Categories []Category `json:"categories"`
}

// TransactionRequest creates a transaction, or updates the one named by ID.
type TransactionRequest struct {
	ID         string            `json:"id,omitempty"`
	OwnerRef   string            `json:"owner_ref"`
	WalletID   string            `json:"wallet_id"`
	Kind       string            `json:"kind"`
	Amount     string            `json:"amount"`
	CategoryID string            `json:"category_id,omitempty"`
	Note       string            `json:"note"`
	OccurredAt time.Time         `json:"occurred_at"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

type TransactionResponse struct {
	Transaction Transaction `json:"transaction"`
}

type ListTransactionsRequest struct {
	OwnerRef string    `json:"owner_ref"`
	WalletID string    `json:"wallet_id,omitempty"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
}

type ListTransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type CreateTransferRequest struct {
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
	Out Transaction `json:"out"`
	In  Transaction `json:"in"`
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
	Amount     string `json:"amount"`
	Percent    string `json:"percent"`
}

type KindSummary struct {
	Total      string          `json:"total"`
	Categories []CategoryShare `json:"categories"`
}

type CurrencySummary struct {
	Currency string      `json:"currency"`
	Income   KindSummary `json:"income"`
	Expense  KindSummary `json:"expense"`
}

type SummaryResponse struct {
	Wallets    []WalletBalance   `json:"wallets"`
	Currencies []CurrencySummary `json:"currencies"`
}

// Empty is the request or response of an operation that carries nothing.
type Empty struct{}
