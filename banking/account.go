package banking

type Account interface {
	Deposit(amount int64)
	Withdraw(amount int64)
	PrintStatement()
}

type defaultAccount struct {
	balance int64
	printer printer
}

func NewAccount() Account {
	return &defaultAccount{
		balance: 0,
		printer: &consolePrinter{},
	}
}

func testAccount(printer printer) Account {
	return &defaultAccount{
		balance: 0,
		printer: printer,
	}
}

func (account *defaultAccount) Deposit(amount int64) {

}

func (account *defaultAccount) Withdraw(amount int64) {

}

func (account *defaultAccount) PrintStatement() {
	account.printer.println("Date\tAmount\tBalance")
}
