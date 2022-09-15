package blockscan

import (
	"errors"
	"math/big"
	"strconv"
)

func processToBlock(endBlock any) (toBlock string, err error) {
	switch v := endBlock.(type) {
	case int:
		toBlock = strconv.FormatInt(int64(v), 10)
	case int8:
		toBlock = strconv.FormatInt(int64(v), 10)
	case int16:
		toBlock = strconv.FormatInt(int64(v), 10)
	case int32:
		toBlock = strconv.FormatInt(int64(v), 10)
	case int64:
		toBlock = strconv.FormatInt(v, 10)
	case uint:
		toBlock = strconv.FormatInt(int64(v), 10)
	case uint8:
		toBlock = strconv.FormatInt(int64(v), 10)
	case uint16:
		toBlock = strconv.FormatInt(int64(v), 10)
	case uint32:
		toBlock = strconv.FormatInt(int64(v), 10)
	case uint64:
		toBlock = strconv.FormatInt(int64(v), 10)
	case *big.Int:
		toBlock = v.String()
	case string:
		if v != "latest" {
			err = errors.New(`endblock should be integer or "latest"`)
			return
		}
		toBlock = "latest"
	default:
		err = errors.New(`endblock should be integer or "latest"`)
		return
	}
	return
}
