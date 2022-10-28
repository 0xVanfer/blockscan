package blockscan

// ------------------------------ gas price ------------------------------

type gasPriceReq struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}
