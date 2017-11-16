// Package clock provides utility for time manipulation
package clock

import (
	"fmt"
)

const minInDay = 60 * 24
const testVersion = 4

// Clock keeps tracks of time
type Clock struct {
	min int // only save minutes in a day
}

// normalizeMinutes takes minutes as input and returns minutes since the start of the day as output
func normalizeMinutes(minutes int) int {
	// as the result has to be positive, first get mod value for minutes in day then
	// as the first result can be negative, add minutes in day and again take mod value
	return ((minutes % minInDay) + minInDay) % minInDay
}

// New takes hour and minute and creates a new clock. Supports negative hours and minutes as well
func New(hour, minute int) Clock {
	return Clock{
		min: normalizeMinutes(hour*60 + minute),
	}
}

// GetHours return hours since start of the day
func (c Clock) GetHours() int {
	return c.min / 60
}

// GetMinutes return minutes since start of the day
func (c Clock) GetMinutes() int {
	return c.min % 60
}

// String converts clock time to string and returns it
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.GetHours(), c.GetMinutes())
}

// Add method can add and subtract minutes from clock time
func (c Clock) Add(minutes int) Clock {
	// calculate change in minutes
	min := normalizeMinutes(c.min + minutes)
	return New(min/60, min%60)
}
