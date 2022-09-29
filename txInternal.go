package blockscan

import (
	"github.com/0xVanfer/blockscan/internal/constants"
	"github.com/0xVanfer/blockscan/internal/regularcheck"
	"github.com/0xVanfer/types"
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

// Return up to 10000 internal txs of an address.
func (s *Scanner) GetInternalTransactions(address any, startBlock int, endBlock any) ([]BlockscanInternalTxs, error) {
	toBlock, err := regularcheck.RegularCheckToBlock(endBlock)
	if err != nil {
		return nil, err
	}
	addressStr, err := regularcheck.RegularCheckAddress(address)
	if err != nil {
		return nil, err
	}
	url := s.UrlHead + `module=account&action=txlistinternal&address=` + addressStr + `&startblock=` + types.ToString(startBlock) + `&endblock=` + toBlock + `&sort=asc&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var res BlockscanGetInternalTxsReq
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

// Return all the internal txs of an address.
func (s *Scanner) GetInternalTransactionsAll(address any) ([]BlockscanInternalTxs, error) {
	res, err := s.GetInternalTransactions(address, 0, constants.UnreachableBlock)
	if err != nil {
		return nil, err
	}
	txs := res
	for len(res) == 10000 {
		lastEndBlock := res[len(res)-1].BlockNumber
		startBlock := types.ToInt(lastEndBlock) + 1
		res, err = s.GetInternalTransactions(address, startBlock, constants.UnreachableBlock)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res...)
	}
	return txs, nil
}
