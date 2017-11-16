// Package leap provides handly methods to evaluate leap year
package leap

const testVersion = 3

// IsLeapYear checks whether an year is a leap year or not
func IsLeapYear(year int) bool {
	if year%4 == 0 && !(year%100 == 0 && year%400 != 0) {
		return true
	}
	return false
}
