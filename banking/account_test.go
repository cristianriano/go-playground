package banking

import "testing"

// Either in the same package or in <pkg>_test
func TestPrintStatement_whenEmpty(t *testing.T) {
	printer := &testPrinter{text: ""}
	account := testAccount(printer)

	account.PrintStatement()
	if printer.text != "\nDate | Amount | Balance" {
		t.Error("Header doesn't match")
	}
}

func TestPrintStatement_afterDeposit(t *testing.T) {
	printer := &testPrinter{text: ""}
	account := testAccount(printer)

	account.Deposit(10)
	account.PrintStatement()
	if printer.text != "\nDate | Amount | Balance\n17.03.2024 | +10 | 10" {
		t.Error("Statement is not correct")
	}
}
