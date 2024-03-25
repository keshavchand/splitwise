package simplifier

import (
	"math/rand"

	"github.com/keshavchand/swsimplify/models"
)

type UserGroup struct {
	users map[uint64]string
	name  map[string]uint64
}

func NewUserGroup() UserGroup {
	return UserGroup{
		users: map[uint64]string{},
		name:  map[string]uint64{},
	}
}

func (u *UserGroup) ConvertRaw(raw models.RawTransactionInfo) models.Transaction {
	return models.Transaction{
		From:   u.newUser(raw.From),
		To:     u.newUser(raw.To),
		Amount: raw.Amount,
	}
}

func (u *UserGroup) ConvertTxn(txn models.Transaction) models.RawTransactionInfo {
	return models.RawTransactionInfo{
		From:   u.getName(txn.From),
		To:     u.getName(txn.To),
		Amount: txn.Amount,
	}
}

func (u *UserGroup) newUser(name string) uint64 {
	id, ok := u.name[name]
	if ok {
		return id
	}

	for {
		id := rand.Uint64()
		if _, ok := u.users[id]; !ok {
			u.users[id] = name
			u.name[name] = id
			return id
		}
	}
}

func (u *UserGroup) getName(id uint64) string {
	return u.users[id]
}
