package blockscan

import (
	"github.com/0xVanfer/blockscan/internal/constants"
	"github.com/0xVanfer/blockscan/internal/regularcheck"
	"github.com/0xVanfer/types"

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

// Return up to 10000 erc20 txs of an address.
func (s *Scanner) GetErc20Transactions(address any, startBlock int, endBlock any) ([]BlockscanErc20Txs, error) {
	addressStr, err := regularcheck.RegularCheckAddress(address)
	if err != nil {
		return nil, err
	}
	toBlock, err := regularcheck.RegularCheckToBlock(endBlock)
	if err != nil {
		return nil, err
	}
	url := s.UrlHead + `module=account&action=tokentx&address=` + addressStr + `&startblock=` + types.ToString(startBlock) + `&endblock=` + toBlock + `&sort=asc&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var res BlockscanGetErc20TxsReq
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

// Return all the erc20 txs of an address.
func (s *Scanner) GetErc20TransactionsAll(address any) ([]BlockscanErc20Txs, error) {
	res, err := s.GetErc20Transactions(address, 0, constants.UnreachableBlock)
	if err != nil {
		return nil, err
	}
	txs := res
	for len(res) == 10000 {
		lastEndBlock := res[len(res)-1].BlockNumber
		startBlock := types.ToInt(lastEndBlock) + 1
		res, err = s.GetErc20Transactions(address, startBlock, constants.UnreachableBlock)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res...)
	}
	return txs, nil
}
