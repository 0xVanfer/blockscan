## Contracts

##### 1. GetContractAbi

Return the contract abi.

```go
res, _ := ethScanner.GetContractAbi("0xae7ab96520de3a18e5e111b5eaab095312d7fe84")
fmt.Println("res:", res)
```

Output:

```bash
res: [{[] pure function false proxyType [{ proxyTypeId uint256}]} {[] view function false isDepositable [{  bool}]} {[] view function false implementation [{  address}]} {[] view function false appId [{  bytes32}]} {[] view function false kernel [{  address}]} {[{ _kernel address} { _appId bytes32} { _initializePayload bytes}] nonpayable constructor false  []} {[] payable fallback false  []} {[{ sender address} { value uint256}]  event false ProxyDeposit []}]
```

##### 2. GetSourceCode

Return the contract's source code.

```go
res, _ := ethScanner.GetSourceCode("0xae7ab96520de3a18e5e111b5eaab095312d7fe84")
fmt.Println("res:", res)
```

Output:

```bash
res: {1 OK [{// File: contracts/common/UnstructuredStorage.sol

/*
 * SPDX-License-Identitifer:    MIT
 */

pragma solidity ^0.4.24;

...
}]}
```

##### 3. GetContractName

Return the contract's name.

```go
res, _ := ethScanner.GetContractName("0xae7ab96520de3a18e5e111b5eaab095312d7fe84")
fmt.Println("res:", res)
```

Output:

```bash
res: AppProxyUpgradeable
```

##### 4. IsVerifiedContract

Return whether the address is a contract.

If the contract is not verified, will still return false.

```go
res, _ := ethScanner.IsVerifiedContract("0xae7ab96520de3a18e5e111b5eaab095312d7fe84")
fmt.Println("res:", res)
```

Output:

```bash
res: true
```
