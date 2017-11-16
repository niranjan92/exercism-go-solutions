// Package gigasecond has utilities to check whether a person has lived for 10^9 seconds
package gigasecond

// import path for the time package from the standard library
import (
	"math"
	"time"
)

const testVersion = 4

// AddGigasecond adds 10^9 seconds to the date and returns the result
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(math.Pow10(9)))
}
