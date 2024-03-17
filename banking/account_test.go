package banking

import "testing"

// Either in the same package or in <pkg>_test
func TestPrintStatement_whenEmpty(t *testing.T) {
	printer := &testPrinter{text: ""}
	account := testAccount(printer)

	account.PrintStatement()
	if printer.text != "\nDate\tAmount\tBalance" {
		t.Error("Header doesn't match")
	}
}
