package simplifier

import "github.com/keshavchand/swsimplify/models"

type Simplifier interface {
	SimplifyTransactions(txns []models.Transaction) ([]models.Transaction, error)
}

type Service struct {
	UserGroup  UserGroup
	Txn        []models.Transaction
	Simplifier Simplifier
}

func New() *Service {
	return &Service{
		Simplifier: &Naive{},
		UserGroup:  NewUserGroup(),
	}
}

func (s *Service) AddTransaction(t models.RawTransactionInfo) {
	s.Txn = append(s.Txn, s.UserGroup.ConvertRaw(t))
}

func (s *Service) Simplify() ([]models.RawTransactionInfo, error) {
	txn, err := s.Simplifier.SimplifyTransactions(s.Txn)
	if err != nil {
		return nil, err
	}
	result := []models.RawTransactionInfo{}
	for _, t := range txn {
		result = append(result, s.UserGroup.ConvertTxn(t))
	}
	return result, nil
}
