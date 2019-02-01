package businessdays

import (
	"testing"
	"time"
)

func TestChristmasEveThurday(t *testing.T) {
	fridayChristmas := time.Date(2020, 12, 24, 12, 0, 0, 0, time.UTC)
	s := IsChristmas(fridayChristmas)
	if s == true {
		t.Errorf("Thursday December 24th should not be an observed holiday.")
	}
}

func TestChristmasEveFriday(t *testing.T) {
	saturdayChristmas := time.Date(2021, 12, 24, 12, 0, 0, 0, time.UTC)
	s := IsChristmas(saturdayChristmas)
	if s == false {
		t.Errorf("Friday December 24th should be an observed holiday.")
	}
}
