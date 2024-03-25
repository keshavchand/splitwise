package simplifier

import (
	"container/heap"
	"errors"

	"github.com/keshavchand/swsimplify/models"
)

type Owes struct {
	Person uint64
	Amount float64
}

type Naive struct{}

func (n *Naive) SimplifyTransactions(txns []models.Transaction) ([]models.Transaction, error) {
	finalAmountForUser := map[uint64]float64{}
	for _, txn := range txns {
		from := txn.From
		to := txn.To
		amount := txn.Amount

		finalAmountForUser[from] -= amount
		finalAmountForUser[to] += amount
	}

	var min MinOwes
	var max MaxOwes

	var verifier float64
	for user, amount := range finalAmountForUser {
		verifier += amount
		if amount == 0 {
			continue
		}
		if amount < 0 {
			min.Push(Owes{Person: user, Amount: -amount})
		} else {
			max.Push(Owes{Person: user, Amount: amount})
		}
	}

	if verifier != 0 {
		return nil, errors.New("Inconsistent transactions")
	}

	heap.Init(&min)
	heap.Init(&max)

	var result []models.Transaction
	for len(min) > 0 || len(max) > 0 {
		from := min.Pop().(Owes)
		to := max.Pop().(Owes)

		result = append(result, models.Transaction{From: from.Person, To: to.Person, Amount: to.Amount})
		to.Amount -= from.Amount

		if to.Amount == 0 {
			continue
		}

		if to.Amount < 0 {
			min.Push(Owes{Person: to.Person, Amount: to.Amount})
		} else {
			max.Push(Owes{Person: to.Person, Amount: to.Amount})
		}
	}

	return result, nil
}
