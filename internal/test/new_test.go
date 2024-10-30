package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/0xVanfer/blockscan"
	"github.com/0xVanfer/ethaddr"
)

func TestNew(t *testing.T) {
	scanners := newScanners()
	// Assert all News to succeed.
	assert.Equal(t, len(scanners), len(ChainIDs))
}

var ChainIDs = []int64{
	ethaddr.ChainEthereum,  // 1
	ethaddr.ChainOptimism,  // 10
	ethaddr.ChainBSC,       // 56
	ethaddr.ChainPolygon,   // 137
	ethaddr.ChainArbitrum,  // 42161
	ethaddr.ChainAvalanche, // 43114
	ethaddr.ChainScroll,    // 534352
}

func newScanners(chainIDs ...int64) map[int64]*blockscan.Scanner {
	if len(chainIDs) == 0 {
		chainIDs = ChainIDs
	}
	res := make(map[int64]*blockscan.Scanner)
	for _, chainID := range chainIDs {
		scanner, err := blockscan.New(chainID, "")
		if err != nil {
			fmt.Println(err)
			continue
		}
		res[chainID] = scanner
	}
	return res
}
