package holidays

import (
	"testing"
	"time"
)

func TestIsNewYears(t *testing.T) {
	newYears := time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)
	s := IsNewYears(newYears)
	if s == false {
		t.Errorf("January 1st should be New Years. Was %s", newYears.String())
	}
}
