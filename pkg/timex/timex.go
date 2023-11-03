package timex

import (
	"time"
)

// CurrentUnixMillis returns current timestamp in millis
func CurrentUnixMillis() int64 {
	return time.Now().UnixMilli()
}

// CurrentUnixNanos returns current timestamp in nanos
func CurrentUnixNanos() int64 {
	return time.Now().UnixNano()
}
