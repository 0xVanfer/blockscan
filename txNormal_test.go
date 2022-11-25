package blockscan

import (
	"testing"

	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/utils"
)

func TestTxNormal(t *testing.T) {
	scanner, _ := New(chainId.AvalancheChainName, "")
	txs, _ := scanner.GetNormalTransactions(
		"0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c", // aave v2 lending pool
		11334400, 11334405,
	)
	utils.PrettyJsonPrintln(txs)
	// [
	// 	 {
	// 		"blockNumber": "11334402",
	// 		"timeStamp": "1645705330",
	// 		"hash": "0xb03ea72ef23affc306081cde182e62cdadef855ec4f6595477484631eb0d3fd0",
	// 		"nonce": "323",
	// 		"blockHash": "0x0dc09d53fa6a914b482d2d4f9117768fdcc5baf9a3d7bdee1b32f545ac7183b3",
	// 		"transactionIndex": "8",
	// 		"from": "0x27e10e6b7520858fe6a59d8dd1d36bfae8857148",
	// 		"to": "0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c",
	// 		"value": "0",
	// 		"gas": "2447798",
	// 		"gasPrice": "26500000000",
	// 		"isError": "0",
	// 		"txreceipt_status": "1",
	// 		"input":"0xab9c4b5d...00000"
	// 		"contractAddress": "",
	// 		"cumulativeGasUsed": "2526892",
	// 		"gasUsed": "1854646",
	// 		"confirmations": "10302702"
	// 	 }
	// ]
}
