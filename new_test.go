package blockscan

import (
	"fmt"

	"github.com/0xVanfer/chainId"
)

func ExampleNew() {
	// case 1
	fmt.Println("\ncase 1:")
	scanner, err := New(chainId.AvalancheChainName, "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*scanner)
	}
	// case 2, wrong key
	fmt.Println("\ncase 2:")
	scanner, err = New(chainId.AvalancheChainName, "asaasfas")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*scanner)
	}

	// Output:
	//
	// case 1:
	// You do not have a blockscan api key. Unecpected errors may occur when running.
	// {https://api.snowtrace.io/api? K6SR1G96B2SQBRWPI4JP8WXM6BCG62EPQ7}
	//
	// case 2:
	// api key length should be 34
}
