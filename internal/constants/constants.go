package constants

import (
	"github.com/0xVanfer/ethaddr"
)

const (
	UnreachableBlock     int = 1e12
	UnreachableTimestamp int = 1e12
)

const (
	Vitalik             string = "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
	JustinSun           string = "0x3DdfA8eC3052539b6C9549F12cEA2C295cfF5296"
	HardhatInternalAddr string = "0x0965E063086A431771A1AE0d3f8A9d2498D04134"
)

type blockscanInfo struct {
	URLHead string `json:"urlHead"`
	APIKey  string `json:"apiKey"`
}

// Use my own api keys.
//
// Each key has a limit of 5 times/sec. Every one uses default keys will share the limit.
var BlockscanCnf = map[int64]blockscanInfo{
	ethaddr.ChainArbitrum:  {URLHead: "https://api.arbiscan.io/api?", APIKey: "BCJX4984KDQ8A3U9WWH8P2EQKN76YV8T5T"},
	ethaddr.ChainAvalanche: {URLHead: "https://api.snowscan.xyz/api?", APIKey: "K6SR1G96B2SQBRWPI4JP8WXM6BCG62EPQ7"},
	ethaddr.ChainBSC:       {URLHead: "https://api.bscscan.com/api?", APIKey: "34WQ9W88CSRHX4ZK59WNKKZYDAP5CEXINP"},
	ethaddr.ChainEthereum:  {URLHead: "https://api.etherscan.io/api?", APIKey: "RAFFHS8XP7K1FCBDFZS4TI2C9VHQBG77RP"},
	ethaddr.ChainFantom:    {URLHead: "https://api.ftmscan.com/api?", APIKey: "Q61VSDQJ7R6WQF5NJJRWDFZJJQDMW4ZR15"},
	ethaddr.ChainHeco:      {URLHead: "https://api.hecoinfo.com/api?", APIKey: "KC7PPFXMH6V187SGUDWD2XMGPCJXQWSPM5"},
	ethaddr.ChainOptimism:  {URLHead: "https://api-optimistic.etherscan.io/api?", APIKey: "53HC59E71EQ3FJS12IP8YTJGMRU59MXXJ7"},
	ethaddr.ChainPolygon:   {URLHead: "https://api.polygonscan.com/api?", APIKey: "Q31QP1KF13D6E298AJQ6X5ZHN373VDMF2S"},
	ethaddr.ChainGnosis:    {URLHead: "https://api.gnosisscan.io/api?", APIKey: "7N4F1SF35MC9ANIXQFGUTETAYI8H92X48Q"},
	ethaddr.ChainScroll:    {URLHead: "https://api.scrollscan.com/api?", APIKey: "XKZMEPQJNH9QRQ3H6K8TIBIXV1MWX29NKQ"},
	ethaddr.ChainPolygonZk: {URLHead: "https://api-zkevm.polygonscan.com/api?", APIKey: "AU75QXXFPGF4Q8PJAA1552WXK8E8AZ3ZUC"},
}
