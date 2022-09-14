package blockscan

import (
	"strconv"

	"github.com/0xVanfer/blockscan/internal/constants"

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

// Get up to 10000 erc20 txs of an address.
func (s *Scanner) GetErc20Transactions(address string, startBlock int, endBlock any) (res BlockscanGetErc20TxsReq, err error) {
	toBlock, err := processToBlock(endBlock)
	if err != nil {
		return
	}
	url := s.UrlHead + `module=account&action=tokentx&address=` + address + `&startblock=` + types.ToString(startBlock) + `&endblock=` + toBlock + `&sort=asc&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	err = r.ToJSON(&res)
	return
}

// Get all the erc20 txs of an address.
func (s *Scanner) GetErc20TransactionsAll(address string) (txs []BlockscanErc20Txs, err error) {
	res, err := s.GetErc20Transactions(address, 0, constants.UnreachableBlock)
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
		res, err = s.GetErc20Transactions(address, startBlock, constants.UnreachableBlock)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res.Result...)
	}
	return
}
