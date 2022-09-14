package blockscan

import (
	"errors"

	"github.com/0xVanfer/blockscan/internal/constants"
)

// Read the config and get api key.
func GetUrlAndKey(network string) (urlHead string, apiKey string, err error) {
	for n, info := range constants.BlockscanCnf {
		if n == network {
			urlHead = info.URLHead
			apiKey = info.APIKey
			return
		}
	}
	err = errors.New("network not supported")
	return
}
