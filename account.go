package blockscan

import (
	"fmt"

	"github.com/0xVanfer/blockscan/internal/constants"
	"github.com/0xVanfer/blockscan/internal/regularcheck"
	"github.com/0xVanfer/types"
	"github.com/0xVanfer/utils"
	"github.com/shopspring/decimal"
)

// Return the balance of a single address.
func (s *Scanner) GetBalance(address string) (decimal.Decimal, error) {
	url := fmt.Sprintf("%smodule=account&action=balance&address=%s&tag=latest&apikey=%s", s.UrlHead, address, s.ApiKey)
	var res decimal.Decimal
	err := s.httpGetEtherscan(url, &res)
	if err != nil {
		return decimal.Zero, err
	}
	return res.Shift(-18), nil
}

// Return balances of up to 20 addresses. map[address] = balance.
func (s *Scanner) GetBalances(addresses ...string) (map[string]decimal.Decimal, error) {
	addressesString := utils.ConnectArray(",", addresses...)
	url := fmt.Sprintf("%smodule=account&action=balancemulti&address=%s&tag=latest&apikey=%s", s.UrlHead, addressesString, s.ApiKey)
	var res []struct {
		Account string          `json:"account"` // address
		Balance decimal.Decimal `json:"balance"` // balance
	}

	err := s.httpGetEtherscan(url, &res)
	if err != nil {
		return nil, err
	}

	resMap := make(map[string]decimal.Decimal)
	for _, info := range res {
		resMap[info.Account] = info.Balance.Shift(-18)
	}
	return resMap, nil
}

type Erc20Txs struct {
	BlockNumber       string `json:"blockNumber"`       // block number
	TimeStamp         string `json:"timeStamp"`         // timestamp
	Hash              string `json:"hash"`              // tx hash
	Nonce             string `json:"nonce"`             // tx nonce
	BlockHash         string `json:"blockHash"`         // block hash
	From              string `json:"from"`              // who sent this token
	ContractAddress   string `json:"contractAddress"`   // token contract
	To                string `json:"to"`                // who received this token
	Value             string `json:"value"`             // value in WEI
	TokenName         string `json:"tokenName"`         // token name
	TokenSymbol       string `json:"tokenSymbol"`       // token symbol
	TokenDecimal      string `json:"tokenDecimal"`      // token decimals
	TransactionIndex  string `json:"transactionIndex"`  // tx index
	Gas               string `json:"gas"`               // gas
	GasPrice          string `json:"gasPrice"`          // gas price
	GasUsed           string `json:"gasUsed"`           // gas used
	CumulativeGasUsed string `json:"cumulativeGasUsed"` //
	Input             string `json:"input"`             //
	Confirmations     string `json:"confirmations"`     //
}

// Return up to 10000 erc20 txs of an address.
func (s *Scanner) GetErc20Transactions(address string, startBlock int64, endBlock any) ([]Erc20Txs, error) {
	toBlock, err := regularcheck.RegularCheckToBlock(endBlock)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%smodule=account&action=tokentx&address=%s&startblock=%d&endblock=%s&sort=asc&apikey=%s", s.UrlHead, address, startBlock, toBlock, s.ApiKey)
	var res []Erc20Txs
	err = s.httpGetEtherscan(url, &res)
	return res, err
}

// Return all the erc20 txs of an address.
func (s *Scanner) GetErc20TransactionsAll(address string) ([]Erc20Txs, error) {
	res, err := s.GetErc20Transactions(address, 0, constants.UnreachableBlock)
	if err != nil {
		return nil, err
	}
	txs := res
	for len(res) == 10000 {
		lastEndBlock := res[len(res)-1].BlockNumber
		startBlock := types.ToInt64(lastEndBlock) + 1
		res, err = s.GetErc20Transactions(address, startBlock, constants.UnreachableBlock)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res...)
	}
	return txs, nil
}

type InternalTxs struct {
	BlockNumber     string `json:"blockNumber"`     // tx block number
	TimeStamp       string `json:"timeStamp"`       // tx timestamp
	Hash            string `json:"hash"`            // tx hash
	From            string `json:"from"`            // who sent the chain token
	To              string `json:"to"`              // who received the chain token
	Value           string `json:"value"`           // value in WEI
	ContractAddress string `json:"contractAddress"` // contract address
	Input           string `json:"input"`           //
	Type            string `json:"type"`            //
	Gas             string `json:"gas"`             //
	GasUsed         string `json:"gasUsed"`         //
	TraceID         string `json:"traceId"`         //
	IsError         string `json:"isError"`         //
	ErrCode         string `json:"errCode"`         //
}

// Return up to 10000 internal txs of an address.
func (s *Scanner) GetInternalTransactions(address string, startBlock int64, endBlock any) ([]InternalTxs, error) {
	toBlock, err := regularcheck.RegularCheckToBlock(endBlock)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%smodule=account&action=txlistinternal&address=%s&startblock=%d&endblock=%s&sort=asc&apikey=%s", s.UrlHead, address, startBlock, toBlock, s.ApiKey)
	var res []InternalTxs
	err = s.httpGetEtherscan(url, &res)
	return res, err
}

// Return all the internal txs of an address.
func (s *Scanner) GetInternalTransactionsAll(address string) ([]InternalTxs, error) {
	res, err := s.GetInternalTransactions(address, 0, constants.UnreachableBlock)
	if err != nil {
		return nil, err
	}
	txs := res
	for len(res) == 10000 {
		lastEndBlock := res[len(res)-1].BlockNumber
		startBlock := types.ToInt64(lastEndBlock) + 1
		res, err = s.GetInternalTransactions(address, startBlock, constants.UnreachableBlock)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res...)
	}
	return txs, nil
}

type NormalTxs struct {
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

// Return up to 10000 txs of an address.
func (s *Scanner) GetNormalTransactions(address string, startBlock int64, endBlock any) ([]NormalTxs, error) {
	toBlock, err := regularcheck.RegularCheckToBlock(endBlock)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%smodule=account&action=txlist&address=%s&startblock=%d&endblock=%s&sort=asc&apikey=%s", s.UrlHead, address, startBlock, toBlock, s.ApiKey)
	var res []NormalTxs
	err = s.httpGetEtherscan(url, &res)
	return res, err
}

// Return all the txs of an address.
func (s *Scanner) GetNormalTransactionsAll(address string) ([]NormalTxs, error) {
	res, err := s.GetNormalTransactions(address, 0, constants.UnreachableBlock)
	if err != nil {
		return nil, err
	}
	txs := res
	for len(res) == 10000 {
		lastEndBlock := res[len(res)-1].BlockNumber
		startBlock := types.ToInt64(lastEndBlock) + 1
		res, err = s.GetNormalTransactions(address, startBlock, constants.UnreachableBlock)
		if err != nil {
			return nil, err
		}
		txs = append(txs, res...)
	}
	return txs, nil
}
