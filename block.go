package blockscan

import (
	"fmt"

	"github.com/0xVanfer/types"
)

// Return the block number.
func (s *Scanner) GetBlockNumberByTimestamp(timestamp int64) (int64, error) {
	url := fmt.Sprintf("%smodule=block&action=getblocknobytime&timestamp=%d&closest=before&apikey=%s", s.UrlHead, timestamp, s.ApiKey)
	var res string
	err := s.httpGetEtherscan(url, &res)
	if err != nil {
		return 0, err
	}
	blockNumber := types.ToInt64(res)
	return blockNumber, nil
}
