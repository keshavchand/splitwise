package simplifier

import (
	"testing"

	"github.com/keshavchand/swsimplify/models"
	"github.com/stretchr/testify/assert"
)

func TestSimplify(t *testing.T) {
	service := New()
	service.AddTransaction(models.RawTransactionInfo{From: "A", To: "B", Amount: 10})
	service.AddTransaction(models.RawTransactionInfo{From: "B", To: "A", Amount: 10})

	result, err := service.Simplify()
	assert.NoError(t, err)
	if len(result) != 0 {
		t.Fatal("Expected 0 transactions, got", len(result))
	}
}

func TestComplex(t *testing.T) {
	t.Run("TestSimplifyZero", func(t *testing.T) {
		service := New()
		service.AddTransaction(models.RawTransactionInfo{From: "1", To: "2", Amount: 10})
		service.AddTransaction(models.RawTransactionInfo{From: "2", To: "3", Amount: 10})
		service.AddTransaction(models.RawTransactionInfo{From: "3", To: "1", Amount: 10})

		result, err := service.Simplify()
		assert.NoError(t, err)
		if len(result) != 0 {
			t.Fatal("Expected 0 transactions, got", len(result))
		}
	})

	t.Run("TestSimplifyNonZero", func(t *testing.T) {
		service := New()
		service.AddTransaction(models.RawTransactionInfo{From: "1", To: "2", Amount: 11})
		service.AddTransaction(models.RawTransactionInfo{From: "2", To: "3", Amount: 10})
		service.AddTransaction(models.RawTransactionInfo{From: "3", To: "1", Amount: 10})

		result, err := service.Simplify()
		assert.NoError(t, err)
		if len(result) != 1 {
			t.Fatal("Expected 1 transactions, got", len(result))
		}

		if result[0].From != "1" || result[0].To != "2" || result[0].Amount != 1 {
			t.Fatal("Expected 1 -> 1 1, got", result)
		}
	})
}
