package banking

import (
	"fmt"
	"testing"
)

// Either in the same package or in <pkg>_test
func TestPrintStatement_whenEmpty(t *testing.T) {
	printer := &testPrinter{text: ""}
	account := testAccount(printer)

	account.PrintStatement()
	if printer.text != "\nDate | Amount | Balance" {
		t.Error(fmt.Sprintf("'%s' header is not correct", printer.text))
	}
}

func TestPrintStatement_afterDeposit(t *testing.T) {
	printer := &testPrinter{text: ""}
	account := testAccount(printer)

	account.Deposit(10)
	account.PrintStatement()
	if printer.text != "\nDate | Amount | Balance\n24.03.2024 | +10 | 10" {
		t.Error(fmt.Sprintf("'%s' statement is not correct", printer.text))
	}
}

func TestPrintStatement_afterWithdraw(t *testing.T) {
	printer := &testPrinter{text: ""}
	account := testAccount(printer)

	account.Withdraw(20)
	account.PrintStatement()
	if printer.text != "\nDate | Amount | Balance\n24.03.2024 | -20 | -20" {
		t.Error(fmt.Sprintf("'%s' statement is not correct", printer.text))
	}
}
