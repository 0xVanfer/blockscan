package blockscan

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
	return &Scanner{
		UrlHead: urlHead,
		ApiKey:  apiKey,
	}, nil
}
