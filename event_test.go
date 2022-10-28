package blockscan

import (
	"fmt"
	"testing"

	"github.com/0xVanfer/chainId"
	"github.com/0xVanfer/utils"
)

func TestGetEvents(t *testing.T) {
	scanner, _ := New(chainId.AvalancheChainName, "")
	events, err := scanner.GetEvents(
		"0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a",
		"0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c", // aave v2 lending pool
		12000000, 12000040,
	)
	if err != nil {
		fmt.Println(err)
	}
	utils.PrettyJsonPrintln(events)
	// Output:
	// You do not have a blockscan api key. Unecpected errors may occur when running.
	// [
	//     {
	//         "address": "0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c",
	//         "topics": [
	//             "0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a",
	//             "0x000000000000000000000000c7198437980c041c805a1edcba50c1ce5db95118"
	//         ],
	//         "data": "0x00000000000000000000000000000000000000000009e752462e65b243b73650000000000000000000000000000000000000000000279e374041684d5f543ebc00000000000000000000000000000000000000000015555ae5f6c7a0b8a87d77000000000000000000000000000000000000000003458e1e9549689b201469890000000000000000000000000000000000000000034b26a85a78ea2327fd2033",
	//         "blockNumber": "0xb71b24",
	//         "timeStamp": "0x622c2359",
	//         "gasPrice": "0x62b85e900",
	//         "gasUsed": "0x57d33",
	//         "logIndex": "0x2d",
	//         "transactionHash": "0xf6e0e44e717495136801eb00155d6f47fc56d6d7d6196916bae5c96d19b45f1c",
	//         "transactionIndex": "0xd"
	//     }
	// ]
}
