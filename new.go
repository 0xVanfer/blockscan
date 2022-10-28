// Request info from blockscan.
//
// Example:
//
//	// To get the balance of address "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7"
//	scanner, err:=blockscan.New("avalanche","")
//	if err != nil { return err }
//	res, err := scanner.GetBalance("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7")
//	if err != nil { return err }
//	fmt.Println(res)
package blockscan

import (
	"errors"
	"fmt"

	"github.com/0xVanfer/blockscan/internal/constants"
)

type Scanner struct {
	UrlHead string // Url head based on network.
	ApiKey  string // User's api key, length should be 34.
}

// Create a new scanner.
//
// "network" should be the full name of the chain,
// such as "ethereum", "avalanche", etc.
//
// "apiKey" is the key for blockscan requests.
// When the key is "", default key will be used.
// Otherwise, the key length should be 34.
//
// To get an apikey on ethereum, visit
// https://docs.etherscan.io/getting-started/viewing-api-usage-statistics
func New(network string, apiKey string) (*Scanner, error) {
	var urlHead string
	var defaultKey string
	for n, info := range constants.BlockscanCnf {
		if n == network {
			urlHead = info.URLHead
			defaultKey = info.APIKey
			break
		}
	}
	if urlHead == "" {
		return nil, errors.New("network not supported")
	}
	if apiKey == "" {
		apiKey = defaultKey
		fmt.Println("You do not have a blockscan api key. Unecpected errors may occur when running.")
	}
	if len(apiKey) != 34 {
		err := errors.New("api key length should be 34")
		return nil, err
	}
	return &Scanner{
		UrlHead: urlHead,
		ApiKey:  apiKey,
	}, nil
}
