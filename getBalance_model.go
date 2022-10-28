package blockscan

// ------------------------------ balance ------------------------------

type getBalanceReq struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"` // result, balance
}

// Get up to 20 addresses' balance.
type getBalancesReq struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Result  []getBalancesRes `json:"result"` // result, balances
}

type getBalancesRes struct {
	Account string `json:"account"` // address
	Balance string `json:"balance"` // balance
}
