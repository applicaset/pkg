// Package billingapi is the wire shape of the billing service's API: types and paths only, so the
// service and its client share one definition without either importing the other.
//
// Amounts and rates are decimal strings, never JSON numbers. A consumer would read a number as a
// float.
package billingapi

import "time"

const (
	PathGetAccess        = "/v1/get-access"
	PathListWallets      = "/v1/list-wallets"
	PathListAllWallets   = "/v1/list-all-wallets"
	PathListTransactions = "/v1/list-transactions"
	PathDeposit          = "/v1/deposit"
	PathWithdraw         = "/v1/withdraw"
	PathExchange         = "/v1/exchange"
	PathCreateCharge     = "/v1/create-charge"
	PathUpdateCharge     = "/v1/update-charge"
	PathDeleteCharge     = "/v1/delete-charge"
	PathPostCharge       = "/v1/post-charge"
	PathGetCharge        = "/v1/get-charge"
	PathListCharges      = "/v1/list-charges"
	PathDisputeCharge    = "/v1/dispute-charge"
	PathDismissDispute   = "/v1/dismiss-dispute"
	PathVoidCharge       = "/v1/void-charge"
)

type Wallet struct {
	OwnerRef string `json:"owner_ref"`
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
}

type Transaction struct {
	Ref             string    `json:"ref"`
	ID              string    `json:"id"`
	OwnerRef        string    `json:"owner_ref"`
	Currency        string    `json:"currency"`
	Amount          string    `json:"amount"`
	Kind            string    `json:"kind"`
	GroupID         string    `json:"group_id"`
	CounterpartyRef string    `json:"counterparty_ref,omitempty"`
	ChargeID        string    `json:"charge_id,omitempty"`
	ReversesID      string    `json:"reverses_id,omitempty"`
	Rate            string    `json:"rate,omitempty"`
	Note            string    `json:"note,omitempty"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
}

type Charge struct {
	Ref           string     `json:"ref"`
	ID            string     `json:"id"`
	BillerRef     string     `json:"biller_ref"`
	PayerRef      string     `json:"payer_ref"`
	Description   string     `json:"description"`
	Amount        string     `json:"amount"`
	Currency      string     `json:"currency"`
	Status        string     `json:"status"`
	DisputeReason string     `json:"dispute_reason,omitempty"`
	DismissedBy   string     `json:"dismissed_by,omitempty"`
	VoidedBy      string     `json:"voided_by,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	PostedAt      *time.Time `json:"posted_at,omitempty"`
	DisputedAt    *time.Time `json:"disputed_at,omitempty"`
	DismissedAt   *time.Time `json:"dismissed_at,omitempty"`
	VoidedAt      *time.Time `json:"voided_at,omitempty"`
}

// ActorRequest names who asks. Every request carries it.
type ActorRequest struct {
	ActorRef string `json:"actor_ref"`
}

type AccessResponse struct {
	Admin bool `json:"admin"`
}

type ListWalletsRequest struct {
	ActorRequest
	OwnerRef string `json:"owner_ref"`
}

type ListWalletsResponse struct {
	Wallets []Wallet `json:"wallets"`
}

type ListTransactionsRequest struct {
	ActorRequest
	OwnerRef string `json:"owner_ref,omitempty"`
	Currency string `json:"currency,omitempty"`
	ChargeID string `json:"charge_id,omitempty"`
}

type ListTransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type MoveRequest struct {
	ActorRequest
	OwnerRef string `json:"owner_ref"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	Note     string `json:"note,omitempty"`
}

type TransactionResponse struct {
	Transaction Transaction `json:"transaction"`
}

type ExchangeRequest struct {
	ActorRequest
	OwnerRef     string `json:"owner_ref"`
	FromCurrency string `json:"from_currency"`
	ToCurrency   string `json:"to_currency"`
	Amount       string `json:"amount"`
	Rate         string `json:"rate"`
	Note         string `json:"note,omitempty"`
}

type ChargeRequest struct {
	ActorRequest
	// ID names the charge to update. Create ignores it.
	ID          string `json:"id,omitempty"`
	PayerRef    string `json:"payer_ref"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
}

// ChargeIDRequest names one charge.
type ChargeIDRequest struct {
	ActorRequest
	ID string `json:"id"`
}

type DisputeRequest struct {
	ActorRequest
	ID     string `json:"id"`
	Reason string `json:"reason"`
}

type ListChargesRequest struct {
	ActorRequest
	BillerRef string `json:"biller_ref,omitempty"`
	PayerRef  string `json:"payer_ref,omitempty"`
	Status    string `json:"status,omitempty"`
}

type ChargeResponse struct {
	Charge Charge `json:"charge"`
}

type ListChargesResponse struct {
	Charges []Charge `json:"charges"`
}

type Empty struct{}
