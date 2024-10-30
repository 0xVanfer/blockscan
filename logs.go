package blockscan

import (
	"fmt"

	"github.com/0xVanfer/blockscan/internal/regularcheck"
	"github.com/0xVanfer/types"
)

type Events struct {
	Address          string   `json:"address"`          // address
	Topics           []string `json:"topics"`           // topics of this event
	Data             string   `json:"data"`             // call data of the event
	BlockNumber      string   `json:"blockNumber"`      // block number
	TimeStamp        string   `json:"timeStamp"`        // timestamp of the event
	GasPrice         string   `json:"gasPrice"`         // gas price when the event emitted
	GasUsed          string   `json:"gasUsed"`          // gas used
	LogIndex         string   `json:"logIndex"`         // log index
	TransactionHash  string   `json:"transactionHash"`  // tx hash
	TransactionIndex string   `json:"transactionIndex"` // tx index
}

// Return up to 1000 events of an address.
//
// Param:
//
//	topic0:     The topic0 of the event. In most cases means the hash of the function name.
//	address:    The address. Can be string or common.Address
//	startBlock: The block to start from.
//	endBlock:   The block to end. If greater than latest block number, or use "latest", will use the latest block number.
func (s *Scanner) GetEvents(topic0 string, address string, startBlock int64, endBlock any) ([]Events, error) {
	toBlock, err := regularcheck.RegularCheckToBlock(endBlock)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%smodule=logs&action=getLogs&fromBlock=%d&toBlock=%s&address=%s&topic0=%s&apikey=%s", s.UrlHead, startBlock, toBlock, address, topic0, s.ApiKey)
	var res []Events
	err = s.httpGetEtherscan(url, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Return all the events of an address.
//
// Param:
//
//	topic0:     The topic0 of the event. In most cases means the hash of the function name.
//	address:    The address. Can be string or common.Address
func (s *Scanner) GetEventsAll(topic0 string, address string) ([]Events, error) {
	res, err := s.GetEvents(topic0, address, 0, "latest")
	if err != nil {
		return nil, err
	}
	events := res
	for len(res) == 1000 {
		lastBlock := res[len(res)-1].BlockNumber
		startBlock := types.ToInt64(lastBlock) + 1
		res, err = s.GetEvents(topic0, address, startBlock, "latest")
		if err != nil {
			return nil, err
		}
		events = append(events, res...)
	}
	return events, nil
}
