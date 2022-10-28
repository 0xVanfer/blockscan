package blockscan

type getErc20TxsReq struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Result  []erc20Txs `json:"result"`
}

type erc20Txs struct {
	BlockNumber       string `json:"blockNumber"`       // block number
	TimeStamp         string `json:"timeStamp"`         // timestamp
	Hash              string `json:"hash"`              // tx hash
	Nonce             string `json:"nonce"`             // tx nonce
	BlockHash         string `json:"blockHash"`         // block hash
	From              string `json:"from"`              // who sent this token
	ContractAddress   string `json:"contractAddress"`   // token contract
	To                string `json:"to"`                // who received this token
	Value             string `json:"value"`             // value in WEI
	TokenName         string `json:"tokenName"`         // token name
	TokenSymbol       string `json:"tokenSymbol"`       // token symbol
	TokenDecimal      string `json:"tokenDecimal"`      // token decimals
	TransactionIndex  string `json:"transactionIndex"`  // tx index
	Gas               string `json:"gas"`               // gas
	GasPrice          string `json:"gasPrice"`          // gas price
	GasUsed           string `json:"gasUsed"`           // gas used
	CumulativeGasUsed string `json:"cumulativeGasUsed"` //
	Input             string `json:"input"`             //
	Confirmations     string `json:"confirmations"`     //
}
