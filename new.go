// Request info from blockscan.
package blockscan

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/0xVanfer/blockscan/internal/constants"
	"github.com/imroc/req/v3"
)

type Scanner struct {
	UrlHead string // Url based on network.
	ApiKey  string // User's api key, length should be 34.

	reqClient *req.Client
}

var reqDumpOpts = &req.DumpOptions{
	Output:         os.Stdout,
	RequestHeader:  false,
	ResponseBody:   false,
	RequestBody:    false,
	ResponseHeader: false,
	Async:          false,
}

// Create a new scanner.
//
// To get an apikey on ethereum, visit
// https://docs.etherscan.io/getting-started/viewing-api-usage-statistics
func New(chainID int64, apiKey string) (*Scanner, error) {
	info, exist := constants.BlockscanCnf[chainID]
	if !exist {
		return nil, errors.New("network not supported")
	}
	if len(apiKey) != 34 {
		apiKey = info.APIKey
	}

	reqClient := req.C().
		SetTimeout(time.Minute).
		SetCommonDumpOptions(reqDumpOpts).
		SetCommonRetryCount(3).
		SetCommonRetryFixedInterval(time.Second * 5).
		SetCommonRetryCondition(func(resp *req.Response, err error) bool {
			return err != nil || resp.StatusCode > http.StatusOK
		}).
		EnableDumpAll()

	return &Scanner{
		UrlHead:   info.URLHead,
		ApiKey:    apiKey,
		reqClient: reqClient,
	}, nil
}
