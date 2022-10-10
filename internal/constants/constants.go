package constants

import "github.com/0xVanfer/chainId"

const UnreachableBlock int = 99999999
const UnreachableTimestamp int = 9999999999

type BlockscanInfo struct {
	URLHead string `json:"urlHead"`
	APIKey  string `json:"apiKey"`
}

// Use my own api keys.
//
// Each key has a limit of 5 times/sec. Every one uses default keys will share the limit.
var BlockscanCnf = map[string]BlockscanInfo{
	chainId.ArbitrumChainName:     {URLHead: "https://api.arbiscan.io/api?", APIKey: "BCJX4984KDQ8A3U9WWH8P2EQKN76YV8T5T"},
	chainId.AvalancheChainName:    {URLHead: "https://api.snowtrace.io/api?", APIKey: "K6SR1G96B2SQBRWPI4JP8WXM6BCG62EPQ7"},
	chainId.BinanceSmartChainName: {URLHead: "https://api.bscscan.com/api?", APIKey: "34WQ9W88CSRHX4ZK59WNKKZYDAP5CEXINP"},
	chainId.EthereumChainName:     {URLHead: "https://api.etherscan.io/api?", APIKey: "RAFFHS8XP7K1FCBDFZS4TI2C9VHQBG77RP"},
	chainId.FantomChainName:       {URLHead: "https://api.ftmscan.com/api?", APIKey: "Q61VSDQJ7R6WQF5NJJRWDFZJJQDMW4ZR15"},
	chainId.HecoChainName:         {URLHead: "https://api.hecoinfo.com/api?", APIKey: "KC7PPFXMH6V187SGUDWD2XMGPCJXQWSPM5"},
	chainId.OptimismChainName:     {URLHead: "https://api-optimistic.etherscan.io/api?", APIKey: "RAFFHS8XP7K1FCBDFZS4TI2C9VHQBG77RP"},
	chainId.PolygonChainName:      {URLHead: "https://api.polygonscan.com/api?", APIKey: "Q31QP1KF13D6E298AJQ6X5ZHN373VDMF2S"},
}
