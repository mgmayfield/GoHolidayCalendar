package holidays

import "time"

// IsNewYears falls on the 1st of January, or is observed Friday or Monday
func IsNewYears(input time.Time) bool {
	if input.Month() == time.January {
		if input.Day() == 1 {
			return true
		}

		// When New Year's is a Sunday, observe Monday the 2nd
		if input.Day() == 2 && input.Weekday() == time.Monday {
			return true
		}
	}

	if input.Month() == time.December {
		// When New Year's is a Saturday, observe Friday the 31st
		if input.Day() == 31 && input.Weekday() == time.Friday {
			return true
		}
	}

	return false
}

// IsMLK falls on the 3rd Monday of January
func IsMLK(input time.Time) bool {
	if input.Month() == time.January {
		if input.Day() >= 15 && input.Day() <= 21 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}

// IsValentines falls on the 14th of February
func IsValentines(input time.Time) bool {
	if input.Month() == time.February {
		if input.Day() == 14 {
			return true
		}
	}

	return false
}

// IsPresidentsDay falls on the 3rd Monday of February
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

// IsFourthOfJuly falls on the 4th of July, or is observed Friday or Monday
func IsFourthOfJuly(input time.Time) bool {
	if input.Month() == time.July {
		if input.Day() == 4 {
			return true
		}

		// When the 4th is a Saturday, observe Friday the 3rd
		if input.Day() == 3 && input.Weekday() == time.Friday {
			return true
		}

		// When the 4th is a Sunday, observe Monday the 5th
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

// IsVeteransDay falls on the 11th of November, or is observed Friday or Monday
func IsVeteransDay(input time.Time) bool {
	if input.Month() == time.November {
		if input.Day() == 11 {
			return true
		}

		// When the 11th is a Saturday, observe Friday the 10th
		if input.Day() == 10 && input.Weekday() == time.Friday {
			return true
		}

		// When the 11th is a Sunday, observe Monday the 12th
		if input.Day() == 12 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}

// IsThanksgiving falls on the 4th Thursday of November
func IsThanksgiving(input time.Time) bool {
	if input.Month() == time.November {
		if input.Day() >= 22 && input.Day() <= 28 && int(input.Weekday()) == 4 {
			return true
		}
	}

	return false
}

// IsBlackFriday falls on the day after Thanksgiving
func IsBlackFriday(input time.Time) bool {
	if input.Month() == time.November {
		if input.Day() >= 23 && input.Day() <= 29 && input.Weekday() == time.Friday {
			return true
		}
	}

	return false
}

// IsChristmas falls on the 25th of December, or is observed Friday or Monday
func IsChristmas(input time.Time) bool {
	if input.Month() == time.December {
		if input.Day() == 25 {
			return true
		}

		// When the 25th is a Saturday, observe Friday the 24th
		if input.Day() == 24 && input.Weekday() == time.Friday {
			return true
		}

		// When the 25th is a Sunday, observe Monday the 26th
		if input.Day() == 26 && input.Weekday() == time.Monday {
			return true
		}
	}

	return false
}
