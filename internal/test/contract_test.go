package test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/0xVanfer/ethaddr"
)

// Module = contract
// https://docs.etherscan.io/api-endpoints/contracts

// action = getabi
func TestGetAbi(t *testing.T) {
	scanners := newScanners()
	wg := &sync.WaitGroup{}
	wg.Add(len(scanners))
	for chainID, scanner := range scanners {
		go func() {
			defer wg.Done()
			if ethaddr.WETHList[chainID] == "" {
				return
			}
			abi, err := scanner.GetContractAbi(ethaddr.WETHList[chainID])
			if err != nil {
				fmt.Printf("chainID = [%d]; err = [%s].\n", chainID, err.Error())
				return
			}
			fmt.Printf("chainID = [%d]; len = [%d].\n", chainID, len(abi))
		}()
	}
	wg.Wait()
}

// action = getsourcecode
func TestGetSourceCode(t *testing.T) {
	scanners := newScanners(ethaddr.ChainEthereum)
	wg := &sync.WaitGroup{}
	wg.Add(len(scanners))
	for chainID, scanner := range scanners {
		go func() {
			defer wg.Done()
			if ethaddr.WETHList[chainID] == "" {
				return
			}
			sourceCode, err := scanner.GetSourceCode(ethaddr.WETHList[chainID])
			if err != nil {
				fmt.Printf("chainID = [%d]; err = [%s].\n", chainID, err.Error())
				return
			}
			var name string
			if len(sourceCode) > 0 {
				name = sourceCode[0].ContractName
			}
			fmt.Printf("chainID = [%d]; len = [%d]; name = [%s].\n", chainID, len(sourceCode), name)
		}()
	}
	wg.Wait()
}

// action = getcontractcreation
func TestGetContractCreation(t *testing.T) {}
