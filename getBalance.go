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
// If "userApiKey" is "", use default api key.
func GetBalance(network string, address string, userApiKey string) (balance int, err error) {
	urlHead, apiKey, err := GetUrlAndKey(network)
	if err != nil {
		return
	}
	if userApiKey != "" {
		apiKey = userApiKey
	}
	url := urlHead + `module=account&action=balance&address=` + address + `&tag=latest&apikey=` + apiKey
	r, _ := req.Get(url)
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
// If "userApiKey" is "", use default api key.
func GetBalances(network string, addresses []string, userApiKey string) (res BlockscanGetBalancesReq, err error) {
	urlHead, apiKey, err := GetUrlAndKey(network)
	if err != nil {
		return
	}
	if userApiKey != "" {
		apiKey = userApiKey
	}
	addressesString := array.ConnectArray(addresses, ",")
	url := urlHead + `module=account&action=balancemulti&address=` + addressesString + `&tag=latest&apikey=` + apiKey
	r, _ := req.Get(url)
	err = r.ToJSON(&res)
	return
}
