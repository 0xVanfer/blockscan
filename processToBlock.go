package blockscan

import (
	"errors"
	"strconv"
)

func processToBlock(endBlock any) (toBlock string, err error) {
	switch v := endBlock.(type) {
	case int:
		toBlock = strconv.Itoa(v)
	case string:
		if v != "latest" {
			err = errors.New(`endblock should be int or "latest"`)
			return
		}
		toBlock = "latest"
	default:
		err = errors.New(`endblock should be int or "latest"`)
		return
	}
	return
}
