package blockscan

import (
	"encoding/json"

	"github.com/0xVanfer/blockscan/internal/regularcheck"
	"github.com/imroc/req"
)

type BlockscanResultStringReq struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type AbiStruct []struct {
	Inputs []struct {
		InternalType string `json:"internalType"`
		Name         string `json:"name"`
		Type         string `json:"type"`
	} `json:"inputs"`
	StateMutability string `json:"stateMutability"`
	Type            string `json:"type"`
	Anonymous       bool   `json:"anonymous"`
	Name            string `json:"name"`
	Outputs         []struct {
		InternalType string `json:"internalType"`
		Name         string `json:"name"`
		Type         string `json:"type"`
	} `json:"outputs"`
}

// Return the contract's abi.
func (s *Scanner) GetContractAbi(address any) (AbiStruct, error) {
	addressStr, err := regularcheck.RegularCheckAddress(address)
	if err != nil {
		return nil, err
	}
	url := s.UrlHead + `module=contract&action=getabi&address=` + addressStr + `&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var abiReq BlockscanResultStringReq
	err = r.ToJSON(&abiReq)
	if err != nil {
		return nil, err
	}
	rawMassage := json.RawMessage(abiReq.Result)
	abiByte, err := json.Marshal(rawMassage)
	if err != nil {
		return nil, err
	}
	var abi AbiStruct
	err = json.Unmarshal(abiByte, &abi)
	return abi, err
}

type BlockscanSourceCodeReq struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Result  []BlockscanSourceCode `json:"result"`
}
type BlockscanSourceCode struct {
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

// Return the source code of a contract.
func (s *Scanner) GetSourceCode(address any) (BlockscanSourceCode, error) {
	var res BlockscanSourceCodeReq
	addressStr, err := regularcheck.RegularCheckAddress(address)
	if err != nil {
		return BlockscanSourceCode{}, err
	}
	url := s.UrlHead + `module=contract&action=getsourcecode&address=` + addressStr + `&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return BlockscanSourceCode{}, err
	}
	err = r.ToJSON(&res)
	return res.Result[0], err
}

// Return the contract's name.
func (s *Scanner) GetContractName(address any) (string, error) {
	sourceCode, err := s.GetSourceCode(address)
	if err != nil {
		return "", err
	}
	name := sourceCode.ContractName
	return name, nil
}

// Return whether the address is a verified contract.
//
// Some contracts may not be verified, will be considered not contract.
func (s *Scanner) IsVerifiedContract(address any) (bool, error) {
	sourceCode, err := s.GetSourceCode(address)
	if err != nil {
		return false, err
	}
	// if verified, has contract name
	isVerified := !(sourceCode.ContractName == "")
	return isVerified, nil
}
