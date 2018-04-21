// Package gigasecond implements a library that manipulates time using gigaseconds (10^9 seconds)
package gigasecond

// import path for the time package from the standard library
import "time"

const gigasecond = time.Second * 1e9

// AddGigasecond returns the parameter time + one gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
