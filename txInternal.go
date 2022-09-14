package blockscan

import (
	"strconv"

	"github.com/0xVanfer/blockscan/internal/constants"
	"github.com/0xVanfer/blockscan/internal/types"
	"github.com/imroc/req"
)

type BlockscanGetInternalTxsReq struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Result  []BlockscanInternalTxs `json:"result"`
}

type BlockscanInternalTxs struct {
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	Hash            string `json:"hash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	Input           string `json:"input"`
	Type            string `json:"type"`
	Gas             string `json:"gas"`
	GasUsed         string `json:"gasUsed"`
	TraceID         string `json:"traceId"`
	IsError         string `json:"isError"`
	ErrCode         string `json:"errCode"`
}

// Get up to 10000 internal txs of an address.
// If "userApiKey" is "", use default api key.
func GetInternalTransactions[T int | string](network string, address string, startBlock int, endBlock T, userApiKey string) (res BlockscanGetInternalTxsReq, err error) {
	urlHead, apiKey, err := GetUrlAndKey(network)
	if err != nil {
		return
	}
	if userApiKey != "" {
		apiKey = userApiKey
	}
	url := urlHead + `module=account&action=txlistinternal&address=` + address + `&startblock=` + types.ToString(startBlock) + `&endblock=` + types.ToString(endBlock) + `&sort=asc&apikey=` + apiKey
	r, _ := req.Get(url)
	err = r.ToJSON(&res)
	return
}

// Get all the internal txs of an address.
// If "userApiKey" is "", use default api key.
func GetInternalTransactionsAll(network string, address string, userApiKey string) (txs []BlockscanInternalTxs, err error) {
	res, err := GetInternalTransactions(network, address, 0, constants.UnreachableBlock, userApiKey)
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
		res, err = GetInternalTransactions(network, address, startBlock, constants.UnreachableBlock, userApiKey)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res.Result...)
	}
	return
}
