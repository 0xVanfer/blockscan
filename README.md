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

To see more on [etherscan api](https://docs.etherscan.io/api-endpoints/accounts).

## [Accounts](./function_descriptions/accounts.md)

##### [1. GetBalance](./function_descriptions/accounts.md#1-getbalance)

Return the chain token(eth for ethereum) balance of an address.

##### [2. GetBalances](./function_descriptions/accounts.md#2-getbalances)

Return the chain token(eth for ethereum) balances of a list of addresses.

##### [3. GetErc20Transactions](./function_descriptions/accounts.md#3-geterc20transactions)

Return the erc20 transactions(up to 10000) of an address between sepecific blocks.

##### [4. GetErc20TransactionsAll](./function_descriptions/accounts.md#4-geterc20transactionsall)

Return all the erc20 transactions(no amount limit) of an address from block 0 till now.

##### [5. GetInternalTransactions](./function_descriptions/accounts.md#5-getinternaltransactions)

Return the internal transactions(up to 10000) of an address between sepecific blocks.

##### [6. GetInternalTransactionsAll](./function_descriptions/accounts.md#6-getinternaltransactionsall)

Return all the internal transactions(no amount limit) of an address from block 0 till now.

##### [7. GetNormalTransactions](./function_descriptions/accounts.md#7-getnormaltransactions)

Return the normal transactions(up to 10000) of an address between sepecific blocks.

##### [8. GetNormalTransactionsAll](./function_descriptions/accounts.md#8-getnormaltransactionsall)

Return all the normal transactions(no amount limit) of an address from block 0 till now.

---

## [Contracts](./function_descriptions/contracts.md)

##### [1. GetContractAbi](./function_descriptions/contracts.md#1-getcontractabi)

Return the contract abi.

##### [2. GetSourceCode](./function_descriptions/contracts.md#2-getsourcecode)

Return the contract's source code.

##### [3. GetContractName](./function_descriptions/contracts.md#3-getcontractname)

Return the contract's name.

##### [4. IsVerifiedContract](./function_descriptions/contracts.md#4-isverifiedcontract)

Return whether the address is a contract.

If the contract is not verified, will still return false.

---

## [Blocks](./function_descriptions/blocks.md)

##### [1. GetBlockNumberByTimestamp](./function_descriptions/blocks.md#1-getblocknumberbytimestamp)

Return the block number at specific timestamp.

---

## [Logs](./function_descriptions/logs.md)

##### [1. GetEvents](./function_descriptions/logs.md#1-getevents)

Return the events(up to 1000) of a topic0 between sepecific blocks.

##### [2. GetEventsAll](./function_descriptions/logs.md#2-geteventsall)

Return all the events(no amount limit) of a topic0 from block 0 till now.

---

## [Gas](./function_descriptions/gas.md)

##### [1. GetGasPrice](./function_descriptions/gas.md#1-getgasprice)

Return the gas price in wei.
