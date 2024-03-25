package models

import "fmt"

type Transaction struct {
	From   uint64
	To     uint64
	Amount float64
}

func (t *Transaction) String() string {
	return fmt.Sprintf("%d -> %d %f", t.From, t.To, t.Amount)
}

type RawTransactionInfo struct {
	From, To    string
	Amount      float64
	Description string
}

func (r *RawTransactionInfo) String() string {
	return fmt.Sprintf("%s -> %s %.2f (%s)", r.From, r.To, r.Amount, r.Description)
}
