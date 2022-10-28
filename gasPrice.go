package blockscan

import (
	"github.com/0xVanfer/types"
	"github.com/imroc/req"
)

// Return gas price.
func (s *Scanner) GetGasPrice() (int64, error) {
	url := s.UrlHead + `module=proxy&action=eth_gasPrice&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return 0, err
	}
	var res gasPriceReq
	err = r.ToJSON(&res)
	if err != nil {
		return 0, err
	}
	gasPrice := types.ToInt64(res.Result)
	return gasPrice, nil
}
