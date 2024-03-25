package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/keshavchand/swsimplify/simplifier"
	"github.com/keshavchand/swsimplify/view"
)

func main() {
	screen := view.IntialMode()
	if _, err := tea.NewProgram(screen).Run(); err != nil {
		panic(err)
	}

	service := simplifier.New()
	for _, txn := range screen.TxnInfo {
		service.AddTransaction(txn)
	}
	txnInfo, err := service.Simplify()
	if err != nil {
		log.Println(err)
	}
	for _, txn := range txnInfo {
		log.Println(txn.From, "->", txn.To, txn.Amount)
	}
}
