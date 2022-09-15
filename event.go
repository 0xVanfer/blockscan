package blockscan

import (
	"strconv"

	"github.com/0xVanfer/blockscan/internal/utils"
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
func (s *Scanner) GetEvents(topic0 string, address any, startBlock int, endBlock any) (res BlockscanGetEventsReq, err error) {
	toBlock, err := utils.ProcessToBlock(endBlock)
	if err != nil {
		return
	}
	addressStr, err := utils.CheckAddress(address)
	if err != nil {
		return
	}
	url := s.UrlHead + `module=logs&action=getLogs&fromBlock=` + strconv.Itoa(startBlock) + `&toBlock=` + toBlock + `&address=` + addressStr + `&topic0=` + topic0 + `&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	err = r.ToJSON(&res)
	return
}

// Get all the events of an address.
func (s *Scanner) GetEventsAll(topic0 string, address string) (events []BlockscanEvents, err error) {
	res, err := s.GetEvents(topic0, address, 0, "latest")
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
		res, err = s.GetEvents(topic0, address, startBlock, "latest")
		if err != nil {
			return nil, err
		}
		events = append(events, res.Result...)
	}
	return
}
