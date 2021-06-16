package core

import (
	"bytes"
	"fmt"
)

// 自定义错误类型
type APIError struct {
	StatusCode int
	Code       string
	Message    string
}

func (e *APIError) Error() string {
	buf := new(bytes.Buffer)
	_, _ = fmt.Fprintf(buf, "error http response:[StatusCode: %d Code: %s", e.StatusCode, e.Code)
	if e.Message != "" {
		_, _ = fmt.Fprintf(buf, "\nMessage: %s", e.Message)
	}

	return buf.String()
}
