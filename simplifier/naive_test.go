package simplifier

import (
	"testing"

	"github.com/keshavchand/swsimplify/models"
	"github.com/stretchr/testify/assert"
)

func TestSimplifySimple(t *testing.T) {
	n := Naive{}
	txn := []models.Transaction{
		{From: 1, To: 2, Amount: 10},
		{From: 2, To: 1, Amount: 10},
	}

	result, err := n.SimplifyTransactions(txn)
	assert.NoError(t, err)
	if len(result) != 0 {
		t.Fatal("Expected 0 transactions, got", len(result))
	}
}

func TestSimplifyComplex(t *testing.T) {
	t.Run("TestSimplifyZero", func(t *testing.T) {
		n := Naive{}
		txn := []models.Transaction{
			{From: 1, To: 2, Amount: 10},
			{From: 2, To: 3, Amount: 10},
			{From: 3, To: 1, Amount: 10},
		}

		result, err := n.SimplifyTransactions(txn)
		assert.NoError(t, err)
		if len(result) != 0 {
			t.Fatal("Expected 0 transactions, got", len(result))
		}
	})

	t.Run("TestSimplifyNonZero", func(t *testing.T) {
		n := Naive{}
		txn := []models.Transaction{
			{From: 1, To: 2, Amount: 11},
			{From: 2, To: 3, Amount: 10},
			{From: 3, To: 1, Amount: 10},
		}

		result, err := n.SimplifyTransactions(txn)
		assert.NoError(t, err)
		if len(result) != 1 {
			t.Fatal("Expected 1 transactions, got", len(result))
		}

		if result[0].From != 1 || result[0].To != 2 || result[0].Amount != 1 {
			t.Fatal("Expected 1 -> 1 1, got", result)
		}
	})
}
