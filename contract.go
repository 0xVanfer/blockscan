package blockscan

import (
	"encoding/json"

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

// Get the contract's abi(if it is a verified contract)
func (s *Scanner) GetContractAbi(address any) (abi AbiStruct, err error) {
	addressStr, err := checkAddress(address)
	if err != nil {
		return
	}
	url := s.UrlHead + `module=contract&action=getabi&address=` + addressStr + `&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	var abiReq BlockscanResultStringReq
	err = r.ToJSON(&abiReq)
	if err != nil {
		return
	}
	rawMassage := json.RawMessage(abiReq.Result)
	abiByte, err := json.Marshal(rawMassage)
	if err != nil {
		return
	}
	err = json.Unmarshal(abiByte, &abi)
	return
}

type BlockscanSourceCodeReq struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
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
	} `json:"result"`
}

// Get the source code of a contract.
func (s *Scanner) GetSourceCode(address any) (res BlockscanSourceCodeReq, err error) {
	addressStr, err := checkAddress(address)
	if err != nil {
		return
	}
	url := s.UrlHead + `module=contract&action=getsourcecode&address=` + addressStr + `&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return
	}
	err = r.ToJSON(&res)
	return
}

// Get the contract's name by its source code.
func (s *Scanner) GetContractName(address any) (name string, err error) {
	sourceCode, err := s.GetSourceCode(address)
	if err != nil {
		return
	}
	name = sourceCode.Result[0].ContractName
	return
}

// Some contracts may not be verified, will be considered not contract.
func (s *Scanner) IsVerifiedContract(address any) (isContract bool, err error) {
	sourceCode, err := s.GetSourceCode(address)
	if err != nil {
		return
	}
	// if verified, has contract name
	isContract = !(sourceCode.Result[0].ContractName == "")
	return
}
