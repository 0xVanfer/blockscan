package blockscan

import (
	"fmt"
	"testing"

	"github.com/0xVanfer/chainId"
)

func TestGetGasPrice(t *testing.T) {
	scanner, _ := New(chainId.AvalancheChainName, "")
	fmt.Println(scanner.GetGasPrice()) // 25000000000 <nil>
}
