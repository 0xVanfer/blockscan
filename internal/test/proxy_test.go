package test

import (
	"fmt"
	"sync"
	"testing"
)

// Module = proxy
// https://docs.etherscan.io/api-endpoints/geth-parity-proxy

// action = eth_blockNumber
func TestETHBlockNumber(t *testing.T) {

}

// action = eth_gasPrice
func TestETHGasPrice(t *testing.T) {
	scanners := newScanners()
	wg := &sync.WaitGroup{}
	wg.Add(len(scanners))
	for chainID, scanner := range scanners {
		go func() {
			defer wg.Done()
			gas, err := scanner.GetGasPrice()
			if err != nil {
				fmt.Printf("chainID = [%d]; err = [%s].\n", chainID, err.Error())
				return
			}
			fmt.Printf("chainID = [%d]; gas = [%d].\n", chainID, gas)
		}()
	}
	wg.Wait()
}
