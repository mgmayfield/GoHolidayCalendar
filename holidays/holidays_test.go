package holidays

import (
	"testing"
	"time"
)

func TestIsNewYears_NotObserved(t *testing.T) {
	newYears := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	s := IsNewYears(newYears, false)
	if s == false {
		t.Errorf("January 1st should be New Years. Was %s", newYears.String())
	}
}

func TestIsNewYears_Observed(t *testing.T) {
	newYears := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	s := IsNewYears(newYears, true)
	if s == false {
		t.Errorf("January 1st should be New Years. Was %s", newYears.String())
	}
}

func TestIsNewYears_ObservedFriday_OnFriday(t *testing.T) {
	newYears := time.Date(2021, 12, 31, 12, 0, 0, 0, time.UTC)
	s := IsNewYears(newYears, true)
	if s == false {
		t.Errorf("December 31st should be observed. Was %s", newYears.String())
	}
}

func TestIsNewYears_ObservedFriday_OnThursday(t *testing.T) {
	newYears := time.Date(2020, 12, 31, 12, 0, 0, 0, time.UTC)
	s := IsNewYears(newYears, true)
	if s == true {
		t.Errorf("December 31st should not be observed. Was %s", newYears.String())
	}
}

func TestIsNewYears_ObservedMonday_OnMonday(t *testing.T) {
	newYears := time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC)
	s := IsNewYears(newYears, true)
	if s == false {
		t.Errorf("January 2nd should be observed. Was %s", newYears.String())
	}
}

func TestIsNewYears_ObservedMonday_OnSunday(t *testing.T) {
	newYears := time.Date(2022, 1, 2, 12, 0, 0, 0, time.UTC)
	s := IsNewYears(newYears, true)
	if s == true {
		t.Errorf("January 2nd should not be observed. Was %s", newYears.String())
	}
}
