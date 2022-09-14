package blockscan

import (
	"strconv"

	"github.com/0xVanfer/blockscan/internal/types"
	"github.com/imroc/req"
)

// Use timestamp to get block number.
// If "userApiKey" is "", use default api key.
func GetBlockNumberByTimestamp[T types.Integer | string](network string, timestamp T, userApiKey string) (blockNumber int, err error) {
	urlHead, apiKey, err := GetUrlAndKey(network)
	if err != nil {
		return 0, err
	}
	if userApiKey != "" {
		apiKey = userApiKey
	}
	url := urlHead + `module=block&action=getblocknobytime&timestamp=` + types.ToString(timestamp) + `&closest=before&apikey=` + apiKey
	r, err := req.Get(url)
	if err != nil {
		return 0, err
	}
	var res BlockscanResultStringReq
	err = r.ToJSON(&res)
	if err != nil {
		return 0, err
	}
	blockNumber, err = strconv.Atoi(res.Result)
	return
}
