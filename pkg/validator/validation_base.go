package validator

import (
	"bytes"
	"strings"
)

type ValidationError struct {
	Message []string
}

// Implementation error interface
func (v ValidationError) Error() string {
	buff := bytes.NewBufferString("")

	for i := 0; i < len(v.Message); i++ {
		errMessage := v.Message[i]
		buff.WriteString(errMessage)
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}
