package helper

import (
	"fmt"
	"time"
)

// MicrosecondsStr 将 time.Duration 类型转为毫秒
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}
