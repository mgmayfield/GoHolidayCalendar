package businessdays

import "time"

// The United States has 10 business days / federal holidays
// New Years, MLK, President's Day, Memorial Day, Independence Day
// Labor Day, Columbus Day, Veterans Day, Thanksgiving, Christmas

// IsNewYears falls on the 1st of January
// It is an observed holiday on Friday Dec 31 or Monday Jan 2
func IsNewYears(input time.Time) bool {
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

	return false
}

// IsMLK falls on the third Monday of January
func IsMLK(input time.Time) bool {
	if input.Month() == time.January {
		if input.Day() >= 15 && input.Day() <= 21 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}

// IsPresidentsDay falls on the third Monday of February
func IsPresidentsDay(input time.Time) bool {
	if input.Month() == time.February {
		if input.Day() >= 15 && input.Day() <= 21 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}

// IsMemorialDay falls on the last Monday of May
func IsMemorialDay(input time.Time) bool {
	if input.Month() == time.May {
		if input.Day() >= 25 && input.Day() <= 31 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}

// IsIndependenceDay falls on the 4th of July
// It is an observed holiday on Friday July 3 or Monday July 5
func IsIndependenceDay(input time.Time) bool {
	if input.Month() == time.July {
		if input.Day() == 4 {
			return true
		}

		if input.Day() == 3 && input.Weekday() == time.Friday {
			return true
		}

		if input.Day() == 5 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}

// IsLaborDay falls on the first Monday of September
func IsLaborDay(input time.Time) bool {
	if input.Month() == time.September {
		if input.Day() >= 1 && input.Day() <= 7 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}

// IsColumbusDay falls on the second Monday of October
func IsColumbusDay(input time.Time) bool {
	if input.Month() == time.October {
		if input.Day() >= 8 && input.Day() <= 14 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}

// IsVeteransDay falls on the 11th of November
// It is an observed holiday on Friday Nov 10 or Monday Nov 12
func IsVeteransDay(input time.Time) bool {
	if input.Month() == time.November {
		if input.Day() == 11 {
			return true
		}

		if input.Day() == 10 && input.Weekday() == time.Friday {
			return true
		}

		if input.Day() == 12 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}

// IsThanksgiving falls on the fourth Thursday of November
func IsThanksgiving(input time.Time) bool {
	if input.Month() == time.November {
		if input.Day() >= 22 && input.Day() <= 28 && int(input.Weekday()) == 4 {
			return true
		}
	}

	return false
}

// IsChristmas falls on the 25th of December
// It is an observed holiday on Friday Dec 24 or Monday Dec 26
func IsChristmas(input time.Time) bool {
	if input.Month() == time.December {
		if input.Day() == 25 {
			return true
		}

		if input.Day() == 24 && input.Weekday() == time.Friday {
			return true
		}

		if input.Day() == 26 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}
