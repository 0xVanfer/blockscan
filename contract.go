package blockscan

import (
	"encoding/json"
	"fmt"
)

type ContractAbi []struct {
	Name            string           `json:"name"`
	StateMutability string           `json:"stateMutability"`
	Type            string           `json:"type"`
	Anonymous       bool             `json:"anonymous"`
	Inputs          []AbiInputOutPut `json:"inputs"`
	Outputs         []AbiInputOutPut `json:"outputs"`
}

type AbiInputOutPut struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type SourceCode struct {
	SourceCode           string `json:"SourceCode"`
	Abi                  string `json:"ABI"`
	ContractName         string `json:"ContractName"`
	CompilerVersion      string `json:"CompilerVersion"`
	OptimizationUsed     string `json:"OptimizationUsed"`
	Runs                 string `json:"Runs"`
	ConstructorArguments string `json:"ConstructorArguments"`
	EVMVersion           string `json:"EVMVersion"`
	Library              string `json:"Library"`
	LicenseType          string `json:"LicenseType"`
	Proxy                string `json:"Proxy"`
	Implementation       string `json:"Implementation"`
	SwarmSource          string `json:"SwarmSource"`
}

// Return the contract's abi.
func (s *Scanner) GetContractAbi(address string) (ContractAbi, error) {
	url := fmt.Sprintf("%smodule=contract&action=getabi&address=%s&apikey=%s", s.UrlHead, address, s.ApiKey)
	var res string
	err := s.httpGetEtherscan(url, &res)
	if err != nil {
		return nil, err
	}
	rawMassage := json.RawMessage(res)
	abiByte, err := json.Marshal(rawMassage)
	if err != nil {
		return nil, err
	}
	var abi ContractAbi
	err = json.Unmarshal(abiByte, &abi)
	return abi, err
}

// Return the source code of a contract.
func (s *Scanner) GetSourceCode(address string) ([]SourceCode, error) {
	url := fmt.Sprintf("%smodule=contract&action=getsourcecode&address=%s&apikey=%s", s.UrlHead, address, s.ApiKey)
	var res []SourceCode
	err := s.httpGetEtherscan(url, &res)
	if err != nil {
		return nil, err
	}
	return res, err
}

// Return the contract's name.
func (s *Scanner) GetContractName(address string) (string, error) {
	sourceCode, err := s.GetSourceCode(address)
	if err != nil {
		return "", err
	}
	if len(sourceCode) == 0 {
		return "", nil
	}
	name := sourceCode[0].ContractName
	return name, nil
}

// Return whether the address is a verified contract.
//
// Some contracts may not be verified, will be considered not contract.
func (s *Scanner) IsVerifiedContract(address string) (bool, error) {
	sourceCode, err := s.GetSourceCode(address)
	if err != nil {
		return false, err
	}
	return len(sourceCode) > 0 && sourceCode[0].ContractName != "", nil
}
