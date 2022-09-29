package blockscan

import (
	"errors"

	"github.com/0xVanfer/blockscan/internal/regularcheck"
	"github.com/0xVanfer/types"
	"github.com/0xVanfer/utils"
	"github.com/imroc/req"
)

type BlockscanGetBalanceReq struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// Return the balance of a single address.
func (s *Scanner) GetBalance(address any) (int64, error) {
	addressStr, err := regularcheck.RegularCheckAddress(address)
	if err != nil {
		return 0, err
	}
	url := s.UrlHead + `module=account&action=balance&address=` + addressStr + `&tag=latest&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return 0, err
	}
	var res BlockscanGetBalanceReq
	err = r.ToJSON(&res)
	if err != nil {
		return 0, err
	}
	balance := types.ToInt64(res.Result)
	return balance, nil
}

type BlockscanGetBalancesReq struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Account string `json:"account"`
		Balance string `json:"balance"`
	} `json:"result"`
}

// Return balances of up to 20 addresses. map[address] = balance.
func (s *Scanner) GetBalances(addresses []string) (map[string]int64, error) {
	addressesString := utils.ConnectArray(addresses, ",")
	url := s.UrlHead + `module=account&action=balancemulti&address=` + addressesString + `&tag=latest&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var res BlockscanGetBalancesReq
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	if res.Message != "OK" {
		err = errors.New(res.Message)
		return nil, err
	}
	resMap := make(map[string]int64)
	for _, info := range res.Result {
		resMap[info.Account] = types.ToInt64(info.Balance)
	}
	return resMap, nil
}
