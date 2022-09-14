package blockscan

import (
	"strconv"

	"github.com/imroc/req"
)

// Use timestamp to get block number.
func (s *Scanner) GetBlockNumberByTimestamp(timestamp string) (blockNumber int, err error) {
	url := s.UrlHead + `module=block&action=getblocknobytime&timestamp=` + timestamp + `&closest=before&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	var res BlockscanResultStringReq
	err = r.ToJSON(&res)
	if err != nil {
		return
	}
	blockNumber, err = strconv.Atoi(res.Result)
	return
}
