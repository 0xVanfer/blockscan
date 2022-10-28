package utils

import (
	"bytes"
	"encoding/json"
)

func FormatJson(obj any) string {
	data, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return PrettifyJson(string(data))
}

func FormatJsonByte(obj any) []byte {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(data), "", "    ")
	return str.Bytes()
}

func PrettifyJson(raw string) string {
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(raw), "", "    ")
	return str.String()
}

func SimpleStringify(obj any) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func SimpleMarshal(obj any) []byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil
	}
	return bytes
}

func Stringify(obj any) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return "null"
	}
	return string(bytes)
}

func PrettyJsonPrintln(obj any) {
	println(FormatJson(obj))
}
