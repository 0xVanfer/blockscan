package blockscan

import (
	"errors"
	"strconv"

	"github.com/0xVanfer/blockscan/internal/utils"
	"github.com/imroc/req"
)

type BlockscanGetBalanceReq struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// Get balance of a single address.
func (s *Scanner) GetBalance(address any) (balance int64, err error) {
	addressStr, err := utils.CheckAddress(address)
	if err != nil {
		return
	}
	url := s.UrlHead + `module=account&action=balance&address=` + addressStr + `&tag=latest&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	var res BlockscanGetBalanceReq
	err = r.ToJSON(&res)
	if err != nil {
		return
	}
	balance, err = strconv.ParseInt(res.Result, 10, 64)
	return
}

type BlockscanGetBalancesReq struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		Account string `json:"account"`
		Balance string `json:"balance"`
	} `json:"result"`
}

// Get balances of up to 20 addresses in one call.
func (s *Scanner) GetBalances(addresses []string) (balanceMap map[string]int64, err error) {
	addressesString := utils.ConnectArray(addresses, ",")
	url := s.UrlHead + `module=account&action=balancemulti&address=` + addressesString + `&tag=latest&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	var res BlockscanGetBalancesReq
	err = r.ToJSON(&res)
	if err != nil {
		return
	}
	if res.Message != "OK" {
		err = errors.New(res.Message)
		return
	}
	resMap := make(map[string]int64)
	for _, info := range res.Result {
		balance, _ := strconv.ParseInt(info.Balance, 10, 64)
		resMap[info.Account] = balance
	}
	return resMap, nil
}
