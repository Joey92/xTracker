package main

import (
	"bytes"
)

func concatString(str string, added ...string) string {
	var buffer bytes.Buffer

	buffer.WriteString(str)
	for _, strings := range added {
		buffer.WriteString(strings)
	}

	return buffer.String()
}
