package blockscan

// ------------------------------ contract abi ------------------------------

type ContractAbi []struct {
	Inputs          []AbiInput  `json:"inputs"`
	StateMutability string      `json:"stateMutability"`
	Type            string      `json:"type"`
	Anonymous       bool        `json:"anonymous"`
	Name            string      `json:"name"`
	Outputs         []AbiOutput `json:"outputs"`
}

type AbiInput struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type AbiOutput struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

// ------------------------------ source code ------------------------------

type sourceCodeReq struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Result  []SourceCode `json:"result"`
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
