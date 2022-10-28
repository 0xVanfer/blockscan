package blockscan

import (
	"math/big"

	"github.com/0xVanfer/types"
	"github.com/imroc/req"
)

// Return gas price in WEI.
func (s *Scanner) GetGasPrice() (*big.Int, error) {
	url := s.UrlHead + `module=proxy&action=eth_gasPrice&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var res gasPriceReq
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	gasPrice := types.ToBigInt(res.Result)
	return gasPrice, nil
}
