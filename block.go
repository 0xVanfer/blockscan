package blockscan

import (
	"github.com/0xVanfer/types"
	"github.com/imroc/req"
)

// Return the block number.
func (s *Scanner) GetBlockNumberByTimestamp(timestamp string) (int64, error) {
	url := s.UrlHead + `module=block&action=getblocknobytime&timestamp=` + timestamp + `&closest=before&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return 0, err
	}
	var res BlockscanResultStringReq
	err = r.ToJSON(&res)
	if err != nil {
		return 0, err
	}
	blockNumber := types.ToInt64(res.Result)
	return blockNumber, nil
}
