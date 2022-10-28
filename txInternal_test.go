package blockscan

import (
	"testing"

	"github.com/0xVanfer/blockscan/internal/utils"
	"github.com/0xVanfer/chainId"
)

func TestTxInternal(t *testing.T) {
	scanner, _ := New(chainId.AvalancheChainName, "")
	txs, _ := scanner.GetInternalTransactions(
		"0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c", // aave v2 lending pool
		4600000, 4700000,
	)
	utils.PrettyJsonPrintln(txs)
	// [
	// 	 {
	// 		"blockNumber": "4607005",
	// 		"timeStamp": "1632143980",
	// 		"hash": "0x5db8b8c3026d4a433ca67cbc120540ab6f8897b3aff37e78ba014ac505d167bc",
	// 		"from": "0xb6a86025f0fe1862b372cb0ca18ce3ede02a318f",
	// 		"to": "",
	// 		"value": "0",
	// 		"contractAddress": "0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c",
	// 		"input": "",
	// 		"type": "create",
	// 		"gas": "531504",
	// 		"gasUsed": "356381",
	// 		"traceId": "0_1",
	// 		"isError": "0",
	// 		"errCode": ""
	// 	 }
	// ]
}
