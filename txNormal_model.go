package blockscan

type getNormalTxsReq struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Result  []normalTxs `json:"result"`
}

type normalTxs struct {
	BlockNumber       string `json:"blockNumber"`       // tx block number
	TimeStamp         string `json:"timeStamp"`         // tx timestamp
	Hash              string `json:"hash"`              // tx hash
	Nonce             string `json:"nonce"`             // tx nonce
	BlockHash         string `json:"blockHash"`         // tx block hash
	TransactionIndex  string `json:"transactionIndex"`  // tx index
	From              string `json:"from"`              // who call the contract
	To                string `json:"to"`                // who received the call
	Value             string `json:"value"`             // value in WEI
	Gas               string `json:"gas"`               //
	GasPrice          string `json:"gasPrice"`          //
	IsError           string `json:"isError"`           //
	TxreceiptStatus   string `json:"txreceipt_status"`  //
	Input             string `json:"input"`             // call data
	ContractAddress   string `json:"contractAddress"`   //
	CumulativeGasUsed string `json:"cumulativeGasUsed"` //
	GasUsed           string `json:"gasUsed"`           //
	Confirmations     string `json:"confirmations"`     //
}
