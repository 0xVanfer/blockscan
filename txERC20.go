package blockscan

import (
	"github.com/0xVanfer/blockscan/internal/constants"
	"github.com/0xVanfer/blockscan/internal/regularcheck"
	"github.com/0xVanfer/types"

	"github.com/imroc/req"
)

// Return up to 10000 erc20 txs of an address.
func (s *Scanner) GetErc20Transactions(address any, startBlock int, endBlock any) ([]Erc20Txs, error) {
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
	var res getErc20TxsReq
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

// Return all the erc20 txs of an address.
func (s *Scanner) GetErc20TransactionsAll(address any) ([]Erc20Txs, error) {
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
