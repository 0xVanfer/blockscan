## Accounts

To see more on [etherscan api](https://docs.etherscan.io/api-endpoints/accounts).

#### 1. GetBalance

Return the chain token(eth for ethereum) balance of an address.

```go
balance, _ := ethScanner.GetBalance("YOUR ADDRESS")

```

#### 2. GetBalances

Return the chain token(eth for ethereum) balances of a list of addresses.

```go
addresses:=[]string{"xxx","xxx"}
balanceMap, _ := ethScanner.GetBalances(addresses)
```

#### 3. GetErc20Transactions

Return the erc20 transactions(up to 10000) of an address between sepecific blocks.

```go
// endblock can be integer or "latest"
res, _ := ethScanner.GetErc20Transactions("0xae7ab96520de3a18e5e111b5eaab095312d7fe84", 15530000, 15536300)
fmt.Println("res:", res)
```

Output:

```bash
res: {1 OK [{15531041 1663130893 0x0f4784426ef3e0ff6573b400c9e31cf90d33fcfdb65c891b9e0eaa3201fc427b 21 0x9772413a1b27de0d6a39551f587a4d645bbe8d6e44b66696579bbbec014e50da 0x32467b144a687759ffdc1c6ffc8f9e0177d8f728 0xae7ab96520de3a18e5e111b5eaab095312d7fe84 0xae7ab96520de3a18e5e111b5eaab095312d7fe84 100000000000000000 stETH stETH 18 250 98345 15860000000 72111 26443498 deprecated 5161}]}
```

#### 4. GetErc20TransactionsAll

Return all the erc20 transactions(no amount limit) of an address from block 0 till now.

```go
res, _ := ethScanner.GetErc20TransactionsAll("0xae7ab96520de3a18e5e111b5eaab095312d7fe84")
fmt.Println("res:", res)
```

Output:

```bash
res: [{12708841 1624699236 0x001e4ff44c2dee05a4ca0bfb19e833d2efc43b3e0401c1a9f8f2f85e034af0a8 27 0xed1ca3bb9147cd70e5eb710cb81451a96b1d4d3d4bd13692448659dbec3ae228 0x6001755d9e1934c478e0de6739868bcaee6b8e44 0xae7ab96520de3a18e5e111b5eaab095312d7fe84 0xae7ab96520de3a18e5e111b5eaab095312d7fe84 1 stETH stETH 18 43 81615 15000000000 64225 2297067 deprecated 2827387} ...]
```

#### 5. GetInternalTransactions

Return the internal transactions(up to 10000) of an address between sepecific blocks.

```go
res, _ := ethScanner.GetInternalTransactions("0xae7ab96520de3a18e5e111b5eaab095312d7fe84", 15535000, 15536300)
fmt.Println("res:", res)
```

Output:

```bash
res: {1 OK [{15535230 1663191866 0x79d252a710526fa27d776b8d53439da7ae0c5f30ef30ef9b134941b683595266 0xae7ab96520de3a18e5e111b5eaab095312d7fe84 0x00000000219ab540356cbb839cbe05303d7705fa 32000000000000000000   call 13761466 23350 0_1_1_1 0 }...]}
```

#### 6. GetInternalTransactionsAll

Return all the internal transactions(no amount limit) of an address from block 0 till now.

```go
res, _ := ethScanner.GetInternalTransactionsAll("0xae7ab96520de3a18e5e111b5eaab095312d7fe84")
fmt.Println("res:", res)
```

Output:

```bash
res: [{11473216 1608242396 0x3feabd79e8549ad68d1827c074fa7123815c80206498946293d5373a160fd866 0xb8ffc3cd6e7cf5a098a1c92f48009765b24088dc  0 0xae7ab96520de3a18e5e111b5eaab095312d7fe84  create 1602466 254707 27_0_1 0 }...]
```

#### 7. GetNormalTransactions

Return the normal transactions(up to 10000) of an address between sepecific blocks.

```go
res, _ := ethScanner.GetNormalTransactions("0xae7ab96520de3a18e5e111b5eaab095312d7fe84", 15536000, 15536300)
fmt.Println("res:", res)
```

Output:

```bash
res: {1 OK [{15536024 1663203504 0x2db42b1952e54c81aa65929a2df6db7f1f9d90f389b08eaa0e001a8c253df71b 10 0xfbb2df98de835cf2d1f80cb30f168e7a9d163e2d3c107825b21502a0bb673c2c 1 0x1e0f47b58c3216fa85ddc94d137b9f16988440d0 0xae7ab96520de3a18e5e111b5eaab095312d7fe84 0 89347 17900000000 0 1 0x095ea7b30000000000000000000000001111111254fb6c44bac0bed2854e76f90643097dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff  88850 67850 555}...]}
```

#### 8. GetNormalTransactionsAll

Return all the normal transactions(no amount limit) of an address from block 0 till now.

```go
res, _ := ethScanner.GetNormalTransactionsAll("0xae7ab96520de3a18e5e111b5eaab095312d7fe84")
fmt.Println("res:", res)
```

Output:

```bash
res: {1 OK [{11480180 1608334322 0xe4f4a779647229e935928abffe882bc7a943fd36a102bfe2b0b01f66afb863cd 1 0x80f64e4f946f6249859efd14c90ac6b4bf1e985839fc11210e011dabbb1f9b36 96 0x2d07c0a3c8033af7e3eee470b15a3a7831009268 0xae7ab96520de3a18e5e111b5eaab095312d7fe84 10000000000000000 300000 28000000000 0 1 0xa1903eab0000000000000000000000000000000000000000000000000000000000000000  8394526 114496 4056419}...]}
```
