package helper

import (
	"time"
)

// DefaultTimezone timezone
const DefaultTimezone = "Asia/Ho_Chi_Minh"

// TimeIn tz
func TimeIn(t time.Time, tz string, fallbackTZ ...string) time.Time {
	loc, err := time.LoadLocation(tz)
	if err == nil {
		t = t.In(loc)
		return t
	}

	var fallback = DefaultTimezone
	if len(fallbackTZ) > 0 {
		fallback = fallbackTZ[0]
	}

	loc, err = time.LoadLocation(fallback)
	if err == nil {
		t = t.In(loc)
		return t
	}

	return t
}

// NowIn now
func NowIn(tz string) time.Time {
	return TimeIn(time.Now(), tz)
}
