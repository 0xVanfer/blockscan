package blockscan

import "errors"

type Scanner struct {
	UrlHead string
	ApiKey  string
}

// Create a new scanner.
func New(network string, apiKey string) (*Scanner, error) {
	urlHead, defaultKey, err := GetUrlAndKey(network)
	if err != nil {
		return nil, err
	}
	if apiKey == "" {
		apiKey = defaultKey
	}
	if len(apiKey) != 34 {
		err = errors.New("api key length should be 34")
		return nil, err
	}
	return &Scanner{
		UrlHead: urlHead,
		ApiKey:  apiKey,
	}, nil
}
