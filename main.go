package main

import (
	"fmt"
	"time"

	"github.com/mgmayfield/GoHolidayCalendar/holidays"
)

type Holiday struct {
	Name string
}

func IsHoliday(input time.Time) Holiday {
	// stuff to determine if holiday
	if holidays.IsChristmas(input) {
		return NewStruct("Christmas")
	} else {
		return NewStruct("")
	}
}

func NewStruct(name string) Holiday {
	return Holiday{Name: name}
}

func (s *Holiday) GetName() string {
	return s.Name
}

func main() {
	str := NewStruct("testing")
	fmt.Println(str.GetName())
}
