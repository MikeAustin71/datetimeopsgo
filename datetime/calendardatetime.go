package datetime

import (
	"sync"
)

type CalendarDateTime struct {
	year                 int64  // Number of Years
	month                int    // Number of Months
	weeks                int    // Number of weeks since beginning of year
	weekDays             int    // Number of Week-WeekDays. Total WeekDays/7 + Remainder WeekDays
	dateDays             int    // Total Number of Days. Weeks x 7 plus WeekDays
	yearDayNumber        int    // Day number for current year
	hours                int    // Number of Hours.
	minutes              int    // Number of Minutes
	seconds              int    // Number of Seconds
	milliseconds         int    // Number of Milliseconds
	microseconds         int    // Number of Microseconds
	nanoseconds          int    // Remaining Nanoseconds after Milliseconds & Microseconds
	totSubSecNanoseconds int    // Total Nanoseconds. Millisecond NanoSecs + Microsecond NanoSecs
	//  plus remaining Nanoseconds
	totTimeNanoseconds int64  // Total Number of equivalent Nanoseconds for Hours + Minutes
	//  + Seconds + Milliseconds + Nanoseconds
	julianDayNumber        JulianDayNoDto
	timeZone TimeZoneDefinition // Contains a detailed definition and descriptions of the Time
	//                             Zone and Time Zone Location associated with this date time.
	lock          *sync.Mutex // Used for coordinating thread safe operations.
}
