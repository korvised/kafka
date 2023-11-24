package commands

type OpenAccountCommand struct {
	AccountHolder  string  `json:"account_holder"`
	AccountType    int     `json:"account_type"`
	OpeningBalance float64 `json:"opening_balance"`
}

type DepositFundCommand struct {
	ID     string
	Amount float64 `json:"amount"`
}

type WithdrawFundCommand struct {
	ID     string
	Amount float64 `json:"amount"`
}

type CloseAccountCommand struct {
	ID string
}
