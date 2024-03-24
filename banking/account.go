package banking

import (
	"fmt"
	"sort"
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
	clock        clock
	total        int64
}

type transaction struct {
	amount int64
	date   time.Time
	total  int64
}

func NewAccount() Account {
	return &defaultAccount{
		transactions: make([]transaction, 0),
		printer:      &consolePrinter{},
		clock:        &defaultClock{},
		total:        0,
	}
}

func (acc *defaultAccount) Deposit(amount int64) {
	acc.total = acc.total + amount
	acc.transactions = append(acc.transactions, transaction{date: acc.clock.now(), amount: amount, total: acc.total})
}

func (acc *defaultAccount) Withdraw(amount int64) {
	acc.total = acc.total - amount
	acc.transactions = append(acc.transactions, transaction{date: acc.clock.now(), amount: amount * -1, total: acc.total})
}

func (acc *defaultAccount) PrintStatement() {
	acc.printer.println("Date | Amount | Balance")

	acc.sortTransactions()
	for _, tx := range acc.transactions {
		line := fmt.Sprintf(
			"%s | %s%d | %d",
			tx.date.Format("02.01.2006"),
			tx.symbol(),
			tx.amount,
			tx.total,
		)
		acc.printer.println(line)
	}
}

func (acc *defaultAccount) sortTransactions() {
	sort.Slice(acc.transactions, func(i, j int) bool {
		return acc.transactions[i].date.After(acc.transactions[j].date)
	})
}

func (t *transaction) symbol() string {
	if t.amount < 0 {
		return ""
	}
	return "+"
}
