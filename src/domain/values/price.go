package values

type Price struct {
	Amount   uint   `bun:"amount" json:"amount"`
	Currency string `bun:"currency" json:"currency"`
}
