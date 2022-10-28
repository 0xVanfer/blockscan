package blockscan

// ------------------------------ event ------------------------------

type getEventsReq struct {
	Status  string   `json:"status"`  // req status
	Message string   `json:"message"` // req message
	Result  []events `json:"result"`  // req result: events
}

// Everything in hex. Should convert before further use.
type events struct {
	Address          string   `json:"address"`          // address
	Topics           []string `json:"topics"`           // topics of this event
	Data             string   `json:"data"`             // call data of the event
	BlockNumber      string   `json:"blockNumber"`      // block number
	TimeStamp        string   `json:"timeStamp"`        // timestamp of the event
	GasPrice         string   `json:"gasPrice"`         // gas price when the event emitted
	GasUsed          string   `json:"gasUsed"`          // gas used
	LogIndex         string   `json:"logIndex"`         // log index
	TransactionHash  string   `json:"transactionHash"`  // tx hash
	TransactionIndex string   `json:"transactionIndex"` // tx index
}
