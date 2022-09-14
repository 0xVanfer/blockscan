package blockscan

import (
	"strconv"

	"github.com/0xVanfer/blockscan/internal/constants"

	"github.com/0xVanfer/blockscan/internal/types"
	"github.com/imroc/req"
)

type BlockscanGetErc20TxsReq struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Result  []BlockscanErc20Txs `json:"result"`
}

type BlockscanErc20Txs struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	Value             string `json:"value"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}

// Get up to 10000 erc20 txs of an address.
// If "userApiKey" is "", use default api key.
func GetErc20Transactions[T int | string](network string, address string, startBlock int, endBlock T, userApiKey string) (res BlockscanGetErc20TxsReq, err error) {
	urlHead, apiKey, err := GetUrlAndKey(network)
	if err != nil {
		return
	}
	if userApiKey != "" {
		apiKey = userApiKey
	}
	url := urlHead + `module=account&action=tokentx&address=` + address + `&startblock=` + types.ToString(startBlock) + `&endblock=` + types.ToString(endBlock) + `&sort=asc&apikey=` + apiKey
	r, _ := req.Get(url)
	err = r.ToJSON(&res)
	return
}

// Get all the erc20 txs of an address.
// If "userApiKey" is "", use default api key.
func GetErc20TransactionsAll(network string, address string, userApiKey string) (txs []BlockscanErc20Txs, err error) {
	res, err := GetErc20Transactions(network, address, 0, constants.UnreachableBlock, userApiKey)
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
		res, err = GetErc20Transactions(network, address, startBlock, constants.UnreachableBlock, userApiKey)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res.Result...)
	}
	return
}
