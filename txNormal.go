package blockscan

import (
	"strconv"

	"github.com/0xVanfer/blockscan/internal/constants"
	"github.com/0xVanfer/blockscan/internal/types"
	"github.com/imroc/req"
)

type BlockscanGetNormalTxsReq struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Result  []BlockscanNormalTxs `json:"result"`
}

type BlockscanNormalTxs struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxreceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
}

// Get up to 10000 txs of an address.
// If "userApiKey" is "", use default api key.
func GetNormalTransactions[T int | string](network string, address string, startBlock int, endBlock T, userApiKey string) (res BlockscanGetNormalTxsReq, err error) {
	urlHead, apiKey, err := GetUrlAndKey(network)
	if err != nil {
		return
	}
	if userApiKey != "" {
		apiKey = userApiKey
	}
	url := urlHead + `module=account&action=txlist&address=` + address + `&startblock=` + types.ToString(startBlock) + `&endblock=` + types.ToString(endBlock) + `&sort=asc&apikey=` + apiKey
	r, _ := req.Get(url)
	err = r.ToJSON(&res)
	return
}

// Get all the txs of an address.
// If "userApiKey" is "", use default api key.
func GetNormalTransactionsAll(network string, address string, userApiKey string) (txs []BlockscanNormalTxs, err error) {
	res, err := GetNormalTransactions(network, address, 0, constants.UnreachableBlock, userApiKey)
	if err != nil {
		return
	}
	txs = append(txs, res.Result...)
	for len(res.Result) == 10000 {
		lastEndBlock := res.Result[len(res.Result)-1].BlockNumber
		lastEndBlockInt, err := strconv.Atoi(lastEndBlock)
		if err != nil {
			return nil, err
		}
		startBlock := lastEndBlockInt + 1
		res, err = GetNormalTransactions(network, address, startBlock, constants.UnreachableBlock, userApiKey)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res.Result...)
	}
	return
}
