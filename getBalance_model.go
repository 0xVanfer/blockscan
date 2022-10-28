package blockscan

// ------------------------------ balance ------------------------------

type getBalanceReq struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type getBalancesReq struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Result  []getBalancesRes `json:"result"`
}

type getBalancesRes struct {
	Account string `json:"account"`
	Balance string `json:"balance"`
}
