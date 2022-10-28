package blockscan

import (
	"github.com/0xVanfer/blockscan/internal/regularcheck"
	"github.com/0xVanfer/types"
	"github.com/imroc/req"
)

// Return up to 1000 events of an address.
//
// Param:
//
//	topic0:     The topic0 of the event. In most cases means the hash of the function name.
//	address:    The address. Can be string or common.Address
//	startBlock: The block to start from.
//	endBlock:   The block to end. If greater than latest block number, or use "latest", will use the latest block number.
func (s *Scanner) GetEvents(topic0 string, address any, startBlock int, endBlock any) ([]events, error) {
	toBlock, err := regularcheck.RegularCheckToBlock(endBlock)
	if err != nil {
		return nil, err
	}
	addressStr, err := regularcheck.RegularCheckAddress(address)
	if err != nil {
		return nil, err
	}
	url := s.UrlHead + `module=logs&action=getLogs&fromBlock=` + types.ToString(startBlock) + `&toBlock=` + toBlock + `&address=` + addressStr + `&topic0=` + topic0 + `&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var res getEventsReq
	err = r.ToJSON(&res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

// Return all the events of an address.
//
// Param:
//
//	topic0:     The topic0 of the event. In most cases means the hash of the function name.
//	address:    The address. Can be string or common.Address
func (s *Scanner) GetEventsAll(topic0 string, address string) ([]events, error) {
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
