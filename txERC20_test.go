package blockscan

import (
	"testing"

	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/utils"
)

func TestTxERC20(t *testing.T) {
	scanner, _ := New(chainId.AvalancheChainName, "")
	txs, _ := scanner.GetErc20Transactions(
		"0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c", // aave v2 lending pool
		5738000, 5739000,
	)
	utils.PrettyJsonPrintln(txs)
	// [
	// 	 {
	// 		"blockNumber": "5738742",
	// 		"timeStamp": "1634451346",
	// 		"hash": "0xf9a298eb9cedf27aa85e0eefcf6b64824916645f2ad7eb5096577cb493cf4932",
	// 		"nonce": "11",
	// 		"blockHash": "0x97734318edba49f4c38bcb0ea5037f28993bed9d2ef4a03aab0f79bee13e7373",
	// 		"from": "0x1f8b517e8e3cbb26125416a958dbb44c7a5387d4",
	// 		"contractAddress": "0xc7198437980c041c805a1edcba50c1ce5db95118",
	// 		"to": "0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c",
	// 		"value": "1772206585",
	// 		"tokenName": "Tether USD",
	// 		"tokenSymbol": "USDT.e",
	// 		"tokenDecimal": "6",
	// 		"transactionIndex": "1",
	// 		"gas": "78682",
	// 		"gasPrice": "25000000000",
	// 		"gasUsed": "52455",
	// 		"cumulativeGasUsed": "73455",
	// 		"input": "deprecated",
	// 		"confirmations": "15897970"
	// 	 }
	// ]
}
