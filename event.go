package blockscan

import (
	"strconv"

	"github.com/0xVanfer/blockscan/internal/types"
	"github.com/imroc/req"
)

type BlockscanGetEventsReq struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Result  []BlockscanEvents `json:"result"`
}

type BlockscanEvents struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TimeStamp        string   `json:"timeStamp"`
	GasPrice         string   `json:"gasPrice"`
	GasUsed          string   `json:"gasUsed"`
	LogIndex         string   `json:"logIndex"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

// Get up to 1000 events of an address.
// If "userApiKey" is "", use default api key.
func GetEvents[T int | string](network string, topic0 string, address string, startblock int, endblock T, userApiKey string) (res BlockscanGetEventsReq, err error) {
	urlHead, apiKey, err := GetUrlAndKey(network)
	if err != nil {
		return
	}
	if userApiKey != "" {
		apiKey = userApiKey
	}
	url := urlHead + `module=logs&action=getLogs&fromBlock=` + strconv.Itoa(startblock) + `&toBlock=` + types.ToString(endblock) + `&address=` + address + `&topic0=` + topic0 + `&apikey=` + apiKey
	r, _ := req.Get(url)
	err = r.ToJSON(&res)
	return
}

// Get all the events of an address.
// If "userApiKey" is "", use default api key.
func GetEventsAll(network string, topic0 string, address string, userApiKey string) (events []BlockscanEvents, err error) {
	res, err := GetEvents(network, topic0, address, 0, "latest", userApiKey)
	if err != nil {
		return
	}
	events = append(events, res.Result...)
	for len(res.Result) == 1000 {
		lastBlock_ := res.Result[len(res.Result)-1].BlockNumber
		lastBlock, err := strconv.ParseInt(lastBlock_[2:], 16, 64)
		if err != nil {
			return nil, err
		}
		startBlock := int(lastBlock) + 1
		res, err = GetEvents(network, topic0, address, startBlock, "latest", userApiKey)
		if err != nil {
			return nil, err
		}
		events = append(events, res.Result...)

	}
	return
}
