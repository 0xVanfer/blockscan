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
// If "userApiKey" is "", use default api key
func GetContractAbi(network string, address string, userApiKey string) (abi AbiStruct, err error) {
	urlHead, apiKey, err := GetUrlAndKey(network)
	if err != nil {
		return
	}
	if userApiKey != "" {
		apiKey = userApiKey
	}
	url := urlHead + `module=contract&action=getabi&address=` + address + `&apikey=` + apiKey
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
// If "userApiKey" is "", use default api key.
func GetSourceCode(network string, address string, userApiKey string) (res BlockscanSourceCodeReq, err error) {
	urlHead, apiKey, err := GetUrlAndKey(network)
	if err != nil {
		return
	}
	if userApiKey != "" {
		apiKey = userApiKey
	}
	url := urlHead + `module=contract&action=getsourcecode&address=` + address + `&apikey=` + apiKey
	r, _ := req.Get(url)
	err = r.ToJSON(&res)
	return
}

// Get the contract's name by its source code.
// If "userApiKey" is "", use default api key.
func GetContractName(network string, address string, userApiKey string) (name string, err error) {
	sourceCode, err := GetSourceCode(network, address, userApiKey)
	if err != nil {
		return
	}
	name = sourceCode.Result[0].ContractName
	return
}

// Some contracts may not be verified, will be considered not contract.
// If "userApiKey" is "", use default api key.
func IsVerifiedContract(network string, address string, userApiKey string) (isContract bool, err error) {
	sourceCode, err := GetSourceCode(network, address, userApiKey)
	if err != nil {
		return
	}
	// if verified, has contract name
	isContract = !(sourceCode.Result[0].ContractName == "")
	return
}
