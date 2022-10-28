package blockscan

import (
	"fmt"
	"testing"

	"github.com/0xVanfer/chainId"
)

func TestGetBalances(t *testing.T) {
	scanner, _ := New(chainId.AvalancheChainName, "")

	res, _ := scanner.GetBalances([]string{
		"0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7",
		"0x4aefa39caeadd662ae31ab0ce7c8c2c9c0a013e8",
	})
	for addr, balance := range res {
		fmt.Println("addr:   ", addr)
		fmt.Println("balance:", balance)
	}
}
