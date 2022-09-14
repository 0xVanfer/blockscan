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

// Get gas price.
func (s *Scanner) GetGasPrice() (gasPrice int64, err error) {
	url := s.UrlHead + `module=proxy&action=eth_gasPrice&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	var res BlockscanGasPriceReq
	err = r.ToJSON(&res)
	if err != nil {
		return
	}
	gasPrice, err = strconv.ParseInt(res.Result[2:], 16, 64)
	return
}
