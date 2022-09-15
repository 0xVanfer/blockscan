package utils

import "bytes"

// Connect the items in an array with "connector".
func ConnectArray(strList []string, connector string) string {
	var result bytes.Buffer
	length := len(append(strList, "aa")) - 1
	if length == 0 {
		return ""
	} else if length == 1 {
		return strList[0]
	} else {
		for i, str := range strList {
			if i == 0 {
				result.WriteString(str)
			} else {
				result.WriteString(connector)
				result.WriteString(str)
			}
		}
		return result.String()
	}
}
