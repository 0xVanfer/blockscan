package blockscan

import (
	"github.com/0xVanfer/blockscan/internal/constants"
	"github.com/0xVanfer/blockscan/internal/utils"
	"github.com/0xVanfer/types"
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

// Return up to 10000 txs of an address.
func (s *Scanner) GetNormalTransactions(address any, startBlock int, endBlock any) ([]BlockscanNormalTxs, error) {
	toBlock, err := utils.CheckToBlock(endBlock)
	if err != nil {
		return nil, err
	}
	addressStr, err := utils.CheckAddress(address)
	if err != nil {
		return nil, err
	}
	url := s.UrlHead + `module=account&action=txlist&address=` + addressStr + `&startblock=` + types.ToString(startBlock) + `&endblock=` + toBlock + `&sort=asc&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var res BlockscanGetNormalTxsReq
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

// Return all the txs of an address.
func (s *Scanner) GetNormalTransactionsAll(address any) ([]BlockscanNormalTxs, error) {
	res, err := s.GetNormalTransactions(address, 0, constants.UnreachableBlock)
	if err != nil {
		return nil, err
	}
	txs := res
	for len(res) == 10000 {
		lastEndBlock := res[len(res)-1].BlockNumber
		startBlock := types.ToInt(lastEndBlock) + 1
		res, err = s.GetNormalTransactions(address, startBlock, constants.UnreachableBlock)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res...)
	}
	return txs, nil
}
