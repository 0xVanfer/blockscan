package blockscan

type getInternalTxsReq struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []InternalTxs `json:"result"`
}

type InternalTxs struct {
	BlockNumber     string `json:"blockNumber"`     // tx block number
	TimeStamp       string `json:"timeStamp"`       // tx timestamp
	Hash            string `json:"hash"`            // tx hash
	From            string `json:"from"`            // who sent the chain token
	To              string `json:"to"`              // who received the chain token
	Value           string `json:"value"`           // value in WEI
	ContractAddress string `json:"contractAddress"` // contract address
	Input           string `json:"input"`           //
	Type            string `json:"type"`            //
	Gas             string `json:"gas"`             //
	GasUsed         string `json:"gasUsed"`         //
	TraceID         string `json:"traceId"`         //
	IsError         string `json:"isError"`         //
	ErrCode         string `json:"errCode"`         //
}
