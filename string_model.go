package blockscan

// ------------------------------ result is string ------------------------------

type resultStringReq struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}
