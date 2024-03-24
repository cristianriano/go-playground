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

func (t *transaction) symbol() string {
	if t.amount < 0 {
		return ""
	}
	return "+"
}

func testAccount(printer printer) Account {
	return &defaultAccount{
		transactions: make([]transaction, 0),
		printer:      printer,
	}
}

func (acc *defaultAccount) Deposit(amount int64) {
	acc.transactions = append(acc.transactions, transaction{date: time.Now(), amount: amount})
}

func (acc *defaultAccount) Withdraw(amount int64) {
	acc.transactions = append(acc.transactions, transaction{date: time.Now(), amount: amount * -1})
}

func (acc *defaultAccount) PrintStatement() {
	acc.printer.println("Date | Amount | Balance")

	for _, tx := range acc.transactions {
		line := fmt.Sprintf(
			"%s | %s%d | %d",
			tx.date.Format("02.01.2006"),
			tx.symbol(),
			tx.amount,
			tx.amount,
		)
		acc.printer.println(line)
	}
}
