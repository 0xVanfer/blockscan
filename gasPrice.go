package blockscan

import (
	"strconv"

	"github.com/imroc/req"
)

type BlockscanGasPriceReq struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

// If "userApiKey" is "", use default api key.
func GetGasPrice(network string, userApiKey string) (gasPrice int64, err error) {
	urlHead, apiKey, err := GetUrlAndKey(network)
	if err != nil {
		return
	}
	if userApiKey != "" {
		apiKey = userApiKey
	}
	url := urlHead + `module=proxy&action=eth_gasPrice&apikey=` + apiKey
	r, _ := req.Get(url)
	var res BlockscanGasPriceReq
	err = r.ToJSON(&res)
	if err != nil {
		return
	}
	gasPrice, err = strconv.ParseInt(res.Result[2:], 16, 64)
	return
}

// Get and print all chains' gas price by Gwei.
// If "userApiKey" is "", use default api key.
func GetGasPriceAll(networks []string, userApiKey string) (gasPriceMap map[string]int64, err error) {
	priceMap := make(map[string]int64)
	for _, network := range networks {
		gasPrice, err := GetGasPrice(network, userApiKey)
		if err != nil {
			return nil, err
		}
		priceMap[network] = gasPrice
	}
	return priceMap, nil
}
