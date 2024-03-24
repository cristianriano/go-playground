package banking

import (
	"fmt"
	"testing"
	"time"
)

// Either in the same package or in <pkg>_test
func TestPrintStatement_whenEmpty(t *testing.T) {
	printer := &testPrinter{text: ""}
	clock := &testClock{dates: []time.Time{marchOf(24)}}
	account := testAccount(printer, clock)

	account.PrintStatement()
	if printer.text != "\nDate | Amount | Balance" {
		t.Error(fmt.Sprintf("'%s' header is not correct", printer.text))
	}
}

func TestPrintStatement_afterDeposit(t *testing.T) {
	printer := &testPrinter{text: ""}
	clock := &testClock{dates: []time.Time{marchOf(24)}}
	account := testAccount(printer, clock)

	account.Deposit(10)
	account.PrintStatement()
	if printer.text != "\nDate | Amount | Balance\n24.03.2024 | +10 | 10" {
		t.Error(fmt.Sprintf("'%s' statement is not correct", printer.text))
	}
}

func TestPrintStatement_afterWithdraw(t *testing.T) {
	printer := &testPrinter{text: ""}
	clock := &testClock{dates: []time.Time{marchOf(24)}}
	account := testAccount(printer, clock)

	account.Withdraw(20)
	account.PrintStatement()
	if printer.text != "\nDate | Amount | Balance\n24.03.2024 | -20 | -20" {
		t.Error(fmt.Sprintf("'%s' statement is not correct", printer.text))
	}
}

func TestPrintStatement_showsTransactionsInOrder(t *testing.T) {
	printer := &testPrinter{text: ""}
	clock := &testClock{
		dates: []time.Time{marchOf(23), marchOf(22), marchOf(21)},
	}
	account := testAccount(printer, clock)

	account.Deposit(50)
	account.Deposit(10)
	account.Withdraw(20)

	account.PrintStatement()
	if printer.text != "\nDate | Amount | Balance\n23.03.2024 | -20 | 40\n22.03.2024 | 10 | 60\n21.03.2024 | 50 | 50" {
		t.Error(fmt.Sprintf("'%s' statement is not correct", printer.text))
	}
}

func marchOf(day int) time.Time {
	return time.Date(2024, 03, day, 0, 0, 0, 0, time.UTC)
}

func testAccount(printer printer, clock clock) Account {
	return &defaultAccount{
		transactions: make([]transaction, 0),
		printer:      printer,
		clock:        clock,
	}
}

type testPrinter struct {
	text string
}

type testClock struct {
	dates []time.Time
}

func (c *testPrinter) println(text string) {
	c.text = c.text + "\n" + text
}

func (c *testClock) now() time.Time {
	lastIndex := len(c.dates) - 1
	now := c.dates[lastIndex]

	c.dates = c.dates[:lastIndex]

	return now
}
