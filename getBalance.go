package blockscan

import (
	"errors"
	"math/big"

	"github.com/0xVanfer/blockscan/internal/regularcheck"
	"github.com/0xVanfer/types"
	"github.com/0xVanfer/utils"
	"github.com/imroc/req"
)

// Return the balance of a single address.
func (s *Scanner) GetBalance(address any) (*big.Int, error) {
	addressStr, err := regularcheck.RegularCheckAddress(address)
	if err != nil {
		return nil, err
	}
	url := s.UrlHead + `module=account&action=balance&address=` + addressStr + `&tag=latest&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var res getBalanceReq
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	balance := types.ToBigInt(res.Result)
	return balance, nil
}

// Return balances of up to 20 addresses. map[address] = balance.
func (s *Scanner) GetBalances(addresses []string) (map[string]*big.Int, error) {
	addressesString := utils.ConnectArray(addresses, ",")
	url := s.UrlHead + `module=account&action=balancemulti&address=` + addressesString + `&tag=latest&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var res getBalancesReq
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	if res.Message != "OK" {
		err = errors.New(res.Message)
		return nil, err
	}
	resMap := make(map[string]*big.Int)
	for _, info := range res.Result {
		resMap[info.Account] = types.ToBigInt(info.Balance)
	}
	return resMap, nil
}
