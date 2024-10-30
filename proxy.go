package blockscan

import (
	"fmt"

	"github.com/0xVanfer/types"
)

// Return gas price in WEI.
func (s *Scanner) GetGasPrice() (int64, error) {
	url := fmt.Sprintf("%smodule=proxy&action=eth_gasPrice&apikey=%s", s.UrlHead, s.ApiKey)
	var res string
	err := s.httpGetGethProxy(url, &res)
	if err != nil {
		return 0, err
	}

	return types.ToInt64(res), err
}
