package blockscan

import (
	"encoding/json"
	"errors"
)

type etherscanResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type gethResp struct {
	JsonRPC string `json:"jsonrpc"`
	ID      int64  `json:"id"`
	Result  any    `json:"result"`
}

func (s *Scanner) httpGetEtherscan(url string, result any) error {
	var resp etherscanResp
	// fmt.Println(url)
	_, err := s.reqClient.R().SetSuccessResult(&resp).Get(url)
	if err != nil {
		return err
	}
	if resp.Message != "OK" {
		return errors.New(resp.Message)
	}
	// fmt.Println(resp.Result)
	b, err := json.Marshal(resp.Result)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, result)
}

func (s *Scanner) httpGetGethProxy(url string, result any) error {
	var resp gethResp
	// fmt.Println(url)
	_, err := s.reqClient.R().SetSuccessResult(&resp).Get(url)
	if err != nil {
		return err
	}
	// fmt.Println(resp.Result)
	b, err := json.Marshal(resp.Result)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, result)
}
