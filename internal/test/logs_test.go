package test

import (
	"fmt"
	"testing"

	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/utils"
)

// Module = logs
// https://docs.etherscan.io/api-endpoints/logs

// action = getLogs
func TestGetLogs(t *testing.T) {
	scanners := newScanners(ethaddr.ChainAvalanche)
	events, err := scanners[ethaddr.ChainAvalanche].GetEvents(
		"0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a",
		"0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c", // aave v2 lending pool
		12000000, 12000040,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	utils.PrettyJsonPrintln(events)
}
