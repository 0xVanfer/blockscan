package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Module = block
// https://docs.etherscan.io/api-endpoints/blocks

// action = getblocknobytime
func TestGetBlockNoByTime(t *testing.T) {
	scanners := newScanners()
	wg := &sync.WaitGroup{}
	wg.Add(len(scanners))
	for chainID, scanner := range scanners {
		go func() {
			defer wg.Done()
			block, err := scanner.GetBlockNumberByTimestamp(time.Now().Unix() - 1000)
			if err != nil {
				fmt.Printf("chainID = [%d]; err = [%s].\n", chainID, err.Error())
				return
			}
			fmt.Printf("chainID = [%d]; block = [%d].\n", chainID, block)
		}()
	}
	wg.Wait()
}
