package types

import (
	"strconv"
)

func ToString(input any) string {
	var str string
	switch v := input.(type) {
	case string:
		str = v
	case int:
		str = strconv.Itoa(v)
	case int8:
		str = strconv.Itoa((int(v)))
	case int16:
		str = strconv.Itoa((int(v)))
	case int32:
		str = strconv.Itoa((int(v)))
	case int64:
		str = strconv.FormatInt(v, 10)
	case float64:
		str = strconv.FormatFloat(v, 'f', -1, 64)
	case []byte:
		str = string(v)
	}
	return str
}
