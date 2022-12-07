package blockscan

import (
	"encoding/json"

	"github.com/0xVanfer/blockscan/internal/regularcheck"
	"github.com/imroc/req"
)

// Return the contract's abi.
func (s *Scanner) GetContractAbi(address any) (ContractAbi, error) {
	addressStr, err := regularcheck.RegularCheckAddress(address)
	if err != nil {
		return nil, err
	}
	url := s.UrlHead + `module=contract&action=getabi&address=` + addressStr + `&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	var abiReq resultStringReq
	err = r.ToJSON(&abiReq)
	if err != nil {
		return nil, err
	}
	rawMassage := json.RawMessage(abiReq.Result)
	abiByte, err := json.Marshal(rawMassage)
	if err != nil {
		return nil, err
	}
	var abi ContractAbi
	err = json.Unmarshal(abiByte, &abi)
	return abi, err
}

// Return the source code of a contract.
func (s *Scanner) GetSourceCode(address any) (SourceCode, error) {
	var res sourceCodeReq
	addressStr, err := regularcheck.RegularCheckAddress(address)
	if err != nil {
		return SourceCode{}, err
	}
	url := s.UrlHead + `module=contract&action=getsourcecode&address=` + addressStr + `&apikey=` + s.ApiKey
	r, err := req.Get(url)
	if err != nil {
		return SourceCode{}, err
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
