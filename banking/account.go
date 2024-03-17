package banking

import (
	"fmt"
	"time"
)

type Account interface {
	Deposit(amount int64)
	Withdraw(amount int64)
	PrintStatement()
}

type defaultAccount struct {
	transactions []transaction
	printer      printer
}

type transaction struct {
	amount int64
	date   time.Time
}

func NewAccount() Account {
	return &defaultAccount{
		transactions: make([]transaction, 0),
		printer:      &consolePrinter{},
	}
}

func testAccount(printer printer) Account {
	return &defaultAccount{
		transactions: make([]transaction, 0),
		printer:      printer,
	}
}

func (account *defaultAccount) Deposit(amount int64) {
	account.transactions = append(account.transactions, transaction{date: time.Now(), amount: amount})
}

func (account *defaultAccount) Withdraw(amount int64) {

}

func (account *defaultAccount) PrintStatement() {
	account.printer.println("Date | Amount | Balance")

	for _, transaction := range account.transactions {
		line := fmt.Sprintf("%s | +%d | %d", transaction.date.Format("02.01.2006"), transaction.amount, transaction.amount)
		account.printer.println(line)
	}
}
