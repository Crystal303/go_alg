package alipay

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	now := time.Now()
	t.Log(now.Format(dateFormat))
}
