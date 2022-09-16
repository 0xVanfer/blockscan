package blockscan

import (
	"github.com/0xVanfer/blockscan/internal/utils"
	"github.com/0xVanfer/types"
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

// Return up to 1000 events of an address.
func (s *Scanner) GetEvents(topic0 string, address any, startBlock int, endBlock any) ([]BlockscanEvents, error) {
	toBlock, err := utils.CheckToBlock(endBlock)
	if err != nil {
		return nil, err
	}
	addressStr, err := utils.CheckAddress(address)
	if err != nil {
		return nil, err
	}
	url := s.UrlHead + `module=logs&action=getLogs&fromBlock=` + types.ToString(startBlock) + `&toBlock=` + toBlock + `&address=` + addressStr + `&topic0=` + topic0 + `&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var res BlockscanGetEventsReq
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

// Return all the events of an address.
func (s *Scanner) GetEventsAll(topic0 string, address string) ([]BlockscanEvents, error) {
	res, err := s.GetEvents(topic0, address, 0, "latest")
	if err != nil {
		return nil, err
	}
	events := res
	for len(res) == 1000 {
		lastBlock := res[len(res)-1].BlockNumber
		startBlock := types.ToInt(lastBlock) + 1
		res, err = s.GetEvents(topic0, address, startBlock, "latest")
		if err != nil {
			return nil, err
		}
		events = append(events, res...)
	}
	return events, nil
}
