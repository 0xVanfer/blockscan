package utils

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

func CheckAddress(address any) (addressStr string, err error) {
	switch v := address.(type) {
	case string:
		if len(v) != 42 {
			err = errors.New(`address length should be 42, including "0x`)
			return
		}
		return v, nil
	case common.Address:
		addressStr = v.String()
		return addressStr, nil
	default:
		err = errors.New(`address type should be string or common.Address`)
		return
	}
}
