package banking

import "testing"

// Either in the same package or in <pkg>_test
func TestPrintStatement(t *testing.T) {
	account := NewAccount()

	account.PrintStatement()
}
