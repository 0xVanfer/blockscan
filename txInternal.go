package blockscan

import (
	"strconv"

	"github.com/0xVanfer/blockscan/internal/constants"
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
func (s *Scanner) GetInternalTransactions(address any, startBlock int, endBlock any) (res BlockscanGetInternalTxsReq, err error) {
	toBlock, err := processToBlock(endBlock)
	if err != nil {
		return
	}
	addressStr, err := checkAddress(address)
	if err != nil {
		return
	}
	url := s.UrlHead + `module=account&action=txlistinternal&address=` + addressStr + `&startblock=` + strconv.FormatInt(int64(startBlock), 10) + `&endblock=` + toBlock + `&sort=asc&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	err = r.ToJSON(&res)
	return
}

// Get all the internal txs of an address.
func (s *Scanner) GetInternalTransactionsAll(address any) (txs []BlockscanInternalTxs, err error) {
	res, err := s.GetInternalTransactions(address, 0, constants.UnreachableBlock)
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
		res, err = s.GetInternalTransactions(address, startBlock, constants.UnreachableBlock)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res.Result...)
	}
	return
}
