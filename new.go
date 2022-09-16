package blockscan

import (
	"errors"

	"github.com/0xVanfer/blockscan/internal/constants"
)

type Scanner struct {
	UrlHead string
	ApiKey  string
}

// Create a new scanner.
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
