package holidays

import (
	"time"
)

// IsNewYears falls on the 1st of January
func IsNewYears(input time.Time, observed bool) bool {
	if observed {
		if input.Month() == time.January {
			if input.Day() == 1 {
				return true
			}

			if input.Day() == 2 && input.Weekday() == time.Monday {
				return true
			}
		}

		if input.Month() == time.December {
			if input.Day() == 31 && input.Weekday() == time.Friday {
				return true
			}
		}
	}

	return input.Month() == time.January && input.Day() == 1
}

// IsMLK falls on the third Monday of January
func IsMLK(input time.Time) bool {
	return input.Month() == time.January && input.Weekday() == time.Monday &&
		input.Day() >= 15 && input.Day() <= 21
}

// IsValentinesDay falls on the 14th of February
func IsValentinesDay(input time.Time) bool {
	return input.Month() == time.February && input.Day() == 14
}

// IsPresidentsDay falls on the third Monday of February
func IsPresidentsDay(input time.Time) bool {
	return input.Month() == time.February && input.Weekday() == time.Monday &&
		input.Day() >= 15 && input.Day() <= 21
}

// IsMemorialDay falls on the last Monday of May
func IsMemorialDay(input time.Time) bool {
	return input.Month() == time.May && input.Weekday() == time.Monday &&
		input.Day() >= 25 && input.Day() <= 31
}

// IsIndependenceDay falls on the 4th of July
func IsIndependenceDay(input time.Time) bool {
	return input.Month() == time.July && input.Day() == 4
}

// IsLaborDay falls on the first Monday of September
func IsLaborDay(input time.Time) bool {
	return input.Month() == time.September && input.Weekday() == time.Monday &&
		input.Day() >= 1 && input.Day() <= 7
}

// IsColumbusDay falls on the second Monday of October
func IsColumbusDay(input time.Time) bool {
	return input.Month() == time.October && input.Weekday() == time.Monday &&
		input.Day() >= 8 && input.Day() <= 14
}

// IsVeteransDay falls on the 11th of November
func IsVeteransDay(input time.Time) bool {
	return input.Month() == time.November && input.Day() == 11
}

// IsThanksgiving falls on the fourth Thursday of November
func IsThanksgiving(input time.Time) bool {
	return input.Month() == time.November && input.Weekday() == time.Thursday &&
		input.Day() >= 22 && input.Day() <= 28
}

// IsChristmas falls on the 25th of December
func IsChristmas(input time.Time) bool {
	return input.Month() == time.December && input.Day() == 25
}
