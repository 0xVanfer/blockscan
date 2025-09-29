# Blockscan

# DEPRECATED: switch to new repo using v2 api

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
| 43114    | avalanche | [snowscan](https://snowscan.xyz/)            |

---

## New a Scanner

[How to get an api key?](https://docs.etherscan.io/getting-started/viewing-api-usage-statistics)

```go
ethScanner, _ := blockscan.New(1, "YOUR API KEY")
avaxScanner, _ := blockscan.New(43114, "YOUR API KEY")
polyScanner, _ := blockscan.New(137, "YOUR API KEY")
...
```
