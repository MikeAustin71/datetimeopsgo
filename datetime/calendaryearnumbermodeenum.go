package datetime

import "sync"

type CalendarYearNumMode int

var lockCalendarYearNumMode sync.Mutex


// None - Signals that the CalendarYearNumMode Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
//
func (calYrNum CalendarYearNumMode) None() CalendarYearNumMode {

	lockCalendarYearNumMode.Lock()

	defer lockCalendarYearNumMode.Unlock()

	return CalendarYearNumMode(0)
}

// Astronomical - Signals that the year numbering system includes
// a year zero. In other words, the date January 1, year 1 is
// immediately preceded by the date December 31, year 0.
// As its name implies, Astronomical Year numbering is frequently
// used in astronomical calculations.
//
// Reference:
//      https://en.wikipedia.org/wiki/Astronomical_year_numbering
//
// This method is part of the standard enumeration.
//
func (calYrNum CalendarYearNumMode) Astronomical() CalendarYearNumMode {

	lockCalendarYearNumMode.Lock()

	defer lockCalendarYearNumMode.Unlock()

	return CalendarYearNumMode(1)
}

