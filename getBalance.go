package blockscan

import (
	"strconv"

	"github.com/0xVanfer/blockscan/internal/array"

	"github.com/imroc/req"
)

type BlockscanGetBalanceReq struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// Get balance of a single address.
func (s *Scanner) GetBalance(address any) (balance int, err error) {
	addressStr, err := checkAddress(address)
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
	balance, err = strconv.Atoi(res.Result)
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
func (s *Scanner) GetBalances(addresses []string) (res BlockscanGetBalancesReq, err error) {
	addressesString := array.ConnectArray(addresses, ",")
	url := s.UrlHead + `module=account&action=balancemulti&address=` + addressesString + `&tag=latest&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	err = r.ToJSON(&res)
	return
}
