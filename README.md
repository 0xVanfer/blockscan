# Blockscan

---

#### Chains supported

| chain id | network   | blockscan                                    |
| -------- | --------- | -------------------------------------------- |
| 1        | ethereum  | [etherscan](https://etherscan.io/)           |
| 10       | optimism  | [optimism](https://optimistic.etherscan.io/) |
| 56       | binance   | [bscscan](https://bscscan.com/)              |
| 128      | heco      | [hecoscan](https://hecoinfo.com/)            |
| 137      | polygon   | [polygonscan](https://polygonscan.com/)      |
| 250      | fantom    | [ftmscan](https://ftmscan.com/)              |
| 42161    | arbitrum  | [arbiscan](https://arbiscan.io/)             |
| 43114    | avalanche | [snowtrace](https://snowtrace.io/)           |

---

## New a Scanner

```go
ethScanner, _ := blockscan.New(chainId.EthereumChainName, "YOUR API KEY")
```

---

## [Accounts](./function_descriptions/accounts.md)

To see more on [etherscan api](https://docs.etherscan.io/api-endpoints/accounts).

##### [1. GetBalance](./function_descriptions/accounts.md)

Return the chain token(eth for ethereum) balance of an address.

##### [2. GetBalances](./function_descriptions/accounts.md)

Return the chain token(eth for ethereum) balances of a list of addresses.

##### [3. GetErc20Transactions](./function_descriptions/accounts.md)

Return the erc20 transactions(up to 10000) of an address between sepecific blocks.

##### [4. GetErc20TransactionsAll](./function_descriptions/accounts.md)

Return all the erc20 transactions(no amount limit) of an address from block 0 till now.

##### [5. GetInternalTransactions](./function_descriptions/accounts.md)

Return the internal transactions(up to 10000) of an address between sepecific blocks.

##### [6. GetInternalTransactionsAll](./function_descriptions/accounts.md)

Return all the internal transactions(no amount limit) of an address from block 0 till now.

##### [7. GetNormalTransactions](./function_descriptions/accounts.md)

Return the normal transactions(up to 10000) of an address between sepecific blocks.

##### [8. GetNormalTransactionsAll](./function_descriptions/accounts.md)

Return all the normal transactions(no amount limit) of an address from block 0 till now.

---

## [Contracts](./function_descriptions/contracts.md)

##### [1. GetContractAbi](./function_descriptions/contracts.md)

Return the contract abi.

##### [2. GetSourceCode](./function_descriptions/contracts.md)

Return the contract's source code.

##### [3. GetContractName](./function_descriptions/contracts.md)

Return the contract's name.

##### [4. IsVerifiedContract](./function_descriptions/contracts.md)

Return whether the address is a contract.

If the contract is not verified, will still return false.

---

## [Blocks](./function_descriptions/blocks.md)

##### [1. GetBlockNumberByTimestamp](./function_descriptions/blocks.md)

Return the block number at specific timestamp.

---

## [Logs](./function_descriptions/logs.md)

##### [1. GetEvents](./function_descriptions/logs.md)

Return the events(up to 1000) of a topic0 between sepecific blocks.

##### [2. GetEventsAll](./function_descriptions/logs.md)

Return all the events(no amount limit) of a topic0 from block 0 till now.

---

## [Gas](./function_descriptions/gas.md)

##### [1. GetGasPrice](./function_descriptions/gas.md)

Return the gas price in wei.
