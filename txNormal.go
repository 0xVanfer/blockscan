package blockscan

import (
	"strconv"

	"github.com/0xVanfer/blockscan/internal/constants"
	"github.com/0xVanfer/blockscan/internal/utils"
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
func (s *Scanner) GetNormalTransactions(address any, startBlock int, endBlock any) (res BlockscanGetNormalTxsReq, err error) {
	toBlock, err := utils.ProcessToBlock(endBlock)
	if err != nil {
		return
	}
	addressStr, err := utils.CheckAddress(address)
	if err != nil {
		return
	}
	url := s.UrlHead + `module=account&action=txlist&address=` + addressStr + `&startblock=` + strconv.FormatInt(int64(startBlock), 10) + `&endblock=` + toBlock + `&sort=asc&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	err = r.ToJSON(&res)
	return
}

// Get all the txs of an address.
func (s *Scanner) GetNormalTransactionsAll(address any) (txs []BlockscanNormalTxs, err error) {
	res, err := s.GetNormalTransactions(address, 0, constants.UnreachableBlock)
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
		res, err = s.GetNormalTransactions(address, startBlock, constants.UnreachableBlock)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res.Result...)
	}
	return
}
