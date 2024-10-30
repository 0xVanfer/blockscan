package test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/0xVanfer/blockscan/internal/constants"
	"github.com/0xVanfer/ethaddr"
	"github.com/0xVanfer/utils"
)

// Module = account
// https://docs.etherscan.io/api-endpoints/accounts

// action = balance
func TestBalance(t *testing.T) {
	scanners := newScanners()
	wg := &sync.WaitGroup{}
	wg.Add(len(scanners))
	for chainID, scanner := range scanners {
		go func() {
			defer wg.Done()
			balance, err := scanner.GetBalance(constants.Vitalik)
			if err != nil {
				fmt.Printf("chainID = [%d]; err = [%s].\n", chainID, err.Error())
				return
			}
			fmt.Printf("chainID = [%d]; balance = [%s].\n", chainID, balance)
		}()
	}
	wg.Wait()
}

// action = balancemulti
func TestBalanceMulti(t *testing.T) {
	scanners := newScanners()
	wg := &sync.WaitGroup{}
	wg.Add(len(scanners))
	for chainID, scanner := range scanners {
		go func() {
			defer wg.Done()
			balances, err := scanner.GetBalances(constants.Vitalik, constants.JustinSun)
			if err != nil {
				fmt.Printf("chainID = [%d]; err = [%s].\n", chainID, err.Error())
				return
			}
			for owner, balance := range balances {
				fmt.Printf("chainID = [%d]; owner = [%s]; balance = [%s].\n", chainID, owner, balance)
			}
		}()
	}
	wg.Wait()
}

// action = tzlist
func TestTxList(t *testing.T) {
	scanners := newScanners(ethaddr.ChainAvalanche)
	txs, err := scanners[ethaddr.ChainAvalanche].GetNormalTransactions(
		"0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c", // aave v2 lending pool
		11334400, 11334405,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	utils.PrettyJsonPrintln(txs)
}

// action = tzlistinternal
func TestTxListInternal(t *testing.T) {
	scanners := newScanners(ethaddr.ChainAvalanche)
	txs, err := scanners[ethaddr.ChainAvalanche].GetInternalTransactions(
		"0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c", // aave v2 lending pool
		4600000, 4700000,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	utils.PrettyJsonPrintln(txs)
}

// action = tokentx
func TestTokenTx(t *testing.T) {
	scanners := newScanners(ethaddr.ChainAvalanche)
	txs, err := scanners[ethaddr.ChainAvalanche].GetErc20Transactions(
		"0x4f01aed16d97e3ab5ab2b501154dc9bb0f1a5a2c", // aave v2 lending pool
		5738000, 5739000,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	utils.PrettyJsonPrintln(txs)
}

// action = tokennfttx
func TestTokenNFTTx(t *testing.T) {

}

// action = token1155tx
func TestToken1155Tx(t *testing.T) {

}

// action = getminedblocks
func TestGetMinedBlocks(t *testing.T) {

}

// action = txsBeaconWithdrawal
func TestTxsBeaconWithdrawal(t *testing.T) {

}
